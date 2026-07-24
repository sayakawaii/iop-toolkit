# Managed Entity

## Identity
- ME ID: 272
- ME Name: GAL Ethernet profile
- Source Section: 9.2.7
- Source Page: 111

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Create, Delete, Get, Set
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: Maximum GEM payload size
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
9.2.7 GAL Ethernet profile  This ME organizes data that describe the gigabit-capable passive optical network transmission  convergence layer (GTC) adaptation layer processing functions of the ONU for Ethernet services. It  is used with the GEM IW TP ME.  Instances of this ME are created and deleted on request of the OLT.  Relationships  An instance of this ME may be associated with zero or more instances of the GEM IW TP  ME.  Attributes  Managed entity  ID: This attribute uniquely identifies each instance of this ME. (R,  set-by-create) (mandatory) (2 bytes)  Maximum GEM payload size: This attribute defines the maximum payload size generated  in the associated GEM IW TP ME. (R, W, set-by-create) (mandatory) (2 bytes)  Actions  Create, delete, get, set  Notifications  None. 

---- page break ---- 
```
