# Managed Entity

## Identity
- ME ID: 445
- ME Name: TWDM channel XGEM performance monitoring history data
- Source Section: 9.16.4
- Source Page: 491

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
- Name: Threshold data 64 bit ID
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 3
- Name: Total transmitted XGEM frames
- Size: 8 byte
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 4
- Name: Transmitted XGEM frames with LF bit not set
- Size: 8 byte
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 5
- Name: Total received XGEM frames
- Size: 8 byte
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 6
- Name: Received XGEM frames with XGEM header HEC errors
- Size: 8 byte
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 7
- Name: FS words lost to XGEM header HEC errors
- Size: 8 byte
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 8
- Name: XGEM encryption key errors
- Size: 8 byte
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 9
- Name: Total transmitted bytes in non -idle XGEM frames
- Size: 8 byte
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 10
- Name: Total received bytes in non-idle XGEM frames
- Size: 8 byte
- Format: needs_review
- Access: R
- Category: mandatory

## Extraction Status
- Status: auto_extracted
- Attributes found: 10
- Review needed: false

## Raw Source

```
9.16.4 TWDM channel XGEM performance monitoring history data  This ME collects certain XGEM-related PM data associated with the slot/circuit pack, hosting one or  more ANI-G MEs, for a specific TWDM channel. Instances of this ME are created and deleted by  the OLT.  For a complete discussion of generic PM architecture, refer to clause I.4.  Relationships  An instance of this ME is associated with an instance of TWDM channel ME.   Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. Through an  identical ID, this ME is implicitly linked to an instance of the TWDM channel  ME. (R, set-by-create) (mandatory) (2 bytes)  Interval end time : This attribute identifies the most recently finished 15  min interval. (R)  (mandatory) (1 byte)  Threshold data 64 bit ID: This attribute points to an instance of the threshold data 64 bit ME  that contains PM threshold values. (R, W, set-by-create) (mandatory) (2 bytes)  Total transmitted XGEM frames : The counter aggregated across all XGEM ports of the  given ONU. (R) (mandatory) (8 byte)  Transmitted XGEM frames with LF bit not set: The counter aggregated across all XGEM  ports of the given ONU identifies the number of fragmentation operations. (R)  (mandatory) (8 byte)  Total received XGEM frames: The counter aggregated across all XGEM ports of the given  ONU. (R) (mandatory) (8 byte)  Received XGEM frames with XGEM header HEC errors: The counter aggregated across  all XGEM ports of the given ONU identifies the number of loss XGEM frame  delineation events. (R) (mandatory) (8 byte)  FS words lost to XGEM header HEC errors : The counter of the FS fra me words lost due  to XGEM frame header errors that cause loss of XGEM frame delineation. (R)  (mandatory) (8 byte)  XGEM encryption key errors: The counter aggregated across all XGEM ports of the given  ONU identifies the number of received XGEM frames that have to be  discarded because of unknown or invalid encryption key . The number is  included into the Total received XGEM frame count above. (R) (mandatory)  (8 byte)  Total transmitted bytes in non -idle XGEM frames : The counter aggregated across all  XGEM ports of the given. (R) (mandatory) (8 byte)  Total received bytes in non-idle XGEM frames: The counter aggregated across all XGEM  ports of the given ONU. (R) (mandatory) (8 byte)  Actions  Create, delete, get, set  Get current data (optional) 

---- page break ---- Notifications  Threshold crossing alert  Alarm number Threshold crossing alert Threshold value attribute No.  (Note)  0 Received XGEM header HEC errors  1  1 FS words lost to XGEM header HEC errors  2  2 XGEM encryption key errors 3  NOTE – This number associates the TCA with the specified threshold value attribute of the threshold data  64 bit managed entity.  
```
