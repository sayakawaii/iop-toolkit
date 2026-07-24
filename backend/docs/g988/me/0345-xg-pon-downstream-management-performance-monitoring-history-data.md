# Managed Entity

## Identity
- ME ID: 345
- ME Name: XG-PON downstream management performance monitoring history data
- Source Section: 9.2.16
- Source Page: 124

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
- Name: Downstream PLOAM messages count
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 4
- Name: Profile messages received
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 5
- Name: Ranging_time messages received
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 6
- Name: Deactivate_ONU-ID messages received
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 7
- Name: Disable_serial_number messages received
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 8
- Name: Request_registration messages received
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 9
- Name: Assign_alloc-ID messages received
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 10
- Name: Key_control messages received
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 11
- Name: Sleep_allow messages received
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 12
- Name: Baseline OMCI messages received count
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 13
- Name: Extended OMCI messages received count
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 14
- Name: Assign_ONU-ID messages received
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 15
- Name: OMCI MIC error count
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: optional

## Extraction Status
- Status: auto_extracted
- Attributes found: 15
- Review needed: false

## Raw Source

```
9.2.16 XG-PON downstream management performance monitoring history data  This ME collects PM data associated with the XG-PON TC layer. It collects counters associated with  downstream PLOAM and OMCI messages.  For a complete discussion of generic PM architecture, refer to clause I.4.  Relationships  An instance of this ME is associated with an ANI-G.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. Through an  identical ID, this ME is implicitly linked to an instance of the ANI-G. (R, set- by-create) (mandatory) (2 bytes)  Interval end time : This attribute identifies the most recently finished 15  min interval. (R)  (mandatory) (1 byte)  Threshold data 1/2 ID: This attribute points to an instance of the threshold data 1 ME that  contains PM threshold values. Since no threshold value attribute number  exceeds 7, a threshold data 2 ME is optional. (R, W, set-by-create) (mandatory)  (2 bytes)  PLOAM message integrity check ( MIC) error count : This attribute counts MIC errors  detected in downstream PLOAM messages, either directed to this ONU or  broadcast to all ONUs. (R) (optional) (4 bytes)  Downstream PLOAM messages count : This attribute counts PLOAM messages received,  either directed to this ONU or broadcast to all ONUs. (R) (optional) (4 bytes)  Profile messages received : This attribute counts the number of profile messages received,  either directed to this ONU or broadcast to all ONUs. In [ITU-T G.9807.1],  this attribute is used for received burst_profile message count. (R) (optional)  (4 bytes)  Ranging_time messages received : This attribute counts the number of ranging_time  messages received, either directed to this ONU or broadcast to all ONUs. (R)  (mandatory) (4 bytes)  Deactivate_ONU-ID messages received : This attribute counts the number of  deactivate_ONU-ID messages received, either directed to this ONU or 

---- page break ---- broadcast to all ONUs. Deactivate_ONU -ID messages do not reset this  counter. (R) (optional) (4 bytes)  Disable_serial_number messages received : This attribute counts the number of  disable_serial_number messages received, whose serial number specified this  ONU. (R) (optional) (4 bytes)  Request_registration messages received : This attribute counts the number of  request_registration messages received. (R) (optional) (4 bytes)  Assign_alloc-ID messages received : This attribute counts the number of assign_alloc -ID  messages received. (R) (optional) (4 bytes)  Key_control messages received: This attribute counts the number of key_control messages  received, either directed to this ONU or broadcast to all ONUs. (R) (optional)  (4 bytes)  Sleep_allow messages received: This attribute counts the number of sleep_allow messages  received, either directed to this ONU or broadcast to all ONUs. (R) (optional)  (4 bytes)  Baseline OMCI messages received count : This attribute counts the number of OMCI  messages received in the baseline message format. (R) (optional) (4 bytes)  Extended OMCI messages received count : This attribute counts the number of OMCI  messages received in the extended message format. (R) (optional) (4 bytes)  Assign_ONU-ID messages received : This attribute counts the number of assign_ONU -ID  messages received since the last re-boot. (R) (optional) (4 bytes)  OMCI MIC error count : This attribute counts MIC errors detected in OMCI messages  directed to this ONU. (R) (optional) (4 bytes)  Actions  Create, delete, get, set  Get current data (optional)  Notifications  Threshold crossing alert  Alarm  number Threshold crossing alert Threshold value attribute No. (Note)  1 PLOAM MIC error count 1  2 OMCI MIC error count 2  NOTE – This number associates the TCA with the specified threshold value attribute of the  threshold data 1 managed entity.  
```
