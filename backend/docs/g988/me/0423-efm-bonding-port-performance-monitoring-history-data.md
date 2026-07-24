# Managed Entity

## Identity
- ME ID: 423
- ME Name: EFM bonding port performance monitoring history data
- Source Section: 9.7.46
- Source Page: 330

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
- Name: Rx frames
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 4
- Name: Tx frames
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 5
- Name: Rx bytes
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 6
- Name: Tx bytes
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 7
- Name: Tx discarded frames
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 8
- Name: Tx discarded bytes
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

## Extraction Status
- Status: auto_extracted
- Attributes found: 8
- Review needed: false

## Raw Source

```
9.7.46 EFM bonding port performance monitoring history data  This ME collects PM data as seen at the xTU-C. Instances of this ME are created and deleted by the  OLT.  Relationships  An instance of this ME is associated with an xDSL UNI.  Attributes  Managed entity ID : This attribute uniquely identifies each instance of this ME. The two  MSBs of the first byte are the bearer channel ID. Excluding the first 2 bits of  the first byte, the remaining part of the ME ID is identical to that of this ME's  parent PPTP xDSL UNI part 1. (R, set-by-create) (mandatory) (2 bytes)  Interval end time : This attribute identifies the most recently finished 15  min interval. (R)  (mandatory) (1 byte)  Threshold data 1/2 ID: This attribute points to an instance of the threshold data 1 and 2 MEs  that contain PM threshold values. (R, W, set-by-create) (mandatory) (2 bytes)  Rx frames: Number of Ethernet frames received over this port. (R) (mandatory) (4 bytes)  Tx frames: Number of Ethernet frames transmitted over this port. (R) (mandatory) (4 bytes)  Rx bytes : Number of bytes contained in the Ethernet frames received over this port. (R)  (mandatory) (4 bytes)  Tx bytes: Number of bytes contained in the Ethernet frames transmitted over this port. (R)  (mandatory) (4 bytes)  Tx discarded frames : Number of Ethernet frames discarded by the port transmit function.  (R) (mandatory) (4 bytes)  Tx discarded bytes: Number of bytes contained in the Ethernet frames discarded by the port  transmit function. (R) (mandatory) (4 bytes)  Actions  Create, delete, get, set  Get current data (optional) 

---- page break ---- Notifications  Threshold crossing alert  Alarm number Threshold crossing alert Threshold value attribute No.  (Note)  0 Bad fragments 1  1 Lost fragments 2  2 Lost starts 3  3 Lost ends 4  4-207 Reserved   NOTE – This number associates the TCA with the specified threshold value attribute of the  threshold data 1/2 managed entities.  
```
