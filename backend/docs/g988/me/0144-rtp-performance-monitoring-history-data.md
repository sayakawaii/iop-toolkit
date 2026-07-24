# Managed Entity

## Identity
- ME ID: 144
- ME Name: RTP performance monitoring history data
- Source Section: 9.9.13
- Source Page: 406

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
- Name: RTP errors
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 4
- Name: Packet loss
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 5
- Name: Buffer underflows
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 6
- Name: Buffer overflows
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

## Extraction Status
- Status: auto_extracted
- Attributes found: 6
- Review needed: false

## Raw Source

```
9.9.13 RTP performance monitoring history data  This ME collects PM data related to an RTP session. Instances of this ME are created and deleted by  the OLT.  For a complete discussion of generic PM architecture, refer to clause I.4.  Relationships  An instance of this ME is associated with an instance of the PPTP POTS UNI ME.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. Through an  identical ID, this ME is implicitly linked to an instance of the PPTP POTS UNI  ME. (R, set-by-create) (mandatory) (2 bytes)  Interval end time : This attribute identifies the most recently finished 15  min interval. (R)  (mandatory) (1 byte)  Threshold data 1/2 ID: This attribute points to an instance of the threshold data  1 ME that  contains PM threshold values. Since no threshold value attribute number  exceeds 7, a threshold data 2 ME is optional. (R, W, set-by-create) (mandatory)  (2 bytes)  RTP errors: This attribute counts RTP packet errors. (R) (mandatory) (4 bytes)  Packet loss: This attribute represents the fraction of packets lost. This attribute is calculated  at the end of the 15  min interval, and is undefined under the get current data  action. The value 0 indicates no packet loss, scaling linearly to 0xFFFF FFFF  to indicate 100% packet loss (zero divided by zero is defined to be zero). (R)  (mandatory) (4 bytes) 

---- page break ---- Maximum jitter : This attribute is a high water -mark that represents the maximum jitter  identified during the measured interval, expressed in RTP timestamp units. (R)  (mandatory) (4 bytes)  Maximum time between real-time transport control protocol ( RTCP) packets: This  attribute is a high water -mark that represents the maximum time between  RTCP packets during the measured interval, in milliseconds. (R) (mandatory)  (4 bytes)  Buffer underflows : This attribute counts the number of times the reassembly buffer  underflows. In the case of continuous underflow caused by a loss of IP packets,  a single buffer underflow should be counted. If the IW function is implemented  with multiple buffers, such as a packet level buffer and a bit level buffer, then  the underflow of either buffer increments this counter. (R) (mandatory)  (4 bytes)  Buffer overflows: This attribute counts the number of times the reassembly buffer overflows.  If the IW function is implemented with multiple buffers, such as a packet level  buffer and a bit level buffer, then the overflow of either buffer increments this  counter. (R) (mandatory) (4 bytes)  Actions  Create, delete, get, set  Get current data (optional)  Notifications  Threshold crossing alert  Alarm  number Threshold crossing alert Threshold value attribute No. (Note 2)  0 RTP errors  1  1 Packet loss(Note 1) 2  2 Maximum jitter 3  3 Max time between RTCP packets 4  4 Buffer underflows 6  5 Buffer overflows 7  NOTE 1 – Since packet loss is undefined until the end of the interval, this TCA may only be issued  at the interval boundary, whereupon it is then immediately cleared.  NOTE 2 – This number associates the TCA with the specified threshold value attribute of the  threshold data 1 managed entity.  
```
