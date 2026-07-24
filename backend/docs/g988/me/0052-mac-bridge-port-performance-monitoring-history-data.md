# Managed Entity

## Identity
- ME ID: 52
- ME Name: MAC bridge port performance monitoring history data
- Source Section: 9.3.9
- Source Page: 149

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
- Name: Forwarded frame counter
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 4
- Name: Delay exceeded discard counter
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 5
- Name: Received frame counter
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 6
- Name: Received and discarded counter
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

## Extraction Status
- Status: auto_extracted
- Attributes found: 6
- Review needed: false

## Raw Source

```
9.3.9 MAC bridge port performance monitoring history data  This ME collects PM data associated with a MAC bridge port. Instances of this ME are created and  deleted by the OLT.  For a complete discussion of generic PM architecture, refer to clause I.4.  Relationships  An instance of this ME is associated with an instance of a MAC bridge port configuration data  ME.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. Through an  identical ID, this ME is implicitly linked to an instance of the MAC bridge port  configuration data ME. (R, set-by-create) (mandatory) (2 bytes)  Interval end time : This attribute identifies the most recently finished 15  min interval. (R)  (mandatory) (1 byte)  Threshold data 1/2 ID: This attribute points to an instance of the threshold data  1 ME that  contains PM threshold values. Since no threshold value attribute number  exceeds 7, a threshold data 2 ME is optional. (R, W, set-by-create) (mandatory)  (2 bytes)  Forwarded frame counter: This attribute counts frames transmitted successfully on this port.  (R) (mandatory) (4 bytes)  Delay exceeded discard counter: This attribute counts frames discarded on this port because  transmission was delayed. (R) (mandatory) (4 bytes)  Maximum transmission unit  (MTU) exceeded discard counter : This attribute counts  frames discarded on this port because the MTU was exceeded. (R) (mandatory)  (4 bytes)  Received frame counter: This attribute counts frames received on this port. (R) (mandatory)  (4 bytes)  Received and discarded counter: This attribute counts frames received on this port that were  discarded due to errors. (R) (mandatory) (4 bytes)  Actions  Create, delete, get, set  Get current data (optional) 

---- page break ---- Notifications    Threshold crossing alert  Alarm number Threshold crossing alert Threshold value attribute No. (Note)  1 Delay exceeded discard 1  2 MTU exceeded discard 2  4 Received and discarded 3  NOTE – This number associates the TCA with the specified threshold value attribute of the threshold  data 1 managed entity.  
```
