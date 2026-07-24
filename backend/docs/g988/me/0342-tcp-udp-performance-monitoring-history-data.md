# Managed Entity

## Identity
- ME ID: 342
- ME Name: TCP/UDP performance monitoring history data
- Source Section: 9.4.4
- Source Page: 224

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
- Name: Socket failed
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 4
- Name: Listen failed
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 5
- Name: Bind failed
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 6
- Name: Accept failed
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 7
- Name: Select failed
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
9.4.4 TCP/UDP performance monitoring history data  This ME collects PM data related to a TCP or UDP port. Instances of this ME are created and deleted  by the OLT.  For a complete discussion of generic PM architecture, refer to clause I.4.  Relationships  An instance of this ME is associated with an instance of the TCP/UDP config data ME.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. Through an  identical ID, this ME is implicitly linked to an instance of the TCP/UDP config  data ME. (R, set-by-create) (mandatory) (2 bytes)  Interval end time : This attribute identifies the most recently finished 15  min interval. (R)  (mandatory) (1 byte)  Threshold data 1/2 ID: This attribute points to an instance of the threshold data  1 ME that  contains PM threshold values. Since no threshold value attribute number  exceeds 7, a threshold data 2 ME is optional. (R, W, set-by-create) (mandatory)  (2 bytes)  Socket failed: This attribute is incremented when an attempt to create a socket associated with  a port fails. (R) (mandatory) (2 bytes)  Listen failed: This attribute is incremented when an attempt by a service to listen for a  request on a port fails. (R) (mandatory) (2 bytes)  Bind failed: This attribute is incremented when an attempt by a service to bind to a port  fails. (R) (mandatory) (2 bytes)  Accept failed: This attribute is incremented when an attempt to accept a connection on a port  fails. (R) (mandatory) (2 bytes)  Select failed: This attribute is incremented when an attempt to perform a select on a group  of ports fails. (R) (mandatory) (2 bytes)  Actions  Create, delete, get, set  Get current data (optional) 

---- page break ---- Notifications  Threshold crossing alert  Alarm  number Threshold crossing alert Threshold value attribute No. (Note)  0 N/A   1 Socket failed 1  2 Listen failed 2  3 Bind failed 3  4 Accept failed 4  5 Select failed 5  NOTE – This number associates the TCA with the specified threshold value attribute of the  threshold data 1 managed entity.  
```
