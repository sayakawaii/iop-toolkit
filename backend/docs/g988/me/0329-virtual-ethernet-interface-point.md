# Managed Entity

## Identity
- ME ID: 329
- ME Name: Virtual Ethernet interface point
- Source Section: 9.5.5
- Source Page: 239

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Get, Set
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: Administrative state
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 2
- Name: Operational state
- Size: 1 byte
- Format: needs_review
- Access: R
- Category: optional

### Attribute 3
- Name: Interdomain name
- Size: 25 bytes
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 4
- Name: TCP/UDP pointer
- Size: 2 bytes
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 5
- Name: IANA assigned port
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: mandatory

## Extraction Status
- Status: auto_extracted
- Attributes found: 5
- Review needed: false

## Raw Source

```
9.5.5 Virtual Ethernet interface point  This ME represents the data plane hand-off point in an ONU to a separate (non-OMCI) management  domain. The VEIP is managed by the OMCI, and is potentially known to the non-OMCI management  domain. One or more Ethernet traffic flows are present at this boundary.  Instances of this ME are automatically created and deleted by the ONU. This is necessary because  the required downstream priority queues are subject to physical implementation constraints. The OLT  may use one or more of the VEIPs created by the ONU.  It is expected that the ONU w ill create one VEIP for each non-OMCI management domain. At the  vendor's discretion, a VEIP may be created for each traffic class.  Relationships  An instance of this ME is associated with an instance of a virtual Ethernet interface between  OMCI and non-OMCI management domains.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. When used  independently of a cardholder and circuit pack, the ONU should assign IDs in  the sequence 1, 2, .... When used in conjunction with a cardholder and circuit  pack, this 2 byte number indicates the physical position of the VEIP. The first  byte is the slot ID (defined in clause 9.1.5). The second byte is the port ID,  with the range 1..255. The values 0 and 0xFFFF are reserved. (R) (mandatory)  (2 bytes)  Administrative state: This attribute locks (1) and unlocks (0) the functions performed by this  ME. Administrative state is further described in clause A.1.6. (R,  W)  (mandatory) (1 byte)  Operational state : This attribute indicates whether the ME is capable of performing its  function. Valid values are enabled (0) and disabled (1). (R) (optional) (1 byte)  Interdomain name : This attribute is a character string that provides an optional way to  identify the VEIP to a non-OMCI management domain. The interface may also  be identified by its ME ID,   [b-IANA] assigned port and possibly other ways. If the vendor offers no  information in this attribute, it should be set to a sequence of null bytes. (R, W)  (optional) (25 bytes)  TCP/UDP pointer: This attribute points to an instance of the TCP/UDP config data ME,  which provides for OMCI management of the non -OMCI management  domain's IP connectivity. If no OMCI management of the non-OMCI domain's 

---- page break ---- IP connectivity is required, this attribute may be omitted or set to its default, a  null pointer. (R, W) (optional) (2 bytes)  IANA assigned port : This attribute contains the TCP or UDP port value as assigned by   [b-IANA] for the management protocol associated with this virtual Ethernet  interface. This attribute is to be regarded as a hint, not as a requirement that  management communications use this port; the actual port and protocol are  specified in the associated TCP/UD P config data ME. If no port has been  assigned or if the management protocol is free to be chosen at run -time, this  attribute should be set to 0xFFFF. (R) (mandatory) (2 bytes)  NOTE – This attribute does not apply to USP management protocol. For USP  MTP discovery, please refer to clause 9.12.19.  Actions  Get, set  Notifications  Attribute value change  Number Attribute value change Description  0..1 N/A   2 Op state Operational state  3 N/A   4..16 Reserved     Alarm  Alarm  number  Alarm Description  0 Connecting function fail Indicates a failure of the connecting function. May be used to  signal faults from the non-OMCI management domain into  the OMCI.  1..207 Reserved   208..223 Vendor-specific alarms Not to be standardized  
```
