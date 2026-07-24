# Managed Entity

## Identity
- ME ID: 434
- ME Name: FAST vectoring line configuration extensions
- Source Section: 9.7.55
- Source Page: 347

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
9.7.55 FAST vectoring line configuration extensions  This ME extends FAST line configuration MEs with attributes that are specific to vectoring. An  instance of this ME is created and deleted by the OLT.  Relationships  An instance of this ME is associated with a PPTP UNI part 3.  The overall FAST line configuration MEs are modelled in several parts, all of which are  associated together through a common ME ID (The client PPTP xDSL UNI part 3 refers to  the entire set of line configuration parts).  Attributes  Managed entity ID : This attribute uniquely identifies each instance of this ME. All FAST  line configuration profiles and extensions that pertain to a given PPTP xDSL  UNI must share a common ME ID. (R, set-by-create) (mandatory) (2 bytes)  FEXT cancellation enabling/disabling upstream ( FEXT_TO_CANCEL_ENABLEus):  A value of 1 enables and a value of 0 disables FEXT cancellation in the  upstream direction from all the other vectored lines into the line in the vectored  group. See clause 7.1.7.2 of [ITU-T G.997.2]. (R, W) (mandatory) (1 byte)  FEXT cancellation enabling/disabling downstream (FEXT_TO_CANCEL_ENABLEds):  A value of 1 enables and a value of 0 disables FEXT cancellation in the  downstream direction from all the other vectored lines into the line in the  vectored group. See clause 7.1. 7.1 of [ITU-T G.997.2]. (R, W) (mandatory)  (1 byte)  Actions  Create, delete, get, set  Notifications  None.  
```
