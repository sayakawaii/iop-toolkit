# Managed Entity

## Identity
- ME ID: 336
- ME Name: ONU dynamic power management control
- Source Section: 9.1.14
- Source Page: 90

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Get, Set
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: Power reduction management capability
- Size: 1 byte
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 2
- Name: Codepoints are assigned as follows
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 3
- Name: Itransinit
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 4
- Name: Maximum sleep interval
- Size: 4 bytes
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 5
- Name: Maximum receiver-off interval
- Size: 4 bytes
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 6
- Name: Minimum aware interval
- Size: 4 bytes
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 7
- Name: Minimum active held interval
- Size: 2 bytes
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 8
- Name: Maximum sleep interval extension
- Size: 8 bytes
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 9
- Name: Missing consecutive bursts threshold
- Size: 4 bytes
- Format: needs_review
- Access: R, W
- Category: mandatory

## Extraction Status
- Status: auto_extracted
- Attributes found: 9
- Review needed: false

## Raw Source

```
9.1.14 ONU dynamic power management control  This ME models the ONU's ability to enter power conservation modes in cooperation with the OLT  in an ITU -T G.987 system. [ITU -T G.987.3] originally specified two alternative modes, doze and  cyclic sleep. The subsequent revision of [ITU -T G.987.3] simplified the s pecification providing a  single power conservation mode, watchful sleep.  An ONU that supports power conservation modes automatically creates an instance of this ME.  Relationships  One instance of this ME is associated with the ONU ME.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. There is only  one instance, number 0. (R) (mandatory) (2 bytes)  Power reduction management capability : This attribute declares the ONU's support for  managed power conservation modes, as defined in [ITU-T G.987.3]. It is a bit  map in which the bit value 0 indicates no support for the specified mode, while  the bit value 1 indicates that the ONU does support the specified mo de. (R)  (mandatory) (1 byte)  Codepoints are assigned as follows:  Value Meaning  0 No support for power reduction  1 Doze mode supported  2 Cyclic sleep mode supported  3 Both doze and cyclic sleep modes supported  4 Watchful sleep mode supported  5..255 Reserved  Power reduction management mode : This attribute enables one or more of the ONU's  managed power conservation modes. It is a bit map in which the bit value 0  disables the mode, while the value 1 enables the mode. Bit assignments are the  same as those of the power reduction management capa bility attribute. The  default value of each bit is 0. (R, W) (mandatory) (1 byte)  Itransinit: This attribute is the ONU vendor's statement of the complete transceiver  initialization time: the worst -case time required for the ONU to regain full  functionality when leaving the asleep state in cyclic sleep mode or low-power  state in watchful sleep mode (i.e., turning on both the receiver and the  transmitter and acquiring synchronization to the downstream flow), measured  in units of 125 µs frames. The value zero indicates that the sleeping ONU can  respond to a bandwidth grant without delay. (R) (mandatory) (2 bytes) 

---- page break ---- Itxinit: This attribute is the ONU vendor's statement of the transmitter initialization  time: the time required for the ONU to regain full functionality when leaving  the listen state (i.e., turning on the transmitter), measured in units of 125  µs  frames. The value zero indicates that the dozing ONU can respond to a  bandwidth grant without delay. If watchful sleep is enabled, the ONU ignores  this attribute. (R) (mandatory) (2 bytes)  Maximum sleep interval : The Isleep/Ilowpower attribute specifies the maximum time the  ONU spends in its asleep, listen, or low -power states, as a count of 125  µs  frames. Local or remote events may truncate the ONU's sojourn in these states.  The default value of this attribute is 0. (R, W) (mandatory) (4 bytes)  Maximum receiver-off interval: The Irxoff attribute specifies the maximum time the OLT  can afford to wait from the moment it decides to wake up an ONU in the low- power state of the watchful sleep mode until the ONU is fully operational,  specified as a count of 125 µs frames. (R, W) (mandatory) (4 bytes)  Minimum aware interval : The Iaware attribute specifies the time the ONU spends in its  aware state, as a count of 125  µs frames, before it re -enters asleep or listen  states. Local or remote events may independently cause the ONU to enter an  active state rather than returning to a sleep state. The default value of this  attribute is 0. (R, W) (mandatory) (4 bytes)  Minimum active held interval: The Ihold attribute specifies the minimum time during which  the ONU remains in the active held state, as a count of 125 µs frames. Its initial  value is zero. (R, W) (mandatory) (2 bytes)  Maximum sleep interval extension: This attribute designates maximum sleep interval values  for doze mode and cyclic sleep mode separately. When it supports this attribute,  the ONU ignores the value of the maximum sleep interval attribute.  Maximum sleep interval for doze mode 4 bytes  Maximum sleep interval for cyclic sleep mode 4 bytes  Maximum sleep interval for doze mode specifies the maximum time the ONU  spends in its listen state, as a count of 125  µs frames. Local or remote events  may truncate the ONU's sojourn in these states. The default value is 0.  Maximum sleep interval for cyclic sleep mode specifies the maximum time the  ONU spends in its asleep state, as a count of 125  µs frames. Local or remote  events may truncate the ONU's sojourn in these states. The default value is 0.  If watchful sleep is enabled, the ONU ignores this attribute.  (R, W) (optional) (8 bytes)  Ethernet passive optical network (EPON) capability extension: This attribute declares  EPON-specific capabilities for the dynamic power management control.  Bits are assigned as follows.  Bit Name Setting  1 (LSB) AckCapable 0: not supported   1: supported  2 Sleep indication capability  0: not supported   1: supported  3 Early wake-up capability  0: not supported   1: supported  4 Sleep mode selection at ONU's discretion   0: not supported   1: supported 

---- page break ---- 5..8 Reserved 0    AckCapable has the value of supported if the ONU is capable of sending a  SLEEP_ACK message, which is defined in [IEEE P1904.1], in response to the  SLEEP_ALLOW message from the OLT. The ONU may select the  appropriate power conservation method by itself if AckCapable is supported.  Sleep indication capability represents ability to send a SLEEP_INDICATION  message, defined in [IEEE P1904.1], to initiate the power saving cycle from  the ONU.  Early wake-up capability shows whether the ONU has a function in which the  ONU can awaken from the sleep mode based on local conditions such as off - hook condition on SIP ports and power down.  ONU self-sleep mode selection indicates whether the ONU has 
```
