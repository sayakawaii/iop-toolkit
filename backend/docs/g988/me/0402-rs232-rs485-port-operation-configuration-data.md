# Managed Entity

## Identity
- ME ID: 402
- ME Name: RS232/RS485 port operation configuration data
- Source Section: 9.15.2
- Source Page: 485

## Classification
- Access: needs_review
- Type: needs_review
- Actions: needs_review
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: Socket mode
- Size: 1 byte
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 2
- Name: TCP/UDP pointer
- Size: 2 bytes
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 3
- Name: PPTP pointer
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

## Extraction Status
- Status: auto_extracted
- Attributes found: 3
- Review needed: true

## Raw Source

```
9.15.2 RS232/RS485 port operation configuration data  This ME specifies the RS232/RS485 port operation mode. The ONU automatically creates instances  of this ME if RS232/RS485 data acquisition services are available.  Relationships  An instance of this ME is associated with a TCP/UDP config data ME and a PPTP  RS232/RS485 UNI.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. (R, set-by- create) (mandatory) (2 bytes)  Socket mode: This attribute identifies the RS232/RS485 port operation mode as follows:  0x00 TCP/server mode;  0x01 TCP/client mode;  0x02 UDP mode,  Other values are reserved.  (R) (mandatory) (1 byte)  TCP/UDP pointer: This pointer associates the RS232/RS485 port operation  configuration  with the TCP/UDP config data ME to be used for communication with the  serial server. The default value is 0xFFFF, a null pointer. (R, W) (mandatory)  (2 bytes)  PPTP pointer: This attribute points to the PPTP RS232/RS485 UNI ME that serves the serial  data acquisition function. (R, W, set-by-create) (mandatory) (2 bytes)  
```
