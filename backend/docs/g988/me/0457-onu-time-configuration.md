# Managed Entity

## Identity
- ME ID: 457
- ME Name: ONU time configuration
- Source Section: 9.1.17
- Source Page: 96

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Get, Set
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: Current local ONU time
- Size: 7 bytes
- Format: needs_review
- Access: R
- Category: mandatory

## Extraction Status
- Status: auto_extracted
- Attributes found: 1
- Review needed: false

## Raw Source

```
9.1.17 ONU time configuration  This ME provides characterization and manipulation of OLT timestamp information. An ONU that  uses OLT-based time synchronization methods automatically creates an instance of this ME. There  is no intention that this ME be used to establish a precise time of day reference. 

---- page break ---- Relationships  The single instance of this ME is associated with the ONU ME.   Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. There is  only one instance, number 0. (R) (mandatory) (2 bytes)  Current local ONU time: If the ONU has a real-time clock, it returns the local ONU time.  This attribute returns the current ONU time. The local ONU time and  synchronize time time-format is the same. (R) (mandatory) (7 bytes)  Byte Description  1-2 Year, e.g., 2009  3 Month, range 1..12  4 Day of month, range 1..31  5 Hour of day, range 0..23  6 Minute of hour, range 0..59  7 Second of minute, range 0..59  Time qualification block: This attribute describes the time-qualification to be applied to the  ONU RTC local time. The following fields are supported:    Bits  16 15 14 13 12 11 10 9 8 7 6 5 4 3 2 1  Time  qualification  Timezone  offset  Reserved RFC  3339  time  offset  sign  RFC 3339 minutes offset (from UTC  time)  where  Time  qualification  (1-bit)  This attribution qualifies the OLT time. Valid values are:  0 The OLT is locally timed  1 The OLT timestamps are UTC timestamped    Timezone  information  (1-bit)  This attribute governs the time offset behaviour of the ONU RTC:  0 No timezone offset information  1 Timezone offset provided in the RFC 3339 time offset      RFC 3339  Time offset  sign (from  UTC time)  (1-bit)  RFC 3339 Timesone offset sign. A value of 0 indicates a positive offset and a  1 indicates a negative offset  RFC 3339  minutes  offset (from  UTC time)  (11-bits)  Describes the RFC 3339 minutes offset 

---- page break ---- (R, W) (mandatory) (2 bytes)  Actions  Get, set  Notifications  None.  
```
