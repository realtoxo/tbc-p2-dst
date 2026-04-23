const fs = require('fs');
const path = require('path');

const SPECS = [
  ['fury_warrior', 'p1', 'https://www.wowhead.com/tbc/guide/fury-warrior-dps-karazhan-best-in-slot-gear-burning-crusade-classic-wow'],
  ['fury_warrior', 'p2', 'https://www.wowhead.com/tbc/guide/classes/warrior/fury/dps-bis-gear-pve-phase-2'],
  ['arms_warrior', 'p1', 'https://www.wowhead.com/tbc/guide/arms-warrior-dps-karazhan-best-in-slot-gear-burning-crusade-classic-wow'],
  ['arms_warrior', 'p2', 'https://www.wowhead.com/tbc/guide/classes/warrior/arms/dps-bis-gear-pve-phase-2'],
  ['rogue', 'p1', 'https://www.wowhead.com/tbc/guide/rogue-dps-karazhan-best-in-slot-gear-burning-crusade-classic-wow'],
  ['rogue', 'p2', 'https://www.wowhead.com/tbc/guide/classes/rogue/dps-bis-gear-pve-phase-2'],
  ['bm_hunter', 'p1', 'https://www.wowhead.com/tbc/guide/beast-mastery-hunter-dps-karazhan-best-in-slot-gear-burning-crusade-classic-wow'],
  ['bm_hunter', 'p2', 'https://www.wowhead.com/tbc/guide/classes/hunter/beast-mastery/dps-bis-gear-pve-phase-2'],
  ['survival_hunter', 'p1', 'https://www.wowhead.com/tbc/guide/survival-hunter-dps-karazhan-best-in-slot-gear-burning-crusade-classic-wow'],
  ['survival_hunter', 'p2', 'https://www.wowhead.com/tbc/guide/classes/hunter/survival/dps-bis-gear-pve-phase-2'],
  ['feral_cat', 'p1', 'https://www.wowhead.com/tbc/guide/feral-druid-dps-karazhan-best-in-slot-gear-burning-crusade-classic-wow'],
  ['feral_cat', 'p2', 'https://www.wowhead.com/tbc/guide/classes/druid/feral/dps-bis-gear-pve-phase-2'],
  ['enhancement_shaman', 'p1', 'https://www.wowhead.com/tbc/guide/enhancement-shaman-dps-karazhan-best-in-slot-gear-burning-crusade-classic-wow'],
  ['enhancement_shaman', 'p2', 'https://www.wowhead.com/tbc/guide/classes/shaman/enhancement/dps-bis-gear-pve-phase-2'],
  ['ret_paladin', 'p1', 'https://www.wowhead.com/tbc/guide/retribution-paladin-dps-karazhan-best-in-slot-gear-burning-crusade-classic-wow'],
  ['ret_paladin', 'p2', 'https://www.wowhead.com/tbc/guide/classes/paladin/retribution/dps-bis-gear-pve-phase-2'],
];

const SLOT_RULES = [
  ['head', /Best in Slot Head/i, 1],
  ['shoulder', /Best in Slot Shoulder/i, 1],
  ['back', /Best in Slot Back|Best in Slot Cloak/i, 1],
  ['chest', /Best in Slot Chest/i, 1],
  ['wrist', /Best in Slot Wrist/i, 1],
  ['hands', /Best in Slot Hand/i, 1],
  ['waist', /Best in Slot Waist/i, 1],
  ['legs', /Best in Slot Leg/i, 1],
  ['feet', /Best in Slot Feet/i, 1],
  ['neck', /Best in Slot Neck/i, 1],
  ['rings', /Best in Slot Ring/i, 2],
  ['trinkets', /Best in Slot Trinket/i, 2],
  ['two_hand', /Best in Slot Two-Handed|Best in Slot Two Handed|Best in Slot Weapons? for Feral|Best in Slot Weapons? for Retribution/i, 1],
  ['main_hand', /Best in Slot Main[- ]Hand|Best in Slot Melee Weapons|Best in Slot Weapons for Enhancement|Weapons for Survival/i, 1],
  ['off_hand', /Best in Slot Off[- ]Hand|Best in Slot Offhands|Best in Slot Melee Weapons|Best in Slot Weapons for Enhancement|Weapons for Survival/i, 1],
  ['ranged', /Best in Slot Ranged|Best in Slot Idols|Best in Slot Totems|Best in Slot Librams/i, 1],
];

const SPEC_USES_TWO_HAND = new Set(['arms_warrior', 'ret_paladin', 'feral_cat']);

const MANUAL_OVERRIDES = {
  rogue_p1: {
    main_hand: [28438, 'Wowhead Main-Hand section: first BiS row.'],
  },
  rogue_p2: {
    main_hand: [30082, 'Wowhead Main-Hand section: first BiS row.'],
  },
  bm_hunter_p2: {
    off_hand: [29948, 'Wowhead P2 Main-Hand section lists Best MH and Best OH in one table.'],
  },
  survival_hunter_p1: {
    main_hand: [28263, 'Wowhead P1 Weapons section: first Best MH row.'],
    off_hand: [28572, 'Wowhead P1 Weapons section: first Best OH row.'],
    ranged: [28772, 'Wowhead P1 Weapons section: first ranged Best row.'],
    wrist: [29246, 'Sim fallback: Wowhead-listed next Best non-random alternative to Ravager\'s Bands of Agility.'],
    waist: [28750, 'Sim fallback: Wowhead-listed next Best non-random alternative to Lurker\'s Belt of Agility.'],
    feet: [28545, 'Sim fallback: Wowhead-listed next Best non-random alternative to Glider\'s Sabatons of Agility.'],
  },
  survival_hunter_p2: {
    off_hand: [29948, 'Wowhead P2 Melee Weapons section: Best OH - Personal DPS row.'],
  },
  enhancement_shaman_p1: {
    main_hand: [28313, 'Wowhead P1 Weapons section: first Best MH row.'],
    off_hand: [28308, 'Wowhead P1 Weapons section: first Best non-MH 2.6 speed row.'],
  },
  enhancement_shaman_p2: {
    main_hand: [29996, 'Wowhead P2 Weapons section: first 2.7 MH/OH row.'],
    off_hand: [29924, 'Wowhead P2 Weapons section: next legal one-hand 2.6 speed row; Talon of the Phoenix is main-hand only in WoWSims.'],
  },
};

function decodeEntities(text) {
  return text
    .replace(/&amp;/g, '&')
    .replace(/&lt;/g, '<')
    .replace(/&gt;/g, '>')
    .replace(/&quot;/g, '"')
    .replace(/&#039;/g, "'");
}

function csvCell(value) {
  const s = String(value ?? '');
  return /[",\n]/.test(s) ? `"${s.replace(/"/g, '""')}"` : s;
}

function itemData(html) {
  const result = {};
  for (const match of html.matchAll(/WH\.Gatherer\.addData\(3,\s*5,\s*(\{[\s\S]*?\})\);/g)) {
    try {
      const parsed = JSON.parse(match[1]);
      for (const [id, data] of Object.entries(parsed)) result[id] = data.name_enus || '';
    } catch {}
  }
  return result;
}

function guideBody(html) {
  const match = html.match(/WH\.markup\.printHtml\(("(?:\\.|[^"\\])*")\s*,\s*"guide-body"/s);
  if (!match) throw new Error('Could not find guide-body payload');
  return JSON.parse(match[1]);
}

function sections(body) {
  return [...body.matchAll(/\[h3[^\]]*\]([^\[]+)\[\/h3\]([\s\S]*?)(?=\[h3|$)/g)]
    .map((m) => ({ title: decodeEntities(m[1].trim()), body: m[2] }));
}

function rows(sectionBody) {
  const out = [];
  const rowRegex = /\[tr\]\[td\]([\s\S]*?)\[\/td\]\s*\[td\]([\s\S]*?)\[\/td\]/g;
  for (const match of sectionBody.matchAll(rowRegex)) {
    const rank = match[1].replace(/\[[^\]]+\]/g, '').trim();
    const ids = [...match[2].matchAll(/\[item=(\d+)\]/g)].map((m) => Number(m[1]));
    if (ids.length) out.push({ rank, ids });
  }
  return out;
}

function bestItems(sectionBody, count) {
  const allRows = rows(sectionBody);
  const candidates = allRows.filter((r) => /^Best\b/i.test(r.rank) || /^BiS\b/i.test(r.rank) || /^Game Best\b/i.test(r.rank) || /^Never Un-equip\b/i.test(r.rank));
  const ids = [];
  for (const row of candidates) {
    for (const id of row.ids) {
      if (!ids.includes(id)) ids.push(id);
      if (ids.length >= count) return ids;
    }
  }
  for (const row of allRows) {
    for (const id of row.ids) {
      if (!ids.includes(id)) ids.push(id);
      if (ids.length >= count) return ids;
    }
  }
  return ids;
}

function setSlot(selected, slot, id, note) {
  const index = selected.findIndex((item) => item.slot === slot);
  const entry = { slot, id, note };
  if (index >= 0) selected[index] = entry;
  else selected.push(entry);
}

async function main() {
  const outRows = [];
  const gearsets = {};
  const rawDir = path.join('sources', 'wowhead-html');
  fs.mkdirSync(rawDir, { recursive: true });

  for (const [spec, phase, url] of SPECS) {
    const response = await fetch(url, { headers: { 'user-agent': 'Mozilla/5.0' } });
    if (!response.ok) throw new Error(`${response.status} ${url}`);
    const html = await response.text();
    fs.writeFileSync(path.join(rawDir, `${spec}-${phase}.html`), html);

    const names = itemData(html);
    const body = guideBody(html);
    const sec = sections(body);
    const selected = [];

    for (const [slot, pattern, count] of SLOT_RULES) {
      if (SPEC_USES_TWO_HAND.has(spec) && (slot === 'main_hand' || slot === 'off_hand')) continue;
      if (!SPEC_USES_TWO_HAND.has(spec) && slot === 'two_hand') continue;
      const found = sec.find((s) => pattern.test(s.title));
      if (!found) continue;
      const ids = bestItems(found.body, count);
      for (let idx = 0; idx < ids.length; idx++) {
        const id = ids[idx];
        selected.push({ slot: slot === 'rings' ? `ring${idx + 1}` : slot === 'trinkets' ? `trinket${idx + 1}` : slot, id });
        outRows.push([spec, phase, slot, idx + 1, id, names[id] || '', found.title, url]);
      }
    }

    const key = `${spec}_${phase}`;
    if (MANUAL_OVERRIDES[key]) {
      for (const [slot, [id, note]] of Object.entries(MANUAL_OVERRIDES[key])) {
        setSlot(selected, slot, id, note);
        outRows.push([spec, phase, slot, 'override', id, names[id] || '', note, url]);
      }
    }

    gearsets[key] = { spec, phase, url, selected };
  }

  fs.writeFileSync(
    path.join('sources', 'wowhead-bis-items.csv'),
    [['spec', 'phase', 'slot', 'rank_index', 'item_id', 'item_name', 'section', 'url'], ...outRows]
      .map((row) => row.map(csvCell).join(','))
      .join('\n') + '\n'
  );
  fs.writeFileSync(path.join('sources', 'wowhead-bis-gearsets.json'), JSON.stringify(gearsets, null, 2) + '\n');
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
