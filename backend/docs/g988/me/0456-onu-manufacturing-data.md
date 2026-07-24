# Managed Entity

## Identity
- ME ID: 456
- ME Name: ONU manufacturing data
- Source Section: 9.1.16
- Source Page: 96

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Get
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

## Extraction Status
- Status: auto_extracted
- Attributes found: 5
- Review needed: false

## Raw Source

```
9.1.16 ONU manufacturing data  This ME contains additional manufacturing attributes associated with a PON ONU. The  manufacturing data is expected to match the content of an ONU label. The ONU automatically creates  an instance of this ME. Its attributes are populated according to data within the ONU itself.  Relationships  This ME is paired with the ONU-G entity.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. There is  only one instance, number 0. (R) (mandatory) (2 bytes)  Manufacturer name: This attribute contains the manufacturer name of this physical ONU.  The preferred value is the manufacturer name string printed on the ONU  itself (if present). (R) (optional) (25 bytes)  Serial number part 1, serial number part 2: These two attributes may be regarded as an  ASCII string of up to 32 bytes whose length is a left justified manufacturer's  serial number for this physical ONU. The preferred value is the manufacturer  serial number string printed on the ONU itself (if present). (R) (optional)  (25 bytes*2 attributes)  Model name: This attribute contains the vendor specific model name identifier string. The  preferred value is the customer-visible part number which may be printed on  the component itself. (R) (optional) (25 bytes)  Manufacturing date: This attribute contains the date of manufacturer of this physical ONU.  The preferred value is the date of the manufacturer printed on the ONU itself  (if present). (R) (optional) (25 bytes)  Hardware-revision: This attribute contains the hardware revision of this physical ONU.  The preferred value is the hardware revision printed on the ONU itself (if  present). (R) (optional) (25 bytes)  Firmware-revision: This attribute contains the vendor specific firmware revision of this  physical ONU. (R) (optional) (25 bytes)  Actions  Get  Notifications  None  
```
