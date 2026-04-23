package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"

	_ "github.com/wowsims/tbc/sim/common"
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/items"
	"github.com/wowsims/tbc/sim/core/proto"
	feral "github.com/wowsims/tbc/sim/druid/feral"
	feraltank "github.com/wowsims/tbc/sim/druid/tank"
	"github.com/wowsims/tbc/sim/hunter"
	ret "github.com/wowsims/tbc/sim/paladin/retribution"
	"github.com/wowsims/tbc/sim/rogue"
	enh "github.com/wowsims/tbc/sim/shaman/enhancement"
	warrior "github.com/wowsims/tbc/sim/warrior/dps"
	googleProto "google.golang.org/protobuf/proto"
)

const (
	DST          int32 = 28830
	Brooch      int32 = 29383
	Abacus      int32 = 28288
	Hourglass   int32 = 28034
	Tsunami     int32 = 30627
	Berserkers  int32 = 33831
	WarpSpring  int32 = 30450
	BadgeTenacy int32 = 32658
	DMCWrath    int32 = 31857
	DMCCrusade  int32 = 31856
)

type specCase struct {
	Name        string
	Class       proto.Class
	Race        proto.Race
	Gear        *proto.EquipmentSpec
	Spec        interface{}
	RaidBuffs   *proto.RaidBuffs
	PartyBuffs  *proto.PartyBuffs
	PlayerBuffs *proto.IndividualBuffs
	Consumes    *proto.Consumes
	Debuffs     *proto.Debuffs
	Other       int32
	Alts        []int32
}

type row struct {
	Spec       string
	Other      string
	Alt        string
	WithDST    float64
	WithoutDST float64
	Delta      float64
	Pct        float64
	UptimePct  float64
}

func main() {
	hunter.RegisterHunter()
	rogue.RegisterRogue()
	enh.RegisterEnhancementShaman()
	ret.RegisterRetributionPaladin()
	feral.RegisterFeralDruid()
	feraltank.RegisterFeralTankDruid()
	warrior.RegisterDpsWarrior()

	cases := []specCase{
		{
			Name:        "Fury Warrior (Orc, P1 preset)",
			Class:       proto.Class_ClassWarrior,
			Race:        proto.Race_RaceOrc,
			Gear:        warrior.FuryP1Gear,
			Spec:        warrior.PlayerOptionsFury,
			RaidBuffs:   warrior.FullRaidBuffs,
			PartyBuffs:  warrior.FullPartyBuffs,
			PlayerBuffs: warrior.FullIndividualBuffs,
			Consumes:    warrior.FullConsumes,
			Debuffs:     warrior.FullDebuffs,
			Other:       Brooch,
			Alts:        []int32{Abacus, Hourglass, Tsunami, Berserkers},
		},
		{
			Name:        "Arms Warrior (Orc, P1 preset)",
			Class:       proto.Class_ClassWarrior,
			Race:        proto.Race_RaceOrc,
			Gear:        warrior.FuryP1Gear,
			Spec:        warrior.PlayerOptionsArmsSlam,
			RaidBuffs:   warrior.FullRaidBuffs,
			PartyBuffs:  warrior.FullPartyBuffs,
			PlayerBuffs: warrior.FullIndividualBuffs,
			Consumes:    warrior.FullConsumes,
			Debuffs:     warrior.FullDebuffs,
			Other:       Brooch,
			Alts:        []int32{Abacus, Hourglass, Tsunami, Berserkers},
		},
		{
			Name:        "Combat Rogue (P1 preset)",
			Class:       proto.Class_ClassRogue,
			Race:        proto.Race_RaceHuman,
			Gear:        rogue.P1Gear,
			Spec:        rogue.PlayerOptionsBasic,
			RaidBuffs:   rogue.FullRaidBuffs,
			PartyBuffs:  rogue.FullPartyBuffs,
			PlayerBuffs: rogue.FullIndividualBuffs,
			Consumes:    rogue.FullConsumes,
			Debuffs:     rogue.FullDebuffs,
			Other:       Brooch,
			Alts:        []int32{Abacus, Hourglass, Tsunami, WarpSpring, Berserkers},
		},
		{
			Name:        "BM Hunter (P1 preset)",
			Class:       proto.Class_ClassHunter,
			Race:        proto.Race_RaceOrc,
			Gear:        hunter.P1Gear,
			Spec:        hunter.PlayerOptionsBasic,
			RaidBuffs:   hunter.FullRaidBuffs,
			PartyBuffs:  hunter.FullPartyBuffs,
			PlayerBuffs: hunter.FullIndividualBuffs,
			Consumes:    hunter.FullConsumes,
			Debuffs:     hunter.FullDebuffs,
			Other:       Brooch,
			Alts:        []int32{Abacus, Hourglass, Tsunami, Berserkers},
		},
		{
			Name:        "SV Hunter (P1 preset gear)",
			Class:       proto.Class_ClassHunter,
			Race:        proto.Race_RaceOrc,
			Gear:        hunter.P1Gear,
			Spec:        hunter.PlayerOptionsSV,
			RaidBuffs:   hunter.FullRaidBuffs,
			PartyBuffs:  hunter.FullPartyBuffs,
			PlayerBuffs: hunter.FullIndividualBuffs,
			Consumes:    hunter.FullConsumes,
			Debuffs:     hunter.FullDebuffs,
			Other:       Brooch,
			Alts:        []int32{Abacus, Hourglass, Tsunami, Berserkers},
		},
		{
			Name:        "Enhancement Shaman (Orc, P2 preset)",
			Class:       proto.Class_ClassShaman,
			Race:        proto.Race_RaceOrc,
			Gear:        enh.Phase2Gear,
			Spec:        enh.PlayerOptionsBasic,
			RaidBuffs:   enh.FullRaidBuffs,
			PartyBuffs:  enh.FullPartyBuffs,
			PlayerBuffs: enh.FullIndividualBuffs,
			Consumes:    enh.FullConsumes,
			Debuffs:     enh.FullDebuffs,
			Other:       Brooch,
			Alts:        []int32{Abacus, Hourglass, Tsunami, Berserkers},
		},
		{
			Name:        "Retribution Paladin (P4 preset)",
			Class:       proto.Class_ClassPaladin,
			Race:        proto.Race_RaceBloodElf,
			Gear:        ret.Phase4Gear,
			Spec:        ret.DefaultOptions,
			RaidBuffs:   ret.FullRaidBuffs,
			PartyBuffs:  ret.FullPartyBuffs,
			PlayerBuffs: ret.FullIndividualBuffs,
			Consumes:    ret.FullConsumes,
			Debuffs:     ret.FullDebuffs,
			Other:       Brooch,
			Alts:        []int32{Abacus, Hourglass, Tsunami, DMCCrusade, Berserkers},
		},
		{
			Name:        "Feral Cat (P1 preset)",
			Class:       proto.Class_ClassDruid,
			Race:        proto.Race_RaceTauren,
			Gear:        feral.P1Gear,
			Spec:        feral.PlayerOptionsBiteweave,
			RaidBuffs:   feral.FullRaidBuffs,
			PartyBuffs:  feral.FullPartyBuffs,
			PlayerBuffs: feral.FullIndividualBuffs,
			Consumes:    feral.FullConsumes,
			Debuffs:     feral.FullDebuffs,
			Other:       Brooch,
			Alts:        []int32{Abacus, Hourglass, Tsunami, BadgeTenacy, Berserkers},
		},
		{
			Name:        "Feral Bear threat (P1 preset)",
			Class:       proto.Class_ClassDruid,
			Race:        proto.Race_RaceTauren,
			Gear:        feraltank.P1Gear,
			Spec:        feraltank.PlayerOptionsDefault,
			RaidBuffs:   feraltank.FullRaidBuffs,
			PartyBuffs:  feraltank.FullPartyBuffs,
			PlayerBuffs: feraltank.FullIndividualBuffs,
			Consumes:    feraltank.FullConsumes,
			Debuffs:     feraltank.FullDebuffs,
			Other:       BadgeTenacy,
			Alts:        []int32{Brooch, Abacus, Hourglass, Tsunami, Berserkers},
		},
	}

	var rows []row
	for _, c := range cases {
		with, uptime := run(c, DST, c.Other)
		for _, alt := range c.Alts {
			without, _ := run(c, alt, c.Other)
			rows = append(rows, row{
				Spec:       c.Name,
				Other:      itemName(c.Other),
				Alt:        itemName(alt),
				WithDST:    with,
				WithoutDST: without,
				Delta:      with - without,
				Pct:        (with - without) / without * 100,
				UptimePct:  uptime / 300 * 100,
			})
		}
	}

	sort.SliceStable(rows, func(i, j int) bool {
		if rows[i].Spec == rows[j].Spec {
			return rows[i].Delta > rows[j].Delta
		}
		return rows[i].Spec < rows[j].Spec
	})

	w := csv.NewWriter(os.Stdout)
	_ = w.Write([]string{"spec", "fixed_other_trinket", "replacement_for_dst", "with_dst_dps", "without_dst_dps", "dst_delta_dps", "dst_delta_pct", "dst_uptime_pct"})
	for _, r := range rows {
		_ = w.Write([]string{
			r.Spec,
			r.Other,
			r.Alt,
			fmt.Sprintf("%.2f", r.WithDST),
			fmt.Sprintf("%.2f", r.WithoutDST),
			fmt.Sprintf("%.2f", r.Delta),
			fmt.Sprintf("%.2f", r.Pct),
			fmt.Sprintf("%.2f", r.UptimePct),
		})
	}
	w.Flush()
}

func run(c specCase, trinket1 int32, trinket2 int32) (float64, float64) {
	gear := googleProto.Clone(c.Gear).(*proto.EquipmentSpec)
	eq := items.ProtoToEquipment(*gear)
	eq[items.ItemSlotTrinket1] = items.ByID[trinket1]
	eq[items.ItemSlotTrinket2] = items.ByID[trinket2]
	gear = eq.ToEquipmentSpecProto()

	raid := core.SinglePlayerRaidProto(
		core.WithSpec(&proto.Player{
			Race:      c.Race,
			Class:     c.Class,
			Equipment: gear,
			Consumes:  c.Consumes,
			Buffs:     c.PlayerBuffs,
		}, c.Spec),
		c.PartyBuffs,
		c.RaidBuffs,
		c.Debuffs,
	)
	if c.Name == "Feral Bear threat (P1 preset)" {
		raid.Tanks = append(raid.Tanks, &proto.RaidTarget{TargetIndex: 0})
	}

	result := core.RunRaidSim(&proto.RaidSimRequest{
		Raid:      raid,
		Encounter: core.MakeSingleTargetEncounter(5),
		SimOptions: &proto.SimOptions{
			Iterations: 20000,
			IsTest:     true,
			RandomSeed:  20260422,
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

func itemName(id int32) string {
	if item, ok := items.ByID[id]; ok {
		return fmt.Sprintf("%s (%d)", item.Name, id)
	}
	return fmt.Sprintf("Item %d", id)
}
