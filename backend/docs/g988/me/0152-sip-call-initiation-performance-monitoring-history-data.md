# Managed Entity

## Identity
- ME ID: 152
- ME Name: SIP call initiation performance monitoring history data
- Source Section: 9.9.15
- Source Page: 409

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
- Name: Failed to connect counter
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 4
- Name: Failed to validate counter
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 5
- Name: Timeout counter
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 6
- Name: Failure received counter
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 7
- Name: Failed to authenticate counter
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
9.9.15 SIP call initiation performance monitoring history data  This ME collects PM data related to call initiations of a VoIP SIP agent. Instances of this ME are  created and deleted by the OLT.  For a complete discussion of generic PM architecture, refer to clause I.4.  Relationships  An instance of this ME is associated with an instance of the SIP agent config data or SIP  config portal ME.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. Through an  identical ID, this ME is implicitly linked to an instance of the SIP agent config  data or the SIP config portal ME. If a non-OMCI configuration method is used  for VoIP, there can be only one live ME instance, associated with the SIP  config portal, and with ME ID 0. (R, set-by-create) (mandatory) (2 bytes)  Interval end time : This attribute identifies the most recently finished 15  min interval. (R)  (mandatory) (1 byte)  Threshold data 1/2 ID: This attribute points to an instance of the threshold data  1 ME that  contains PM threshold values. Since no threshold value attribute number  exceeds 7, a threshold data 2 ME is optional. (R, W, set-by-create) (mandatory)  (2 bytes)  Failed to connect counter: This attribute counts the number of times that the SIP UA failed  to reach/connect its TCP/UDP peer during SIP call initiations. (R) (mandatory)  (4 bytes)  Failed to validate counter: This attribute counts the number of times that the SIP UA failed  to validate its peer during SIP call initiations. (R) (mandatory) (4 bytes)  Timeout counter: This attribute counts the number of times that the SIP UA timed out during  SIP call initiations. (R) (mandatory) (4 bytes)  Failure received counter: This attribute counts the number of times that the SIP UA received  a failure error code during SIP call initiations. (R) (mandatory) (4 bytes)  Failed to authenticate counter : This attribute counts the number of times that the SIP UA  failed to authenticate itself during SIP call initiations. (R) (mandatory)  (4 bytes) 

---- page break ---- Actions  Create, delete, get, set  Get current data (optional)  Notifications  Threshold crossing alert  Alarm  number Threshold crossing alert Threshold value attribute No. (Note)  0 SIP call PM failed connect 1  1 SIP call PM failed to validate 2  2 SIP call PM timeout 3  3 SIP call PM failure error code received 4  4 SIP call PM failed to authenticate 5  NOTE – This number associates the TCA with the specified threshold value attribute of the  threshold data 1 managed entity.  
```
