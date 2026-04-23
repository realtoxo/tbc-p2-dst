# DST Sim Assumption Log

## Rejected / Experimental Baseline

`sources/wowsims-dstbench.csv` was generated from built-in WoWSims TBC presets. It is useful for checking that the local harness runs and for seeing how WoWSims models `Dragonspine Trophy`, but it is not decision-grade for loot priority.

Reasons:
- Several presets are Phase 1 defaults, while Enhancement uses a Phase 2 preset and Retribution uses a Phase 4 preset.
- Some test presets use Horde races or Horde-specific assumptions.
- The project goal is an Alliance loot-council policy.
- The user requested Wowhead gearsets and full group buffs with no Draenei hit aura.

Replacement baseline:
- Gearsets should be built from Wowhead class/spec BiS guides where available.
- Races should use realistic Alliance choices.
- Raid/party buffs should be full physical-DPS support, but exclude Draenei hit.
- Outputs should clearly separate sim deltas from LC policy.

## Accepted Baseline

`sources/wowsims-wowhead-dstbench.csv` is the current decision-grade numeric baseline for this repository.

Properties:
- Phase 1 and Phase 2 are modeled separately.
- BM Hunter and Survival Hunter are separate rows.
- Fury Warrior and Arms Warrior are separate rows.
- Combat Rogue, Feral Cat, Enhancement Shaman, and Retribution Paladin are included.
- Enhancement Shaman uses a no-Draenei-hit proxy race because WoWSims automatically grants Draenei hit to Draenei characters.
- Enhancement Shaman P2 uses `Netherbane` as the offhand after validating that `Talon of the Phoenix` is main-hand only in WoWSims.
