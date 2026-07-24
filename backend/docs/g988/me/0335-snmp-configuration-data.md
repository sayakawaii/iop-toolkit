# Managed Entity

## Identity
- ME ID: 335
- ME Name: SNMP configuration data
- Source Section: 9.12.15
- Source Page: 439

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Create, Delete, Set, Get
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: SNMP version
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 2
- Name: SNMP agent address
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 3
- Name: SNMP server address
- Size: 4 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 4
- Name: SNMP server port
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 5
- Name: Security name pointer
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 6
- Name: Community for read
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 7
- Name: Community for write
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 8
- Name: Sys name pointer
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

## Extraction Status
- Status: auto_extracted
- Attributes found: 8
- Review needed: false

## Raw Source

```
9.12.15 SNMP configuration data  The SNMP configuration data ME provides a way for the OLT to provision an IP path for an SNMP  management agent.  The SNMP configuration data ME is created and deleted by the OLT.  Relationships  One instance of this ME is created by the OLT for each SNMP management path termination.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. The ME IDs  0 and 0xFFFF are reserved. (R, set-by-create) (mandatory) (2 bytes)  SNMP version: This integer attribute is the SNMP protocol version to be supported. (R, W,  set-by-create) (mandatory) (2 bytes)  SNMP agent address : This attribute is a pointer to a TCP/UDP config data ME, which  provides the SNMP agent. (R, W, set-by-create) (mandatory) (2 bytes)  SNMP server address : This attribute is the IP address of the SNMP server. (R, W,  set-by-create) (mandatory) (4 bytes)  SNMP server port : This attribute is the UDP port number of the SNMP server. (R, W,  set-by-create) (mandatory) (2 bytes)  Security name pointer : This attribute points to a large string whose content represents the  SNMP security name in a human -readable format that is independent of the 

---- page break ---- security model. SecurityName is defined in [b -IETF RFC 2571]. (R, W,  set-by-create) (mandatory) (2 bytes)  Community for read: This attribute is a pointer to a large string that contains the name of  the read community. (R, W, set-by-create) (mandatory) (2 bytes)  Community for write: This attribute is a pointer to a large string that contains the name of  the write community. (R, W, set-by-create) (mandatory) (2 bytes)  Sys name pointer: This attribute points to a large string whose content identifies the SNMP  system name. SysName is defined in [b -IETF RFC  3418]. (R, W,  set-by-create) (mandatory) (2 bytes)  Actions  Create, delete, Set, get  Notifications  None.  
```
