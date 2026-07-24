# Managed Entity

## Identity
- ME ID: 151
- ME Name: SIP agent performance monitoring history data
- Source Section: 9.9.14
- Source Page: 407

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
- Name: Transactions
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 4
- Name: Rx invite reqs
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 5
- Name: Rx invite retrans
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 6
- Name: Rx noninvite reqs
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 7
- Name: Rx noninvite retrans
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 8
- Name: Rx response
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 9
- Name: Rx response retransmissions
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 10
- Name: Tx invite reqs
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 11
- Name: Tx invite retrans
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 12
- Name: Tx noninvite reqs
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 13
- Name: Tx noninvite retrans
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 14
- Name: Tx response
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 15
- Name: Tx response retransmissions
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
9.9.14 SIP agent performance monitoring history data  This ME collects PM data for the associated VoIP SIP agent. Instances of this ME are created and  deleted by the OLT.  For a complete discussion of generic PM architecture, refer to clause I.4.  Relationships  An instance of this ME is associated with a SIP agent config data or SIP config portal object.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. Through an  identical ID, this ME is implicitly linked to an instance of the corresponding 

---- page break ---- SIP agent config data or to the SIP config portal. If a non-OMCI configuration  method is used for VoIP, there can be only one live ME instance, associated  with the SIP config portal, and with ME ID 0. (R, set-by-create) (mandatory)  (2 bytes)  Interval end time : This attribute identifies the most recently finished 15  min interval. (R)  (mandatory) (1 byte)  Threshold data 1/2 ID: This attribute points to an instance of the threshold data  1 ME that  contains PM threshold values. Since no threshold value attribute number  exceeds 7, a threshold data 2 ME is optional. (R, W, set-by-create) (mandatory)  (2 bytes)  Transactions: This attribute counts the number of new transactions that were initiated. (R)  (optional) (4 bytes)  Rx invite reqs: This attribute counts received invite messages, including retransmissions. (R)  (optional) (4 bytes)  Rx invite retrans : This attribute counts received invite retransmission messages. (R)  (optional) (4 bytes)  Rx noninvite reqs : This attribute counts received non -invite messages, including  retransmissions. (R) (optional) (4 bytes)  Rx noninvite retrans: This attribute counts received non-invite retransmission messages. (R)  (optional) (4 bytes)  Rx response: This attribute counts total responses received. (R) (optional) (4 bytes)  Rx response retransmissions: This attribute counts total response retransmissions received.  (R) (optional) (4 bytes)  Tx invite reqs: This attribute counts transmitted invite messages, including retransmissions.  (R) (optional) (4 bytes)  Tx invite retrans : This attribute counts transmitted invite retransmission messages. (R)  (optional) (4 bytes)  Tx noninvite reqs : This attribute counts transmitted non -invite messages, including  retransmissions. (R) (optional) (4 bytes)  Tx noninvite retrans: This attribute counts transmitted non-invite retransmission messages.  (R) (optional) (4 bytes)  Tx response: This attribute counts the total responses sent. (R) (optional) (4 bytes)  Tx response retransmissions: This attribute counts total response retransmissions sent. (R)  (optional) (4 bytes)  Actions  Create, delete, get, set  Get current data (optional)  Notifications  Threshold crossing alert  Alarm  number Threshold crossing alert Threshold value attribute No. (Note)  0 SIPAMD Rx invite req 1 

---- page break ---- Threshold crossing alert  Alarm  number Threshold crossing alert Threshold value attribute No. (Note)  1 SIPAMD Rx invite req retransmission 2  2 SIPAMD Rx noninvite req 3  3 SIPAMD Rx noninvite req  retransmission  4  4 SIPAMD Rx response 5  5 SIPAMD Rx response retransmission 6  NOTE – This number associates the TCA with the specified threshold value attribute of the  threshold data 1 managed entity.  
```
