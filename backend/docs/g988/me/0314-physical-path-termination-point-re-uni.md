# Managed Entity

## Identity
- ME ID: 314
- ME Name: Physical path termination point RE UNI
- Source Section: 9.14.2
- Source Page: 472

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
- Name: RE ANI -G pointer
- Size: 2 bytes
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 6
- Name: Total optical receive signal level table
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 7
- Name: Upper receive optical threshold
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
- Name: Additional preamble
- Size: 1 byte
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 12
- Name: Additional guard time
- Size: 1 byte
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 13
- Name: Connected ONUs table
- Size: 1 byte
- Format: needs_review
- Access: W
- Category: optional

## Extraction Status
- Status: auto_extracted
- Attributes found: 13
- Review needed: false

## Raw Source

```
9.14.2 Physical path termination point RE UNI  This ME represents an S'/R' interface in a mid-span PON RE that supports OEO regeneration in at  least one direction, w here physical paths terminate and physical path level functions are performed  (transmit or receive).  Such an RE automatically creates an instance of this ME for each S'/R' interface port as follows.  • When the RE has mid-span PON RE UNI interface ports built into its factory configuration.  • When a cardholder is provisioned to expect a circuit pack of the mid-span PON RE UNI type.  • When a cardholder provisioned for plug -and-play is equipped with a circuit pack of  the  mid-span PON RE UNI type. Note that the installation of a plug-and-play card may indicate  the presence of a mid-span PON RE UNI port via equipment ID as well as its type attribute,  and indeed may cause the management ONU to instantiate a port-mapping package to specify  the ports precisely.  The management ONU automatically deletes instances of this ME when a cardholder is neither  provisioned to expect a mid-span PON RE UNI circuit pack, nor is it equipped with a mid-span PON  RE UNI circuit pack.  As illustrated in Figure 8.2.10-3, a PPTP RE UNI may share the physical port with an RE upstream  amplifier. The ONU declares a shared configuration through the port -mapping package combined  port table, whose structure defines one ME as the master. It is recommended that the PPTP RE UNI  be the master, with the RE upstream amplifier as a secondary ME.  The administrative state, operational state and ARC attributes of the master ME override similar  attributes in secondary MEs associated with the same port. In the secondary ME, these attributes are 

---- page break ---- present, but cause no action when written and have undefined values when read. The RE upstream  amplifier should use its provisionable upstream alarm thresholds and should declare upstream alarms  as necessary; other isomorphic alarms should be declared by the PPTP RE UNI. The test action should  be addressed to the master ME.  Relationships  An instance of this ME is associated with each instance of a mid-span PON RE S'/R' physical  interface of an RE that includes OEO regeneration in either direction, and it may also be  associated with an RE upstream amplifier.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. This 2 byte  number indicates the physical position of the UNI. The first byte is the slot ID  (defined in clause 9.1.5). The second byte is the port ID, with the range 1..255.  (R) (mandatory) (2 bytes)  NOTE 1 – This ME ID may be identical to that of an RE upstream amplifier if it shares  the same physical slot and port.  Administrative state: This attribute locks (1) and unlocks (0) the functions performed by this  ME. Administrative state is further described in clause A.1.6. (R,  W)  (mandatory) (1 byte)  NOTE 2 – Administrative lock of a PPTP RE UNI results in loss of signal to any  downstream ONUs.  Operational state : This attribute indicates whether the ME is capable of performing its  function. Valid values are enabled (0) and disabled (1). (R) (optional) (1 byte)  ARC: See clause A.1.4.3. (R, W) (optional) (1 byte)  ARC interval: See clause A.1.4.3. (R, W) (optional) (1 byte)  RE ANI -G pointer : This  attribute points to a n RE ANI -G instance. (R, W) (mandatory)  (2 bytes)  Total optical receive signal level table: This table attribute reports a series of measurements  of time averaged received upstream optical signal power. The measurement  circuit should have a temporal response similar to a simple 1  pole low pass  filter, with an effective time constant o f the order of a GTC frame time. Each  table entry has a 2 byte frame counter field (most significant end), and a 2 byte  power measurement field. The frame counter field contains the least significant  16 bits of the superframe counter received closest to the time of the  measurement. The power measurement field is a 2s complement integer  referred to 1  mW (i.e., dBm), with 0.002  dB granularity. (Coding –32768 to  +32767, where 0x00 = 0  dBm, 0x03e8 = +2  dBm, etc.) The RE equipment  should add entries to this table as frequently as is reasonable. The RE should  clear the table once it is read by the OLT. (R) (optional) (4 * N bytes, where N  is the number of measurements present.)   Per burst receive signal level table: This table attribute reports the most recent measurement  of received burst upstream optical signal power. Each table entry has a 2 byte  ONU-ID field (most significant end), and a 2 byte power measurement field.  The power measurement field is a 2s  complement integer referred to 1  mW  (i.e., dBm), with 0.002  dB granularity. (Coding –32768 to +32767, where  0x00 = 0 dBm, 0x03e8 = +2 dBm, etc.) (R) (optional) (4 * N bytes, where N is  the number of distinct ONUs connected to the S'/R' interface.) 

---- page break ---- Lower receive optical threshold: This attribute specifies the optical level that the RE uses to  declare the burst mode low received optical power alarm. Valid values are   –127 dBm (coded as 254) to 0  dBm (coded as 0) in 0.5  dB increments. The  default value 0xFF selects the RE's internal policy. (R, W) (optional) (1 byte)  Upper receive optical threshold: This attribute specifies the optical level that the RE uses to  declare the burst mode high optical power alarm. Valid values are   –127 dBm (coded as 254) to 0  dBm (coded as 0) in 0.5 dB increments. The  default value 0xFF selects the RE's internal policy. (R, W) (optional) (1 byte)  Transmit optical level : This attribute reports the current measurement of the downstream  mean optical launch power. Its value is a 2s complement integer referred to  1 mW (i.e., dBm), with 0.002 dB granularity. (R) (optional) (2 bytes)  Lower transmit power threshold : This attribute specifies the downstream minimum mean  optical launch power at the S'/R' interface that the RE uses to declare t
```
