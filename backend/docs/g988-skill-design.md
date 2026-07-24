# G.988 ME Extraction Skill Design

## Goal

Build a reusable workflow that converts a full G.988 source document into OMCI ME definitions that can eventually be rendered as JSON files like `resource/me/schema/IPv6HostConfigData.json`.

This document treats the end-to-end job as a workflow, not as a single skill. The stable and reusable parts are split into several skills with clear inputs and outputs.

## Conclusion

For this problem, the better skill-oriented solution is not:

- read one large PDF
- generate all ME JSON files in one pass

The better solution is:

- extract a complete ME index first
- use the index as the only official entry point
- extract each ME definition in a separate step
- normalize each ME definition into a structured intermediate format
- convert the structured intermediate format into the target JSON schema
- validate the generated result

In other words:

- Solution 1 is a workflow, but not a good single skill
- Solution 2 is a good skill architecture

## Why Solution 2 Fits Skill Design Better

Solution 2 is a better fit for skills because it has:

1. Stable boundaries
2. Repeatable inputs and outputs
3. Clear checkpoints for manual review and automated validation
4. Better fault isolation when parsing quality is poor
5. Better reuse across different G.988 editions or vendor supplements

If one ME fails to extract, the workflow can continue. If one attribute format is ambiguous, only the affected ME needs to be corrected. This is much more practical than retrying a full-document one-shot generation.

## Recommended Architecture

### End-to-End Workflow

The full pipeline should be designed as:

1. Parse the source PDF into accessible text and tables
2. Extract the master ME index
3. Generate a canonical Markdown index file
4. Extract each ME definition according to the canonical index
5. Generate one Markdown file per ME as the normalized source of truth
6. Convert each Markdown ME definition into the target JSON schema
7. Run validation on names, IDs, attribute sizes, formats, access flags, and structure layout

### Skill Breakdown

The workflow should be split into the following skills.

#### Skill A: Extract ME Index

Purpose:

Extract the complete list of Managed Entities from the G.988 source and produce a canonical index.

Input:

- full G.988 PDF, or OCR/text extracted document

Output:

- one Markdown file containing all ME IDs, names, and location references

Suggested output file:

- `docs/g988/index/me-index.md`

Typical output fields:

- ME ID
- ME name
- class or entity name as written in the standard
- chapter number
- section number
- page number if available
- notes for ambiguity or parser confidence

Why this is a skill:

- it has one narrow purpose
- it is reusable for every standard revision
- it creates a stable interface for downstream extraction

#### Skill B: Extract Single ME Definition

Purpose:

Read the G.988 source using the master index and extract one ME definition into a normalized Markdown representation.

Input:

- a single ME entry from `me-index.md`
- the source G.988 document or preprocessed text

Output:

- one Markdown file for that ME

Suggested output location:

- `docs/g988/me/<me-id>-<me-name>.md`

Typical output fields:

- ME ID
- ME name
- ME access mode
- ME type
- supported actions
- instance semantics
- attribute list
- attribute order
- attribute name
- attribute size in bytes
- attribute format
- access flags
- category such as Mandatory or Optional
- table structure fields when the attribute is an array or structure
- raw citation text if confidence is low

Why this is a skill:

- it is a repeatable per-entity extraction step
- failures are isolated per ME
- it is easy to review and rerun

#### Skill C: Normalize ME Markdown To JSON Schema

Purpose:

Convert one normalized ME Markdown file into the target JSON format used by this repository.

Input:

- one normalized ME Markdown file

Output:

- one JSON schema file matching the repository format

Suggested output location:

- `resource/me/schema/<MeName>.json`

Target example:

- same structure as `resource/me/schema/IPv6HostConfigData.json`

Why this is a skill:

- it maps from a stable intermediate format to a stable target schema
- it is deterministic compared with direct PDF parsing
- it can be validated mechanically

#### Skill D: Validate Generated ME Definition

Purpose:

Check whether generated JSON and Markdown outputs are internally consistent and compatible with project expectations.

Input:

- one ME Markdown file
- one generated JSON file

Output:

- pass or fail report
- normalized warnings
- repair suggestions if needed

Validation targets:

- ME ID exists and is numeric
- ME name is not empty
- attribute order is preserved
- size units are normalized correctly
- `Array Structure` attributes include `structure.fields`
- access lists are present
- Mandatory and Optional categories are preserved
- field sizes add up correctly when the standard is explicit
- target JSON is valid and parseable

Why this is a skill:

- it is reusable for every generated ME
- it prevents silent corruption from OCR or table parsing errors

## Skill Versus Workflow

This distinction should stay explicit.

### What is a skill here

A skill is a reusable unit such as:

- extract the master ME index
- extract one ME definition
- convert one normalized ME definition into JSON
- validate one generated ME artifact

### What is not a good single skill here

This is too broad to be a single stable skill:

- read the full G.988 PDF and generate every ME JSON file directly

That task is better described as a workflow or orchestrator that calls several skills in sequence.

## Recommended File Structure

The following layout keeps the pipeline explicit and reviewable:

```text
docs/
  g988/
    index/
      me-index.md
    me/
      002-ont-data.md
      006-circuit-pack.md
      347-ipv6-host-config-data.md
    prompts/
      extract-me-index.md
      extract-single-me.md
      markdown-to-json.md
      validate-me.md
resource/
  me/
    schema/
      OntData.json
      CircuitPack.json
      IPv6HostConfigData.json
```

## Intermediate Markdown Contract

The most important design choice is to introduce a strict intermediate format. The PDF should not be converted directly into the final JSON target.

Each ME Markdown file should contain the following sections in a fixed order.

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

### Attribute 2
- Name:
- Size:
- Format:
- Access:
- Category:

### Attribute N
- Name:
- Size:
- Format:
- Access:
- Category:

#### Structure Fields
| Field Name | Size | Format |
|---|---:|---|
```

This contract is important because it gives downstream conversion a stable source regardless of PDF formatting differences.

## Mapping Rules To Repository JSON

The normalized Markdown should be mapped to JSON with the following rules.

### Top-level mapping

- `ME ID` -> `Me.id`
- `ME Name` -> `Me.name`
- `Access` -> `Me.access`
- `Type` -> `Me.type`
- `Actions` -> `Me.action`
- `Instance Type` -> `Me.instance.type`
- `Instance Value` -> `Me.instance.value`

### Attribute mapping

For each attribute:

- `Name` -> `name`
- normalized byte size -> `size`
- standardized format label -> `format`
- access list -> `access`
- mandatory or optional -> `category`

If the attribute is a structured table:

- add `structure.fields`
- each field includes `name`, `size`, and `format`

### Format normalization examples

- `octets` -> `Bytes`
- `string` -> `String`
- `unsigned integer` -> `UnsignedInteger`
- table with row layout -> `Array Structure`
- fixed repeated scalar values -> `Array UnsignedInteger N` only if the source is explicit and the repository already uses this convention

## Prompt Templates

The following templates are suitable if this pipeline is driven by LLM-based skills.

### Prompt Template For Skill A

Task:

Extract the full Managed Entity index from the provided G.988 source.

Requirements:

- include every Managed Entity ID and name
- preserve the official naming used in the source
- record section and page references when available
- do not infer missing IDs
- if a table row is ambiguous, mark it as `needs_review`
- output only the agreed Markdown index format

Expected output shape:

```markdown
# G.988 Managed Entity Index

| ME ID | ME Name | Section | Page | Status | Notes |
|---:|---|---|---:|---|---|
```

### Prompt Template For Skill B

Task:

Using one ME entry from the canonical index, extract the complete definition for that ME from the G.988 source and render it in the agreed normalized Markdown contract.

Requirements:

- preserve attribute order exactly
- keep official ME and attribute names
- normalize sizes to bytes when possible
- keep access flags as a list
- mark uncertain fields explicitly with `needs_review`
- if an attribute is a structured table, list all structure fields
- do not emit JSON

### Prompt Template For Skill C

Task:

Convert the provided normalized ME Markdown into repository JSON schema format.

Requirements:

- output valid JSON only
- preserve field order used by existing schema files
- keep names exactly as defined in the normalized Markdown
- include `structure.fields` when the format is `Array Structure`
- do not invent missing values
- if a required field is missing, stop and report an error instead of guessing

### Prompt Template For Skill D

Task:

Validate a normalized ME Markdown file and its generated JSON artifact.

Requirements:

- compare ME ID, name, actions, instance metadata, attribute count, and attribute order
- verify every structured attribute contains all fields
- verify sizes and formats are consistent
- report exact mismatches
- return `PASS`, `PASS_WITH_WARNINGS`, or `FAIL`

## Recommended Execution Strategy

The safest rollout is incremental.

### Phase 1

Generate only the master ME index.

Success criteria:

- the index contains all ME IDs from the source table
- manual spot checks confirm page references and naming

### Phase 2

Generate Markdown files for a small sample of MEs.

Recommended sample:

- one simple ME
- one ME with table attributes
- one ME with optional attributes

Success criteria:

- Markdown structure stays consistent across the sample
- ambiguous fields are explicitly marked instead of silently guessed

### Phase 3

Generate JSON for the sample set and compare against known-good repository conventions.

Success criteria:

- JSON shape matches existing files in `resource/me/schema`
- array structures and access flags are represented consistently

### Phase 4

Run the full batch generation using the validated pipeline.

## Risk Analysis

### Main risks in direct PDF to JSON generation

1. Table boundaries may be parsed incorrectly.
2. Attribute rows may be merged or split by OCR.
3. Bit lengths and byte lengths may be confused.
4. Structured table fields may be partially lost.
5. Action and access fields may be normalized inconsistently.
6. One extraction error can corrupt the final JSON silently.

### How the recommended design reduces those risks

1. The ME index provides a stable entry point.
2. Per-ME Markdown files isolate extraction errors.
3. JSON generation becomes deterministic after normalization.
4. Validation can run independently of extraction.
5. Human review can focus only on flagged entities.

## Final Recommendation

If the target is long-term maintainability, then the correct interpretation is:

- Solution 2 is the right skill-oriented design
- Solution 1 should only exist as a high-level workflow name

The best implementation model is:

1. one workflow orchestrator
2. one skill for ME index extraction
3. one skill for single-ME extraction
4. one skill for Markdown-to-JSON conversion
5. one skill for validation

If needed later, the workflow can still provide a one-command experience such as:

- ingest G.988 source
- generate or update ME index
- generate missing or changed ME Markdown files
- generate JSON artifacts
- produce a validation report

But internally, it should still be composed of multiple skills rather than a single large prompt.

## Create-Skills Ready Specification

This section converts the design into a format that is directly usable as input for a skill generation tool.

The key rule is:

- generate four independent skills
- do not generate one monolithic skill
- each skill must have a fixed contract
- each skill must fail explicitly when required inputs are missing

## Shared Conventions

These conventions apply to all four skills.

### General rules

- use English in all generated artifacts unless the caller explicitly requests another language
- preserve official G.988 naming exactly when source text is clear
- do not guess missing values
- if source content is ambiguous, return `needs_review`
- keep field order stable
- all outputs must be deterministic for the same input

### Normalization rules

- normalize sizes to bytes when possible
- if the source uses bits for a field inside a structure, keep the original numeric unit only if conversion would lose meaning
- normalize format names to repository conventions where a repository convention already exists
- preserve attribute order exactly as defined in the source
- preserve action order if the source defines it explicitly

### Error policy

Every skill must use explicit failure categories:

- `INVALID_INPUT`: required input is missing or malformed
- `SOURCE_NOT_FOUND`: required source section, page, or document fragment cannot be located
- `AMBIGUOUS_SOURCE`: source exists but cannot be resolved confidently
- `UNSUPPORTED_PATTERN`: source pattern exists but is outside current parser scope
- `VALIDATION_FAILED`: output was produced but does not satisfy the skill contract

### Review markers

When a value is uncertain but partial output is still useful, mark it with one of these tokens:

- `needs_review`
- `missing_in_source`
- `not_parsed`

These tokens must appear only in fields that are explicitly allowed to be uncertain.

## Skill Spec A: ME Index Extractor

### Skill identity

- Skill name: `g988-me-index-extractor`
- Purpose: Extract the full Managed Entity index from G.988 and output the canonical Markdown index file
- Invocation scope: one full source document per run

### Input contract

Required inputs:

- `source_document_path`: path or identifier for the PDF or preprocessed text source

Optional inputs:

- `source_format`: one of `pdf`, `ocr_text`, `markdown`, `plain_text`
- `edition_label`: string label such as `G.988-2024`
- `page_range`: optional page range to narrow extraction

Input example:

```yaml
source_document_path: input/G988.pdf
source_format: pdf
edition_label: G.988
```

### Output contract

The output must be one Markdown document and no JSON.

Required sections:

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

Column rules:

- `ME ID`: integer as shown in source
- `ME Name`: official source name
- `Section`: section identifier if present
- `Page`: source page number if present
- `Status`: `ok` or `needs_review`
- `Notes`: empty when not needed

### Output constraints

- include every ME entry that appears in the source index table
- do not invent missing ME IDs
- preserve row order if the source table order is stable
- if duplicate IDs appear in source, keep both rows and flag them in `Notes`

### Failure rules

- return `INVALID_INPUT` if `source_document_path` is missing
- return `SOURCE_NOT_FOUND` if the source cannot be opened or located
- return `AMBIGUOUS_SOURCE` if no reliable ME index table can be identified
- return `VALIDATION_FAILED` if the Markdown table is malformed or missing required columns

### Minimal acceptance tests

- the document contains the top-level title
- the output contains the `Managed Entities` table
- every row has six columns
- every `ME ID` is numeric unless the row is explicitly marked `needs_review`

### Few-shot example

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

## Skill Spec B: Single ME Extractor

### Skill identity

- Skill name: `g988-single-me-extractor`
- Purpose: Extract one Managed Entity definition from the source and emit normalized Markdown
- Invocation scope: one ME per run

### Input contract

Required inputs:

- `source_document_path`
- `me_id`
- `me_name`

Optional inputs:

- `source_format`
- `source_section`
- `source_page`
- `edition_label`
- `index_row_markdown`: original row from the canonical index

Input example:

```yaml
source_document_path: input/G988.pdf
source_format: pdf
me_id: 347
me_name: IPv6 host config data
source_section: 9.3.x
source_page: 420
```

### Output contract

The output must be one normalized Markdown file with this exact section order:

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

### Output constraints

- preserve ME and attribute names from source spelling when clear
- preserve attribute order exactly
- if an attribute is a structure or table, include `#### Structure Fields`
- if an attribute has no structure, omit the `Structure Fields` subsection for that attribute
- `Actions` must be rendered as a comma-separated list in source order
- `Access` must be rendered as a comma-separated list
- use `needs_review` instead of guessed values

### Failure rules

- return `INVALID_INPUT` if `me_id` or `me_name` is missing
- return `SOURCE_NOT_FOUND` if the target ME section cannot be located
- return `AMBIGUOUS_SOURCE` if multiple candidate sections exist and cannot be disambiguated
- return `UNSUPPORTED_PATTERN` if the ME uses an unsupported layout that cannot be normalized safely
- return `VALIDATION_FAILED` if required sections are missing from the output

### Minimal acceptance tests

- `Identity`, `Classification`, and `Attributes` sections all exist
- `ME ID` and `ME Name` are present
- attributes are numbered sequentially starting from `Attribute 1`
- every attribute contains `Name`, `Size`, `Format`, `Access`, and `Category`

### Few-shot example

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

## Skill Spec C: ME Markdown To JSON

### Skill identity

- Skill name: `g988-me-markdown-to-json`
- Purpose: Convert one normalized ME Markdown file into repository JSON schema format
- Invocation scope: one ME Markdown file per run

### Input contract

Required inputs:

- `me_markdown_content`

Optional inputs:

- `me_markdown_path`
- `target_schema_style`: default `omcianalyzer`
- `strict_mode`: boolean, default `true`

Input example:

```yaml
me_markdown_path: docs/g988/me/347-ipv6-host-config-data.md
target_schema_style: omcianalyzer
strict_mode: true
```

### Output contract

The output must be valid JSON only.

Required top-level shape:

```json
{
  "Me": {
    "id": 0,
    "name": "",
    "access": "",
    "type": "",
    "action": [],
    "instance": {
      "type": "",
      "value": 0
    },
    "attributes": {
      "attribute": []
    }
  }
}
```

Attribute object rules:

- always include `name`, `size`, `format`, `access`, and `category`
- include `structure.fields` only when the format is structural
- each structure field must contain `name`, `size`, and `format`

### Output constraints

- preserve field order used by existing repository schema files
- output integer values as numbers, not strings
- `action` must be a JSON array of strings
- `access` must be a JSON array of strings for each attribute
- do not include extra metadata not used by the repository schema
- if a required Markdown field is absent, fail instead of guessing

### Failure rules

- return `INVALID_INPUT` if no Markdown content is provided
- return `VALIDATION_FAILED` if the Markdown contract is incomplete or malformed
- return `UNSUPPORTED_PATTERN` if a format cannot be mapped to repository JSON safely

### Minimal acceptance tests

- JSON parses successfully
- `.Me.id` is numeric
- `.Me.action` is an array
- `.Me.attributes.attribute` is an array
- every attribute contains required keys

### Few-shot example

Input:

```yaml
me_markdown_path: docs/g988/me/347-ipv6-host-config-data.md
target_schema_style: omcianalyzer
strict_mode: true
```

Expected output snippet:

```json
{
  "Me": {
    "id": 347,
    "name": "IPv6HostConfigData",
    "access": "CreatedByONU",
    "type": "Config",
    "action": [
      "Set",
      "Get"
    ],
    "instance": {
      "type": "Dynamic",
      "value": 0
    },
    "attributes": {
      "attribute": [
        {
          "name": "IPv6 options",
          "size": 1,
          "format": "UnsignedInteger",
          "access": [
            "R",
            "W",
            "Dummy"
          ],
          "category": "Mandatory"
        }
      ]
    }
  }
}
```

## Skill Spec D: ME Artifact Validator

### Skill identity

- Skill name: `g988-me-artifact-validator`
- Purpose: Validate one normalized ME Markdown file and one generated JSON file for consistency
- Invocation scope: one ME pair per run

### Input contract

Required inputs:

- `me_markdown_content`
- `me_json_content`

Optional inputs:

- `me_markdown_path`
- `me_json_path`
- `strict_mode`: boolean, default `true`

Input example:

```yaml
me_markdown_path: docs/g988/me/347-ipv6-host-config-data.md
me_json_path: resource/me/schema/IPv6HostConfigData.json
strict_mode: true
```

### Output contract

The output must be Markdown with this exact section order:

```markdown
# Validation Result

## Summary
- Status:
- ME ID:
- ME Name:

## Checks
| Check | Result | Notes |
|---|---|---|

## Mismatches
- None
```

Allowed status values:

- `PASS`
- `PASS_WITH_WARNINGS`
- `FAIL`

Required checks:

- identity match
- classification match
- action match
- instance metadata match
- attribute count match
- attribute order match
- attribute field match
- structure field match when applicable
- JSON parse validity

### Output constraints

- each failed check must appear in `Checks`
- each mismatch must be described explicitly in `Mismatches`
- if there are no mismatches, the `Mismatches` section must contain `- None`
- do not auto-correct data inside the validator output

### Failure rules

- return `INVALID_INPUT` if Markdown or JSON content is missing
- return `VALIDATION_FAILED` if the JSON cannot be parsed or the Markdown contract is broken

### Minimal acceptance tests

- the output title is `Validation Result`
- `Summary`, `Checks`, and `Mismatches` sections all exist
- the `Status` value is one of the allowed values

### Few-shot example

Input:

```yaml
me_markdown_path: docs/g988/me/347-ipv6-host-config-data.md
me_json_path: resource/me/schema/IPv6HostConfigData.json
strict_mode: true
```

Expected output snippet:

```markdown
# Validation Result

## Summary
- Status: PASS
- ME ID: 347
- ME Name: IPv6HostConfigData

## Checks
| Check | Result | Notes |
|---|---|---|
| identity match | pass | |
| classification match | pass | |
| action match | pass | |
| attribute count match | pass | |

## Mismatches
- None
```

## Direct Create-Skills Input Block

If a skill generation tool accepts concise skill definitions, the following block can be used as the source specification.

### Skill 1

- Name: `g988-me-index-extractor`
- Input: full G.988 source document
- Output: canonical ME index Markdown
- Reject when: source missing, source unreadable, index table not confidently found
- Never do: generate per-ME details, generate JSON

### Skill 2

- Name: `g988-single-me-extractor`
- Input: one ME ID, one ME name, source document, optional section and page
- Output: one normalized ME Markdown file
- Reject when: target ME cannot be identified confidently
- Never do: generate repository JSON directly

### Skill 3

- Name: `g988-me-markdown-to-json`
- Input: one normalized ME Markdown document
- Output: one repository JSON schema file
- Reject when: required Markdown fields are missing or unmappable
- Never do: infer missing semantic values from unrelated MEs

### Skill 4

- Name: `g988-me-artifact-validator`
- Input: one normalized ME Markdown document and one generated JSON document
- Output: validation Markdown report
- Reject when: either artifact is missing or malformed
- Never do: silently correct source data

## Recommended Next Step For Skill Generation

When using create-skills, start with only these two skills first:

1. `g988-me-index-extractor`
2. `g988-single-me-extractor`

After those outputs are stable, generate:

3. `g988-me-markdown-to-json`
4. `g988-me-artifact-validator`

This order keeps the extraction layer stable before schema conversion and validation are introduced.