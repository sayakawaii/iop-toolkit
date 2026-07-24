# Managed Entity

## Identity
- ME ID: 65302
- ME Name: SSH Server Operation
- Source Section: 5.1.2 Example ONU Service Configuration
- Source Page: 44

## Classification
- Access: needs_review
- Type: vendor-specific configuration
- Actions: Set, AVC
- Instance Type: singleton
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: Server action
- Size: 1 byte
- Format: enum
- Access: W
- Category: mandatory

### Attribute 2
- Name: Server operational state (inferred)
- Size: needs_review
- Format: state_or_bitfield
- Access: R
- Category: mandatory

## Extraction Status
- Status: inferred_from_omci_log_and_ciena_docs
- Attributes found: 2
- Review needed: true
- Validation: Ciena draft ONU requirements names this custom ME as SSH Server Operation and shows server_action=1; OMCI log shows Set on attribute mask 32768 followed by AVC on attribute mask 16384; raw OMCI hex confirms the Set payload starts with 0x01 and the AVC payload carries changing bytes 0x20/0x80 in the corresponding attribute slot, implying a second AVC-capable status attribute.

## Raw Source

```text
Ciena draft PDF, page 19:
Custom ME SSH Server Operation

Ciena draft PDF, page 44:
"SshServerOperation": {
  "0": {
    "server_action": 1
  }
}

OMCI log evidence:
2026-02-26T19:42:46.327142Z Set MeClass=65302 MeInst=0 AttrMask=32768
2026-02-26T19:42:46.428949Z Set response Result=0
2026-02-26T19:40:02.651741Z AVC MeClass=65302 MeInst=0 AttrMask=16384
2026-02-26T19:42:47.067623Z AVC MeClass=65302 MeInst=0 AttrMask=16384

Raw OMCI hex:
- 0063480aff1600008000010000... -> class 0xFF16, inst 0, mask 0x8000, first attr byte 0x01
- 0000110aff1600004000200000... -> AVC for class 0xFF16, inst 0, mask 0x4000, first returned bytes 0x20 0x00
- 0000110aff1600004000800000... -> AVC for class 0xFF16, inst 0, mask 0x4000, first returned bytes 0x80 0x00

Interpretation:
- The official example exposes one configured field, server_action, with value 1.
- The live OMCI trace shows the OLT setting only attribute 1 to value 1, then receiving AVC notifications on attribute 2.
- Raw hex confirms attribute 1 is directly driven by a value byte of 0x01. This matches the draft JSON field server_action=1.
- Attribute 2 is not named in the public draft. The raw AVC payload changes from 0x20 to 0x80 in the attribute-2 slot, so this field is at least a read-only state or event bitfield, but its exact width and semantic mapping remain unresolved.
```