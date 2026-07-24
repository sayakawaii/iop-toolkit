# Managed Entity

## Identity
- ME ID: 408
- ME Name: xDSL xTU-C performance monitoring history data part 2
- Source Section: 9.7.21
- Source Page: 294

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
- Name: Loss of frame seconds
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 4
- Name: Loss of signal seconds
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 5
- Name: Loss of link seconds
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 6
- Name: Loss of power seconds
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 7
- Name: Severely errored seconds
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 8
- Name: Line initializations
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 9
- Name: Failed line initializations
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 10
- Name: Short initializations
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 11
- Name: Failed short initializations
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 12
- Name: FEC seconds
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 13
- Name: Unavailable seconds
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 14
- Name: SOS success count, near end
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 15
- Name: Interval end time
- Size: 1 byte
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 16
- Name: Threshold data 1/2 ID
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 17
- Name: Loss of signal seconds
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 18
- Name: Loss of power seconds
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 19
- Name: Errored seconds
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 20
- Name: Severely errored seconds
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 21
- Name: FEC seconds
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 22
- Name: Unavailable seconds
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 23
- Name: Error-free bits counter
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 24
- Name: Interval end time
- Size: 1 byte
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 25
- Name: Threshold data 1/2 ID
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 26
- Name: Corrected blocks
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 27
- Name: Uncorrected blocks
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 28
- Name: Transmitted blocks
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 29
- Name: Received blocks
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 30
- Name: Code violations
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 31
- Name: Interval end time
- Size: 1 byte
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 32
- Name: Threshold data 1/2 ID
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 33
- Name: Corrected blocks
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 34
- Name: Uncorrected blocks
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 35
- Name: Transmitted blocks
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 36
- Name: Received blocks
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 37
- Name: Forward error corrections
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: mandatory

## Extraction Status
- Status: auto_extracted
- Attributes found: 37
- Review needed: false

## Raw Source

```
9.7.21 xDSL xTU-C performance monitoring history data part 1  This ME collects PM data on the xTU -C to xTU-R path as seen from the xTU -C. Instances of this  ME are created and deleted by the OLT.  For a complete discussion of generic PM architecture, refer to clause I.4.  Relationships  An instance of this ME is associated with an xDSL UNI.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. Through an  identical ID, this ME is implicitly linked to an instance of the PPTP xDSL UNI  part 1. (R, set-by-create) (mandatory) (2 bytes)  Interval end time : This attribute identifies the most recently finished 15  min interval. (R)  (mandatory) (1 byte)  Threshold data 1/2 ID: This attribute points to an instance of the threshold data 1 and 2 MEs  that contain PM threshold values. (R, W, set-by-create) (mandatory) (2 bytes)  Loss of frame seconds: (R) (mandatory) (2 bytes)  Loss of signal seconds: (R) (mandatory) (2 bytes)  Loss of link seconds: (R) (mandatory) (2 bytes)  Loss of power seconds: (R) (mandatory) (2 bytes) 

---- page break ---- Errored seconds (ES): This attribute counts 1 s intervals with one or more CRC-8 anomalies  summed over all received bearer channels, or one or more loss of signal (LOS)  defects, or one or more SEF defects, or one or more LPR defects. (R)  (mandatory) (2 bytes)  Severely errored seconds: This attribute counts severely errored seconds (SES -L). An SES  is declared if, during a 1 s interval, there were 18 or more CRC-8 anomalies in  one or more of the received bearer channels, or one or more LOS defects, or  one or more SEF defects, or one or more LPR defects.  If the relevant Recommendation ([ITU -T G.992.3], [ITU -T G.992.5]  or  [ITU-T G.993.2]) supports a 1 s normalized CRC -8 anomaly counter  increment, the 1 s SES counter follows this value instead of counting CRC -8  anomalies directly.  If a common CRC is applied over multiple bearer channels, then each related  CRC-8 anomaly is counted only once for the whole set of bearer channels over  which the CRC is applied.  (R) (mandatory) (2 bytes)  Line initializations: This attribute counts the total number of full initializations attempted on  the line, both successful and failed. (R) (mandatory) (2 bytes)  Failed line initializations: This attribute counts the total number of failed full initializations  during the accumulation period. A f ailed full initialization occurs when  showtime is not reached at the end of the full initialization procedure. (R)  (mandatory) (2 bytes)  Short initializations : This attribute counts the total number of fast retrains or short  initializations attempted on the line, successful and failed. Fast retrain is  defined in [ITU-T G.992.2]. Short initialization is defined in [ITU-T G.992.3]  and [ITU-T G.992.4]. (R) (optional) (2 bytes)  Failed short initializations : This attribute counts the total number of failed fast retrains or  short initializations during the accumulation period, e.g., when:  – a CRC error is detected;  – a timeout occurs;  – a fast retrain profile is unknown.  (R) (optional) (2 bytes)  FEC seconds: This attribute counts seconds during which at least one uncorrectable FEC  codeword was received. (R) (mandatory) (2 bytes)  Unavailable seconds : This attribute counts 1 s intervals during which the xDSL UNI is  unavailable. The line becomes unavailable at the onset of 10 contiguous SES- Ls. The 10 SES-Ls are included in unavailable time. Once unavailable, the line  becomes available at the onset of 10 contiguous seconds that are not severely  errored. The 10 s with no SES -Ls are excluded from unavailable time. Some  attribute counts are inhibited during unavailability – see clause 7.2.7.13 of  [ITU-T G.997.1]. (R) (mandatory) (2 bytes)  SOS success count, near end : The SOS -SUCCESS-NE attribute is a count of the total  number of successful SOS procedures initiated by the near-end xTU on the line  during the accumulation period. Successful SOS is defined in clause 12.1.4 of  [ITU-T G.993.2]. (R) (optional) (2 bytes) 

---- page break ---- SOS success count, far end: The SOS-SUCCESS-FE attribute is a count of the total number  of successful SOS procedures initiated by the far -end xTU on the line during  the accumulation period. Successful SOS is defined in clause 12.1.4 of  [ITU-T G.993.2]. (R) (optional) (2 bytes)  Actions  Create, delete, get, set  Get current data (optional)  Notifications  Threshold crossing alert  Alarm number Threshold crossing alert Threshold value attribute No.  (Note)  0 Loss of frame seconds 1  1 Loss of signal seconds 2  2 Loss of link seconds 3  3 Loss of power seconds 4  4 Errored seconds 5  5 Severely errored seconds 6  6 Line initializations 7  7 Failed line initializations 8  8 Short initializations 9  9 Failed short initializations 10  10 FEC seconds 11  11 Unavailable seconds 12  NOTE – This number associates the TCA with the specified threshold value attribute of the  threshold data 1/2 managed entities.  9.7.22 xDSL xTU-R performance monitoring history data  This ME collects PM data of the xTU -C to xTU -R path as seen from the xTU -R. Instances of this  ME are created and deleted by the OLT.  For a complete discussion of generic PM architecture, refer to clause I.4.  Relationships  An instance of this ME is associated with an xDSL UNI.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. Through an  identical ID, this ME is implicitly linked to an instance of the PPTP xDSL UNI  part 1. (R, set-by-create) (mandatory) (2 bytes)  Interval end time : This attribute identifies the most recently finished 15  min interval. (R)  (mandatory) (1 byte)  Threshold data 1/2 ID: This attribute points to an instance of the threshold data  1 ME that  contains PM threshold values. Since no threshold value attribute number  exceeds 7, a threshold data 2 ME is optional. (R, W, set-by-create) (mandatory)  (2 bytes) 

---- page break ---- Loss of frame se
```
