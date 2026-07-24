# Managed Entity

## Identity
- ME ID: 407
- ME Name: SIP agent config data 2
- Source Section: 9.9.3
- Source Page: 388

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Create, Delete, Get, Set
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: Proxy server address pointer
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 2
- Name: Outbound proxy address pointer
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 3
- Name: Primary SIP DNS
- Size: 4 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 4
- Name: Secondary SIP DNS
- Size: 4 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 5
- Name: SIP reg exp time
- Size: 4 bytes
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 6
- Name: SIP rereg head start time
- Size: 4 bytes
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 7
- Name: Host part URI
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 8
- Name: SIP status
- Size: 1 byte
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 9
- Name: SIP registrar
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 10
- Name: Softswitch
- Size: 4 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 11
- Name: SIP response table
- Size: 1 byte
- Format: needs_review
- Access: R, W, set-by-create
- Category: optional

### Attribute 12
- Name: SIP URI format
- Size: 1 byte
- Format: needs_review
- Access: R, W, set-by-create
- Category: optional

### Attribute 13
- Name: Redundant SIP agent pointer
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: optional

## Extraction Status
- Status: auto_extracted
- Attributes found: 13
- Review needed: false

## Raw Source

```
9.9.3 SIP agent config data  The SIP agent config data ME models a SIP signalling agent. It defines the configuration necessary  to establish communication for signalling between the SIP user agent (UA) and a SIP server.  NOTE 1 – If a non-OMCI interface is used to manage SIP for VoIP, this ME is unnecessary. The non -OMCI  interface supplies the necessary data, which may be read back to the OLT via the SIP config portal ME.  Instances of this ME are created and deleted by the OLT.  Relationships  An instance of this ME serves one or more SIP user data MEs and points to a TCP/UDP config  data that carries signalling messages. Other pointers establish additional agent parameters  such as proxy servers.  Attributes  Managed entity  ID: This attribute uniquely identifies each instance of this ME. (R ,  set-by-create) (mandatory) (2 bytes)  Proxy server address pointer : This attribute points to a large string ME that contains the  name (IP address or URI) of the SIP proxy server for SIP signalling messages.  (R, W, set-by-create) (mandatory) (2 bytes)  Outbound proxy address pointer : An outbound SIP proxy may or may not be required  within a given network. If an outbound SIP proxy is used, the outbound proxy  address pointer attribute must be set to point to a valid large string ME  that  contains the name (IP address or URI) of the outbound proxy server for SIP  signalling messages. If an outbound SIP proxy is not used, the outbound proxy  address pointer attribute must be set to a null pointer.  (R, W, set-by-create)  (mandatory) (2 bytes)  Primary SIP DNS: This attribute specifies the primary SIP DNS IP address. If the value of  this attribute is 0, the primary DNS server is defined in the corresponding IP  host config data or IPv6 host config data ME. If the value is non-zero, it takes  precedence over the primary DNS server defined in the IP host config data or  IPv6 host config data ME. (R, W, set-by-create) (mandatory) (4 bytes)  Secondary SIP DNS: This attribute specifies the secondary SIP DNS IP address. If the value  of this attribute is 0, the secondary DNS server is defined in the corresponding  IP host config data or IPv6 host config data ME. If the value is non -zero, it  takes precedence over the secondary DNS server defined in the IP host config  data or IPv6 host config data ME. (R, W, set-by-create) (mandatory) (4 bytes) 

---- page break ---- TCP/UDP pointer: This pointer associates the SIP agent with the TCP/UDP config data ME  to be used for communication with the SIP server. The default value is 0xFFFF,  a null pointer. (R, W) (mandatory) (2 bytes)  SIP reg exp time: This attribute specifies the SIP registration expiration time in seconds. If  its value is 0, the SIP agent does not add an expiration time to the registration  requests and does not perform re -registration. The default value is 3600  s.  (R, W) (mandatory) (4 bytes)  SIP rereg head start time : This attribute specifies the time in seconds prior to timeout that  causes the SIP agent to start the re -registration process. The default value is  360 s. (R, W) (mandatory) (4 bytes)  Host part URI : This attribute points to a large string ME that contains the host or domain  part of the SIP address of record for users connected to this ONU. A null  pointer indicates that the current address in the IP host config ME is to be used.  (R, W, set-by-create) (mandatory) (2 bytes)  SIP status: This attribute shows the current status of the SIP agent. Values are as follows.  0 Ok/initial  1 Connected  2 Failed – ICMP error  3 Failed – Malformed response  4 Failed – Inadequate info response  5 Failed – Timeout  6 Redundant, offline: this instance of the SIP agent config data occupies  the role of a redundant server, and is not presently in use.  (R) (mandatory) (1 byte)  SIP registrar : This attribute points to a network address ME that contains the name (IP  address or resolved name) of the registrar server for SIP signalling messages.  Examples: "10.10.10.10" and "proxy.voip.net". (R, W, set-by-create)  (mandatory) (2 bytes)  Softswitch: This attribute identifies the SIP gateway softswitch vendor. The format is four  ASCII coded alphabetic characters [A..Z] as defined in [ATIS -0300220]. A  value of four null bytes indicates an unknown or unspecified vendor. (R,  W,  set-by-create) (mandatory) (4 bytes)  SIP response table: This attribute specifies the tone and text to be presented to the subscriber  upon receipt of various SIP messages (normally 4xx, 5xx, 6xx message codes).  The table is a sequence of entries, each of which is defined as follows.  SIP response code (2 bytes): This field is the value of the SIP message code.  It also serves as the index into the SIP response table. When a set operation  is performed with the value 0 in this field, the table is cleared.  Tone (1 byte): This field specifies one of the tones in the tone pattern table of  the associated voice service profile. The specified tone is played to the  subscriber.  Text message (2 bytes): This field is a pointer to a large string that contains a  message to be displayed to the subscriber. If the value of this field is a null  pointer, text pre -associated with the tone may be displayed, or no text at  all.  (R, W) (optional) (N * 5 bytes) 

---- page break ---- NOTE 2 – This model assumes that SIP response tones and text are common to all  POTS lines that share a given SIP agent.  SIP option transmit control: This Boolean attribute specifies that the ONU is (true) or is not  (false) enabled to transmit SIP options. The default value is recommended to  be false. (R, W, set-by-create) (optional) (1 byte)  SIP URI format : This attribute specifies the format of the URI in outgoing SIP messages.  The recommended default value 0 specifies TEL URIs; the value 1 specifies  SIP URIs. Other values are reserved. (R, W, set-by-create) (optional) (1 byte)  Redundant SIP agent pointer : This attribute points to another SIP agent config data ME,  which is 
```
