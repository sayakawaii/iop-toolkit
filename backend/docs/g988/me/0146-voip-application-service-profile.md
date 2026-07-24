# Managed Entity

## Identity
- ME ID: 146
- ME Name: VoIP application service profile
- Source Section: 9.9.8
- Source Page: 399

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Create, Delete, Get, Set
- Instance Type: singleton
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: CID features
- Size: 1 byte
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 2
- Name: Call waiting features
- Size: 1 byte
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 3
- Name: Call progress or transfer features
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 4
- Name: Call presentation features
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 5
- Name: Direct connect feature
- Size: 1 byte
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 6
- Name: Direct connect URI pointer
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 7
- Name: Bridged line agent URI pointer
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 8
- Name: Conference factory URI pointer
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 9
- Name: IP host pointer
- Size: 2 bytes
- Format: needs_review
- Access: R, W
- Category: optional

## Extraction Status
- Status: auto_extracted
- Attributes found: 9
- Review needed: false

## Raw Source

```
9.9.8 VoIP application service profile  The VoIP application service profile defines attributes of calling features used in conjunction with a  VoIP line service. It is optional for ONUs that support VoIP services. If a non-OMCI interface is used  to manage SIP for VoIP, this ME is unnecessary.  An instance of this ME is created and deleted by the OLT. A VoIP application service profile instance  is needed for each unique set of profile attributes.  Relationships  An instance of this ME is associated with zero or more SIP user data MEs.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME.  (R, set-by-create) (mandatory) (2 bytes)  CID features: This attribute contains a bit map of caller ID features. Except as noted, the bit  value 0 disables the feature; 1 enables it.  0x01 Calling number  0x02 Calling name  0x04 CID blocking (both number and name)  0x08 CID number – Permanent presentation status for  number (0 = public, 1 = private)  0x10 CID name – Permanent presentation status for name   (0 = public, 1 = private)  0x20 Anonymous CID blocking (ACR). It may not be  possible to support this in the ONU.  0x40 GR-1188 Absence of Calling Name Reason disable.  (0=transmit "O" ASCII code in MDMF message, 1=do  not transmit "O" ASCII code.  0x80 Not used  The recommended default value is 0x00. (R, W, set-by-create) (mandatory)  (1 byte)  Call waiting features: This attribute contains a bit map of call waiting features. The bit  value 0 disables the feature; 1 enables it.  0x01 Call waiting  0x02 Caller ID announcement  0x04..0x80 Not used  The recommended default value is 0x00. (R, W, set-by-create) (mandatory)  (1 byte)  Call progress or transfer features: This attribute is a bit map of call processing features.  The bit value 0 disables the feature; 1 enables it.  0x0001 3way  0x0002 Call transfer 

---- page break ---- 0x0004 Call hold (RFC 3264 sendonly SDP call hold)  0x0008 Call park  0x0010 Do not disturb  0x0020 Flash on emergency service call (flash is to be  processed during an emergency service call)  0x0040 Emergency service originating hold (determines  whether call clearing is to be performed on on-hook  during an emergency service call)  0x0080 6way  0x0100 Call hold (RFC 2543 connection address 0.0.0.0 call  hold)  0x0200..0x8000 Not used  The recommended default value is 0x0000. (R, W, set-by-create)  (mandatory) (2 bytes)  Call presentation features: This attribute is a bit map of call presentation features. The bit  value 0 disables the feature; 1 enables it.  0x0001 Message waiting indication splash ring  0x0002 Message waiting indication special dial tone  0x0004 Message waiting indication visual indication  0x0008 Call forwarding indication  0x0010 DC voltage based visual message waiting indicator  (vmwi) (e.g., neon lamp on a phone to indicate a  message waiting). For backwards compatibility  reasons, the value 0x0010 is a companion value to  0x0004. If an ONU does not support DC voltage  vmwi, the ONU uses other existing vmwi methods. If  the ONU supports DC voltage vmwi and needs to  apply DC voltage to turn on the phone lamp (to  indicate message waiting), the values 0x0004 and  0x0010 are set.  0x0020..0x8000 Not used  The recommended default value is 0x0000. (R, W, set-by-create)  (mandatory) (2 bytes)  Direct connect feature: This attribute is a bit map of characteristics associated with the  direct connect feature. The bit value 0 disables the feature; 1 enables it.  0x01 Direct connect feature enabled  0x02 Dial tone feature delay option  The recommended default value is 0x00. (R, W, set-by-create) (mandatory)  (1 byte)  Direct connect URI pointer: This attribute points to a network address ME that specifies  the URI of the direct connect. If this attribute is set to a null pointer, no URI  is defined. (R, W, set-by-create) (mandatory) (2 bytes)  Bridged line agent URI pointer: This attribute points to a network address ME that  specifies the URI of the bridged line agent. If this attribute is set to a null  pointer, no URI is defined. (R, W, set-by-create) (mandatory) (2 bytes)  Conference factory URI pointer: This attribute points to a network address ME that  specifies the URI of the conference factory. If this attribute is set to a null  pointer, no URI is defined. (R, W, set-by-create) (mandatory) (2 bytes) 

---- page break ---- Dial tone feature delay/warmline timer (new): This attribute defines the warmline  timer/dial tone feature delay timer (seconds). The default value 0 specifies  vendor-specific implementation. (R, W) (optional) (2 bytes)  IP host pointer: This attribute points to the IP host config data or IPv6 host config data ME  associated with this VoIP config data ME. This attribute is only relevant when  the VoIP configuration method used attribute of this ME is set to configuration  file retrieval (2) OR IETF sipping config framework (4). Upon instantiation  ONU sets this value to NULL (0xFFFF) pointer. (R, W) (optional) (2 bytes)  Actions  Create, delete, get, set  Notifications  None.  
```
