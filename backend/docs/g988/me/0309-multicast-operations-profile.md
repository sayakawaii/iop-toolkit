# Managed Entity

## Identity
- ME ID: 309
- ME Name: Multicast operations profile
- Source Section: 9.3.27
- Source Page: 195

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Create, Delete, Get, Set
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: IGMP version
- Size: 1 byte
- Format: needs_review
- Access: R,W, set-by-create
- Category: mandatory

### Attribute 2
- Name: IGMP function
- Size: 1 byte
- Format: needs_review
- Access: R,W, set-by-create
- Category: mandatory

### Attribute 3
- Name: Immediate leave
- Size: 1 byte
- Format: needs_review
- Access: R,W, set-by-create
- Category: mandatory

### Attribute 4
- Name: Upstream IGMP TCI
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: optional

### Attribute 5
- Name: Upstream IGMP rate
- Size: 4 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: optional

### Attribute 6
- Name: Dynamic access control list table
- Size: 1 byte
- Format: needs_review
- Access: R, W, set-by-create
- Category: optional

### Attribute 7
- Name: Querier IP address
- Size: 4 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: optional

### Attribute 8
- Name: Query interval
- Size: 4 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: optional

### Attribute 9
- Name: Query max response time
- Size: 4 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: optional

### Attribute 10
- Name: Last member query interval
- Size: 4 bytes
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 11
- Name: Unauthorized join request behaviour
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: optional

## Extraction Status
- Status: auto_extracted
- Attributes found: 11
- Review needed: false

## Raw Source

```
9.3.27 Multicast operations profile  This ME expresses multicast policy. A multi -dwelling unit ONU may have several such policies,  which are linked to subscribers as required. Some of the attributes  configure IGMP snooping and  proxy parameters if the defaults do not suffice, as described in [IETF RFC 2236], [IETF RFC 3376],  [IETF RFC 3810] and [IETF RFC 5519]. Instances of this ME are created and deleted by the OLT.  Relationships  An instance of this ME may be associated with zero or more instances of the multicast  subscriber config info ME.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. The values 0  and 0xFFFF are reserved. (R, set-by-create) (mandatory) (2 bytes)  IGMP version: This attribute specifies the version of IGMP to be supported. Support of a  given version implies compatible support of previous versions. If the ONU  cannot support the version requested, it should deny an attempt to set the  attribute. (R,W, set-by-create) (mandatory) (1 byte)  1 IGMP version 1 (deprecated)  2 IGMP version 2  3 IGMP version 3  16 MLD version 1  17 MLD version 2  Other values are reserved.  IGMP function: This attribute enables an IGMP function. The value 0 specifies transparent  IGMP snooping only. The value 1 specifies snooping with proxy reporting  (SPR); the value 2 specifies IGMP proxy. The function must be consistent with  the capabilities specified by the other IGMP configuration attributes. (R,W,  set-by-create) (mandatory) (1 byte)  Immediate leave: This Boolean attribute controls the immediate leave function. The value  false disables immediate leave; true enables immediate leave. (R,W,  set-by-create) (mandatory) (1 byte)  Upstream IGMP TCI : Under control of the upstream IGMP tag control attribute, the  upstream IGMP TCI attribute defines a VLAN ID and P-bits to add to upstream  IGMP messages. (R, W, set-by-create) (optional) (2 bytes) 

---- page break ---- Upstream IGMP tag control : This attribute controls the upstream IGMP TCI attribute. If  this attribute is non-zero, a possible extended VLAN tagging operation ME is  ignored for upstream frames containing IGMP/MLD packets. (R,  W,  set-by-create) (optional) (1 byte)  Value Meaning  0 Pass upstream IGMP/MLD traffic transparently, neither adding,  stripping nor modifying tags that may be present.  1 Add a VLAN tag (including P bits) to upstream IGMP/MLD  traffic. The tag is specified by the upstream IGMP TCI attribute.  2 Replace the entire TCI (VLAN ID plus P bits) on upstream  IGMP/MLD traffic. The new tag is specified by the upstream  IGMP/MLD TCI attribute. If the received IGMP/MLD traffic  is untagged, an add operation is performed.  3 Replace only the VLAN ID on upstream IGMP/MLD traffic,  retaining the original DEI and P bits. The new VLAN ID is  specified by the VLAN ID field of the upstream IGMP TCI  attribute. If the received IGMP/MLD traffic is untagged, an add  operation is performed, with DEI and P bits also taken from the  upstream IGMP TCI attribute.  Other values are reserved.  Upstream IGMP rate : This attribute limits the maximum rate of upstream IGMP traffic.  Traffic in excess of this limit is silently discarded. The attribute value is  specified in messages/second. The recommended default value 0 imposes no  rate limit on this traffic. (R, W, set-by-create) (optional) (4 bytes)  Dynamic access control list table: This attribute is a list that specifies one or more multicast  group address ranges. Each row in the list comprises up to three row parts,  where each row part is 24 bytes long. Each entry must include row part 0. The  ONU may also support row parts 1-2, thus allowing the table to contain logical  rows that exceed the 24 byte definition of row part 0.  Table control (2 bytes)  16 15 14 13 12 11 10 9 8 7 6 5 4 3 2 1  Set ctrl Row part ID Test Row key  The first 2 bytes of each row part is the table control field, which comprises a  key into the row, the row part identifier, and fields to define the result of a set  operation and to test whether the ONU supports the extended table format.  It is the responsibility of the OLT to assign and track row keys and content.  The ONU should deny set operations that create range overlaps.  Set ctrl  The two MSBs of this field determine the meaning of a set operation. These  bits are returned as 00 during get next operations.  Bits 16..15 Meaning  00 Reserved  01 Write this entry into the table. Overwrite any  existing entry with the same row part ID and row  key.  10 Delete this entry from the table, including all row  parts. The remaining fields are not meaningful. 

---- page break ---- Bits 16..15 Meaning  11 Clear all entries from the table. The remaining fields  are not meaningful.  Row part ID  The row part ID field distinguishes the row part associated with the current set  or get operation.  Row part 0 is backward compatible with earlier versions of this ME definition.  Row parts 1 -2 are optional on a row by row basis. They can be set by using  values 001-010 as the row part ID. Row parts 3-7 are reserved.    Bits 14..12 Meaning  000 The associated row part has format 0.  001 The associated row part has format 1.  010 The associated row part has format 2.  011..111 Reserved  Test  This bit allows the OLT to determine whether an ONU supports the extended  format access control list. If the ONU does not support the extended format, it  should be possible to set the test bit to 1 and read it back with a get and get  next operation. If the ONU does support the extended format, this bit should  always return the value 0 under a get next operation.  Row key  The row key distinguishes rows in the table.  Row part definition    Byte Row part 0 Row part 1 Row part 2  1 Table control  (2 bytes)  Table control  (2 bytes)  Table control  (2 bytes) 2  3 GEM port ID  (2 bytes)  Leading bytes of  IPv6 source address  (12 bytes)  Leading bytes of  IPv6 destination  address  (12 bytes)  4  5 VLAN ID (ANI)  (2 bytes) 
```
