# Managed Entity

## Identity
- ME ID: 113
- ME Name: xDSL xTU-R performance monitoring history data
- Source Section: 9.7.22
- Source Page: 296

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
- Name: Loss of signal seconds
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 4
- Name: Loss of power seconds
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 5
- Name: Errored seconds
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 6
- Name: Severely errored seconds
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 7
- Name: FEC seconds
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 8
- Name: Unavailable seconds
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 9
- Name: Error-free bits counter
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 10
- Name: Interval end time
- Size: 1 byte
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 11
- Name: Threshold data 1/2 ID
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 12
- Name: Corrected blocks
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 13
- Name: Uncorrected blocks
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 14
- Name: Transmitted blocks
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 15
- Name: Received blocks
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 16
- Name: Code violations
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 17
- Name: Interval end time
- Size: 1 byte
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 18
- Name: Threshold data 1/2 ID
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 19
- Name: Corrected blocks
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 20
- Name: Uncorrected blocks
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 21
- Name: Transmitted blocks
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 22
- Name: Received blocks
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 23
- Name: Forward error corrections
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: mandatory

## Extraction Status
- Status: auto_extracted
- Attributes found: 23
- Review needed: false

## Raw Source

```
9.7.22 xDSL xTU-R performance monitoring history data  This ME collects PM data of the xTU -C to xTU -R path as seen from the xTU -R. Instances of this  ME are created and deleted by the OLT.  For a complete discussion of generic PM architecture, refer to clause I.4.  Relationships  An instance of this ME is associated with an xDSL UNI.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. Through an  identical ID, this ME is implicitly linked to an instance of the PPTP xDSL UNI  part 1. (R, set-by-create) (mandatory) (2 bytes)  Interval end time : This attribute identifies the most recently finished 15  min interval. (R)  (mandatory) (1 byte)  Threshold data 1/2 ID: This attribute points to an instance of the threshold data  1 ME that  contains PM threshold values. Since no threshold value attribute number  exceeds 7, a threshold data 2 ME is optional. (R, W, set-by-create) (mandatory)  (2 bytes) 

---- page break ---- Loss of frame seconds: (R) (mandatory) (2 bytes)  Loss of signal seconds: (R) (mandatory) (2 bytes)  Loss of power seconds: (R) (mandatory) (2 bytes)  Errored seconds: This attribute counts 1 s intervals with one or more far end block error  (FEBE) anomalies summed ov er all transmitted bearer channels, or one or  more LOS-FE defects, or one or more RDI defects, or one or more LPR -FE  defects. (R) (mandatory) (2 bytes)  Severely errored seconds : This attribute counts severely errored seconds (SES -LFE). An  SES is declared if, during a 1 s interval, 18 or more FEBE anomalies were  reported across the totality of bearer channels, or there were one or more far - end LOS defects, one or more RDI defects or one or more LPR-FE defects.  If the relevant Recommendation ([ITU -T G.992.3], [ITU -T G.992.5] or  [ITU-T G.993.2]) supports a 1 s normalized CRC -8 anomaly counter  increment, the 1 s SES counter follows this value instead of counting FEBE  anomalies directly.  If a CRC is applied for multiple bearer channels, then each related FEBE  anomaly is counted only once for the whole set of related bearer channels.  (R) (mandatory) (2 bytes)  FEC seconds: This attribute counts seconds during which at least one uncorrectable FEC  codeword was received. (R) (mandatory) (2 bytes)  Unavailable seconds : This attribute counts 1 s intervals during which the far -end xDSL  termination is unavailable.  The far -end xDSL termination becomes unavailable at the onset of 10  contiguous SES -LFEs. The 10 SES-LFEs are included in unavailable time.  Once unavailable, the far -end line becomes available at the onset of 10  contiguous seconds with no SES -LFEs. The 10 s with no SES -LFEs are  excluded from unavailable time. Some attribute counts are inhibited during  unavailability – see clause 7.2.7.13 of [ITU-T G.997.1].  (R) (mandatory) (2 bytes)  "leftr" defect seconds: If retransmission is used, this parameter is a count of the seconds  with a near-end ''leftr'' defect present – see clause 7.2.1.1.6 of  [ITU-T G.997.1]. (R) (optional) (2 bytes)  Error-free bits counter: If retransmission is used, this parameter is a count of the number  of error-free bits passed over the β1 reference point, divided by 216 – see  clause 7.2.1.1.7 of [ITU-T G.997.1]. (R) (optional) (4 bytes)  Minimum error-free throughput (MINEFTR): If retransmission is used, this parameter is  the minimum error-free throughput in bits per second – see clause 7.2.1.1.8 of  [ITU-T G.997.1]. (R) (optional) (4 bytes)  Actions  Create, delete, get, set  Get current data (optional) 

---- page break ---- Notifications  Threshold crossing alert  Alarm  number Threshold crossing alert Threshold value attribute No. (Note)  0 Loss of frame seconds 1  1 Loss of signal seconds 2  2 Loss of power seconds 3  3 Errored seconds 4  4 Severely errored seconds 5  5 FEC seconds 6  6 Unavailable seconds 7  7 "leftr" defect seconds 8  NOTE – This number associates the TCA with the specified threshold value attribute of the  threshold data 1 managed entity.  9.7.23 xDSL xTU-C channel performance monitoring history data  This ME collects PM data of an xTU-C to xTU-R channel as seen from the xTU-C. Instances of this  ME are created and deleted by the OLT.  For a complete discussion of generic PM architecture, refer to clause I.4.  Relationships  An instance of this ME is associated with an xDSL bearer channel. Several instances may  therefore be associated with an xDSL UNI.  Attributes  Managed entity  ID: This attribute uniquely identifies each instance of this ME. The two  MSBs of the first byte are the bearer channel ID. Excluding the first 2 bits of  the first byte, the remaining part of the ME ID is identical to that of this ME's  parent PPTP xDSL UNI part 1. (R, set-by-create) (mandatory) (2 bytes)  Interval end time : This attribute identifies the most recently finished 15  min interval. (R)  (mandatory) (1 byte)  Threshold data 1/2 ID: This attribute points to an instance of the threshold data  1 ME that  contains PM threshold values. Since no threshold value attribute number  exceeds 7, a threshold data 2 ME is optional. (R, W, set-by-create) (mandatory)  (2 bytes)  Corrected blocks: This attribute counts blocks received with errors that were corrected on  this channel. (R) (mandatory) (4 bytes)  Uncorrected blocks: This attribute counts blocks received with uncorrectable errors on this  channel. (R) (mandatory) (4 bytes)  Transmitted blocks: This attribute counts encoded blocks transmitted on this channel. (R)  (mandatory) (4 bytes)  Received blocks : This attribute counts encoded blocks received on this channel. (R)  (mandatory) (4 bytes)  Code violations : This attribute counts CRC -8 anomalies in the bearer channel. (R)  (mandatory) (2 bytes) 

---- page break ---- Forward error corrections: This attribute counts FEC anomalies in the bearer channel. (R)  (mandatory) (2 bytes)  Actions  Create, delete, get, set  Get current data (optional)  Notifications  Threshold crossing alert  Alarm  number Th
```
