# Managed Entity

## Identity
- ME ID: 333
- ME Name: MPLS pseudowire termination point
- Source Section: 9.8.14
- Source Page: 375

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Create, Delete, Get, Set
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
- Name: TP pointer
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 3
- Name: MPLS label indicator
- Size: 1 byte
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 4
- Name: MPLS PW direction
- Size: 1 byte
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 5
- Name: MPLS PW uplink label
- Size: 4 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 6
- Name: MPLS PW downlink label
- Size: 4 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 7
- Name: MPLS PW TC
- Size: 1 byte
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 8
- Name: MPLS tunnel direction
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 9
- Name: Pseudowire control word preference
- Size: 1 byte
- Format: needs_review
- Access: R, W, set-by-create
- Category: optional

### Attribute 10
- Name: Administrative state
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 11
- Name: Operational state
- Size: 1 byte
- Format: needs_review
- Access: R
- Category: optional

## Extraction Status
- Status: auto_extracted
- Attributes found: 11
- Review needed: false

## Raw Source

```
9.8.14 MPLS pseudowire termination point  This ME contains the configuration data of a pseudowire whose underlying transport method is  MPLS. Instances of this ME are created and deleted by the OLT.  Relationships  Zero or one instance of this ME is associated with each instance of the pseudowire TP ME.  Attributes  Managed entity  ID: This attribute uniquely identifies each instance of this ME. (R,  set-by-create) (mandatory) (2 bytes)  TP type: This attribute specifies the type of ANI-side TP associated with this ME.   1 Ethernet flow termination point  2 GEM IW TP  3 TCP/UDP config data  4 MPLS pseudowire termination point 

---- page break ---- NOTE – If this instance of the MPLS PW TP is pointed to by another instance of the  MPLS PW TP (i .e., whose TP type  = 4), this instance represents a tunne lled MPLS  flow, and the following attributes are not meaningful: MPLS PW direction ; MPLS  PW uplink label ; MPLS PW downlink label ; and MPLS PW TC. These attributes  should be set to the proper number of 0x00 bytes by the OLT and ignored by the ONU.  (R, W, set-by-create) (mandatory) (1 byte)  TP pointer: This attribute points to the instance of the TP associated with this MPLS PW  TP. The type of the associated TP is determined by the TP type attribute. (R,  W, set-by-create) (mandatory) (2 bytes)  MPLS label indicator: This attribute specifies the MPLS label stacking situation.  0 Single MPLS labelled  1 Double MPLS labelled  (R, W, set-by-create) (mandatory) (1 byte)  MPLS PW direction: This attribute specifies the inner MPLS direction.  0 Upstream only  1 Downstream only  2 Bidirectional  (R, W, set-by-create) (mandatory) (1 byte)  MPLS PW uplink label: This attribute specifies the label  of the inner MPLS pseudowire   upstream. The attribute is not meaningful for unidirectional downstream PWs.  (R, W, set-by-create) (mandatory) (4 bytes)  MPLS PW downlink label: This attribute specifies the label of the inner MPLS pseudowire  downstream. The attribute is not meaningful for unidirectional upstream PWs.  (R, W, set-by-create) (mandatory) (4 bytes)  MPLS PW TC: This attribute specifies the inner MPLS TC value in the upstream direction.  The attribute is not meaningful for unidirectional downstream PWs . (R, W,  set-by-create) (mandatory) (1 byte)  NOTE 1 – The TC field was previously known as EXP. Refer to [b-IETF RFC 5462].  MPLS tunnel direction: This attribute specifies the direction of the (outer) MPLS tunnel.  0 Upstream only  1 Downstream only  2 Bidirectional  (R, W, set-by-create) (mandatory for double-labelled case) (1 byte)  MPLS tunnel uplink label: This attribute specifies the (outer) label for the upstream MPLS  tunnel. If the MPLS tunnel is downstream only, this attribute should be set to  0. (R, W, set-by-create) (mandatory for double-labelled case) (4 bytes)  MPLS tunnel downlink label: This attribute specifies the (outer) label  for the downstream  MPLS tunnel. If the MPLS tunnel is upstream only, this attribute should be set  to 0. (R, W, set-by-create) (mandatory for double-labelled case) (4 bytes)  MPLS tunnel TC: This attribute specifies the TC value of the upstream MPLS tunnel. If the  MPLS tunnel is downstream only , this attribute should be set to 0 . (R, W,  set-by-create) (mandatory for double MPLS labelled case) (1 byte)  NOTE 2 – The TC field was previously known as EXP. Refer to [b-IETF RFC 5462].  Pseudowire type: This attribute specifies the emulated service to be carried over this PW.   The values are from [IETF RFC 4446].  2 ATM AAL5 SDU VCC transport  

---- page break ---- 3 ATM transparent cell transport   5 Ethernet   9 ATM n-to-one VCC cell transport   10 ATM n-to-one VPC cell transport   12 ATM one-to-one VCC cell mode   13 ATM one-to-one VPC cell mode  14 ATM AAL5 PDU VCC transport   All other values are reserved.  (R, W, set-by-create) (mandatory) (2 bytes)  Pseudowire control word preference: When set to true, this Boolean attribute specifies that  a control word is to be sent with each packet. Some PW types mandate the use  of a control word in any event. In such cases, the value configured for this  attribute has no effect on the presence of the control word . (R, W,  set-by-create) (optional) (1 byte)  Administrative state: This attribute locks (1) and unlocks (0) the functions performed by the  MPLS pseudowire TP. Administrative state is further described in  clause A.1.6. (R, W) (optional) (1 byte)  Operational state: This attribute reports whether the ME is currently capable of performing  its function. Valid values are enabled (0) and disabled (1). (R) (optional)  (1 byte)  Actions  Create, delete, get, set  Notifications  Attribute value change  Number Attribute value change Description  1..14 N/A   15 Op state Operational state  16 Reserved   
```
