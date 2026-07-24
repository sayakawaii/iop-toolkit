# Skill Spec: g988-single-me-extractor

## Purpose

Extract one Managed Entity definition from a G.988 source document and emit one normalized Markdown artifact.

## Scope

- One ME per run
- Extract one complete ME definition
- Preserve source order and naming when clear
- Do not generate repository JSON

## Required Inputs

- `source_document_path`
- `me_id`
- `me_name`

## Optional Inputs

- `source_format`
- `source_section`
- `source_page`
- `edition_label`
- `index_row_markdown`

## Input Example

```yaml
source_document_path: input/G988.pdf
source_format: pdf
me_id: 347
me_name: IPv6 host config data
source_section: 9.3.x
source_page: 420
```

## Output Contract

Output one normalized Markdown document only.

Required structure:

```markdown
# Managed Entity

## Identity
- ME ID:
- ME Name:
- Source Section:
- Source Page:

## Classification
- Access:
- Type:
- Actions:
- Instance Type:
- Instance Value:

## Attributes
### Attribute 1
- Name:
- Size:
- Format:
- Access:
- Category:

#### Structure Fields
| Field Name | Size | Format |
|---|---:|---|
```

## Output Rules

- Preserve attribute order exactly
- Preserve official ME and attribute names when clear
- `Actions` must be comma-separated in source order
- `Access` must be comma-separated
- Include `#### Structure Fields` only for structural attributes
- Use `needs_review` instead of guessing uncertain values
- Do not emit JSON

## Failure Policy

Use one of these explicit failures when output cannot be produced correctly:

- `INVALID_INPUT`: required input is missing or malformed
- `SOURCE_NOT_FOUND`: target ME section cannot be located
- `AMBIGUOUS_SOURCE`: multiple candidate sections exist and cannot be disambiguated
- `UNSUPPORTED_PATTERN`: source ME layout cannot be normalized safely
- `VALIDATION_FAILED`: output was produced but does not satisfy the required format

## Review Markers

If a field is uncertain but partial extraction is still useful, use one of:

- `needs_review`
- `missing_in_source`
- `not_parsed`

## Acceptance Criteria

- `Identity`, `Classification`, and `Attributes` sections all exist
- `ME ID` and `ME Name` are present
- Attributes are numbered sequentially starting from `Attribute 1`
- Every attribute contains `Name`, `Size`, `Format`, `Access`, and `Category`

## Few-Shot Example

Input:

```yaml
source_document_path: input/G988.pdf
source_format: pdf
me_id: 347
me_name: IPv6 host config data
source_section: 9.3.x
source_page: 420
```

Expected output snippet:

```markdown
# Managed Entity

## Identity
- ME ID: 347
- ME Name: IPv6HostConfigData
- Source Section: 9.3.x
- Source Page: 420

## Classification
- Access: CreatedByONU
- Type: Config
- Actions: Set, Get
- Instance Type: Dynamic
- Instance Value: 0

## Attributes
### Attribute 1
- Name: IPv6 options
- Size: 1
- Format: UnsignedInteger
- Access: R, W, Dummy
- Category: Mandatory
```

## Direct Create-Skills Prompt

Create a skill named `g988-single-me-extractor`.

Its only purpose is to extract one Managed Entity definition from a G.988 source document and output one normalized Markdown artifact.

Required inputs:

- `source_document_path`
- `me_id`
- `me_name`

Optional inputs:

- `source_format`
- `source_section`
- `source_page`
- `edition_label`
- `index_row_markdown`

The output must be one Markdown document with this exact structure:

```markdown
# Managed Entity

## Identity
- ME ID:
- ME Name:
- Source Section:
- Source Page:

## Classification
- Access:
- Type:
- Actions:
- Instance Type:
- Instance Value:

## Attributes
### Attribute 1
- Name:
- Size:
- Format:
- Access:
- Category:

#### Structure Fields
| Field Name | Size | Format |
|---|---:|---|
```

Behavior rules:

- preserve attribute order exactly
- preserve official names when clear
- use `needs_review` instead of guessing
- include structure fields only for structural attributes
- output Markdown only
- never generate repository JSON directly

Failure categories:

- `INVALID_INPUT`
- `SOURCE_NOT_FOUND`
- `AMBIGUOUS_SOURCE`
- `UNSUPPORTED_PATTERN`
- `VALIDATION_FAILED`