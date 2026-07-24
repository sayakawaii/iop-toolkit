# Managed Entity

## Identity
- ME ID: 49
- ME Name: MAC bridge port filter table data
- Source Section: 9.3.6
- Source Page: 146

## Classification
- Access: needs_review
- Type: needs_review
- Actions: needs_review
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
Notifications  None.  9.3.6 MAC bridge port filter table data  This ME organizes data associated with a bridge port. The ONU automatically creates or deletes an  instance of this ME upon the creation or deletion of a MAC bridge port configuration data ME.  NOTE – The OLT should disable the learning mode in the MAC bridge service profile before writing to the  MAC filter table.  Relationships  An instance of this ME is associated with an instance of a MAC bridge port configuration data  ME.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. Through an  identical ID, this ME is implicitly linked to an instance of the MAC bridge port  configuration data ME. (R) (mandatory) (2 bytes)  MAC filter table : This attribute lists MAC addresses associated with the bridge port, each  with an allow/disallow forwarding indicator for traffic flowing out of the  bridge port. Additionally, the forwarding action may be based on a MAC  source or DA. In this way, upstream traffic is filtered on ANI-side bridge ports,  and downstream traffic is filtered on UNI -side bridge ports. The setting of an  entry with a forward action implies that all other addresses are filtered.  Conversely, the setting of an en try with a filter action implies that all other  addresses are forwarded. The behaviour is unspecified if forward and filter  actions are mixed.  Each entry contains:  – the entry number, an index into this attribute list (1 byte);  – filter byte (1 byte);  – MAC address (6 bytes).  The bits of the filter byte are assigned as follows.  Bit Name Setting  1 (LSB) Filter/forward 0: forward    1: filter    2  0: MAC DAs    1: MAC source addresses    3..6 Reserved 0    7..8 Add/remove 10: Clear entire table (set operation)    00: Remove this entry (set operation)    01: Add this entry  Upon ME instantiation, the ONU sets this attribute to an empty table.  (R, W) (Mandatory) (8N bytes, where N is the number of entries in the list)  Actions  Get, get next, set  Set table (optional) 

---- page break ---- Notifications  None.  
```
