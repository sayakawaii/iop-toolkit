# Managed Entity

## Identity
- ME ID: 139
- ME Name: VoIP voice CTP
- Source Section: 9.9.4
- Source Page: 391

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Create, Delete, Get, Set
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: User protocol pointer
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 2
- Name: PPTP pointer
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 3
- Name: VoIP media profile pointer
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 4
- Name: Signalling code
- Size: 1 byte
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

## Extraction Status
- Status: auto_extracted
- Attributes found: 4
- Review needed: false

## Raw Source

```
9.9.4 VoIP voice CTP  The VoIP voice CTP defines the attributes necessary to associate a specified VoIP service (SIP,  ITU-T H.248) with a POTS UNI. This entity is conditionally required for ONUs that offer VoIP   services. If a non-OMCI interface is used to manage VoIP signalling, this ME is unnecessary.  An instance of this ME is created and deleted by the OLT. A VoIP voice CTP ME is needed for each  PPTP POTS UNI served by VoIP.  Relationships  An instance of this ME links a PPTP POTS UNI ME with a VoIP media profile and a SIP  user data or media gateway controller (MGC) config data ME.  Attributes  Managed entity  ID: This attribute uniquely identifies each instance of this ME. (R,  set-by-create) (mandatory) (2 bytes)  User protocol pointer : This attribute points to signalling protocol data. If the signalling  protocol used attribute of the VoIP config data ME specifies that the ONU 's  signalling protocol is SIP, this attribute points to a SIP user data ME, which in  turn points to a SIP agent config data  ME. If the signalling protocol is ITU - T H.248, this attribute points directly to an MGC config data ME. (R,  W,  set-by-create) (mandatory) (2 bytes)  PPTP pointer: This attribute points to the PPTP POTS UNI ME that serves the analogue  telephone port. (R, W, set-by-create) (mandatory) (2 bytes)  VoIP media profile pointer: This attribute points to an associated VoIP media profile. (R, W,  set-by-create) (mandatory) (2 bytes)  Signalling code: This attribute specifies the POTS-side signalling as follows.  1 Loop start  2 Ground start  3 Loop reverse battery  4 Coin first  5 Dial tone first  6 Multi-party  (R, W, set-by-create) (mandatory) (1 byte)  Actions  Create, delete, get, set  Notifications  None.  
```
