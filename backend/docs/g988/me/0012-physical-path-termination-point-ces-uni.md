# Managed Entity

## Identity
- ME ID: 12
- ME Name: Physical path termination point CES UNI
- Source Section: 9.8.1
- Source Page: 356

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Get, Set
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: Expected type
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 2
- Name: Sensed type
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 3
- Name: Framing
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 4
- Name: Encoding
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 5
- Name: DS1 mode
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 6
- Name: ARC
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 7
- Name: ARC interval
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: optional

## Extraction Status
- Status: auto_extracted
- Attributes found: 7
- Review needed: false

## Raw Source

```
9.8.1 Physical path termination point CES UNI  This ME represents the point at a CES UNI in the ONU where the physical path terminates and  physical level functions are performed.  The ONU automatically creates an instance of this ME per port:  • when the ONU has CES ports built into its factory configuration;  • when a cardholder is provisioned to expect a circuit pack of a CES type;  • when a cardholder provisioned for plug -and-play is equipped with a circuit pack of a CES  type. Note that the installation of a plug -and-play card may indicate the presence of CES  ports via equipment ID as well as its type and indeed may cause the ONU to instantiate a  port-mapping package that specifies CES ports.  The ONU automatically deletes instances of this ME when a cardholder is neither provisioned to  expect a CES circuit pack, nor is it equipped with a CES circuit pack.  Relationships  An instance of this ME is associated with each real or pre -provisioned CES port. It can be  linked from a GEM IW TP, a pseudowire TP or a logical N × 64 kbit/s CTP. 

---- page break ---- Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. This 2 byte  number indicates the physical position of the UNI. The first byte is the slot ID  (defined in clause 9.1.5). The second byte is the port ID, with the range 1..255.  (R) (mandatory) (2 bytes)  Expected type: The following coding is used for this attribute-  0 Autosense  1 to 254 One of the values from Table 9.1.5-1 that is compatible with a  CES circuit pack  Upon ME instantiation, the ONU sets this attribute to 0. (R,  W) (mandatory)  (1 byte)  Sensed type: If the value of expected type is not 0, then the value of sensed type equals the  value of expected type. If expected type  = 0, then the value of sensed type is  one of the compatible values from Table 9.1.5-1. Upon ME instantiation, the  ONU sets this attribute to 0 or to the value that reflects the physically present  equipment. (R) (mandatory if the ONU supports circuit packs with  configurable interface types, e.g., C1.5/2/6.3) (1 byte)  CES loopback configuration: This attribute specifies and reports the loopback configuration  of the physical interface.  0 No loopback  1 Payload loopback  2 Line loopback  3 Operations system -directed ( OS-directed) loopback 1 (loopback  from/to PON side)  4 OS-directed loopback 2 (loopback from/to CES UNI side)  5 OS-directed loopback 3 (loopback of both PON side and CES UNI side)  6 Manual button-directed loopback [read only (RO)]  7 Network-side code inband-directed loopback (RO)  8 SmartJack-directed loopback (RO)  9 Network-side code inband-directed loopback (armed; RO)  10 Remote-line loopback via facility data link (FDL)  11 Remote-line loopback via inband code  12 Remote-payload loopback  Upon ME instantiation, the ONU sets this attribute to 0. (R,  W) (mandatory)  (1 byte)  G.988(12)_F9.8.1-1 ONU PHY framer PON interface PON CES UNI Loopback 2 Loopback 1 Loopback 3   Figure 9.8.1-1 – CES loopback configuration  Administrative state: This attribute locks (1) and unlocks (0) the functions performed by this  ME. Administrative state is further described in clause A.1.6. (R,  W)  (mandatory) (1 byte) 

---- page break ---- Operational state : This attribute indicates whether the ME is capable of performing its  function. Valid values are enabled (0) and disabled (1). (R) (optional) (1 byte)  Framing: This attribute specifies the framing structure.  These code points are for use with DS1 services. Code point 2 may also be  used for an unframed E1 service.  0 Extended superframe  1 Superframe  2 Unframed  3 ITU-T G.704  NOTE – [ITU-T G.704] describes both SF and ESF framing for DS1 signals.  This code point is retained for backward compatibility, but its meaning is  undefined.  4 JT-G.704  The following code points are for use with E1 services.  5 Basic framing: clause 2.3.2 of [ITU-T G.704]  6 Basic framing with CRC-4: clause 2.3.3 of [ITU-T G.704]  7 Basic framing with TS16 multiframe  8 Basic framing with CRC-4 and TS16 multiframe  Upon ME instantiation, the ONU sets this attribute to a value that reflects the  vendor's default. (R, W) (optional) (1 byte)  Encoding: This attribute specifies the line coding scheme. Valid values are as follows.  0 B8ZS  1 AMI  2 HDB3  3 B3ZS  Upon ME instantiation, the ONU sets this attribute to 0. (R, W) (mandatory for  DS1 and DS3 interfaces) (1 byte)  Line length: This attribute specifies the length of the twisted pair cable from a DS1 physical  UNI to the DSX -1 cross-connect point or the length of coaxial cable from a  DS3 physical UNI to the DSX-3 cross-connect point. Valid values are given in  Table 9.8.1-1. Upon ME instantiation for a DS1 interface, the ONU assigns the  value 0 for non-power feed type DS1 and the value 6 for power feed type DS1.  Upon ME instantiation for a DS3 interface, the ONU sets this attribute to 0x0F.  (R, W) (optional) (1 byte)  DS1 mode: This attribute specifies the mode of a DS1. Valid values are as follows.    Value Mode Connect Line length Power Loopback  0 No.1 DS1 CPE Short haul No power feed Smart jack  1 No.2 DS1 CPE Long haul No power feed Smart jack  2 No.3 DS1 NIU CPE Long haul No power feed Intelligent office repeater.  Transparent to FDL.  3 No.4 DS1 NIU CPE Long haul With power feed Intelligent office repeater.  Transparent to FDL.  In the event of conflicting values between this attribute and the (also optional)  line length attribute, the line length attribute is taken to be valid. This permits  the separation of line build-out (LBO) and power settings from smart jack and 

---- page break ---- FDL behaviour. Upon ME instantiation, the ONU sets this attribute to 0. (R, W)  (optional) (1 byte)  ARC: See clause A.1.4.3. (R, W) (optional) (1 byte)  ARC interval: See clause A.1.4.3. (R, W) (optional) (1 byte)  Line type: This attribute specifies the line type used in a DS3 or E3 application or when  the sensed type of the PPTP is configurable. Val
```
