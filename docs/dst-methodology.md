# DST Methodology

This document describes how the current Dragonspine Trophy priority numbers were produced.

## Item Mechanics

`Dragonspine Trophy` is item `28830`.

Wowhead lists it as:

- +40 Attack Power.
- Chance on melee/ranged attack to gain +325 Haste Rating for 10 seconds.
- 20 second cooldown.

WoWSims models it as:

- Static `AttackPower: 40` and `RangedAttackPower: 40`.
- A `Dragonspine Trophy Proc` temporary aura with `MeleeHaste: 325`.
- 10 second aura duration.
- 20 second internal cooldown.
- 1.0 PPM manager on melee or ranged landed hits.

At level 70, Wowhead's TBC stats overview lists 15.8 Haste Rating as 1% haste. That makes the proc roughly 20.6% haste while active before uptime and stacking behavior are considered.

## Gear Baseline

Gearsets are extracted from Wowhead P1/P2 class/spec BiS guides:

- Arms Warrior
- Fury Warrior
- Combat Rogue
- BM Hunter
- Survival Hunter
- Feral Cat
- Enhancement Shaman
- Retribution Paladin

The extractor writes:

- [sources/wowhead-bis-gearsets.json](../sources/wowhead-bis-gearsets.json)
- [sources/wowhead-bis-items.csv](../sources/wowhead-bis-items.csv)

Manual corrections are recorded in [scripts/extract-wowhead-bis.js](../scripts/extract-wowhead-bis.js). The most important correction is Enhancement Shaman P2: `Talon of the Phoenix` is main-hand only in WoWSims, so the sim baseline uses `Netherbane` as the legal offhand.

## Sim Baseline

Harness: [sources/wowsims-wowhead-dstbench.go](../sources/wowsims-wowhead-dstbench.go)

Raw exploratory output: [sources/wowsims-wowhead-dstbench.csv](../sources/wowsims-wowhead-dstbench.csv)

Final LC-facing outputs:

- [sources/final-dst-spreads.csv](../sources/final-dst-spreads.csv)
- [sources/final-p2-mutual-exclusion-ladder.csv](../sources/final-p2-mutual-exclusion-ladder.csv)

The harness:

- Loads the Wowhead gearsets.
- Replaces DST with candidate trinkets while holding the other trinket fixed.
- Runs 10,000 iterations per comparison.
- Uses the same random seed for each with-DST and without-DST pair.
- Uses a 5 minute single-target encounter.
- Sets target mob type to neutral/unknown so `Mark of the Champion` is not accidentally active.
- Uses rare-quality P1/P2 gems and normal physical DPS enchants.

The final LC tables were then normalized for the guild policy question:

- Warriors are treated as having **no Badge of the Swarmguard**.
- Warriors are treated as having **no Slayer's Crest**.
- Phase 2 assumes **Dragonspine Trophy and Tsunami Talisman are mutually exclusive LC awards**.
- The LC-facing spread asks what the player wears with DST, what they wear without DST, and what the DPS gap is under that loot rule.

This means [sources/wowsims-wowhead-dstbench.csv](../sources/wowsims-wowhead-dstbench.csv) is retained as the broad raw comparison matrix, while [sources/final-dst-spreads.csv](../sources/final-dst-spreads.csv) and [sources/final-p2-mutual-exclusion-ladder.csv](../sources/final-p2-mutual-exclusion-ladder.csv) are the officer-facing policy tables.

## Simulation Assumptions

Use these assumptions when presenting the numbers:

| Assumption | Value |
|---|---|
| Simulator | Local WoWSims TBC Go engine |
| Encounter | 5 minute single-target fight |
| Target type | Neutral/unknown mob type |
| Iterations | 10,000 per direct comparison in the retained P2 harness |
| Output metric | Personal DPS, not raid DPS |
| Raid style | Normalized full physical-DPS support environment |
| Faction | Alliance assumptions |
| Draenei hit aura | Excluded |
| Bloodlust | Enabled in party buffs |
| Drums | Drums of Battle enabled |
| Consumables | WoWSims full-consume preset per class/spec |
| Gear source | Wowhead P1/P2 BiS guide extraction plus documented manual overrides |
| Gems | Rare-quality gems selected by agility or strength profile |
| Enchants | Standard physical DPS enchants selected by slot/profile |
| Warrior trinket policy in final report | No Badge of the Swarmguard, no Slayer's Crest |
| Phase 2 LC rule | DST and Tsunami are mutually exclusive awards |
| Survival Hunter raid value | Expose Weakness raid DPS is not credited |
| Support-spec raid value | Ret/Enh/Feral utility is not credited |

The goal is not to reconstruct a legal five-player party. The goal is a consistent trinket opportunity-cost comparison across specs.

Race handling:

- Warriors, Rogue, and Ret use Human.
- Hunters and Feral use Night Elf.
- Enhancement uses a Tauren no-hit proxy because WoWSims automatically grants Draenei hit to Draenei characters. This is a sim workaround, not a Horde assumption.

Buff handling:

- Full raid buffs are enabled.
- Physical support buffs include Bloodlust, Drums, Improved Battle Shout, Improved Leader of the Pack, Grace of Air, Strength of Earth, Trueshot Aura, Ferocious Inspiration, and related physical support where applicable.
- No Draenei hit aura is included.

This is a normalized physical DPS environment, not an exact legal five-player party reconstruction. The goal is trinket opportunity cost, not full raid composition modeling.

## Commands

Regenerate Wowhead gearsets:

```bash
node scripts/extract-wowhead-bis.js
```

Run the sim matrix from the local WoWSims checkout:

```bash
export RESEARCH_REPO="$(pwd)"
export WOWSIMS_TBC="/path/to/wowsims-tbc"

mkdir -p "$WOWSIMS_TBC/cmd/dstbench-wowhead"
cp "$RESEARCH_REPO/sources/wowsims-wowhead-dstbench.go" "$WOWSIMS_TBC/cmd/dstbench-wowhead/main.go"
cd "$WOWSIMS_TBC"

DST_GEARSETS_JSON="$RESEARCH_REPO/sources/wowhead-bis-gearsets.json" \
  go run ./cmd/dstbench-wowhead > "$RESEARCH_REPO/sources/wowsims-wowhead-dstbench.csv"
```

## Caveats

- These are personal DPS deltas, not raid DPS deltas.
- Survival's `Expose Weakness` raid value is not credited as a raid gain here; DST has no Agility, so this is acceptable for item-specific personal value but not a complete raid-value model.
- `Mark of the Champion` needs a separate Demon/Undead target pass if LC wants encounter-specific swap guidance.
- Phase 3+ claims are outside this pass. The docs only assert P1/P2 retention.
- Exact raid comp, kill times, player execution, and trinket availability can move individual decisions.
