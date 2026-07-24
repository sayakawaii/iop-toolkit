#!/usr/bin/env python3
"""G.988 per-ME batch extractor — Skill B runner.

Reads:   docs/g988/index/me-index.md
PDF:     docs/g988/input/T-REC-G.988-202403-I!Amd1!PDF-E.pdf
Writes:  docs/g988/me/<id:04d>-<slug>.md  (one file per ME)
State:   /tmp/me_extract_progress.json    (checkpoint — re-run to resume)

Usage:
  python3 docs/g988/extract_all_mes.py
  python3 docs/g988/extract_all_mes.py > /tmp/me_extract.log 2>&1 &

To check progress while running:
  tail -f /tmp/me_extract.log
  python3 -c "import json; d=json.load(open('/tmp/me_extract_progress.json')); \
    print(f'Done: {len(d[\"done\"])}, Failed: {len(d[\"failed\"])}')"
"""

import os
import re
import json
import sys
import unicodedata

WORKSPACE  = "/home/mingheh/project/omcianalyzer"
PDF_PATH   = os.path.join(WORKSPACE, "docs/g988/input/T-REC-G.988-202403-I!Amd1!PDF-E.pdf")
INDEX_PATH = os.path.join(WORKSPACE, "docs/g988/index/me-index.md")
OUTPUT_DIR = os.path.join(WORKSPACE, "docs/g988/me")
CHECKPOINT = "/tmp/me_extract_progress.json"

# PDF page 0-index to printed page:  printed = pdf_idx - PAGE_OFFSET
PAGE_OFFSET = 6

# How many pages ahead to capture after the section heading page
MAX_WINDOW  = 8


# ─────────────────────────────────────────────────────────────── helpers

def slugify(s):
    s = unicodedata.normalize("NFKD", s).encode("ascii", "ignore").decode()
    s = re.sub(r"[^\w]+", "-", s.lower()).strip("-")
    return s[:64]


def parse_index(path):
    """Return [(me_id, me_name), ...] for individual (non-range) MEs only."""
    entries = []
    with open(path, encoding="utf-8") as f:
        for line in f:
            # Only rows with a plain integer ME ID
            m = re.match(r"^\|\s*(\d+)\s*\|([^|]+)\|", line)
            if not m:
                continue
            me_id   = int(m.group(1))
            me_name = m.group(2).strip()
            # Skip placeholder entries
            if re.search(r"\(intentionally left blank\)|\(reserved\)", me_name, re.I):
                continue
            entries.append((me_id, me_name))
    return entries


# ─────────────────────────────────────────────────────────────── PDF loading

def load_pages(pdf_path):
    from pypdf import PdfReader
    reader = PdfReader(pdf_path)
    pages  = []
    total  = len(reader.pages)
    for i, page in enumerate(reader.pages):
        if i % 100 == 0:
            print(f"  loading pages {i}/{total} ...", flush=True)
        pages.append(page.extract_text() or "")
    return pages


# ─────────────────────────────────────────────────────────────── text normalization

def normalize(text):
    """Collapse PDF column-padding artifacts into readable prose.

    pypdf extracts the two-column G.988 text with large runs of spaces at the
    right edge of each line (the 'gutter').  This function:
    1. Collapses 3+ consecutive spaces into one space.
    2. Joins continuation lines (lines that do NOT start a new word/heading)
       by replacing bare newlines with a space.
    3. Preserves double-newlines (paragraph / attribute boundaries).
    """
    # Step 1: collapse column padding (3+ spaces → single space)
    text = re.sub(r" {3,}", " ", text)
    # Step 2: join wrapped lines — a newline preceded/followed by lowercase or
    #         mid-sentence text is treated as a line-wrap, not a paragraph break.
    #         Double newlines are preserved as paragraph separators.
    text = re.sub(r"(?<!\n)\n(?!\n)", " ", text)
    # Step 3: collapse multiple blank lines to one
    text = re.sub(r"\n{3,}", "\n\n", text)
    return text.strip()


# Matches both forms of the G.988 page header (page number on left or right):
#   "  Rec. ITU-T G.988 (2022) Amd. 1 (03/2024) 141 "
#   "142 Rec. ITU-T G.988 (2022) Amd. 1 (03/2024) "
_PAGE_HDR2 = re.compile(r"(?m)^\s*(?:\d+\s+)?Rec\.\s+ITU-T\s+G\.988[^\n]*\n?")


# ─────────────────────────────────────────────────────────────── section search

# G.988 section name cleanup: strip B-PON/G-PON footnote suffixes that appear
# in the index but not in the PDF section headings.
_BPON_SUFFIX = re.compile(r"B-?PON\s*$", re.I)
_PARENS      = re.compile(r"\([^)]*\)")
_EXTRA_SPACE = re.compile(r"\s{2,}")


def clean_for_search(me_name):
    s = _BPON_SUFFIX.sub("", me_name)
    s = _PARENS.sub("", s)
    s = _EXTRA_SPACE.sub(" ", s).strip()
    return s


def build_heading_pattern(me_name):
    """Build a regex that matches a G.988 section heading for this ME."""
    clean = clean_for_search(me_name)
    words = clean.split()
    if not words:
        return None
    # Match section number + ME name (first 4 significant words, allow extra whitespace)
    sig_words = [w for w in words if len(w) > 1][:5]
    if not sig_words:
        return None
    # Allow flexible whitespace between words (PDF may insert newlines/spaces)
    word_pat = r"\s+".join(re.escape(w) for w in sig_words)
    pat = re.compile(r"(\d+(?:\.\d+){2,})\s+" + word_pat, re.I)
    return pat


def _is_definition_page(pages, page_idx):

    """Return True if this page looks like a real ME definition (not a TOC entry).

    A definition page (or the page immediately following) contains at least one
    of the canonical G.988 section keywords: Relationships, Attributes, Actions.
    """
    check = "".join(pages[page_idx : page_idx + 2])  # this page + next
    return bool(re.search(r"\b(Relationships|Attributes|Actions)\b", check))


def find_section(pages, me_name):
    """Return (start_page_idx, section_num_str, end_page_idx) or None.

    Skips TOC / cross-reference pages that contain only the heading without the
    actual definition body.  A page is accepted only when it (or the next page)
    contains at least one of: Relationships, Attributes, Actions.
    """
    pat = build_heading_pattern(me_name)
    if pat is None:
        return None

    for i, raw_text in enumerate(pages):
        m = pat.search(raw_text)
        if m and _is_definition_page(pages, i):
            section_num = m.group(1)
            depth       = len(section_num.split("."))
            # Scan forward to find where this section ends (next heading at same depth)
            end = i
            for j in range(i + 1, min(i + MAX_WINDOW + 1, len(pages))):
                next_heading = re.search(r"(\d+(?:\.\d+)+)\s+[A-Z]", pages[j])
                if next_heading:
                    nd = len(next_heading.group(1).split("."))
                    if nd <= depth:
                        break  # next same-level or parent section starts here
                end = j
            return i, section_num, min(end + 1, len(pages) - 1)

    # Fallback: partial name match (first 5 significant words, ≥ 8 chars) on a
    # page that also has a section number AND definition keywords
    clean   = clean_for_search(me_name)
    words   = clean.split()
    partial = " ".join(words[:5])[:40]
    if len(partial) >= 8:
        for i, raw_text in enumerate(pages):
            if (partial.lower() in raw_text.lower()
                    and re.search(r"\d+\.\d+\.\d+", raw_text)
                    and _is_definition_page(pages, i)):
                return i, "needs_review", min(i + MAX_WINDOW, len(pages) - 1)

    return None


# ─────────────────────────────────────────────────────────────── section clipping

def clip_to_section(text, section_num, me_name):
    """Return only the content of this ME's section in the window text.

    A. Trim start: skip text that precedes the section heading (tail of a
       previous section that happens to share the same page).
    B. Trim end: truncate at the first sibling/parent section heading.
    """
    if not section_num or section_num == "needs_review":
        return text

    # Heading pattern: section number + first 2 significant words of ME name
    clean = clean_for_search(me_name)
    words = [w for w in clean.split() if len(w) > 1][:2]
    if words:
        head_pat = re.compile(
            re.escape(section_num) + r"\s+[\s\S]{0,30}?" + re.escape(words[0]),
            re.I
        )
    else:
        head_pat = re.compile(re.escape(section_num) + r"\s", re.I)

    sm = head_pat.search(text)
    if sm and sm.start() > 50:
        text = text[sm.start():]

    # Truncate at the next sibling or parent section heading.
    # E.g., current = 9.2.10 → match 9.2.<any> that appears after pos 100.
    parts = section_num.split(".")
    if len(parts) >= 2:
        sib_prefix = re.escape(".".join(parts[:-1])) + r"\.\d+"
    else:
        sib_prefix = r"\d+"
    next_sec_pat = re.compile(r"\b" + sib_prefix + r"\s+[A-Z]")
    for m in next_sec_pat.finditer(text):
        if m.start() > 100:   # skip the current heading itself
            text = text[: m.start()]
            break

    return text


# ─────────────────────────────────────────────────────────────── attribute parsing

# Attribute triplet at end of definition:  (access) (optionality) (size)
# access:      e.g.  R  /  R, W  /  R, W, set-by-create  /  W
# optionality: mandatory | optional | conditional
# size:        N bytes | N byte | N bits | variable
_ATTR_TRIPLET = re.compile(
    r"\(([A-Za-z,\s/\-]+?)\)"          # (access)
    r"\s*\((mandatory|optional|conditional)\)"  # (optionality)
    r"\s*\((\d+(?:\.\d+)?\s*(?:bytes?|bits?)|variable)\)",  # (size)
    re.I
)

# Section heading (for stripping page headers from the window)
_PAGE_HEADER  = re.compile(r"Rec\.\s+ITU-T\s+G\.988[^\n]*\n?", re.I)

# Actions list
_ACTIONS_BLOCK = re.compile(
    r"\bActions?\b\s*(.*?)(?=\s*\bNotifications?\b|\Z)",
    re.I | re.S
)
_ACTION_VERBS = re.compile(
    r"\b(Create|Delete|Get|Set|Get next|Get current data|Set table|Test|Sync time|"
    r"Alarm retrieval)\b",
    re.I
)

# Notifications (None / list)
_NOTIF_BLOCK = re.compile(
    r"(?:^|\n)Notifications?\s*\n?([^\n]+(?:\n[^\n]+){0,5})",
    re.I
)


def parse_attributes(norm_text):
    """Extract attribute list from a normalized ME section text.

    Strategy: find ALL (access)(optionality)(size) triplets in the full text.
    For each triplet k the attribute name lives in the text block BETWEEN the end
    of triplet k-1 and the start of triplet k.  That block always starts with:
        '  Name: description…'
    so the name is simply the text before the FIRST colon in the block.
    """
    triplets = list(_ATTR_TRIPLET.finditer(norm_text))
    if not triplets:
        return []

    attrs = []

    for k, tm in enumerate(triplets):
        if k == 0:
            # First attribute: find the "Attributes" keyword, then look for Name:
            kw_pos = norm_text.lower().rfind("attributes", 0, tm.start())
            search_start = kw_pos if kw_pos >= 0 else max(0, tm.start() - 400)
            block = norm_text[search_start : tm.start()]
        else:
            block = norm_text[triplets[k - 1].end() : tm.start()]

        # The attribute name ends at the first colon in the block
        colon_pos = block.find(":")
        if colon_pos < 0:
            continue

        name = block[:colon_pos]
        # Remove any leading sentence fragments (text up to last \n or ". ")
        for sep in ("\n", ". ", "? ", "! "):
            idx = name.rfind(sep)
            if idx >= 0:
                name = name[idx + len(sep):]
        name = name.strip()

        # Validate: must be a clean phrase (letter-start, no periods/parens, ≤ 80 chars)
        if (not name
                or len(name) < 2
                or len(name) > 80
                or not re.match(r"^[A-Za-z\d]", name)
                or re.search(r"[.()\[\]]", name)):
            continue

        # Strip leading "Attributes" section keyword if it leaked into the name
        name = re.sub(r"^Attributes\s+", "", name, flags=re.I).strip()
        # Strip isolated leading page numbers (e.g. "220 Type" → "Type")
        name = re.sub(r"^\d+\s+", "", name).strip()
        if not name or len(name) < 2:
            continue

        access = re.sub(r"\s+", " ", tm.group(1)).strip()
        opt    = tm.group(2).strip()
        size   = re.sub(r"\s+", " ", tm.group(3)).strip()

        attrs.append({
            "name":     name,
            "access":   access,
            "category": opt,
            "size":     size,
            "format":   "needs_review",
        })

    return attrs


def parse_section(raw_window, me_id, me_name, section_num, pdf_start_page):
    """Parse the raw text window for one ME into structured fields."""

    # Remove page headers (both forms: page-number left and right)
    text = _PAGE_HDR2.sub("", raw_window)
    text = normalize(text)
    # Clip to only the current ME's section (remove preceding/following sections)
    text = clip_to_section(text, section_num, me_name)

    result = {
        "me_id":          me_id,
        "me_name":        me_name,
        "source_section": section_num,
        "source_page":    max(1, pdf_start_page - PAGE_OFFSET),
        "access":         "needs_review",
        "type":           "needs_review",
        "actions":        "needs_review",
        "notifications":  "needs_review",
        "instance_type":  "needs_review",
        "instance_value": "needs_review",
        "attributes":     [],
        "raw_text":       text[:6000],   # cap raw for file size
    }

    # Attributes (parse first so we know where the last triplet ends)
    attrs = parse_attributes(text)
    numbered = []
    num = 1
    for a in attrs:
        # Skip the meta "Managed entity ID" attribute (it's the identity, not a user attr)
        if re.match(r"managed entity\s+(id|identifier)", a["name"], re.I):
            # Extract instance info from this attribute instead
            result["instance_type"]  = "singleton" if "0x0000" in text.lower() else "per-instance"
            result["instance_value"] = a["size"]
            continue
        a["number"] = num
        numbered.append(a)
        num += 1

    # Actions — search only in the TAIL of the text after all attribute triplets,
    # to avoid matching the word "actions" inside attribute descriptions.
    triplets = list(_ATTR_TRIPLET.finditer(text))
    tail_start = triplets[-1].end() if triplets else 0
    tail = text[tail_start:]

    m = _ACTIONS_BLOCK.search(tail)
    if m:
        verbs = _ACTION_VERBS.findall(m.group(1)[:300])
        if verbs:
            result["actions"] = ", ".join(dict.fromkeys(
                v.capitalize() if v.islower() else v for v in verbs
            ))

    # Notifications
    m = _NOTIF_BLOCK.search(tail)
    if m:
        notif = re.sub(r"\s+", " ", m.group(1)).strip()
        result["notifications"] = notif
    result["attributes"] = numbered

    return result


# ─────────────────────────────────────────────────────────────── Markdown formatter

def format_md(p):
    lines = [
        "# Managed Entity",
        "",
        "## Identity",
        f"- ME ID: {p['me_id']}",
        f"- ME Name: {p['me_name']}",
        f"- Source Section: {p['source_section']}",
        f"- Source Page: {p['source_page']}",
        "",
        "## Classification",
        f"- Access: {p['access']}",
        f"- Type: {p['type']}",
        f"- Actions: {p['actions']}",
        f"- Instance Type: {p['instance_type']}",
        f"- Instance Value: {p['instance_value']}",
        "",
        "## Attributes",
    ]

    if p["attributes"]:
        for a in p["attributes"]:
            lines += [
                "",
                f"### Attribute {a['number']}",
                f"- Name: {a['name']}",
                f"- Size: {a['size']}",
                f"- Format: {a['format']}",
                f"- Access: {a['access']}",
                f"- Category: {a['category']}",
            ]
    else:
        lines += [
            "",
            "<!-- No attributes successfully parsed; see Raw Source section -->",
        ]

    n_attrs = len(p["attributes"])
    has_review = (
        p["source_section"] == "needs_review"
        or not p["attributes"]
        or p["actions"] == "needs_review"
    )
    lines += [
        "",
        "## Extraction Status",
        f"- Status: auto_extracted",
        f"- Attributes found: {n_attrs}",
        f"- Review needed: {str(has_review).lower()}",
        "",
        "## Raw Source",
        "",
        "```",
        p["raw_text"].replace("```", "'''"),
        "```",
    ]
    return "\n".join(lines) + "\n"


def write_stub(out_file, me_id, me_name, reason):
    content = (
        "# Managed Entity\n\n"
        "## Identity\n"
        f"- ME ID: {me_id}\n"
        f"- ME Name: {me_name}\n"
        "- Source Section: needs_review\n"
        "- Source Page: needs_review\n\n"
        "## Classification\n"
        "- Access: needs_review\n"
        "- Type: needs_review\n"
        "- Actions: needs_review\n"
        "- Instance Type: needs_review\n"
        "- Instance Value: needs_review\n\n"
        "## Attributes\n"
        f"<!-- {reason} -->\n\n"
        "## Extraction Status\n"
        f"- Status: {reason}\n"
        "- Attributes found: 0\n"
        "- Review needed: true\n"
    )
    with open(out_file, "w", encoding="utf-8") as f:
        f.write(content)


# ─────────────────────────────────────────────────────────────── checkpoint

def load_cp():
    if os.path.exists(CHECKPOINT):
        with open(CHECKPOINT) as f:
            return json.load(f)
    return {"done": [], "failed": []}


def save_cp(cp):
    with open(CHECKPOINT, "w") as f:
        json.dump(cp, f, indent=2)


# ─────────────────────────────────────────────────────────────── main

def main():
    os.makedirs(OUTPUT_DIR, exist_ok=True)

    print("=== G.988 per-ME batch extractor ===", flush=True)
    print(f"Index:  {INDEX_PATH}", flush=True)
    print(f"PDF:    {PDF_PATH}", flush=True)
    print(f"Output: {OUTPUT_DIR}", flush=True)
    print(f"State:  {CHECKPOINT}", flush=True)
    print("", flush=True)

    entries = parse_index(INDEX_PATH)
    print(f"Index: {len(entries)} individual MEs to process.", flush=True)

    cp      = load_cp()
    done_ids = set(cp["done"])
    todo    = [(id_, name) for id_, name in entries if id_ not in done_ids]
    print(f"To do: {len(todo)}  (already done: {len(done_ids)})", flush=True)

    if not todo:
        print("Nothing left to do — all MEs already extracted.", flush=True)
        return

    print("Loading all PDF pages into memory (30–60 s) ...", flush=True)
    pages = load_pages(PDF_PATH)
    print(f"Loaded {len(pages)} pages.\n", flush=True)

    for seq, (me_id, me_name) in enumerate(todo, 1):
        slug     = slugify(me_name)
        out_file = os.path.join(OUTPUT_DIR, f"{me_id:04d}-{slug}.md")
        label    = f"[{seq:3d}/{len(todo)}] ME {me_id:5d}  {me_name[:48]}"
        print(label, end=" ... ", flush=True)

        try:
            result = find_section(pages, me_name)

            if result is None:
                print("NOT FOUND", flush=True)
                write_stub(out_file, me_id, me_name, "SOURCE_NOT_FOUND")
            else:
                start_p, section_num, end_p = result
                raw_window = "\n\n---- page break ----\n\n".join(
                    pages[start_p : end_p + 1]
                )
                parsed = parse_section(raw_window, me_id, me_name, section_num, start_p)
                md     = format_md(parsed)
                with open(out_file, "w", encoding="utf-8") as f:
                    f.write(md)
                n_attrs = len(parsed["attributes"])
                sec     = parsed["source_section"]
                print(f"OK  section={sec}  attrs={n_attrs}", flush=True)

        except Exception as exc:
            import traceback
            print(f"ERROR: {exc}", flush=True)
            traceback.print_exc()
            write_stub(out_file, me_id, me_name, f"ERROR")
            cp["failed"].append({"id": me_id, "name": me_name, "reason": str(exc)})

        cp["done"].append(me_id)
        save_cp(cp)

    fails = len(cp["failed"])
    print(f"\n=== Complete ===", flush=True)
    print(f"Done: {len(cp['done'])}   Not-found / errors: {fails}", flush=True)
    if fails:
        print("Failed IDs:", [x["id"] for x in cp["failed"]], flush=True)
    print(f"Output in: {OUTPUT_DIR}", flush=True)


if __name__ == "__main__":
    main()
