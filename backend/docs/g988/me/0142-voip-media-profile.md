# Managed Entity

## Identity
- ME ID: 142
- ME Name: VoIP media profile
- Source Section: 9.9.5
- Source Page: 391

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Create, Delete, Get, Set
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: Fax mode
- Size: 1 byte
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 2
- Name: Voice service profile pointer
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 3
- Name: OOB DTMF
- Size: 1 byte
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 4
- Name: RTP profile pointer
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

## Extraction Status
- Status: auto_extracted
- Attributes found: 4
- Review needed: false

## Raw Source

```
9.9.5 VoIP media profile  The VoIP media profile ME contains settings that apply to VoIP voice encoding.  This entity is  conditionally required for ONUs that offer VoIP services. If a non-OMCI interface is used to manage  VoIP signalling, this ME is unnecessary.  An instance of this ME is created and deleted by the OLT. A VoIP media profile is needed for each  unique set of profile attributes.  Relationships  An instance of this ME may be associated with one or more VoIP voice CTP MEs. 

---- page break ---- Attributes  Managed entity  ID: This attribute uniquely identifies each instance of this ME. (R,  set-by-create) (mandatory) (2 bytes)  Fax mode: Selects the fax mode; values are as follows.  0 Passthru  1 ITU-T T.38  (R, W, set-by-create) (mandatory) (1 byte)  Voice service profile pointer : Pointer to a voice service profile, which defines parameters  such as jitter buffering and echo cancellation. (R,  W, set-by-create)  (mandatory) (2 bytes)  Codec selection (1st order) : This attribute specifies codec selection as defined by  [IETF RFC 3551].    Value Encoding name Clock  rate (Hz)  0 PCMU 8000  1 reserved   2 reserved   3 GSM 8000  4 ITU-T G.723 8000  5 DVI4 8000  6 DVI4 16000  7 LPC 8000  8 PCMA 8000  9 ITU-T G.722 8000  10 L16, 2 channels 44100  11 L16, 1 channel 44100  12 QCELP 8000  13 CN 8000  14 MPA 90000  15 ITU-T G.728 8000  16 DVI4 11025  17 DVI4 22050  18 ITU-T G.729 8000  (R, W, set-by-create) (mandatory) (1 byte)  Packet period selection (1st order) : This attribute specifies the packet period selection  interval in milliseconds. The recommended default value is 10  ms. Valid  values are 10..30 ms. (R, W, set-by-create) (mandatory) (1 byte)  Silence suppression (1st order): This attribute specifies whether silence suppression is on or  off. Valid values are 0  = off and 1  = on. (R,  W, set-by-create) (mandatory)  (1 byte)  Three more groups of three attributes are defined, with definitions identical to the preceding three: 

---- page break ---- Codec selection (2nd order): (R, W, set-by-create) (mandatory) (1 byte)  Packet period selection (2nd order): (R, W, set-by-create) (mandatory) (1 byte)  Silence suppression (2nd order): (R, W, set-by-create) (mandatory) (1 byte)  Codec selection (3rd order): (R, W, set-by-create) (mandatory) (1 byte)  Packet period selection (3rd order): (R, W, set-by-create) (mandatory) (1 byte)  Silence suppression (3rd order): (R, W, set-by-create) (mandatory) (1 byte)  Codec selection (4th order): (R, W, set-by-create) (mandatory) (1 byte)  Packet period selection (4th order): (R, W, set-by-create) (mandatory) (1 byte)  Silence suppression (4th order): (R, W, set-by-create) (mandatory) (1 byte)    OOB DTMF: This attribute specifies out-of-band DMTF carriage. When enabled (1), DTMF  signals are carried out of band via RTP or the associated signalling protocol.  When disabled (0), DTMF tones are carried in the PCM stream. (R,  W,  set-by-create) (mandatory) (1 byte)  RTP profile pointer : This attribute points to the associated RTP profile data ME. (R,  W,  set-by-create) (mandatory) (2 bytes)  Actions  Create, delete, get, set  Notifications  None.  
```
