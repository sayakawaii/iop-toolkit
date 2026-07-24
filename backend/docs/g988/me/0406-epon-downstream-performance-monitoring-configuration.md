# Managed Entity

## Identity
- ME ID: 406
- ME Name: EPON downstream performance monitoring configuration
- Source Section: 9.2.20
- Source Page: 133

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Get, Set
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: Errored symbol period  window
- Size: 8 bytes
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 2
- Name: Errored symbol period threshold
- Size: 8 bytes
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 3
- Name: Errored frame  window
- Size: 2 bytes
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 4
- Name: Errored frame threshold
- Size: 4 bytes
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 5
- Name: Errored frame period window
- Size: 4 bytes
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 6
- Name: Errored frame period threshold
- Size: 4 bytes
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 7
- Name: Errored frame seconds summary window
- Size: 2 bytes
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 8
- Name: Errored frame seconds summary threshold
- Size: 2 bytes
- Format: needs_review
- Access: R, W
- Category: optional

## Extraction Status
- Status: auto_extracted
- Attributes found: 8
- Review needed: false

## Raw Source

```
9.2.20  EPON downstream performance monitoring configuration  This ME represents window sizes and threshold values for EPON downstream PM operations which  are defined in [IEEE 802.3] as: errored symbol period, errored frame, errored frame period and  errored frame seconds summary.  The EPON ONU automatically instantiates an instance of this ME for each ANI-E. 

---- page break ---- Relationships  An instance of this ME is associated with an ANI-E.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. Through an  identical ID, this ME is implicitly linked to an instance of the ANI-E. (R, set- by-create) (mandatory) (2 bytes)  Errored symbol period  window: This attribute specifies the number of symbols in the  measurement period, as defined in clause 57.5.3.1 of [IEEE 802.3].   (R, W) (optional) (8 bytes)  Errored symbol period threshold: This attribute specifies the threshold of errored symbols  for generating an event report, as defined in clause 57.5.3.1 of [IEEE 802.3].  (R, W) (optional) (8 bytes)  Errored frame  window: This attribute specifies the duration in units of 100  ms of the  measurement period, as defined in clause 57.5.3.2 of [IEEE 802.3].   (R, W) (optional) (2 bytes)  Errored frame threshold : This attribute specifies the threshold of errored frames for  generating an event report, as defined in clause 57.5.3.2 of [IEEE 802.3].   (R, W) (optional) (4 bytes)  Errored frame period window: This attribute specifies the duration in terms of frames of  the measurement period, as defined in clause 57.5.3.3 of [IEEE 802.3].   (R, W) (optional) (4 bytes)  Errored frame period threshold: This attribute specifies the threshold of errored frames for  generating an event report, as defined in clause 57.5.3.3 of [IEEE 802.3].   (R, W) (optional) (4 bytes)  Errored frame seconds summary window: This attribute specifies the duration in units of  100 ms of the measurement period, as defined in clause 57.5.3.4 of  [IEEE 802.3]. (R, W) (optional) (2 bytes)  Errored frame seconds summary threshold: This attribute specifies the threshold of errored  frame seconds for generating an event report, as defined in clause 57.5.3.4 of  [IEEE 802.3]. (R, W) (optional) (2 bytes)  Actions  Get, set  Notifications  None.  
```
