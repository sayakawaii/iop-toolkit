# Managed Entity

## Identity
- ME ID: 44
- ME Name: Vendor specific
- Source Section: needs_review
- Source Page: 96

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Get, Set
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: Manufacturer name
- Size: 25 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 2
- Name: Serial number part 1, serial number part 2
- Size: 25 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 3
- Name: Manufacturing date
- Size: 25 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 4
- Name: Hardware-revision
- Size: 25 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 5
- Name: Firmware-revision
- Size: 25 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 6
- Name: Current local ONU time
- Size: 7 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 7
- Name: Interval end time
- Size: 1 byte
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 8
- Name: Threshold data 1/2 ID
- Size: 2 bytes
- Format: needs_review
- Access: R, W, Set by create
- Category: mandatory

### Attribute 9
- Name: Temperature sensor value
- Size: 1 byte
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 10
- Name: RAM size available
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 11
- Name: RAM utilization
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 12
- Name: FLASH size available
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 13
- Name: FLASH utilization
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 14
- Name: Software errors
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 15
- Name: Errors in operations
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 16
- Name: Sustained downstream MAC client rate
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 17
- Name: Downstream RAB size
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 18
- Name: Sustained upstream MAC client rate
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 19
- Name: Upstream RAB size
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 20
- Name: SR indication
- Size: 1 byte
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 21
- Name: Total T-CONT number
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 22
- Name: GEM block length
- Size: 2 bytes
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 23
- Name: Piggyback DBA reporting
- Size: 1 byte
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 24
- Name: Deprecated
- Size: 1 byte
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 25
- Name: ARC interval
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 26
- Name: Optical signal level
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 27
- Name: Lower optical threshold
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 28
- Name: Upper optical threshold
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 29
- Name: ONU response time
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 30
- Name: Transmit optical level
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 31
- Name: Lower transmit power threshold
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 32
- Name: Upper transmit power threshold
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 33
- Name: Actions  Get, set  Test
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 34
- Name: Alloc-ID
- Size: 2 bytes
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 35
- Name: Deprecated
- Size: 1 byte
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 36
- Name: Policy
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: mandatory

## Extraction Status
- Status: auto_extracted
- Attributes found: 36
- Review needed: true

## Raw Source

```
Alarm  Alarm  number Alarm Description  4 Ground Fault Ground fault; ONU has detected a loss of grounding or a  degradation in the ground connection.  9.1.16 ONU manufacturing data  This ME contains additional manufacturing attributes associated with a PON ONU. The  manufacturing data is expected to match the content of an ONU label. The ONU automatically creates  an instance of this ME. Its attributes are populated according to data within the ONU itself.  Relationships  This ME is paired with the ONU-G entity.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. There is  only one instance, number 0. (R) (mandatory) (2 bytes)  Manufacturer name: This attribute contains the manufacturer name of this physical ONU.  The preferred value is the manufacturer name string printed on the ONU  itself (if present). (R) (optional) (25 bytes)  Serial number part 1, serial number part 2: These two attributes may be regarded as an  ASCII string of up to 32 bytes whose length is a left justified manufacturer's  serial number for this physical ONU. The preferred value is the manufacturer  serial number string printed on the ONU itself (if present). (R) (optional)  (25 bytes*2 attributes)  Model name: This attribute contains the vendor specific model name identifier string. The  preferred value is the customer-visible part number which may be printed on  the component itself. (R) (optional) (25 bytes)  Manufacturing date: This attribute contains the date of manufacturer of this physical ONU.  The preferred value is the date of the manufacturer printed on the ONU itself  (if present). (R) (optional) (25 bytes)  Hardware-revision: This attribute contains the hardware revision of this physical ONU.  The preferred value is the hardware revision printed on the ONU itself (if  present). (R) (optional) (25 bytes)  Firmware-revision: This attribute contains the vendor specific firmware revision of this  physical ONU. (R) (optional) (25 bytes)  Actions  Get  Notifications  None  9.1.17 ONU time configuration  This ME provides characterization and manipulation of OLT timestamp information. An ONU that  uses OLT-based time synchronization methods automatically creates an instance of this ME. There  is no intention that this ME be used to establish a precise time of day reference. 

---- page break ---- Relationships  The single instance of this ME is associated with the ONU ME.   Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. There is  only one instance, number 0. (R) (mandatory) (2 bytes)  Current local ONU time: If the ONU has a real-time clock, it returns the local ONU time.  This attribute returns the current ONU time. The local ONU time and  synchronize time time-format is the same. (R) (mandatory) (7 bytes)  Byte Description  1-2 Year, e.g., 2009  3 Month, range 1..12  4 Day of month, range 1..31  5 Hour of day, range 0..23  6 Minute of hour, range 0..59  7 Second of minute, range 0..59  Time qualification block: This attribute describes the time-qualification to be applied to the  ONU RTC local time. The following fields are supported:    Bits  16 15 14 13 12 11 10 9 8 7 6 5 4 3 2 1  Time  qualification  Timezone  offset  Reserved RFC  3339  time  offset  sign  RFC 3339 minutes offset (from UTC  time)  where  Time  qualification  (1-bit)  This attribution qualifies the OLT time. Valid values are:  0 The OLT is locally timed  1 The OLT timestamps are UTC timestamped    Timezone  information  (1-bit)  This attribute governs the time offset behaviour of the ONU RTC:  0 No timezone offset information  1 Timezone offset provided in the RFC 3339 time offset      RFC 3339  Time offset  sign (from  UTC time)  (1-bit)  RFC 3339 Timesone offset sign. A value of 0 indicates a positive offset and a  1 indicates a negative offset  RFC 3339  minutes  offset (from  UTC time)  (11-bits)  Describes the RFC 3339 minutes offset 

---- page break ---- (R, W) (mandatory) (2 bytes)  Actions  Get, set  Notifications  None.  9.1.18 ONU operational performance monitoring history data  This managed entity collects performance monitoring data associated with the ONU instances of this  managed entity are created and deleted by the OLT.  For a complete discussion of generic PM architecture, refer to clause I.4.  Relationships  An instance of this ME is associated with the ONU-G managed entity.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. There is only  one instance, number 0. (R) (mandatory) (2 bytes)  Interval end time : This attribute identifies the most recently finished 15  min interval. (R)  (mandatory) (1 byte)  Threshold data 1/2 ID: This attribute points to an instance of the threshold data 1 managed  entity that contains PM threshold values. Since no threshold value attribute  number exceeds 7, a threshold data 2 ME is optional. (R, W, Set by create)  (mandatory) (2 bytes)  Temperature sensor value: A table of one -byte temperature sensor values, each being  represented by a 2s complement integer that specifies the average temperature  of the ONU temperature sensor (s) during the measurement interval . Valid  values are –40 to +127 °C in 1 °C increments. The special values: 0x80  indicates that the temperature sensor is not available; 0x81, that the sensor has  malfunctioned. (R) (mandatory) (N byte)  NOTE – Sampling rates of temperature are outside the scope of this  specification).  Temperature sensor description: A table of 25 -byte long temperature sensor descriptions,  each represented by a character string that includes the physical location on the  ONU or the component being measured. Strings shorter than 25 bytes are  padded with null characters. (R) (mandatory) (25N bytes).  BEST PRACTICE NOTE: The order of the temperature sensor description and  temperature sensor values should match and be maintained. Care should be  taken that the temperature sensor names do not change across measurem
```
