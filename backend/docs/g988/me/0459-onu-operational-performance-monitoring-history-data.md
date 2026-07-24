# Managed Entity

## Identity
- ME ID: 459
- ME Name: ONU operational performance monitoring history data
- Source Section: 9.1.18
- Source Page: 98

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Create, Delete, Get, Set
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: Interval end time
- Size: 1 byte
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 2
- Name: Threshold data 1/2 ID
- Size: 2 bytes
- Format: needs_review
- Access: R, W, Set by create
- Category: mandatory

### Attribute 3
- Name: Temperature sensor value
- Size: 1 byte
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 4
- Name: RAM size available
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 5
- Name: RAM utilization
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 6
- Name: FLASH size available
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 7
- Name: FLASH utilization
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 8
- Name: Software errors
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 9
- Name: Errors in operations
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

## Extraction Status
- Status: auto_extracted
- Attributes found: 9
- Review needed: false

## Raw Source

```
9.1.18 ONU operational performance monitoring history data  This managed entity collects performance monitoring data associated with the ONU instances of this  managed entity are created and deleted by the OLT.  For a complete discussion of generic PM architecture, refer to clause I.4.  Relationships  An instance of this ME is associated with the ONU-G managed entity.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. There is only  one instance, number 0. (R) (mandatory) (2 bytes)  Interval end time : This attribute identifies the most recently finished 15  min interval. (R)  (mandatory) (1 byte)  Threshold data 1/2 ID: This attribute points to an instance of the threshold data 1 managed  entity that contains PM threshold values. Since no threshold value attribute  number exceeds 7, a threshold data 2 ME is optional. (R, W, Set by create)  (mandatory) (2 bytes)  Temperature sensor value: A table of one -byte temperature sensor values, each being  represented by a 2s complement integer that specifies the average temperature  of the ONU temperature sensor (s) during the measurement interval . Valid  values are –40 to +127 °C in 1 °C increments. The special values: 0x80  indicates that the temperature sensor is not available; 0x81, that the sensor has  malfunctioned. (R) (mandatory) (N byte)  NOTE – Sampling rates of temperature are outside the scope of this  specification).  Temperature sensor description: A table of 25 -byte long temperature sensor descriptions,  each represented by a character string that includes the physical location on the  ONU or the component being measured. Strings shorter than 25 bytes are  padded with null characters. (R) (mandatory) (25N bytes).  BEST PRACTICE NOTE: The order of the temperature sensor description and  temperature sensor values should match and be maintained. Care should be  taken that the temperature sensor names do not change across measurement  intervals nor ONU reboots. A temperature sensor name may change as a result  of a ONU software upgrade (when a new sensor is reported upon).  CPU percent utilization: The maximum system CPU utilization (high water mark) during  the measurement interval. For multi -processor systems, the CPU utilization  values are global maximum across all processors (e.g., if one core is 50% CPU  utilization, 3 other cores are 10% CPU util ization, the global CPU utilization  is 50%). This attribute is an integer ranging from 0 to 100. A value of 100  indicates that at least one CPU was fully utilized, and a value of 0 indicates all 

---- page break ---- the CPUs were idle during the measurement interval. The value of 0xFF  indicates that no reliable measurement is available. (R) (mandatory) (1 byte)  RAM size available: The minimum RAM size (in Megabytes ) available during the  measurement interval. This attribute is an integer from 1 to 2 32 – 2. The value  of 0xFFFFFFFF indicates that RAM size report is not reliable. (R) (mandatory)  (4 bytes)   RAM utilization: The maximum RAM size (in Megabytes) utilized during the measurement  interval. This attribute is an integer from 0 to 2 32 – 2. The value of  0xFFFFFFFF indicates that no reliable measurement is available. (R)  (mandatory) (4 bytes)  FLASH size available: The minimum  FLASH size (in Megabytes ) available during the  measurement interval. This attribute is an integer from 1 to 2 32 – 2. The value  of 0xFFFFFFFF indicates that FLASH size report is not reliable. (R)  (mandatory) (4 bytes)  FLASH utilization: The maximum  FLASH size (in Megabytes) utilized during the  measurement interval. This attribute is an integer from 0 to 2 32 – 2. The value  of 0xFFFFFFFF indicates that no reliable measurement is available. (R)  (mandatory) (4 bytes)  Software errors: A count of the number of software errors  detected. A software error is an  error, flaw, failure or fault in a computer program that causes it to produce an  incorrect or unexpected result, or to behave in unintended ways. Examples  include internal logical inconsistencies, divi sion by zero, referencing to non - existent memory, writing to read -only-memory and "exceptions" in certain  programming languages (such as C++ and Java) (R) (mandatory) (4 bytes)  Errors in operations: A count of the number of detected errors in operations, not due to a  software error. Examples include reading MEs that do not exist, provisioning  services on entities that do not exist, deleting entities that do not exist. (R)  (mandatory) (4 bytes)  Actions  Create, delete, get, get next, set  Notifications  Threshold crossing alert  Alarm  number Threshold crossing alert Threshold value attribute No. (Note)  0 CPU utilization 1  1 RAM utilization 2  2 FLASH utilization 3  3 Software errors 4  4 Errors in operations 5  NOTE – This number associates the TCA with the specified threshold value attribute of the  threshold data 1 managed entity.   

---- page break ---- 
```
