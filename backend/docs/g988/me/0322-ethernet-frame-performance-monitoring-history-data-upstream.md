# Managed Entity

## Identity
- ME ID: 322
- ME Name: Ethernet frame performance monitoring history data upstream
- Source Section: 9.3.30
- Source Page: 208

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
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 3
- Name: Drop events
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 4
- Name: Octets
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 5
- Name: Packets
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 6
- Name: Broadcast packets
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 7
- Name: Multicast packets
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 8
- Name: CRC errored packets
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 9
- Name: Undersize packets
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 10
- Name: Oversize packets
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 11
- Name: Packets 64 octets
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 12
- Name: Packets 65 to 127 octets
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 13
- Name: Packets 128 to 255 octets
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 14
- Name: Packets 512 to 1023 octets
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 15
- Name: Packets 1024 to 1518 octets
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

## Extraction Status
- Status: auto_extracted
- Attributes found: 15
- Review needed: false

## Raw Source

```
9.3.30 Ethernet frame performance monitoring history data upstream  This ME collects PM data associated with upstream Ethernet frame delivery. It is based on the  Etherstats group of [IETF RFC 2819]. Instances of this ME are created and deleted by the OLT.  For a complete discussion of generic PM architecture, refer to clause I.4.  NOTE 1 – Implementers are encouraged to consider the Ethernet frame extended PM ME defined in  clause 9.3.32, which collects the same counters in a more generalized way.  Relationships  An instance of this ME is associated with an instance of a MAC bridge port configuration  data. 

---- page break ---- Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. Through an  identical ID, this ME is implicitly linked to an instance of a MAC bridge port  configuration data. (R, set-by-create) (mandatory) (2 bytes)  Interval end time : This attribute identifies the most recently finished 15  min interval. (R)  (mandatory) (1 byte)  Threshold data 1/2 ID: This attribute points to an instance of the threshold data  1 ME that  contains PM threshold values. Since no threshold value attribute number  exceeds 7, a threshold data 2 ME is optional. (R, W, set-by-create) (mandatory)  (2 bytes)  Drop events: The total number of events in which packets were dropped due to a lack of  resources. This is not necessarily the number of packets dropped; it is the  number of times this event was detected. (R) (mandatory) (4 bytes)   Octets: The total number of upstream octets received, including those in bad packets,  excluding framing bits, but including FCS. (R) (mandatory) (4 bytes)  Packets: The total number of upstream packets received, including bad packets,  broadcast packets and multicast packets. (R) (mandatory) (4 bytes)  Broadcast packets: The total number of upstream good packets received that were directed  to the broadcast address. This does not include multicast packets. (R)  (mandatory) (4 bytes)  Multicast packets: The total number of upstream good packets received that were directed to  a multicast address. This does not include broadcast packets. (R) (mandatory)  (4 bytes)  CRC errored packets : The total number of upstream packets received that had a length  (excluding framing bits, but including FCS octets) of between 64  octets and  1518 octets, inclusive, but had either a bad FCS with an integral number of  octets (FCS error) or a bad FCS with a non -integral number of octets  (alignment error). (R) (mandatory) (4 bytes)  Undersize packets : The total number of upstream packets received that were less than  64 octets long, but were otherwise well formed (excluding framing bits, but  including FCS). (R) (mandatory) (4 bytes)  Oversize packets : The total number of upstream packets received that were longer than  1518 octets (excluding framing bits, but including FCS) and were otherwise  well formed. (R) (mandatory) (4 bytes)  NOTE 2 – If 2 000 byte Ethernet frames are supported, counts in this performance  parameter are not necessarily errors.  Packets 64 octets : The total number of upstream received packets (including bad packets)  that were 64  octets long, excluding framing bits but including FCS. (R)  (mandatory) (4 bytes)  Packets 65 to 127 octets : The total number of upstream received packets (including bad  packets) that were 65..127 octets long, excluding framing bits but including  FCS. (R) (mandatory) (4 bytes)  Packets 128 to 255 octets : The total number of upstream packets (including bad packets)  received that were 128..255 octets long, excluding framing bits but including  FCS. (R) (mandatory) (4 bytes) 

---- page break ---- Packets 256 to 511 octets : The total number of upstream packets (including bad packets)  received that were 256..511 octets long, excluding framing bits but including  FCS. (R) (mandatory) (4 bytes)  Packets 512 to 1023 octets : The total number of upstream packets (including bad packets)  received that were 512..1 023 octets long, excluding framing bits but including  FCS. (R) (mandatory) (4 bytes)   Packets 1024 to 1518 octets: The total number of upstream packets (including bad packets)  received that were 1024..1518 octets long, excluding framing bits , but  including FCS. (R) (mandatory) (4 bytes)   Actions  Create, delete, get, set  Get current data (optional)  Notifications  Threshold crossing alert  Alarm  number Threshold crossing alert Threshold data counter No. (Note)  0 Drop events 1  1 CRC errored packets 2  2 Undersize packets 3  3 Oversize packets 4  NOTE – This number associates the TCA with the specified threshold value attribute of the  threshold data 1 managed entity.  
```
