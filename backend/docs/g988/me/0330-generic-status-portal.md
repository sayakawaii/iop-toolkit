# Managed Entity

## Identity
- ME ID: 330
- ME Name: Generic status portal
- Source Section: 9.12.14
- Source Page: 438

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Create, Delete, Get
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
9.12.14 Generic status portal  The generic status portal (GSP) ME provides a way for the OLT to discover the status and  configuration information of a non -OMCI management domain within an ONU. The non -OMCI  management domain is indicated by the VEIP associated with this GSP.  The GSP ME uses two table attributes to convey status and configuration from a non-OMCI managed  domain to the OMCI. Each of these attributes uses an XML document to present this information.  These XML documents are not required to be understood by the OLT or EMS. The schema for the  documents may be used in the creation of tools that parse and interpret the contents of the document.  Relationships  One instance of this ME is created by the OLT for each separate non -OMCI management  domain whose information is desired to be visible.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. Through an  identical ID, the GSP ME is implicitly linked to an instance of the VEIP ME.  (R, set-by-create) (mandatory) (2 bytes). (R)  Status document table : This attribute is used to pass a textual representation of the non - OMCI managed domain status back to the OLT. The contents are  vendor-specific and formatted as an XML document. The first element of the  document must contain an XML declaration indicating the version of XML  and encoding used in the remainder of the document. The second element of  the document must contain a schema reference to the vendor-supplied schema  used by the remainder of the document. The get, get next sequence must be  used with this attribute since its size is unspecified. (R) (mandatory) (N bytes)  Configuration document table: This attribute is used to pass a textual representation of the  non-OMCI managed domain configuration back to the OLT. The contents are  vendor-specific and formatted as an XML document. The first element of the  document must contain an XML declaration indicating the version of XML  and encoding used in the remainder of the document. The second element of  the document must contain a schema reference to the vendor-supplied schema  used by the remainder of the document. The get, get next sequence must be  used with this attribute since its size is unspecified. (R) (mandatory) (M bytes) 

---- page break ---- AVC report rate : This attribute governs the rate at which the GSP generates AVC  notifications. The default value 0 disables AVCs, while the highest value 3,  which may be useful for debugging, generates an AVC on every change seen  in the non -OMCI management domain. As a guideline, the value 1 should  collect changes into not more than one notification in 10 min, while value 2  should generate an AVC not more than once per second. (R, W, set-by-create)  (optional) (1 byte)  Actions  Create, delete, Get, get next  Notifications  Attribute value change  Number Attribute value  change  Description  1 Status document table Indicates an update to the status document table from a  non-OMCI interface. Because the attribute is a table, the  AVC does not contain information about its value. The OLT  must use the get, get next action sequence if it wishes to  obtain the updated attribute content.  2 Configuration  document table  Indicates an update to the configuration document table from  a non-OMCI interface. Because the attribute is a table, the  AVC does not contain information about its value. The OLT  must use the get, get next action sequence if it wishes to  obtain the updated attribute content.  3..16 Reserved   
```
