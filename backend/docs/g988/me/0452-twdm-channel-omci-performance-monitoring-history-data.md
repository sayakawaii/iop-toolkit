# Managed Entity

## Identity
- ME ID: 452
- ME Name: TWDM channel OMCI performance monitoring history data
- Source Section: 9.16.11
- Source Page: 502

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
- Name: OMCI baseline message count
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 4
- Name: OMCI extended message count
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 5
- Name: OMCI MIC error count
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

## Extraction Status
- Status: auto_extracted
- Attributes found: 5
- Review needed: false

## Raw Source

```
9.16.11 TWDM channel OMCI performance monitoring history data  This ME collects OMCI-related PM data associated with the slot/circuit pack, hosting on e or more  ANI-G MEs, for a specific TWDM channel. Instances of this ME are created and deleted by the OLT.  The counters maintained by this ME are characterized as optional in clause 14 of [ITU-T G.989.3].  For a complete discussion of generic PM architecture, refer to clause I.4.  Relationships  An instance of this ME is associated with an instance of TWDM channel ME.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. Through an  identical ID, this ME is implicitly linked to an instance of the TWDM channel  ME. (R, set-by-create) (mandatory) (2 bytes)  Interval end time : This attribute identifies the most recently finished 15  min interval. (R)  (mandatory) (1 byte)  Threshold data 1/2 ID: This attribute points to an instance of the threshold data 1 and 2 MEs  that contains PM threshold values. (R, W, set-by-create) (mandatory) (2 bytes)  OMCI baseline message count: The counter of baseline format OMCI messages directed to  the given ONU. (R) (mandatory) (4 bytes)  OMCI extended message count: The counter of extended format OMCI messages directed  to the given ONU. (R) (mandatory) (4 bytes)  OMCI MIC error count : The counter of OMCI messages received with MIC errors. (R)  (mandatory) (4 bytes)  Actions  Create, delete, get, set  Get current data (optional)  Notifications  Threshold crossing alert  Alarm  number Threshold crossing alert Threshold value  attribute No. (Note)  0 OMCI MIC error count 1  NOTE – This number associates the TCA with the specified threshold value attribute of the threshold data  1/2 managed entities.  10 This clause is intentionally left blank  11 ONU management and control protocol  11.1 Baseline and extended messages  This clause defines two formats for OMCI messages, baseline and extended.  ITU-T GTC based PON systems are free to use either the baseline or the extended OMCI message  format. The baseline format is the default at initialization. Use of the extended format is then  negotiated between the OLT and ONU. 

---- page break ---- The conventions for the use of baseline and extended messages by systems complying with other  Recommendations are for further study.  Baseline messages have 48  byte fixed length PDUs, while extended messages have variable length  PDUs. A receiver that does not support extended messages may therefore reject an extended message  based on nothing more than their length.  Both baseline and extended messages carry a n MIC in their final 4 bytes. This facilitates ad hoc  recovery of both MTs by a receiver. In ITU-T G.984 systems, the MIC is an ITU-T I.363.5 CRC; in  the subsequent ITU-T PON systems, the MIC is a cryptographic hash as specified in the respective  TC layer specification.  Baseline and extended messages are distinguished from one another by the device identifier field,  which is in the same byte location in both MTs. Baseline messages contain device identifier 0x0A,  while extended messages employ device identifier 0x0B.  All G-PON ONUs and OLTs are required to support the baseline format. During initialization, and  whenever the ONU is re -ranged on to the PON, both entities use the baseline format to establish  communications and to negotiate their capabilities. If both endpoints support extended messages, they  may or may not choose to conduct all or some subsequent communications in the extended mess age  set. Baseline messages may be used for any transaction, i.e., any exchange of one or more related  messages such as a get/get-next sequence.  Figure 11.1-1 illustrates the negotiation and the exchange of messages in one or the other message  format.
```
