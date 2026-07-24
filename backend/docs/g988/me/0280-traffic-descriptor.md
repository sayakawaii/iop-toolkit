# Managed Entity

## Identity
- ME ID: 280
- ME Name: Traffic descriptor
- Source Section: 9.2.12
- Source Page: 118

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Create, Delete, Get, Set
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: CIR
- Size: 4 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: optional

### Attribute 2
- Name: PIR
- Size: 4 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: optional

### Attribute 3
- Name: CBS
- Size: 4 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: optional

### Attribute 4
- Name: PBS
- Size: 4 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: optional

### Attribute 5
- Name: Colour mode
- Size: 1 byte
- Format: needs_review
- Access: R, W, set-by-create
- Category: optional

### Attribute 6
- Name: Ingress colour marking
- Size: 1 byte
- Format: needs_review
- Access: R, W, set-by-create
- Category: optional

### Attribute 7
- Name: Egress colour marking
- Size: 1 byte
- Format: needs_review
- Access: R, W, set-by-create
- Category: optional

### Attribute 8
- Name: Meter type
- Size: 1 byte
- Format: needs_review
- Access: R, set-by-create
- Category: optional

## Extraction Status
- Status: auto_extracted
- Attributes found: 8
- Review needed: false

## Raw Source

```
9.2.12 Traffic descriptor  The traffic descriptor is a profile that allows for traffic management. A priority controlled ONU can  point from a MAC bridge port configuration data ME to a traffic descriptor in order to implement  traffic management (marking, policing). A rate controlled ONU can point to a traffic descriptor from  either a MAC bridge port configuration data ME or a GEM port network CTP to implement traffic  management (marking, shaping).  Packets are determined to be green, yellow or red as a function of the ingress packet rate and the  settings in this ME. The colour indicates drop precedence (eligibility), subsequently used by the  priority queue ME to drop packets conditionally during cong estion conditions. Packet colour is also  used by the optional mode 1 DBA status reporting function described in [ITU -T G.984.3]. Red  packets are dropped immediately. Yellow packets are marked as drop eligible, and green packets are  marked as not drop eligible, according to the egress colour marking attribute. 

---- page break ---- The algorithm used to determine the colour marking is specified by the meter type attribute. If  [b-IETF RFC 4115] is used, then:  CIR4115 = CIR  EIR4115 = PIR – CIR (EIR: excess information rate)  CBS4115 = CBS  EBS4115 = PBS – CBS.  Relationships  This ME is associated with a GEM port network CTP or a MAC bridge port configuration  data ME.  Attributes  Managed entity  ID: This attribute uniquely identifies each instance of this ME. (R,  set-by-create) (mandatory) (2 bytes)  CIR: This attribute specifies the committed information rate, in byte s per second.  The default is 0. (R, W, set-by-create) (optional) (4 bytes)  PIR: This attribute specifies the peak information rate, in byte s per second. The  default value 0 accepts the ONU 's factory policy. (R,  W, set-by-create)  (optional) (4 bytes)  CBS: This attribute specifies the committed burst size, in bytes. The default is 0.  (R, W, set-by-create) (optional) (4 bytes)  PBS: This attribute specifies the peak burst size, in bytes. The default value 0 accepts  the ONU's factory policy. (R, W, set-by-create) (optional) (4 bytes)  Colour mode: This attribute specifies whether the colour marking algorithm considers pre - existing marking on ingress packets (colour-aware) or ignores it (colour-blind).  In colour-aware mode, packets can only be demoted (from green to yellow or  red, or from yellow to red). The default value is 0.  0 Colour-blind  1 Colour-aware  (R, W, set-by-create) (optional) (1 byte)  Ingress colour marking: This attribute is meaningful in colour-aware mode. It identifies how  pre-existing drop precedence is marked on ingress packets. For DEI and PCP  marking, a drop eligible indicator is equivalent to yellow; otherwise, the colour  is green. For DSCP AF markin g, the lowest drop precedence is equivalent to  green; otherwise, the colour is yellow. The default value is 0.  0 No marking (ignore ingress marking)  2 DEI [IEEE 802.1ad]  3 PCP 8P0D [IEEE 802.1ad]  4 PCP 7P1D [IEEE 802.1ad]  5 PCP 6P2D [IEEE 802.1ad]  6 PCP 5P3D [IEEE 802.1ad]  7 DSCP AF class [IETF RFC 2597]  (R, W, set-by-create) (optional) (1 byte)  Egress colour marking: This attribute specifies how drop precedence is to be marked by the  ONU on egress packets. If set to internal marking only, the externally visible  packet contents are not modified, but the packet is identified in a vendor - specific local way that indicate s its colour to the priority queue ME. It is 

---- page break ---- possible for the egress marking to differ from the ingress marking; for example,  ingress PCP marking could be translated to DEI egress marking. The default  value is 0.  0 No marking  1 Internal marking only  2 DEI [IEEE 802.1ad]  3 PCP 8P0D [IEEE 802.1ad]  4 PCP 7P1D [IEEE 802.1ad]  5 PCP 6P2D [IEEE 802.1ad]  6 PCP 5P3D [IEEE 802.1ad]  7 DSCP AF class [IETF RFC 2597]  (R, W, set-by-create) (optional) (1 byte)  Meter type: This attribute specifies the algorithm used to determine the colour of the  packet. The default value is 0.  0 Not specified  1 [b-IETF RFC 4115]  2 [b-IETF RFC 2698]  (R, set-by-create) (optional) (1 byte)  Actions  Create, delete, get, set  Notifications  None.  
```
