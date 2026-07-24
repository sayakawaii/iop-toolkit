# Managed Entity

## Identity
- ME ID: 303
- ME Name: Dot1ag MEP status
- Source Section: 9.3.23
- Source Page: 188

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Get
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: MEP MAC address
- Size: 6 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 2
- Name: Fault notification generator state
- Size: 1 byte
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 3
- Name: Highest priority defect observed
- Size: 1 byte
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 4
- Name: Current defects
- Size: 1 byte
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 5
- Name: Last received errored CCM table
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 6
- Name: CCMs transmitted count
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 7
- Name: Unexpected LTRs count
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 8
- Name: Next loopback transaction identifier
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 9
- Name: Next linktrace transaction identifier
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

## Extraction Status
- Status: auto_extracted
- Attributes found: 9
- Review needed: false

## Raw Source

```
9.3.23 Dot1ag MEP status  This ME is the read-only twin of the dot1ag MEP. Its purpose is to return information that may help  in system- or network-level troubleshooting. It is automatically created and deleted by the ONU at  the time its MEP is created or deleted.  As the reporter of ephemeral information, the dot1ag MEP status ME does not retain its attribute  values across initializations and is not included in MIB uploads.  Relationships  A dot1ag MEP status ME is associated with a dot1ag MEP ME.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. Through an  identical ID, this ME is implicitly linked to an instance of the dot1ag MEP ME.  (R) (mandatory) (2 bytes)  MEP MAC address : This attribute records the MEP 's MAC address . (R) (mandatory)  (6 bytes)  Fault notification generator state: This attribute records the current state of the MEP's fault  notification generator state machine . States are defined in clause 20.35 of  [IEEE 802.1ag]. 

---- page break ---- 1 Reset  2 Defect  3 Report defect  4 Defect reported  5 Defect clearing  (R) (mandatory) (1 byte)  Highest priority defect observed: This attribute records the highest priority defect observed  since the fault notification state machine was last in reset state. In increasing  priority order, possible values are as follows.  0 No defect observed  1 Received a CCM from a remote MEP in which the remote defect  indication (RDI) bit was set  2 Received a CCM from a remote MEP in which the port status or  interface status TLV reported an error  3 No CCMs received for at least 3.5 * CCM interval from at least one  remote MEP in the MA  4 Received invalid CCMs for at least 3.5 * CCM interval  5 Received CCMs for at least 3.5 * CCM interval that could be from  another MA  (R) (mandatory) (1 byte)  Current defects: This attribute is a bit field that signals several events of interest in real time.  Bit Meaning when set  1 (LSB) Another MEP in the same MA is currently transmitting an RDI.  2 A port status or interface status TLV received from another  MEP in the MA is currently indicating an error condition.  3 CCMs have not been received for at least 3.5 * CCM interval  from at least one of the expected remote MEPs.  4 Erroneous CCMs have been received for at least 3.5 * CCM  interval from at least one of the remote MEPs in this MA.  5 CCMs have been received for at least 3.5 * CCM interval from  an MEP that is not configured into the current MA.  6..8 Reserved  (R) (mandatory) (1 byte)  Last received errored CCM table: This attribute contains the most recently received CCM  that contributed to a defErrorCCM fault . If no such CCM has been received,  this attribute is null. The format of the CCM is defined in clause 21.6 of [IEEE  802.1ag]. (R) (mandatory) (N bytes, not to exceed 128)  Last received Xcon CCM table: This attribute contains the most recently received CCM that  contributed to a defXconCCM fault . If no such CCM has been received, this  attribute is null. (R) (mandatory) (N bytes, not to exceed 128)  Out of sequence CCMs count: This attribute records the number of out of sequence CCMs  received. When the counter is full, it rolls over to 0. (R) (optional) (4 bytes)  CCMs transmitted count: This attribute records the number of CCMs transmitted. It may be  used as the sequence number of transmitted CCMs. When the counter is full,  it rolls over to 0. (R) (mandatory) (4 bytes)  Unexpected LTRs count: This attribute records the number of unexpected LTRs received.  When the counter is full, it rolls over to 0. (R) (mandatory) (4 bytes) 

---- page break ---- Loopback replies ( LBRs) transmitted count: This attribute records the number of LBRs  transmitted. When the counter is full, it rolls over to 0 . (R) (mandatory)  (4 bytes)  Next loopback transaction identifier : This attribute is the value of the transaction number  sent in the next LBM to be transmitted . At ONU initialization, it should be  initialized to a random value. It increments with each LBM sent, and rolls over  when full. (R) (mandatory) (4 bytes)  Next linktrace transaction identifier : This attribute is the value of the transaction number  sent in the next LTM to be transmitted. It increments with each LTM sent, and  rolls over when full. (R) (mandatory) (4 bytes)  Actions  Get, get next  Notifications  None. This ME does not generate AVCs because its attributes change frequently in real time,  but are generally only of interest after the corresponding MEP declares an alarm.  
```
