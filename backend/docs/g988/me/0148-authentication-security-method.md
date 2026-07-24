# Managed Entity

## Identity
- ME ID: 148
- ME Name: Authentication security method
- Source Section: 9.12.4
- Source Page: 428

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Create, Delete, Get, Set
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: Validation scheme
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 2
- Name: Username 1
- Size: 25 bytes
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 3
- Name: Password
- Size: 25 bytes
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 4
- Name: Realm
- Size: 25 bytes
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 5
- Name: Username 2
- Size: 25 bytes
- Format: needs_review
- Access: R, W
- Category: optional

## Extraction Status
- Status: auto_extracted
- Attributes found: 5
- Review needed: false

## Raw Source

```
9.12.4 Authentication security method  The authentication security method defines the user ID and password configuration to establish a  session between a client and a server. This object may be used in the role of the client or server. An  instance of this ME is created by the OLT if authenticated communication is necessary.  Relationships  One instance of this management entity may be associated with a network address ME. This  ME may also be cited by other MEs that require authentication parameter management.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. The value  0xFFFF is reserved. (R, set-by-create) (mandatory) (2 bytes)  Validation scheme : This attribute specifies the validation scheme used when the ONU  validates a challenge. Validation schemes are defined as follows.  0 Validation disabled  1 Validate using MD5 digest authentication as defined in  [IETF RFC 2617] (recommended)  3 Validate using basic authentication as defined in [IETF RFC 2617]  (R, W) (mandatory) (1 byte)  Username 1: This string attribute is the user name. If the string is shorter than 25  bytes, it  must be null terminated (Note). (R, W) (mandatory) (25 bytes)  Password: This string attribute is the password. If the string is shorter than 25  bytes, it  must be null terminated. (R, W) (mandatory) (25 bytes)  Realm: This string attribute specifies the realm used in digest authentication. If the  string is shorter than 25 bytes, it must be null terminated. (R, W) (mandatory)  (25 bytes)  Username 2: This string attribute allows for continuation of the user name beyond  25 characters ( Note). Its default value is a null string.  (R, W) (optional)  (25 bytes)  NOTE – The total username is the concatenation of the username 1 and username 2 attributes if and  only if: a) username 1 comprises 25 non-null characters; b) username 2 is supported by the ONU; and  c) username 2 contains a leading non-null character string. Otherwise, the total username is simply the  value of the username 1 attribute.  Actions  Create, delete, get, set  Notifications  None.  
```
