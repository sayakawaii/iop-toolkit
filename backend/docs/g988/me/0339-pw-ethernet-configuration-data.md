# Managed Entity

## Identity
- ME ID: 339
- ME Name: PW Ethernet configuration data
- Source Section: 9.8.17
- Source Page: 382

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Create, Delete, Get, Set
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: MPLS pseudowire TP pointer
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 2
- Name: TP type
- Size: 1 byte
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 3
- Name: UNI pointer
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

## Extraction Status
- Status: auto_extracted
- Attributes found: 3
- Review needed: false

## Raw Source

```
9.8.17 PW Ethernet configuration data  This ME contains the Ethernet pseudowire configuration data. Instances of this ME are created and  deleted by the OLT.  Relationships  An instance of this ME is associated with an instance of the MPLS pseudowire TP ME with  a pseudowire type attribute equal to the following.  5 Ethernet  Attributes  Managed entity  ID: This attribute uniquely identifies each instance of this ME. (R,  set-by-create) (mandatory) (2 bytes)  MPLS pseudowire TP pointer: This attribute points to an instance of the MPLS pseudowire  TP ME associated with this ME. (R, W, set-by-create) (mandatory) (2 bytes)  TP type: This attribute identifies the type of UNI associated with this Ethernet PW.  Valid values are as follows.  1 Physical path termination point Ethernet UNI  3 IEEE 802.1p mapper service profile  7 Physical path termination point xDSL UNI part 1  11 Virtual Ethernet interface point  12 Physical path termination point MoCA UNI  13 MAC bridge port configuration data  Other values are reserved  (R, W, set-by-create) (mandatory) (1 byte)  UNI pointer: This attribute points to the associated instance of a UNI-side ME. The type of  UNI is determined by the TP type attribute. (R, W, set-by-create) (mandatory)  (2 bytes)  Actions  Create, delete, get, set  Notifications  None.  
```
