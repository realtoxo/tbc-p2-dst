# YouTube Corroboration

This document records transcript-backed video evidence for DST priority. It is a corroboration layer, not the numeric baseline. The numeric baseline remains [DST Sim Results](dst-sim-results.md).

Auto-captions were downloaded with `yt-dlp` into [sources/youtube-transcripts](../sources/youtube-transcripts). Timestamps below refer to those local VTT files and the linked YouTube videos.

## High-Value Sources

### Simonize: Dragonspine Trophy Loot Prio Guide

URL: [Dragonspine Trophy Loot Prio Guide](https://www.youtube.com/watch?v=nkzItWU13kc)

Transcript: [nkzItWU13kc.en.vtt](../sources/youtube-transcripts/nkzItWU13kc.en.vtt)

Source quality: High. This is the most directly relevant DST-priority video found so far.

Useful claims:

- `00:00:12-00:00:19`: DST is P1 BiS for many classes and used in later phases.
- `00:01:15-00:02:19`: The method compares best P1 setup with DST to the next-best setup without DST, using class-specific tools for Fury, Hunter, Rogue, Shaman, and Ret.
- `00:02:19-00:04:22`: Simonize argues percentage gain is safer than raw gain when combining different class tools and actual player performance.
- `00:05:20-00:05:33`: Hunter value is much higher when the hunter melee-weaves correctly.
- `00:06:07-00:07:20`: Hunters need extra execution to realize full DST value; Rogues and Warriors do not.
- `00:07:20-00:07:48`: Enhancement and Ret have haste soft-cap concerns around Windfury weapon and seal twisting.
- `00:08:43-00:09:09`: At a 10-15% drop rate over 16 weeks, expect only about 2-3 DSTs.
- `00:09:17-00:09:36`: Simonize says Enhance, Ret, and Hunters get close alternatives later, while for Rogues and Warriors DST remains "far and away" best.
- `00:09:41-00:10:06`: Raid DPS gain is not the only LC factor.

How it affects this doc:

- Supports keeping player reliability and execution in the LC rule set.
- Supports splitting Hunters by actual ability to melee-weave/handle haste, not just by class.
- Supports higher long-term confidence for Rogue and Warrior than the P1/P2-only sim alone proves.
- Adds caution for Enhancement and Ret despite their good P1/P2 deltas.

### Ragebtw: Fury Warrior DST PRIO

URL: [Fury Warrior DST PRIO](https://www.youtube.com/watch?v=ZPtTpvf3Leg)

Transcript: [ZPtTpvf3Leg.en.vtt](../sources/youtube-transcripts/ZPtTpvf3Leg.en.vtt)

Source quality: Medium-high for Fury-specific context.

Useful claims:

- `00:01:24-00:01:53`: Cites sim data showing Fury Warriors getting the biggest DPS increase from DST in that chart.
- `00:02:31-00:02:57`: Describes Fury trinket pairing logic around DST, Bloodlust Brooch, and Badge of the Swarmguard.
- `00:06:42-00:07:02`: Argues melee classes may show bigger DST increases than the old hunter-first assumption.
- `00:07:40-00:08:36`: Makes the LC argument that Fury can justify DST over the default BM Hunter expectation.
- `00:08:40-00:08:53`: Frames haste as especially valuable for Fury.

How it affects this doc:

- Strongly supports Fury over Arms for Warrior DST priority.
- Supports not treating "hunter first" as automatic.
- Does not by itself prove Fury over Rogue or Ret; it is a Fury-oriented source.

### SNO x ROGUECRAFT: Prepared Podcast DST Discussion

URL: [Best guild loot systems, who gets DST (pt 2) & TBC is broken](https://www.youtube.com/watch?v=E2KfS7YgBVU)

Transcript: [E2KfS7YgBVU.en.vtt](../sources/youtube-transcripts/E2KfS7YgBVU.en.vtt)

Source quality: Medium-high. Discussion format, but RogueCraft is useful rogue-specialist context.

Useful claims:

- `00:13:45-00:13:53`: Discusses DST as roughly a 54-64 DPS increase depending on factors.
- `00:15:07-00:15:13`: Mentions `DST + Bloodlust Brooch` as best in the discussed context.
- `00:20:03-00:20:07`: Frames a realistic DST candidate pool as hunters, warrior, rogue, and possibly enhancement shaman.
- `00:26:24-00:26:29`: Says DST is very good for Rogues and Warriors because it scales resource generation and attack speed.
- `00:42:14-00:42:21`: Describes DST as best during P1, while noting close options exist.

How it affects this doc:

- Supports Rogue as a real contender, not a leftovers recipient.
- Supports including Enhancement in the DST claimant pool.
- Reinforces LC context over rigid spreadsheet-only ordering.

### SNO x ROGUECRAFT: Rogue P2 BIS Guide

URL: [TBC Rogue P2 BIS Gear Guide, Talents & Cheatsheet](https://www.youtube.com/watch?v=pPtzxRTCUlQ)

Transcript: [pPtzxRTCUlQ.en.vtt](../sources/youtube-transcripts/pPtzxRTCUlQ.en.vtt)

Source quality: High for Rogue P2.

Useful claims:

- `00:16:56-00:17:00`: If the rogue has DST, they are still using it in P2, now with `Warp-Spring Coil`.
- `00:25:19-00:25:25`: `Warp-Spring Coil` is a bigger upgrade for rogues without DST because they replace a weaker trinket.

How it affects this doc:

- Supports the P2 Rogue baseline of `DST + Warp-Spring Coil`.
- Supports the conclusion that Rogue remains a strong P2 DST holder.

### Veramos: Survival Hunter Phase 3 BiS And Prio

URL: [Phase 3 Survival Hunter BiS and Prio Guide](https://www.youtube.com/watch?v=yqQagrZrgos)

Transcript: [yqQagrZrgos.en.vtt](../sources/youtube-transcripts/yqQagrZrgos.en.vtt)

Source quality: Medium-high for Survival Hunter retention context. This is Phase 3, so it should not replace the P1/P2 sim baseline.

Useful claims:

- `00:08:07-00:08:12`: DST remains first BiS trinket for Survival in the guide's Phase 3 setup.
- `00:07:02-00:07:10`: A contested leg item is called out as something smart guilds should prioritize to Survival Hunters first.

How it affects this doc:

- Supports Survival as a serious priority spec, not just a BM-adjacent hunter category.
- Supports DST retention beyond P2 for Survival, but this repo still only simmed P1/P2.

### Veramos: BM Hunter Phase 3 BiS And Prio

URL: [Phase 3 Beast Mastery Hunter BiS and Prio Guide](https://www.youtube.com/watch?v=hsIsJVIyOCE)

Transcript: [hsIsJVIyOCE.en.vtt](../sources/youtube-transcripts/hsIsJVIyOCE.en.vtt)

Source quality: Medium-high for BM Hunter retention context. This is Phase 3, so it should not replace the P1/P2 sim baseline.

Useful claims:

- `00:01:46-00:01:51`: BM's Phase 3 trinkets include DST and replace Brooch with `Madness of the Betrayer`.

How it affects this doc:

- Supports BM as a continuing DST user.
- The priority signal is weaker than Survival because the cited section is a BiS listing rather than an explicit cross-class LC argument.

### Jambrosay: BM Hunter Physical DPS Loot Guide

URL: [Phase 3 + 4 TBC Loot Guide for Physical DPS - BM Hunter Edition](https://www.youtube.com/watch?v=VpQyyTCc89M)

Transcript: [VpQyyTCc89M.en.vtt](../sources/youtube-transcripts/VpQyyTCc89M.en.vtt)

Source quality: Medium-high for later-phase BM Hunter trinket context.

Useful claims:

- `00:11:52-00:11:55`: Describes DST as still the top Hunter trinket.
- `00:11:55-00:12:10`: Notes later alternatives get closer by Phase 4.

How it affects this doc:

- Supports the idea that DST has long phase life for Hunters.
- Also supports the caveat that later alternatives eventually compress the gap.

### Sarthe: Hunter T5 Loot Prio & BIS

URL: [Hunter T5 Loot Prio & BiS List Guide](https://www.youtube.com/watch?v=W2J4DOYFbl0)

Transcript: [W2J4DOYFbl0.en.vtt](../sources/youtube-transcripts/W2J4DOYFbl0.en.vtt)

Source quality: Medium-high for Hunter P2 loot context.

Useful claims:

- `00:01:45-00:01:55`: Hunters should not replace Beast Lord until they have the T5 four-piece set.
- `00:04:00-00:04:14`: `Tsunami Talisman` is not a major Hunter upgrade and is not an upgrade over `Bloodlust Brooch`.
- `00:04:27-00:04:40`: Do not focus on Tsunami as a Hunter priority item.

How it affects this doc:

- Supports keeping Hunter P2 trinket discussion centered on DST/Brooch rather than treating Tsunami as a must-have replacement.
- Adds a caveat to the sim table: our conservative P2 replacement uses Tsunami as the best tested no-DST replacement, but Hunter loot council should not over-prioritize Tsunami itself.
- Also preserves the broader T5 Hunter nuance: set timing matters, and individual pieces should not be treated without the four-piece constraint.

### Scottejaye: Feral DPS Items For TBC

URL: [Classic WoW Feral DPS Druid Items you want for TBC](https://www.youtube.com/watch?v=kLUQ5eUnOcE)

Transcript: [kLUQ5eUnOcE.en.vtt](../sources/youtube-transcripts/kLUQ5eUnOcE.en.vtt)

Source quality: Medium for Feral entry/P1 context.

Useful claims:

- `00:16:15-00:16:30`: Treats DST as the likely TBC replacement for `Drake Fang Talisman` and recommends trying to get the relevant trinkets.

How it affects this doc:

- Supports that Feral can use DST seriously.
- Does not contradict the sim result that Feral is a lower early LC priority than Rogue/Fury/Hunter/Enh/Ret in the P1/P2 matrix.

### Blayst: Ret Paladin Seal Twisting And Rotation

URL: [Classic TBC Retribution Paladin: Seal Twisting & Rotation Guide](https://www.youtube.com/watch?v=PIjJ2pcnLVQ)

Transcript: [PIjJ2pcnLVQ.en.vtt](../sources/youtube-transcripts/PIjJ2pcnLVQ.en.vtt)

Source quality: Medium-high for Ret execution context.

Useful claims:

- `00:04:33-00:05:08`: Ret twists differently inside Bloodlust/Heroism than outside it, including back-to-back twist windows.
- `00:05:21-00:05:30`: Recommends aligning Avenging Wrath, trinkets, and DPS potions right before an attack lands.

How it affects this doc:

- Supports the idea that Ret's DST value is execution-sensitive and haste-window-sensitive.
- Does not provide a direct DST priority claim, so it should not be used to move Ret above other specs by itself.

## Lower-Value / Context Sources

### MemeDaddy: Dragonstrikes, Mongoose, DST, and Flurry

URL: [How fast do they go? Dragonstrikes, mongoose, dragonspine trophy, and flurry](https://www.youtube.com/watch?v=7kt7Poj6-kw)

Transcript: [7kt7Poj6-kw.en.vtt](../sources/youtube-transcripts/7kt7Poj6-kw.en.vtt)

Use: Mechanics color only. It demonstrates DST stacking with other haste effects for Enhancement-style setups, but it is not a loot-priority source.

### Reggie VS Solace: Madness and DST Uptime Breakdown

URL: [Madness of the Betrayer and Dragonspine Trophy Uptime Breakdown](https://www.youtube.com/watch?v=t_vy83IEBnk)

Transcript: [t_vy83IEBnk.en.vtt](../sources/youtube-transcripts/t_vy83IEBnk.en.vtt)

Use: Uptime/mechanics context only. It is later-phase warrior trinket testing, not P1/P2 LC priority.

### Guzu: Warrior Trinket Combo

URL: [The BEST Trinket Combo! | TBC Classic Warrior](https://www.youtube.com/watch?v=VmGqBJPIJwk)

Transcript: [VmGqBJPIJwk.en.vtt](../sources/youtube-transcripts/VmGqBJPIJwk.en.vtt)

Use: Warrior P1 alternative-trinket context around `Bloodlust Brooch` and `Badge of the Swarmguard`, not DST priority.

## Consensus And Divergence

The transcript-backed video layer mostly supports the numeric pass:

- Rogue and Fury have strong external support.
- Hunters have strong retention support, with Survival having the clearest spec-specific signal from the sources reviewed.
- Enhancement and Ret qualify, but Simonize's haste-soft-cap warning argues against ranking them by raw sim delta alone.
- No credible video found so far argues Arms should outrank Fury for DST.
- Feral Cat has some support as a legitimate DST user, but not as an early cross-class priority.

The main divergence is Ret. The P1/P2 WoWSims baseline gives Ret a very high marginal gain, but Simonize's longer-term commentary says Ret gets closer alternatives later and has seal-twist haste constraints. LC should keep Ret in the main pool only for a stable, high-performing core ret.

## Unverified Leads

- Holderhek, `This is why you main Enh Shaman in TBC Fresh`: surfaced as an Enhancement Shaman specialist lead, but the YouTube video is private and no transcript was available locally. It is not used as evidence in the priority recommendation.
