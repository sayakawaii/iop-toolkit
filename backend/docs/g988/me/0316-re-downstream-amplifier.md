# Managed Entity

## Identity
- ME ID: 316
- ME Name: RE downstream amplifier
- Source Section: 9.14.4
- Source Page: 478

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Get, Set, Test
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: Operational state
- Size: 1 byte
- Format: needs_review
- Access: R
- Category: optional

### Attribute 2
- Name: ARC
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 3
- Name: ARC interval
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 4
- Name: Operational mode
- Size: 1 byte
- Format: needs_review
- Access: R,W
- Category: mandatory

### Attribute 5
- Name: Input optical signal level
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 6
- Name: Lower input optical threshold
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 7
- Name: Upper input optical threshold
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 8
- Name: Output optical signal level
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 9
- Name: Lower output optical threshold
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 10
- Name: Upper output optical threshold
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: optional

## Extraction Status
- Status: auto_extracted
- Attributes found: 10
- Review needed: false

## Raw Source

```
9.14.4 RE downstream amplifier  This ME organizes data associated with each  OA for downstream data  supported by the RE. The  management ONU automatically creates one instance of this ME for each downstream OA as follows.  • When the RE has mid-span PON RE downstream OA ports built into its factory  configuration.  • When a cardholder is provisioned to expect a circuit pack of the mid-span PON RE  downstream OA type.  • When a cardholder provisioned for plug -and-play is equipped with a circuit pack of  the  mid-span PON RE downstream OA type. Note that the installation of a plug -and-play card  may indicate the presence of a mid-span PON RE downstream OA via equipment ID as well  as its type attribute, and indeed may cause the management ONU to instantiate a port - mapping package to specify the ports precisely.  The management ONU automatically deletes instances of this ME when a cardholder is neither  provisioned to expect a mid-span PON RE downstream OA circuit pack, nor is it equipped with a  mid-span PON RE downstream OA circuit pack.  Relationships  An instance of this ME is associated with a downstream OA and with an instance of a circuit  pack. If the RE includes OEO regeneration in either direction, the RE downstream amplifier  is also associated with an RE ANI-G. Refer to clause 9.14.1 for further discussion.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. Its value  indicates the physical position of the downstream OA. The first byte is the slot  ID (defined in clause 9.1.5  of [ITU-T G.984.4]). The second byte is the port  ID. (R) (mandatory) (2 bytes)  NOTE 1 – This ME ID may be identical to that of an RE ANI-G if it shares the same  physical slot-port.  

---- page break ---- Administrative state: This attribute locks (1) and unlocks (0) the functions performed by this  ME. Administrative state is further described in clause A.1.6. (R,  W)  (mandatory) (1 byte)  NOTE 2– When an RE supports multiple PONs, or protected access to a single PON,  its primary ANI-G cannot be completely shut down, due to a loss of the management  communications capability. Complete blocking of service and removal of power may  nevertheless be appropriate for secondary RE ANI -Gs. Administrative lock  suppresses alarms and notifications for both primary and secondary RE ANI -Gs.  Administrative lock suppresses alarms and notifications for an RE downstream  amplifier, be it either primary or secondary.  Operational state : This attribute indicates whether the ME is capable of performing its  function. Valid values are enabled (0) and disabled (1). (R) (optional) (1 byte)  ARC: See clause A.1.4.3. (R, W) (optional) (1 byte)  ARC interval: See clause A.1.4.3. (R, W) (optional) (1 byte)  Operational mode: This attribute indicates the operational mode as follows.  0 Constant gain  1 Constant output power  2 Autonomous  (R,W) (mandatory) (1 byte)  Input optical signal level: This attribute reports the current measurement of the input optical  signal power  of the downstream OA. Its value is a 2s  complement integer  referred to 1  mW (i.e., dBm), with 0.002  dB granularity. (Coding –32768 to  +32767, where 0x00 = 0 dBm, 0x03e8 = +2 dBm, etc.) (R) (optional) (2 bytes)  Lower input optical threshold : This attribute specifies the optical level the RE uses to  declare the low received optical power alarm. Valid values are –127 dBm  (coded as 254) to 0 dBm (coded as 0) in 0.5 dB increments. The default value  0xFF selects the RE's internal policy. (R, W) (optional) (1 byte)  Upper input optical threshold : This attribute specifies the optical level the RE uses to  declare the high  received optical power alarm. Valid values are –127 dBm  (coded as 254) to 0 dBm (coded as 0) in 0.5 dB increments. The default value  0xFF selects the RE's internal policy. (R, W) (optional) (1 byte)  Output optical signal level : This attribute reports the current measurement of the mean  optical launch power  of the downstream OA. Its value is a 2s  complement  integer referred to 1  mW (i.e., dBm), with 0.002  dB granularity. (Coding  −32768 to +32767, where 0x00 = 0  dBm, 0x03e8 = +2  dBm, etc.) (R)  (optional) (2 bytes)  Lower output optical threshold: This attribute specifies the minimum mean optical launch  power that the RE uses to declare the low transmit optical power alarm. Its  value is a 2s complement integer referred to 1  mW (i.e., dBm), with 0.5  dB  granularity. The default value 0x7F selects the RE 's internal policy. (R,  W)  (optional) (1 byte)  Upper output optical threshold: This attribute specifies the maximum mean optical launch  power that the RE uses to declare the high  transmit optical power alarm. Its  value is a 2s complement integer referred to 1  mW (i.e., dBm), with 0.5  dB  granularity. The default value 0x7F selects the RE 's internal policy. (R,  W)  (optional) (1 byte) 

---- page break ---- R'S' splitter coupling ratio: This attribute reports the coupling ratio of the splitter at the R'/S'  interface that connects the embedded management ONU and the amplifiers to  the OTL. Valid values are 99:1 (coded as 99  decimal) to 1:99 (coded as 1  decimal), where the first value is the value encoded and is the percentage of  the optical signal connected to the amplifier. The default value 0xFF indicates  that there is no splitter connected to this upstream/downstream amplifier pair.  (R) (optional) (1 byte)  Actions  Get, set  Test Test the RE downstream amplifier . The test action can be used to perform  optical line supervision tests; refer to the test and test result message  descriptions in Annex A.  Notifications  Attribute value change  Number Attribute value change Description  1 N/A   2 Op state Operational state of RE downstream amplifier  3 ARC Alarm-reporting control cancellation  4..12 N/A   13..16 Reserved     Alarm  Alarm  number Alarm Description  0 Low received optical power Received downstream optical power below threshold  1 High received
```
