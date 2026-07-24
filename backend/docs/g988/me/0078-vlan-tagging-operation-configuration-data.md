# Managed Entity

## Identity
- ME ID: 78
- ME Name: VLAN tagging operation configuration data
- Source Section: 9.3.12
- Source Page: 155

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Create, Delete, Get, Set
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: Upstream VLAN tag TCI value
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 2
- Name: Downstream VLAN tagging operation mode
- Size: 1 byte
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 3
- Name: Association type
- Size: 1 byte
- Format: needs_review
- Access: R, W, set-by-create
- Category: optional

## Extraction Status
- Status: auto_extracted
- Attributes found: 3
- Review needed: false

## Raw Source

```
9.3.12 VLAN tagging operation configuration data  This ME organizes data associated with VLAN tagging. Instances of this ME are created and deleted  by the OLT.  NOTE 1 – The extended VLAN tagging operation configuration data of clause 9.3.13 is preferred for new  implementations.  Relationships  Zero or one instance of this ME may exist for an instance of any ME that can terminate or  modify an Ethernet stream.  When this ME is associated with a UNI -side TP, it performs its upstream classification and  tagging operations before offering the upstream frame to other filtering, bridging or switching  functions. In the downstream direction, the defined inverse operation is the last operation  performed on the frame before offering it to the UNI-side termination.  When this ME is associated with an ANI-side TP, it performs its upstream classification and  tagging operations as the last step before queu eing for transmission to the OLT, after having  received the upstream frame from other filtering, bridging or switching functions. In the  downstream direction, the defined inverse operation is the first operation performed on the  frame before offering it to possible filter, bridge or switch functions.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. When the  optional association type attribute is 0 or undefined, this attribute's value is the  same as the ID of the ME with which this VLAN tagging operation  configuration data instance is associated, which may be either a PPTP Ethernet  UNI or an IP host config data or an IPv6 host config data ME. Otherwise, the  value of the ME ID is unconstrained except by the need to be  unique. (R, set- by-create) (mandatory) (2 bytes) 

---- page break ---- Upstream VLAN tagging operation mode: This attribute controls upstream VLAN tagging.  Valid values are as follows.  0 Upstream frame is sent as is, regardless of tag.  1 The upstream frame is tagged, whether or not the received frame was  tagged. A tagged frame's TCI is overwritten with the upstream VLAN tag  TCI value. An untagged frame is prepended with a tag whose values are  taken from the upstream VLAN tag TCI value attribute.  2 A tag is prepended to the upstream frame, whether or not the received  frame was tagged. If the received frame is tagged, a second tag (Q -in- Q) is added to the frame. If the received frame is not tagged, a tag is  attached to the frame. The added tag is defined by the upstream VLAN  tag TCI value attribute.  (R, W, set-by-create) (mandatory) (1 byte)  Upstream VLAN tag TCI value : This attribute specifies the TCI for upstream VLAN  tagging. It is used when the upstream VLAN tagging operation mode is 1 or 2.  (R, W, set-by-create) (mandatory) (2 bytes)  Downstream VLAN tagging operation mode : This attribute controls downstream VLAN  tagging. Valid values are as follows.  0 Downstream frame is sent as is, regardless of tag.  1 If the received frame is tagged, the outer tag is stripped. An untagged  frame is forwarded unchanged.  (R, W, set-by-create) (mandatory) (1 byte)  Association type: This attribute specifies the type of ME that is associated with this VLAN  tagging operation configuration data ME. Values are assigned in accordance  with the following list.  0 (Default) Physical path termination point Ethernet UNI (for backward  compatibility, may also be an IP host config data ME; they must not  have the same ME ID). The associated ME instance is implicit; its  identifier is the same as that of this VLAN tagging operation  configuration data.  1 IP host config data or IPv6 host config data  2 IEEE 802.1p mapper service profile  3 MAC bridge port configuration data  4 Physical path termination point xDSL UNI  5 GEM IW termination point  6 Multicast GEM IW termination point  7 Physical path termination point MoCA UNI  8 Reserved  9 Ethernet flow termination point  10 Physical path termination point Ethernet UNI  11 Virtual Ethernet interface point  12 MPLS pseudowire termination point  13 EFM bonding group  The associated ME instance is identified by the associated ME pointer. (R, W,  set-by-create) (optional) (1 byte) 

---- page break ---- Associated ME pointer: When the association type attribute is non-zero, this attribute points  to the ME with which this VLAN tagging operation configuration data ME is  associated. Otherwise, this attribute is undefined, and the association is  implicit through the ME ID. (R, W, set-by-create) (optional) (2 bytes)  NOTE 2 – Implicit association is retained for legacy compatibility. Explicit pointers  are preferred for new implementations.  NOTE 3 – When the association type is xDSL, the two MSBs may be used to indicate  a bearer channel.  Actions  Create, delete, get, set  Notifications  None.  
```
