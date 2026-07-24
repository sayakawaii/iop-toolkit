# Managed Entity

## Identity
- ME ID: 154
- ME Name: MGC config portal
- Source Section: 9.9.20
- Source Page: 417

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
9.9.20 MGC config portal  The MGC config portal ME provides a way for the OLT to discover the configuration text delivered  to an ONU by a non -OMCI ITU-T H.248 VoIP configuration method. Text retrieved from this ME  is not required to be understood by the OLT or EMS, but it may be useful for human or vendor - specific analysis tools. See also the SIP config portal ME.  An instance of this ME may be created by an ONU that supports non -OMCI ITU -T H.248  configuration. It is not reported during an MIB upload.  Relationships  One instance of this ME is associated with the ONU.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. There is one  instance, number 0. (R, set-by-create) (mandatory) (2 bytes)  Configuration text table: This attribute is used to pass a textual representation of the VoIP  configuration back to the OLT. The contents are vendor-specific. The get, get  next sequence must be used with this attribute since its size is unspecified.  Upon ME instantiation, the ONU sets this attribute to 0. (R) (mandatory) ( x  bytes)  Actions  Get, get next  Notifications  Attribute value change  Number Attribute value change Description  1 Configuration text Indicates an update to the VoIP configuration from a non- OMCI interface. Because the attribute is a table, the AVC  does not contain information about its value. The OLT must  use the get, get next action sequence if it wishes to obtain the  updated attribute content.  2..16 Reserved   
```
