# Managed Entity

## Identity
- ME ID: 162
- ME Name: Physical path termination point MoCA UNI
- Source Section: 9.10.1
- Source Page: 419

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Get, Set
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: Loopback configuration
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 2
- Name: Operational state
- Size: 1 byte
- Format: needs_review
- Access: R
- Category: optional

### Attribute 3
- Name: Max frame size
- Size: 2 bytes
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 4
- Name: ARC
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 5
- Name: ARC interval
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 6
- Name: PPPoE filter
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 7
- Name: Network status
- Size: 1 byte
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 8
- Name: Password
- Size: 17 bytes
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 9
- Name: Privacy enabled
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 10
- Name: Minimum bandwidth alarm threshold
- Size: 2 bytes
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 11
- Name: Frequency mask
- Size: 4 bytes
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 12
- Name: RF channel
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 13
- Name: Last operational frequency
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: mandatory

## Extraction Status
- Status: auto_extracted
- Attributes found: 13
- Review needed: false

## Raw Source

```
9.10.1 Physical path termination point MoCA UNI  This ME represents an MoCA UNI, where physical paths terminate and physical path level functions  are performed.  The ONU automatically creates an instance of this ME per port as follows.  • When the ONU has MoCA ports built into its factory configuration.  • When a cardholder is provisioned to expect a circuit pack of the MoCA type.  • When a cardholder provisioned for plug -and-play is equipped with a circuit pack of the  MoCA type. Note that the installation of a plug -and-play card may indicate the presence of  MoCA ports via equipment ID as well as its type, and indeed may cause the ONU to  instantiate a port-mapping package that specifies MoCA ports.  The ONU automatically deletes instances of this ME when a cardholder is neither provisioned to  expect an MoCA circuit pack, nor is it equipped with an MoCA circuit pack.  Relationships  An instance of this ME is associated with each real or pre-provisioned MoCA port.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. This 2 byte  number is directly associated with the physical position of the UNI. The first  byte is the slot ID (defined in clause 9.1.5). The second byte is the port ID,  with the range 1..255. (R) (mandatory) (2 bytes)  Loopback configuration: This attribute sets the MoCA loopback configuration. Note that  normal bridge behaviour may defeat the loopback signal unless broadcast  MAC addresses are used.  0 No loopback  3 Loopback3, loopback of downstream traffic after PHY transceiver,  depicted in Figure 9.10.1-1.  Upon ME instantiation, the ONU sets this attribute to 0. (R,  W) (optional)  (1 byte)  G.988(12)_F9.10.1-1 PHY transceiver ONU MoCA UNI PON Loopback 3   Figure 9.10.1-1 – MoCA loopback configuration 

---- page break ---- Administrative state: This attribute locks (1) and unlocks (0) the functions performed by this  ME. Administrative state is further described in clause A.1.6. (R,  W)  (mandatory) (1 byte)  Operational state : This attribute indicates whether the ME is capable of performing  its  function. Valid values are enabled (0) and disabled (1). (R) (optional) (1 byte)  Max frame size: This attribute denotes the maximum frame size allowed across this interface.  Upon ME instantiation, the ONU sets this attribute to 1518. (R,  W)  (mandatory) (2 bytes)  ARC: See clause A.1.4.3. (R, W) (optional) (1 byte)  ARC interval: See clause A.1.4.3. (R, W) (optional) (1 byte)  PPPoE filter: This attribute controls filtering of PPPoE packets on this MoCA port. When its  value is 1, all packets other than PPPoE packets are discarded. The default 0  accepts packets of all types. (R, W) (optional) (1 byte)  Network status : This attribute indicates the networking state of the MoCA interface  as  follows.  0 The interface has not joined an MoCA network.  1 The interface has joined an MoCA network.  2 The interface has joined a n MoCA network and is currently the  network coordinator.  (R) (mandatory) (1 byte)  Password: This attribute specifies the MoCA encryption key. It is an ASCII string of 17  decimal digits. Upon ME instantiation, the ONU sets this attribute to 17 null  bytes. (R, W) (mandatory) (17 bytes)  Privacy enabled : This attribute activates (1) link -layer security. The default value 0  deactivates it. (R, W) (mandatory) (1 byte)  Minimum bandwidth alarm threshold: This attribute specifies the minimum desired PHY  link bandwidth between two nodes. If the actual bandwidth is lower, a n LL  alarm is declared. Valid values are 0 to 0x0410 (260  Mbit/s) in 0.25  Mbit/s  increments. The default value is 0x02D0 (180  Mbit/s). The value 0 disables  the threshold. (R, W) (optional) (2 bytes)  Frequency mask: This attribute is a bit map of the centre frequencies that the interface is  permitted to use, where each bit represents a centre frequency. The LSB (b[1])  corresponds to centre frequency 800  MHz. The next significant bit ( b[2])  corresponds to centre frequency 825 MHz. The 28th bit (b[28]) corresponds to  centre frequency 1500  MHz. The four MSBs are not used. (R,  W) (optional)  (4 bytes)  RF channel: This attribute reports the frequency to which the MoCA interface is currently  tuned, in megahertz. (R) (mandatory) (2 bytes)  Last operational frequency : This attribute reports the frequency to which the MoCA  interface was tuned when last operational, in  megahertz. (R) (mandatory)  (2 bytes)  Actions  Get, set
```
