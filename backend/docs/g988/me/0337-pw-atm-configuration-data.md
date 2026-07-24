# Managed Entity

## Identity
- ME ID: 337
- ME Name: PW ATM configuration data
- Source Section: 9.8.15
- Source Page: 377

## Classification
- Access: needs_review
- Type: needs_review
- Actions: needs_review
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: TP type
- Size: 1 byte
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 2
- Name: PPTP ATM UNI pointer
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 3
- Name: Max cell concatenation
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 4
- Name: Far-end max cell concatenation
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: optional

### Attribute 5
- Name: Timeout mode
- Size: 1 byte
- Format: needs_review
- Access: R, W, set-by-create
- Category: optional

## Extraction Status
- Status: auto_extracted
- Attributes found: 5
- Review needed: true

## Raw Source

```
9.8.15 PW ATM configuration data  This ME contains generic configuration data for an ATM pseudowire . Definitions of attributes are  from PW-ATM-MIB [IETF RFC 5605]. Instances of this ME are created and deleted by the OLT.  Relationships  An instance of this ME is associated with an instance of the MPLS pseudowire TP ME with  a pseudowire type attribute equal to one of the following.  2 ATM AAL5 SDU VCC transport  3 ATM transparent cell transport  9 ATM n-to-one VCC cell transport  10 ATM n-to-one VPC cell transport  12 ATM one-to-one VCC cell mode  13 ATM one-to-one VPC cell mode  14 ATM AAL5 PDU VCC transport 

---- page break ---- Alternatively, an instance of this ME may be associated with an Ethernet flow TP or a  TCP/UDP config data ME, depending on the transport layer of the pseudowire.  Attributes  Managed entity  ID: This attribute uniquely identifies each instance of this ME. (R,  set-by-create) (mandatory) (2 bytes)  TP type: This attribute specifies the type of the underlying transport layer. (R, W,  set-by-create) (mandatory) (1 byte)  0 MPLS pseudowire termination point  1 Ethernet flow termination point  2 TCP/UDP config data  Transport TP pointer: This attribute points to an associated instance of the transport layer  TP, whose type is specified by the TP type attribute . (R, W, set-by-create)  (mandatory) (2 bytes)  PPTP ATM UNI pointer : This attribute points to an associated instance of the ITU -T  G.983.2 PPTP ATM UNI. Refer to [ITU -T G.983.2] for the definition of the  target ME. (R, W, set-by-create) (mandatory) (2 bytes)  Max cell concatenation: This attribute specifies the maximum number of ATM cells that can  be concatenated into one PW packet in the upstream direction . (R, W,  set-by-create) (mandatory) (2 bytes)  Far-end max cell concatenation: This attribute specifies the maximum number of ATM cells  that can be concatenated into one PW packet as provisioned at the far end. This  attribute may be used for error checking of downstream traffic. The value 0  specifies that the ONU uses its internal default. (R, W, set-by-create) (optional)  (2 bytes)  ATM cell loss priority (CLP) QoS mapping: This attribute specifies whether the CLP bits  should be considered when setting the value in the QoS fields of the  encapsulating protocol (e.g., TC fields of the MPLS label stack).  1 ATM CLP bits mapping to QoS fields of the encapsulating protocol  2 Not applicable  The value 0 specifies that the ONU use s its internal default. (R, W,  set-by-create) (optional) (1 byte)  Timeout mode : This attribute specifies whether a packet is transmitted in the upstream  direction based on timeout expiration for collecting cells. The actual handling  of the timeout is implementation specific; as such, this attribute may be  changed at any time with proper consideration of the traffic disruption effect.  1 Disabled. The ONU does not generate packets based on timeout cells.  2 Enabled. The ONU generates packets based on timeout cells.  The value 0 specifies that the ONU use s its internal default. (R, W,  set-by-create) (optional) (1 byte)   PW ATM mapping table : This attribute lists ATM VPI/VCI mapping entries in both the  upstream and downstream directions. In the upstream direction, ATM cells that  match no entry's upstream VPI (and conditionally VCI) values are discarded;  conversely in the downstream direction. Upon ME instantiation, the ONU sets  this attribute to an empty table, which discards all cells in both directions.  The table can contain up to N entries when the pseudowire type is equal to one  of the following:
```
