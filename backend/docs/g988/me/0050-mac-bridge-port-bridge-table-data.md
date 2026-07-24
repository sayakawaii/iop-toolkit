# Managed Entity

## Identity
- ME ID: 50
- ME Name: MAC bridge port bridge table data
- Source Section: 9.3.8
- Source Page: 148

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
9.3.8 MAC bridge port bridge table data  This ME reports status data associated with a bridge port. The ONU automatically creates or deletes  an instance of this ME upon the creation or deletion of a MAC bridge port configuration data.  Relationships  An instance of this ME is associated with an instance of a MAC bridge port configuration data  ME.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. Through an  identical ID, this ME is implicitly linked to an instance of the MAC bridge port  configuration data ME. (R) (mandatory) (2 bytes)  Bridge table: This attribute lists known MAC DAs, whether they are learned or statically  assigned, whether packets that have them as DAs are filtered or forwarded, and  their ages. Each entry contains:   – Information (2 bytes);   – MAC address (6 bytes).  The information bits are assigned as described as follows.  Bit Name Setting  1 (LSB) Filter/forward 0: forward    1: filter  2 Reserved 0  3 Dynamic/static 0: this entry is statically assigned    1: this entry is dynamically learned 

---- page break ---- 4 Reserved 0  16..5 Age Age in seconds (1..4095)  Upon ME instantiation, this attribute is an empty list. (R) (mandatory)  (8 * M bytes, where M is the number of entries in the list.)  Actions  Get, get next  Notifications  None.  
```
