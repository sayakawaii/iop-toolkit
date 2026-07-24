# Managed Entity

## Identity
- ME ID: 163
- ME Name: MoCA Ethernet performance monitoring history data
- Source Section: 9.10.2
- Source Page: 421

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
- Name: Incoming unicast packets
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 4
- Name: Incoming discarded packets
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 5
- Name: Incoming errored packets
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 6
- Name: Incoming unknown packets
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 7
- Name: Incoming multicast packets
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 8
- Name: Incoming broadcast packets
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 9
- Name: Incoming octets
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 10
- Name: Outgoing discarded packets
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 11
- Name: Outgoing errored packets
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 12
- Name: Outgoing unknown packets
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 13
- Name: Outgoing multicast packets
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 14
- Name: Outgoing broadcast packets
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 15
- Name: Outgoing octets
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
9.10.2 MoCA Ethernet performance monitoring history data  This ME collects PM data for a n MoCA Ethernet interface. Instances of this ME are created and  deleted by the OLT.  For a complete discussion of generic PM architecture, refer to clause I.4.  Relationships  An instance of this ME is associated with an instance of the PPTP MoCA UNI ME.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. Through an  identical ID, this ME is implicitly linked to an instance of the PPTP MoCA  UNI. (R, set-by-create) (mandatory) (2 bytes)  Interval end time : This attribute identifies the most recently finished 15  min interval. (R)  (mandatory) (1 byte)  Threshold data 1/2 ID: This attribute points to an instance of the threshold data 1 and 2 MEs  that contains PM threshold values. (R, W, set-by-create) (mandatory) (2 bytes)  Incoming PM refers to upstream traffic received on the UNI; outgoing PM refers to downstream  traffic transmitted on the UNI.  Incoming unicast packets: (R) (optional) (4 bytes)  Incoming discarded packets: (R) (optional) (4 bytes)  Incoming errored packets: (R) (optional) (4 bytes)  Incoming unknown packets: (R) (optional) (4 bytes)  Incoming multicast packets: (R) (optional) (4 bytes)  Incoming broadcast packets: (R) (optional) (4 bytes)  Incoming octets: (R) (optional) (4 bytes) 

---- page break ---- Outgoing unicast packets: (R) (optional) (4 bytes)  Outgoing discarded packets: (R) (optional) (4 bytes)  Outgoing errored packets: (R) (optional) (4 bytes)  Outgoing unknown packets: (R) (optional) (4 bytes)  Outgoing multicast packets: (R) (optional) (4 bytes)  Outgoing broadcast packets: (R) (optional) (4 bytes)  Outgoing octets: (R) (optional) (4 bytes)  Actions  Create, delete, get, set  Notifications  Threshold crossing alert  Alarm  number Threshold crossing alert Threshold value attribute No. (Note)  0 Incoming unicast packets 1  1 Incoming discarded packets 2  2 Incoming errored packets 3  3 Incoming unknown packets 4  4 Incoming multicast packets 5  5 Incoming broadcast packets 6  6 Incoming octets  7  7 Outgoing unicast packets 8  8 Outgoing discarded packets 9  9 Outgoing errored packets 10  10 Outgoing unknown packets 11  11 Outgoing multicast packets 12  12 Outgoing broadcast packets 13  13 Outgoing octets  14  NOTE – This number associates the TCA with the specified threshold value attribute of the  threshold data 1/2 managed entities.  
```
