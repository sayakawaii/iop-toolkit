# Managed Entity

## Identity
- ME ID: 130
- ME Name: IEEE 802.1p mapper service profile
- Source Section: 9.3.10
- Source Page: 150

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Create, Delete, Get, Set
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: TP pointer
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 2
- Name: Interwork TP pointer for P-bit priority 0
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 3
- Name: Interwork TP pointer for P-bit priority 1
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 4
- Name: Interwork TP pointer for P-bit priority 2
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 5
- Name: Interwork TP pointer for P-bit priority 3
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 6
- Name: Interwork TP pointer for P-bit priority 4
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 7
- Name: Interwork TP pointer for P-bit priority 5
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 8
- Name: Interwork TP pointer for P-bit priority 6
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 9
- Name: Unmarked frame option
- Size: 1 byte
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 10
- Name: DSCP to P-bit mapping
- Size: 24 bytes
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 11
- Name: Default P-bit assumption
- Size: 1 byte
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 12
- Name: TP type
- Size: 1 byte
- Format: needs_review
- Access: R, W, set-by-create
- Category: optional

## Extraction Status
- Status: auto_extracted
- Attributes found: 12
- Review needed: false

## Raw Source

```
9.3.10 IEEE 802.1p mapper service profile  This ME associates the priorities of IEEE 802.1p [IEEE 802.1D] priority tagged frames with specific  connections. This ME directs upstream traffic to the designated GEM ports. Downstream traffic  arriving on any of the IEEE 802.1p mapper 's GEM ports is directed to the mapper 's root TP. Other  mechanisms exist to direct downstream traffic, specifically a direct pointer to a downstream queue  from the GEM port network CTP. If such an alternative is used, it should be provisioned to be  consistent with the flow model of the mapper.  Instances of this ME are created and deleted by the OLT.  Relationships  At its root, an instance of this ME may be associated with zero or one instance of a PPTP  UNI, MAC bridge port configuration data, or any type of IW TP ME that carries IEEE 802  traffic. Each of its eight branches is associated with zero or one GEM IW TP.  Attributes  Managed entity  ID: This attribute uniquely identifies each instance of this ME. (R,  set-by-create) (mandatory) (2 bytes)  TP pointer: This attribute points to an instance of the associated TP.  If the optional TP type attribute is not supported, the TP pointer indicates  bridging mapping with the value 0xFFFF; the TP pointer may also point to a  PPTP Ethernet UNI.  The TP type value 0 also indicates bridging mapping, and the TP pointer  should be set to 0xFFFF.  In all other cases, the TP type is determined by the TP type attribute.  (R, W, set-by-create) (mandatory) (2 bytes)  Each of the following eight attributes points to the GEM IW TP associated with the stated P-bit value.  The null pointer 0xFFFF specifies that frames with the associated priority are to be discarded.  Interwork TP pointer for P-bit priority 0: (R, W, set-by-create) (mandatory) (2 bytes)  Interwork TP pointer for P-bit priority 1: (R, W, set-by-create) (mandatory) (2 bytes)  Interwork TP pointer for P-bit priority 2: (R, W, set-by-create) (mandatory) (2 bytes)  Interwork TP pointer for P-bit priority 3: (R, W, set-by-create) (mandatory) (2 bytes)  Interwork TP pointer for P-bit priority 4: (R, W, set-by-create) (mandatory) (2 bytes)  Interwork TP pointer for P-bit priority 5: (R, W, set-by-create) (mandatory) (2 bytes)  Interwork TP pointer for P-bit priority 6: (R, W, set-by-create) (mandatory) (2 bytes) 

---- page break ---- Interwork TP pointer for P-bit priority 7: (R, W, set-by-create) (mandatory) (2 bytes)  Unmarked frame option : This attribute specifies how the ONU should handle untagged  Ethernet frames received across the associated interface. Although it does not  alter the frame in any way, the ONU routes the frame as if it were tagged with  P bits (PCP field) according to the following code points.  0 Derive implied PCP field from DSCP bits of received frame  1 Set implied PCP field to a fixed value specified by the default P -bit  assumption attribute  (R, W, set-by-create) (mandatory) (1 byte)  Untagged downstream frames are passed through the mapper transparently.   DSCP to P-bit mapping: This attribute is valid when the unmarked frame option attribute is  set to 0. The DSCP to P-bit attribute can be considered a bit string sequence of  64 3 bit groupings. The 64 sequence entries represent the possible values of  the 6 bit DSCP field. Each 3 bit grouping specifies the P-bit value to which the  associated DSCP value should be mapped. The unmarked frame is then  directed to the GEM IW TP indicated by the interwork TP pointer mappings.  (R, W) (mandatory) (24 bytes)  NOTE – If certain bits in the DSCP field are to be ignored in the mapping process,  the attribute should be provisioned such that all possible values of those bits produce  the same P-bit mapping. This can be applied to the case where instead of full DSCP,  the operator wishes to adopt the priority mechanism based on IP precedence, which  needs only the three MSBs of the DSCP field.  Default P-bit assumption: This attribute is valid when the unmarked frame option attribute  is set to 1. In its LSBs, the default P -bit assumption attribute contains the  default PCP field to be assumed. The unmodified frame is then directed to the  GEM IW TP indicated by the interwork TP pointer mappings. (R, W,  set-by-create) (mandatory) (1 byte)  TP type: This attribute identifies the type of TP associated with the mapper.  0 Mapper used for bridging-mapping  1 Mapper directly associated with a PPTP Ethernet UNI  2 Mapper directly associated with an IP host config data or IPv6 host  config data ME  3 Mapper directly associated with an Ethernet flow termination point  4 Mapper directly associated with a PPTP xDSL UNI  5 Reserved  6 Mapper directly associated with a PPTP MoCA UNI  7 Mapper directly associated with a virtual Ethernet interface point  8 Mapper directly associated with an IW VCC termination point  9 Mapper directly associated with an EFM bonding group  (R, W, set-by-create) (optional) (1 byte)  Actions  Create, delete, get, set  Notifications  None. 

---- page break ---- 
```
