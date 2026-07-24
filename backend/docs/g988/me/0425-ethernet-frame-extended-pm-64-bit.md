# Managed Entity

## Identity
- ME ID: 425
- ME Name: Ethernet frame extended PM 64 bit
- Source Section: 9.3.34
- Source Page: 215

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Create, Delete, Get, Set
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: Interval end time
- Size: 1 byte
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 2
- Name: Control block
- Size: 16 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 3
- Name: Drop events
- Size: 8 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 4
- Name: Octets
- Size: 8 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 5
- Name: Frames
- Size: 8 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 6
- Name: Broadcast frames
- Size: 8 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 7
- Name: Multicast frames
- Size: 8 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 8
- Name: CRC errored frames
- Size: 8 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 9
- Name: Undersize frames
- Size: 8 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 10
- Name: Oversize frames
- Size: 8 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 11
- Name: Frames 64 octets
- Size: 8 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 12
- Name: Frames 65 to 127 octets
- Size: 8 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 13
- Name: Frames 256 to 511 octets
- Size: 8 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 14
- Name: Frames 512 to 1023 octets
- Size: 8 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 15
- Name: Frames 1024 to 1518 octets
- Size: 8 bytes
- Format: needs_review
- Access: R
- Category: mandatory

## Extraction Status
- Status: auto_extracted
- Attributes found: 15
- Review needed: false

## Raw Source

```
9.3.34 Ethernet frame extended PM 64 bit  This ME collects some of the PM data at a point where an Ethernet flow can be observed. It is based  on the Etherstats group of [IETF RFC 2819] and [IETF RFC 2863]. Instances of this ME are created  and deleted by the OLT. References to received frames are to be interpreted as the number of frames  entering the monitoring point in the direction specified by the control block.  For a complete discussion of generic PM architecture, refer to clause I.4.  Relationships  An instance of this ME may be associated with an instance of an ME at any Ethernet  interface within the ONU. The specific ME is identified in the control block attribute.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. To facilitate  discovery, it is encouraged to identify instances sequentially starting with 1.  (R, set-by-create) (mandatory) (2 bytes)  Interval end time : This attribute identifies the most recently finished 15  min interval. If  continuous accumulation is enabled in the control block, this attribute is not  used and has the fixed value 0. (R) (mandatory) (1 byte)  Control block: This attribute contains fields defined as follows.  Threshold data 64 bit ID: (2 bytes). This attribute points to an instance of the  threshold data  64-bit ME that contains PM threshold values. When PM is  collected on a continuously  running basis, rather than in 15  min intervals,  counter thresholds should not be established. There is no mechanism to clear 

---- page break ---- a TCA, and any counter parameter may eventually be expected to cross any  given threshold value.  Parent ME class : (2 bytes) . This field contains the enumerated value of the  ME class of the PM ME's parent. Together with the parent ME instance field,  this permits a given PM ME to be associated with any OMCI ME. The  supported ME classes are as follows.  46 MAC bridge configuration data  47 MAC bridge port configuration data  11 Physical path termination point Ethernet UNI  98 Physical path termination point xDSL UNI part 1  266 GEM IW termination point  281 Multicast GEM IW termination point  329 Virtual Ethernet interface point  162 Physical path termination point MoCA UNI  Parent ME instance : (2 bytes) . This field identifies the specific parent ME  instance to which the PM ME is attached.  Accumulation disable: (2 bytes). This bit field allows PM accumulation to be  disabled; refer to Table 9.3.32-1. The default value 0 enables PM collection. If  bit 15 is set to 1, no PM is collected by this ME instance. If bit 15 = 0 and any  of bits 14..1 are set to 1, PM collection is inhibited for the attributes indicated  by the 1 bits. Inhibiting PM collection does not change the value of a PM  attribute, but if PM is accumulated in 15  min intervals, the value is lost at the  next 15 min interval boundary.  Bit 16 is an action bit that always reads back as 0. When written to 1, it resets  all PM attributes in the ME, and clears any TCAs that may be outstanding.  TCA disable: (2 bytes) . Also clarified in Table 9.3.32 -1, this field permits  TCAs to be inhibited, either individually or for the complete ME instance. As  with the accumulation disable field, the default value 0 enables TCAs, and  setting the global disable bit overrides the settings of the individual thresholds.  Unlike the accumulation disable field, the bits are mapped to the thresholds  defined in the associated threshold data 1 and 2 ME instances. When the global  or attribute-specific value changes from 0 to 1, outstanding TCAs are cleared,  either for the ME instance globally or for the individual disabled threshold.  These bits affect only notifications, not the underlying parameter accumulation  or storage.  If the threshold data 64 bit ID attribute does not contain a valid pointer, this  field is not meaningful.  Thresholds should be used with caution if PM attributes are accumulated  continuously.  Control fields : (2 bytes) . This field is a bit map whose values govern the  behaviour of the PM ME. Bits are assigned as follows:  Bit 1 (LSB) The value 1 specifies continuous accumulation, regardless  of 15 min intervals. There is no concept of current and  historic accumulators; get and get current data (if  supported) both return current values. The value 0  specifies 15 min accumulators exactly like those of  classical PM.  Bit 2 This bit indicates directionality for the collection of data.  The value 0 indicates that data are to be collected for 

---- page break ---- upstream traffic. The value 1 indicates that data are to be  collected for downstream traffic.  Bits 3..14 Reserved, should be set to 0 by the OLT and ignored by  the ONU.  Bit 15 When this bit is 1, the P bits of the TCI field are used to  filter the PM data collected. The value 0 indicates that PM  is collected without regard to P bits.  Bit 16 When this bit is 1, the VID bits of the TCI field are used to  filter the PM data collected. The value 0 indicates that PM  is collected without regard to VID.  TCI: (2 bytes). This field contains the value optionally used as a filter for the  PM data collected, under the control of bits 15..16 of the control fields. This  value is matched to the outer tag of a frame. Untagged frames are not counted  when this field is used.  Reserved: (2 bytes). Not used; should be set to 0 by the OLT and ignored by  the ONU.  (R, W, set-by-create) (mandatory) (16 bytes)  Drop events: The total number of events in which frames were dropped due to lack of  resources. This is not necessarily the number of frames dropped; it is the  number of times this event was detected. (R) (mandatory) (8 bytes)  Octets: The total number of octets received, including those in bad frames, excluding  framing bits, but including FCS. (R) (mandatory) (8 bytes)  Frames: The total number of frames received, including bad frames, broadcast frames  and multicast frames. (R) (mandatory) (8 bytes)  Broadcast frames : The total numbe
```
