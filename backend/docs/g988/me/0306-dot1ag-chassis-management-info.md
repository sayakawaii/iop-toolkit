# Managed Entity

## Identity
- ME ID: 306
- ME Name: Dot1ag chassis-management info
- Source Section: 9.3.26
- Source Page: 193

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Get, Set
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: Chassis ID length
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 2
- Name: Chassis ID subtype
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 3
- Name: Chassis ID part 1, Chassis ID part 2
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 4
- Name: Management address domain 1, Management address domain 2
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: mandatory

## Extraction Status
- Status: auto_extracted
- Attributes found: 4
- Review needed: false

## Raw Source

```
9.3.26 Dot1ag chassis-management info  This ME represents the system -level chassis ID or management address for [IEEE  802.1ag] CFM  messages, and potentially for other IEEE 802-based functions. Although [IEEE 802.1AB] allows for  several management addresses (synonyms in different formats or with granularity to the component  level), [IEEE 802.1ag] does not provide for more than one. Nor is it expected that an ONU would  require more than one format. Accordingly, this ME provides for only one.  According to sender ID permission attributes in several dot1ag MEs, transmitted IEEE 802.1ag CFM  messages may include either or both of the chassis ID or management address fields.[IEEE 802.1ag]  requires that CCMs do not exceed 128 bytes, of which 74 are separately allocated to other purposes;  the sender ID TLV, if present, must accommodate this requirement. The chassis info and management 

---- page break ---- info must fit, with a minimum of 4 additional overhead bytes, into the remaining 54 bytes. This limit  is exploited in defining the maximum size of the ME's attributes.  Relationships  If an ONU supports [IEEE 802.1ag] functionality, it automatically creates an instance of this  ME.  Attributes  Managed entity ID: This attribute uniquely identifies this ME. There is at most one instance,  whose value is 0. (R) (mandatory) (2 bytes)  Chassis ID length: The length of the chassis ID attribute (not including the chassis ID subtype  attribute), default value 0. (R, W) (mandatory) (1 byte)  Chassis ID subtype : The format of the chassis ID attribute, default value 7, as defined in  [IEEE 802.1AB]:    1 Chassis  component  A particular instance of the entPhysicalAlias object (defined in  [IETF RFC 4133]) for a chassis component.  2 Interface  alias  A particular instance of the ifAlias object (defined in  [IETF RFC 2863]) for an interface on the containing chassis.  3 Port  component  A particular instance of the entPhysicalAlias object (defined in  [IETF RFC 4133]) for a port or backplane component within  the containing chassis.  4 Mac address A particular unicast source address (encoded in network byte  order and IEEE 802.3 canonical bit order), of a port on the  containing chassis as defined in [IEEE 802].  5 Network  address  A particular network address, encoded in network byte order,  associated with one or more ports on the containing chassis.  The first octet contains the Internet Assigned Numbers  Authority [b-IANA] address family numbers enumeration value  for the specific address type, and octets 2 to N contain the  network address value in network byte order.  6 Interface  name  A particular instance of the ifName object (defined in  [IETF RFC 2863]) for an interface on the containing chassis.  7 Local Locally assigned chassis ID  (R, W) (mandatory) (1 byte)  Chassis ID part 1, Chassis ID part 2: These two attributes may be regarded as an octet string  of up to 50 bytes whose length is given by the chassis ID length attribute and  whose value is the left -justified chassis ID. (R, W) (mandatory) (25 bytes * 2  attributes)  Management address domain length : The length of the management address domain  attribute, default value 0. If this attribute has the value 0, all of the other  management address attributes are undefined. (R, W) (mandatory) (1 byte)  Management address domain 1, Management address domain 2: These two attributes may  be regarded as an octet string of up to 50  bytes whose length is given by the  management address domain length attribute and whose value is the left - justified management address domain. The attribute is coded as an object  identifier (OID) as per [ITU-T X.690], referring to a TDomain as defined in  [IETF RFC 2579]. Typical domain values include snmpUDPDomain (from 

---- page break ---- SNMPv2-TM [IETF RFC 3417]) and snmpIeee802Domain (from SNMP - IEEE 802-TM-MIB [IETF RFC 4789]). (R,  W) (mandatory) (25  bytes * 2  attributes)  Management address length: The length of the management address attribute, default value  0. (R, W) (mandatory) (1 byte)  Management address 1, Management address 2: These two attributes may be regarded as  an octet string of up to 50  bytes whose length is given by the management  address length attribute and whose value is the left -justified management  address. (R, W) (mandatory) (25 bytes * 2 attributes)  Actions  Get, set  Notifications  None.  
```
