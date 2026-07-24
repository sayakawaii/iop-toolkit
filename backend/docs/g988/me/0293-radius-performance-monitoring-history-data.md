# Managed Entity

## Identity
- ME ID: 293
- ME Name: Radius performance monitoring history data
- Source Section: 9.3.17
- Source Page: 177

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Create, Delete, Get, Set
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: Threshold data 1/2 ID
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 2
- Name: Access-request packets transmitted
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 3
- Name: Access-request retransmission count
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 4
- Name: Access-challenge packets received
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 5
- Name: Access-accept packets received
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 6
- Name: Access-reject packets received
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 7
- Name: Invalid radius packets received
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

## Extraction Status
- Status: auto_extracted
- Attributes found: 7
- Review needed: false

## Raw Source

```
9.3.17 Radius performance monitoring history data  This ME collects performance statistics on an ONU 's radius client, particularly as related to its  IEEE 802.1X operation.  Instances of this ME are created and deleted by the OLT.  For a complete discussion of generic PM architecture, refer to clause I.4.  Relationships  An instance of this ME is associated with an ONU.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. Through an  identical ID (namely 0), this ME is implicitly linked to an instance of a dot1X  configuration profile. (R, set-by-create) (mandatory) (2 bytes) 

---- page break ---- Interval end time : This attribute identifies the most recently finished 15  min interval. (R)  (mandatory) (1 byte)  Threshold data 1/2 ID: This attribute points to an instance of the threshold data  1 ME that  contains PM threshold values. Since no threshold value attribute number  exceeds 7, a threshold data 2 ME is optional. (R, W, set-by-create) (mandatory)  (2 bytes)  Access-request packets transmitted: This attribute counts transmitted radius access-request  messages, including retransmissions. (R) (mandatory) (4 bytes)  Access-request retransmission count : This attribute counts radius access -request  retransmissions. (R) (mandatory) (4 bytes)  Access-challenge packets received : This attribute counts received radius access -challenge  messages. (R) (mandatory) (4 bytes)  Access-accept packets received : This attribute counts received radius access -accept  messages. (R) (mandatory) (4 bytes)  Access-reject packets received: This attribute counts received radius access-reject messages.  (R) (mandatory) (4 bytes)  Invalid radius packets received: This attribute counts received invalid radius messages. (R)  (mandatory) (4 bytes)  Actions  Create, delete, get, set  Get current data (optional)  Notifications  Threshold crossing alert  Alarm  number Threshold crossing alert Threshold value attribute No. (Note)  1 Retransmission count 2  5 Invalid radius packets received 6  NOTE – This number associates the TCA with the specified threshold value attribute of the  threshold data 1 managed entity.  
```
