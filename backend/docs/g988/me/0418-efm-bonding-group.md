# Managed Entity

## Identity
- ME ID: 418
- ME Name: EFM bonding group
- Source Section: 9.7.41
- Source Page: 325

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Create, Delete, Get, Set
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: Group ID
- Size: 6 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 2
- Name: Minimum upstream group rate
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
9.7.41 EFM bonding group  The EFM bonding group represents a group of links that are bonded. In [IEEE 802.3], a bonding  group is known as a PAF [physical medium entity (PME) aggregation function] and a link is known  as a PME instance of this ME are created and deleted by the OLT.  Relationships  An instance of this ME may be associated with zero or more instances of an EFM bonding  link.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. The value 0  is reserved. (R, set-by-create) (mandatory) (2 bytes)  Group ID: This attribute is the unique number representing this bonding group. See clause  C.3.1.1 of [ITU-T G.998.2]. (R, W, set-by-create) (mandatory) (6 bytes)  Minimum upstream group rate : This attribute sets the minimum upstream group rate, in  bits per second, for this EFM Group. This attribute is used to determine the  group US rate low alarm status. The group US rate low alarm means that the  aggregate upstream rate of all active links associated with this group is less  than the minimum upstream group rate. The default value for this rate is zero.  (R, W) (mandatory, set-by-create) (4 bytes)  Minimum downstream group rate: This attribute sets the minimum downstream group rate,  in bits per second, for this EFM Group. This attribute is used to determine the  group DS rate low alarm status. The group DS rate low alarm means that the  aggregate downstream rate of all active links associated with this group is less  than the minimum downstream group rate. The default value for this rate is  zero. (R, W) (mandatory) (4 bytes, set-by-create)  Group alarm enable: This bit mapped attribute enables the various group alarms. A bit value  of 1 means "enable".  Bit Meaning  1 (LSB) Group down  2 Group partial  3 Group US rate low  4 Group DS rate low 

---- page break ---- 5 4x rate ratio  6-8 Reserved  (R, W, set-by-create) (mandatory) (1 byte)  Actions  Create, delete, get, set  Notifications  Alarm  Alarm number Alarm Description  0 Group down No links associated with this group are active  1 Group partial Not all links associated with this group are active  2 Group US rate low Aggregate upstream rate is less than the minimum upstream  group rate  3 Group DS rate low Aggregate downstream rate is less than the minimum  downstream group rate  4 4x rate ratio In this group, ratio of max link rate to min link rate > 4  5..207 Reserved   NOTE – An "active" link means that the port is trained and fragments can flow across the link in both  directions.  
```
