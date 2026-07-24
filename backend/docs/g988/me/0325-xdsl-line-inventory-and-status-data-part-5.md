# Managed Entity

## Identity
- ME ID: 325
- ME Name: xDSL line inventory and status data part 5
- Source Section: 9.7.12
- Source Page: 273

## Classification
- Access: needs_review
- Type: needs_review
- Actions: needs_review
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: xTU-C system vendor ID
- Size: 8 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 2
- Name: xTU-R system vendor ID
- Size: 8 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 3
- Name: xTU-C version number
- Size: 16 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 4
- Name: xTU-R version number
- Size: 16 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 5
- Name: xTU-C serial number part 1
- Size: 16 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 6
- Name: xTU-C serial number part 2
- Size: 16 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 7
- Name: xTU-R serial number part 1
- Size: 16 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 8
- Name: xTU-R serial number part 2
- Size: 16 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 9
- Name: xTU-C self-test results
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

## Extraction Status
- Status: auto_extracted
- Attributes found: 9
- Review needed: true

## Raw Source

```
9.7.12 xDSL line inventory and status data part 1  This ME contains part 1 of the line inventory and status data for an xDSL UNI. The ONU  automatically creates or deletes an instance of this ME upon the creation or deletion of a PPTP xDSL  UNI part 1.  Relationships  An instance of this ME is associated with an xDSL UNI.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. Through an  identical ID, this ME is implicitly linked to an instance of the PPTP xDSL UNI  part 1. (R) (mandatory) (2 bytes)  xTU-C G.994.1 vendor ID : This is the vendor ID as inserted by the xTU -C in the  ITU-T G.994.1 CL message. It comprises 8 octets, including a country code  followed by a (regionally allocated) provider code, as defined in [ITU-T T.35].  (R) (mandatory) (8 bytes) 

---- page break ---- xTU-R G.994.1 vendor ID : This is the vendor ID as inserted by the xTU -R in the  ITU-T G.994.1 CLR message. It comprises 8 binary octets, with the same  format as the xTU-C ITU-T G.994.1 vendor ID. (R) (mandatory) (8 bytes)  xTU-C system vendor ID : This is the vendor ID as inserted by the xTU -C in the overhead  messages of [ITU -T G.992.3] and [ITU -T G.992.4]. It comprises 8 binary  octets, with the same format as the xTU -C ITU -T G.994.1 vendor ID. (R)  (mandatory) (8 bytes)  xTU-R system vendor ID: This is the vendor ID as inserted by the xTU -R in the embedded  operations channel and overhead messages of [ITU -T G.992.3] and  [ITU-T G.992.4]. It comprises 8 binary octets, with the same format as the  xTU-C ITU-T G.994.1 vendor ID. (R) (mandatory) (8 bytes)  xTU-C version number: This is the vendor-specific version number as inserted by the xTU-C  in the overhead messages of [ITU -T G.992.3] and [ITU -T G.992.4]. It  comprises up to 16 binary octets. (R) (mandatory) (16 bytes)  xTU-R version number : This is the version number as inserted by the xTU -R in the  embedded operations channel of [ITU-T G.992.1] or [ITU-T G.992.2], or the  overhead messages of [ITU -T G.992.3], [ITU -T G.992.4], [ITU -T G.992.5]  and [ITU -T G.993.2]. The attribute value may be ve ndor-specific, but is  recommended to comprise up to 16 ASCII characters, null -terminated if it is  shorter than 16. The string should contain the xTU-R firmware version and the  xTU-R model, encoded in that order and separated by a space character:  "<xTU-R firmware version><xTU -R model>". It is recognized that legacy  xTU-Rs may not support this format. (R) (mandatory) (16 bytes)  xTU-C serial number part 1 : The vendor-specific serial number inserted by the xTU -C in  the overhead messages of [ITU-T G.992.3] and [ITU-T G.992.4] comprises up  to 32 ASCII characters, null terminated if it is shorter than 32 characters. This  attribute contains the first 16 characters. (R) (mandatory) (16 bytes)  xTU-C serial number part 2: This attribute contains the second 16 characters of the xTU-C  serial number. (R) (mandatory) (16 bytes)  xTU-R serial number part 1 : The serial number inserted by the xTU -R in the embedded  operations channel of [ITU -T G.992.1] or [ITU -T G.992.2], or the overhead  messages of [ITU -T G.992.3], [ITU -T G.992.4], [ITU -T G.992.5] and  [ITU-T G.993.2], comprises up to 32 ASCII characters, null-terminated if it is  shorter than 32. It is recommended that the equipment serial number, the  equipment model and the equipment firmware version, encoded in that order  and separated by space characters , be contained : "<equipment s erial  number><equipment model><equipment firmware version>". It is recognized  that legacy xTU -Rs may not support this format. This attribute contains the  first 16 characters. (R) (mandatory) (16 bytes)  xTU-R serial number part 2: This attribute contains the second 16 characters of the xTU-R  serial number. (R) (mandatory) (16 bytes)  xTU-C self-test results: This parameter reports the xTU-C self-test result. It is coded in two  fields. The most significant octet is 0 if the self -test passed and 1 if it failed.  The three least significant octets are a vendor-discretionary integer that can be  interpreted in com bination with [ITU -T G.994.1] and the system vendor ID.  (R) (mandatory) (4 bytes)
```
