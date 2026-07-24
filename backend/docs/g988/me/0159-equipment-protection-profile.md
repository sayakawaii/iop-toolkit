# Managed Entity

## Identity
- ME ID: 159
- ME Name: Equipment protection profile
- Source Section: 9.1.11
- Source Page: 87

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Create, Delete, Get, Set
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: Protect slot 1, Protect slot 2
- Size: 1 byte
- Format: needs_review
- Access: RWSC
- Category: optional

### Attribute 2
- Name: Wait to restore time
- Size: 1 byte
- Format: needs_review
- Access: RWSC
- Category: optional

## Extraction Status
- Status: auto_extracted
- Attributes found: 2
- Review needed: false

## Raw Source

```
9.1.11 Equipment protection profile  This ME supports equipment protection. There can be as many as two protection slots protecting as  many as eight working slots. Each of the working and protect cardholder MEs should refer to the  equipment protection profile that defines its protection group. Instances of this ME are created and  deleted by the OLT.  An ONU should deny pre-provisioning that would create impossible protection groupings because of  slot or equipment incompatibilities. In the same way, the ONU should deny creation or addition to  protection groups that cannot be supported by the current equ ipped configuration . Even so, an  inconsistent card type alarm is defined, for example, to cover the case of a plug-and-play circuit pack  installed in a protection group cardholder that cannot support it.  Relationships  An instance of this object points to the working and protect cardholders, which in turn point  back to this ME.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. The first byte  is 0. The second byte is assigned by the OLT, and must be unique and non-zero.  (R, set-by-create) (mandatory) (2 bytes)  Protect slot 1, Protect slot 2:This pair of attributes describes the protecting cardholder  entities in an equipment protection group. There can be one or two protecting  entities.  0 Undefined entry (default), a place -holder if there are fewer than  two protecting entities in the protection group.  1..254 Slot number of the protecting circuit pack.  (RWSC) (protect slot 1 mandatory, protect slot 2 optional) (1  byte * 2  attributes)  Working slot 1, Working slot 2, Working slot 3, Working slot 4, Working slot 5,  Working slot 6, Working slot 7, Working slot 8: This group of attributes  describes the working cardholder entities in an equipment protection group.  There can be up to eight working entities.  0 Undefined entry (default), a place -holder if there are fewer than  eight working entities in the protection group.  1..254 Slot number of the working circuit pack.  (RWSC) (working slot 1 mandatory, other working slots optional) (1 byte * 8  attributes)  Protect status 1, Protect status 2: This pair of attributes indicates whether each protection  cardholder is currently protecting another cardholder, and if so, which one.  0 Not protecting any other cardholder.  1..254 Slot number of the working cardholder currently being protected  by this ME.  (R) (mandatory) (1 byte * 2 attributes)  Revertive ind: This attribute specifies whether equipment protection is revertive. The default  value 0 indicates revertive switching; any other value indicates non -revertive  switching. (RWSC) (optional) (1 byte)  Wait to restore time : This attribute specifies the time, in minutes, during which a working  equipment must be free of error before a revertive switch occurs. It defaults to  0. (RWSC) (optional) (1 byte) 

---- page break ---- Actions  Create, delete, get, set  Notifications  Alarm  Alarm  number  Alarm Description  0 Inconsistent card type The expected or actual circuit pack type in a slot is incapable  of participating in the equipment protection group, either  because it is not subject to equipment protection or because  its type or equipment ID differs from that previously defined  for the other cardholders of the group. When possible, the  ONU should deny provisioning attempts that would create  incompatibilities, for example, in the case of plug-and-play,  it may not be possible to forestall the inconsistency.  1..207 Reserved   208..223 Vendor-specific alarms Not to be standardized  
```
