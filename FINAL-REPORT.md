# Dragonspine Trophy Final LC Report

This is the officer-facing version of the DST research. It uses readable trinket pairs instead of "replacement" language.

## Assumptions

- Alliance raid.
- Full physical raid/group buffs.
- No Draenei hit aura.
- Wowhead Phase 1 and Phase 2 BiS gear baselines.
- BM Hunter and Survival Hunter are separate.
- Fury Warrior and Arms Warrior are separate.
- Warriors are assumed to have **no Badge of the Swarmguard** and **no Slayer's Crest**.
- In Phase 2, **Dragonspine Trophy and Tsunami Talisman are mutually exclusive LC awards**. If a player gets one, they are not eligible for the other.

Source CSVs:

- [Final DST spreads](sources/final-dst-spreads.csv)
- [P2 DST/Tsunami ladder](sources/final-p2-mutual-exclusion-ladder.csv)

## TLDR Priority

### Dragonspine Trophy

Default DST order under these assumptions:

1. **Combat Rogue**
2. **Survival Hunter**
3. **Fury Warrior**
4. **Retribution Paladin**, if core and high-execution
5. **Enhancement Shaman**, if not being handled with Tsunami
6. **Arms Warrior**
7. **BM Hunter**
8. **Feral Cat**

Ret sims very high, but it is more execution-sensitive and has seal-twist/haste-window caveats. Do not give it to a casual or unstable Ret over a core Rogue/Hunter/Fury.

Arms is now a real mid-priority candidate because the Warrior no-Badge assumption raises its DST value. It still does not pass Fury.

### Tsunami Talisman

Best practical Tsunami targets:

1. **Feral Cat**
2. **Enhancement Shaman**
3. **Survival Hunter / BM Hunter / Fury Warrior**, only if DST is going elsewhere

Avoid treating Tsunami as a meaningful consolation prize for everyone. For most specs it is only a small step above the next fallback.

Do **not** prioritize Tsunami to Ret or Arms as a consolation if Abacus is available. In this sim, both prefer `Abacus + Bloodlust Brooch` over `Tsunami + Bloodlust Brooch`.

## Phase 1 DST Spread

No Tsunami exists in P1. This table asks: if this player gets DST, what are they wearing; if they do not get DST, what is the fallback pair?

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

This is the key table for the mutual-exclusion rule. It shows what each spec gets if they are awarded DST, what they get if they are awarded Tsunami instead, and what the next fallback is if they get neither.

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

Negative `Tsunami over next` means the listed fallback is better than Tsunami.

## LC Interpretation

### Feral Cat

Feral is the cleanest Tsunami assignment.

Feral gains only **+5.62 DPS** by getting DST instead of Tsunami in P2. That means Tsunami almost fully satisfies the Feral trinket claim while preserving DST for specs with a much larger DST gap.

Practical ruling: **Feral gets Tsunami, not DST**, unless every high-value DST user is already covered.

### Enhancement Shaman

Enhancement is a good Tsunami recipient if the LC wants to separate DST and Tsunami pools.

Enh loses **+33.94 DPS** going from DST to Tsunami, so DST is still much better for personal DPS. But Tsunami is the best non-DST package tested, and Enh is a natural candidate if your loot policy wants Feral/Enh on Tsunami and Rogue/Hunter/Warrior on DST.

Practical ruling: **Enh can be handled with Tsunami**, but do not pretend Tsunami is equal to DST.

### Fury Warrior

With no Badge, Fury's DST value rises.

P2:

- DST + Brooch: 1583.43 DPS
- Tsunami + Brooch: 1549.50 DPS
- Abacus + Brooch: 1544.86 DPS

Fury loses **+33.93 DPS** if they get Tsunami instead of DST, while Tsunami is only **+4.64 DPS** over Abacus.

Practical ruling: **Fury is a real DST candidate. Tsunami is only a small consolation.**

### Arms Warrior

With no Badge, Arms is more competitive than the earlier spread suggested.

P2:

- DST + Brooch: 1394.65 DPS
- Tsunami + Brooch: 1364.89 DPS
- Abacus + Brooch: 1366.91 DPS

Arms loses **+29.75 DPS** if they get Tsunami instead of DST, but Abacus is slightly better than Tsunami.

Practical ruling: **Arms qualifies for DST after stronger claims are handled. Do not award Tsunami to Arms if Abacus is available.**

### Combat Rogue

Rogue's fallback after DST is poor in relative terms.

P2:

- DST + Warp-Spring Coil: 1952.96 DPS
- Tsunami + Warp-Spring Coil: 1908.94 DPS
- Bloodlust Brooch + Warp-Spring Coil: 1906.58 DPS

Tsunami is only **+2.36 DPS** over Brooch for Rogue, while DST is **+44.02 DPS** over Tsunami.

Practical ruling: **Rogue wants DST, not Tsunami.**

### Survival Hunter

Survival remains a strong DST target.

P2:

- DST + Brooch: 1867.10 DPS
- Tsunami + Brooch: 1835.67 DPS
- Abacus + Brooch: 1830.85 DPS

Tsunami is a small fallback, but DST is still **+31.43 DPS** over Tsunami.

Practical ruling: **Survival is a strong DST candidate if the player executes properly.**

### BM Hunter

BM still uses DST, but its marginal gap is smaller than Survival's.

P2:

- DST + Brooch: 1943.72 DPS
- Tsunami + Brooch: 1921.09 DPS
- Hourglass + Brooch: 1916.23 DPS

Practical ruling: **BM qualifies, but should generally wait behind Rogue, Survival, Fury, and strong Ret/Enh claims.**

### Ret Paladin

Ret has the biggest measured DST gap, but is not a Tsunami target.

P2:

- DST + Brooch: 2190.45 DPS
- Abacus + Brooch: 2135.82 DPS
- Tsunami + Brooch: 2132.27 DPS

Practical ruling: **Give DST to Ret only if the Ret is core and high-execution. If not DST, use Abacus/Brooch, not Tsunami.**

## Final Practical Award Plan

Use this if you want a simple loot-council policy:

1. Put **Feral Cat** on the Tsunami track.
2. Put **Enhancement Shaman** on the Tsunami track if you want cleaner item separation, but recognize it is a real DST loss.
3. Prioritize DST to **Combat Rogue, Survival Hunter, Fury Warrior**, and **core Ret**.
4. Treat **Arms Warrior** as a legitimate mid-priority DST candidate under the no-Badge assumption.
5. Treat **BM Hunter** as valid but behind Survival.
6. Do not spend early DST on Feral if Tsunami is available.
7. Do not spend Tsunami on Ret or Arms if Abacus is available.
