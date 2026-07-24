# Managed Entity

## Identity
- ME ID: 171
- ME Name: Extended VLAN tagging operation configuration data
- Source Section: 9.3.13
- Source Page: 157

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Create, Delete, Get, Set
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: Association type
- Size: 1 byte
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 2
- Name: Received frame VLAN tagging operation table max size
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 3
- Name: Input TPID
- Size: 2 bytes
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 4
- Name: Output TPID
- Size: 2 bytes
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 5
- Name: Downstream mode
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 6
- Name: Received frame VLAN tagging operation table
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 7
- Name: DSCP to P-bit mapping
- Size: 24 bytes
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 8
- Name: Enhanced mode
- Size: 1 byte
- Format: needs_review
- Access: R, Set-by-create
- Category: optional

## Extraction Status
- Status: auto_extracted
- Attributes found: 8
- Review needed: false

## Raw Source

```
9.3.13 Extended VLAN tagging operation configuration data  This ME organizes data associated with VLAN classification and tagging operations. Regardless of  its point of attachment, the specified tagging operations refer to the upstream direction.  Instances of  this ME are created and deleted by the OLT.  Through separate attributes, this ME supports either a Received frame VLAN tagging operation table  attribute in its backward compatible form, or an enhanced frame classification and processing  capability. The OLT can determine whether the ONU supports the enhanced capability through the  Enhanced mode attribute of the ONU3-G ME.  Relationships  Zero or one instance of this ME may exist for an instance of any ME that can terminate or  modify an Ethernet stream.  When this ME is associated with a UNI -side TP, it performs its upstream classification and  tagging operations before offering the upstream frame to other filtering, bridging or switching  functions. In the downstream direction, the defined inverse operation is the last operation  performed on the frame before offering it to the UNI-side termination.  When this ME is associated with an ANI-side TP, it performs its upstream classification and  tagging operations as the last step before transmission to the OLT, after having received the  upstream frame from other filtering, bridging or switching functions. In the downstream  direction, the defined inver se operation is the first operation performed on the frame before  offering it to possible filter, bridge or switch functions.  Attributes  Managed entity ID: This attribute provides a unique number for each instance of this ME.  (R, set-by-create) (mandatory) (2 bytes)  Association type: This attribute identifies the type of the ME associated with this extended  VLAN tagging ME. Values are assigned as follows.  0 MAC bridge port configuration data  1 IEEE 802.1p mapper service profile  2 Physical path termination point Ethernet UNI  3 IP host config data or IPv6 host config data  4 Physical path termination point xDSL UNI  5 GEM IW termination point  6 Multicast GEM IW termination point  7 Physical path termination point MoCA UNI 

---- page break ---- 8 Reserved  9 Ethernet flow termination point  10 Virtual Ethernet interface point  11 MPLS pseudowire termination point  12 EFM bonding group  (R, W, set-by-create) (mandatory) (1 byte)  NOTE 1 – If a MAC bridge is configured, code points 1, 5 , 6 and 11 are associated  with the ANI side of the MAC bridge. Code point 0 is associated with the ANI or UNI  side, depending on the location of the MAC bridge port. The other code points are  associated with the UNI side.  When the extended VLAN tagging ME is associated with the ANI side, it behaves as  an upstream egress rule, and as a downstream ingress rule when the downstream mode  attribute is equal to 0. When the extended VLAN tagging ME is associated with the  UNI side, the extended VLAN tagging ME behaves as an upstream ingress rule, and  as a downstream egress rule when the downstream mode attribute is equal to 0.  Received frame VLAN tagging operation table max size : This attribute indicates the  maximum number of entries that can be set in the received frame VLAN  tagging operation table. (R) (mandatory) (2 bytes)  Input TPID: This attribute gives the special TPID value for operations on the input  (filtering) side of the table. Typical values include 0x88A8 and 0x9100. (R, W)  (mandatory) (2 bytes)  Output TPID : This attribute gives the special TPID value for operations on the output  (tagging) side of the table. Typical values include 0x88A8 and 0x9100. (R, W)  (mandatory) (2 bytes)  Downstream mode: Regardless of its association, the rules of the received frame VLAN  tagging operation table attribute pertain to upstream traffic. The downstream  mode attribute defines the tagging action to be applied to downstream frames.  The received frame VLAN tagging operation table installs defaults upstream  rules. In the downstream direction, the upstream default rules with the default  treatment do not apply. It should be noted that downstream frame treatment is  defined by the downstream mode attribute and is not affected by the upstream  default rules.  The received frame VLAN tagging operation table can result in two types of  rule mappings:  • One to one mapping: A table contains one or more rules that result in  unique mappings between the ingress and egress flows.  • Many to one mapping: A table contains more than one rule that results in  the same ANI-side tag configuration.  For one-to-one mappings, the inverse operation to apply in the downstream  direction (in the case of bidirectional flows) is the inverse operation of the  upstream rule.  Many-to-one mappings are possible however, and these are treated as  follows.  • If an upstream many-to-one mapping results from multiple operation  rules producing the same ANI-side tag configuration, then the first  matching rule in the list defines the inverse operation. The meaning of  match depends on the value of the downstream mode attribute. 

---- page break ---- • If the many-to-one mapping results from "don't care" fields in the filter  being replaced with provisioned fields in the ANI side tags, then the  inverse is defined to set the corresponding fields on the ANI side to their  lowest legal value.  If the upstream rule merely copies (i.e., no explicit value is specified in the  filter field) an inbound tag value to an outbound tag value, the comparison in  the downstream direction applies to all tag values. This applies separately to  the VID and P-bit fields. For example, with a downstream mode of 2 and an  upstream rule that translates the VID while carrying forward the P-bit value,  downstream frames that match the specified WAN-side VID will match any  P-bit value and will translate the VID.  0 The operation performed in the downstream direction is the inverse of  that performed in the upstream direction. Which trea
```
