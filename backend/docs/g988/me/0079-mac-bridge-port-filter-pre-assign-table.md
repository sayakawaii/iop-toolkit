# Managed Entity

## Identity
- ME ID: 79
- ME Name: MAC bridge port filter pre-assign table
- Source Section: 9.3.7
- Source Page: 147

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Get, Set
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: IPv6 multicast filtering
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 2
- Name: IPv4 broadcast filtering
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 3
- Name: RARP filtering
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 4
- Name: IPX filtering
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 5
- Name: NetBEUI filtering
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 6
- Name: AppleTalk filtering
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 7
- Name: Bridge management information filtering
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 8
- Name: ARP filtering
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: mandatory

## Extraction Status
- Status: auto_extracted
- Attributes found: 8
- Review needed: false

## Raw Source

```
Notifications  None.  9.3.7 MAC bridge port filter pre-assign table  This ME provides an alternate approach to DA filtering from that supported through the MAC bridge  port filter table data ME. This alternate approach is useful when all groups of addresses are stored  beforehand in the ONU, and the MAC bridge port filter pre-assign table ME designates which groups  are valid or invalid for filtering. On a circuit pack in which all groups of addresses are pre -assigned  and stored locally, the ONU creates or deletes an instance of this ME automatically upon creation or  deletion of a MAC bridge port configuration data ME.  Relationships  An instance of this ME is associated with an instance of a MAC bridge port configuration data  ME.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. Through an  identical ID, this ME is implicitly linked to an instance of the MAC bridge port  configuration data ME. (R) (mandatory) (2 bytes)  The following 10 attributes have similar definitions. Each permits the OLT to specify whether MAC  DAs or Ethertypes of the named type are forwarded (0) or filtered (1). In each case, the initial value  of the attribute is 0.    No. Protocol MAC address Ethertype Standard  1 IPv4 multicast 01.00.5E.00.00.00 –  01.00.5E.7F.FF.FF  – [b-IETF RFC 3232]  2 IPv6 multicast (Note) 33.33.00.00.00.00 –  33.33.FF.FF.FF.FF  – [IETF RFC 2464]  3 IPv4 broadcast  FF.FF.FF.FF.FF.FF 0x0800 [b-IETF RFC 3232]  4 RARP  FF.FF.FF.FF.FF.FF 0x8035 [b-IETF RFC 3232]  5 IPX  FF.FF.FF.FF.FF.FF 0x8137 [b-IETF RFC 3232]    09.00.1B.FF.FF.FF,  09.00.4E.00.00.02  –   6 NetBEUI  03.00.00.00.00.01 –   7 AppleTalk  FF.FF.FF.FF.FF.FF 0x809B,  0x80F3  [b-IETF RFC 3232]    09.00.07.00.00.00 –  09.00.07.00.00.FC,  09.00.07.FF.FF.FF  –   8 Bridge management  information  01.80.C2.00.00.00 –  01.80.C2.00.00.FF  – [IEEE 802.1D]  9 Address resolution  protocol (ARP)  FF.FF.FF.FF.FF.FF 0x0806 [b-IETF RFC 3232]  10 PPPoE broadcast FF.FF.FF.FF.FF.FF 0x8863 [b-IETF RFC 2516]  NOTE – The specified MAC address range does not distinguish network control traffic from user traffic.  The dot1 rate limiter may be a preferable way to limit the flow of IPv6 multicast traffic. 

---- page break ---- IPv4 multicast filtering: (R, W) (mandatory) (1 byte)  IPv6 multicast filtering: (R, W) (mandatory) (1 byte)  IPv4 broadcast filtering: (R, W) (mandatory) (1 byte)  RARP filtering: (R, W) (mandatory) (1 byte)  IPX filtering:  (R, W) (mandatory) (1 byte)  NetBEUI filtering: (R, W) (mandatory) (1 byte)  AppleTalk filtering: (R, W) (mandatory) (1 byte)  Bridge management information filtering: (R, W) (mandatory) (1 byte)  Note that some destination MAC addresses should never be forwarded,  considering the following rules of [IEEE 802.1D].  1 Addresses from 01.80.C2.00.00.00 to 01.80.C2.00.00.0F are reserved.  2 Addresses from 01.80.C2.00.00.20 to 01.80.C2.00.00.2F are used for  generic attribute registration protocol (GARP) applications.  ARP filtering: (R, W) (mandatory) (1 byte)  Point-to-point protocol over Ethernet (PPPoE) broadcast filtering: (R, W)  (mandatory) (1 byte)  Actions  Get, set  Notifications  None.  
```
