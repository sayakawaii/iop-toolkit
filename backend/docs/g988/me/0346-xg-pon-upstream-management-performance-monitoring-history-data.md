# Managed Entity

## Identity
- ME ID: 346
- ME Name: XG-PON upstream management performance monitoring history data
- Source Section: 9.2.17
- Source Page: 125

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
- Name: Upstream PLOAM message count
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 4
- Name: Serial_number_ONU message count
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 5
- Name: Registration message count
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 6
- Name: Key_report message count
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 7
- Name: Acknowledge message count
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 8
- Name: Sleep_request message count
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: optional

## Extraction Status
- Status: auto_extracted
- Attributes found: 8
- Review needed: false

## Raw Source

```
9.2.17 XG-PON upstream management performance monitoring history data  This ME collects PM data associated with the XG -PON TC layer. It counts upstream PLOAM  messages transmitted by the ONU.   For a complete discussion of generic PM architecture, refer to clause I.4.  Relationships  An instance of this ME is associated with an ANI-G. 

---- page break ---- Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. Through an  identical ID, this ME is implicitly linked to an instance of the ANI-G. (R, set- by-create) (mandatory) (2 bytes)  Interval end time : This attribute identifies the most recently finished 15  min interval. (R)  (mandatory) (1 byte)  Threshold data 1/2 ID: No thresholds are defined for this ME. For uniformity with other PM,  the attribute is retained and shown as mandatory, but it should be set to a null  pointer. (R, W, set-by-create) (mandatory) (2 bytes)  Upstream PLOAM message count : This attribute counts PLOAM messages transmitted  upstream, excluding acknowledge messages. (R) (optional) (4 bytes)  Serial_number_ONU message count: This attribute counts Serial_number_ONU PLOAM  messages transmitted. (R) (optional) (4 bytes)  Registration message count : This attribute counts Registration PLOAM messages  transmitted. (R) (optional) (4 bytes)  Key_report message count: This attribute counts key_report PLOAM messages transmitted.  (R) (optional) (4 bytes)  Acknowledge message count : This attribute counts acknowledge PLOAM messages  transmitted. It includes all forms of acknowledgement  (AK), including those  transmitted in response to a PLOAM grant when the ONU has nothing to send.  (R) (optional) (4 bytes)  Sleep_request message count : This attribute counts sleep_request PLOAM messages  transmitted. (R) (optional) (4 bytes)  Actions  Create, delete, get, set  Get current data (optional)  Notifications  None.  
```
