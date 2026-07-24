# Managed Entity

## Identity
- ME ID: 141
- ME Name: VoIP line status
- Source Section: 9.9.11
- Source Page: 403

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Get
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: Voip codec used
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 2
- Name: Voip voice server status
- Size: 1 byte
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 3
- Name: Voip port session type
- Size: 1 byte
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 4
- Name: Voip call 1 packet period
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 5
- Name: Voip call 2 packet period
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 6
- Name: Voip call 1 dest addr
- Size: 25 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 7
- Name: Voip call 2 dest addr
- Size: 25 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 8
- Name: Voip line state
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
9.9.11 VoIP line status  The VoIP line status ME contains line status information for POTS ports using VoIP service s. An  ONU that supports VoIP automatically creates or deletes an instance of this ME upon creation or  deletion of a PPTP POTS UNI.  Relationships  An instance of this ME is associated with a PPTP POTS UNI.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. Through an  identical ID, this ME is implicitly linked to an instance of the PPTP POTS  UNI. (R) (mandatory) (2 bytes)  Voip codec used : Reports the current codec used for a VoIP POTS port. Valid values are  taken from [IETF RFC 3551], and are the same as specified in the codec  selection attribute of the VoIP media profile. This attribute is meaningful only  if the VoIP port session type attribute is not idle.  0 PCMU  1 reserved  2 reserved  3 GSM  4 ITU-T G.723  5 DVI4, 8 kHz  6 DVI4, 16 kHz  7 LPC  8 PCMA  9 ITU-T G.722  10 L16, 2 channels  11 L16, 1 channel  12 QCELP  13 CN  14 MPA  15 ITU-T G.728  16 DVI4, 11.025 kHz  17 DVI4, 22.050 kHz  18 ITU-T G.729  (R) (mandatory) (2 bytes)  Voip voice server status: Status of the VoIP session for this POTS port:  0 None/initial  1 Registered 

---- page break ---- 2 In session  3 Failed registration – icmp error  4 Failed registration – failed tcp  5 Failed registration – failed authentication  6 Failed registration – timeout  7 Failed registration – server fail code  8 Failed invite – icmp error  9 Failed invite – failed tcp  10 Failed invite – failed authentication  11 Failed invite – timeout  12 Failed invite – server fail code  13 Port not configured  14 Config done  15 Disabled by switch  (R) (mandatory) (1 byte)  Voip port session type: This attribute reports the current state of a VoIP POTS port session:  0 Idle/none  1 2way  2 3way  3 Fax/modem  4 Telemetry  5 Conference  (R) (mandatory) (1 byte)  Voip call 1 packet period : This attribute reports the packet period for the first call on the  VoIP POTS port. The value is defined in milliseconds. (R) (mandatory)  (2 bytes)  Voip call 2 packet period: This attribute reports the packet period for the second call on the  VoIP POTS port. The value is defined in milliseconds. (R) (mandatory)  (2 bytes)  Voip call 1 dest addr: This attribute reports the DA for the first call on the VoIP POTS port.  The value is an ASCII string. (R) (mandatory) (25 bytes)  Voip call 2 dest addr : This attribute reports the DA for the second call on the VoIP POTS  port. The value is an ASCII string. (R) (mandatory) (25 bytes)  Voip line state : This attribute reports the state of the POTS line. This attribute may not be  meaningful if the POTS port is administratively locked, is operationally  disabled, or is being tested. Code points are assigned as follows:  0 Idle, on-hook  1 Off-hook dial tone  2  Dialling  3 Ringing or FSK alerting/data  4 Audible ringback  5 Connecting  6 Connected  7 Disconnecting, audible indication  8 ROH, no tone  9 ROH with tone  10 Unknown or undefined  (R) (optional) (1 byte) 

---- page break ---- Emergency call status: This attribute reports the current state of an emergency call session  (when the ONU is the call originator) on the VoIP POTS port. The ONU  determines the presence of an originating emergency call on the basis of the  Emergency service number attribute of the VoIP feature access codes ME.  0 No emergency call in progress  1 Emergency call in progress  NOTE – The ONU may also be able to determine the presence of an emergency call  on the basis of other, unspecified information.  (R) (Optional) (1 byte)  Actions  Get  Notifications  Attribute value change  Number Attribute value change Description  1..8 N/A   9 Emergency call status Indicates an update to the Emergency call status attribute  10..16 Reserved   
```
