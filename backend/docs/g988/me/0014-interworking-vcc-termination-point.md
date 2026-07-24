# Managed Entity

## Identity
- ME ID: 14
- ME Name: Interworking VCC termination point
- Source Section: 9.13.4
- Source Page: 457

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Create, Delete, Get, Set
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: VCI value
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 2
- Name: VP network CTP connectivity pointer
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 3
- Name: Deprecated 1
- Size: 1 byte
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 4
- Name: Deprecated 2
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 5
- Name: AAL5 profile pointer
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 6
- Name: Deprecated 3
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 7
- Name: AAL loopback configuration
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 8
- Name: PPTP counter
- Size: 1 byte
- Format: needs_review
- Access: R
- Category: optional

## Extraction Status
- Status: auto_extracted
- Attributes found: 8
- Review needed: false

## Raw Source

```
Actions  Get, set  Notifications  None.  9.13.4 Interworking VCC termination point  An instance of this ME represents a point in the ONU where the IW of a service or underlying physical  infrastructure (e.g., ADSL) to an ATM layer takes place. At this point, ATM cells are generated from  a bit stream (e.g., Ethernet) or a bit stream is reconstructed from ATM cells.  Instances of this ME are created and deleted by the OLT.  Relationships  One instance of this ME exists for each occurrence of transformation of a data stream into  ATM cells and vice versa.  Attributes  Managed entity  ID: This attribute uniquely identifies each instance of this ME.  (R, set-by-create) (mandatory) (2 bytes)  VCI value: This attribute identifies the VCI value associated with this IW VCC TP. (R, W,  set-by-create) (mandatory) (2 bytes)  VP network CTP connectivity pointer : This attribute points to the VP network CTP  associated with this IW VCC TP. (R, W, set-by-create) (mandatory) (2 bytes)  Deprecated 1: Not used; should be set to 0. (R, W, set-by-create) (mandatory) (1 byte)  Deprecated 2: Not used; should be set to 0. (R, W, set-by-create) (mandatory) (2 bytes)  AAL5 profile pointer : This attribute points to an instance of the AAL5 profile. (R,  W,  set-by-create) (mandatory) (2 bytes)  Deprecated 3: Not used; should be set to 0. (R, W, set-by-create) (mandatory) (2 bytes)  AAL loopback configuration: This attribute sets the ATM loopback configuration. All code  points are retained for backward compatibility, but some are not expected to  be needed in current and future applications.  0 No loopback  1 Loopback 1, loopback of downstream traffic before FEC of AAL1  2 Loopback 2, loopback of downstream traffic after FEC of AAL1  3 Loopback after AAL, loopback of downstream traffic after any AAL.  Loopback after AAL is depicted in Figure 9.13.4-1.  G.988(12)_F9.13.4-1 Ethernet over GEM ONU xDSL UNI Loopback after AAL PONATM interworking   Figure 9.13.4-1 – AAL loopback configuration  The default value of this attribute is 0. (R, W) (mandatory) (1 byte)  PPTP counter : This value is the number of instances of PPTP MEs associated with this  instance of the IW VCC TP. (R) (optional) (1 byte) 

---- page break ---- Operational state : This attribute indicates whether the ME is capable of performing its  function. Valid values are enabled (0) and disabled (1). (R) (optional) (1 byte)  Actions  Create, delete, get, set  Notifications  Attribute value change  Number Attribute value change Description  1..8 N/A   9 Op state Operational state change  10..16 Reserved     Alarm  Alarm  number Alarm Description  0 End-to-end VC AIS layer  management indication  receiving (LMIR)  End-to-end VC-AIS receiving indication (optional)  1 End-to-end VC RDI LMIR End-to-end VC-RDI receiving indication (optional)  2 End-to-end VC AIS layer  management indication  generation (LMIG)  End-to-end VC-AIS generation indication (optional)  3 End-to-end VC RDI LMIG End-to-end VC-RDI generation indication (optional)  4 Segment loss of continuity Loss of continuity detected when the interworking  VCC termination point is a segment end point  (optional)  5 End-to-end loss of continuity Loss of continuity detected at the interworking VCC  termination point (optional)  6 CSA Cell starvation alarm  7..207 Reserved   208..223 Vendor-specific alarms Not to be standardized  
```
