# Managed Entity

## Identity
- ME ID: 465
- ME Name: Precision Time Protocol
- Source Section: 9.12.22
- Source Page: 448

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Get, Set
- Instance Type: singleton
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: Administrative state
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 2
- Name: Network transport protocol
- Size: 2 bytes
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 3
- Name: Multicast address mode
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 4
- Name: PTP source address
- Size: 4 bytes
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 5
- Name: Downstream Packet TCI
- Size: 3 bytes
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 6
- Name: The first byte defines the control type
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 7
- Name: Log a nnounce interval
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 8
- Name: ToD update interval
- Size: 2 byte
- Format: needs_review
- Access: R, W
- Category: mandatory

## Extraction Status
- Status: auto_extracted
- Attributes found: 8
- Review needed: false

## Raw Source

```
9.12.22 Precision Time Protocol  This ME configures the ONU's ability to control the Precision Time Protocol (PTP) capability of the  ONU as defined in IEEE 1588v2.  An ONU that supports PTP automatically creates an instance of this ME.  Relationships  The single instance of this ME is associated with the ONU ME.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. There is only  one instance, number 0. (R) (mandatory) (2 bytes)  Administrative state: This attribute locks (1) and unlocks (0) the functions performed by this  ME. Administrative state is further described in clause A.1.6. (R,  W)  (mandatory) The default value is 1. (1 byte)  Communication model: This attribute specifies the  communication model of PTP packets  encapsulation. See clause 6.2 in [IEEE 1588v2]. Its values are defined as  follows:  0 multicast  1 unicast  The default value is 0. (R, W) (mandatory) (1 byte)  Network transport protocol : This attribute specifies the network protocol of PTP packets  encapsulation. See clause 7.4.1 Network transport protocol in [IEEE 1588v2].  Its values are defined as follows:  0x0000 Reserved  0x0001 UDP/IPv4  0x0002 UDP/IPv6  0x0003 IEEE 802.3  All other values are reserved.  The default value is 0x0003. (R, W) (mandatory) (2 bytes)  Multicast address mode: This attribute specifies the transmitted multicast destination MAC  address when transmitting PTP messages to a remote PTP port (when  IEEE 802.3 Network Transport protocol is selected). Two multicast addresses  are supported to handle operator preferences of whether to forward PTP  messages in a PTP unaware network. See [IEEE 1588v2] clause F.3 and  [ITU-T G.8275.1] Appendix III for more details. The following values are  supported:  0 Non-forwardable multicast address  (multicast address 01 -80-C2-00- 00-0E) will be used on transmitted PTP message  

---- page break ---- 1 Forwardable multicast address (multicast address 01-1B-19-00-00-00)  will be used on transmitted PTP message   All other values are reserved.  The default value is 0. (R, W) (mandatory) (1 byte)  PTP source address: This attribute specifies the source address for PTP packets. If the value  is 0.0.x.y, where x and y are not both 0, then x.y is to be interpreted as a pointer  to a large string ME that represents an IPv6 address. Otherwise, the address is  an IPv4 address.  The default value is 0. (R, W) (mandatory) (4 bytes)  Downstream Packet TCI: This attribute specifies the VLAN of PTP packets encapsulation.  Default value is 0. (R, W) (mandatory) (3 bytes)   The first byte defines the control type:   0 PTP packets are untagged. The remaining 2 bytes of the attribute are  ignored.  1 PTP packets are tagged with the TCI specified in bytes 2 and 3 of this  attribute.  The second and third bytes specifies the TCI (VLAN ID and P -bits) to be  applied to the PTP packet if the control byte indicates a tagged format.  Clock step: This attribute specifies the ONU clock is one-step clock or two-step clock. Default  value is 0.   0 one-step  1 two-step  (R, W) (mandatory) (1 byte)  Log a nnounce interval : This attribute specifies the base -2 logarithm of the mean time  interval in seconds between successive ONU Announce messages (see  clause 7.7.2.2 in [IEEE 1588v2]). The format is a signed integer. Default value  is 0 (implement specific). (Refer to clause 8.2.5.4.1 of [IEEE 1588v2]).  (R, W) (mandatory)) (1 byte)  Log s ync interval: This attribute specifies base -2 logarithm of the mean SyncInterval in  seconds for multicast messages. T he rates for unicast transmissions are  negotiated separately on a per -port basis and are not constrained by this  attribute(see clause 7.7.2.3 in [IEEE 1588v2]). The format is a signed integer.  Default value is 0. (Refer to clause 8.2.5.4.3 of [IEEE 1588v2]).  (R, W) (mandatory) (1 byte)   ToD update interval: This attribute specifies the expected time of day (ToD) update interval  in minutes. This interval is used for determining whether the ONU clock state  enters holdover mode (see clause 9.12.23). The format is a non-signed integed.  Valid values are 0 to 1440 minutes, as recommended in clause 10.4.6.2 of  [ITU-T G.984.3] for ToD refresh interval. Default value is 0.  (R, W) (mandatory) (2 byte)  Actions  Get, set  Notifications  None 

---- page break ---- 
```
