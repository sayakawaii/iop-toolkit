# Managed Entity

## Identity
- ME ID: 6
- ME Name: Circuit pack
- Source Section: 9.1.6
- Source Page: 77

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Get, Set, Create, Delete, Test
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: Type
- Size: 1 byte
- Format: needs_review
- Access: R, set-by-create if applicable
- Category: mandatory

### Attribute 2
- Name: Number of ports
- Size: 1 byte
- Format: needs_review
- Access: R
- Category: optional

### Attribute 3
- Name: Serial number
- Size: 8 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 4
- Name: Version
- Size: 14 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 5
- Name: Administrative state
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 6
- Name: Operational state
- Size: 1 byte
- Format: needs_review
- Access: R
- Category: optional

### Attribute 7
- Name: Bridged or IP ind
- Size: 20 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 8
- Name: Card configuration
- Size: 4 bytes
- Format: needs_review
- Access: R, W
- Category: optional

## Extraction Status
- Status: auto_extracted
- Attributes found: 8
- Review needed: false

## Raw Source

```
9.1.6 Circuit pack  This ME models a real or virtual circuit pack that is equipped in a real or virtual ONU slot. For ONUs  with integrated interfaces, this ME may be used to distinguish available types of interfaces (the port- mapping package is another way).  For ONUs with integrated interfaces, the ONU automatically creates an instance of this ME for each  instance of the virtual cardholder ME. The ONU also creates an instance of this ME when the OLT  provisions the cardholder to expect a circuit pack, i.e., when the OLT sets the expected plug -in unit  type or equipment ID of the cardholder to a circuit pack type, as defined in Table 9.1.5-1. The ONU  also creates an instance of this ME when a circuit pack is installed in a cardholder whose expected  plug-in unit type is 255 = plug-and-play, and whose equipment ID is not provisioned. Finally, when  the cardholder is provisioned for plug-and-play, an instance of this ME can be created at the request  of the OLT.  The ONU deletes an instance of this ME when the OLT de-provisions the circuit pack (i.e., when the  OLT sets the expected plug-in unit type or equipment ID of the cardholder to 0 = no LIM). The ONU  also deletes an instance of this ME on request of the OLT if the expected plug -in unit type attribute  of the corresponding cardholder is equal to 255, plug -and-play, and the expected equipment ID is  blank (a string of all spaces). ONUs with integrated interfaces do not delete circuit pack instances.  NOTE – Creation and deletion by the OLT is retained for backward compatibility.  Relationships  An instance of this ME is contained by an instance of the cardholder ME.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. Its value is  the same as that of the cardholder ME containing this circuit pack instance. (R,  set-by-create if applicable) (mandatory) (2 bytes)  Type: This attribute identifies the circuit pack type. This attribute is a code as defined  in Table 9.1.5-1. The value 255 means unknown or undefined, i.e., the inserted  circuit pack is not recognized by the ONU or is not mapped to an entry in Table  9.1.5-1. In the latter case, the equipment ID attribute may contain inventory  information. Upon autonomous ME instantiation, the ONU sets this attribute  to 0 or to the type of the circuit pack that is physically present. (R, set-by-create  if applicable) (mandatory) (1 byte)  Number of ports: This attribute is the number of access ports on the circuit pack. If the port- mapping package is supported for this circuit pack, this attribute should be set  to the total number of ports of all types. (R) (optional) (1 byte)  Serial number: The serial number is expected to be unique for each circuit pack, at least  within the scope of the given vendor. Note that the serial number may contain  the vendor ID or version number. For integrated ONUs, this value is identical  to the value of the serial number attribute of the ONU-G ME. Upon creation in  the absence of a physical circuit pack, this attribute comprises all spaces. (R)  (mandatory) (8 bytes)  Version: This attribute is a string that identifies the version of the circuit pack as defined  by the vendor. The value 0 indicates that version information is not available  or applicable. For integrated ONUs, this value is identical to the value of the  version attribute of the ONU-G ME. Upon creation in the absence of a physical  circuit pack, this attribute comprises all spaces. (R) (mandatory) (14 bytes) 

---- page break ---- Vendor ID: This attribute identifies the vendor of the circuit pack. For ONUs with  integrated interfaces, this value is identical to the value of the vendor ID  attribute of the ONU-G ME. Upon creation in the absence of a physical circuit  pack, this attribute comprises all spaces. (R) (optional) (4 bytes)  Administrative state: This attribute locks (1) and unlocks (0) the functions performed by this  ME. Administrative state is further described in clause A.1.6. (R,  W)  (mandatory) (1 byte)  Operational state: This attribute indicates whether the circuit pack is capable of performing  its function. Valid values are enabled (0), disabled (1) and unknown (2).  Pending completion of initialization and self -test on an installed circuit pack,  the ONU sets this attribute to 2. (R) (optional) (1 byte)  Bridged or IP ind: This attribute specifies whether an Ethernet interface is bridged or derived  from an IP router function.  0 Bridged  1 IP router  2 Both bridged and IP router functions  (R, W) (optional, only applicable for circuit packs with Ethernet interfaces)  (1 byte)  Equipment ID: This attribute may be used to identify the vendor 's specific type of circuit  pack. In some environments, this attribute may include the CLEI code. Upon  ME instantiation, the ONU sets this attribute to all spaces or to the equipment  ID of the circuit pack that is physically present. (R) (optional) (20 bytes)  Card configuration : This attribute selects the appropriate configuration of configurable  circuit packs. Table 9.1.5-1 specifies two configurable card types: C -DS1/E1  (code 16), and C -DS1/E1/J1 (code 17). Values are indicated below for the  allowed card types and configurations.  Card Type Configuration Value  C-DS1/E1 DS1 0   E1 1  C-DS1/E1/J1 DS1 0   E1 1   J1 2  Upon autonomous instantiation, this attribute is set to 0. (R,  W, set-by-create  if applicable) (mandatory for configurable circuit packs) (1 byte)  Total T-CONT buffer number: This attribute reports the total number of T -CONT buffers  associated with the circuit pack. Upon ME instantiation, the ONU sets this  attribute to 0 or to the value supported by the physical circuit pack. (R)  (mandatory for circuit packs that provide a traffic scheduler function) (1 byte)  Total priority queue number: This value reports the total number of priority queues  associated with the circuit pack. Upon ME instantiation, the ONU sets the  attribute to 0 or to the value su
```
