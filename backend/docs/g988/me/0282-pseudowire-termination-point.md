# Managed Entity

## Identity
- ME ID: 282
- ME Name: Pseudowire termination point
- Source Section: 9.8.5
- Source Page: 364

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Create, Delete, Get, Set
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: Underlying transport
- Size: 1 byte
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 2
- Name: Service type
- Size: 1 byte
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 3
- Name: Signalling
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 4
- Name: North-side pointer
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 5
- Name: Far-end IP info
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 6
- Name: ARC
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 7
- Name: ARC interval
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: optional

## Extraction Status
- Status: auto_extracted
- Attributes found: 7
- Review needed: false

## Raw Source

```
9.8.5 Pseudowire termination point  The pseudowire TP supports packetized (rather than TDM) transport of TDM services, transported  either directly over Ethernet, over UDP/IP or over MPLS. Instances of this ME are created and deleted  by the OLT.  Relationships  One pseudowire TP ME exists for each distinct TDM service that is mapped to a pseudowire.  Attributes  Managed entity  ID: This attribute uniquely identifies each instance of this ME. (R,  set-by-create) (mandatory) (2 bytes)  Underlying transport:  0 Ethernet, MEF 8  1 UDP/IP  2 MPLS  (R, W, set-by-create) (mandatory) (1 byte)  Service type: This attribute specifies the basic service type, either a transparent bit pipe or  an encapsulation that recognizes the underlying structure of the payload.  0 Basic unstructured (also known as structure agnostic)  1 Octet-aligned unstructured, structure agnostic. Applicable only to DS1,  a mode in which each frame of 193 bits is encapsulated in 25 bytes with  7 padding bits.  2 Structured (structure-locked)  (R, W, set-by-create) (mandatory) (1 byte)  Signalling:  0 No signalling visible at this layer  1 CAS, to be carried in the same packet stream as the payload  2 CAS, to be carried in a separate signalling channel  (R, W, set-by-create) (mandatory for structured service type) (1 byte)  TDM UNI pointer: If service type = structured, this attribute points to a logical N × 64 kbit/s  sub-port CTP. Otherwise, this attribute points to a PPTP CES UNI. (R,  W,  set-by-create) (mandatory) (2 bytes)  North-side pointer: When the pseudowire service is transported via IP, as indicated by the  underlying transport attribute, the north -side pointer attribute points to an  instance of the TCP/UDP config data ME. When the pseudowire service is 

---- page break ---- transported directly over Ethernet, the north -side pointer attribute is not used  – the linkage to the Ethernet flow TP is implicit in the ME IDs. When the  pseudowire service is transported over MPLS, the north -side pointer attribute  points to an instance of the MPLS PW TP. (R, W, set-by-create) (mandatory)  (2 bytes)  Far-end IP info: When the pseudowire service is transported via IP, this attribute points to a  large string ME that contains the URI of the far-end TP, e.g.,   udp://192.168.100.221:5000   udp://pwe3srvr.int.example.net:2222  A null pointer is appropriate if the pseudowire is not transported via IP. (R, W,  set-by-create) (mandatory for IP transport) (2 bytes)  Payload size: Number of payload bytes per packet. Valid only if service type  = basic  unstructured or octet-aligned unstructured. Valid choices depend on the TDM  service, but must include the following. Other choices are at the vendor 's  discretion.  DS1 192  DS1 200, required only if an octet-aligned unstructured service is  supported  E1 256  DS3 1024  E3 1024  (R, W, set-by-create) (mandatory for unstructured service) (2 bytes)  Payload encapsulation delay : Number of 125  µs frames to be encapsulated in each  pseudowire packet. Valid only if service type  = structured. The minimum set  of choices for various TDM services is listed in the following table , and is  affected by the possible presence of in -band signalling. Other choices are at  the vendor's discretion.    Payload encapsulation delay Payload type  64 required (8 ms),   40 desired (5 ms)   NxDS0, no signalling, N = 1  32 (4 ms) NxDS0, no signalling, N = 2..4  8 (1 ms) NxDS0, no signalling, N > 4  24 (3 ms) NxDS0 with DS1 CAS  16 (2 ms) NxDS0 with E1 CAS  (R, W, set-by-create) (mandatory for structured service) (1 byte)  Timing mode: This attribute selects the timing mode of the TDM service. If RTP is used, this  attribute must be set to be consistent with the value of the RTP timestamp mode  attribute in the RTP pseudowire parameters ME, or its equivalent, at the far  end.  0 Network timing (default)  1 Differential timing  2 Adaptive timing  3 Loop timing: local TDM transmit clock derived from local TDM  receive stream  (R, W) (mandatory) (1 byte) 

---- page break ---- Transmit circuit ID : This attribute is a pair of emulated circuit ID  (ECID) values that the  ONU transmits in the direction from the TDM termination towards the packet- switched network (PSN). MEF 8 ECIDs lie in the range 1..1048575 (2 20 – 1).  To allow for the possibility of other transport (L2TP) in the future, each ECID  is allocated 4 bytes.  The first value is used for the payload ECID; the second is used for the optional  separate signalling ECID. The first ECID is required for all MEF 8  pseudowires; the second is required only if signalling is to be carried in a  distinct channel. If signalling is not present, or is carried in the same channel  as the payload, the second ECID should be set to 0.  (R, W) (mandatory for MEF 8 transport) (8 bytes)  Expected circuit ID: This attribute is a pair of ECID values that the ONU can expect in the  direction from the PSN towards the TDM termination. Checking ECIDs may  be a way to detect circuit misconnection. MEF 8 ECIDs lie in the range  1..1048575 (220 – 1). To allow for the possibility of other transport (L2TP) in  the future, each ECID is allocated 4 bytes.  The first value is used for the payload ECID; the second is used for the optional  separate signalling ECID. In both cases, the default value 0 indicates that no  ECID checking is expected.  (R, W) (optional for MEF 8 transport) (8 bytes)  Received circuit ID: This attribute indicates the actual ECID(s) received on the payload and  signalling channels, respectively. It may be used for diagnostic purposes. (R)  (optional for MEF 8 transport) (8 bytes)  Exception policy: This attribute points to an instance of the pseudowire maintenance profile  ME. If the pointer has its default value 0, the ONU 's internal defaults apply.  (R, W) (optional) (2 bytes)  ARC: See clause A.1.4.3. (R, W) (optional) (1 byte)  ARC interval: See clause A.1.4.3. (R, W) (optional) (1 byte)  Actions  Create, delete, get, set  Notifications  Attrib
```
