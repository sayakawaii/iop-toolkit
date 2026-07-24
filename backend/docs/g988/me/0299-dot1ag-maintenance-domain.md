# Managed Entity

## Identity
- ME ID: 299
- ME Name: Dot1ag maintenance domain
- Source Section: 9.3.19
- Source Page: 179

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Create, Delete, Get, Set
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: MD level
- Size: 1 byte
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 2
- Name: MD name format
- Size: 1 byte
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 3
- Name: Sender ID permission
- Size: 1 byte
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

## Extraction Status
- Status: auto_extracted
- Attributes found: 3
- Review needed: false

## Raw Source

```
9.3.19 Dot1ag maintenance domain  In [IEEE 802.1ag], a maintenance domain (MD) is a context within which configuration fault  management (CFM) connectivity verification can occur. Individual services (maintenance  associations, MAs) exist within an MD. A n MD is created and deleted by the OLT. The MD ME is  specified by [IEEE 802.1ag] in such a way that the same provisioning can be used for all associated  systems in a network; the OMCI definition accordingly avoids ONU -specific information such as  pointers.  Relationships  Several MDs may be associated with a given bridge, at various MD levels, and a given MD  may be associated with any number of bridges.  Attributes  Managed entity ID: This attribute uniquely identifies an instance of this ME. The values 0  and 0xFFFF are reserved. (R, set-by-create) (mandatory) (2 bytes)  MD level: This attribute ranges from 0..7 and specifies the maintenance level of this MD.  Higher numbers have wider geographic scope. (R,  W, set-by-create)  (mandatory) (1 byte)  MD name format: This attribute specifies one of several possible formats for the MD name  attribute. (R, W, set-by-create) (mandatory) (1 byte)    Value MD name format MD name attribute Defined in  1 None No MD name present [IEEE 802.1ag] 

---- page break ---- Value MD name format MD name attribute Defined in  2 DNS-like name Globally unique text string derived  from a DNS name  "  3 MAC addr and  UINT  MAC address, followed by a 2-octet  unsigned integer, total length 8 bytes  "  4 Character string String of printable characters. This is  recommended to be the default value.  "  32 ICC-based ITU carrier code followed by locally  assigned UMC code, 13 bytes with  trailing nulls as needed  Annex A of   [ITU-T Y.1731]  Others Reserved  MD name 1 , MD name 2 :These two attributes may be regarded as a 50  byte octet string  whose value is the left-justified maintenance domain name. The MD name may  or may not be a printable character string, so an octet string is the appropriate  representation. If the MD name format specifies a DNS -like name or a  character st ring, the string is null -terminated; otherwise, its length is  determined by the MD name format. If the MD has no name (MD name  format = 0), this attribute is undefined. Note that binary comp arisons of the  MD name are made in other CFM state machines, so blanks, alphabetic case,  etc., are significant. Also, note that the MD name and the MA name must be  packed (with additional bytes) into 48  byte CFM message headers. (R,  W)  (mandatory if MD name format is not 1) (25 bytes * 2 attributes)  Maintenance domain intermediate point half function  (MHF) creation: This attribute  determines whether an associated bridge creates an M HF for this MD, under  circumstances defined in clause 22.2.3 of [IEEE 802.1ag]. This attribute is an  enumeration with the following values.  1 None  2 Default (IEEE 802.1ag term). The bridge can create MHFs on an  associated VID on any port through which the VID can pass, where: i)  there are no lower active MD levels or ii) there is a  maintenance  association end point (MEP) at the next lower active MD level on the  port.  3 Explicit. The bridge can create MHFs on an associated VID on any port  through which the VID can pass, but only if a n MEP exists at some  lower maintenance level.  (R, W, set-by-create) (mandatory) (1 byte)  Sender ID permission: This attribute determines the content s of the sender ID type-length- value (TLV) included in CFM messages transmitted by maintenance points  (MPs) controlled by this MD. Chassis ID and management address information  is available from the dot1ag chassis-management info ME. The attribute is an  enumeration with the following values.  1 None: the sender ID TLV is not to be sent.  2 Chassis: the chassis ID length, chassis ID subtype, and chassis ID fields  of the sender ID TLV are to be sent, but not the management address  fields.  3 Manage: the management address fields of the sender ID TLV are to  be sent, but the chassis ID length is to be transmitted with the value 0,  and the chassis ID subtype, and chassis ID fields are not to be sent. 

---- page break ---- 4 ChassisManage: all chassis ID and management address fields are to  be sent.  (R, W, set-by-create) (mandatory) (1 byte)  Actions  Create, delete, get, set  Notifications  None.  
```
