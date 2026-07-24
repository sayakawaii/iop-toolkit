# Managed Entity

## Identity
- ME ID: 455
- ME Name: Link aggregation service profile
- Source Section: 9.3.35
- Source Page: 218

## Classification
- Access: needs_review
- Type: needs_review
- Actions: needs_review
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: Frame distribution and collection mode
- Size: 1 byte
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 2
- Name: MAC bridge service profile ID pointer
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 3
- Name: IP options
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 4
- Name: MAC address
- Size: 6 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 5
- Name: Onu identifier
- Size: 25 bytes
- Format: needs_review
- Access: R, W
- Category: mandatory

## Extraction Status
- Status: auto_extracted
- Attributes found: 5
- Review needed: true

## Raw Source

```
9.3.35 Link aggregation service profile   This ME organizes data associated with a link aggregation group (LAG) in PON channel bonding.  The OLT creates one instance of this managed entity for each channel bonding LAG.  Relationships  One instance of this managed entity exists for each LAG. An instance of this managed entity  is associated with one or more instances of MAC bridge port configuration data. An instance  of this managed entity is associated with one instances of MAC bridge service profile.   Attributes  Managed entity ID: This attribute uniquely identifies each instance of this managed entity.  Its value is created by the OLT. (R, set-by-create) (mandatory) (2 bytes)  Frame distribution and collection mode: This attribute indicates the frame distribution  and collection mode supported on this LAG. Valid values are:  1 Per-service frame distribution. This mode distributes frames to the  bonded links according to their service identifier. The algorithm is  specified in IEEE 802.1AX clause 8.  2 Round-robin frame distribution. This mode distributes frames to the  bonded links in circular order.  (R, W, set-by-create) (mandatory) (1 byte) 

---- page break ---- Type:  This attribute indicates the LASP ME type. Valid values are:  0 ANI-G interface channel bonding.   1 UNI-G interface channel bonding.  (R ,W, Set-by-create) (mandatory) (1 byte)  MAC bridge service profile ID pointer: This attribute points to an instance of the MAC  bridge service profile ME. (R, W, set-by-create) (mandatory) (2 bytes)  Actions  Create, delete, get, set  Notifications  None.  9.4 Layer 3 data services  9.4.1 IP host config data  The IP host config data configures IPv4 based services offered on the ONU. The ONU automatically  creates instances of this ME if IP host services are available. A possible IPv6 stack is supported  through the IPv6 host config data ME. In this clause, references to IP addresses are understood to  mean IPv4.  Relationships  An instance of this ME is associated with the ONU ME. Any number of TCP/UDP config  data MEs can point to the IP host config data, to model any number of ports and protocols.  Performance may be monitored through an implicitly linked IP host PM history data ME.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. The ONU  creates as many instances as there are independent IPv4 stacks on the ONU.  To facilitate discovery, IP host config data MEs should be numbered from 0  upwards. The ONU should create IP(v4) and IPv6 host config data MEs with  separate ME IDs, such that other MEs can use a single TP type attribute to link  with either. (R) (mandatory) (2 bytes)  IP options: This attribute is a bit map that enables or disables IP-related options. The value  1 enables the option while 0 disables it. The default value of this attribute is 0.  0x01 Enable DHCP  0x02 Respond to pings  0x04 Respond to traceroute messages  0x08 Enable IP stack  0x10..0x80 Reserved  (R, W) (mandatory) (1 byte)  MAC address: This attribute indicates the MAC address used by the IP node. (R) (mandatory)  (6 bytes)  Onu identifier: A unique ONU identifier string. If set to a non -null value, this string is used  instead of the MAC address in retrieving dynamic host configuration protocol  (DHCP) parameters. If the string is shorter than 25 characters, it must be null  terminated. Its default value is 25 null bytes. (R, W) (mandatory) (25 bytes)  Several attributes of this ME may be paired together into two categories, manual settings and current  values.
```
