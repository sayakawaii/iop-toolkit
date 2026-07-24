# Managed Entity

## Identity
- ME ID: 311
- ME Name: Multicast subscriber monitor
- Source Section: 9.3.29
- Source Page: 207

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Create, Delete, Get, Set
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: ME type
- Size: 1 byte
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 2
- Name: Current multicast bandwidth
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 3
- Name: Join messages counter
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 4
- Name: Bandwidth exceeded counter
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: optional

## Extraction Status
- Status: auto_extracted
- Attributes found: 4
- Review needed: false

## Raw Source

```
9.3.29 Multicast subscriber monitor  This ME provides the current status of each port with respect to its multicast subscriptions. It may be  useful for status monitoring or debugging purposes. The status table includes all dynamic groups  currently subscribed by the port.  Relationships  Instances of this ME are created and deleted at the request of the OLT. One instance may exist  for each IEEE 802.1 UNI configured to support multicast subscription.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. Through an  identical ID, this ME is implicitly linked to an instance of the MAC bridge port  configuration data or IEEE 802.1p mapper ME. (R, set-by-create) (mandatory)  (2 bytes)  ME type: This attribute indicates the type of the ME implicitly linked by the ME ID  attribute.  0 MAC bridge port config data  1 IEEE 802.1p mapper service profile  (R, W, set-by-create) (mandatory) (1 byte)  Current multicast bandwidth : This attribute is the ONU 's (BE) estimate of the actual  bandwidth currently being delivered to this particular MAC bridge port over  all dynamic multicast groups. (R) (optional) (4 bytes)  Join messages counter : This attribute counts the number of times the corresponding  subscriber sent a join message that was accepted. When full, the counter rolls  over to 0. (R) (optional) (4 bytes)  Bandwidth exceeded counter : This attribute counts the number of join messages that did  exceed, or would have exceeded, the max multicast bandwidth, whether  accepted or denied. When full, the counter rolls over to 0. (R) (optional)  (4 bytes)  IPv4 active group list table: This attribute lists the groups from one of the related dynamic  access control list tables or the allowed preview groups table that are currently  being actively forwarded, along with the actual bandwidth of each. If a join  has been recognized from more than one IPv4 source address for a given group 

---- page break ---- on this UNI, there will be one table entry for each. Each table entry has the  following form.  – VLAN ID, 0 if not used (2 bytes)  – Source IP address, 0.0.0.0 if not used (4 bytes)  – Multicast destination IP address (4 bytes)  – Best efforts actual bandwidth estimate, bytes per second (4 bytes)  – Client (set -top box) IP address, i.e., the IP address of the device  currently joined (4 bytes)  – Time since the most recent join of this client to the IP channel, in  seconds (4 bytes)  – Reserved (2 bytes)  (R) (mandatory) (24N bytes)  IPv6 active group list table: This attribute lists the groups from one of the related dynamic  access control list tables or the allowed preview groups table that are currently  being actively forwarded, along with the actual bandwidth of each. If a join  has been recognized from more than one IPv6 source address for a given group  on this UNI, there will be one table entry for each. In mixed IPv4 -IPv6  scenarios, it is possible that some fields might be IPv4, in which case their 12  most significant bytes of the  given field are set to zero. Each table entry has  the form:  – VLAN ID, 0 if not used (2 bytes)  – Source IP address, 0 if not used (16 bytes)  – Multicast destination IP address (16 bytes)  – Best efforts actual bandwidth estimate, bytes per second (4 bytes)  – Client (set -top box) IP address, i.e., the IP address of the device  currently joined (16 bytes)  – Time since the most recent join of this client to the IP channel, in  seconds (4 bytes)  (R) (optional) (58N bytes)  Actions  Create, delete, get, get next, set  Notifications  None.  
```
