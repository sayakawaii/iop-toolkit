# Managed Entity

## Identity
- ME ID: 422
- ME Name: EFM bonding link performance monitoring history data
- Source Section: 9.7.45
- Source Page: 329

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
- Name: Rx errored fragments
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 4
- Name: Rx small fragments
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 5
- Name: Rx large fragments
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 6
- Name: Rx discarded fragments
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 7
- Name: Rx FCS errors
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 8
- Name: Rx coding errors
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 9
- Name: Rx fragments
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 10
- Name: Tx fragments
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

## Extraction Status
- Status: auto_extracted
- Attributes found: 10
- Review needed: false

## Raw Source

```
9.7.45 EFM bonding link performance monitoring history data  This ME collects PM data as seen at the xTU-C. Instances of this ME are created and deleted by the  OLT.  Relationships  An instance of this ME is associated with an xDSL UNI.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. Through an  identical ID, this ME is implicitly linked to an instance of the EFM bonding  link. (R, set-by-create) (mandatory) (2 bytes)  Interval end time : This attribute identifies the most recently finished 15  min interval. (R)  (mandatory) (1 byte)  Threshold data 1/2 ID: This attribute points to an instance of the threshold data 1 and 2 MEs  that contain PM threshold values. (R, W, set-by-create) (mandatory) (2 bytes)  Rx errored fragments: Clause 45.2.3.29 of [IEEE 802.3]. (R) (mandatory) (4 bytes)  Rx small fragments: Clause 45.2.3.30 of [IEEE 802.3]. (R) (mandatory) (4 bytes)  Rx large fragments: Clause 45.2.3.31 of [IEEE 802.3]. (R) (mandatory) (4 bytes)  Rx discarded fragments: Clause 45.2.3.32 of [IEEE 802.3]. (R) (mandatory) (4 bytes)  Rx FCS errors: Clause 45.2.6.11 of [IEEE 802.3]. (R) (mandatory) (4 bytes)  Rx coding errors: Clause 45.2.6.12 of [IEEE 802.3]. (R) (mandatory) (4 bytes)  Rx fragments: Number of fragments received over this link. (R) (mandatory) (4 bytes)   Tx fragments: Number of fragments transmitted over this link. (R) (mandatory) (4 bytes)  Actions  Create, delete, get, set  Get current data (optional) 

---- page break ---- Notifications  Threshold crossing alert  Alarm number Threshold crossing alert Threshold value attribute No.  (Note)  0 Rx errored fragments 1  1 Rx small fragments 2  2 Rx large fragments 3  3 Rx discarded fragments 4  4 Rx FCS errors 5  5 Rx coding errors 6  6-207 Reserved   NOTE – This number associates the TCA with the specified threshold value attribute of the  threshold data 1/2 managed entities.  
```
