# Managed Entity

## Identity
- ME ID: 302
- ME Name: Dot1ag MEP
- Source Section: 9.3.22
- Source Page: 185

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Create, Delete, Get, Set, Test
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: Layer 2 entity pointer
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 2
- Name: Layer 2 type
- Size: 1 byte
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 3
- Name: MA pointer
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 4
- Name: MEP ID
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 5
- Name: MEP control
- Size: 1 byte
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 6
- Name: Primary VLAN
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 7
- Name: Administrative state
- Size: 1 byte
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 8
- Name: CCM and LTM priority
- Size: 1 byte
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 9
- Name: Egress identifier
- Size: 8 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 10
- Name: Peer MEP IDs
- Size: 24 bytes
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 11
- Name: Alarm declaration soak time
- Size: 2 bytes
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 12
- Name: Alarm clear soak time
- Size: 2 bytes
- Format: needs_review
- Access: R, W
- Category: mandatory

## Extraction Status
- Status: auto_extracted
- Attributes found: 12
- Review needed: false

## Raw Source

```
9.3.22 Dot1ag MEP  This ME models an MEP as defined primarily in [IEEE 802.1ag] and secondarily in [ITU-T Y.1731].  It is created and deleted by the OLT. An MEP exists at one of eight possible maintenance levels, and  resides at the boundary of a MD. It inherits a name, and optionally a set of associated VLANs, from  its associated MA.  Relationships  One or more MEPs may be associated with a MAC bridge port or an IEEE 802.1p mapper in  the absence of a MAC bridge. A n MEP is also associated with zero or more VLANs and a n  MA.  Attributes  Managed entity  ID: This attribute uniquely identifies each instance of this ME. (R,  set-by-create) (mandatory) (2 bytes)  Layer 2 entity pointer : Depending on the value of the layer 2 type attribute, this pointer  specifies the MAC bridge port configuration data ME or the IEEE 802.1p  mapper service profile ME with which this MEP is associated. (R,  W,  set-by-create) (mandatory) (2 bytes)  Layer 2 type: This attribute specifies whether the MA is associated with a MAC bridge port  (value 0) or an IEEE 802.1p mapper (value 1). (R,  W, set-by-create)  (mandatory) (1 byte)  MA pointer: This pointer specifies the maintenance association with which this MEP is  associated. (R, W, set-by-create) (mandatory) (2 bytes)  MEP ID: This attribute specifies the MEP 's own identity in the MA. For a given MA,  the MEP ID must be unique throughout the network defined by the MD. The  MEP ID is defined in the range 1..8191. The value 0 indicates that no MEP ID  is (yet) configured. (R, W, set-by-create) (mandatory) (2 bytes)  MEP control: This attribute specifies some of the overall behavioural aspects of the MEP. It  is interpreted as follows. 

---- page break ----   Bit Interpretation when bit value = 1  1 (LSB) Reserved  2 MEP generates CCMs  3 Enable ITU-T Y.1731 server MEP function  4 Enable generation of Ethernet AIS  5 This is an up MEP, facing toward the core of the bridge. If more  than one MEP exists on a given maintenance association and on a  given bridge, all such MEPs must face the same direction.  6..8 Reserved  (R, W, set-by-create) (mandatory) (1 byte)  Primary VLAN: This attribute is a 12  bit VLAN ID. The value 0 indicates that the MEP  inherits its primary VLAN from its parent MA. CFM messages, except  forwarded LTMs, are tagged with the primary VLAN ID. If explicitly specified,  the value of this attribute must be one of the VLANs associated with the parent  MA. (R, W, set-by-create) (mandatory) (2 bytes)  Administrative state: This attribute locks (1) and unlocks (0) the functions performed by this  ME. Administrative state is further described in clause A.1.6. (R,  W,  set-by-create) (mandatory) (1 byte)  CCM and LTM priority : Ranging from 0..7, this attribute permits CCM and LTM frames  to be explicitly prioritized, which may be needed if flows are separated,  e.g.,  by 802.1p priority. The priority specified in this attribute is also used in  linktrace reply (LTR) frames originated by this MEP. The value 0xFF selects  the IEEE 802.1ag default, whereby CCM and LTM frames are transmitted with  the highest Ethernet priority available. (R,  W, set-by-create) (mandatory)  (1 byte)  Egress identifier : This attribute comprises 8  bytes to be included in LTMs. They allow  received LTRs to be directed to the correct originator. The attribute includes  the originator MAC address and a locally  defined identifier. If this field is 0,  the ONU uses the MEP's MAC address, with 0 as the locally defined identifier.  (R, W, set-by-create) (mandatory) (8 bytes)  Peer MEP IDs: This attribute lists the expected peer MEPs for CCMs, 2  bytes per MEP ID.  [IEEE 802.1ag] allows for multipoint networks, and therefore a list of peer  MEPs. This attribute allows for up to 12 peers for a given MEP, though G-PON  applications are expected to need only a single peer. Missing or unexpected  messages trigger alarm declaration after a soak interval . Unused peer MEP  slots should be set to 0. (R, W) (mandatory) (24 bytes) 

---- page break ---- ETH AIS control: This attribute controls the generation of Ethernet alarm indication signal  (AIS) frames when they are enabled through the MEP control attribute. It is  interpreted as follows:    Bit Interpretation  1 (LSB) Transmission period   0: once per second   1: once per minute  2..4 P-bit priority of transmitted ETH AIS frames  5..7 The maintenance level at which the client MEP exists  8 Reserved  (R, W, set-by-create) (mandatory if ETH AIS is enabled) (1 byte)  Fault alarm threshold : This attribute specifies the lowest priority alarm that is allowed to  generate a fault alarm. The value 0 specifies that the ONU use s its internal  default. It is defined as follows.  1 All defects generate alarms after suitable soaking, including AIS and  RDICCM.  2 Alarm generated only by one of: MACstatus, RemoteCCM, ErrorCCM,  XconCCM. This value is recommended as the default in [IEEE  802.1ag].  3 Alarm generated only by one of: RemoteCCM, ErrorCCM, XconCCM.  4 Alarm generated only by one of: ErrorCCM, XconCCM.  5 Alarm generated only by: XconCCM.  6 No alarms are to be reported. This setting may be useful during  configuration of services across the network when spurious alarms  could otherwise be generated.  (R, W, set-by-create) (optional) (1 byte)  Alarm declaration soak time : This attribute defines the defect soak time that must elapse  before the MEP declares an alarm. It is expressed in 10 ms units with a range  of 250 to 1000, i.e.,  2.5 s to 10  s. The default is recommended to be 2.5  seconds. (R, W) (mandatory) (2 bytes)  Alarm clear soak time : This attribute defines the defect -free soak time that must elapse  before the MEP clears an alarm . It is expressed in intervals of 10  ms with a  range of 250 to 1 000, i.e., 2.5 s to 10 s. The default is recommended to be 10 s.  (R, W) (mandatory) (2 bytes)  Actions  Create, delete, get, set  Test: The test operation causes the MEP to originate one or more loopback messages  (
```
