# TBC Dragonspine Trophy Loot Research

This repository is a loot-council research pack for **Dragonspine Trophy** in TBC Classic, focused on Phase 1 and Phase 2 gearing.

The current baseline is:

- Alliance assumptions.
- No Draenei `Heroic Presence` hit aura.
- Full raid/physical DPS buff support.
- Phase 1 BiS and Phase 2 BiS gearsets sourced from Wowhead class/spec guides.
- Separate treatment for BM Hunter vs Survival Hunter, Fury Warrior vs Arms Warrior, and every normal physical DPS DST claimant.

## Documents

- [LC Cheat Sheet](LC-CHEATSHEET.md): Quick officer-night DST/Tsunami decisions and DPS deltas.
- [Final LC Report](FINAL-REPORT.md): Human-readable DST/Tsunami award plan and DPS spreads.
- [DST Loot Priority](docs/dst-loot-priority.md): LC-facing priority recommendation, class/spec notes, and tradeoffs.
- [DST Sim Results](docs/dst-sim-results.md): Parsed results from the Wowhead gearset sim matrix.
- [DST Methodology](docs/dst-methodology.md): Assumptions, source files, sim harness details, and caveats.
- [YouTube Corroboration](docs/video-corroboration.md): Transcript-backed creator/source notes.
- [Sim vs Video Corroboration](docs/sim-video-corroboration.md): Where the sim numbers support or conflict with the video layer.
- [Sim Assumption Log](notes/sim-assumption-log.md): Notes on rejected presets and accepted baseline changes.

## Bottom Line

The final no-Badge/no-Slayer numbers support **Combat Rogue, Survival Hunter, Fury Warrior, core/high-execution Retribution Paladin, and Enhancement Shaman** as the main DST priority pool. Arms Warrior and BM Hunter still qualify, but their measured marginal value is lower under these assumptions. Feral Cat is the cleanest Tsunami-track claimant rather than an early DST target.

Do not merge Warrior specs or Hunter specs together:

- Fury is materially ahead of Arms for DST value.
- Survival is materially ahead of BM in this profile.
- Rogue is a legitimate high-priority claimant, not a leftover category.

## Key Data Sources

Primary source artifacts in this repo:

- [Final DST spread CSV](sources/final-dst-spreads.csv)
- [Final P2 DST/Tsunami ladder CSV](sources/final-p2-mutual-exclusion-ladder.csv)
- [Wowhead gearset JSON](sources/wowhead-bis-gearsets.json)
- [Wowhead extracted item CSV](sources/wowhead-bis-items.csv)
- [WoWSims DST matrix CSV](sources/wowsims-wowhead-dstbench.csv)
- [WoWSims harness](sources/wowsims-wowhead-dstbench.go)
- [Wowhead extraction script](scripts/extract-wowhead-bis.js)

External sources used:

- [Dragonspine Trophy - Wowhead](https://www.wowhead.com/tbc/item=28830/dragonspine-trophy)
- [TBC Classic stats overview - Wowhead](https://www.wowhead.com/tbc/guide/classic-the-burning-crusade-stats-overview)
- [WoWSims TBC](https://wowsims.github.io/tbc/)
- [WoWSims TBC GitHub](https://github.com/wowsims/tbc)
- [Simonize: Dragonspine Trophy Loot Prio Guide](https://www.youtube.com/watch?v=nkzItWU13kc)
- [Ragebtw: Fury Warrior DST PRIO](https://www.youtube.com/watch?v=ZPtTpvf3Leg)
- [SNO x ROGUECRAFT: Prepared Podcast DST Discussion](https://www.youtube.com/watch?v=E2KfS7YgBVU)
- [SNO x ROGUECRAFT: Rogue P2 BIS Guide](https://www.youtube.com/watch?v=pPtzxRTCUlQ)
- [Veramos: Survival Hunter Phase 3 BiS and Prio Guide](https://www.youtube.com/watch?v=yqQagrZrgos)
- [Veramos: BM Hunter Phase 3 BiS and Prio Guide](https://www.youtube.com/watch?v=hsIsJVIyOCE)
- [Jambrosay: BM Hunter Physical DPS Loot Guide](https://www.youtube.com/watch?v=VpQyyTCc89M)
- [Sarthe: Hunter T5 Loot Prio & BiS List Guide](https://www.youtube.com/watch?v=W2J4DOYFbl0)
- [Scottejaye: Feral DPS Items For TBC](https://www.youtube.com/watch?v=kLUQ5eUnOcE)
- [Blayst: Ret Paladin Seal Twisting & Rotation Guide](https://www.youtube.com/watch?v=PIjJ2pcnLVQ)

Wowhead P1/P2 class/spec BiS guide URLs are stored on every row of [sources/wowhead-bis-items.csv](sources/wowhead-bis-items.csv) and [sources/wowsims-wowhead-dstbench.csv](sources/wowsims-wowhead-dstbench.csv).

## Current Research Status

The numeric baseline is complete for:

- Arms Warrior
- Fury Warrior
- Combat Rogue
- BM Hunter
- Survival Hunter
- Feral Cat
- Enhancement Shaman
- Retribution Paladin

YouTube/class-specialist corroboration has been started and is tracked in [YouTube Corroboration](docs/video-corroboration.md). More class-specific sources can still be added, especially for Ret, Enhancement, and Feral.
