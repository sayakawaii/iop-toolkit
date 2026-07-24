# Managed Entity

## Identity
- ME ID: 440
- ME Name: Time Status Message
- Source Section: 9.12.19
- Source Page: 443

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Get, Set
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: Domain number
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 2
- Name: Flag Field
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 3
- Name: Priority1
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 4
- Name: clockClass
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 5
- Name: offsetScaledLogVariance
- Size: 2 bytes
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 6
- Name: Priority2
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 7
- Name: Grandmaster ID
- Size: 8 bytes
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 8
- Name: Steps removed
- Size: 2 bytes
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 9
- Name: Time source
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: mandatory

## Extraction Status
- Status: auto_extracted
- Attributes found: 9
- Review needed: false

## Raw Source

```
9.12.19 Time status message  This ME provides status and characterization information  for the time -transmitting node and its  grandmaster. An ONU that supports time synchronization automatically creates an instance of this  ME. The best practise is to set all the attributes at the same time.  Relationships  The single instance of this ME is associated with the ONU ME.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. There is  only one instance, number 0. (R) (mandatory) (2 bytes)  Domain number: Using the format of clause 7.1 of [IEEE 1588]. The default value is 0. (R,  W) (mandatory) (1 byte)  Flag Field:  The field  format is  given in the  table. Value  1 represents  "true". (R,  W)  (mandatory)  (1 byte)Bits  Usage  1 (LSB) Leap61  2 Leap59  3 currentUtcOffsetValid  4 PTP timescale  5 Time traceable  6 Frequency traceable  7 Reserved  8 (MSB) Reserved  currentUtcOffset: Provides the UTC offset value between the TAI and UTC timescales  (UTC Offset = TAI – UTC), as specified in clause 7.2.3 of [IEEE 1588].  (R, W) (mandatory) (2 bytes)  Priority1: As specified in clause 7.6.2.2 of [IEEE 1588]. (R, W) (mandatory) (1 byte)  clockClass: Provides the clockClass information denoting the traceability of the time  distributed by the grandmaster clock , as specified in clause 7.6.2.4 of  [IEEE 1588]. (R, W) (mandatory) (1 byte) 

---- page break ---- Accuracy: Indicates the expected accuracy of a clock when it is the grandmaster, as  specified in clause 7.6.2.5 of [IEEE 1588]. (R, W) (mandatory) (1 byte)  offsetScaledLogVariance: Provides the estimate of the time variance, as specified in  clause 7.6.3 of [IEEE 1588]. (R, W) (mandatory) (2 bytes)  Priority2: As specified in clause 7.6.2.3 of [IEEE 1588]. (R, W) (mandatory) (1 byte)  Grandmaster ID : The clockIdentity attribute of the grandmaster, taken from the IEEE  EUI-64 individual assigned numbers. (R, W) (mandatory) (8 bytes)  Steps removed: Provides the number of boundary clocks between the local clock and the  master. (R, W) (mandatory) (2 bytes)  Time source: Indicates the source of time used by the grandmaster clock, as specified in  clause 7.6.2.6 of [IEEE 1588]. (R, W) (mandatory) (1 byte)  Actions  Get, set  Notifications  None.  Alarm  Alarm  number  Alarm Description  0 Clock unlock ITU-T G.781 clock unlock defect: If the status of the  system clock is unlocked, a clock unlock is declared.  The defect is cleared if the status of the system clock  is "locked". Note that "unlocked" is a status of the  clock, not a clock mode.  1 ESMC loss ITU-T G.781 loss of ESMC channel. If no valid ESM  PDU is received for 5 seconds, a loss of ESMC  (LOESMC) is detected. The LOESMC defect is  cleared upon the first ESMC PDU.  2 Time unlock ITU-T G.781.1 time unlock (1PPS+TOD input).  Alarm is declared when the system time is unlocked to  the time source.  3..207 Reserved Reserved for vendor-specific alarms.  208..223 Vendor-specific alarms Not to be standardized  
```
