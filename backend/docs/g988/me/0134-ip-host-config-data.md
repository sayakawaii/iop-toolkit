# Managed Entity

## Identity
- ME ID: 134
- ME Name: IP host config data
- Source Section: 9.4.1
- Source Page: 219

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Get, Set
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
- Name: IP address
- Size: 4 bytes
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 5
- Name: Mask
- Size: 4 bytes
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 6
- Name: Gateway
- Size: 4 bytes
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 7
- Name: Primary DNS
- Size: 4 bytes
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 8
- Name: Secondary DNS
- Size: 4 bytes
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 9
- Name: Current address
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 10
- Name: Current mask
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 11
- Name: Current gateway
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 12
- Name: Current primary DNS
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 13
- Name: Current secondary DNS
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 14
- Name: Domain name
- Size: 25 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 15
- Name: Host name
- Size: 25 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 16
- Name: Relay agent options
- Size: 2 bytes
- Format: needs_review
- Access: R, W
- Category: optional

## Extraction Status
- Status: auto_extracted
- Attributes found: 16
- Review needed: false

## Raw Source

```
9.4.1 IP host config data  The IP host config data configures IPv4 based services offered on the ONU. The ONU automatically  creates instances of this ME if IP host services are available. A possible IPv6 stack is supported  through the IPv6 host config data ME. In this clause, references to IP addresses are understood to  mean IPv4.  Relationships  An instance of this ME is associated with the ONU ME. Any number of TCP/UDP config  data MEs can point to the IP host config data, to model any number of ports and protocols.  Performance may be monitored through an implicitly linked IP host PM history data ME.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. The ONU  creates as many instances as there are independent IPv4 stacks on the ONU.  To facilitate discovery, IP host config data MEs should be numbered from 0  upwards. The ONU should create IP(v4) and IPv6 host config data MEs with  separate ME IDs, such that other MEs can use a single TP type attribute to link  with either. (R) (mandatory) (2 bytes)  IP options: This attribute is a bit map that enables or disables IP-related options. The value  1 enables the option while 0 disables it. The default value of this attribute is 0.  0x01 Enable DHCP  0x02 Respond to pings  0x04 Respond to traceroute messages  0x08 Enable IP stack  0x10..0x80 Reserved  (R, W) (mandatory) (1 byte)  MAC address: This attribute indicates the MAC address used by the IP node. (R) (mandatory)  (6 bytes)  Onu identifier: A unique ONU identifier string. If set to a non -null value, this string is used  instead of the MAC address in retrieving dynamic host configuration protocol  (DHCP) parameters. If the string is shorter than 25 characters, it must be null  terminated. Its default value is 25 null bytes. (R, W) (mandatory) (25 bytes)  Several attributes of this ME may be paired together into two categories, manual settings and current  values.   

---- page break ---- Manual settings Current values  IP address Current address  Mask Current mask  Gateway Current gateway  Primary DNS Current primary DNS  Secondary DNS Current secondary DNS  While the IP stack is disabled, there is no IP connectivity to the external world from this ME instance.  While DHCP is disabled, the current values are always the same as the manual settings. While DHCP  is enabled, the current values are those assigned by DHCP, or undefined (0) if DHCP has never  assigned values.  IP address: The address used for IP host services; this attribute has the default value 0.  (R, W) (mandatory) (4 bytes)  Mask: The subnet mask for IP host services; this attribute has the default value 0.  (R, W) (mandatory) (4 bytes)  Gateway: The default gateway address used for IP host services; this attribute has the  default value 0. (R, W) (mandatory) (4 bytes)  Primary DNS: The address of the primary DNS server; this attribute has the default value 0.  (R, W) (mandatory) (4 bytes)  Secondary DNS: The address of the secondary DNS server; this attribute has the default value  0. (R, W) (mandatory) (4 bytes)  Current address: Current address of the IP host service. (R) (optional) (4 bytes)  Current mask: Current subnet mask for the IP host service. (R) (optional) (4 bytes)  Current gateway: Current default gateway address for the IP host service. (R) (optional)  (4 bytes)  Current primary DNS: Current primary DNS server address. (R) (optional) (4 bytes)  Current secondary DNS: Current secondary DNS server address. (R) (optional) (4 bytes)  Domain name: If DHCP indicates a domain name, it is presented here. If no domain name is  indicated, this attribute is set to a null string. If the string is shorter than  25 bytes, it must be null terminated. The default value is 25 null bytes. (R)  (mandatory) (25 bytes)  Host name: If DHCP indicates a host name, it is presented here. If no host name is  indicated, this attribute is set to a null string. If the string is shorter than  25 bytes, it must be null terminated. The default value is 25 null bytes. (R)  (mandatory) (25 bytes)  Relay agent options: This attribute is a pointer to a large string ME whose content specifies  one or more DHCP relay agent options. (R, W) (optional) (2 bytes)  The contents of the large string are parsed by the ONU and converted into text  strings. Variable substitution is based on defined three-character groups, each  of which begins with the '%' character. The string '%%' is an escape mechanism  whose output is a single '%' character. When the ONU cannot perform variable  substitution on a substring of the large string, it generates the specified option  as an exact quotation of the provisioned substring value. 

---- page break ---- Provisioning of the large string is separate from the operation of setting the  pointer in this attribute. It is the responsibility of the OLT to ensure that the  large string contents are correct and meaningful.  Three-character variable definitions are as follows. The first variable in the  large string must specify one of the option types. Both options for a given IP  version may be present if desired, each introduced by its option identifier.  Terminology is taken from clause 3.9.3 of [b-BBF TR-101].  %01, %18  Specifies that the following string is for option 82 sub-option 1,  agent circuit-ID (IPv4) or option 18, interface-ID (IPv6). The  equivalence permits the same large string to be used in both IP  environments.  %02, %37  Specifies that the following string is for option 82 sub-option 2,  relay agent remote-ID (IPv4) or option 37, relay agent remote- ID (IPv6). The equivalence permits the same large string to be  used in both IP environments.  %SL In [b-BBF TR-101], this is called a slot. In an ONU, this variable  refers to a shelf. It would be meaningful if the ONU has multiple  shelves internally or is daisy -chained to multiple equipment  modules. The range of this variable is "0".. "99"  %SU In TR -101, this is called a sub -slot. In fact, it represents
```
