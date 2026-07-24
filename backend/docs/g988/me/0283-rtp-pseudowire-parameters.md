# Managed Entity

## Identity
- ME ID: 283
- ME Name: RTP pseudowire parameters
- Source Section: 9.8.6
- Source Page: 367

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Create, Delete, Get, Set
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: Clock reference
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 2
- Name: RTP timestamp mode
- Size: 1 byte
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 3
- Name: PTYPE
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 4
- Name: Expected PTYPE
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: optional

### Attribute 5
- Name: Expected SSRC
- Size: 8 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: optional

## Extraction Status
- Status: auto_extracted
- Attributes found: 5
- Review needed: false

## Raw Source

```
9.8.6 RTP pseudowire parameters  If a pseudowire service uses RTP, the RTP pseudowire parameters ME provides configuration  information for the RTP layer. Instances of this ME are created and deleted by the OLT. The use of  RTP on a pseudowire is optional, and is determined by the existence of the RTP pseudowire  parameters ME.  Relationships  An instance of the RTP pseudowire parameters ME may exist for each pseudowire TP ME,  to which it is implicitly bound by a common ME ID.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. Through an  identical ID, this ME is implicitly linked to an instance of the pseudowire TP  ME. (R, set-by-create) (mandatory) (2 bytes)  Clock reference: This attribute specifies the frequency of the common timing reference, in  multiples of 8 kHz. (R, W, set-by-create) (mandatory) (2 bytes)  RTP timestamp mode : This attribute determines the mode in which RTP timestamps are  generated in the TDM to the PSN direction.  0 Unknown or not applicable.  1 Absolute. Timestamps are based on the timing of the incoming TDM  signal.  2 Differential. Timestamps are based on the ONU 's reference clock,  which is understood to be stratum -traceable along with the reference  clock at the far end.  (R, W, set-by-create) (mandatory) (1 byte)  PTYPE: This attribute specifies the RTP payload type in the TDM to the PSN direction.  It comprises two 1 byte values. The first is for the payload channel, the second,  for the optional separate signalling channel. Assignable PTYPEs lie in the  dynamic range 96..127. If signalling is not transported in its own channel, the  second value should be set to 0. (R, W, set-by-create) (mandatory) (2 bytes) 

---- page break ---- SSRC: This attribute specifies the RTP synchronization source in the TDM to the PSN  direction. It comprises two 4 byte values. The first is for the payload channel,  the second, for the optional separate signalling channel. If signalling is not  transported in its own channel, the second value should be set to 0. (R,  W,  set-by-create) (mandatory) (8 bytes)  Expected PTYPE: This attribute specifies the RTP payload type in the PSN to the TDM  direction. The received payload type may be used to detect malformed packets.  It comprises two 1 byte values. The first is for the payload channel, the second,  for the optional separate signalling channel. To disable either or both of the  check functions, set the corresponding value to its default value 0. (R,  W,  set-by-create) (optional) (2 bytes)  Expected SSRC: This attribute specifies the RTP synchronization source in the PSN to the  TDM direction. The received SSRC may be used to detect misconnection  (stray packets). It comprises two 4 byte values. The first is for the payload  channel, the second, for the optional separate signalling channel. To disable  either or both of the check functions, set the corresponding value to its default  value 0. (R, W, set-by-create) (optional) (8 bytes)  Actions  Create, delete, get, set  Notifications  None.  
```
