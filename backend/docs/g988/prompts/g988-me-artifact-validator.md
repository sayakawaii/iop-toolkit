# Skill Spec: g988-me-artifact-validator

## Purpose

Validate one normalized ME Markdown artifact and one generated JSON artifact for internal consistency.

## Scope

- One ME Markdown and JSON pair per run
- Validation only
- Do not mutate either artifact
- Do not auto-correct source data

## Required Inputs

- `me_markdown_content`
- `me_json_content`

## Optional Inputs

- `me_markdown_path`
- `me_json_path`
- `strict_mode`: default `true`

## Input Example

```yaml
me_markdown_path: docs/g988/me/347-ipv6-host-config-data.md
me_json_path: resource/me/schema/IPv6HostConfigData.json
strict_mode: true
```

## Output Contract

Output one Markdown validation report only.

Required structure:

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

## Required Checks

- identity match
- classification match
- action match
- instance metadata match
- attribute count match
- attribute order match
- attribute field match
- structure field match when applicable
- JSON parse validity

## Output Rules

- Every failed check must appear in the `Checks` table
- Every mismatch must be described explicitly in `Mismatches`
- If there are no mismatches, `Mismatches` must contain `- None`
- Do not auto-correct or rewrite the source artifacts

## Failure Policy

Use one of these explicit failures when output cannot be produced correctly:

- `INVALID_INPUT`: Markdown or JSON content is missing
- `VALIDATION_FAILED`: JSON cannot be parsed or Markdown contract is broken

## Acceptance Criteria

- The title is `Validation Result`
- `Summary`, `Checks`, and `Mismatches` sections all exist
- `Status` is `PASS`, `PASS_WITH_WARNINGS`, or `FAIL`
- The `Checks` table exists

## Few-Shot Example

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

## Direct Create-Skills Prompt

Create a skill named `g988-me-artifact-validator`.

Its only purpose is to validate one normalized ME Markdown artifact and one generated JSON artifact for internal consistency.

Required inputs:

- `me_markdown_content`
- `me_json_content`

Optional inputs:

- `me_markdown_path`
- `me_json_path`
- `strict_mode`

The output must be one Markdown report with this exact structure:

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

Behavior rules:

- check identity, classification, actions, instance metadata, attribute count, attribute order, attribute fields, and structure fields when applicable
- use `PASS`, `PASS_WITH_WARNINGS`, or `FAIL`
- describe each mismatch explicitly
- never auto-correct source data

Failure categories:

- `INVALID_INPUT`
- `VALIDATION_FAILED`