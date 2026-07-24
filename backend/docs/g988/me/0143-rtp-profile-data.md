# Managed Entity

## Identity
- ME ID: 143
- ME Name: RTP profile data
- Source Section: 9.9.7
- Source Page: 398

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Create, Delete, Get, Set
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: Local port min
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 2
- Name: Local port max
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set- by-create
- Category: optional

### Attribute 3
- Name: DSCP mark
- Size: 1 byte
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 4
- Name: Piggyback events
- Size: 1 byte
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 5
- Name: Tone events
- Size: 1 byte
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 6
- Name: DTMF events
- Size: 1 byte
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 7
- Name: CAS events
- Size: 1 byte
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 8
- Name: IP host config pointer
- Size: 2 bytes
- Format: needs_review
- Access: R, W
- Category: optional

## Extraction Status
- Status: auto_extracted
- Attributes found: 8
- Review needed: false

## Raw Source

```
9.9.7 RTP profile data  This ME configures RTP. It  is conditionally required for ONUs that offer VoIP service. If a non - OMCI interface is used to manage VoIP, this ME is unnecessary.  An instance of this ME is created and deleted by the OLT. An RTP profile is needed for each unique  set of attributes.  Relationships  An instance of this ME may be associated with one or more VoIP media profile MEs.  Attributes  Managed entity  ID: This attribute uniquely identifies each instance of this ME. (R,  set-by-create) (mandatory) (2 bytes)  Local port min : This attribute defines the base UDP port that should be used by RTP for  voice traffic. The recommended default is 50000 (R,  W, set-by-create)  (mandatory) (2 bytes)  Local port max: This attribute defines the highest UDP port used by RTP for voice traffic.  The value must be greater than the local port minimum. The value 0 specifies  that the local port max imum be equal to the local port min imum. (R, W, set- by-create) (optional) (2 bytes)  DSCP mark: Diffserv code point to be used for outgoing RTP packets for this profile. The  recommended default value is expedited forwarding (EF)  = 0x2E. (R,  W,  set-by-create) (mandatory) (1 byte)  Piggyback events: Enables or disables RTP piggyback events.  0 Disabled (recommended default)  1 Enabled  (R, W, set-by-create) (mandatory) (1 byte)  Tone events: Enables or disables the handling of tones via RTP tone events per [IETF RFC  4733], (see also [IETF RFC 4734]).  0 Disabled (recommended default)  1 Enabled  (R, W, set-by-create) (mandatory) (1 byte)  DTMF events : Enables or disables the handling of DTMF via RTP DTMF events per  [IETF RFC 4733], (see also [IETF RFC 4734]). This attribute is ignored unless  the OOB DTMF attribute in the VoIP media profile is enabled.  0 Disabled  1 Enabled  (R, W, set-by-create) (mandatory) (1 byte)  CAS events: Enables or disables the handling of CAS via RTP CAS events per  [IETF RFC 4733], (see also [IETF RFC 4734]).  0 Disabled  1 Enabled  (R, W, set-by-create) (mandatory) (1 byte)  IP host config pointer : This optional pointer associates the bearer (voice) flow with an IP  host config data or IPv6 host config data ME. If this attribute is not present or  is not populated with a valid pointer value, the bearer flow uses the same IP  stack that is used for sign alling, indicated by the TCP/UDP pointer in the 

---- page break ---- associated SIP agent or MGC config data. The default value is 0xFFFF, a null  pointer. (R, W) (optional) (2 bytes)  Actions  Create, delete, get, set  Notifications  None.  
```
