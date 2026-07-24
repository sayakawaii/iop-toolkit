# Managed Entity

## Identity
- ME ID: 304
- ME Name: Dot1ag MEP CCM database
- Source Section: 9.3.24
- Source Page: 190

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Get
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

<!-- No attributes successfully parsed; see Raw Source section -->

## Extraction Status
- Status: auto_extracted
- Attributes found: 0
- Review needed: true

## Raw Source

```
9.3.24 Dot1ag MEP CCM database  This ME records the recent history of remote MEPs, as deduced by the local parent MEP. Because  records are of variable length, and are constantly updated, a separate attribute is defined for each  remote MEP. The dot1ag MEP CCM database is automatically created or deleted by the ONU at the  time an MEP is created or deleted.  As the reporter of ephemeral information, the dot1ag MEP CCM database ME does not retain its  attribute values across initializations and is not included in MIB uploads.  Relationships  A dot1ag MEP CCM database ME is associated with a dot1ag MEP ME.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. Through an  identical ID, this ME is implicitly linked to an instance of the dot1ag MEP  ME. (R) (mandatory) (2 bytes)  Each of the following RMEP database table attributes records information for one of the possible  remote MEPs. It is expected that there will be only one remote MEP per MA in G-PON applications,  but the ME is defined in a way that permits several RMEPs. The  optional attributes are instantiated  by the ONU when additional remote MEPs are provisioned on the local MEP. Remote MEP records  appear in no particular order, and the order is not guaranteed to persist across ONU initializations.  NOTE – Although each attribute is shown with a single indeterminate length N, it is understood that the length  of each attribute varies in real time, and independently from the length of the other attributes.  RMEP 1 database table: (R) (mandatory) (N bytes)  RMEP 2 database table: (R) (optional) (N bytes)  RMEP 3 database table: (R) (optional) (N bytes)  RMEP 4 database table: (R) (optional) (N bytes)  RMEP 5 database table: (R) (optional) (N bytes)  RMEP 6 database table: (R) (optional) (N bytes)  RMEP 7 database table: (R) (optional) (N bytes) 

---- page break ---- RMEP 8 database table: (R) (optional) (N bytes)  RMEP 9 database table: (R) (optional) (N bytes)  RMEP 10 database table: (R) (optional) (N bytes)  RMEP 11 database table: (R) (optional) (N bytes)  RMEP 12 database table: (R) (optional) (N bytes)  Each attribute is a record that comprises the following fields.  RMep identifier: The MEP ID of the remote MEP. (2 bytes)  RMep state: An enumeration with the following meaning (1 byte).  1 Idle. Momentary state during reset.  2 Start. The timer has not expired since the state machine was  reset, but no valid CCM has yet been received.  3 Failed. The timer has expired since the state machine was reset  and since a valid CCM was received.  4 Ok. The timer has not expired since a valid CCM was received.  Failed-ok time: A timestamp, the value of the local ONU's SysUpTime at  which the remote MEP state last entered either the failed or ok state.  SysUpTime is a count of 10 ms intervals since ONU initialization. The  value is 0 if it has not been in either of these states since ONU  initialization. (4 bytes)  MAC address: The MAC address of the remote MEP. If no CCM has been  received from the remote MEP, this field has the value 0. (6 bytes)  RDI: Boolean indicating whether the RDI bit in the most recently received  CCM was set. (1 byte)  Port status : The port status from the most recently received CCM, as  defined in clause 21.5.4 of [IEEE 802.1ag]. The absence of a received  port status TLV is indicated by the value 0. (1 byte)  Interface status : The interface status from the most recently received  CCM, as defined in clause 21.5.5 of [IEEE 802.1ag]. The absence of a  received interface status TLV is indicated by the value 0. (1 byte)  Sender ID TLV: This is the actual sender ID TLV from the most recently  received CCM, as defined in clause 21.5.3 of [IEEE  802.1ag]. The  absence of a received sender ID TLV is indicated by a single byte of  value 0. (M bytes)  Actions  Get, get next  Notifications  None. The MEP CCM database table attributes do not generate AVCs because they change  constantly in real time, usually in ways that are of no immediate interest.  
```
