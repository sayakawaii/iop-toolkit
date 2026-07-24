# Managed Entity

## Identity
- ME ID: 18
- ME Name: AAL5 performance monitoring history data
- Source Section: 9.13.6
- Source Page: 459

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
- Name: Sum of invalid CS field errors
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 4
- Name: CRC violations
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 5
- Name: Reassembly timer expirations
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 6
- Name: Encap protocol errors
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

## Extraction Status
- Status: auto_extracted
- Attributes found: 6
- Review needed: false

## Raw Source

```
9.13.6 AAL5 performance monitoring history data  This ME collects PM data as a result of  performing segmentation and reassembly (SAR) and  convergence sublayer (CS) level protocol monitoring. Instances of this ME are created and deleted  by the OLT.  For a complete discussion of generic PM architecture, refer to clause I.4.  Relationships  An instance of this ME is associated with an instance of an IW VCC TP that represents AAL5  functions.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. Through an  identical ID, this ME is implicitly linked to an instance of the IW VCC TP. (R,  set-by-create) (mandatory) (2 bytes)  Interval end time : This attribute identifies the most recently finished 15  min interval. (R)  (mandatory) (1 byte)  Threshold data 1/2 ID: This attribute points to an instance of the threshold data  1 ME that  contains PM threshold values. Since no threshold value attribute number  exceeds 7, a threshold data 2 ME is optional. (R, W, set-by-create) (mandatory)  (2 bytes)  Sum of invalid CS field errors: This attribute counts the sum of invalid CS field errors. For  AAL type 5, this attribute is a single count of the number of CS PDUs  discarded due to one of the following error conditions: invalid common part  indicator (CPI), oversized received SDU, or length violation. (R) (mandatory)  (4 bytes)  CRC violations: This attribute counts CRC violations detected on incoming SAR PDUs. (R)  (mandatory) (4 bytes)  Reassembly timer expirations : This attribute counts reassembly timer expirations. (R)  (mandatory if reassembly timer is implemented) (4 bytes) 

---- page break ---- Buffer overflows: This attribute counts the number of times where there was not enough  buffer space for a reassembled packet. (R) (mandatory) (4 bytes)  Encap protocol errors : This attribute counts the number of times that [IETF RFC 2684]  encapsulation protocol detected a bad header. (R) (mandatory) (4 bytes)  Actions  Create, delete, get, set  Get current data (optional)  Notifications  Threshold crossing alert  Alarm  number Threshold crossing alert Threshold value attribute No. (Note)  0 Invalid fields 1  1 CRC violation 2  2 Reassembly timer  expirations  3  3 Buffer overflows 4  4 Encap protocol errors 5  NOTE – This number associates the TCA with the specified threshold value attribute of the  threshold data 1 managed entity.  
```
