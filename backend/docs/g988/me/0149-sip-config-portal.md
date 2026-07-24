# Managed Entity

## Identity
- ME ID: 149
- ME Name: SIP config portal
- Source Section: 9.9.19
- Source Page: 416

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Get
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

<!-- No attributes successfully parsed; see Raw Source section -->

## Extraction Status
- Status: auto_extracted
- Attributes found: 0
- Review needed: true

## Raw Source

```
9.9.19 SIP config portal  The SIP config portal ME provides a way for the OLT to discover the configuration text delivered to  an ONU by a non -OMCI SIP VoIP configuration method ([BBF TR -069]/[BBF TR-369], sipping  framework, etc.). Text retrieved from this ME is not required to be understood by the OLT or EMS,  but it may be useful for human or vendor-specific analysis tools. See also the MGC config portal ME.  An instance of this ME may be created by an ONU that supports non -OMCI SIP configuration. It is  not reported during an MIB upload.  Relationships  One instance of this ME is associated with the ONU.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. There is one  instance, number 0. (R) (mandatory) (2 bytes)  Configuration text table: This attribute is used to pass a textual representation of the VoIP  configuration back to the OLT. The contents are vendor-specific. The get, get  next sequence must be used with this attribute since its size is unspecified.  Upon ME instantiation, the ONU sets this attribute to 0. (R) (mandatory)  (x bytes)  Actions  Get, get next 

---- page break ---- Notifications  Attribute value change  Number Attribute value  change Description  1 Configuration text Indicates an update to the VoIP configuration from a non- OMCI interface. Because the attribute is a table, the AVC  does not contain information about its value. The OLT must  use the get, get next action sequence if it wishes to obtain the  updated attribute content.  2..16 Reserved   
```
