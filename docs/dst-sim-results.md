# DST Sim Results

This document summarizes the LC-facing Phase 1/Phase 2 DST numbers. The raw exploratory matrix is still available at [sources/wowsims-wowhead-dstbench.csv](../sources/wowsims-wowhead-dstbench.csv), but the tables below match the final guild-facing assumptions in [FINAL-REPORT.md](../FINAL-REPORT.md).

Source outputs:

- [sources/final-dst-spreads.csv](../sources/final-dst-spreads.csv)
- [sources/final-p2-mutual-exclusion-ladder.csv](../sources/final-p2-mutual-exclusion-ladder.csv)

## Final Assumptions

- Alliance raid.
- Full physical raid/group buffs.
- No Draenei hit aura.
- Five-minute single-target neutral-mob sim target.
- Personal DPS only; raid utility is not credited.
- Warriors are assumed to have **no Badge of the Swarmguard** and **no Slayer's Crest**.
- Phase 2 treats **Dragonspine Trophy and Tsunami Talisman as mutually exclusive LC awards**.

## Phase 1 DST Spread

No Tsunami exists in Phase 1.

| Rank | Spec | With DST | Without DST | DST value |
|---:|---|---|---|---:|
| 1 | Survival Hunter | DST + Bloodlust Brooch | Abacus + Bloodlust Brooch | +48.32 DPS |
| 2 | Combat Rogue | DST + Bloodlust Brooch | Abacus + Bloodlust Brooch | +47.15 DPS |
| 3 | Ret Paladin | DST + Scrolls of Blinding Light | Bloodlust Brooch + Scrolls | +43.39 DPS |
| 4 | Enhancement Shaman | DST + Bloodlust Brooch | Abacus + Bloodlust Brooch | +35.21 DPS |
| 5 | Fury Warrior | DST + Bloodlust Brooch | Abacus + Bloodlust Brooch | +34.98 DPS |
| 6 | BM Hunter | DST + Bloodlust Brooch | Hourglass + Bloodlust Brooch | +29.02 DPS |
| 7 | Arms Warrior | DST + Bloodlust Brooch | Abacus + Bloodlust Brooch | +26.36 DPS |
| 8 | Feral Cat | DST + Bloodlust Brooch | Hourglass + Bloodlust Brooch | +8.15 DPS |

## Phase 2 DST Spread

This table treats Tsunami as the no-DST alternative where it is a reasonable alternative. For Ret and Arms, the actual best no-DST option is Abacus, not Tsunami.

| Rank | Spec | With DST | Without DST | DST value |
|---:|---|---|---|---:|
| 1 | Ret Paladin | DST + Bloodlust Brooch | Abacus + Bloodlust Brooch | +54.63 DPS |
| 2 | Combat Rogue | DST + Warp-Spring Coil | Tsunami + Warp-Spring Coil | +44.01 DPS |
| 3 | Enhancement Shaman | DST + Bloodlust Brooch | Tsunami + Bloodlust Brooch | +33.94 DPS |
| 4 | Fury Warrior | DST + Bloodlust Brooch | Tsunami + Bloodlust Brooch | +33.93 DPS |
| 5 | Survival Hunter | DST + Bloodlust Brooch | Tsunami + Bloodlust Brooch | +31.44 DPS |
| 6 | Arms Warrior | DST + Bloodlust Brooch | Tsunami + Bloodlust Brooch | +29.75 DPS |
| 7 | BM Hunter | DST + Bloodlust Brooch | Tsunami + Bloodlust Brooch | +22.62 DPS |
| 8 | Feral Cat | DST + Bloodlust Brooch | Tsunami + Bloodlust Brooch | +5.62 DPS |

## Phase 2 Tsunami Step-Down Ladder

This is the key table for the mutual-exclusion rule. It shows the DST package, the Tsunami package, and the next fallback if the player gets neither.

| Spec | DST package | Tsunami package | Next fallback | DST over Tsunami | Tsunami over next |
|---|---|---|---|---:|---:|
| Combat Rogue | DST + Warp-Spring Coil | Tsunami + Warp-Spring Coil | Bloodlust Brooch + Warp-Spring Coil | +44.02 | +2.36 |
| Survival Hunter | DST + Bloodlust Brooch | Tsunami + Bloodlust Brooch | Abacus + Bloodlust Brooch | +31.43 | +4.82 |
| Fury Warrior | DST + Bloodlust Brooch | Tsunami + Bloodlust Brooch | Abacus + Bloodlust Brooch | +33.93 | +4.64 |
| Enhancement Shaman | DST + Bloodlust Brooch | Tsunami + Bloodlust Brooch | Abacus + Bloodlust Brooch | +33.94 | +5.92 |
| Ret Paladin | DST + Bloodlust Brooch | Tsunami + Bloodlust Brooch | Abacus + Bloodlust Brooch | +58.18 | -3.55 |
| Arms Warrior | DST + Bloodlust Brooch | Tsunami + Bloodlust Brooch | Abacus + Bloodlust Brooch | +29.76 | -2.02 |
| BM Hunter | DST + Bloodlust Brooch | Tsunami + Bloodlust Brooch | Hourglass + Bloodlust Brooch | +22.63 | +4.86 |
| Feral Cat | DST + Bloodlust Brooch | Tsunami + Bloodlust Brooch | Empty Mug of Direbrew + Bloodlust Brooch | +5.61 | +5.47 |

Negative `Tsunami over next` means the fallback beats Tsunami.

## Cross-Spec Observations

- Combat Rogue, Survival Hunter, Fury Warrior, Ret Paladin, and Enhancement Shaman are the main practical DST pool.
- Feral Cat is the cleanest Tsunami target because DST is only a small gain over Tsunami.
- Enhancement Shaman can be placed on the Tsunami track for loot separation, but DST remains a real personal DPS upgrade.
- Ret Paladin has a very high measured personal gain, but execution and roster stability matter.
- Arms Warrior qualifies under the no-Badge assumption, but it should not receive Tsunami over Abacus.
