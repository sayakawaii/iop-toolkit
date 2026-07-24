# Managed Entity

## Identity
- ME ID: 136
- ME Name: TCP/UDP config data
- Source Section: 9.4.3
- Source Page: 223

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Create, Delete, Get, Set
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: Port ID
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 2
- Name: Protocol
- Size: 1 byte
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 3
- Name: TOS/diffserv field
- Size: 1 byte
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

## Extraction Status
- Status: auto_extracted
- Attributes found: 3
- Review needed: false

## Raw Source

```
9.4.3 TCP/UDP config data  The TCP/UDP config data ME configures services based on the transmission control protocol (TCP)  and user datagram protocol (UDP) that are offered from an IP host. If a non-OMCI interface is used  to manage an IP service, this ME is unnecessary; the non-OMCI interface supplies the necessary data.  An instance of this ME is created and deleted on request of the OLT.  Relationships  One or more instances of this ME may be associated with an instance of an IP host config data  or IPv6 host config data ME.  Attributes   Managed entity  ID: This attribute uniquely identifies each instance of this ME. It is  recommended that the ME ID be the same as the port number. (R,  set-by-create) (mandatory) (2 bytes)  Port ID: This attribute specifies the port number that offers the TCP/UDP service.  (R, W, set-by-create) (mandatory) (2 bytes)  Protocol: This attribute specifies the protocol type as defined by [b-IANA] (protocol  numbers), for example UDP (0x11). (R,  W, set-by-create) (mandatory)  (1 byte)  TOS/diffserv field: This attribute specifies the value of the TOS/diffserv field of the IPv4  header. The contents of this attribute may contain the type of service per [IETF  RFC 2474] or a DSCP. Valid values for DSCP are as defined by [b-IANA]  (differentiated services field code points). (R,  W, set-by-create) (mandatory)  (1 byte) 

---- page break ---- IP host pointer: This attribute points to the IP host config data or IPv6 host config data ME  associated with this TCP/UDP data. Any number of ports and protocols may  be associated with an IP host. (R, W, set-by-create) (mandatory) (2 bytes)  Actions  Create, delete, get, set  Notifications  None.  
```
