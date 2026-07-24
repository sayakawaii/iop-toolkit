# Managed Entity

## Identity
- ME ID: 420
- ME Name: EFM bonding group performance monitoring history data
- Source Section: 9.7.43
- Source Page: 327

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
- Name: Rx bad fragments
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 4
- Name: Rx lost fragments
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 5
- Name: Rx lost starts
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 6
- Name: Rx lost ends
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 7
- Name: Rx frames
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 8
- Name: Tx frames
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 9
- Name: Rx bytes
- Size: 8 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 10
- Name: Tx bytes
- Size: 8 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 11
- Name: Tx discarded frames
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 12
- Name: Tx discarded bytes
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

## Extraction Status
- Status: auto_extracted
- Attributes found: 12
- Review needed: false

## Raw Source

```
9.7.43 EFM bonding group performance monitoring history data  This ME collects PM data as seen at the xTU-C. Instances of this ME are created and deleted by the  OLT.  Relationships  An instance of this ME is associated with an xDSL UNI.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. Through an  identical ID, this ME is implicitly linked to an instance of the EFM bonding  group. (R, set-by-create) (mandatory) (2 bytes)  Interval end time : This attribute identifies the most recently finished 15  min interval. (R)  (mandatory) (1 byte)  Threshold data 1/2 ID: This attribute points to an instance of the threshold data 1 and 2 MEs  that contain PM threshold values. (R, W, set-by-create) (mandatory) (2 bytes)  Rx bad fragments: Clause 45.2.3.33 of [IEEE 802.3]. (R) (mandatory) (4 bytes)  Rx lost fragments: Clause 45.2.3.34 of [IEEE 802.3]. (R) (mandatory) (4 bytes)  Rx lost starts: Clause 45.2.3.35 of [IEEE 802.3]. (R) (mandatory) (4 bytes)  Rx lost ends: Clause 45.2.3.36 of [IEEE 802.3]. (R) (mandatory) (4 bytes)  Rx frames: Number of Ethernet frames received over this group. (R) (mandatory) (4 bytes)  Tx frames: Number of Ethernet frames transmitted over this group. (R) (mandatory) (4 bytes)  Rx bytes: Number of bytes contained in the Ethernet frames received over this group. (R)  (mandatory) (8 bytes)  Tx bytes: Number of bytes contained in the Ethernet frames transmitted over this group. (R)  (mandatory) (8 bytes)  Tx discarded frames: Number of Ethernet frames discarded by the group transmit function.  (R) (mandatory) (4 bytes)  Tx discarded bytes: Number of bytes contained in the Ethernet frames discarded by the group  transmit function. (R) (mandatory) (4 bytes)  Actions  Create, delete, get, set  Get current data (optional) 

---- page break ---- Notifications  Threshold crossing alert  Alarm number Threshold crossing alert Threshold value attribute No.  (Note)  0 Rx bad fragments 1  1 Rx lost fragments 2  2 Rx lost starts 3  3 Rx lost ends 4  4..207 Reserved   NOTE – This number associates the TCA with the specified threshold value attribute of the  threshold data 1/2 managed entities.  
```
