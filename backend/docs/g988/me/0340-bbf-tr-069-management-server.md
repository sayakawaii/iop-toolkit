# Managed Entity

## Identity
- ME ID: 340
- ME Name: BBF TR-069 management server
- Source Section: 9.12.16
- Source Page: 440

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Get, Set
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: Administrative state
- Size: 1 byte
- Format: needs_review
- Access: R,W
- Category: mandatory

### Attribute 2
- Name: ACS network address
- Size: 2 bytes
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 3
- Name: Associated tag
- Size: 2 bytes
- Format: needs_review
- Access: R, W
- Category: mandatory

## Extraction Status
- Status: auto_extracted
- Attributes found: 3
- Review needed: false

## Raw Source

```
9.12.16 BBF TR-069 management server  If functions within the ONU are managed by [BBF TR-069], this ME allows OMCI configuration of  the autoconfiguration server (ACS) URL and related authentication information for an ACS  connection initiated by the ONU. [BBF TR-069] supports other means to discover its ACS, so not all  BBF TR-069-compatible ONUs necessarily support this ME. Further more, even if the ONU does  support this ME, some operators may choose not to use it.  An ONU that supports OMCI configuration of ACS information automatically creates instances of  this ME.  Relationships  An instance of the BBF TR -069 management server ME exists for each instance of a BBF  TR-069 management domain within the ONU.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. Through an  identical ID, this ME is implicitly linked to an instance of a VEIP that links to  the BBF TR-069 management domain. (R) (mandatory) (2 bytes)  Administrative state: This attribute locks (1) and unlocks (0) the functions performed by this  ME. When the administrative state is locked, the functions of this ME are  disabled. BBF TR-069 connectivity to an ACS may be possible through means  that do not depend on this ME. The default value of this attribute is locked.  (R,W) (mandatory) (1 byte)   ACS network address : This attribute points to an instance of a network address ME that  contains URL and authentication information associated with the ACS URL.  (R, W) (mandatory) (2 bytes)  Associated tag: This attribute is a TCI value for BBF TR -069 management traffic passing  through the VEIP. A TCI, comprising user priority, CFI and VID, is  represented by 2  bytes. The value 0xFFFF specifies that BBF TR -069  management traffic passes through the VEIP with neither a VLAN nor a  priority tag. (R, W) (mandatory) (2 bytes)  Actions  Get, set 

---- page break ---- 
```
