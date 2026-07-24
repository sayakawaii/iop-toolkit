# Managed Entity

## Identity
- ME ID: 16
- ME Name: AAL5 profile
- Source Section: 9.13.5
- Source Page: 458

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Create, Delete, Get, Set
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: Max CPCS PDU size
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 2
- Name: SSCS type
- Size: 1 byte
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

## Extraction Status
- Status: auto_extracted
- Attributes found: 2
- Review needed: false

## Raw Source

```
9.13.5 AAL5 profile  This ME organizes data that describe the AAL type 5 processing functions of the ONU. It is used  with the IW VCC TP ME.  This ME is created and deleted by the OLT.  Relationships  An instance of this ME may be associated with zero or more instances of the IW VCC TP.  Attributes  Managed entity  ID: This attribute uniquely identifies each instance of this ME. (R,  set-by-create) (mandatory) (2 bytes)  Max CPCS PDU size : This attribute specifies the maximum CPCS PDU size to be  transmitted over the connection in both upstream and downstream directions.  (R, W, set-by-create) (mandatory) (2 bytes) 

---- page break ---- AAL mode: This attribute specifies the AAL mode as follows.  0 Message assured  1 Message unassured  2 Streaming assured  3 Streaming non assured  (R, W, set-by-create) (mandatory) (1 byte)  SSCS type: This attribute specifies the SSCS type for the AAL. Valid values are as follows.  0 Null  1 Data SSCS based on SSCOP, assured operation  2 Data SSCS based on SSCOP, non-assured operation  3 Frame relay SSCS  (R, W, set-by-create) (mandatory) (1 byte)  Actions  Create, delete, get, set  Notifications  None.  
```
