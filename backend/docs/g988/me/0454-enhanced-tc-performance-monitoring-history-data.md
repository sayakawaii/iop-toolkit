# Managed Entity

## Identity
- ME ID: 454
- ME Name: Enhanced TC performance monitoring history data
- Source Section: 9.2.23
- Source Page: 135

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
- Name: Threshold data 64 bit ID
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 3
- Name: PSBd HEC error count
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 4
- Name: XGTC HEC error count
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 5
- Name: Unknown profile count
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 6
- Name: Transmitted XGEM frames
- Size: 8 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 7
- Name: Fragment XGEM frames
- Size: 8 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 8
- Name: XGEM HEC lost words count
- Size: 8 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 9
- Name: XGEM key errors
- Size: 8 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 10
- Name: XGEM HEC error count
- Size: 8 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 11
- Name: Transmitted bytes in non -idle XGEM frames
- Size: 8 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 12
- Name: Received bytes in non -idle XGEM frames
- Size: 8 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 13
- Name: LODS event count
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 14
- Name: ONU reactivation by LODS events
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: optional

## Extraction Status
- Status: auto_extracted
- Attributes found: 14
- Review needed: false

## Raw Source

```
9.2.23 Enhanced TC performance monitoring history data  This ME collects PM data associated with the XGS-PON and subsequent ITU -T PON systems' TC  layer.   For a complete discussion of generic PM architecture, refer to clause I.4.  Relationships  An instance of this ME is associated with an ANI-G. 

---- page break ---- Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. Through an  identical ID, this ME is implicitly linked to an instance of the ANI-G. (R, set- by-create) (mandatory) (2 bytes)  Interval end time : This attribute identifies the most recently finished 15  min interval. (R)  (mandatory) (1 byte)  Threshold data 64 bit ID: This attribute points to an instance of the threshold data 64 bit ME  that contains PM threshold values. (R, W, set-by-create) (mandatory) (2 bytes)  PSBd HEC error count : This attribute counts HEC errors in any of the fields of the  downstream physical sync block. (R) (optional) (4 bytes)  XGTC HEC error count: This attribute counts HEC errors detected in the XGTC header. In  [ITU-T G.9807.1], this attribute is used for FS HEC error count management.  (R) (optional) (4 bytes)  Unknown profile count: This attribute counts the number of grants received whose specified  profile was not known to the ONU. (R) (optional) (4 bytes)  Transmitted XGEM frames : This attribute counts the number of non -idle XGEM frames  transmitted. If an SDU is fragmented, each fragment is an XGEM frame and  is counted as such. (R) (mandatory) (8 bytes)  Fragment XGEM frames: This attribute counts the number of XGEM frames that represent  fragmented SDUs, as indicated by the LF bit = 0. (R) (optional) (8 bytes)  XGEM HEC lost words count: This attribute counts the number of 4 byte words lost because  of an XGEM frame HEC error. In general, all XGTC payload following the  error is lost, until the next PSBd event. (R) (optional) (8 bytes)  XGEM key errors: This attribute counts the number of downstream XGEM frames received  with an invalid key specification. The key may be invalid for several reasons,  among which are:  a) GEM port provisioned for clear text and key index not equal to 00;  b) no multicast key of the specified key index has been provided via the  OMCI for a multicast GEM port;  c) no unicast key of the specified key index has been successfully negotiated  (see clause 15.5 of [ITU-T G.987.3] or clause C.15.5 of [ITU-T G.9807.1]  for key negotiation state machine);  d) GEM port specified to be encrypted and key index = 00;  e) key index = 11, a reserved value.  (R) (mandatory) (8 bytes)  XGEM HEC error count: This attribute counts the number of instances of an XGEM frame  HEC error. (R) (mandatory) (8 bytes)  Transmitted bytes in non -idle XGEM frames : This attribute counts the number of  transmitted bytes in non-idle XGEM frames. (R) (mandatory) (8 bytes)  Received bytes in non -idle XGEM frames : This attribute counts the number of received  bytes in non-idle XGEM frames. (R) (optional) (8 bytes)  LODS event count : This attribute counts the number of state transitions from O5.1 to O6.  (R) (optional) (4 bytes) 

---- page break ---- LODS event restored count: This attribute counts the number of LODS cleared events. (R)  (optional) (4 bytes)  ONU reactivation by LODS events : This attribute counts the number of LODS events  resulting in ONU reactivation without synchronization being reacquired. (R)  (optional) (4 bytes)  Actions  Create, delete, get, set  Get current data (optional)  Notifications  Threshold crossing alert  Alarm  number Threshold crossing alert Threshold value attribute No. (Note)  1 PSBd HEC error count 1  2 XGTC HEC error count 2  3 Unknown profile count 3  4 XGEM HEC loss count 4  5 XGEM key errors 5  6 XGEM HEC error count 6  NOTE – This number associates the TCA with the specified threshold value attribute of the  threshold data 64 bit managed entity.  
```
