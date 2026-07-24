# Managed Entity

## Identity
- ME ID: 155
- ME Name: MGC config data
- Source Section: 9.9.16
- Source Page: 410

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Create, Delete, Get, Set
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: Primary MGC
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 2
- Name: Secondary MGC
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 3
- Name: TCP/UDP pointer
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 4
- Name: Version
- Size: 1 byte
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 5
- Name: Maximum retry time
- Size: 2 bytes
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 6
- Name: Maximum retry attempts
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: optional

### Attribute 7
- Name: Service change delay
- Size: 2 bytes
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 8
- Name: Termination ID base
- Size: 25 bytes
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 9
- Name: Softswitch
- Size: 4 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 10
- Name: Message ID pointer
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: optional

## Extraction Status
- Status: auto_extracted
- Attributes found: 10
- Review needed: false

## Raw Source

```
9.9.16 MGC config data  The MGC config data ME defines the MGC configuration associated with an MG subscriber. It is  conditionally required for ONUs that support ITU -T H.248 VoIP services. If a non-OMCI interface  is used to manage VoIP signalling, this ME is unnecessary.  Instances of this ME are created and deleted by the OLT.  Relationships  An instance of this ME may be associated with one or more VoIP voice CTP MEs.  Attributes  Managed entity  ID: This attribute uniquely identifies each instance of this ME. (R,  set-by-create) (mandatory) (2 bytes)  Primary MGC : This attribute points to a network address ME that contains the name  (IP address or resolved name) of the primary MGC that controls the signalling  messages. The port is optional and defaults to 2944 for text message formats  and 2955 for binary message formats. (R,  W, set-by-create) (mandatory)  (2 bytes)  Secondary MGC : This attribute points to a network address ME that contains the name  (IP address or resolved name) of the secondary or backup MGC that controls  the signalling messages. The port is optional and defaults to 2944 for text  message formats and 2955 for binary message formats. (R,  W, set-by-create)  (mandatory) (2 bytes)  TCP/UDP pointer : This attribute points to the TCP/UDP config data ME to be used for  communication with the MGC. (R, W, set-by-create) (mandatory) (2 bytes)  Version: This integer attribute reports the version of the Megaco protocol in use. The  ONU should deny an attempt by the OLT to set or create a value that it does  not support. The value 0 indicates that no particular version is specified. (R, W,  set-by-create) (mandatory) (1 byte) 

---- page break ---- Message format: This attribute defines the message format. Valid values are as follows.  0 Text long  1 Text short  2 Binary  The default value is recommended to be 0. (R,  W, set-by-create) (mandatory)  (1 byte)  Maximum retry time: This attribute specifies the maximum retry time for MGC transactions,  in seconds. The default value 0 specifies vendor -specific implementation.  (R, W) (optional) (2 bytes)  Maximum retry attempts: This attribute specifies the maximum number of times a message  is retransmitted to the MGC. The recommended default value 0 specifies  vendor-specific implementation. (R, W, set-by-create) (optional) (2 bytes)  Service change delay: This attribute specifies the service status delay time for changes in line  service status. This attribute is specified in seconds. The default value 0  specifies no delay. (R, W) (optional) (2 bytes)  Termination ID base: This attribute specifies the base string for the ITU -T H.248 physical  termination ID(s) for this ONU. This string is intended to uniquely identify an  ONU. Vendor-specific termination identifiers (port IDs) are optionally added  to this string to uniquely identify a termination on a specific ONU. (R,  W)  (optional) (25 bytes)  Softswitch: This attribute identifies the gateway softswitch vendor. The format is four  ASCII coded alphabetic characters [A..Z] as defined in [ATIS -0300220]. A  value of four null bytes indicates an unknown or unspecified vendor. (R,  W,  set-by-create) (mandatory) (4 bytes)  Message ID pointer: This attribute points to a large string whose value specifies the message  identifier string for ITU -T H.248 messages originated by the ONU. (R, W,  set-by-create) (optional) (2 bytes)  Actions  Create, delete, get, set  Notifications  Alarm  Alarm  number Alarm Description  0 Timeout Timeout of association with MG  1..207 Reserved   208..223 Vendor-specific alarms Not to be standardized  
```
