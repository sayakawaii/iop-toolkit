# Managed Entity

## Identity
- ME ID: 347
- ME Name: IPv6 host config data
- Source Section: 9.4.5
- Source Page: 225

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Get, Set, Test
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: IP options
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 2
- Name: MAC address
- Size: 6 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 3
- Name: Onu identifier
- Size: 25 bytes
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 4
- Name: IPv6 link local address
- Size: 16 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 5
- Name: IPv6 address
- Size: 16 bytes
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 6
- Name: Default router
- Size: 16 bytes
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 7
- Name: Primary DNS
- Size: 16 bytes
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 8
- Name: Secondary DNS
- Size: 16 bytes
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 9
- Name: Current address table
- Size: 25 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 10
- Name: On-link prefix
- Size: 17 bytes
- Format: needs_review
- Access: R,W
- Category: optional

### Attribute 11
- Name: Current on-link prefix table
- Size: 2 bytes
- Format: needs_review
- Access: R, W
- Category: optional

## Extraction Status
- Status: auto_extracted
- Attributes found: 11
- Review needed: false

## Raw Source

```
9.4.5 IPv6 host config data  The IPv6 host config data configures IPv6 based services offered on the ONU. The ONU  automatically creates instances of this ME if IPv6 host services are available. If an IPv4 stack is  present, it is independently supported through the IP host config data ME.  This ME may be statically provisioned or may derive its parameters from router advertisements (RAs)  or DHCPv6.  Relationships  One or more instances of this ME are associated with the ONU ME. Any number of TCP/UDP  config data MEs can point to the IPv6 host config data, to model any number of ports and  protocols. Performance may be monitored through an implicitly linked IP host PM history  data ME.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. The ONU  creates as many instances as there are independent IP stacks on the ONU. To  facilitate discovery, IP and IPv6 host config data MEs should be numbered  from 0 upwards. The ONU must create IP(v4) and IPv6 host config data MEs  with separate ME IDs, such that other MEs can use a single TP type attribute  to link with either. (R) (mandatory) (2 bytes)  IP options: This attribute is a bit map that enables or disables IPv6 related options. The value  1 enables the option, while 0 disables it. The default value of this attribute is  0. (R, W) (mandatory) (1 byte)  0x01 IPv6 stack administrative unlock.  0x02 Enable router solicitation  (RS). The host generates RS  messages, if necessary, and responds to RAs. If the RA message has  the M flag set to 1, the ONU is expected to request the address and  other configuration information via DHCPv6. If the RA message has  the O flag set to 1 and M to 0, the ONU is expected to only request  additional configuration information via DHCPv6, but not addresses.  0x04 Enable DHCPv6.  0x08 Respond to pings (ICMPv6 echo replies)  0x10..0x80 Reserved  The following IP stack initialization flow is expected. 

---- page break ---- 1. If the IPv6 stack is administratively unlocked (0x01), establish a link-local  address [self-assign address, duplicate address detection (DAD) to confirm  that address is unique within the local link ]. This process is defined in [b - IETF RFC 4862], and is a part of stateless address autoconfiguration   (SLAAC). However, no IP options are set to enable or disable this function  – it always happens.  2. If RS and DHCPv6 are both disabled, do nothing. Manual settings are to  be used and are required for this IPv6 stack to be fully functional.  3. If RS is enabled (0x02) and DHCPv6 is disabled, send RS and listen for an  RA. If no RA is received, the ONU never attempts DHCPv6 and cannot  complete automated initialization of IPv6. If an RA is received  then the  following occur.  a. The ONU builds its default router table per [b -IETF RFC 4861],  reported via the current default router table attribute.   b. If the received RA includes information option(s) with an "A" prefix,  then ONU assigns itself an address from all such A prefixes, per [b - IETF RFC 4862] (also part of SLAAC) (reported via the current  address table attribute).   c. If the received RA includes DNS information [b-IETF RFC 6106], the  ONU accepts it (reported via the current DNS table attribute). Support  for RFC 6106 is strongly recommended.   d. If the received RA has M = 1, then the ONU requests identity  association for non -temporary addresses  (IA_NA) and other options  (which could include DNS) via DHCPv6. The access network provider  is responsible for ensuring that, if it sends DNS information both in the  RA and DHCPv6, it sends the same DNS information; it must not rely  on the ONU to figure out whether RA DNS is preferred over DHCPv6  DNS or vice versa . If different DNS information is received via  DHCPv6, it also goes into the current DNS table attribute. If the ONU  gets IA_NA via DHCPv6, this goes into the current address table  attribute.  e. If the received RA has M = 0 and O = 1, then the ONU requests stateless  options (which could include DNS) via DHCPv6.   4. If RS and DHCPv6 are both enabled, the ONU does RS (as described in a- c above, if RA is received) and DHCPv6 (requesting IA_NA and other  options, as described in d above) simultaneously, effectively ignoring M  and O flags.  5. If RS is disabled and DHCPv6 is enabled, then the ONU does not send RS  and it does send DHCPv6 (requesting IA_NA and other options, as  described in d above). If an unsolicited RA is received, it is ignored.  MAC address: This attribute indicates the MAC address used by the IP node. (R) (mandatory)  (6 bytes)  Onu identifier: A unique ONU identifier string. If set to a non -null value, this string is used  instead of the MAC address in retrieving DHCPv6 parameters. If the string is  shorter than 25 characters, it must be null terminated. Its default value is 25  null bytes. (R, W) (mandatory) (25 bytes) 

---- page break ---- Several attributes of this ME may be paired together into two categories, manual settings and current  values.    Manual settings Current values  IPv6 address Current address table  Default router Current default router table  Primary DNS Current DNS table  Secondary DNS   On-link prefix Current on-link prefix table  While this ME instance is administratively locked, it provides no IPv6 connectivity to the external  world. Especially if manual provisioning is to be used, it is important that the ME remain locked until  provisioning is complete.  While autoconfiguration is disabled, the current values are the same as the manual settings. While  autoconfiguration is enabled, the current values are those autoconfigured on the basis of RAs,  assigned by DHCPv6, or undefined (empty tables) if no values have (yet) been assigned.  IPv6 link local address : The address used for on -link IP host services, such as RS and  DHCPv6. [b-IETF RFC 4862] specifies how to automatically establish a link- local address. (R) (mandatory) (16 bytes)  IPv6 address: The manually prov
```
