# Managed Entity

## Identity
- ME ID: 21
- ME Name: CES service profile
- Source Section: 9.8.3
- Source Page: 361

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Create, Delete, Get, Set
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: CES buffered CDV tolerance
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

## Extraction Status
- Status: auto_extracted
- Attributes found: 1
- Review needed: false

## Raw Source

```
9.8.3 CES service profile  NOTE – In [ITU-T G.984.4], this ME is called a CES service profile-G.  An instance of this ME organizes data that describe the CES service functions of the ONU. Instances  of this ME are created and deleted by the OLT. 

---- page break ---- Relationships  An instance of this ME may be associated with zero or more instances of a GEM IW TP.  Attributes  Managed entity  ID: This attribute uniquely identifies each instance of this ME. (R,  set-by-create) (mandatory) (2 bytes)  CES buffered CDV tolerance : This attribute represents the duration of user data that must  be buffered by the CES IW entity to offset packet delay variation. It is  expressed in 10  µs increments. 75 (750  μs) is suggested as a default value.  (R, W, set-by-create) (mandatory) (2 bytes)  Channel associated signalling (CAS): This attribute selects the signalling format. It applies  to structured interfaces only. For unstructured interfaces, this value, if present,  must be set to the default 0. Valid values are as follows.  0 Basic  1 E1 CAS  2 SF CAS  3 DS1 ESF CAS  4 J2 CAS  (R, W, set-by-create) (optional) (1 byte)  Actions  Create, delete, get, set  Notifications  None.  
```
