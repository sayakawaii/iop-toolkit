# Managed Entity

## Identity
- ME ID: 137
- ME Name: Network address
- Source Section: 9.12.3
- Source Page: 427

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Create, Delete, Get, Set
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: Security pointer
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 2
- Name: Address pointer
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

## Extraction Status
- Status: auto_extracted
- Attributes found: 2
- Review needed: false

## Raw Source

```
9.12.3 Network address  The network address ME associates a network address with security methods required to access a  server. It is conditionally required for ONUs that support VoIP service s. The address may take the  form of a URL, a fully qualified path or IP address represented as an ACII string.  If a non-OMCI interface is used to manage VoIP signalling, this ME is unnecessary.  Instances of this ME are created and deleted by the OLT or the ONU, depending on the method used  and case.  Relationships  Any ME that requires a network address may link to an instance of this ME.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. Instances of  this ME created autonomously by the ONU have IDs in the range 0..0x7FFF.  Instances created by the OLT have IDs in the range 0x8000..0xFFFE.  The  value 0xFFFF is reserved. (R, set-by-create) (mandatory) (2 bytes)  Security pointer : This attribute points to an authentication security method ME. The  authentication security method indicates the username and password to be used  when retrieving the network address indicated by this ME. A null pointer  indicates that security attributes are not defined for this network address.  (R, W, set-by-create) (mandatory) (2 bytes)  Address pointer : This attribute points to the large string ME that contains the network  address. It may contain a fully qualified domain name, URI or IP address. The  URI may also contain a port identifier (e.g., "x.y.z.com:5060"). A null pointer  indicates that no network address is defined. (R, W, set-by-create) (mandatory)  (2 bytes)  Actions  Create, delete, get, set  Notifications  None. 

---- page break ---- 
```
