# Managed Entity

## Identity
- ME ID: 466
- ME Name: Precision Time Protocol status
- Source Section: 9.12.23
- Source Page: 450

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Get
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: Clock identity
- Size: 8 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 2
- Name: Clock state
- Size: 1 byte
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 3
- Name: Clock state change time
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 4
- Name: Tx sync messages
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 5
- Name: Tx follow up messages
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 6
- Name: Rx delay request message s
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 7
- Name: Tx delay response message s
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

## Extraction Status
- Status: auto_extracted
- Attributes found: 7
- Review needed: false

## Raw Source

```
9.12.23 Precision Time Protocol status  This ME supplements the Precision Time Protocol ME. Its purpose is to return information that may  help in system- or network-level troubleshooting. An ONU that supports PTP automatically creates  or deletes an instance of this ME.  Relationships  The single instance of this ME is associated with the PTP ME.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. There is only  one instance, number 0. (R) (mandatory) (2 bytes)  Clock identity: This attribute identifies the clock identity attribute of the ONU PTP clock as  specified in clause 7.5.2.2 of [IEEE 1588v2]. It is expressed as a binary value  of 8 octets.  (R) (mandatory) (8 bytes)  Clock state: This attribute indicates the current ONU PTP clock state.  The following clock  states (based on [ITU-T G.8275 Appendix VIII.2] clock modes) are supported:   0x00 Free-Run mode: The operating condition of the ONU clock has  lost traceability to the OLT -G ME ToD timestamp. This may occur  when PTP is disabled, or PTP is enabled and OLT -G ME ToD  information is invalid.  0x01 Holdover-in-specification mode: The operating condition of the  ONU clock has lost traceability to the controlling OLT -G ME ToD  BUT is using the last valid OLT -G ME ToD  timestamp and its  performance is within desired specification. The ONU clock stopped  receiving OLT-G ME ToD timestamps for more than the 3 times the  ToD update interval.  0x02: Holdover-out-of-specification mode: The operating condition of the  ONU clock has lost traceability to the controlling OLT -G ME ToD  BUT is using the last valid OLT -G ME ToD timestamp and its  performance is outside desired specification.  0x03 Locked mode: The operating condition of the ONU clock is  locked/controlled by received OLT -G ToD. The ONU clock is  receiving OLT-G ME ToD timestamps within 3 times the ToD update  interval)  Note: The [ITU -T G.8275] clock mode of "Acquiring" is indicated as  "Free-run".  (R) (mandatory) (1 byte)  Clock state change time: A UTC timestamp at which the ONU last changed the clock state.   (R) (mandatory) (4 bytes)  Tx sync messages: This attribute increments when a sync message is sent by the ONU. When  full, the counter rolls over to 0.   (R) (mandatory) (4 bytes)  Tx follow up messages : This attribute increments when a follow up message is sent by the  ONU. When full, the counter rolls over to 0.   (R) (mandatory) (4 bytes) 

---- page break ---- Tx announce messages: This attribute increments when an announce message is sent by the  ONU. When full, the counter rolls over to 0.   (R) (mandatory) (4 bytes)  Rx delay request message s: This attribute increments when a delay request message is  received by the ONU. When full, the counter rolls over to 0.   (R) (mandatory) (4 bytes)  Tx delay response message s: This attribute increments when a delay response message is  sent by the ONU. When full, the counter rolls over to 0.   (R) (mandatory) (4 bytes)  PTP slave table: This attribute is a list of synchronized ptp slave nodes detected by ONU.   Note: Only discovered ptp  slave node whose PTP timeout timer has not expired  are included in this table (see clause 7.7.3 (PTP timeout timer) of  [IEEE 1588v2]). Each table entry has the below form:  MAC address: A unique identifier of the slave node within the table (6 bytes)  Source IP address: This component specifies the slave source IP address. May  be either an IPv4 address (first 12 bytes 0) or an IPv6 address. (16 bytes)   Timestamp: A UTC timestamp at which the ONU detected the slave node. (4  bytes)  (R) (mandatory) (26N bytes)  Actions  Get, Get next  Notifications  None  9.13 Miscellaneous services  9.13.1 Physical path termination point video UNI  This ME represents an RF video UNI in the ONU, where physical paths terminate and physical path  level functions are performed.  The ONU automatically creates an instance of this ME per port:  • when the ONU has RF video UNI ports built into its factory configuration;  • when a cardholder is provisioned to expect a circuit pack of the video UNI type;  • when a cardholder provisioned for plug-and-play is equipped with a circuit pack of the video  UNI type. Note that the installation of a plug -and-play card may indicate the presence of  video ports via equipment ID as well as its type, and indeed may cause the ONU to instantiate  a port-mapping package that specifies video ports.  The ONU automatically deletes instances of this ME when a cardholder is neither provisioned to  expect a video circuit pack, nor is it equipped with a video circuit pack.  Relationships  One or more instances of this ME are associated with an instance of a real or virtual circuit  pack classified as video type.
```
