# Managed Entity

## Identity
- ME ID: 269
- ME Name: VP network CTP
- Source Section: 9.13.9
- Source Page: 460

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Create, Delete, Get, Set
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: VPI value
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 2
- Name: UNI pointer
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 3
- Name: Deprecated 1
- Size: 2 bytes
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
- Name: Deprecated 3
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: optional

### Attribute 6
- Name: Deprecated 4
- Size: 1 byte
- Format: needs_review
- Access: R
- Category: optional

## Extraction Status
- Status: auto_extracted
- Attributes found: 6
- Review needed: false

## Raw Source

```
9.13.9 VP network CTP  NOTE – In [ITU-T G.984.4], this ME is called VP network CTP-G.  This ME represents the termination of VP links on an ONU. It aggregates connectivity functionality  from the network view and alarms from the network element view as well as artefacts from trails.  Instances of this ME are created and deleted by the OLT.  An instance of the VP network CTP ME can be deleted only when no ATM IW VCC TP is associated  with it. It is the responsibility of the OLT to ensure that this condition is met.  Relationships  Zero or more instances of the VP network CTP ME may exist for each instance of the IW  VCC TP ME.  Attributes  Managed entity  ID: This attribute uniquely identifies each instance of this ME. (R ,  set-by-create) (mandatory) (2 bytes)  VPI value: This attribute identifies the VPI value associated with the VP link being  terminated. (R, W, set-by-create) (mandatory) (2 bytes)  UNI pointer: This pointer indicates the xDSL PPTP UNI associated with this VP TP. The  bearer channel may be indicated by the two MSBs of the pointer. (R,  W,  set-by-create) (mandatory) (2 bytes) 

---- page break ---- Direction: This attribute specifies whether the VP link is used for UNI-to-ANI (value 1),  ANI-to-UNI (value  2), or bidirectional (value 3) connection. (R,  W,  set-by-create) (mandatory) (1 byte)  Deprecated 1: Not used; should be set to 0. (R, W, set-by-create) (mandatory) (2 bytes)  Deprecated 2: Not used; should be set to 0. (R, W, set-by-create) (mandatory) (2 bytes)  Deprecated 3: Not used; should be set to 0. (R, W, set-by-create) (optional) (2 bytes)  Deprecated 4: Not used; if present, should be set to 0. (R) (optional) (1 byte)  Actions  Create, delete, get, set  Notifications  Alarm  Alarm  number Alarm Description  0 VP AIS LMIR VP-AIS receiving indication  1 VP RDI LMIR VP-RDI receiving indication  2 VP AIS LMIG VP-AIS generation indication  3 VP RDI LMIG VP-RDI generation indication  4 Segment loss of  continuity  Loss of continuity is detected when the VP network CTP is a  segment end point  5 End-to-end loss of  continuity  Loss of continuity can be detected when the VP network CTP  supports an interworking VCC termination point  6..207 Reserved   208..223 Vendor-specific  alarms  Not to be standardized  
```
