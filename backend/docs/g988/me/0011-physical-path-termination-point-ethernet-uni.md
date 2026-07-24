# Managed Entity

## Identity
- ME ID: 11
- ME Name: Physical path termination point Ethernet UNI
- Source Section: 9.5.1
- Source Page: 230

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Get, Set
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: Expected type
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 2
- Name: Sensed type
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 3
- Name: Operational state
- Size: 1 byte
- Format: needs_review
- Access: R
- Category: optional

### Attribute 4
- Name: Configuration ind
- Size: 1 byte
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 5
- Name: Max frame size
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 6
- Name: Pause time
- Size: 2 bytes
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 7
- Name: Bridged or IP ind
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 8
- Name: ARC
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 9
- Name: ARC interval
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 10
- Name: PPPoE filter
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 11
- Name: Power control
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 12
- Name: Disable packet-based time synchronization
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: optional

## Extraction Status
- Status: auto_extracted
- Attributes found: 12
- Review needed: false

## Raw Source

```
9.5.1 Physical path termination point Ethernet UNI  This ME represents the point at an Ethernet UNI where the physical path terminates and Ethernet  physical level functions are performed.  The ONU automatically creates an instance of this ME per port:  • when the ONU has Ethernet ports built into its factory configuration;  • when a cardholder is provisioned to expect a circuit pack of the Ethernet type;  • when a cardholder provisioned for plug -and-play is equipped with a circuit pack of the  Ethernet type. Note that the installation of a plug-and-play card may indicate the presence of  Ethernet ports via equipment ID as well as its type, and indeed may cause the ONU to  instantiate a port-mapping package that specifies Ethernet ports. 

---- page break ---- The ONU automatically deletes instances of this ME when a cardholder is neither provisioned to  expect an Ethernet circuit pack, nor is it equipped with an Ethernet circuit pack.  Relationships  An instance of this ME is associated with each instance of a pre-provisioned or real Ethernet  port.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. This 2 byte  number indicates the physical position of the UNI. The first byte is the slot ID  (defined in clause 9.1.5). The second byte is the port ID, with the range 1..255.  (R) (mandatory) (2 bytes)  Expected type: This attribute supports pre-provisioning. It is coded as follows:  0 Autosense  1 to 254 One of the values from Table 9.1.5-1 that is compatible with an  Ethernet circuit pack  Upon ME instantiation, the ONU sets this attribute to 0. (R,  W) (mandatory)  (1 byte)  Sensed type: When a circuit pack is present, this attribute represents its type as one of the  values from Table 9.1.5-1. If the value of the expected type is not 0, then the  value of the sensed type should be the same as the value of the expected type.  Upon ME instantiation, the ONU sets this attribute to 0. See also the note in  the following AVC table.  (R) (mandatory if the ONU supports circuit packs with configurable interface  types, e.g., 10/100 BASE-T card) (1 byte)  Auto detection configuration: This attribute sets the following Ethernet port configuration.    Code point Rate Duplex  0x00 Auto Auto  0x01 10 Mbit/s only Full duplex only  0x02 100 Mbit/s only Full duplex only  0x03 1000 Mbit/s  only  Full duplex only  0x04 Auto Full duplex only  0x05 10Gb/s only Full duplex only  0x06 2.5Gb/s only Full duplex only  0x07 5Gb/s only Full duplex only  0x08 25Gb/s only Full duplex only  0x09 40Gb/s only Full duplex only  0x10 10 Mbit/s only Auto  0x11 10 Mbit/s only Half duplex only  0x12 100 Mbit/s only Half duplex only  0x13 1000 Mbit/s  only  Half duplex only  0x14 Auto Half duplex only 

---- page break ---- Code point Rate Duplex  0x20 1000 Mbit/s  only  Auto  0x30 100 Mbit/s only Auto  Upon ME instantiation, the ONU sets this attribute to 0. (R, W) (mandatory for  interfaces with autodetection options) (1 byte)  Ethernet loopback configuration : This attribute sets the following Ethernet loopback  configuration.  0 No loopback  3 Loop 3, loopback of downstream traffic after PHY transceiver. Loop 3  is depicted in Figure 9.5.1-1.  Note that normal bridge behaviour may defeat the loopback signal unless  broadcast MAC addresses are used. Although it does not reach the physical  interface, [IEEE 802.1ag] loopback is preferred.  Upon ME instantiation, the ONU sets this attribute to 0. (R,  W) (mandatory)  (1 byte)  G.988(12)_F9.5.1-1 ONU PHY transceiver PON Ethernet UNI Loopback 3   Figure 9.5.1-1 – Ethernet loopback configuration  Administrative state: This attribute locks (1) and unlocks (0) the functions performed by this  ME. Administrative state is further described in clause A.1.6. (R,  W)  (mandatory) (1 byte)  Operational state : This attribute indicates whether the ME is capable of performing its  function. Valid values are enabled (0) and disabled (1). (R) (optional) (1 byte)  Configuration ind: This attribute indicates the configuration status of the Ethernet UNI.  0x01 10BASE-T full duplex  0x02 100BASE-T full duplex  0x03 Gigabit Ethernet full duplex  0x04 10Gb/s Ethernet full duplex  0x05 2.5Gb/s Ethernet full duplex  0x06 5Gb/s Ethernet full duplex  0x07 25Gb/s Ethernet full duplex  0x08 40Gb/s Ethernet full duplex  0x11 10BASE-T half duplex  0x12 100BASE-T half duplex  0x13 Gigabit Ethernet half duplex  The value 0 indicates that the configuration status is unknown (e.g., Ethernet  link is not established or the circuit pack is not yet installed). Upon ME  instantiation, the ONU sets this attribute to 0. (R) (mandatory) (1 byte)  Max frame size: This attribute denotes the maximum frame size allowed across this interface.  Upon ME instantiation, the ONU sets the attribute to 1518. (R, W) (mandatory  for G-PON, optional for ITU-T G.986 systems) (2 bytes) 

---- page break ---- DTE or DCE ind: This attribute specifies the following Ethernet interface wiring.  0 DCE or MDI-X (default).  1 DTE or MDI.  2 Automatic selection  (R, W) (mandatory) (1 byte)  Pause time: This attribute allows the PPTP to ask the subscriber terminal to temporarily  suspend sending data. Units are in pause quanta (1 pause quantum is 512 bit  times of the particular implementation). Values: 0..0xFFFF. Upon ME  instantiation, the ONU sets this attribute to 0. (R, W) (optional) (2 bytes)  Bridged or IP ind : This attribute specifies whether the Ethernet interface is bridged or  derived from an IP router function.  0 Bridged  1 IP router  2 Depends on the parent circuit pack. 2 means that the circuit pack 's  bridged or IP ind attribute is either 0 or 1.  Upon ME instantiation, the ONU sets this attribute to 2. (R,  W) (optional)  (1 byte)  ARC: See clause A.1.4.3. (R, W) (optional) (1 byte)  ARC interval: See clause A.1.4.3. (R, W) (optional) (1 byte)  PPPoE filter: This attribute controls filtering of PPPoE  packets on this Ethernet port. The  value 0 allows packets of all types. The
```
