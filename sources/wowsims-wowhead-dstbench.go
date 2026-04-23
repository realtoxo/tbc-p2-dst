package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"

	_ "github.com/wowsims/tbc/sim/common"
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/items"
	"github.com/wowsims/tbc/sim/core/proto"
	feral "github.com/wowsims/tbc/sim/druid/feral"
	"github.com/wowsims/tbc/sim/hunter"
	ret "github.com/wowsims/tbc/sim/paladin/retribution"
	"github.com/wowsims/tbc/sim/rogue"
	enh "github.com/wowsims/tbc/sim/shaman/enhancement"
	warrior "github.com/wowsims/tbc/sim/warrior/dps"
	googleProto "google.golang.org/protobuf/proto"
)

const (
	DST        int32 = 28830
	Brooch     int32 = 29383
	Abacus     int32 = 28288
	Hourglass  int32 = 28034
	Tsunami    int32 = 30627
	WarpSpring int32 = 30450
	Swarmguard int32 = 21670
	Slayers    int32 = 23041
	MarkChamp  int32 = 23206
	KissSpider int32 = 22954
	Solarian   int32 = 30446
	Scrolls    int32 = 19343
	Direbrew   int32 = 38287
)

type selectedItem struct {
	Slot string `json:"slot"`
	ID   int32  `json:"id"`
	Note string `json:"note"`
}

type gearSet struct {
	Spec     string         `json:"spec"`
	Phase    string         `json:"phase"`
	URL      string         `json:"url"`
	Selected []selectedItem `json:"selected"`
}

type specCase struct {
	Key         string
	Label       string
	Phase       string
	Class       proto.Class
	Race        proto.Race
	RaceLabel   string
	GemProfile  string
	Spec        interface{}
	RaidBuffs   *proto.RaidBuffs
	PartyBuffs  *proto.PartyBuffs
	PlayerBuffs *proto.IndividualBuffs
	Consumes    *proto.Consumes
	Debuffs     *proto.Debuffs
	FixedOther  int32
	Alts        []int32
}

type resultRow struct {
	Spec       string
	Phase      string
	Race       string
	Other      int32
	Alt        int32
	WithDST    float64
	WithoutDST float64
	Delta      float64
	Pct        float64
	UptimePct  float64
	URL        string
}

func main() {
	hunter.RegisterHunter()
	rogue.RegisterRogue()
	enh.RegisterEnhancementShaman()
	ret.RegisterRetributionPaladin()
	feral.RegisterFeralDruid()
	warrior.RegisterDpsWarrior()

	gearsetsPath := os.Getenv("DST_GEARSETS_JSON")
	if gearsetsPath == "" {
		gearsetsPath = "sources/wowhead-bis-gearsets.json"
	}
	gearsets := loadGearsets(gearsetsPath)

	var rows []resultRow
	for caseIndex, c := range cases() {
		gearset, ok := gearsets[c.Key]
		if !ok {
			panic(fmt.Sprintf("missing gearset %s", c.Key))
		}
		fixedOther := c.FixedOther
		if fixedOther == 0 {
			fixedOther = chooseFixedOther(gearset, c.Alts)
		}

		alts := append([]int32{}, c.Alts...)
		for _, trinket := range selectedTrinkets(gearset) {
			alts = append(alts, trinket)
		}

		for altIndex, alt := range uniqueIDs(alts) {
			if alt == 0 || alt == DST || alt == fixedOther {
				continue
			}
			seed := int64(2026042200 + caseIndex*100 + altIndex)
			with, uptime := run(c, gearset, DST, fixedOther, seed)
			without, _ := run(c, gearset, alt, fixedOther, seed)
			rows = append(rows, resultRow{
				Spec:       c.Label,
				Phase:      c.Phase,
				Race:       c.RaceLabel,
				Other:      fixedOther,
				Alt:        alt,
				WithDST:    with,
				WithoutDST: without,
				Delta:      with - without,
				Pct:        (with - without) / without * 100,
				UptimePct:  uptime / 300 * 100,
				URL:        gearset.URL,
			})
		}
	}

	sort.SliceStable(rows, func(i, j int) bool {
		if rows[i].Phase == rows[j].Phase {
			if rows[i].Spec == rows[j].Spec {
				return rows[i].Delta > rows[j].Delta
			}
			return rows[i].Spec < rows[j].Spec
		}
		return rows[i].Phase < rows[j].Phase
	})

	w := csv.NewWriter(os.Stdout)
	_ = w.Write([]string{"spec", "phase", "race", "fixed_other_trinket", "replacement_for_dst", "with_dst_dps", "without_dst_dps", "dst_delta_dps", "dst_delta_pct", "dst_uptime_pct", "wowhead_gear_url"})
	for _, row := range rows {
		_ = w.Write([]string{
			row.Spec,
			row.Phase,
			row.Race,
			itemName(row.Other),
			itemName(row.Alt),
			fmt.Sprintf("%.2f", row.WithDST),
			fmt.Sprintf("%.2f", row.WithoutDST),
			fmt.Sprintf("%.2f", row.Delta),
			fmt.Sprintf("%.2f", row.Pct),
			fmt.Sprintf("%.2f", row.UptimePct),
			row.URL,
		})
	}
	w.Flush()
}

func cases() []specCase {
	return []specCase{
		meleeCase("fury_warrior_p1", "Fury Warrior", "p1", proto.Class_ClassWarrior, proto.Race_RaceHuman, "Human", "strength", warrior.PlayerOptionsFury, warrior.FullConsumes, warrior.FullDebuffs, Swarmguard, []int32{Brooch, Abacus, Hourglass, Slayers}),
		meleeCase("fury_warrior_p2", "Fury Warrior", "p2", proto.Class_ClassWarrior, proto.Race_RaceHuman, "Human", "strength", warrior.PlayerOptionsFury, warrior.FullConsumes, warrior.FullDebuffs, Swarmguard, []int32{Brooch, Tsunami, Abacus, Hourglass, Slayers, Solarian}),
		meleeCase("arms_warrior_p1", "Arms Warrior", "p1", proto.Class_ClassWarrior, proto.Race_RaceHuman, "Human", "strength", warrior.PlayerOptionsArmsSlam, warrior.FullConsumes, warrior.FullDebuffs, Swarmguard, []int32{Brooch, Abacus, Hourglass, Slayers}),
		meleeCase("arms_warrior_p2", "Arms Warrior", "p2", proto.Class_ClassWarrior, proto.Race_RaceHuman, "Human", "strength", warrior.PlayerOptionsArmsSlam, warrior.FullConsumes, warrior.FullDebuffs, Swarmguard, []int32{Brooch, Tsunami, Abacus, Hourglass, Slayers, Solarian}),

		meleeCase("rogue_p1", "Combat Rogue", "p1", proto.Class_ClassRogue, proto.Race_RaceHuman, "Human", "agility", rogue.PlayerOptionsBasic, rogue.FullConsumes, rogue.FullDebuffs, Brooch, []int32{Abacus, Hourglass, MarkChamp, KissSpider}),
		meleeCase("rogue_p2", "Combat Rogue", "p2", proto.Class_ClassRogue, proto.Race_RaceHuman, "Human", "agility", rogue.PlayerOptionsBasic, rogue.FullConsumes, rogue.FullDebuffs, WarpSpring, []int32{Brooch, Tsunami, Abacus, Hourglass, MarkChamp, KissSpider}),

		hunterCase("bm_hunter_p1", "BM Hunter", "p1", hunter.PlayerOptionsBasic, Brooch, []int32{Abacus, Hourglass, Swarmguard}),
		hunterCase("bm_hunter_p2", "BM Hunter", "p2", hunter.PlayerOptionsBasic, Brooch, []int32{Tsunami, Abacus, Hourglass, Swarmguard}),
		hunterCase("survival_hunter_p1", "Survival Hunter", "p1", hunter.PlayerOptionsSV, Brooch, []int32{Abacus, Hourglass, Swarmguard}),
		hunterCase("survival_hunter_p2", "Survival Hunter", "p2", hunter.PlayerOptionsSV, Brooch, []int32{Tsunami, Abacus, Hourglass, Swarmguard}),

		meleeCase("feral_cat_p1", "Feral Cat", "p1", proto.Class_ClassDruid, proto.Race_RaceNightElf, "Night Elf", "agility", feral.PlayerOptionsBiteweave, feral.FullConsumes, feral.FullDebuffs, Brooch, []int32{Abacus, Hourglass, MarkChamp, Slayers}),
		meleeCase("feral_cat_p2", "Feral Cat", "p2", proto.Class_ClassDruid, proto.Race_RaceNightElf, "Night Elf", "agility", feral.PlayerOptionsBiteweave, feral.FullConsumes, feral.FullDebuffs, Tsunami, []int32{Brooch, Abacus, Hourglass, MarkChamp, Direbrew}),

		// Tauren is used as a no-Heroic-Presence proxy; WoWSims automatically adds Draenei hit to Draenei characters.
		meleeCase("enhancement_shaman_p1", "Enhancement Shaman", "p1", proto.Class_ClassShaman, proto.Race_RaceTauren, "Alliance Shaman no-hit proxy", "strength", enh.PlayerOptionsBasic, enh.FullConsumes, enh.FullDebuffs, Brooch, []int32{Abacus, Hourglass, Slayers}),
		meleeCase("enhancement_shaman_p2", "Enhancement Shaman", "p2", proto.Class_ClassShaman, proto.Race_RaceTauren, "Alliance Shaman no-hit proxy", "strength", enh.PlayerOptionsBasic, enh.FullConsumes, enh.FullDebuffs, Brooch, []int32{Tsunami, Abacus, Hourglass, Slayers}),

		meleeCase("ret_paladin_p1", "Retribution Paladin", "p1", proto.Class_ClassPaladin, proto.Race_RaceHuman, "Human", "strength", ret.DefaultOptions, ret.FullConsumes, ret.FullDebuffs, Scrolls, []int32{Brooch, Abacus, Hourglass, MarkChamp, KissSpider}),
		meleeCase("ret_paladin_p2", "Retribution Paladin", "p2", proto.Class_ClassPaladin, proto.Race_RaceHuman, "Human", "strength", ret.DefaultOptions, ret.FullConsumes, ret.FullDebuffs, Brooch, []int32{Tsunami, Abacus, Hourglass, MarkChamp, Direbrew}),
	}
}

func meleeCase(key, label, phase string, class proto.Class, race proto.Race, raceLabel string, gemProfile string, spec interface{}, consumes *proto.Consumes, debuffs *proto.Debuffs, fixedOther int32, alts []int32) specCase {
	return specCase{
		Key:         key,
		Label:       label,
		Phase:       phase,
		Class:       class,
		Race:        race,
		RaceLabel:   raceLabel,
		GemProfile:  gemProfile,
		Spec:        spec,
		RaidBuffs:   fullRaidBuffs(),
		PartyBuffs:  fullPhysicalPartyBuffs(),
		PlayerBuffs: fullIndividualBuffs(),
		Consumes:    googleProto.Clone(consumes).(*proto.Consumes),
		Debuffs:     googleProto.Clone(debuffs).(*proto.Debuffs),
		FixedOther:  fixedOther,
		Alts:        alts,
	}
}

func hunterCase(key, label, phase string, spec interface{}, fixedOther int32, alts []int32) specCase {
	return specCase{
		Key:         key,
		Label:       label,
		Phase:       phase,
		Class:       proto.Class_ClassHunter,
		Race:        proto.Race_RaceNightElf,
		RaceLabel:   "Night Elf",
		GemProfile:  "agility",
		Spec:        spec,
		RaidBuffs:   fullRaidBuffs(),
		PartyBuffs:  fullHunterPartyBuffs(),
		PlayerBuffs: fullIndividualBuffs(),
		Consumes:    googleProto.Clone(hunter.FullConsumes).(*proto.Consumes),
		Debuffs:     googleProto.Clone(hunter.FullDebuffs).(*proto.Debuffs),
		FixedOther:  fixedOther,
		Alts:        alts,
	}
}

func fullRaidBuffs() *proto.RaidBuffs {
	return &proto.RaidBuffs{
		ArcaneBrilliance: true,
		GiftOfTheWild:    proto.TristateEffect_TristateEffectImproved,
		DivineSpirit:     proto.TristateEffect_TristateEffectImproved,
	}
}

func fullIndividualBuffs() *proto.IndividualBuffs {
	return &proto.IndividualBuffs{
		BlessingOfKings:     true,
		BlessingOfWisdom:    proto.TristateEffect_TristateEffectImproved,
		BlessingOfMight:     proto.TristateEffect_TristateEffectImproved,
		BlessingOfSalvation: true,
		UnleashedRage:       true,
	}
}

func fullPhysicalPartyBuffs() *proto.PartyBuffs {
	return &proto.PartyBuffs{
		Bloodlust:            1,
		Drums:                proto.Drums_DrumsOfBattle,
		BattleShout:          proto.TristateEffect_TristateEffectImproved,
		LeaderOfThePack:      proto.TristateEffect_TristateEffectImproved,
		GraceOfAirTotem:      proto.TristateEffect_TristateEffectImproved,
		ManaSpringTotem:      proto.TristateEffect_TristateEffectRegular,
		StrengthOfEarthTotem: proto.StrengthOfEarthType_EnhancingTotems,
		WindfuryTotemRank:    5,
		WindfuryTotemIwt:     2,
		SanctityAura:         proto.TristateEffect_TristateEffectImproved,
		TrueshotAura:         true,
		FerociousInspiration: 2,
		BraidedEterniumChain: true,
	}
}

func fullHunterPartyBuffs() *proto.PartyBuffs {
	return &proto.PartyBuffs{
		Bloodlust:            1,
		Drums:                proto.Drums_DrumsOfBattle,
		BattleShout:          proto.TristateEffect_TristateEffectImproved,
		LeaderOfThePack:      proto.TristateEffect_TristateEffectImproved,
		GraceOfAirTotem:      proto.TristateEffect_TristateEffectImproved,
		ManaSpringTotem:      proto.TristateEffect_TristateEffectRegular,
		StrengthOfEarthTotem: proto.StrengthOfEarthType_EnhancingTotems,
		TrueshotAura:         true,
		FerociousInspiration: 2,
	}
}

func loadGearsets(path string) map[string]gearSet {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	var gearsets map[string]gearSet
	if err := json.Unmarshal(data, &gearsets); err != nil {
		panic(err)
	}
	return gearsets
}

func run(c specCase, gearset gearSet, trinket1 int32, trinket2 int32, seed int64) (float64, float64) {
	gear := buildGear(c, gearset, trinket1, trinket2)
	raid := core.SinglePlayerRaidProto(
		core.WithSpec(&proto.Player{
			Race:      c.Race,
			Class:     c.Class,
			Equipment: gear,
			Consumes:  c.Consumes,
			Buffs:     c.PlayerBuffs,
		}, c.Spec),
		googleProto.Clone(c.PartyBuffs).(*proto.PartyBuffs),
		googleProto.Clone(c.RaidBuffs).(*proto.RaidBuffs),
		googleProto.Clone(c.Debuffs).(*proto.Debuffs),
	)

	encounter := core.MakeSingleTargetEncounter(5)
	encounter.Targets[0].MobType = proto.MobType_MobTypeUnknown

	result := core.RunRaidSim(&proto.RaidSimRequest{
		Raid:      raid,
		Encounter: encounter,
		SimOptions: &proto.SimOptions{
			Iterations: 10000,
			IsTest:     true,
			RandomSeed: seed,
		},
	})

	player := result.RaidMetrics.Parties[0].Players[0]
	uptime := 0.0
	for _, aura := range player.Auras {
		if aura.Id != nil && aura.Id.GetItemId() == DST {
			uptime = aura.UptimeSecondsAvg
		}
	}
	return player.Dps.Avg, uptime
}

func buildGear(c specCase, gearset gearSet, trinket1 int32, trinket2 int32) *proto.EquipmentSpec {
	var eq items.EquipmentSpec
	for _, selected := range gearset.Selected {
		slot, ok := slotFor(selected.Slot)
		if !ok {
			continue
		}
		id := selected.ID
		if selected.Slot == "trinket1" {
			id = trinket1
		}
		if selected.Slot == "trinket2" {
			id = trinket2
		}
		eq[slot] = itemSpecFor(c, slot, id)
	}
	eq[items.ItemSlotTrinket1] = itemSpecFor(c, items.ItemSlotTrinket1, trinket1)
	eq[items.ItemSlotTrinket2] = itemSpecFor(c, items.ItemSlotTrinket2, trinket2)
	equipment := items.NewEquipmentSet(eq)
	return equipment.ToEquipmentSpecProto()
}

func itemSpecFor(c specCase, slot items.ItemSlot, id int32) items.ItemSpec {
	item, ok := items.ByID[id]
	if !ok {
		panic(fmt.Sprintf("unknown item id %d in %s", id, c.Key))
	}
	return items.ItemSpec{
		ID:      id,
		Enchant: enchantFor(c, slot, item),
		Gems:    gemsFor(c.GemProfile, item),
	}
}

func slotFor(slot string) (items.ItemSlot, bool) {
	switch slot {
	case "head":
		return items.ItemSlotHead, true
	case "neck":
		return items.ItemSlotNeck, true
	case "shoulder":
		return items.ItemSlotShoulder, true
	case "back":
		return items.ItemSlotBack, true
	case "chest":
		return items.ItemSlotChest, true
	case "wrist":
		return items.ItemSlotWrist, true
	case "hands":
		return items.ItemSlotHands, true
	case "waist":
		return items.ItemSlotWaist, true
	case "legs":
		return items.ItemSlotLegs, true
	case "feet":
		return items.ItemSlotFeet, true
	case "ring1":
		return items.ItemSlotFinger1, true
	case "ring2":
		return items.ItemSlotFinger2, true
	case "trinket1":
		return items.ItemSlotTrinket1, true
	case "trinket2":
		return items.ItemSlotTrinket2, true
	case "main_hand", "two_hand":
		return items.ItemSlotMainHand, true
	case "off_hand":
		return items.ItemSlotOffHand, true
	case "ranged":
		return items.ItemSlotRanged, true
	default:
		return 0, false
	}
}

func enchantFor(c specCase, slot items.ItemSlot, item items.Item) int32 {
	switch slot {
	case items.ItemSlotHead:
		return 29192
	case items.ItemSlotShoulder:
		return 28888
	case items.ItemSlotBack:
		return 34004
	case items.ItemSlotChest:
		return 24003
	case items.ItemSlotWrist:
		if c.GemProfile == "agility" {
			return 34002
		}
		return 27899
	case items.ItemSlotHands:
		if c.GemProfile == "agility" {
			return 33152
		}
		return 33995
	case items.ItemSlotLegs:
		return 29535
	case items.ItemSlotFeet:
		return 28279
	case items.ItemSlotMainHand, items.ItemSlotOffHand:
		if c.Class == proto.Class_ClassHunter {
			if item.HandType == proto.HandType_HandTypeTwoHand {
				return 22556
			}
			return 33165
		}
		if c.Class == proto.Class_ClassDruid {
			return 22556
		}
		return 22559
	case items.ItemSlotRanged:
		if c.Class == proto.Class_ClassHunter {
			return 23766
		}
	}
	return 0
}

func gemsFor(profile string, item items.Item) []int32 {
	gems := make([]int32, 0, len(item.GemSockets))
	for _, socket := range item.GemSockets {
		switch socket {
		case proto.GemColor_GemColorMeta:
			gems = append(gems, 32409)
		case proto.GemColor_GemColorBlue:
			if profile == "agility" {
				gems = append(gems, 24055)
			} else {
				gems = append(gems, 24054)
			}
		case proto.GemColor_GemColorYellow:
			if profile == "agility" {
				gems = append(gems, 24061)
			} else {
				gems = append(gems, 24058)
			}
		default:
			if profile == "agility" {
				gems = append(gems, 24028)
			} else {
				gems = append(gems, 24027)
			}
		}
	}
	return gems
}

func selectedTrinkets(gearset gearSet) []int32 {
	var trinkets []int32
	for _, selected := range gearset.Selected {
		if strings.HasPrefix(selected.Slot, "trinket") {
			trinkets = append(trinkets, selected.ID)
		}
	}
	return trinkets
}

func chooseFixedOther(gearset gearSet, fallback []int32) int32 {
	for _, id := range selectedTrinkets(gearset) {
		if id != DST && id != MarkChamp {
			return id
		}
	}
	for _, id := range fallback {
		if id != DST && id != MarkChamp {
			return id
		}
	}
	return Brooch
}

func uniqueIDs(ids []int32) []int32 {
	seen := map[int32]bool{}
	var out []int32
	for _, id := range ids {
		if seen[id] {
			continue
		}
		seen[id] = true
		out = append(out, id)
	}
	return out
}

func itemName(id int32) string {
	if item, ok := items.ByID[id]; ok {
		return fmt.Sprintf("%s (%d)", item.Name, id)
	}
	return fmt.Sprintf("Item %d", id)
}
