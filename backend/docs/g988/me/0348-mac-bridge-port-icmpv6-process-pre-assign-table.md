# Managed Entity

## Identity
- ME ID: 348
- ME Name: MAC bridge port ICMPv6 process pre-assign table
- Source Section: 9.3.33
- Source Page: 213

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Get, Set
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: ICMPv6 informational messages processing
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 2
- Name: Router solicitation processing
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 3
- Name: Router advertisement processing
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 4
- Name: Neighbour solicitation processing
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 5
- Name: Neighbour advertisement processing
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 6
- Name: Redirect processing
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 7
- Name: Multicast listener query processing
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 8
- Name: Unknown ICMPv6 processing
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: mandatory

## Extraction Status
- Status: auto_extracted
- Attributes found: 8
- Review needed: false

## Raw Source

```
9.3.33 MAC bridge port ICMPv6 process pre-assign table  This ME provides an approach to ICMPv6 message processing configuration to those ONUs that  support IPv6 awareness. For every message, the MAC bridge port ICMPv6 process pre -assign table  can designate a forward, discard or snoop operation. The ONU creates or deletes an instance of this  ME automatically upon creation or deletion of a MAC bridge port configuration data ME. 

---- page break ---- The MAC bridge port ICMPv6 process pre -assign table ME filters layer 2 traffic between the UNI  and ANI. The operation of this ME is completely independent of the operation and traffic generated  or received by a possible IPv6 host config data ME.  Relationships  An instance of this ME is associated with an instance of a MAC bridge port configuration data  ME.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. Through an  identical ID, this ME is implicitly linked to an instance of the MAC bridge port  configuration data ME. (R) (mandatory) (2 bytes)  The following nine attributes have similar definitions. Each permits the OLT to specify ICMPv6 as  the next header in the IPv6 header and various types in the ICMPv6 header, and whether traffic of  the specified type is forwarded, discarded or snooped, in up stream and downstream directions  separately. The bits of each attribute are assigned as follows.    Bit Name Setting  1..2 (LSB) Process for upstream 00: forward  01: discard  10: snoop  3..4 Process for downstream 00: forward  01: discard  10: snoop  5..8 Reserved 0  The initial value of each attribute is given in the last column of the table.    No. Protocol Next header type Standard Initial value  1 ICMPv6 error messages 58 1-4 [b-IETF RFC 2460]  [b-IETF RFC 2463]  0  2 ICMPv6 informational  messages  58 128,129 [b-IETF RFC 2460]  [b-IETF RFC 2463]  0  3 Neighbour discovery  router solicitation  58 133 [b-IETF RFC 2460]  [b-IETF RFC 4861]  6  4 Neighbour discovery – router advertisement  58 134 [b-IETF RFC 2460]  [b-IETF RFC 4861]  9  5 Neighbour discovery – neighbour solicitation  58 135 [b-IETF RFC 2460]  [b-IETF RFC 4861]  0  6 Neighbour discovery – neighbour  advertisement  58 136 [b-IETF RFC 2460]  [b-IETF RFC 4861]  0  7 Neighbour discovery –  redirect  58 137 [b-IETF RFC 2460]  [b-IETF RFC 4861]  1 

---- page break ---- No. Protocol Next header type Standard Initial value  8 MLD – Multicast  listener query (MLDv1,  MLDv2)  58 130 [b-IETF RFC 2710]  [IETF RFC 3810]  1  9 Unknown ICMPv6 58 – – 5    ICMPv6 error messages processing: (R, W) (mandatory) (1 byte)  ICMPv6 informational messages processing: (R, W) (mandatory) (1 byte)  Router solicitation processing: (R, W) (mandatory) (1 byte)  Router advertisement processing: (R, W) (mandatory) (1 byte)  Neighbour solicitation processing: (R, W) (mandatory) (1 byte)  Neighbour advertisement processing: (R, W) (mandatory) (1 byte)  Redirect processing: (R, W) (mandatory) (1 byte)  Multicast listener query processing: (R, W) (mandatory) (1 byte)  NOTE – If the ONU participates in multicast services, MLD queries should be controlled through the  multicast operations profile ME. In such a case, it is strongly recommended not to provision the  downstream direction of the multicast listener query process ing attribute to any value other than  forwarding.  Unknown ICMPv6 processing: (R, W) (mandatory) (1 byte)  Actions  Get, set.  
```
