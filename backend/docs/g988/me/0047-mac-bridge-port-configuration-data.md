# Managed Entity

## Identity
- ME ID: 47
- ME Name: MAC bridge port configuration data
- Source Section: 9.3.4
- Source Page: 142

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Create, Delete, Get, Set
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: Bridge ID pointer
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 2
- Name: Port num
- Size: 1 byte
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 3
- Name: TP type
- Size: 1 byte
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 4
- Name: TP pointer
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 5
- Name: Port priority
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: optional

### Attribute 6
- Name: Port path cost
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 7
- Name: Port spanning tree ind
- Size: 1 byte
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 8
- Name: Deprecated 1
- Size: 1 byte
- Format: needs_review
- Access: R, W, set-by-create
- Category: optional

### Attribute 9
- Name: Deprecated 2
- Size: 6 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 10
- Name: Outbound TD pointer
- Size: 2 byte
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 11
- Name: Inbound TD pointer
- Size: 2 byte
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 12
- Name: MAC learning depth
- Size: 1 byte
- Format: needs_review
- Access: R, W, set-by-create
- Category: optional

### Attribute 13
- Name: LASP ID pointer
- Size: 2 bytes
- Format: needs_review
- Access: R,W, set-by-create
- Category: optional

## Extraction Status
- Status: auto_extracted
- Attributes found: 13
- Review needed: false

## Raw Source

```
9.3.4 MAC bridge port configuration data  This ME models a port on a MAC bridge. Instances of this ME are created and deleted by the OLT. 

---- page break ---- Relationships  An instance of this ME is linked to an instance of the MAC bridge service profile. Additional  bridge port control capabilities are provided by implicitly linked instances of some or all of:   • MAC bridge port filter table data;   • MAC bridge port filter pre-assign table;   • VLAN tagging filter data;   • Dot1 rate limiter.  Real-time status of the bridge port is provided by implicitly linked instances of:   • MAC bridge port designation data;   • MAC bridge port bridge table data;   • Multicast subscriber monitor.  Bridge port PM collection is provided by implicitly linked instances of:   • MAC bridge port PM history data;   • Ethernet frame PM history data upstream and downstream;   • Ethernet frame extended PM (preferred).  Attributes  Managed entity  ID: This attribute uniquely identifies each instance of this ME. (R,  set-by-create) (mandatory) (2 bytes)  Bridge ID pointer: This attribute points to an instance of the MAC bridge service profile.  (R, W, set-by-create) (mandatory) (2 bytes)  Port num: This attribute is the bridge port number. It must be unique among all ports  associated with a particular MAC bridge service profile. (R, W, set-by-create)  (mandatory) (1 byte)  TP type: This attribute identifies the type of TP associated with this MAC bridge port.  Valid values are as follows.  1 Physical path termination point Ethernet UNI  2 Interworking virtual circuit connection (VCC) termination point  3 IEEE 802.1p mapper service profile  4 IP host config data or IPv6 host config data  5 GEM interworking termination point  6 Multicast GEM interworking termination point  7 Physical path termination point xDSL UNI part 1  8 Physical path termination point VDSL UNI  9 Ethernet flow termination point  10 Reserved  11 Virtual Ethernet interface point  12 Physical path termination point MoCA UNI  13 Ethernet in the first mile (EFM) bonding group  (R, W, set-by-create) (mandatory) (1 byte)  TP pointer: This attribute points to the TP associated with this MAC bridge port. The TP  type attribute indicates the type of the TP; this attribute contains its instance  identifier (ME ID). (R, W, set-by-create) (mandatory) (2 bytes) 

---- page break ---- NOTE 1 – When the TP type is very high-speed digital subscriber line (VDSL) or  xDSL, the two MSBs may be used to indicate a bearer channel.  Port priority: This attribute denotes the priority of the port for use in (rapid) spanning tree  algorithms. The range is 0..255. (R, W, set-by-create) (optional) (2 bytes)  Port path cost: This attribute specifies the contribution of the port to the path cost toward s  the spanning tree root bridge.  The range is 1..65535. (R,  W, set-by-create)  (mandatory) (2 bytes)  Port spanning tree ind : The Boolean value true enables (R)STP LAN topology change  detection at this port. The value false disables topology change detection.  (R, W, set-by-create) (mandatory) (1 byte)  Deprecated 1: This attribute is not used. If present, it should be ignored by both the ONU and  the OLT, except as necessary to comply with OMCI message definitions.  (R, W, set-by-create) (optional) (1 byte)  Deprecated 2: This attribute is not used. If present, it should be ignored by both the ONU and  the OLT, except as necessary to comply with OMCI message definitions .  (R, W, set-by-create) (1 byte) (optional)  Port MAC address : If the TP associated with this port has a MAC address, this attribute  specifies it. (R) (optional) (6 bytes)  Outbound TD pointer: This attribute points to a traffic descriptor that limits the traffic rate  leaving the MAC bridge. (R, W) (optional) (2 byte)  Inbound TD pointer : This attribute points to a traffic descriptor that limits the traffic rate  entering the MAC bridge. (R, W) (optional) (2 byte)  MAC learning depth: This attribute specifies the maximum number of MAC addresses to be  learned by this MAC bridge port. The default value 0 specifies that there is no  administratively imposed limit. (R, W, set-by-create) (optional) (1 byte)  NOTE 2 – If this attribute is not zero, its value overrides the value set in the MAC  learning depth attribute of the MAC bridge service profile.  LASP ID pointer: This attribute points to an instance of the LASP ME. (R,W,  set-by-create) (optional) (2 bytes)  Actions  Create, delete, get, set  Notifications  Alarm  Alarm  number Alarm Description  0 Port blocking This port has been blocked due to loop detection in  accordance with [IEEE 802.1D] (Note).  1..207 Reserved   208..223 Vendor-specific alarms Not to be standardized  NOTE – To determine the state of a MAC bridge port, the OLT can read the port state attribute of  the MAC bridge port designation data. 

---- page break ---- 
```
