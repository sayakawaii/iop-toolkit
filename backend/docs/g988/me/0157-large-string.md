# Managed Entity

## Identity
- ME ID: 157
- ME Name: Large string
- Source Section: 9.12.5
- Source Page: 428

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Create, Delete, Get, Set
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: Number of parts
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: mandatory

## Extraction Status
- Status: auto_extracted
- Attributes found: 1
- Review needed: false

## Raw Source

```
9.12.5 Large string  The large string ME holds character strings longer than 25 bytes, up to 375 bytes. It is maintained in  up to 15 parts, each part containing 25  bytes. If the final part contains fewer than 25  bytes, it is  terminated by at least one null byte. For example:    Number of parts 3  Part 1 sftp://myusername:mypassw  Part 2 ord@config.telecom.com:12 

---- page break ---- Part 3 34/path/to/filename<null>  Or  Number of parts 3  Part 1 sftp://myusername:mypassw  Part 2 ord@config.telecom.com:12  Part 3 34/path/to/longfilename<null>  Instances of this ME are created and deleted by the OLT. Under some circumstances, they may also  be created by the ONU. To use this ME, the OLT or ONU instantiates the large string ME and then  points to the created ME from other ME instances. Systems that maintain the large string should  ensure that the large string ME is not deleted while it is still linked.  Relationships  An instance of this ME may be cited by any ME that requires a text string longer than 25 bytes.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. The value  0xFFFF is reserved. When the large string is to be used as an IPv6 address, the  value 0 is also reserved. The OLT should create large string MEs starting at 1  (or 0), and numbering upward s. The ONU should create large string MEs  starting at 65534 (0xFFFE) and numbering downward s. (R,  set-by-create)  (mandatory) (2 bytes)  Number of parts: This attribute specifies the number of non-empty parts that form the large  string. This attribute defaults to 0 to indicate no large string content is  defined.(R, W) (mandatory) (1 byte)  Fifteen additional attributes are defined in the following; they are identical. The large string is simply  divided into as many parts as necessary, starting at part 1. If the end of the string does not lie at a part  boundary, it is marked with a null byte.  Part 1, Part 2, Part 3, Part 4, Part 5, Part 6, Part 7, Part 8, Part 9,   Part 10, Part 11, Part 12, Part 13, Part 14, Part 15: (R, W) (mandatory)  (25 bytes * 15 attributes)  Actions  Create, delete, get, set  Notifications  Attribute value change  Number Attribute value change Description  1 Number of parts   2 Part 1   3 Part 2   4 Part 3   5 Part 4   6 Part 5   7 Part 6   8 Part 7   9 Part 8  

---- page break ---- Attribute value change  Number Attribute value change Description  10 Part 9   11 Part 10   12 Part 11   13 Part 12   14 Part 13   15 Part 14   16 Part 15   NOTE – Older implementations of the OMCI may not support this notification, which has been  introduced in this version of this Recommendation.  
```
