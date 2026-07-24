# Managed Entity

## Identity
- ME ID: 419
- ME Name: EFM bonding link
- Source Section: 9.7.42
- Source Page: 326

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Create, Delete, Get, Set
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: Associated group ME ID
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 2
- Name: Link alarm enable
- Size: 1 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

## Extraction Status
- Status: auto_extracted
- Attributes found: 2
- Review needed: false

## Raw Source

```
9.7.42 EFM bonding link  The EFM bonding link represents a link that can be bonded with other links to form a group. In  [IEEE 802.3], a bonding group is known as a PAF and a link is known as a PME. Instances of this  ME are created and deleted by the OLT.  Relationships  An instance of this ME may be associated with zero or one instance of an EFM bonding group.  Attributes  Managed entity ID : This attribute uniquely identifies each instance of this ME. The two  MSBs of the first byte are the bearer channel ID. Excluding the first 2 bits of  the first byte, the remaining part of the ME ID is identical to that of this ME's  parent PPTP xDSL UNI part 1.   NOTE – This attribute has the same meaning as the Stream ID in clause C.3.1.2 of  [ITU-T G.998.2], except that it cannot be changed. (R, set-by-create) (mandatory)  (2 bytes)  Associated group ME ID : This attribute is the ME ID of the bonding group to which this  link is associated. Changing this attribute moves the link from one group to  another. Setting this attribute to an ME ID that has not yet been provisioned  will result in this link being place d in a single -link group that contains only  this link. The default value for this attribute is the null pointer, 0xFFFF. (R, W,  set-by-create) (mandatory) (2 bytes)  Link alarm enable : This bit mapped attribute enables the group down and group partial  alarms. A bit value of 1 means "enable".  Bit Meaning  1 (LSB) Link down  2-8 Reserved 

---- page break ---- (R, W, set-by-create) (mandatory) (1 bytes)  Actions  Create, delete, get, set  Notifications  Alarm  Alarm number Alarm Description  0 Link down Link not active. See definition in EFM bonding group ME  1..207 Reserved   
```
