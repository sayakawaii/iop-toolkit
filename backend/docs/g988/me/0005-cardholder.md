# Managed Entity

## Identity
- ME ID: 5
- ME Name: Cardholder
- Source Section: 9.1.5
- Source Page: 70

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Get, Set
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: Expected plug-in unit type
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 2
- Name: Expected port count
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 3
- Name: Expected equipment ID
- Size: 20 bytes
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 4
- Name: Actual equipment ID
- Size: 20 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 5
- Name: Protection profile pointer
- Size: 1 byte
- Format: needs_review
- Access: R
- Category: optional

### Attribute 6
- Name: Invoke protection switch
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: optional

## Extraction Status
- Status: auto_extracted
- Attributes found: 6
- Review needed: false

## Raw Source

```
9.1.5 Cardholder  The cardholder represents the fixed equipment slot configuration of the ONU. Each cardholder can  contain 0 or 1 circuit packs; the circuit pack models equipment information that can change over the  lifetime of the ONU, e.g., through replacement.  One instance of this ME exists for each physical slot in an ONU that has pluggable circuit packs. One  or more instances of this ME may also exist in an integrated ONU, to represent virtual slots. Instances  of this ME are created automatically by the ONU, and the status attributes are populated according to  data within the ONU itself.  Slot 0 is intended to be used only in an integrated ONU. If an integrated ONU is modelled with a  universal slot 0, it is recommended that it does not contain additional (non -zero) virtual slots. A  cardholder for virtual slot 0 is recommended.  There is potential for conflict in the semantics of the expected plug -in unit type, the expected port  count and the expected equipment ID, both when the slot is not populated and when a new circuit  pack is inserted. The expected plug-in unit type and the plug-in type mismatch alarm are mandatory,  although plug-and-play/unknown (circuit pack type 255) may be used as a way to minimize their  significance. It is recommended that an ONU deny the provisioning of inconsistent combinations of  expected equipment attributes.  When a circuit pack is plugged into a cardholder or when a cardholder is pre-provisioned to expect a  circuit pack of a given type, it may trigger the ONU to instantiate a number of MEs and update the  values of others, depending on the circuit pack type. The ONU may also delete a variety of other MEs  when a circuit pack is reprovisioned to not expect a circuit pack or to expect a circuit pack of a  different type. These actions are described in the definitions of the various MEs.  Expected equipment ID and expected port count are alternate ways to trigger the same  pre-provisioning effects. These tools may be useful if an ONU is prepared to accept more than one  circuit pack of a given type but with different port counts, or if a circuit pack is a hybrid that matches  none of the types in Table 9.1.5-1, but whose identification (e.g., part number) is known.  Relationships  An ONU may contain zero or more instances of the cardholder, each of which may contain  an instance of the circuit pack ME. The slot ID, real or virtual, is a fundamental identification  mechanism for MEs that bear some relationship to a physical location.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. The ONU  sets the first byte of this 2 byte identifier to:  0 if the ONU contains pluggable equipment modules  1 if the ONU is a single piece of integrated equipment.  The second byte of this identifier is the slot number. In integrated ONUs, this  byte may be used as a virtual slot or set to 0 to indicate a universal pseudo-slot.  Slot numbering schemes differ among vendors. It is only required that slot  numbers be unique across the ONU. Up to 254 equipment slots are supported  in the range 1..254 ( Note 1). The value 0 is reserved for possible use in an  integrated ONU to indicate a universal pseudo -slot. The value 255 is also  reserved. (R) (mandatory) (2 bytes)  NOTE 1 – Some xDSL MEs use the two MSBs of the slot number for other purposes.  An ONU that supports these services may have slot limitations or restrictions. 

---- page break ---- Actual plug -in unit type : This attribute is equal to the type of the circuit pack in the  cardholder, or 0 if the cardholder is empty. When the cardholder is populated,  this attribute is the same as the type attribute of the corresponding circuit pack  ME. Circuit pack types are defined in Table 9.1.5-1. (R) (mandatory) (1 byte)  The three following attributes permit the OLT to specify its intentions for any future equip ped  configuration of a slot. Once some or all of these are set, the ONU can proceed to instantiate circuit  pack and PPTP MEs, along with other predeterminable MEs, and allow the OLT to create related  discretionary MEs, thereby supporting service pre-provisioning.  Expected plug-in unit type: This attribute provisions the type of circuit pack for the slot. For  type coding, see  Table 9.1.5-1. The value 0 means that the cardholder is not  provisioned to contain a circuit pack. The value 255 means that the cardholder  is configured for plug -and-play. Upon ME instantiation, the ONU sets this  attribute to 0. For integrated interfaces, this attribute may be used to represent  the type of interface. (R, W) (mandatory) (1 byte)  Expected port count: This attribute permits the OLT to specify the number of ports it expects  in a circuit pack. Prior to provisioning by the OLT, the ONU initializes this  attribute to 0. (R, W) (optional) (1 byte)  Expected equipment ID: This attribute provisions the specific type of expected circuit pack.  This attribute applies only to ONUs that do not have integrated interfaces. In  some environments, this may contain the expected CLEI code. Upon ME  instantiation, the ONU sets this attri bute to all spaces. (R,  W) (optional)  (20 bytes)  Actual equipment ID: This attribute identifies the specific type of circuit pack, once it is  installed. This attribute applies only to ONUs that do not have integrated  interfaces. In some environments, this may include the CLEI code. When the  slot is empty or the equipment ID is not known, this attribute should be set to  all spaces. (R) (optional) (20 bytes)  Protection profile pointer: This attribute specifies an equipment protection profile that may  be associated with the cardholder. Its value is the least significant byte of the  ME ID of the equipment protection profile with which it is associated, or 0 if  equipment protection is not used. (R) (optional) (1 byte)  Invoke protection switch: The OLT may use this attribute to control equipment protection  switching. Code points have the
```
