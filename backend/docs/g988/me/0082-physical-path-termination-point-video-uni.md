# Managed Entity

## Identity
- ME ID: 82
- ME Name: Physical path termination point video UNI
- Source Section: 9.13.1
- Source Page: 451

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
- Name: ARC
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 4
- Name: ARC interval
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 5
- Name: Power control
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: optional

## Extraction Status
- Status: auto_extracted
- Attributes found: 5
- Review needed: false

## Raw Source

```
9.13.1 Physical path termination point video UNI  This ME represents an RF video UNI in the ONU, where physical paths terminate and physical path  level functions are performed.  The ONU automatically creates an instance of this ME per port:  • when the ONU has RF video UNI ports built into its factory configuration;  • when a cardholder is provisioned to expect a circuit pack of the video UNI type;  • when a cardholder provisioned for plug-and-play is equipped with a circuit pack of the video  UNI type. Note that the installation of a plug -and-play card may indicate the presence of  video ports via equipment ID as well as its type, and indeed may cause the ONU to instantiate  a port-mapping package that specifies video ports.  The ONU automatically deletes instances of this ME when a cardholder is neither provisioned to  expect a video circuit pack, nor is it equipped with a video circuit pack.  Relationships  One or more instances of this ME are associated with an instance of a real or virtual circuit  pack classified as video type. 

---- page break ---- Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. This 2 byte  number indicates the physical position of the UNI. The first byte is the slot ID  (defined in clause 9.1.5). The second byte is the port ID, with the range 1..255.  (R) (mandatory) (2 bytes)  Administrative state: This attribute locks (1) and unlocks (0) the functions performed by this  ME. Administrative state is further described in clause A.1.6. (R,  W)  (mandatory) (1 byte)  Operational state : This attribute indicates whether the ME is capable of performing its  function. Valid values are enabled (0) and disabled (1). (R) (optional) (1 byte)  ARC: See clause A.1.4.3. (R, W) (optional) (1 byte)  ARC interval: See clause A.1.4.3. (R, W) (optional) (1 byte)  Power control : This attribute controls whether power is provided from the ONU to an  external equipment over the video PPTP. Value 1 enables power over coaxial  cable. The default value 0 disables power feed. (R, W) (optional) (1 byte)  Actions  Get, set  Notifications  Attribute value change  Number Attribute value change Description  1 N/A   2 Op state Operational state of video UNI  3 ARC ARC timer expiration  4..5 N/A   6..16 Reserved     Alarm  Alarm  number  Alarm Description  0 Video-LOS No signal at the video UNI  1 Video-OOR-low RF output below rated value  2 Video-OOR-high RF output above rated value  3..207 Reserved Reserved for vendor-specific alarms  208..223 Vendor-specific alarms Not to be standardized  
```
