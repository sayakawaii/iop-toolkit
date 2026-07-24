# Managed Entity

## Identity
- ME ID: 401
- ME Name: Physical path termination point RS232/RS485 UNI
- Source Section: 9.15.1
- Source Page: 483

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Get, Set
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: Operational state
- Size: 1 byte
- Format: needs_review
- Access: R
- Category: optional

### Attribute 2
- Name: Port mode
- Size: 1 byte
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 3
- Name: Data_bits
- Size: 1 byte
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 4
- Name: Parity
- Size: 1 byte
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 5
- Name: Stop_bits
- Size: 1 byte
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 6
- Name: Flow_control
- Size: 1 byte
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 7
- Name: Min_send_time
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

## Extraction Status
- Status: auto_extracted
- Attributes found: 7
- Review needed: false

## Raw Source

```
9.15.1 Physical path termination point RS232/RS485 UNI  This ME represents an RS232/RS485 UNI in the ONU, where physical paths terminate and physical  path level functions are performed.  The ONU automatically creates an instance of this ME per port as follows.  • When the ONU has RS232/RS485 ports built into its factory configuration.  • When a cardholder is provisioned to expect a circuit pack of RS232/RS485 type.  • When a cardholder provisioned for plug and play is equipped with a circuit pack of  RS232/RS485 type. Note that the installation of a plug and play card may indicate the  presence of RS232/RS485 ports via equipment ID as well as its type, and indeed may cause  the ONU to instantiate a port-mapping package that specifies RS232/RS485 ports.  The ONU automatically deletes instances of this ME when a cardholder is neither provisioned to  expect a RS232/RS485 circuit pack, nor is equipped with a RS232/RS485 circuit pack.  Relationships  An instance of this ME is associated with each real RS232/RS485 port.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. This 2 byte  number is directly associated with the physical position of the UNI. The first  byte is the slot ID (defined in clause 9.1.5). The second byte is the port ID,  with range 1..255. (R) (mandatory) (2 bytes) 

---- page break ---- Administrative state: This attribute locks (1) and unlocks (0) the functions performed by this  ME. Administrative state is further described in clause A.1.6. (R,  W)  (mandatory) (1 byte)  Operational state : This attribute indicates whether the ME is capable of performing its  function. Valid values are enabled (0) and disabled (1). (R) (optional) (1 byte)  Port mode: This attribute indicates the working mode of the RS232/RS485 controller chipset.  Valid values are as follows.  0 half-duplex  1 full duplex  (mandatory) (1 byte)  Baud_rate: This attribute specifies the working baud rate of RS232/RS485 port. Valid values  are as follows.  0 300 bit/s  1 600 bit/s  2 1200 bit/s  3 2400 bit/s  4 4800 bit/s  5 9600 bit/s  6 19200 bit/s  7 38400 bit/s  8 43000 bit/s  9 56000 bit/s  10 57600 bit/s  11 115200 bit/s  (R, W, set-by-create) (mandatory) (1 byte)  Data_bits: This attribute specifies the bits of the data. Valid values are as follows.  5 5 bits  6 6 bits  7 7 bits  8 8 bits  (R, W, set-by-create) (mandatory) (1 byte)  Parity: This attribute specifies the parity of the data. Valid values are as follows.  0 no parity  1 odd parity  2 even parity  (R, W, set-by-create) (mandatory) (1 byte)  Stop_bits: This attribute specifies the number of stop bits of the data. Valid values are as  follows.  1 1 bit  2 2 bits  (R, W, set-by-create) (mandatory) (1 byte)  Flow_control: This attribute specifies the flow control of the data. Valid values are as  follows.  0 no flow control  1 hardware flow control (RTS/CTS)  2 software flow control (Xon/Xoff)  (R, W, set-by-create) (mandatory) (1 byte) 

---- page break ---- Min_send_payload: This attribute specifies the length of serial data acquisition  by  RS232/RS485 controller chipset in the fixed length mode. (R) (mandatory)  (4 bytes)  Min_send_time: This attribute specifies the time of serial data acquisition by RS232/RS485  controller chipset in the timing mode. (R) (mandatory) (4 bytes)  Reserve: This attribute is reserved for future use.  Actions  Get, set  
```
