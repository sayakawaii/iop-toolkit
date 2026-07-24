# Managed Entity

## Identity
- ME ID: 292
- ME Name: Dot1X performance monitoring history data
- Source Section: 9.3.16
- Source Page: 176

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
- Name: EAPOL frames received
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 4
- Name: EAPOL frames transmitted
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 5
- Name: EAPOL start frames received
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 6
- Name: EAPOL logoff frames received
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 7
- Name: EAP resp/id frames received
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 8
- Name: EAP response frames received
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 9
- Name: EAP initial request frames transmitted
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 10
- Name: EAP request frames transmitted
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 11
- Name: EAP length error frames received
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 12
- Name: EAP success frames generated autonomously
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 13
- Name: EAP failure frames generated autonomously
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

## Extraction Status
- Status: auto_extracted
- Attributes found: 13
- Review needed: false

## Raw Source

```
9.3.16 Dot1X performance monitoring history data  This ME collects performance statistics on an ONU 's IEEE 802.1X CPE authentication operation.  Instances of this ME are created and deleted by the OLT.   For a complete discussion of generic PM architecture, refer to clause I.4.  Relationships  An instance of this ME may be associated with each UNI that can perform IEEE  802.1X  authentication of CPE.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. Through an  identical ID, this ME is implicitly linked to an instance of a PPTP. (R,  set-by-create) (mandatory) (2 bytes)  Interval end time : This attribute identifies the most recently finished 15  min interval. (R)  (mandatory) (1 byte)  Threshold data 1/2 ID: This attribute points to an instance of the threshold data 1 and 2 MEs  that contains PM threshold values. (R, W, set-by-create) (mandatory) (2 bytes)  EAPOL frames received: This attribute counts received valid EAPOL frames of any type.  (R) (mandatory) (4 bytes)  EAPOL frames transmitted: This attribute counts transmitted EAPOL frames of any type.  (R) (mandatory) (4 bytes)  EAPOL start frames received : This attribute counts received EAPOL start frames. (R)  (mandatory) (4 bytes)  EAPOL logoff frames received : This attribute counts received EAPOL logoff frames. (R)  (mandatory) (4 bytes) 

---- page break ---- Invalid EAPOL frames received : This attribute counts received EAPOL frames in which  the frame type was not recognized. (R) (mandatory) (4 bytes)  EAP resp/id frames received : This attribute counts received EAP response frames  containing an identifier type field. (R) (mandatory) (4 bytes)  EAP response frames received: This attribute counts received EAP response frames, other  than resp/id frames. (R) (mandatory) (4 bytes)  EAP initial request frames transmitted: This attribute counts transmitted request frames  containing an identifier type field. In [IEEE 802.1X], this is also called ReqId.  (R) (mandatory) (4 bytes)  EAP request frames transmitted : This attribute counts transmitted request frames, other  than request/id frames. (R) (mandatory) (4 bytes)  EAP length error frames received : This attribute counts received EAPOL frames whose  packet body length field was invalid. (R) (mandatory) (4 bytes)  EAP success frames generated autonomously: This attribute counts EAPOL success frames  generated according to the local fallback policy because no radius server was  available. (R) (mandatory) (4 bytes)  EAP failure frames generated autonomously: This attribute counts EAPOL failure frames  generated according to the local fallback policy because no radius server was  available. (R) (mandatory) (4 bytes)  Actions  Create, delete, get, set  Get current data (optional)  Notifications  Threshold crossing alert  Alarm  number Threshold crossing alert Threshold value attribute No. (Note)  4 Invalid EAPOL frames received 5  9 EAP length error frames received 10  NOTE – This number associates the TCA with the specified threshold value attribute of the  threshold data 1/2 managed entities.  
```
