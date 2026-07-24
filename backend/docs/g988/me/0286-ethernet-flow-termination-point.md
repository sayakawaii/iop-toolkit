# Managed Entity

## Identity
- ME ID: 286
- ME Name: Ethernet flow termination point
- Source Section: 9.8.9
- Source Page: 372

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Create, Delete, Get, Set
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: Destination MAC
- Size: 6 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 2
- Name: Source MAC
- Size: 6 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 3
- Name: Tag policy
- Size: 1 byte
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 4
- Name: TCI
- Size: 2 bytes
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 5
- Name: Loopback
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: mandatory

## Extraction Status
- Status: auto_extracted
- Attributes found: 5
- Review needed: false

## Raw Source

```
9.8.9 Ethernet flow termination point  The Ethernet flow TP contains the attributes necessary to originate and terminate Ethernet frames in  the ONU. It is appropriate when transporting pseudowire services via layer  2. Instances of this ME  are created and deleted by the OLT.  Relationships  One Ethernet flow TP ME exists for each distinct pseudowire service that is transported via  layer 2.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. Through an  identical ID, this ME is implicitly linked to a pseudowire TP ME. (R,  set-by-create) (mandatory) (2 bytes)  Destination MAC: This attribute specifies the destination MAC address of upstream Ethernet  frames. (R, W, set-by-create) (mandatory) (6 bytes)  Source MAC : This attribute specifies the near -end MAC address. It is established by  non-OMCI means (e.g., factory programmed into ONU flash memory) and is  included here for information only. (R) (mandatory) (6 bytes)  Tag policy: This attribute specifies the tagging policy to be applied to upstream Ethernet  frames.  0 untagged frame  1 tagged frame 

---- page break ---- (R, W, set-by-create) (mandatory) (1 byte)  TCI: If the tag policy calls for tagging of upstream Ethernet frames, this attribute  specifies the tag control information, which includes the VLAN tag, P bits and  CFI bit. (R, W) (optional) (2 bytes)  Loopback: This attribute sets the loopback configuration as follows.  0 No loopback  1 Loopback of downstream traffic at MAC client  (R, W) (mandatory) (1 byte)  Actions  Create, delete, get, set  Notifications  None.  
```
