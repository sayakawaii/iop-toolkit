# Managed Entity

## Identity
- ME ID: 23
- ME Name: CES physical interface performance monitoring history data
- Source Section: 9.8.4
- Source Page: 362

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
- Name: Severely errored seconds
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 4
- Name: Burst errored seconds
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 5
- Name: Unavailable seconds
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 6
- Name: Controlled slip seconds
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: mandatory

## Extraction Status
- Status: auto_extracted
- Attributes found: 6
- Review needed: false

## Raw Source

```
9.8.4 CES physical interface performance monitoring history data  This ME collects statistics for a CES physical interface. Interfaces include DS1, E1, J1, J2 and  possibly others. The performance management requirements of particular interfaces are described in  the corresponding ITU-T or other standards document, e.g., [ITU-T G.784] or [b-ATIS-0300231].  Instances of this ME are created and deleted by the OLT.  For a complete discussion of generic PM architecture, refer to clause I.4.  Relationships  An instance of this ME is associated with one instance of the PPTP CES UNI.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. Through an  identical ID, this ME is implicitly linked to an instance of the PPTP CES UNI.  (R, set-by-create) (mandatory) (2 bytes)  Interval end time : This attribute identifies the most recently finished 15  min interval. (R)  (mandatory) (1 byte)  Threshold data 1/2 ID: This attribute points to an instance of the threshold data  1 ME that  contains PM threshold values. Since no mandatory threshold value attribute  number exceeds 7, a threshold data 2 ME is optional. (R,  W, set-by-create)  (mandatory) (2 bytes) 

---- page break ---- Errored seconds: When a detailed distinction needs to be made, this attribute corresponds to  near-end line errored seconds, ES -L, also known as LES. (R) (mandatory)  (2 bytes)  Severely errored seconds : When a  detailed distinction needs to be made, this attribute  corresponds to near-end line severely errored seconds, SES-L. (R) (mandatory)  (2 bytes)  Burst errored seconds: A burst errored second (BES) is any second that is not an unavailable  second (UAS), that contains between 2 and 319 error events but no LOS, AIS  or OOF condition. This attribute is also known as ESB -P. (R) (optional)  (2 bytes)  Unavailable seconds: When a detailed distinction needs to be made, this attribute corresponds  to near-end path unavailable seconds, UAS-P. (R) (mandatory) (2 bytes)  Controlled slip seconds : When a  detailed distinction needs to be made, this attribute  corresponds to near-end path controlled slip seconds CSS-P. (R) (mandatory)  (2 bytes)  Each of the following attributes is (R) (optional) (2 bytes)    Attribute name Common acronym  Loss of signal seconds LOSS-L  AIS seconds AISS-P  Errored seconds, path ES-P  Errored seconds, type A ESA-P  Severely errored seconds, path SES-P  Severely errored frame and AIS seconds SAS-P aka SEFS  Code violations, line CV-L aka LCV  Code violations, path CV-P aka PCV  Errored blocks (see [ITU-T G.826]) EB  Actions  Create, delete, get, set  Get current data (optional)  Notifications  Threshold crossing alert  Alarm  number Threshold crossing alert Threshold value attribute No. (Note)  0 ES 1  1 SES 2  2 BES 3  3 UAS 4  4 CSS 5  5 LOSS-L 6  6 AISS-P 7  7 ES-P 8 

---- page break ---- Threshold crossing alert  Alarm  number Threshold crossing alert Threshold value attribute No. (Note)  8 ESA-P 9  9 SES-P 10  10 SAS-P 11  11 CV-L 12  12 CV-P 13  13 EB 14  NOTE – This number associates the TCA with the specified threshold value attribute of the  threshold data 1/2 managed entities.  
```
