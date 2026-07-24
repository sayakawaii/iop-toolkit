# Managed Entity

## Identity
- ME ID: 300
- ME Name: Dot1ag maintenance association
- Source Section: 9.3.20
- Source Page: 181

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Create, Delete, Get, Set
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: MD pointer
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 2
- Name: Short MA name format
- Size: 1 byte
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 3
- Name: Associated VLANs
- Size: 1 byte
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 4
- Name: Sender ID permission
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
9.3.20 Dot1ag maintenance association  This ME models an [IEEE 802.1ag] service defined on a bridge port. An MA is a set of endpoints on  opposite sides of a network, all existing at a defined maintenance level. One of the endpoints resides  on the local ONU; the others are understood to be configured in a consistent way on external  equipment. [ITU-T Y.1731] refers to the MA as a maintenance entity group (MEG).  An MA is created and deleted by the OLT.  Relationships  Any number of MAs may be associated with a given MD, or may stand on their own without  an MD. One or more MAs may be associated with a MAC bridge or an IEEE 802.1p mapper.  An MA exists at one of eight possible maintenance levels.  Attributes  Managed entity ID: This attribute uniquely identifies an instance of this ME. The values 0  and 0xFFFF are reserved. (R, set-by-create) (mandatory) (2 bytes)  MD pointer: This pointer specifies the dot1ag maintenance domain with which this MA is  associated. A null pointer specifies that the MA is not associated with an MD.  (R, W, set-by-create) (mandatory) (2 bytes)  Short MA name format: This attribute specifies one of several possible formats for the short  MA name attribute. Value 1, the primary VLAN ID, is recommended to be the  default. (R, W, set-by-create) (mandatory) (1 byte)    Value Short MA name format Short MA name attribute  1 Primary VID 2 octets, 12 LSBs specify primary VID, 0  if none  2 Character string String of up to 45 printable characters  3 2-octet integer 2 octet unsigned integer  4 Virtual private network  (VPN) ID  7 octets, as defined in [IETF RFC 2685]  32 ICC-based ITU carrier code followed by locally  assigned UMC code, 13 bytes with trailing  nulls as needed. Defined in Annex A of  [ITU-T Y.1731]  Other Reserved   Short MA name 1, Short MA name 2 : These two attributes may be regarded as an octet  string whose value is the left-justified MA name. Because the MA name may  or may not be a printable character string, an octet string is the appropriate  representation. If the short MA name format specifies a character string, the 

---- page break ---- string is null-terminated; otherwise, its length is determined by the short MA  name format. Note that binary comparisons of the short MA name are made in  other CFM state machines, so blanks, alphabetic case, etc., are significant.  Also, note that the MD name and the MA short name must be packed (with  additional bytes) into 48  byte CFM message headers. (R,  W) (mandatory)  (25 bytes * 2 attributes)  Continuity check message  (CCM) interval: If CCMs are enabled on a n MEP, the CCM  interval attribute specifies the rate at which they are generated. The MEP also  expects to receive CCMs from each of the other MEPs in its CC database at  this rate.  0: CCM transmission disabled  1: 3.33 ms   2: 10 ms  3: 100 ms  4: 1 s  5: 10 s  6: 1 min  7: 10 min  Short intervals should be used judiciously, as they can interfere with the  network's ability to handle subscriber traffic. The recommended value is 1  s.  (R, W, set-by-create) (mandatory) (1 byte)  Associated VLANs: This attribute is a list of up to 12 VLAN IDs with which this MA is  associated. Once a set of VLANs is defined, the ONU should deny operations  to other dot1ag MAs or dot1ag default MD level entries that conflict with the  set membership. The all-zeros value indicates that this MA is not associated  with any VLANs. Assuming that the attribute is not 0, the first entry is  understood to be the primary VLAN. Except forwarded linktrace messages   (LTMs), CFM messages emitted by MPs in this MA are t agged with the  primary VLAN ID. (R, W) (mandatory) (2 bytes/entry * 12 entries = 24 bytes)  MHF creation : This attribute determines whether the bridge creates a n MHF, under  circumstances defined in clause 22.2.3 of [IEEE 802.1ag]. This attribute is an  enumeration with the following values:  1 None. No MHFs are created on this bridge for this MA.  2 Default (IEEE 802.1ag term). The bridge can create MHFs on this VID  on any port through which the VID can pass.  3 Explicit. The bridge can create MHFs on this VID on any port through  which the VID can pass, but only if a n MEP exists at some lower  maintenance level.  4 Defer. This value causes the ONU to use the setting of the parent MD.  This is recommended to be the default value.  (R, W, set-by-create) (mandatory) (1 byte)  Sender ID permission: This attribute determines the contents of the sender ID TLV included  in CFM messages transmitted by MPs controlled by this MA. This attribute is  the same as that defined in the description of the dot1ag MD ME, with the  addition of code point 5.  1 None: the sender ID TLV is not to be sent.  2 Chassis: the chassis ID length, chassis ID subtype, and chassis ID fields  of the sender ID TLV are to be sent, but not the management address  fields. 

---- page break ---- 3 Manage: the management address fields of the sender ID TLV are to  be sent, but the chassis ID length is to be transmitted with a 0 value,  and the chassis ID subtype, and chassis ID fields are not to be sent.  4 ChassisManage: all chassis ID and management address fields are to  be sent.  5 Defer: the content s of the sender ID TLV are determined by the  corresponding MD attribute. This is recommended to be the default  value.  (R, W, set-by-create) (mandatory) (1 byte)  Actions  Create, delete, get, set  Notifications  None.  
```
