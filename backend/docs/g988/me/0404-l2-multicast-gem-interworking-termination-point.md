# Managed Entity

## Identity
- ME ID: 404
- ME Name: L2 multicast GEM interworking termination point
- Source Section: 9.2.18
- Source Page: 126

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Create, Delete, Get, Set
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: GEM port network CTP connectivity pointer
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 2
- Name: Interworking option
- Size: 1 byte
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 3
- Name: Service profile pointer
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 4
- Name: Not used 1
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 5
- Name: Operational state
- Size: 1 byte
- Format: needs_review
- Access: R
- Category: optional

### Attribute 6
- Name: GAL profile pointer
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 7
- Name: Not used 2
- Size: 1 byte
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 8
- Name: Multicast MAC address filtering capability
- Size: 5 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 9
- Name: Multicast MAC address registration mode
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 10
- Name: Aging timer
- Size: 3 bytes
- Format: needs_review
- Access: R, W
- Category: optional

## Extraction Status
- Status: auto_extracted
- Attributes found: 10
- Review needed: false

## Raw Source

```
9.2.18 L2 multicast GEM interworking termination point  An instance of this ME represents a point in an EPON ONU where a multicast service interworks  with the GEM layer. At this point, a multicast bit stream is forwarded.  Instances of this ME are created and deleted by the OLT.  Multicast interworking GEM modes of operation  The default multicast operation of the EPON is one in which all the multicast content streams are  carried in one PON layer connection (GEM port). This connection is then specified by the multicast  MAC address filtering table . According to this table, t he ONU filters the traffic based on Ethernet  MAC addresses. The associated GEM port network CTP ME specifies the GEM port-ID that supports  all multicast connections.  In the default multicast operation, all multicast content streams are placed in one PON layer  connection (GEM port). The OLT sets up a completely conventional model, a pointer from the L2  multicast GEM IW TP to a GEM port network CTP. The OLT configures the GEM port -ID of the  GEM port network CTP into the multicast MAC address filtering table attribute, along with the other  table fields that specify multicast destination MAC addresses for filtering. The ONU accepts the entire 

---- page break ---- multicast stream through the designated GEM port and then filters the traffic based on Ethernet MAC  address.  An optional multicast configuration supports separate multicast streams carried over separate PON  layer connections, i.e., on separate GEM ports. This permits the ONU to filter multicast streams at  the GEM level, which is hardware efficient, while ignoring  other multicast streams that may be of  interest to other ONUs on the PON.  After configuring the explicit model for the first multicast GEM port, the OLT supports multiple  multicast GEM ports by configuring additional entries into the multicast MAC address filtering table,  entries with different GEM port-IDs. The OMCI model is defined such that these ports are implicitly  grouped together and served by the single explicit GEM port network CTP. No additional GEM  network CTPs need be created or linked for the additional GEM ports.  Several L2 multicast GEM IW TPs can exist, each linked to separate bridge ports or mappers, to serve  different communities of interest in a complex ONU.  Discovery of multicast support  The OLT uses the L2 multicast GEM IW TP entity as the means to discover the ONU's multicast  capability. This entity is mandatory if multicast is supported by an EPON ONU. If the OLT attempts  to create this entity on an ONU that does not support multicast, the create command fails. The create  command also fails if the OLT attempts to exploit optional features that the ONU does not support,  e.g., in attempting to create multiple L2 multicast GEM IW TPs.  This ME is defined by similarity to the unicast GEM IW TP, and a number of its attributes are not  meaningful in a multicast context. These attributes are set to 0 and not used, as indicated in the  following.  Relationships  An instance of this ME exists for each layer 2 multicast community of interest.  Attributes  Managed entity ID : This attribute uniquely identifies each instance of this ME. The value  0xFFFF is reserved. (R, set-by-create) (mandatory) (2 bytes)  GEM port network CTP connectivity pointer : This attribute points to an instance of the  GEM port network CTP that is associated with this L2 multicast GEM IW TP.  (R, W, set-by-create) (mandatory) (2 bytes)  Interworking option: This attribute identifies the type of non -GEM function that is being  interworked. The option can be:  0 This value is a "no-op" or "don't care". It should be used when the L2  multicast GEM IW TP is associated with several functions of  different types. It can optionally be used in all cases, since the  necessary information is available elsewhere. The previous code  points are retained for backward compatibility:  1 MAC bridged LAN  3 Reserved  5 IEEE 802.1p mapper  (R, W, set-by-create) (mandatory) (1 byte)  Service profile pointer: This attribute is set to 0 and not used. For backward compatibility,  it may also be set to point to a MAC bridge service profile or IEEE 802.1p  mapper service profile. (R, W, set-by-create) (mandatory) (2 bytes)  Not used 1: This attribute is set to 0 and not used. (R, W, set-by-create) (mandatory) (2 bytes) 

---- page break ---- PPTP counter: This attribute represents the number of instances of PPTP MEs associated  with this instance of the L2 multicast GEM IW TP. This attribute conveys no  information that is not available elsewhere; it may be set to 0xFF and not used.  (R) (optional) (1 byte)  Operational state : This attribute indicates whether the ME is capable of performing its  function. Valid values are enabled (0) and disabled (1). (R) (optional) (1 byte)  GAL profile pointer: This attribute is set to 0 and not used. For backward compatibility, it  may also be set to point to a GAL Ethernet profile. (R,  W, set-by-create)  (mandatory) (2 bytes)  Not used 2: This attribute is set to 0 and not used. (R, W, set-by-create) (mandatory) (1 byte)  Multicast MAC address filtering capability : With this attribute, the ONU  reports to the  OLT its supported multicast MAC address registration methods and the  maximum number of filtering table entries at the ONU.  Multicast MAC address registration method 1 byte  Maximum number of static registration entries  2 bytes  Maximum number of dynamic registration entries 2 bytes  The bits of the multicast MAC address registration method  field are assigned  as follows:  Bit Name Setting  1 (LSB) Static MAC address registration 0: not supported   1: supported  2 Dynamic MAC address registration 0: not supported   1: supported  3..8 Reserved 0  (R) (mandatory) (5 bytes)  Multicast MAC address registration mode : This attribute allows the OLT to specify  the  multicast MAC address registration method.  0 Disable multicast MAC address filtering  1 Static MAC a
```
