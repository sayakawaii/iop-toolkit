# Managed Entity

## Identity
- ME ID: 338
- ME Name: PW ATM performance monitoring history data
- Source Section: 9.8.16
- Source Page: 380

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
- Name: Downstream missing packets counter
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 4
- Name: Downstream reordered packets counter
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 5
- Name: Downstream misordered packets  counter
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 6
- Name: Upstream timeout packets counter
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 7
- Name: Upstream transmitted cells counter
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 8
- Name: Upstream dropped cells  counter
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 9
- Name: Upstream received cells  counter
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

## Extraction Status
- Status: auto_extracted
- Attributes found: 9
- Review needed: false

## Raw Source

```
9.8.16 PW ATM performance monitoring history data  This ME collects PM data associated with an ATM pseudowire. Instances of this ME are created and  deleted by the OLT.  For a complete discussion of generic PM architecture, refer to clause I.4. 

---- page break ---- Relationships  An instance of this ME is associated with an instance of the PW ATM configuration data ME.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. Through an  identical ID, this ME is implicitly linked to the instance of the PW ATM   configuration data ME. (R, set-by-create) (mandatory) (2 bytes)  Interval end time : This attribute identifies the most recently finished 15  min interval. (R)  (mandatory) (1 byte)  Threshold data 1/2 ID: This attribute points to an instance of the threshold data  1 ME that  contains PM threshold values. Since no threshold value attribute number  exceeds 7, a threshold data 2 ME is optional. (R, W, set-by-create) (mandatory)  (2 bytes)  Downstream missing packets counter: This attribute counts missing packets, as detected via  control word sequence number gaps. (R) (mandatory) (4 bytes)  Downstream reordered packets counter: This attribute counts packets detected out of  sequence via the control word sequence number, but successfully reordered.   Some implementations may not support this feature. (R) (optional) (4 bytes)  Downstream misordered packets  counter: This attribute counts packets detected out of  order via the control word sequence numbers. (R) (mandatory) (4 bytes)  Upstream timeout packets counter: This attribute counts packets transmitted due to timeout  expiration while attempting to collect cells. (R) (mandatory) (4 bytes)  Upstream transmitted cells counter: This attribute counts transmitted cells. (R) (mandatory)  (4 bytes)  Upstream dropped cells  counter: This attribute counts dropped cells. (R) (mandatory)  (4 bytes)  Upstream received cells  counter: This attribute counts received cells. (R) (mandatory)  (4 bytes)  Actions  Create, delete, get, set  Get current data (optional)  Notifications  Threshold crossing alert  Alarm  number Threshold crossing alert Threshold value attribute No. (Note)  1 Downstream missing packets  1  2 Downstream reordered packets  2  3 Downstream timeout packets 3  4 Upstream dropped cells 4  NOTE – This number associates the TCA with the specified threshold value attribute of the  threshold data 1 managed entity. 

---- page break ---- 
```
