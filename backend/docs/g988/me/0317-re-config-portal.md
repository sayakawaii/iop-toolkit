# Managed Entity

## Identity
- ME ID: 317
- ME Name: RE config portal
- Source Section: 9.14.5
- Source Page: 480

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Get
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: Configuration text table
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
9.14.5 RE config portal  The RE config portal ME provides a way for the OLT to discover the configuration delivered to a n  RE by a non-OMCI configuration method (SNMP, etc.). Text retrieved from this ME is not required  to be understood by the OLT or EMS, but it may be useful for human or vendor -specific analysis  tools.  An instance of this ME may be created by an RE that supports non-OMCI RE configuration. It is not  reported during an MIB upload. 

---- page break ---- Relationships  One instance of this ME is associated with an instance of a TCP/UDP config data ME.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. There is one  instance, number 0. (R) (mandatory) (2 bytes)  Configuration text table : This attribute is used to pass a textual representation of the RE  configuration back to the OLT. The contents are vendor-specific. The get, get  next sequence must be used with this attribute since its size is unspecified.  Upon ME instantiation, the manag ement ONU sets this attribute to 0. (R)  (mandatory) (x bytes)  TCP/UDP pointer: This pointer associates the RE config portal with the TCP/UDP config  data ME to be used for communication with any valid and necessary in-band  protocol server. The default value is 0xFFFF. (R, W) (mandatory) (2 bytes)  Actions  Get, get next  Notifications  Attribute value change  Number Attribute value  change  Description  1 Configuration text Indicates an update to the RE configuration from a non-OMCI  interface. Because the attribute is a table, the AVC does not  contain information about its value. The OLT must use the  get, get next action sequence if it wishes to obtain the updated  attribute content.  2 N/A   3..16 Reserved   
```
