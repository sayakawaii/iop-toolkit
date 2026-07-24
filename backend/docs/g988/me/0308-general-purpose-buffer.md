# Managed Entity

## Identity
- ME ID: 308
- ME Name: General purpose buffer
- Source Section: 9.12.12
- Source Page: 435

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Create, Delete, Get
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: Maximum size
- Size: 4 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: optional

## Extraction Status
- Status: auto_extracted
- Attributes found: 1
- Review needed: false

## Raw Source

```
9.12.12 General purpose buffer  This ME is created by the OLT when needed to store the results of an operation, such as a test  command, that needs to return a block of data of indeterminate size. The buffer is retrieved with get  next operations, since its size is not known a priori. An instance of this ME is created and deleted by  the OLT, and typically made known to an ONU ME or to an action through a pointer.  The ME is defined as generically as possible, such that it can be used for other applications that may  not initially be apparent, such as logging. The format of its content is specific to each application, and  is documented there.  The general purpose buffer is neither captured in an MIB upload, nor retained in a non-volatile ONU  memory.  Relationships  Through a pointer, the OLT may associate a general purpose buffer with an ME or an  operation that has a need to create large or indeterminate blocks of data for subsequent upload.  Attributes  Managed entity  ID: This attribute uniquely identifies each instance of this ME. (R,  set-by-create) (mandatory) (2 bytes)  Maximum size: The ONU determines the actual size of the buffer table in the process of  capturing the data directed to it. The maximum size attribute permits the OLT  to restrict the maximum size of the buffer table. The value 0 indicates that the  OLT imposes no limit on the size; it is recognized that ONU implementations  will impose their own limits. The ONU will not create a buffer table larger than  the value of this attribute. If the ONU cannot allocate enough memory to  accommodate this size, it should deny the ME create action or a write operation  that attempts to expand an existing ME. (R,  W, set-by-create) (optional)  (4 bytes)  Buffer table: This attribute is an octet string that contains the result of some operation  performed on the ONU. The exact content depends on the operation, and is  documented with the definition of each operation. (R) (mandatory) (N bytes)  Actions  Create, delete, get, get next  Notifications  Attribute value change  Number Attribute value  change  Description  1 N/A  

---- page break ---- Attribute value change  Number Attribute value  change  Description  2 Buffer table This AVC indicates that the ONU has completed writing the  buffer, and thereby signals to the OLT that the operation is  complete and the buffer is available for upload.  3..16 Reserved   
```
