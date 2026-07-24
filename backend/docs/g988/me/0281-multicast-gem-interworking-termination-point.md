# Managed Entity

## Identity
- ME ID: 281
- ME Name: Multicast GEM interworking termination point
- Source Section: 9.2.5
- Source Page: 108

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Set
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
- Name: PPTP counter
- Size: 1 byte
- Format: needs_review
- Access: R
- Category: optional

### Attribute 6
- Name: Operational state
- Size: 1 byte
- Format: needs_review
- Access: R
- Category: optional

### Attribute 7
- Name: GAL profile pointer
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 8
- Name: Not used 2
- Size: 1 byte
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

## Extraction Status
- Status: auto_extracted
- Attributes found: 8
- Review needed: false

## Raw Source

```
9.2.5 Multicast GEM interworking termination point  An instance of this ME represents a point in a G -PON ONU where a multicast service interworks  with the GEM layer. At this point, a multicast bit stream is reconstructed from GEM packets.  Instances of this ME are created and deleted by the OLT.  Multicast interworking GEM modes of operation  The default multicast operation of the PON is where all the multicast content streams are carried in  one PON layer connection (GEM port). This connection is then specified in the first entry of the IPv4  or IPv6 multicast address table, as the case may be. This single entry also specifies an all -inclusive  IP multicast destination address (DA) range (e.g., 224.0.0.0 to 239.255.255.255 in the case of IPv4).  The ONU then filters the traffic based on either Ethernet MAC addresses or IP addresses. 

---- page break ---- The associated GEM port network CTP ME specifies the GEM port -ID that supports all multicast  connections.  In the default multicast operation, all multicast content streams are placed in one PON layer  connection (GEM port). The OLT sets up a completely conventional model, a pointer from the  multicast GEM IW termination to a GEM port network CTP. The OLT configures the GEM port-ID  of the GEM port network CTP into the appropriate multicast address table attribute(s), along with the  other table fields that specify the range of IP multicast DAs. The ONU accepts the entire multicast  stream through the designated GEM port, then filters the traffic based on either the Ethernet MAC  address or IP DA.  An optional multicast configuration supports separate multicast streams carried over separate PON  layer connections, i.e., on separate GEM ports. This permits the ONU to filter multicast streams at  the GEM level, which is efficient in hardware, while ignoring other multicast streams that may be of  interest to other ONUs on the PON.  After configuring the explicit model for the first multicast GEM port, the OLT supports multiple  multicast GEM ports by then configuring additional entries into the multicast address table(s), entries  with different GEM port-IDs. The OMCI model is defined such that these ports are implicitly grouped  together and served by the single explicit GEM port network CTP. No additional GEM network CTPs  need be created or linked for the additional GEM ports.  Several multicast GEM IW TPs can exist, each linked to separate bridge ports or mappers to serve  different communities of interest in a complex ONU.  Discovery of multicast support  The OLT uses the multicast GEM IW TP entity as the means to discover the ONU 's multicast  capability. This entity is mandatory if multicast is supported by the ONU. If the OLT attempts to  create this entity on an ONU that does not support multicast, the create command fails. The create or  set command also fails if the OLT attempts to exploit optional features that the ONU does not support,  e.g., in attempting to write a multicast address table with more than a single entry or to create multiple  multicast GEM IW TPs.  This ME is defined by a similarity to the unicast GEM IW TP, and a number of its attributes are not  meaningful in a multicast context. These attributes are set to 0 and not used, as indicated in the  following.  Relationships  An instance of this ME exists for each occurrence of transformation of GEM packets into a  multicast data stream.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. The value  0xFFFF is reserved. (R, set-by-create) (mandatory) (2 bytes)  GEM port network CTP connectivity pointer : This attribute points to an instance of the  GEM port network CTP that is associated with this multicast GEM IW TP.  (R, W, set-by-create) (mandatory) (2 bytes)  Interworking option: This attribute identifies the type of non -GEM function that is being  interworked. The option can be as follows.  0 This value is a "no-op" or "don't care ". It should be used when the  multicast GEM IW TP is associated with several functions of different  types. It can optionally be used in all cases, since the necessary  information is available elsewhere. The previous code points are  retained for backward compatibility: 

---- page break ---- 1 MAC bridged LAN  3 Reserved  5 IEEE 802.1p mapper  (R, W, set-by-create) (mandatory) (1 byte)  Service profile pointer: This attribute is set to 0 and not used. For backward compatibility,  it may also be set to point to a MAC bridge service profile or IEEE 802.1p  mapper service profile. (R, W, set-by-create) (mandatory) (2 bytes)  Not used 1: This attribute is set to 0 and not used. (R, W, set-by-create) (mandatory)  (2 bytes)  PPTP counter: This attribute represents the number of instances of PPTP MEs associated  with this instance of the multicast GEM IW TP. This attribute conveys no  information that is not available elsewhere; it may be set to 0xFF and not used.  (R) (optional) (1 byte)  Operational state : This attribute indicates whether the ME is capable of performing its  function. Valid values are enabled (0) and disabled (1). (R) (optional) (1 byte)  GAL profile pointer: This attribute is set to 0 and not used. For backward compatibility, it  may also be set to point to a GAL Ethernet profile. (R,  W, set-by-create)  (mandatory) (2 bytes)  Not used 2: This attribute is set to 0 and not used. (R,  W, set-by-create) (mandatory)  (1 byte)  IPv4 multicast address table : This attribute maps IP multicast addresses to PON layer  addresses. Each entry contains the following.  GEM port-ID  2 bytes  Secondary key  2 bytes  IP multicast DA range start  4 bytes  IP multicast DA range stop  4 bytes  The first four bytes of each entry are treated as a key into the list. The  secondary key allows the table to contain more than a single range for a given  GEM port.  A set action to a particular value overwrites any existing entry with the same  first four bytes. If the last eight bytes of a set c
```
