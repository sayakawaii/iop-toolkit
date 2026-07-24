# Managed Entity

## Identity
- ME ID: 307
- ME Name: Octet string
- Source Section: 9.12.11
- Source Page: 434

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Create, Delete, Get, Set
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: Length
- Size: 2 bytes
- Format: needs_review
- Access: R, W
- Category: mandatory

## Extraction Status
- Status: auto_extracted
- Attributes found: 1
- Review needed: false

## Raw Source

```
9.12.11 Octet string  The octet string is modelled on the large string ME. The large string is constrained to printable  characters because it uses null as a trailing delimiter. The octet string has a length attribute and is  therefore suitable for arbitrary sequences of bytes.  Instances of this ME are created and deleted by the OLT. To use this ME, the OLT instantiates the  octet string ME and then points to the created ME from other ME instances. Systems that maintain  the octet string should ensure that the octet string ME is not deleted while it is still linked.  Relationships  An instance of this ME may be cited by any ME that requires an octet string that can exceed  25 bytes in length.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. The values 0  and 0xFFFF are reserved. (R, set-by-create) (mandatory) (2 bytes)  Length: This attribute specifies the number of octets that comprise the sequence of  octets. This attribute defaults to 0 to indicate no octet string is defined. The  maximum value of this attribute is 375 (15 parts, 25  bytes each). (R,  W)  (mandatory) (2 bytes)  In the following, 15  additional attributes are defined; they are identical. The octet string is simply  divided into as many parts as necessary, starting at part 1 and left justified. 

---- page break ---- Part 1, Part 2, Part 3, Part 4, Part 5, Part 6, Part 7, Part 8, Part 9,   Part 10, Part 11, Part 12, Part 13, Part 14, Part 15:   (R, W) (part 1 mandatory, others optional) (25 bytes * 15 attributes)  Actions  Create, delete, get, set  Notifications  None.  
```
