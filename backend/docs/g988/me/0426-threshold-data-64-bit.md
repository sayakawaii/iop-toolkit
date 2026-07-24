# Managed Entity

## Identity
- ME ID: 426
- ME Name: Threshold data 64 bit
- Source Section: 9.12.17
- Source Page: 441

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Create, Delete, Get, Set
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: Threshold value 1
- Size: 8 bytes
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 2
- Name: Threshold value 2
- Size: 8 bytes
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 3
- Name: Threshold value 3
- Size: 8 bytes
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 4
- Name: Threshold value 4
- Size: 8 bytes
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 5
- Name: Threshold value 5
- Size: 8 bytes
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 6
- Name: Threshold value 6
- Size: 8 bytes
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 7
- Name: Threshold value 7
- Size: 8 bytes
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 8
- Name: Threshold value 8
- Size: 8 bytes
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 9
- Name: Threshold value 9
- Size: 8 bytes
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 10
- Name: Threshold value 10
- Size: 8 bytes
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 11
- Name: Threshold value 11
- Size: 8 bytes
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 12
- Name: Threshold value 12
- Size: 8 bytes
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 13
- Name: Threshold value 13
- Size: 8 bytes
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 14
- Name: Threshold value 14
- Size: 8 bytes
- Format: needs_review
- Access: R, W
- Category: mandatory

## Extraction Status
- Status: auto_extracted
- Attributes found: 14
- Review needed: false

## Raw Source

```
9.12.17 Threshold data 64 bit  An instance of this ME contains threshold values for counters in PM history data MEs.  Instances of this ME are created and deleted by the OLT.  Relationships  An instance of this ME may be related to multiple instances of PM history data type MEs.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. (R, set-by- create) (mandatory) (2 bytes)  The following attributes specify threshold values for thresholded counters in associated PM history  data MEs. The definition of each PM history ME includes a table that links each thresholded counter  to one of these threshold value attributes. The default values of these attributes are all 1s.  Threshold value 1: (R, W) (mandatory) (8 bytes)  Threshold value 2: (R, W) (mandatory) (8 bytes)  Threshold value 3: (R, W) (mandatory) (8 bytes)   Threshold value 4: (R, W) (mandatory) (8 bytes)   Threshold value 5: (R, W) (mandatory) (8 bytes)   Threshold value 6: (R, W) (mandatory) (8 bytes)   Threshold value 7: (R, W) (mandatory) (8 bytes)   Threshold value 8: (R, W) (mandatory) (8 bytes)   Threshold value 9: (R, W) (mandatory) (8 bytes)  Threshold value 10: (R, W) (mandatory) (8 bytes)  Threshold value 11: (R, W) (mandatory) (8 bytes)   Threshold value 12: (R, W) (mandatory) (8 bytes)   Threshold value 13: (R, W) (mandatory) (8 bytes)   Threshold value 14: (R, W) (mandatory) (8 bytes)   Actions  Create, delete, get, set  Notifications  None  
```
