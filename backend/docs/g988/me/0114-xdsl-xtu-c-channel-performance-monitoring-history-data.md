# Managed Entity

## Identity
- ME ID: 114
- ME Name: xDSL xTU-C channel performance monitoring history data
- Source Section: 9.7.23
- Source Page: 298

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
- Name: Threshold data 1/2 ID
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 3
- Name: Corrected blocks
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 4
- Name: Uncorrected blocks
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 5
- Name: Transmitted blocks
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 6
- Name: Received blocks
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 7
- Name: Code violations
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 8
- Name: Interval end time
- Size: 1 byte
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 9
- Name: Threshold data 1/2 ID
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 10
- Name: Corrected blocks
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 11
- Name: Uncorrected blocks
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 12
- Name: Transmitted blocks
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 13
- Name: Received blocks
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 14
- Name: Forward error corrections
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: mandatory

## Extraction Status
- Status: auto_extracted
- Attributes found: 14
- Review needed: false

## Raw Source

```
9.7.23 xDSL xTU-C channel performance monitoring history data  This ME collects PM data of an xTU-C to xTU-R channel as seen from the xTU-C. Instances of this  ME are created and deleted by the OLT.  For a complete discussion of generic PM architecture, refer to clause I.4.  Relationships  An instance of this ME is associated with an xDSL bearer channel. Several instances may  therefore be associated with an xDSL UNI.  Attributes  Managed entity  ID: This attribute uniquely identifies each instance of this ME. The two  MSBs of the first byte are the bearer channel ID. Excluding the first 2 bits of  the first byte, the remaining part of the ME ID is identical to that of this ME's  parent PPTP xDSL UNI part 1. (R, set-by-create) (mandatory) (2 bytes)  Interval end time : This attribute identifies the most recently finished 15  min interval. (R)  (mandatory) (1 byte)  Threshold data 1/2 ID: This attribute points to an instance of the threshold data  1 ME that  contains PM threshold values. Since no threshold value attribute number  exceeds 7, a threshold data 2 ME is optional. (R, W, set-by-create) (mandatory)  (2 bytes)  Corrected blocks: This attribute counts blocks received with errors that were corrected on  this channel. (R) (mandatory) (4 bytes)  Uncorrected blocks: This attribute counts blocks received with uncorrectable errors on this  channel. (R) (mandatory) (4 bytes)  Transmitted blocks: This attribute counts encoded blocks transmitted on this channel. (R)  (mandatory) (4 bytes)  Received blocks : This attribute counts encoded blocks received on this channel. (R)  (mandatory) (4 bytes)  Code violations : This attribute counts CRC -8 anomalies in the bearer channel. (R)  (mandatory) (2 bytes) 

---- page break ---- Forward error corrections: This attribute counts FEC anomalies in the bearer channel. (R)  (mandatory) (2 bytes)  Actions  Create, delete, get, set  Get current data (optional)  Notifications  Threshold crossing alert  Alarm  number Threshold crossing alert Threshold value attribute No. (Note)  0 Corrected blocks 1  1 Uncorrected blocks 2  2 Code violations 3  3 Forward error corrections 4  NOTE – This number associates the TCA with the specified threshold value attribute of the  threshold data 1 managed entity.  9.7.24 xDSL xTU-R channel performance monitoring history data  This ME collects PM data of the xTU-C to xTU-R channel as seen from the xTU-R. Instances of this  ME are created and deleted by the OLT.  For a complete discussion of generic PM architecture, refer to clause I.4.  Relationships  An instance of this ME is associated with an xDSL bearer channel. Several instances may  therefore be associated with an xDSL UNI.  Attributes  Managed entity  ID: This attribute uniquely identifies each instance of this ME. The two  MSBs of the first byte are the bearer channel ID. Excluding the first 2 bits of  the first byte, the remaining part of the ME ID is identical to that of this ME's  parent PPTP xDSL UNI part 1. (R, set-by-create) (mandatory) (2 bytes)  Interval end time : This attribute identifies the most recently finished 15  min interval. (R)  (mandatory) (1 byte)  Threshold data 1/2 ID: This attribute points to an instance of the threshold data  1 ME that  contains PM threshold values. Since no threshold value attribute number  exceeds 7, a threshold data 2 ME is optional. (R, W, set-by-create) (mandatory)  (2 bytes)  Corrected blocks: This attribute counts blocks received with errors that were corrected on  this channel. (R) (mandatory) (4 bytes)  Uncorrected blocks: This attribute counts blocks received with uncorrectable errors on this  channel. (R) (mandatory) (4 bytes)  Transmitted blocks: This attribute counts encoded blocks transmitted on this channel. (R)  (mandatory) (4 bytes)  Received blocks : This attribute counts encoded blocks received on this channel. (R)  (mandatory) (4 bytes) 

---- page break ---- Code violations: This attribute counts FEBE anomalies reported in the downstream bearer  channel. If the CRC is applied over multiple bearer channels, then each related  FEBE anomaly increments each of the counters related to the individual bearer  channels. (R) (mandatory) (2 bytes)  Forward error corrections : This attribute counts FFEC anomalies reported in the  downstream bearer channel. If FEC is applied over multiple bearer channels,  each related FFEC anomaly increments each of the counters related to the  individual bearer channels. (R) (mandatory) (2 bytes)  Actions  Create, delete, get, set  Get current data (optional)  Notifications  Threshold crossing alert  Alarm  number Threshold crossing alert Threshold value attribute No. (Note)  0 Corrected blocks 1  1 Uncorrected blocks 2  2 Code violations 3  3 Forward error corrections 4  NOTE – This number associates the TCA with the specified threshold value attribute of the  threshold data 1 managed entity.  
```
