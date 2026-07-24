# Managed Entity

## Identity
- ME ID: 2
- ME Name: ONU data
- Source Section: 9.1.3
- Source Page: 65

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Get, Set
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: MIB data sync
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: mandatory

## Extraction Status
- Status: auto_extracted
- Attributes found: 1
- Review needed: false

## Raw Source

```
9.1.3 ONU data  This ME models the MIB itself. Clause I.1.3 explains the use of this ME with respect to MIB  synchronization.  The ONU automatically creates an instance of this ME, and  updates the associated attributes  according to data within the ONU itself.  Relationships  One instance of this ME is contained in an ONU.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. There is only  one instance, number 0. (R) (mandatory) (2 bytes)  MIB data sync: This attribute is used to check the alignment of the MIB of the ONU with the  corresponding MIB in the OLT. MIB data sync relies on this attribute, which  is a sequence number that can be checked by the OLT to see if the MIB  snapshots for the OLT and ONU mat ch. Refer to clause I.1.2.1 for a detailed  description of this attribute. Upon ME instantiation, the ONU sets this attribute  to 0. (R, W) (mandatory) (1 byte)  Actions  Get, set  Get all alarms: Latch a snapshot of the current alarm statuses of all MEs and reset the alarm  message counter.  Get all alarms next : Get the latched alarm status of the next ME(s) within the current  snapshot.  MIB reset: Reset the MIB data sync attribute to 0 and reset the MIB of the ONU to its  default. The default MIB comprises those MEs that are designated mandatory  in the corresponding Recommendation, along with other auto -created MEs  whose existence is implicit in the architecture or physical configuration of the  ONU.  For G-PON applications, the minimum default MIB comprises one instance of  the ONU-G ME pair, one instance of the ONU data ME, and two instances of  the software image ME.  MIB upload: Latch a snapshot (i.e., copy) of the current MIB. Not every ME or every  attribute is included in an MIB upload. Table attributes are excluded. Only the  control block attributes of PM MEs are uploaded. Other MEs and attributes,  such as the PPTP for the local craft terminal  (LCT), are excluded as  documented in their specific definitions. 

---- page break ---- MIB upload next : Get the latched attribute values of the next ME(s) within the current  snapshot.  Notifications  None.  
```
