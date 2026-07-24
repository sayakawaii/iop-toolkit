# Managed Entity

## Identity
- ME ID: 274
- ME Name: Threshold data 2
- Source Section: 9.12.6
- Source Page: 430

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Create, Delete, Get, Set
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: Threshold value 1
- Size: 4 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 2
- Name: Threshold value 2
- Size: 4 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 3
- Name: Threshold value 3
- Size: 4 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 4
- Name: Threshold value 4
- Size: 4 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 5
- Name: Threshold value 5
- Size: 4 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 6
- Name: Threshold value 6
- Size: 4 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 7
- Name: Threshold value 7
- Size: 4 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

## Extraction Status
- Status: auto_extracted
- Attributes found: 7
- Review needed: false

## Raw Source

```
9.12.6 Threshold data 1  Threshold data are partitioned into two MEs for historical reasons. An instance of this ME, together  with an optional instance of the threshold data 2 ME, contains threshold values for counters in PM  history data MEs.  For a complete discussion of generic PM architecture, refer to clause I.4.  Instances of this ME are created and deleted by the OLT.  Relationships  An instance of this ME may be related to multiple instances of PM history data type MEs.  Paired instances of threshold data 1 ME and threshold data 2 ME are implicitly linked together  through a common ME ID.  Attributes  Managed entity  ID: This attribute uniquely identifies each instance of this ME. (R,  set-by-create) (mandatory) (2 bytes)  The following seven attributes specify threshold values for seven thresholded counters in associated  PM history data MEs. The definition of each PM history ME includes a table that links each  thresholded counter to one of these threshold value attributes.  Threshold value 1: (R, W, set-by-create) (mandatory) (4 bytes)  Threshold value 2: (R, W, set-by-create) (mandatory) (4 bytes)  Threshold value 3: (R, W, set-by-create) (mandatory) (4 bytes)  Threshold value 4: (R, W, set-by-create) (mandatory) (4 bytes)  Threshold value 5: (R, W, set-by-create) (mandatory) (4 bytes)  Threshold value 6: (R, W, set-by-create) (mandatory) (4 bytes)  Threshold value 7: (R, W, set-by-create) (mandatory) (4 bytes)  Actions  Create, delete, get, set  Notifications  None. 

---- page break ---- 
```
