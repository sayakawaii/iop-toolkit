# Managed Entity

## Identity
- ME ID: 116
- ME Name: TC adaptor performance monitoring history data xDSL
- Source Section: 9.7.25
- Source Page: 300

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
- Name: Threshold data1/2 ID
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 3
- Name: Near-end HEC violation count
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 4
- Name: Near-end idle cell bit error count
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 5
- Name: Far-end HEC violation count
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 6
- Name: Far-end idle cell bit error count
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
9.7.25 TC adaptor performance monitoring history data xDSL  This ME collects PM data of an xTU-C to xTU-R ATM data path. Instances of this ME are created  and deleted by the OLT.  For a complete discussion of generic PM architecture, refer to clause I.4.  Relationships  An instance of this ME is associated with an xDSL UNI.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. Through an  identical ID, this ME is implicitly linked to an instance of the PPTP xDSL UNI  part 1. (R) (mandatory) (2 bytes)  Interval end time : This attribute identifies the most recently finished 15  min interval. (R)  (mandatory) (1 byte)  Threshold data1/2 ID: This attribute points to an instance of the threshold data 1 ME that  contains PM threshold values. Since no threshold value attribute number  exceeds 7, a threshold data 2 ME is optional. (R, W, set-by-create) (mandatory)  (2 bytes)  Near-end HEC violation count: This attribute counts near-end HEC anomalies in the ATM  data path. (R) (mandatory) (2 bytes)  Near-end delineated total cell count (CD-P): This attribute counts the total number of cells  passed through the cell delineation and HEC function process operating on the  ATM data path while in the SYNC state. (R) (mandatory) (4 bytes) 

---- page break ---- Near-end user total cell count(CU-P): This attribute counts the total number of cells in the  ATM data path delivered at the V-C interface. (R) (mandatory) (4 bytes)  Near-end idle cell bit error count: This attribute counts cells with bit errors in the ATM data  path idle payload received at the near end. (R) (mandatory) (2 bytes)  Far-end HEC violation count: This attribute counts far-end HEC anomalies in the ATM data  path. (R) (mandatory) (2 bytes)  Far-end delineated total cell count  (CD-PFE): This attribute counts the total number of  cells passed through the cell delineation process and HEC function operating  on the ATM data path while in the SYNC state. (R) (mandatory) (4 bytes)  Far-end user total cell count  (CU-PFE): This attribute counts the total number of cells in  the ATM data path delivered at the T-R interface. (R) (mandatory) (4 bytes)  Far-end idle cell bit error count: This attribute counts cells with bit errors in the ATM data  path idle payload received at the far end. (R) (mandatory) (2 bytes)  Actions  Create, delete, get, set  Get current data (optional)  Notifications  Threshold crossing alert  Alarm number Threshold crossing alert Threshold value attribute No. (Note)  0 Near-end HEC violation 1  1 Near-end idle cell bit error count 2  2 Far-end HEC violation count 3  3 Far-end idle cell bit error count 4  4 Near-end delineated total cell count (CD-P) 5  5 Near-end user total cell count (CU-P) 6  6 Far-end delineated total cell count (CD-PFE) 7  7 Far-end user total cell count (CU-PFE) 8  NOTE – This number associates the TCA with the specified threshold value attribute of the threshold  data 1 managed entity.  
```
