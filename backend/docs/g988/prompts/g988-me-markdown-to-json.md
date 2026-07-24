# Skill Spec: g988-me-markdown-to-json

## Purpose

Convert one normalized ME Markdown artifact into the repository JSON schema format used by this project.

## Scope

- One ME Markdown document per run
- Deterministic conversion only
- Do not invent missing semantic values

## Required Inputs

- `me_markdown_content`

## Optional Inputs

- `me_markdown_path`
- `target_schema_style`: default `omcianalyzer`
- `strict_mode`: default `true`

## Input Example

```yaml
me_markdown_path: docs/g988/me/347-ipv6-host-config-data.md
target_schema_style: omcianalyzer
strict_mode: true
```

## Output Contract

Output valid JSON only.

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

## Output Rules

- Preserve field order used by existing repository schema files
- Output integers as numbers, not strings
- `action` must be an array of strings
- attribute `access` must be an array of strings
- Include `structure.fields` only when the attribute format is structural
- Fail instead of guessing if required Markdown fields are missing
- Do not include extra metadata not used by the repository schema

## Failure Policy

Use one of these explicit failures when output cannot be produced correctly:

- `INVALID_INPUT`: Markdown content is missing
- `VALIDATION_FAILED`: Markdown contract is incomplete or malformed
- `UNSUPPORTED_PATTERN`: format cannot be mapped safely to repository JSON

## Acceptance Criteria

- Output parses as valid JSON
- `.Me.id` is numeric
- `.Me.action` is an array
- `.Me.attributes.attribute` is an array
- Every attribute contains `name`, `size`, `format`, `access`, and `category`

## Few-Shot Example

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

## Direct Create-Skills Prompt

Create a skill named `g988-me-markdown-to-json`.

Its only purpose is to convert one normalized ME Markdown artifact into the repository JSON schema format used by this project.

Required input:

- `me_markdown_content`

Optional inputs:

- `me_markdown_path`
- `target_schema_style`
- `strict_mode`

The output must be valid JSON only with this top-level shape:

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

Behavior rules:

- preserve repository field order
- use arrays for `action` and attribute `access`
- include `structure.fields` only when required
- fail instead of guessing missing values
- never add extra metadata

Failure categories:

- `INVALID_INPUT`
- `VALIDATION_FAILED`
- `UNSUPPORTED_PATTERN`