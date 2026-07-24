# Managed Entity

## Identity
- ME ID: 62
- ME Name: VP performance monitoring history data
- Source Section: 9.13.10
- Source Page: 461

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
- Name: Lost C = 0 cells
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 4
- Name: Misinserted cells
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 5
- Name: Transmitted C = 0 + 1 cells
- Size: 5 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 6
- Name: Transmitted C  = 0 cells
- Size: 5 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 7
- Name: Impaired blocks
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: mandatory

## Extraction Status
- Status: auto_extracted
- Attributes found: 7
- Review needed: false

## Raw Source

```
9.13.10 VP performance monitoring history data  This ME collects PM data associated with a VP network CTP. Instances of this ME are created and  deleted by the OLT.  Relationships  An instance of this ME is associated with an instance of the VP network CTP ME. The  performance of upstream ATM flows is reported.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. Through an  identical ID, this ME is implicitly linked to an instance of the VP network CTP.  (R, set-by-create) (mandatory) (2 bytes)  Interval end time : This attribute identifies the most recently finished 15  min interval. (R)  (mandatory) (1 byte)  Threshold data 1/2 ID: This attribute points to an instance of the threshold data  1 ME that  contains PM threshold values. Since no threshold value attribute number  exceeds 7, a threshold data 2 ME is optional. (R, W, set-by-create) (mandatory)  (2 bytes) 

---- page break ---- Lost C = 0 + 1 cells: This attribute counts all cell loss. It cannot distinguish between cells lost  because of header bit errors, ATM-level header errors, cell policing, or buffer  overflows. It records only loss of information, independent of the priority of  the cell. (R) (mandatory) (2 bytes)  Lost C = 0 cells: This attribute counts loss of high priority cells. It cannot distinguish between  cells lost because of header bit errors, ATM-level header errors, cell policing,  or buffer overflows. It records only loss of high priority cells. (R) (mandatory)  (2 bytes)  Misinserted cells : This attribute counts cells that are misrouted to a monitored VP. (R)  (mandatory) (2 bytes)  Transmitted C = 0 + 1 cells : This attribute counts cells originated by the transmitting end  point (i.e., backward reporting is assumed). (R) (mandatory) (5 bytes)  Transmitted C  = 0 cells : This attribute counts high priority cells originated by the  transmitting end point (i.e., backward reporting is assumed). (R) (mandatory)  (5 bytes)  Impaired blocks: This severely errored cell block counter is incremented whenever one of  the following events takes place: the number of misinserted cells reaches its  threshold; the number of bipolar violations reaches its threshold; or the number  of lost cells reaches its threshold. Threshold values are based on vendor - operator negotiation. (R) (mandatory) (2 bytes)  Actions  Create, delete, get, set  Get current data (optional)  Notifications  Threshold crossing alert  Alarm  number  Threshold crossing  alert Threshold value attribute No. (Note)  0 Lost CLP = 0 + 1 cells  1  1 Lost CLP = 0 cells 2  2 Misinserted cells 3  3 Impaired blocks 4  NOTE – This number associates the TCA with the specified threshold value attribute of the  threshold data 1 managed entity.  
```
