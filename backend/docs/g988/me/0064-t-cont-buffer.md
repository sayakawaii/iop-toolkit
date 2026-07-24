# Managed Entity

## Identity
- ME ID: 64
- ME Name: T-CONT buffer
- Source Section: needs_review
- Source Page: 78

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Get, Set, Create, Delete
- Instance Type: needs_review
- Instance Value: needs_review

## Attributes

### Attribute 1
- Name: Vendor ID
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 2
- Name: Administrative state
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
- Name: Bridged or IP ind
- Size: 20 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 5
- Name: Card configuration
- Size: 4 bytes
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 6
- Name: Actions  Get, set  Create, delete
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 7
- Name: Restore power timer reset interval
- Size: 2 bytes
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 8
- Name: Data class shedding interval
- Size: 2 bytes
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 9
- Name: Voice class shedding interval
- Size: 2 bytes
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 10
- Name: Video overlay class shedding interval
- Size: 2 bytes
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 11
- Name: Video return class shedding interval
- Size: 2 bytes
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 12
- Name: ATM class shedding interval
- Size: 2 bytes
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 13
- Name: CES class shedding interval
- Size: 2 bytes
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 14
- Name: Frame class shedding interval
- Size: 2 bytes
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 15
- Name: Sdh-sonet class shedding interval
- Size: 2 bytes
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 16
- Name: Shedding status
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 17
- Name: Max ports
- Size: 1 byte
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 18
- Name: Port list 1
- Size: 16 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 19
- Name: Port list 2
- Size: 16 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 20
- Name: Port list 3
- Size: 16 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 21
- Name: Port list 4
- Size: 16 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 22
- Name: Port list 5
- Size: 16 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 23
- Name: Port list 6
- Size: 16 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 24
- Name: Port list 7
- Size: 16 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 25
- Name: Port list 8
- Size: 16 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 26
- Name: Combined port table
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 27
- Name: Environmental sense
- Size: 2 bytes
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 28
- Name: For R'/S' protection, the protection type must be 1
- Size: 2 bytes
- Format: needs_review
- Access: R, set-by-create if applicable
- Category: mandatory

### Attribute 29
- Name: Working ANI-G pointer
- Size: 2 bytes
- Format: needs_review
- Access: R, W if applicable, set-by-create if applicable
- Category: mandatory

### Attribute 30
- Name: Protection ANI-G pointer
- Size: 2 bytes
- Format: needs_review
- Access: R, W if applicable, set-by-create if applicable
- Category: mandatory

### Attribute 31
- Name: Protection type
- Size: 1 byte
- Format: needs_review
- Access: R, W if applicable, set-by-create if applicable
- Category: mandatory

### Attribute 32
- Name: Revertive ind
- Size: 1 byte
- Format: needs_review
- Access: R, W if applicable, set-by-create if applicable
- Category: mandatory

### Attribute 33
- Name: Wait to restore time
- Size: 2 bytes
- Format: needs_review
- Access: RWSC if applicable
- Category: mandatory

### Attribute 34
- Name: Switching guard time
- Size: 2 bytes
- Format: needs_review
- Access: RWS C if applicable
- Category: optional

## Extraction Status
- Status: auto_extracted
- Attributes found: 34
- Review needed: true

## Raw Source

```
Vendor ID: This attribute identifies the vendor of the circuit pack. For ONUs with  integrated interfaces, this value is identical to the value of the vendor ID  attribute of the ONU-G ME. Upon creation in the absence of a physical circuit  pack, this attribute comprises all spaces. (R) (optional) (4 bytes)  Administrative state: This attribute locks (1) and unlocks (0) the functions performed by this  ME. Administrative state is further described in clause A.1.6. (R,  W)  (mandatory) (1 byte)  Operational state: This attribute indicates whether the circuit pack is capable of performing  its function. Valid values are enabled (0), disabled (1) and unknown (2).  Pending completion of initialization and self -test on an installed circuit pack,  the ONU sets this attribute to 2. (R) (optional) (1 byte)  Bridged or IP ind: This attribute specifies whether an Ethernet interface is bridged or derived  from an IP router function.  0 Bridged  1 IP router  2 Both bridged and IP router functions  (R, W) (optional, only applicable for circuit packs with Ethernet interfaces)  (1 byte)  Equipment ID: This attribute may be used to identify the vendor 's specific type of circuit  pack. In some environments, this attribute may include the CLEI code. Upon  ME instantiation, the ONU sets this attribute to all spaces or to the equipment  ID of the circuit pack that is physically present. (R) (optional) (20 bytes)  Card configuration : This attribute selects the appropriate configuration of configurable  circuit packs. Table 9.1.5-1 specifies two configurable card types: C -DS1/E1  (code 16), and C -DS1/E1/J1 (code 17). Values are indicated below for the  allowed card types and configurations.  Card Type Configuration Value  C-DS1/E1 DS1 0   E1 1  C-DS1/E1/J1 DS1 0   E1 1   J1 2  Upon autonomous instantiation, this attribute is set to 0. (R,  W, set-by-create  if applicable) (mandatory for configurable circuit packs) (1 byte)  Total T-CONT buffer number: This attribute reports the total number of T -CONT buffers  associated with the circuit pack. Upon ME instantiation, the ONU sets this  attribute to 0 or to the value supported by the physical circuit pack. (R)  (mandatory for circuit packs that provide a traffic scheduler function) (1 byte)  Total priority queue number: This value reports the total number of priority queues  associated with the circuit pack. Upon ME instantiation, the ONU sets the  attribute to 0 or to the value supported by the physical circuit pack. (R)  (mandatory for circuit packs that provide a traffic scheduler function) (1 byte) 

---- page break ---- Total traffic scheduler number: This value reports the total number of traffic schedulers  associated with the circuit pack. The ONU supports null function, strict priority  scheduling and WRR from the priority control, and guarantee of minimum rate  control points of view. If the circuit pack has no traffic scheduler, this attribute  should be absent or have the value 0. Upon ME instantiation, the ONU sets the  attribute to 0 or to the value supported by the physical circuit pack. (R)  (mandatory for circuit packs that provide a traffic scheduler function) (1 byte)  Power shed override: This attribute allows ports to be excluded from the power shed control  defined in clause 9.1.7. It is a bit mask that takes port 1 as the MSB; a bit value  of 1 marks the corresponding port to override the power shed timer. For  hardware that cannot shed power per port, this attribute is a slot override rather  than a port override, with any non -zero port value causing the entire circuit  pack to override power shedding. (R, W) (optional) (4 bytes)  Actions  Get, set  Create, delete: Optional, only when plug-and-play is supported.  Reboot: Reboot the circuit pack.  Test: Test the circuit pack (optional). The test action may be used either to perform  equipment diagnostics or to measure parameters , such as received optical  power, video output level and battery voltage. Test and test result messages are  defined in Annex A.  Notifications  Attribute value change  Number Attribute value change Description  1..6 N/A   7 Op state Operational state change  8..14 N/A   15..16 Reserved     Alarm  Alarm  number  Alarm Description  0 Equipment alarm A failure on an internal interface or failed self-test  1 Powering alarm Fuse failure or failure of DC/DC converter  2 Self-test failure Failure of circuit pack autonomous self-test  3 Laser end of life Failure of transmit laser imminent  4 Temperature yellow No service shutdown at present, but the circuit pack is  operating beyond its recommended range.  5 Temperature red Service has been shut down to avoid equipment damage.  The operational state of the affected PPTPs indicates the  affected services.  6..207 Reserved   208..223 Vendor-specific alarms Not to be standardized 

---- page break ---- 9.1.7 ONU power shedding  This ME models the ONU's ability to shed services when the ONU goes into battery operation mode  after AC power failure. Shedding classes are defined in the following table, which may span multiple  circuit pack types. This feature works in conjunction with the power shed override attribute of the  circuit pack ME, which can selectively prevent power shedding of priority ports.  An ONU that supports power shedding automatically creates an instance of this ME.  The following table defines the binding of shedding class and PPTP type. The coding is taken from  Table 9.1.5-1. In the case of hybrid circuit pack types, multiple shedding classes may affect a circuit  pack if the hardware is capable of partial power shedding.  An ONU may choose to model its ports with the port -mapping package of clause 9.1.8, rather than  with real or virtual circuit packs. In this case, power shedding pertains to individual PPTPs (listed in  column 2 of the table).   

---- page break ---- Shedding class PPTP type Coding Content  ATM ATM PPTP 1..12 Various ATM UNIs  CES CES PPTP 13 C1.5 (DS1)  14 C2.0 (E1)  15 C6.3 
```
