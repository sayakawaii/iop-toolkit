# Managed Entity

## Identity
- ME ID: 341
- ME Name: GEM port network CTP performance monitoring history data
- Source Section: 9.2.13
- Source Page: 120

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
- Name: Received GEM frames
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 4
- Name: Received payload bytes
- Size: 8 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 5
- Name: Transmitted payload bytes
- Size: 8 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 6
- Name: Encryption key errors
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: optional

## Extraction Status
- Status: auto_extracted
- Attributes found: 6
- Review needed: false

## Raw Source

```
9.2.13 GEM port network CTP performance monitoring history data  This ME collects GEM frame PM data associated with a GEM port network CTP. Instances of this  ME are created and deleted by the OLT.  NOTE 1 – One might expect to find some form of impaired or discarded frame count associated with a GEM  port. However, the only impairment that might be detected at the GEM frame level would be a corrupted GEM  frame header. In this case, no part of the header c ould be considered reliable including the port ID. For this  reason, there is no impaired or discarded frame count in this ME.  NOTE 2 – This ME replaces the GEM port performance history data ME and is preferred for new  implementations.  For a complete discussion of generic PM architecture, refer to clause I.4.  Relationships  An instance of this ME is associated with an instance of the GEM port network CTP ME.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. Through an  identical ID, this ME is implicitly linked to an instance of the GEM port  network CTP. (R, set-by-create) (mandatory) (2 bytes)  Interval end time : This attribute identifies the most recently finished 15  min interval. (R)  (mandatory) (1 byte)  Threshold data 1/2 ID: This attribute points to an instance of the threshold data  1 ME that  contains PM threshold values. Since no threshold value attribute number  exceeds 7, a threshold data 2 ME is optional. (R, W, set-by-create) (mandatory)  (2 bytes) 

---- page break ---- Transmitted GEM frames: This attribute counts GEM frames transmitted on the monitored  GEM port. (R) (mandatory) (4 bytes)  Received GEM frames : This attribute counts GEM frames received correctly on the  monitored GEM port . A correctly received GEM frame is one that does not  contain uncorrectable errors and has a valid header error check  (HEC). (R)  (mandatory) (4 bytes)  Received payload bytes: This attribute counts user payload bytes received on the monitored  GEM port. (R) (mandatory) (8 bytes)  Transmitted payload bytes : This attribute counts user payload bytes transmitted on the  monitored GEM port. (R) (mandatory) (8 bytes)  Encryption key errors : This attribute is defined in ITU -T G.987 systems only. It counts  GEM frames with erroneous encryption key indexes. If the GEM port is not  encrypted, this attribute counts any frame with a key index not equal to 0. If  the GEM port is encrypted, this attri bute counts any frame whose key index  specifies a key that is not known to the ONU. (R) (optional) (4 bytes)  NOTE 3 – GEM PM ignores idle GEM frames.   NOTE 4 – GEM PM counts each non -idle GEM frame, whether it contains an entire user frame or only a  fragment of a user frame.  Actions  Create, delete, get, set  Get current data (optional)  Notifications  Threshold crossing alert  Alarm  number Threshold crossing alert Threshold value attribute No. (Note)  1 Encryption key errors 1  NOTE – This number associates the TCA with the specified threshold value attribute of the  threshold data 1 managed entity.  
```
