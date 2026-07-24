# Managed Entity

## Identity
- ME ID: 156
- ME Name: MGC performance monitoring history data
- Source Section: 9.9.17
- Source Page: 411

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
- Name: Received messages
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 4
- Name: Received octets
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 5
- Name: Sent messages
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 6
- Name: Sent octets
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 7
- Name: Protocol errors
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 8
- Name: Transport losses
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 9
- Name: Last detected event
- Size: 1 byte
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 10
- Name: Last detected reset time
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
9.9.17 MGC performance monitoring history data  The MGC monitoring data ME provides run-time statistics for an active MGC association. Instances  of this ME are created and deleted by the OLT.  For a complete discussion of generic PM architecture, refer to clause I.4.  Relationships  An instance of this ME is associated with an instance of the MGC config data or MGC config  portal ME. 

---- page break ---- Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. Through an  identical ID, this ME is implicitly linked to an instance of the associated MGC  config data or to the MGC config portal ME. If a non -OMCI configuration  method is used for VoIP, there can be only one live ME instance, associated  with the MGC config portal, and with ME ID 0. (R, set-by-create) (mandatory)  (2 bytes)  Interval end time : This attribute identifies the most recently finished 15  min interval. (R)  (mandatory) (1 byte)  Threshold data 1/2 ID: This attribute points to an instance of the threshold data  1 ME that  contains PM threshold values. Since no threshold value attribute number  exceeds 7, a threshold data 2 ME is optional. (R, W, set-by-create) (mandatory)  (2 bytes)  Received messages: This attribute counts the number of received Megaco messages on this  association, as defined by [ITU-T H.341]. (R) (mandatory) (4 bytes)  Received octets: This attribute counts the total number of octets received on this association,  as defined by [ITU-T H.341]. (R) (mandatory) (4 bytes)  Sent messages: This attribute counts the total number of Megaco  messages sent over this  association, as defined by [ITU-T H.341]. (R) (mandatory) (4 bytes)  Sent octets: This attribute counts the total number of octets sent over this association, as  defined by [ITU-T H.341]. (R) (mandatory) (4 bytes)  Protocol errors: This attribute counts the total number of errors detected on this association,  as defined by [ITU-T H.341]. This includes:   • syntax errors detected in a given received message;   • outgoing transactions that failed for protocol reasons.  (R) (mandatory) (4 bytes)  Transport losses : This attribute counts the total number of transport losses ( e.g., socket  problems) detected on this association. A link loss is defined as loss of  communication with the remote entity due to hardware/transient problems, or  problems in related software. (R) (mandatory) (4 bytes)  Last detected event : This attribute reports the last event detected on this association. This  includes events such as the link failing or being set up. Under normal  circumstances, a get action on this attribute would return 0 to indicate no  abnormal activity. This field is an enumeration as follows.  0 No event – No event has yet been detected during this PM interval.  1 Link up – The transport link underpinning the association came up.  2 Link down – The transport link underpinning the association went  down.  3 Persistent error – A persistent error was detected on the link (such as  the socket/TCP connection to the remote node could not be set up).  4 Local shutdown – The association was brought down intentionally  by the local application.  5 Failover down – The association was brought down as part of  failover processing.  255 Other event – The latest event does not match any in the list.  (R) (mandatory) (1 byte) 

---- page break ---- Last detected event time : This attribute reports the time in seconds since the last event on  this association was detected, as defined by [ITU -T H.341]. (R) (mandatory)  (4 bytes)  Last detected reset time: This attribute reports the time in seconds since these statistics were  last reset, as defined by [ITU -T H.341]. Under normal circumstances, a get  action on this attribute would return 900  s to indicate a completed 15  min  interval. (R) (mandatory) (4 bytes)  Actions  Create, delete, get, set  Get current data (optional)  Notifications  Threshold crossing alert  Alarm  number  Threshold crossing  alert Threshold value attribute No. (Note)  0 MGCP protocol errors 1  1 MGCP transport  losses  2  NOTE – This number associates the TCA with the specified threshold value attribute of the  threshold data 1 managed entity.  
```
