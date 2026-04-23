# P2 DST LC Cheat Sheet

Quick officer-night version of the Phase 1/Phase 2 Dragonspine Trophy research.
Use this alongside the full [Final LC Report](FINAL-REPORT.md).

## Fixed Assumptions

- Alliance raid.
- Full physical raid/group buffs.
- No Draenei `Heroic Presence` hit aura.
- Wowhead Phase 1 and Phase 2 BiS gear baselines.
- BM Hunter and Survival Hunter are separate.
- Fury Warrior and Arms Warrior are separate.
- Warriors are assumed to have **no Badge of the Swarmguard** and **no Slayer's Crest**.
- In Phase 2, **Dragonspine Trophy and Tsunami Talisman are mutually exclusive LC awards**.

## Fast Ruling

### Dragonspine Trophy

Default DST order:

1. **Combat Rogue**
2. **Survival Hunter**
3. **Fury Warrior**
4. **Retribution Paladin**, only if core and high-execution
5. **Enhancement Shaman**, if not being handled with Tsunami
6. **Arms Warrior**
7. **BM Hunter**
8. **Feral Cat**

### Tsunami Talisman

Best practical Tsunami targets:

1. **Feral Cat**
2. **Enhancement Shaman**
3. **Survival Hunter / BM Hunter / Fury Warrior**, only if DST is going elsewhere

Do **not** use Tsunami as a generic consolation prize. For many specs it is only a tiny gain over the next fallback, and for Ret/Arms it is worse than Abacus in this profile.

## Phase 2 Mutual-Exclusion Table

This is the main LC table. It assumes a player who gets DST is not also eligible for Tsunami, and vice versa.

| Spec | DST package | Tsunami package | Next fallback | DST over Tsunami | Tsunami over next | LC read |
|---|---|---|---|---:|---:|---|
| Combat Rogue | DST + Warp-Spring Coil | Tsunami + Warp-Spring Coil | Bloodlust Brooch + Warp-Spring Coil | +44.02 | +2.36 | DST target. Tsunami barely matters. |
| Survival Hunter | DST + Bloodlust Brooch | Tsunami + Bloodlust Brooch | Abacus + Bloodlust Brooch | +31.43 | +4.82 | Strong DST target. Tsunami is a small fallback. |
| Fury Warrior | DST + Bloodlust Brooch | Tsunami + Bloodlust Brooch | Abacus + Bloodlust Brooch | +33.93 | +4.64 | Real DST target under no-Badge assumption. |
| Enhancement Shaman | DST + Bloodlust Brooch | Tsunami + Bloodlust Brooch | Abacus + Bloodlust Brooch | +33.94 | +5.92 | Good Tsunami track target, but DST is still much better. |
| Ret Paladin | DST + Bloodlust Brooch | Tsunami + Bloodlust Brooch | Abacus + Bloodlust Brooch | +58.18 | -3.55 | DST only for core/high-execution Ret. Do not award Tsunami over Abacus. |
| Arms Warrior | DST + Bloodlust Brooch | Tsunami + Bloodlust Brooch | Abacus + Bloodlust Brooch | +29.76 | -2.02 | Valid mid-priority DST candidate. Do not award Tsunami over Abacus. |
| BM Hunter | DST + Bloodlust Brooch | Tsunami + Bloodlust Brooch | Hourglass + Bloodlust Brooch | +22.63 | +4.86 | Valid DST user, behind Survival. |
| Feral Cat | DST + Bloodlust Brooch | Tsunami + Bloodlust Brooch | Empty Mug of Direbrew + Bloodlust Brooch | +5.61 | +5.47 | Clean Tsunami assignment. Lowest DST urgency. |

Negative `Tsunami over next` means the next fallback beats Tsunami.

## Phase 2 DST Spread

This table answers: if we give this player DST, what is the measured gain over their no-DST pair?

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

Ret sims highest, but that should not automatically jump Ret over core Rogue/Survival/Fury. Ret's value depends heavily on player stability, seal twisting, uptime, and raid role.

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

## Spec Calls

### Combat Rogue

Rogue wants DST, not Tsunami. In Phase 2, Tsunami is only **+2.36 DPS** over Brooch when paired with Warp-Spring Coil, while DST is **+44.02 DPS** over Tsunami.

### Survival Hunter

Survival is a strong DST claimant and should not be merged with BM. In Phase 2, DST is **+31.43 DPS** over Tsunami.

### Fury Warrior

Fury is a real DST candidate under the no-Badge/no-Slayer's assumption. In Phase 2, DST is **+33.93 DPS** over Tsunami, while Tsunami is only **+4.64 DPS** over Abacus.

### Retribution Paladin

Ret has the largest measured DST gap, but the LC should only act on that for a core, high-execution Ret. If Ret does not get DST, use `Abacus + Bloodlust Brooch`, not Tsunami.

### Enhancement Shaman

Enhancement can reasonably be placed on the Tsunami track for loot separation, especially if the LC wants DST concentrated on Rogue/Survival/Fury. Do not describe Tsunami as equivalent: DST is still **+33.94 DPS** over Tsunami.

### Arms Warrior

Arms qualifies after stronger claims are handled. Under the no-Badge assumption, its Phase 2 DST value is meaningful, but Tsunami is worse than Abacus.

### BM Hunter

BM is a valid DST user, but Survival has the stronger measured marginal value in this profile. BM can use Tsunami as a small fallback.

### Feral Cat

Feral is the cleanest Tsunami assignment. DST is only **+5.61 DPS** over Tsunami, so awarding Feral Tsunami preserves DST for specs with much larger DST gaps.

## Practical Script

Use this when making calls live:

1. Put **Feral Cat** on Tsunami unless all high-value DST users are already covered.
2. Decide whether **Enhancement Shaman** is on Tsunami track or DST track before loot drops.
3. Award DST primarily to **Combat Rogue, Survival Hunter, Fury Warrior**, and **core/high-execution Ret**.
4. Keep **Arms Warrior** in the DST pool, especially under no-Badge assumptions.
5. Keep **BM Hunter** eligible, but behind Survival.
6. Do not award Tsunami to **Ret** or **Arms** if Abacus is available.

## Source Data

- [Final LC Report](FINAL-REPORT.md)
- [Final DST spread CSV](sources/final-dst-spreads.csv)
- [P2 DST/Tsunami ladder CSV](sources/final-p2-mutual-exclusion-ladder.csv)
- [DST methodology](docs/dst-methodology.md)
- [YouTube corroboration](docs/video-corroboration.md)
