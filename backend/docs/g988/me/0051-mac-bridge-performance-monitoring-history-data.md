# Managed Entity

## Identity
- ME ID: 51
- ME Name: MAC bridge performance monitoring history data
- Source Section: 9.3.3
- Source Page: 142

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
- Name: Bridge learning entry discard count
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

## Extraction Status
- Status: auto_extracted
- Attributes found: 3
- Review needed: false

## Raw Source

```
Actions  Get  Notifications  None.  9.3.3 MAC bridge performance monitoring history data  This ME collects PM data associated with a MAC bridge. Instances of this ME are created and deleted  by the OLT.  For a complete discussion of generic PM architecture, refer to clause I.4.  Relationships  This ME is associated with an instance of a MAC bridge service profile.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. Through an  identical ID, this ME is implicitly linked to an instance of the MAC bridge  service profile. (R, set-by-create) (mandatory) (2 bytes)  Interval end time : This attribute identifies the most recently finished 15  min interval. (R)  (mandatory) (1 byte)  Threshold data 1/2 ID: This attribute points to an instance of the threshold data 1 ME that  contains PM threshold values. Since no threshold value attribute number  exceeds 7, a threshold data 2 ME is optional. Since no threshold value attribute  number exceeds 7, a threshold data 2 ME is optional. (R,  W, set-by-create)  (mandatory) (2 bytes)  Bridge learning entry discard count: This attribute counts forwarding database entries that  have been or would have been learned , but were discarded or replaced due to  a lack of space in the database table. When used with the MAC learning depth  attribute of the MAC bridge service profile, the bridge learning entry discard  count may be particularly useful in detecting MAC spoofing attempts. (R)  (mandatory) (4 bytes)  Actions  Create, delete, get, set  Get current data (optional)  Notifications  Threshold crossing alert  Alarm  number Threshold crossing alert Threshold value attribute No. (Note)  0 Bridge learning entry discard 1  NOTE – This number associates the TCA with the specified threshold value attribute of the  threshold data 1 managed entity.  
```
