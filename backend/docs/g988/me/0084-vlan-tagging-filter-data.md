# Managed Entity

## Identity
- ME ID: 84
- ME Name: VLAN tagging filter data
- Source Section: 9.3.11
- Source Page: 152

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Create, Delete, Get, Set
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: VLAN filter list
- Size: 24 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 2
- Name: Forward operation
- Size: 1 byte
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 3
- Name: Number of entries
- Size: 1 byte
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

## Extraction Status
- Status: auto_extracted
- Attributes found: 3
- Review needed: false

## Raw Source

```
9.3.11 VLAN tagging filter data  This ME organizes data associated with VLAN tagging. Instances of this ME are created and deleted  by the OLT.  Relationships  An instance of this ME is associated with an instance of a MAC bridge port configuration data  ME. By definition, tag filtering occurs closer to the MAC bridge than the tagging operation.  Schematically, the ordering of the functions is as given in Figure 9.3.11-1:  G.988(12)_F9.3.11-1 ANI Tag operation Tag filtering Bridging UNITag operationTag filtering   Figure 9.3.11-1  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. Through an  identical ID, this ME is implicitly linked to an instance of the MAC bridge port  configuration data ME. (R, set-by-create) (mandatory) (2 bytes)  VLAN filter list: This attribute is a list of provisioned tag control information (TCI) values  for the bridge port. A TCI, comprising user priority, canonical format indicator  (CFI) and virtual local area network identifier (VID), is represented by 2 bytes.  This attribute supports up to 12 VLAN entries. The first N are valid, where N  is given by the number of entries attribute. (R,  W, set-by-create) (mandatory)  (24 bytes)  Forward operation : When a frame passes through the MAC bridge port, it is processed  according to the operation specified by this attribute, in accordance with   Table 9.3.11-1. Figure 9.3.11-3 illustrates the treatment of frames according to  the provisioned action possibilities. Tagged and untagged frames are treated  separately, but both in accordance with Figure 9.3.11-3. While all forwarding  operations are plausible, only actions 0x10 and 0x12 are necessary to construct  a VLAN mapper and an 802.1p mapper, respectively. (R,  W, set-by-create)  (mandatory) (1 byte)  Table 9.3.11-1 – Forward operation attribute values  Forward  operation  Type of received frame  Tagged Untagged  0x00 Bridging (a) (no investigation) Bridging (a)  0x01 Discarding (c) Bridging (a)  0x02 Bridging (a) (no investigation) Discarding (c)  0x03 Action (h) (VID investigation) Bridging (a)  0x04 Action (h) (VID investigation) Discarding (c)  0x05 Action (g) (VID investigation) Bridging (a)  0x06 Action (g) (VID investigation) Discarding (c)  0x07 Action (h) (user priority investigation) Bridging (a)  0x08 Action (h) (user priority investigation) Discarding (c)  0x09 Action (g) (user priority investigation) Bridging (a) 

---- page break ---- Table 9.3.11-1 – Forward operation attribute values  Forward  operation  Type of received frame  Tagged Untagged  0x0A Action (g) (user priority investigation) Discarding (c)  0x0B Action (h) (TCI investigation) Bridging (a)  0x0C Action (h) (TCI investigation) Discarding (c)  0x0D Action (g) (TCI investigation) Bridging (a)  0x0E Action (g) (TCI investigation) Discarding (c)  0x0F Action (h) (VID investigation) Bridging (a)  0x10 Action (h) (VID investigation) Discarding (c)  0x11 Action (h) (user priority investigation) Bridging (a)  0x12 Action (h) (user priority investigation) Discarding (c)  0x13 Action (h) (TCI investigation) Bridging (a)  0x14 Action (h) (TCI investigation) Discarding (c)  0x15 Bridging (a) (no investigation) Discarding (c)  0x16 Action (j) (VID investigation) Bridging (a)  0x17 Action (j) (VID investigation) Discarding (c)  0x18 Action (j) (user priority investigation) Bridging (a)  0x19 Action (j) (user priority investigation) Discarding (c)  0x1A Action (j) (TCI investigation) Bridging (a)  0x1B Action (j) (TCI investigation) Discarding (c)  0x1C Action (h) (VID investigation) Bridging (a)  0x1D Action (h) (VID investigation) Discarding (c)  0x1E Action (h) (user priority investigation) Bridging (a)  0x1F Action (h) (user priority investigation) Discarding (c)  0x20 Action (h) (TCI investigation) Bridging (a)  0x21 Action (h) (TCI investigation) Discarding (c)  Table 9.3.11-1 contains duplicate entries due to simplification of the original  set of actions.  Table 9.3.11-1 and the actions listed are discussed in detail in the following.  Number of entries : This attribute specifies the number of valid entries in the VLAN filter  list. (R, W, set-by-create) (mandatory) (1 byte)  Actions  Create, delete, get, set  Notifications  None. 

---- page break ---- Supplementary explanation  This section explains the actions specified in the forward operation attribute.  The format of an Ethernet frame for VLAN services is described in [IEEE 802.1Q] and  [IEEE 802.1ad]. See Figure 9.3.11-2.  G.988(12)_F9.3.11-2 DA SA Length/ type MAC SDU TPID TCI 802.1p (3 bits) CFI (1 bit) VID (802.1Q) – VLAN ID (tag) (12 bits) User priority TCI: Tag control information   Figure 9.3.11-2 – Format of an Ethernet frame for VLAN services  a) Basic MAC bridge operation : All frames are accepted into the MAC bridging entity.  Egress frames are forwarded from this port if either a) the frame 's MAC DA is listed in  the MAC bridge port bridge table data for this port or b) the frame's DA does not appear  in the MAC bridge port bridge table data for any port (flooding). The content s of the  VLAN filter list attribute are not meaningful. See Figure 9.3.11-3.  NOTE – Action (a) on a given port may imply egress flooding of a frame from other ports of the  bridge. The possible VLAN tagging filter data MEs attached to the other ports override this action  however, so the frame is only transmitted from another port if it also satisfies the forward  operation attribute value established on that port.  G.988(12)_F9.3.11-3 802.1 bridging entity MAC bridge service profile Modelled in OMCI as All frames accepted DA match or flooded frames only Only frames whose TCI does not match DA match or flood Only frames whose TCI matches Functionality of (extended) VLAN  tagging operation  configuration data MEs, if any  Action a: accept VLAN filter list  = don’t care Port P1 P2 Action c: discard VLAN filter list  = don’t care P3 P4 Action g:   discard on TCI Action h:   accept on TCI,   else discard All
```
