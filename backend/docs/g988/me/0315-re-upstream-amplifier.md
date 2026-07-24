# Managed Entity

## Identity
- ME ID: 315
- ME Name: RE upstream amplifier
- Source Section: 9.14.3
- Source Page: 475

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
- Name: Operational mode
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 4
- Name: ARC
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 5
- Name: ARC interval
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 6
- Name: RE downstream amplifier pointer
- Size: 2 bytes
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 7
- Name: Total optical receive signal level table
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 8
- Name: Upper receive optical threshold
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 9
- Name: Transmit optical signal level
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 10
- Name: Lower transmit optical threshold
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 11
- Name: Upper transmit optical threshold
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: optional

## Extraction Status
- Status: auto_extracted
- Attributes found: 11
- Review needed: false

## Raw Source

```
9.14.3 RE upstream amplifier  This ME organizes data associated with each  upstream RE optical amplifier (OA) supported by the  RE. The management ONU automatically creates one instance of this ME for each upstream OA as  follows.  • When the RE has mid-span PON RE upstream OA ports built into its factory configuration.  • When a cardholder is provisioned to expect a circuit pack of the mid-span PON RE upstream  OA type.  • When a cardholder provisioned for plug-and-play is equipped with a circuit pack of the mid- span PON RE upstream OA type. Note that the installation of a plug -and-play card may  indicate the presence of a mid-span PON RE upstream OA via equipment ID as well as its  type attribute, and indeed may cause the management ONU to instantiate a port -mapping  package to specify the ports precisely.  The management ONU automatically deletes instances of this ME when a cardholder is neither  provisioned to expect a mid-span PON RE upstream OA circuit pack, nor is it equipped with a mid- span PON RE upstream OA circuit pack. 

---- page break ---- Relationships  An instance of this ME is associated with an upstream OA, and with an instance of a circuit  pack. If the RE includes OEO regeneration in either direction, the RE upstream amplifier is  also associated with a PPTP RE UNI. Refer to clause 9.14.2 for further discussion.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. Its value  indicates the physical position of the upstream OA. The first byte is the slot ID  (defined in clause 9.1.5). The second byte is the port ID. (R) (mandatory)  (2 bytes)  NOTE 1 – This ME ID may be identical to that of a PPTP RE UNI if it shares the  same physical slot and port.  Administrative state: This attribute locks (1) and unlocks (0) the functions performed by this  ME. Administrative state is further described in clause A.1.6. (R,  W)  (mandatory) (1 byte)  NOTE 2 – Administrative lock of an RE upstream amplifier results in LOS from any  downstream ONUs.  Operational state : This attribute indicates whether the ME is capable of performing its  function. Valid values are enabled (0) and disabled (1). (R) (optional) (1 byte)  Operational mode: This attribute indicates the operational mode as follows.  0 Constant gain  1 Constant output power  2 Autonomous  (R, W) (mandatory) (1 byte)  ARC: See clause A.1.4.3. (R, W) (optional) (1 byte)  ARC interval: See clause A.1.4.3. (R, W) (optional) (1 byte)  RE downstream amplifier pointer: This attribute points to a n RE downstream amplifier  instance. The default value is 0xFFFF, a null pointer.  (R, W) (mandatory)  (2 bytes)  Total optical receive signal level table: This table attribute reports a series of measurements  of time -averaged input upstream optical signal power. The measurement  circuit should have a temporal response similar to a simple 1 pole low pass  filter, with an effective time constant on the order of a GTC frame time. Each  table entry has a 2 byte frame counter field (most significant end), and a 2 byte  power measurement field. The frame counter field contains the least significant  16 bits of the superframe cou nter received closest to the time of the  measurement. The power measurement field is a 2s  complement integer  referred to 1  mW (i.e., dBm), with 0.002  dB granularity. (Coding –32768 to  +32767, where 0x00 = 0  dBm, 0x03e8 = +2  dBm, etc.)  The RE equipment  should add entries to this table as frequently as is reasonable. The RE should  clear the table once it is read by the OLT. (R) (optional) (4 * N bytes, where N  is the number of measurements present.)  

---- page break ---- Per burst receive signal level table: This table attribute reports the most recent measurement  of received burst upstream optical signal power. Each table entry has a 2 byte  ONU-ID field (most significant end), and a 2 byte power measurement field.  The power measurement field is a 2s  complement integer referred to 1  mW  (i.e., dBm), with 0.002  dB granularity. (Coding –32768 to +32767, where  0x00 = 0 dBm, 0x03e8 = +2 dBm, etc.) (R) (optional) (4 * N bytes, where N is  the number of distinct ONUs connected to the S'/R' interface.)  Lower receive optical threshold: This attribute specifies the optical level that the RE uses to  declare the low received optical power alarm. Valid values are –127 dBm  (coded as 254) to 0 dBm (coded as 0) in 0.5 dB increments. The default value  0xFF selects the RE's internal policy. (R, W) (optional) (1 byte)  Upper receive optical threshold: This attribute specifies the optical level that the RE uses to  declare the high  received optical power alarm. Valid values are –127 dBm  (coded as 254) to 0 dBm (coded as 0) in 0.5 dB increments. The default value  0xFF selects the RE's internal policy. (R, W) (optional) (1 byte)  Transmit optical signal level : This attribute reports the current measurement of the mean  optical launch power of the upstream OA. Its value is a 2s complement integer  referred to 1  mW (i.e., dBm), with 0.002  dB granularity. (R) (optional)  (2 bytes)  Lower transmit optical threshold: This attribute specifies the minimum mean optical launch  power that the RE uses to declare the low transmit optical power alarm. Its  value is a 2s  complement integer referred to 1  mW (i.e., dBm), with 0.5  dB  granularity. The default value 0x7F selects the RE 's internal policy. (R,  W)  (optional) (1 byte)  Upper transmit optical threshold: This attribute specifies the maximum mean optical launch  power that the RE uses to declare the high  transmit optical power alarm. Its  value is a 2s complement integer referred to 1  mW (i.e., dBm), with 0.5  dB  granularity. The default value 0x7F selects the RE 's internal policy. (R,  W)  (optional) (1 byte)  Actions  Get, get next, set  Test Test the upstream amplifier. The test action can be used to perform optical line  supervision tests; refer to the test and test result message descriptions in  Annex A.  No
```
