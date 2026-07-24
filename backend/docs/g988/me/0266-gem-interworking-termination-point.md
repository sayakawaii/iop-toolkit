# Managed Entity

## Identity
- ME ID: 266
- ME Name: GEM interworking termination point
- Source Section: 9.2.4
- Source Page: 106

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
- Name: Interworking termination point pointer
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
- Name: GAL loopback configuration
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: mandatory

## Extraction Status
- Status: auto_extracted
- Attributes found: 7
- Review needed: false

## Raw Source

```
9.2.4 GEM interworking termination point  An instance of this ME represents a point in the ONU where the IW of a bearer service (usually  Ethernet) to the GEM layer takes place. At this point, GEM packets are generated from the bearer bit  stream (e.g., Ethernet) or the bearer bit stream is reconstructed from GEM packets.  Instances of this ME are created and deleted by the OLT. 

---- page break ---- Relationships  One instance of this ME exists for each transformation of a data stream into GEM frames and  vice versa.  Attributes  Managed entity  ID: This attribute uniquely identifies each instance of this ME. (R,  set-by-create) (mandatory) (2 bytes)  GEM port network CTP connectivity pointer : This attribute points to an instance of the  GEM port network CTP. (R, W, set-by-create) (mandatory) (2 bytes)  Interworking option: This attribute identifies the type of non -GEM function that is being  interworked. The options are as follows.  0 Circuit-emulated TDM  1 MAC bridged LAN  2 Reserved  3 Reserved  4 Video return path  5 IEEE 802.1p mapper  6 Downstream broadcast  7 MPLS PW TDM service  (R, W, set-by-create) (mandatory) (1 byte)  Service profile pointer: This attribute points to an instance of a service profile:  CES service profile if IW option = 0  MAC bridge service profile if IW option = 1  Video return path service profile if IW option = 4  IEEE 802.1p mapper service profile if IW option = 5  Null pointer if IW option = 6  CES service profile if IW option = 7  (R, W, set-by-create) (mandatory) (2 bytes)  NOTE – The video return path (VRP) service profile is defined in [ITU-T G.984.4].  Interworking termination point pointer: This attribute is used for the CES and IEEE 802.1p  mapper service without a MAC bridge. Depending on the service provided, it  points to the associated instance of the following MEs:  PPTP CES UNI  Logical N × 64 kbit/s sub-port CTP  PPTP Ethernet UNI  In all other GEM services, the relationship between the related service TP and  this GEM IW TP is derived from other ME relations; this attribute is set to a  null pointer and not used. (R, W, set-by-create) (mandatory) (2 bytes)  PPTP counter: This value reports the number of PPTP ME instances associated with this  GEM IW TP. (R) (optional) (1 byte)  Operational state : This attribute indicates whether the ME is capable of performing its  function. Valid values are enabled (0) and disabled (1). (R) (optional) (1 byte) 

---- page break ---- GAL profile pointer: This attribute points to an instance of the GAL profile. The relationship  between the IW option and the related GAL profile is as follows.  Interworking option GAL profile type  0 Null pointer  1 GAL Ethernet profile  3 GAL Ethernet profile for data service  4 GAL Ethernet profile for video return path  5 GAL Ethernet profile for IEEE 802.1p mapper  6 Null pointer  7 Null pointer  (R, W, set-by-create) (mandatory) (2 bytes)  GAL loopback configuration: This attribute sets the loopback configuration when using  GEM mode:  0 No loopback  1 Loopback of downstream traffic after GAL  The default value of this attribute is 0. When the IW option is 6 (downstream  broadcast), this attribute is not used. (R, W) (mandatory) (1 byte)  Actions  Create, delete, get, set  Notifications  Attribute value change  Number Attribute value change Description  1..5 N/A   6 Op state Operational state change  7..8 N/A   9..16 Reserved     Alarm  Alarm  number Alarm Description  0 Deprecated   1..207 Reserved   208..223 Vendor-specific alarms Not to be standardized  
```
