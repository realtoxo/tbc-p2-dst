# Dragonspine Trophy Loot Priority

This is the LC-facing recommendation for Phase 1 and Phase 2. It uses the final guild-facing assumptions from [FINAL-REPORT.md](../FINAL-REPORT.md), not the earlier Badge-based exploratory Warrior rows.

## Recommended Priority Pool

### First Wave

| Priority | Spec | Why |
|---:|---|---|
| 1 | Combat Rogue | Strong durable DST claimant: +47.15 DPS in P1 and +44.01 DPS in P2, with excellent DST uptime and weak Tsunami consolation value. |
| 2 | Survival Hunter | Stronger measured marginal value than BM: +48.32 DPS in P1 and +31.44 DPS in P2. Verify execution, especially haste timing and melee-weave comfort. |
| 3 | Fury Warrior | Real DST target under the final no-Badge/no-Slayer policy: +34.98 DPS in P1 and +33.93 DPS in P2. |
| 4 | Retribution Paladin | Highest measured P2 personal delta at +54.63 DPS, but should only be treated as high priority for a core, high-execution Ret. |
| 5 | Enhancement Shaman | Strong user at +35.21 DPS in P1 and +33.94 DPS in P2, though Tsunami can be used as a policy track if DST is going to Rogue/Hunter/Warrior. |

### Later Wave

| Priority | Spec | Why |
|---:|---|---|
| 6 | Arms Warrior | Qualifies under no-Badge assumptions: +26.36 DPS in P1 and +29.75 DPS in P2, but less directly corroborated than Fury. |
| 7 | BM Hunter | Valid DST user, but lower measured marginal value than Survival: +29.02 DPS in P1 and +22.62 DPS in P2. |
| 8 | Feral Cat | Lowest broad DST urgency: +8.15 DPS in P1 and +5.62 DPS over Tsunami in P2. Best handled on Tsunami track. |

## Numbers-First Ranking

| Rank | Spec | P1 delta | P2 delta | Practical read |
|---:|---|---:|---:|---|
| 1 | Combat Rogue | +47.15 | +44.01 | Premium DST claimant |
| 2 | Survival Hunter | +48.32 | +31.44 | Premium DST claimant with execution caveat |
| 3 | Fury Warrior | +34.98 | +33.93 | Premium DST claimant under no-Badge policy |
| 4 | Retribution Paladin | +43.39 | +54.63 | High value, execution-sensitive |
| 5 | Enhancement Shaman | +35.21 | +33.94 | Strong DST user; possible Tsunami track |
| 6 | Arms Warrior | +26.36 | +29.75 | Valid mid-priority claim |
| 7 | BM Hunter | +29.02 | +22.62 | Valid but behind Survival |
| 8 | Feral Cat | +8.15 | +5.62 | Prefer Tsunami |

## Spec Notes

### Combat Rogue

Rogue should qualify strongly. In P2, with `Warp-Spring Coil` paired:

- DST + Warp-Spring Coil: 1952.96 DPS
- Tsunami + Warp-Spring Coil: 1908.94 DPS
- Bloodlust Brooch + Warp-Spring Coil: 1906.58 DPS

Tsunami is only +2.36 DPS over Brooch, while DST is +44.02 DPS over Tsunami. LC takeaway: Rogue wants DST, not Tsunami.

### Survival Hunter

Survival is not BM and should not inherit BM's priority. Under this profile:

- P1: +48.32 DPS over `Abacus + Bloodlust Brooch`.
- P2: +31.44 DPS over `Tsunami + Bloodlust Brooch`.

DST does not directly increase `Expose Weakness` because it has no Agility, so this is mostly personal DPS. Still, SV's measured personal gain is first-wave quality.

### Fury Warrior

Fury should be evaluated separately from Arms:

- P1: +34.98 DPS over `Abacus + Bloodlust Brooch`.
- P2: +33.93 DPS over `Tsunami + Bloodlust Brooch`.

This depends on the final policy assumption that the Warrior does not have Badge of the Swarmguard or Slayer's Crest. LC takeaway: Fury is a real first-wave candidate when the player is core and the raid expects them to remain Fury through P2.

### Retribution Paladin

Ret looks very strong numerically:

- P1: +43.39 DPS over `Bloodlust Brooch + Scrolls of Blinding Light`.
- P2: +54.63 DPS over `Abacus + Bloodlust Brooch`.

The caveat is roster reality and execution. If the Ret is stable, core, and actually seal-twists/executes well, DST is not a meme award. If the Ret slot is inconsistent, do not let the raw sim delta override attendance and role certainty.

### Enhancement Shaman

Enhancement remains a strong DST user:

- P1: +35.21 DPS over `Abacus + Bloodlust Brooch`.
- P2: +33.94 DPS over `Tsunami + Bloodlust Brooch`.

Enhancement is also a reasonable Tsunami-track target if the LC wants cleaner item separation, but Tsunami is not equal to DST.

### Arms Warrior

Arms qualifies under the final no-Badge assumption:

- P1: +26.36 DPS over `Abacus + Bloodlust Brooch`.
- P2: +29.75 DPS over `Tsunami + Bloodlust Brooch`.

Tsunami is worse than Abacus for Arms in this profile. LC takeaway: Arms is eligible after stronger claims are handled, but should not receive Tsunami over Abacus.

### BM Hunter

BM qualifies, but this matrix does not put BM above Survival:

- P1: +29.02 DPS over `Hourglass + Bloodlust Brooch`.
- P2: +22.62 DPS over `Tsunami + Bloodlust Brooch`.

If the raid runs multiple BM hunters and one SV, priority should not be copied across every hunter. Use player quality and attendance, but keep the spec distinction.

### Feral Cat

Feral is the cleanest Tsunami assignment:

- P1: +8.15 DPS over `Hourglass + Bloodlust Brooch`.
- P2: +5.62 DPS over `Tsunami + Bloodlust Brooch`.

LC takeaway: Feral can use DST, but Tsunami solves most of the problem while preserving DST for specs with much larger gaps.

## YouTube Corroboration Overlay

The transcript-backed video pass supports the broad priority bands, not every exact pair:

- Simonize and SNO/RogueCraft support Rogue/Warrior as durable DST holders.
- Ragebtw strengthens Fury's case.
- Hunter videos support Hunter DST usage, but execution matters and Survival should remain separate from BM.
- Ret has high sim value, but execution and haste-window caveats matter.
- Feral's low marginal sim value supports using alternatives instead of early DST.

See [YouTube Corroboration](video-corroboration.md) and [Sim vs Video Corroboration](sim-video-corroboration.md) for source notes and conflicts.

## Practical LC Rules

1. Do not award DST to a class bucket; award it to a stable player in a specific spec.
2. Do not merge Fury with Arms or BM with Survival.
3. Treat Rogue, Survival, Fury, core/high-execution Ret, and Enhancement as the main contested pool.
4. Treat Arms and BM as valid but lower-priority claims.
5. Put Feral Cat on the Tsunami track if possible.
6. Do not award Tsunami to Ret or Arms if Abacus is available.
7. Re-sim if your raid lacks major physical group buffs, uses Draenei hit, or has materially different trinket access.
