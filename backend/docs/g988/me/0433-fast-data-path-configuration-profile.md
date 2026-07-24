# Managed Entity

## Identity
- ME ID: 433
- ME Name: FAST data path configuration profile
- Source Section: 9.7.54
- Source Page: 346

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Create, Delete, Get, Set
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

<!-- No attributes successfully parsed; see Raw Source section -->

## Extraction Status
- Status: auto_extracted
- Attributes found: 0
- Review needed: true

## Raw Source

```
9.7.54 FAST data path configuration profile  This ME contains FAST the data path configuration profile for an xDSL UNI. An instance of this ME  is created and deleted by the OLT.  Relationships  An instance of this ME may be associated with zero or more instances of the PPTP UNI part  3.  The overall FAST line configuration profile is modelled in several parts, all of which are  associated together through a common ME ID . (The client PPTP xDSL UNI part 3 refers to  the entire set of line configuration parts).  Attributes  Managed entity ID : This attribute uniquely identifies each instance of this ME. All FAST  line configuration profiles and extensions that pertain to a given PPTP xDSL  UNI must share a common ME ID. (R, set-by-create) (mandatory) (2 bytes)  TPS-TC testmode (TPS_TESTMODE) : This Boolean attribute specifies whether the  TPS-TC test mode defined in clause 8.3.1 [ITU-T G.9701] is enabled (true=1)  or disabled (false=0). See clause 7.3.1 of [ITU-T G.997.2]. (R, W) (mandatory)  (1 byte)  DRA testmode (DRA_TESTMODE): This Boolean attribute d efines whether the dynamic  resource allocation (DRA) testmode defined in clause  9.8.3.1.2 of [ITU-T  G.9701] is enabled (true=1) or disabled  (false=0). See clause 7.3. 2 of  [ITU-T G.997.2] for detailed specification. (R, W) (optional) (1 byte)  Actions  Create, delete, get, set  Notifications  None. 

---- page break ---- 
```
