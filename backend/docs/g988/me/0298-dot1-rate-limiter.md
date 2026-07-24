# Managed Entity

## Identity
- ME ID: 298
- ME Name: Dot1 rate limiter
- Source Section: 9.3.18
- Source Page: 178

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Create, Delete, Get, Set
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: Parent ME pointer
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 2
- Name: Upstream unicast flood rate pointer
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: optional

### Attribute 3
- Name: Upstream broadcast rate pointer
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: optional

### Attribute 4
- Name: Upstream multicast payload rate pointer
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: optional

## Extraction Status
- Status: auto_extracted
- Attributes found: 4
- Review needed: false

## Raw Source

```
9.3.18 Dot1 rate limiter  This ME allows rate limits to be defined for various types of upstream traffic that are processed by  IEEE 802.1 bridges or related structures.  Relationships  An instance of this ME may be linked to an instance of a MAC bridge service profile or an  IEEE 802.1p mapper.  Attributes  Managed entity  ID: This attribute uniquely identifies each instance of this ME. (R,  set-by-create) (mandatory) (2 bytes)  Parent ME pointer : This attribute points to an instance of a ME. The type of ME is  determined by the TP type attribute. (R,  W, set-by-create) (mandatory)  (2 bytes) 

---- page break ---- TP type: This attribute identifies the type of TP associated with this dot1 rate limiter.  Valid values are:  1 MAC bridge service profile  2 IEEE 802.1p mapper service profile  (R, W, set-by-create) (mandatory) (1 byte)  Upstream unicast flood rate pointer : This attribute points to an instance of the traffic  descriptor that governs the rate of upstream unicast packets whose DA is  unknown to the bridge. A null pointer specifies that no administrative limit is  to be imposed. (R, W, set-by-create) (optional) (2 bytes)  Upstream broadcast rate pointer: This attribute points to an instance of the traffic descriptor  that governs the rate of upstream broadcast packets. A null pointer specifies  that no administrative limit is to be imposed. (R,  W, set-by-create) (optional)  (2 bytes)  Upstream multicast payload rate pointer: This attribute points to an instance of the traffic  descriptor that governs the rate of upstream multicast payload packets. A null  pointer specifies that no administrative limit is to be imposed. (R,  W,  set-by-create) (optional) (2 bytes)  Actions  Create, delete, get, set  Notifications  None.  
```
