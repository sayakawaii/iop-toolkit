# Managed Entity

## Identity
- ME ID: 439
- ME Name: OpenFlow config data
- Source Section: 9.12.18
- Source Page: 441

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Create, Delete, Get, Set
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: TP pointer
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 2
- Name: Version
- Size: 1 byte
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 3
- Name: DS OpenFlow message
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 4
- Name: DS receiving status
- Size: 1 byte
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 5
- Name: US OpenFlow message
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 6
- Name: US forwarding status
- Size: 1 byte
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 7
- Name: Circuit ID
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 8
- Name: Operational state
- Size: 1 byte
- Format: needs_review
- Access: R
- Category: optional

## Extraction Status
- Status: auto_extracted
- Attributes found: 8
- Review needed: false

## Raw Source

```
9.12.18 OpenFlow config data  This ME contains the configuration data whose underlying transport method is OpenFlow. Instances  of this ME are created and deleted by the OLT.  Relationships  An instance of this ME is associated with each instance per OpenFlow transportation channel.  There might be more than one OpenFlow transportation channel per ONU.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. (R, set-by- create) (mandatory) (2 bytes) 

---- page break ---- TP type:  This attribute specifies the type of ANI-side TP associated with this ME.   1 GEM IW TP   (R, W, set-by-create) (mandatory) (1 byte)  TP pointer: This attribute points to the instance of the TP associated with this  OpenFlow configuration data. The type of the associated TP is determined  by the TP type attribute. (R, W, set-by-create) (mandatory) (2 bytes)  Version: This integer attribute reports the version of the OpenFlow protocol in use.  The ONU should deny an attempt by the OLT to set or create a value that  it does not support. The value 0 indicates that no particular version is  specified. (R, W, set-by-create) (mandatory) (1 byte)   DS OpenFlow message: This attribute specifies the DS OpenFlow message which is carried  over the OMCC channel. (R, W) (mandatory) (24N bytes)  DS forwarding control: This Boolean attribute indicates the current DS OpenFlow message  is ready to be sent (true) or not. The default value is false. (R, W) (mandatory)  (1 byte)  DS receiving status : This Boolean attribute indicates the ONU is ready to accept a new  downstream packet (true) or not. The default value is false. (R) (mandatory)  (1 byte)  US OpenFlow message: This attribute specifies the US OpenFlow message which is carried  over the OMCC channel. (R, W) (mandatory) (24N bytes)  US receiving control: This Boolean attribute controls the current US OpenFlow message is  safely received (true) or not. The default value is false. (R, W) (mandatory)  (1 byte)  US forwarding status: This Boolean attribute reports the current US OpenFlow message is  ready to be sent (true) or not. The default value is false. (R) (mandatory) (1 byte)  Circuit ID:  This attribute identifies the first access information of the user  (R, W) (Optional) (24N byte)  Remote ID:  This attribute identifies the second access information of the user as an addition  identifier to circuit ID. (R, W) (Optional) (24N byte)  Administrative state: This attribute locks (1) and unlocks (0) the functions performed by the  MPLS pseudowire  TP. Administrative state is further described in  clause A.1.6. (R, W) (optional) (1 byte)  Operational state: This attribute reports whether the ME is currently capable of performing  its function. Valid values are enabled (0) and disabled (1). (R) (optional)  (1 byte)  Actions  Create, delete, get, get next, set  Notifications  Attribute value change  Number Attribute value change Description  1..5 N/A   6 DS receiving status The DS packets receiving status has changed 

---- page break ---- Attribute value change  Number Attribute value change Description  7..8 N/A   9 US forwarding status A new ONU response has been loaded into the table for the  OLT to retrieve  10-12 N/A   13 Op state Operational state  14..16 reserved   
```
