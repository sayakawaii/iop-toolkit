# Managed Entity

## Identity
- ME ID: 7
- ME Name: Software image
- Source Section: 9.1.4
- Source Page: 66

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Get
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: Version
- Size: 14 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 2
- Name: Is committed
- Size: 1 byte
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 3
- Name: Is valid
- Size: 1 byte
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 4
- Name: Product code
- Size: 25 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 5
- Name: Image hash
- Size: 16 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 6
- Name: Start download
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 7
- Name: Version
- Size: 14 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 8
- Name: Is committed
- Size: 1 byte
- Format: needs_review
- Access: R
- Category: optional

### Attribute 9
- Name: Is active
- Size: 1 byte
- Format: needs_review
- Access: R
- Category: optional

### Attribute 10
- Name: Product code
- Size: 25 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 11
- Name: Image hash
- Size: 16 bytes
- Format: needs_review
- Access: R
- Category: optional

## Extraction Status
- Status: auto_extracted
- Attributes found: 11
- Review needed: false

## Raw Source

```
9.1.4 Software image  This ME models an executable software image stored in the ONU (documented here as its  fundamental usage). It may also be used to represent an opaque vendor-specific file (vendor-specific  usage).  Fundamental usage  The ONU automatically creates two instances of this ME upon the creation of each ME that contains  independently manageable software, either the ONU itself or an individual circuit pack. It populates  ME attributes according to data within the ONU or the circuit pack.  Some pluggable equipment may not contain software. Others may contain software that is  intrinsically bound to the ONU's own software image. No software image ME need exist for such  equipment, though it may be convenient for the ONU to create them to support software version audit  from the OLT. In this case, the dependent MEs would support only the get action.  A slot may contain various equipment over its lifetime, and if software image MEs exist, the ONU  must automatically create and delete them as the equipped configuration changes. The identity of the  software image is tied to the cardholder.   When an ONU controller packs are duplicated, each can be expected to contain two software image  MEs, managed through reference to the individual controller packs themselves. When this occurs, the  ONU should not have a global pair of software images MEs (instance 0), since an action (download,  activate, commit) directed to instance 0 would be ambiguous.  Relationships  Two instances of the software image ME are associated with each instance of the ONU or  cardholder whose software is independently managed.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. The first byte  indicates the physical location of the equipment hosting the software image,  either the ONU (0) or a cardholder (1..254). The second byte distinguishes  between the two software image ME instances (0..1). (R) (mandatory)  (2 bytes)  Version: This string attribute identifies the version of the software. (R) (mandatory)  (14 bytes)  Is committed: This attribute indicates whether the associated software image is committed  (1) or uncommitted (0). By definition, the committed software image is loaded  and executed upon reboot of the ONU or circuit pack. During normal  operation, one software image is always committed, while the other is  uncommitted. Under no circumstances are both software images allowed to be  committed at the same time. On the other hand, both software images could be  uncommitted at the same time if both were invalid. Upon ME instanti ation,  instance 0 is initialized to committed, while instance 1 is initialized to  uncommitted (i.e., the ONU ships from the factory with image 0 committed).  (R) (mandatory) (1 byte) 

---- page break ---- Is active: This attribute indicates whether the associated software image is active (1) or  inactive (0). By definition, the active software image is one that is currently  loaded and executing in the ONU or circuit pack. Under normal operation, one  software image is  always active while the other is inactive. Under no  circumstances are both software images allowed to be active at the same time.  On the other hand, both software images could be inactive at the same time if  both were invalid. (R) (mandatory) (1 byte)  Is valid: This attribute indicates whether the associated software image is valid (1) or  invalid (0). By definition, a software image is valid if it has been verified to be  an executable code image. The verification mechanism is not subject to  standardization; however, it should include at least a data integrity check [e.g.,  a cyclic redundancy check ( CRC)] of the entire code image. Upon ME  instantiation or software download completion, the ONU validates the  associated code image and sets this attribute acc ording to the result. (R)  (mandatory) (1 byte)  Product code: This attribute provides a way for a vendor to indicate product code information  on a file. It is a character string, padded with trailing nulls if it is shorter than  25 bytes. (R) (optional) (25 bytes)  Image hash: This attribute is an MD5 hash of the software image. It is computed at  completion of the end download action. (R) (optional) (16 bytes)  Actions  Get  Software upgrade is described in clause I.3. All of  the following actions are mandatory for ONUs  with remotely manageable software.  Start download : Initiate a software download sequence. This action is valid only for a  software image instance that is neither active nor committed.  Download section: Download a section of a software image. This action is valid only for a  software image instance that is currently being downloaded (image 1 in  state S2, image 0 in state S2').  End download: Signal the completion of a download image sequence, providing both CRC  and version information for final verification. This action is valid only for a  software image instance that is currently being downloaded (image 1 in  state S2, image 0 in state S2').  Activate image: Load/execute a software image. When this action is applied to a software  image that is currently inactive, execution of the current code image is  suspended, the associated software image is loaded from non-volatile memory,  and execution of this new code image is initiated ( i.e., the associated entity  reboots on the previously inactive image). When this action is applied to a  software image that is already active, a soft restart is performed. The software  image is not reloaded from non -volatile memory; the current volatile code  image is simply restarted. This action is only valid for a valid software image.  Commit image: Set the is committed attribute value to 1 for the target software image ME  and set the is committed attribute value to 0 for the other software image. This  causes the committed software image to be loaded and executed by the boot  code upon subsequent start-ups. This action is only applicable 
```
