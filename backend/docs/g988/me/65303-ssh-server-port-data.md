# Managed Entity

## Identity
- ME ID: 65303
- ME Name: SSH Server Port Data
- Source Section: 5.1.2 Example ONU Service Configuration
- Source Page: 44

## Classification
- Access: needs_review
- Type: vendor-specific configuration
- Actions: Set
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: TCP/UDP pointer
- Size: 2 bytes
- Format: pointer
- Access: R, W
- Category: mandatory

### Attribute 2
- Name: SSH server pointer
- Size: 2 bytes
- Format: pointer
- Access: R, W
- Category: mandatory

## Extraction Status
- Status: inferred_from_omci_log_and_ciena_docs
- Attributes found: 2
- Review needed: true
- Validation: Ciena draft ONU requirements names this custom ME as SSH Server Port Data and shows instances 0-3 with tcp_udp_ptr values 7001-7004 and ssh_server_ptr value 0; OMCI log shows Set requests to ME 65303 instances 0-3 immediately after the corresponding TCPUDPConfigData and RS232 port configuration steps; raw OMCI hex confirms attribute 1 carries the 16-bit pointer values 0x1B59-0x1B5C.

## Raw Source

```text
Ciena draft PDF, page 19:
Custom ME SSH Server Port Data

Ciena draft PDF, page 44:
"SshServerPortData": {
  "0": { "tcp_udp_ptr": 7001, "ssh_server_ptr": 0 },
  "1": { "tcp_udp_ptr": 7002, "ssh_server_ptr": 0 },
  "2": { "tcp_udp_ptr": 7003, "ssh_server_ptr": 0 },
  "3": { "tcp_udp_ptr": 7004, "ssh_server_ptr": 0 }
}

OMCI log evidence:
Before ME 65303, the OLT provisions four serial console paths:
- TCPUDPConfigData 7001 + RS232PortOperationConfigurationData 0 + PPTP RS232 UNI 8193
- TCPUDPConfigData 7002 + RS232PortOperationConfigurationData 1 + PPTP RS232 UNI 8194
- TCPUDPConfigData 7003 + RS232PortOperationConfigurationData 2 + PPTP RS232 UNI 8195
- TCPUDPConfigData 7004 + RS232PortOperationConfigurationData 3 + PPTP RS232 UNI 8196

Then the OLT applies ME 65303:
2026-02-26T19:42:45.650852Z Set MeClass=65303 MeInst=0 AttrMask=32768
2026-02-26T19:42:45.820871Z Set MeClass=65303 MeInst=1 AttrMask=32768
2026-02-26T19:42:45.971046Z Set MeClass=65303 MeInst=2 AttrMask=32768
2026-02-26T19:42:46.201901Z Set MeClass=65303 MeInst=3 AttrMask=32768

Raw OMCI hex for those Set requests:
- 005f480aff17000080001b5900... -> class 0xFF17, inst 0, mask 0x8000, attr bytes 0x1B59 = 7001
- 0060480aff17000180001b5a00... -> class 0xFF17, inst 1, mask 0x8000, attr bytes 0x1B5A = 7002
- 0061480aff17000280001b5b00... -> class 0xFF17, inst 2, mask 0x8000, attr bytes 0x1B5B = 7003
- 0062480aff17000380001b5c00... -> class 0xFF17, inst 3, mask 0x8000, attr bytes 0x1B5C = 7004

Interpretation:
- The one-to-one ordering between TCP/UDP port objects 7001-7004 and ME 65303 instances 0-3 matches the draft JSON exactly.
- Raw hex confirms attribute 1 is a 2-byte pointer field carrying the TCP/UDP object instance.
- Attribute 2, SSH server pointer, is documented in the draft JSON and remains at the default singleton server instance 0 in this scenario, but it is not directly exercised in the captured raw Set frames.
```