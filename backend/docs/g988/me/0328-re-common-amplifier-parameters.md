# Managed Entity

## Identity
- ME ID: 328
- ME Name: RE common amplifier parameters
- Source Section: 9.14.6
- Source Page: 481

## Classification
- Access: needs_review
- Type: needs_review
- Actions: needs_review
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: Gain
- Size: 1 byte
- Format: needs_review
- Access: R
- Category: optional

### Attribute 2
- Name: Upper gain threshold
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 3
- Name: Target gain
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 4
- Name: Device temperature
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 5
- Name: Lower device temperature  threshold
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 6
- Name: Upper device temperature threshold
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 7
- Name: Device bias current
- Size: 1 byte
- Format: needs_review
- Access: R
- Category: optional

### Attribute 8
- Name: Amplifier saturation output power
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 9
- Name: Amplifier noise figure
- Size: 1 byte
- Format: needs_review
- Access: R
- Category: optional

### Attribute 10
- Name: Amplifier saturation gain
- Size: 1 byte
- Format: needs_review
- Access: R
- Category: optional

## Extraction Status
- Status: auto_extracted
- Attributes found: 10
- Review needed: true

## Raw Source

```
9.14.6 RE common amplifier parameters  This ME organizes data associated with each  OA supported by the RE. The management ONU  automatically creates one instance of this ME for each upstream or downstream OA.  Relationships  An instance of this ME is associated with an instance of the RE downstream amplifier or RE  upstream amplifier ME.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. Through an  identical ID, this ME is implicitly linked to an instance of an  upstream or  downstream OA. The first byte is the slot ID (defined in clause 9.1.5). The  second byte is the port ID. (R) (mandatory) (2 bytes)  NOTE – The type of the linked ME can be determined by uniqueness of slot and port.  Gain: This attribute reports the current measurement of the OA's gain, in d ecibels.  Its value is a 2s complement integer with 0.25 dB granularity, and with a range  from –32 dB to 31.5 dB. The value 0x7F indicates that the current measured  gain is 0, i.e., negative infinity in decibels terms. (R) (optional) (1 byte) 

---- page break ---- Lower gain threshold: This attribute specifies the gain the RE uses to declare the low gain  alarm. Valid values are 0 dB (coded as 0x00) to 63.5 dB (coded as 0xFE). The  default value 0xFF selects the RE's internal policy. (R, W) (optional) (1 byte)  Upper gain threshold: This attribute specifies the gain the RE uses to declare the high gain  alarm. Valid values are 0 dB (coded as 0x00) to 63.5 dB (coded as 0xFE). The  default value 0xFF selects the RE's internal policy. (R, W) (optional) (1 byte)  Target gain: This attribute specifies the target gain, when the operational mode of the parent  RE downstream or upstream amplifier is set to constant gain mode. Valid  values are 0 dB (coded as 0x00) to 63.5 dB (coded as 0xFE). The default value  0xFF selects the RE's internal policy. (R, W) (optional) (1 byte)  Device temperature: This attribute reports the temperature in degrees Celcius of the active  device (SOA or pump) in the OA. Its value is a 2s complement integer with  granularity 1/256 °C. (R) (optional) (2 bytes)  Lower device temperature  threshold: This attribute is a 2s complement integer that  specifies the temperature the RE uses to declare the low temperature alarm.  Valid values are –64 to +63 °C in 0.5 °C increments. The default value 0x7F  selects the RE's internal policy. (R, W) (optional) (1 byte)  Upper device temperature threshold: This attribute is a 2s complement integer that specifies  the temperature the RE uses to declare the high temperature alarm. Valid  values are –64 to +63 °C in 0.5 °C increments. The default value 0x7F selects  the RE's internal policy. (R, W) (optional) (1 byte)  Device bias current: This attribute contains the measured bias current applied to the SOA or  pump laser. Its value is an unsigned integer with granularity 2  mA. Valid  values are 0 to 512 mA. (R) (optional) (1 byte)  Amplifier saturation output power: This attribute reports the saturation output power of the  amplifier as specified by the manufacturer. Its value is an unsigned integer  referred to 1 mW (i.e., dBm), with 0.1 dB granularity. (R) (optional) (2 bytes)  Amplifier noise figure : This attribute reports the intrinsic noise figure of the amplifier, as  specified by the manufacturer. Its value is an unsigned integer with 0.1  dB  granularity (R) (optional) (1 byte)  Amplifier saturation gain : This attribute reports the gain of the amplifier at saturation, as  specified by the manufacturer. Its value is an unsigned integer with 0.25  dB  granularity, and with a range from 0 to 63.75 dB. (R) (optional) (1 byte)  Actions  Get, set  Notifications  Alarm  Alarm  number Alarm Description  0 Low gain Gain below lower threshold  1 High gain Gain above upper threshold  2 Low temperature Device temperature below lower threshold  3 High temperature Device temperature above upper threshold  4 High bias current Device bias current above threshold determined by vendor;  device end of life pending 

---- page break ---- Alarm  Alarm  number Alarm Description  5 High temperature shutdown Device has shut down due to temperature exceeding  manufacturer's specifications  6 High current shutdown Device has shut down due to bias current exceeding  manufacturer's specifications  7..207 Reserved   208..223 Vendor-specific alarms Not to be standardized  9.15 RS232/RS485 interface service  This clause defines MEs associated with RS232/RS485 UNI, as shown in Figure 9.15-1.    Figure 9.15-1 – Managed entities associated with RS232/RS485 UNI  9.15.1 Physical path termination point RS232/RS485 UNI  This ME represents an RS232/RS485 UNI in the ONU, where physical paths terminate and physical  path level functions are performed.  The ONU automatically creates an instance of this ME per port as follows.  • When the ONU has RS232/RS485 ports built into its factory configuration.  • When a cardholder is provisioned to expect a circuit pack of RS232/RS485 type.  • When a cardholder provisioned for plug and play is equipped with a circuit pack of  RS232/RS485 type. Note that the installation of a plug and play card may indicate the  presence of RS232/RS485 ports via equipment ID as well as its type, and indeed may cause  the ONU to instantiate a port-mapping package that specifies RS232/RS485 ports.  The ONU automatically deletes instances of this ME when a cardholder is neither provisioned to  expect a RS232/RS485 circuit pack, nor is equipped with a RS232/RS485 circuit pack.  Relationships  An instance of this ME is associated with each real RS232/RS485 port.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. This 2 byte  number is directly associated with the physical position of the UNI. The first  byte is the slot ID (defined in clause 9.1.5). The second byte is the port ID,  with range 1..255. (R) (mandatory) (2 bytes)
```
