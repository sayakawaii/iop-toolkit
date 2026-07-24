# Skill Spec: g988-me-index-extractor

## Purpose

Extract the full Managed Entity index from a G.988 source document and produce one canonical Markdown index file.

## Scope

- One full source document per run
- Extract only the ME index
- Do not extract per-ME attribute definitions
- Do not generate JSON

## Required Inputs

- `source_document_path`: path or identifier of the source document

## Optional Inputs

- `source_format`: `pdf`, `ocr_text`, `markdown`, or `plain_text`
- `edition_label`: source edition label such as `G.988-2024`
- `page_range`: optional page range for bounded extraction
- `index_table_hint`: optional description of the known ME index table location

## Input Example

```yaml
source_document_path: docs/g988/input/T-REC-G.988-202403-I!Amd1!PDF-E.pdf
source_format: pdf
edition_label: G.988
page_range: 508-517
index_table_hint: Table 11.2.4-1 - Managed entity identifiers
```

## Output Contract

Output one Markdown document only.

Required structure:

```markdown
# G.988 Managed Entity Index

## Source
- Document:
- Format:
- Edition:

## Managed Entities
| ME ID | ME Name | Section | Page | Status | Notes |
|---:|---|---|---:|---|---|
```

## Output Rules

- Include every ME row that appears in the source index table
- Preserve official ME naming from source when clear
- Preserve source row order when stable
- If the table spans multiple pages, ignore repeated header rows and merge the continued table into one output table
- If `page_range` is provided, prioritize extraction within that range before searching elsewhere
- If `index_table_hint` is provided, treat it as a location hint, not as a substitute for extraction
- If `Section` or `Page` is not present in the source table, leave that cell empty, set `Status` to `needs_review`, and explain the omission in `Notes`
- `ME ID` must be numeric unless the row is explicitly marked `needs_review`
- `Status` must be `ok` or `needs_review`
- `Notes` must be empty unless clarification is needed
- Do not invent missing IDs, names, sections, or pages

## Failure Policy

Use one of these explicit failures when output cannot be produced correctly:

- `INVALID_INPUT`: required input is missing or malformed
- `SOURCE_NOT_FOUND`: source cannot be opened or located
- `AMBIGUOUS_SOURCE`: no reliable ME index table can be identified
- `UNSUPPORTED_PATTERN`: source layout exists but is outside supported extraction scope
- `VALIDATION_FAILED`: output was produced but does not satisfy the required format

## Review Markers

If a row is partially recoverable but uncertain, use one of:

- `needs_review`
- `missing_in_source`
- `not_parsed`

## Acceptance Criteria

- The title is `G.988 Managed Entity Index`
- The `Source` section exists
- The `Managed Entities` table exists
- Every row has exactly six columns
- Every `ME ID` is numeric unless flagged for review

## Few-Shot Example

Input:

```yaml
source_document_path: input/G988-sample.pdf
source_format: pdf
edition_label: G.988 Sample
```

Expected output snippet:

```markdown
# G.988 Managed Entity Index

## Source
- Document: input/G988-sample.pdf
- Format: pdf
- Edition: G.988 Sample

## Managed Entities
| ME ID | ME Name | Section | Page | Status | Notes |
|---:|---|---|---:|---|---|
| 2 | ONU data | 9.1.1 | 100 | ok | |
| 6 | Circuit pack | 9.1.2 | 105 | ok | |
| 347 | IPv6 host config data | 9.3.x | 420 | needs_review | page derived from OCR heading |
```

## Direct Create-Skills Prompt

Create a skill named `g988-me-index-extractor`.

Its only purpose is to extract the full Managed Entity index from a G.988 source document and output one canonical Markdown index.

Required input:

- `source_document_path`

Optional inputs:

- `source_format`
- `edition_label`
- `page_range`
- `index_table_hint`

The output must be one Markdown document with this exact structure:

```markdown
# G.988 Managed Entity Index

## Source
- Document:
- Format:
- Edition:

## Managed Entities
| ME ID | ME Name | Section | Page | Status | Notes |
|---:|---|---|---:|---|---|
```

Behavior rules:

- include every ME row from the source index table
- preserve official source naming when clear
- ignore repeated table headers when the table continues across pages
- prioritize the provided `page_range` when present
- use `index_table_hint` only as a search hint
- if `Section` or `Page` is absent in the source table, leave the cell empty and explain that in `Notes`
- do not invent missing values
- mark uncertain rows with `needs_review`
- `Status` must be `ok` or `needs_review`
- output Markdown only
- never extract per-ME details
- never generate JSON

Failure categories:

- `INVALID_INPUT`
- `SOURCE_NOT_FOUND`
- `AMBIGUOUS_SOURCE`
- `UNSUPPORTED_PATTERN`
- `VALIDATION_FAILED`

## Trial Run Parameters

Use these values for the first real run in this repository:

```yaml
source_document_path: docs/g988/input/T-REC-G.988-202403-I!Amd1!PDF-E.pdf
source_format: pdf
edition_label: G.988
page_range: 508-517
index_table_hint: Table 11.2.4-1 - Managed entity identifiers
```

Expected source location notes:

- table starts on page 508
- table ends on page 517
- table title is `Table 11.2.4-1 - Managed entity identifiers`