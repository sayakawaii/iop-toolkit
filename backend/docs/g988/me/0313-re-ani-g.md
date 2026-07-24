# Managed Entity

## Identity
- ME ID: 313
- ME Name: RE ANI-G
- Source Section: 9.14.1
- Source Page: 469

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Get, Set, Test
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
- Name: Optical signal level
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 6
- Name: Lower optical threshold
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 7
- Name: Upper optical threshold
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 8
- Name: Transmit optical level
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 9
- Name: Lower transmit power threshold
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 10
- Name: Upper transmit power threshold
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 11
- Name: Usage mode
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 12
- Name: Target upstream frequency
- Size: 4 bytes
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 13
- Name: Target downstream frequency
- Size: 4 bytes
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 14
- Name: Upstream signal transmission mode
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: optional

## Extraction Status
- Status: auto_extracted
- Attributes found: 14
- Review needed: false

## Raw Source

```
9.14.1 RE ANI-G  This ME organizes data associated with each R'/S' physical interface of a n RE if the RE supports  OEO regeneration in either direction . The management ONU automatically creates one instance of  this ME for each R'/S' physical port (uni- or bidirectional) as follows.  • When the RE has mid-span PON RE ANI interface ports built into its factory configuration.  • When a cardholder is provisioned to expect a circuit pack of the mid-span PON RE ANI type.  • When a cardholder provisioned for plug -and-play is equipped with a circuit pack of  the  mid-span PON RE ANI type. Note that the installation of a plug-and-play card may indicate  the presence of a mid-span PON RE ANI port via equipment ID as well as its type attribute,  and indeed may cause the management ONU to instantiate a port-mapping package to specify  the ports precisely.  The management ONU automatically deletes instances of this ME when a cardholder is neither  provisioned to expect a mid-span PON RE ANI circuit pack, nor is it equipped with a mid-span PON  RE ANI circuit pack.  As illustrated in Figure 8.2.10-4, an RE ANI-G may share the physical port with an RE downstream  amplifier. The ONU declares a shared configuration through the port -mapping package combined  port table, whose structure defines one ME as the master. It is recommended that the RE ANI -G be  the master, with the RE downstream amplifier as a secondary ME.  The administrative state, operational state and ARC attributes of the master ME override similar  attributes in secondary MEs associated with the same port. In the secondary ME, these attributes are  present, but cause no action when written and have undefined values when read. The RE downstream 

---- page break ---- amplifier should use its provisionable downstream alarm thresholds and should declare downstream  alarms as necessary; other isomorphic alarms should be declared by the RE ANI -G. The test action  should be addressed to the master ME.  Relationships  An instance of this ME is associated with each R'/S' physical interface of an RE that includes  OEO regeneration in either direction, and with one or more instances of the PPTP RE UNI. It  may also be associated with an RE downstream amplifier.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. Its value  indicates the physical position of the R'/S' interface. The first byte is the slot  ID (defined in clause 9.1.5). The second byte is the port ID. (R) (mandatory)  (2 bytes)  NOTE 1 – This ME ID may be identical to that of an RE downstream amplifier if it  shares the same physical slot and port.  Administrative state: This attribute locks (1) and unlocks (0) the functions performed by this  ME. Administrative state is further described in clause A.1.6. (R,  W)  (mandatory) (1 byte)  NOTE 2 – When an RE supports multiple PONs, or protected access to a single PON,  its primary ANI-G cannot be completely shut down, due to a loss of the management  communications capability. Complete blocking of service and removal of power may  nevertheless be appropriate for secondary RE ANI -Gs. Administrative lock  suppresses alarms and notifications for an RE ANI -G, be it either primary or  secondary.  Operational state : This attribute indicates whether the ME is capable of performing its  function. Valid values are enabled (0) and disabled (1). (R) (optional) (1 byte)  ARC: See clause A.1.4.3. (R, W) (optional) (1 byte)  ARC interval: See clause A.1.4.3. (R, W) (optional) (1 byte)  Optical signal level : This attribute reports the current measurement of total downstream  optical power. Its value is a 2s complement integer referred to 1  mW (i.e.,  dBm), with 0.002  dB granularity. (Coding –32768 to +32767, where 0x00 =  0 dBm, 0x03e8 = +2 dBm, etc.) (R) (optional) (2 bytes)  Lower optical threshold: This attribute specifies the optical level that the RE uses to declare  the downstream low received optical power alarm. Valid values are   –127 dBm (coded as 254) to 0  dBm (coded as 0) in 0.5  dB increments. The  default value 0xFF selects the RE's internal policy. (R, W) (optional) (1 byte)  Upper optical threshold: This attribute specifies the optical level that the RE uses to declare  the downstream high received optical power alarm. Valid values are   –127 dBm (coded as 254) to 0  dBm (coded as 0) in 0.5 dB increments. The  default value 0xFF selects the RE's internal policy. (R, W) (optional) (1 byte)  Transmit optical level: This attribute reports the current measurement of mean optical launch  power. Its value is a 2s complement integer referred to 1 mW (i.e., dBm), with  0.002 dB granularity. (Coding –32768 to +32767, where 0x00 = 0  dBm,  0x03e8 = +2 dBm, etc.) (R) (optional) (2 bytes)  Lower transmit power threshold: This attribute specifies the minimum mean optical launch  power that the RE uses to declare the low transmit optical power alarm. Its  value is a 2s  complement integer referred to 1  mW (i.e., dBm), with 0.5  dB 

---- page break ---- granularity. The default value 0x7F selects the RE 's internal policy. (R,  W)  (optional) (1 byte)  Upper transmit power threshold: This attribute specifies the maximum mean optical launch  power that the RE uses to declare the high transmit optical power alarm. Its  value is a 2s  complement integer referred to 1  mW (i.e., dBm), with 0.5  dB  granularity. The default value 0x7F selects the RE 's internal policy. (R,  W)  (optional) (1 byte)  Usage mode: In a mid-span PON RE, an R'/S' interface may be used as the PON interface  for the embedded management ONU  or the uplink interface for an S'/R'  interface. This attribute specifies the usage of the R'/S' interface.  (R, W)  (mandatory) (1 byte)  0 Disable  1 This R'/S' interface is used as the uplink for the embedded management  ONU  2 This R'/S' interface is used as the uplink for one or more PPTP RE  UNI(s)  3 This R'/S' interface is used as the uplink for both the embedded  management ONU a
```
