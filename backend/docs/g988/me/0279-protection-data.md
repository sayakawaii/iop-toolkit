# Managed Entity

## Identity
- ME ID: 279
- ME Name: Protection data
- Source Section: 9.1.10
- Source Page: 85

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Get, Set, Create, Delete
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: Working ANI-G pointer
- Size: 2 bytes
- Format: needs_review
- Access: R, W if applicable, set-by-create if applicable
- Category: mandatory

### Attribute 2
- Name: Protection ANI-G pointer
- Size: 2 bytes
- Format: needs_review
- Access: R, W if applicable, set-by-create if applicable
- Category: mandatory

### Attribute 3
- Name: Protection type
- Size: 1 byte
- Format: needs_review
- Access: R, W if applicable, set-by-create if applicable
- Category: mandatory

### Attribute 4
- Name: Revertive ind
- Size: 1 byte
- Format: needs_review
- Access: R, W if applicable, set-by-create if applicable
- Category: mandatory

### Attribute 5
- Name: Wait to restore time
- Size: 2 bytes
- Format: needs_review
- Access: RWSC if applicable
- Category: mandatory

### Attribute 6
- Name: Switching guard time
- Size: 2 bytes
- Format: needs_review
- Access: RWS C if applicable
- Category: optional

## Extraction Status
- Status: auto_extracted
- Attributes found: 6
- Review needed: false

## Raw Source

```
9.1.10 Protection data  This ME models the capability and parameters of PON protection. An ONU that supports PON  protection automatically creates one instance of this ME. A chassis -based RE could have the  capability of protecting a number of PONs, possibly by way of circuit packs configured in arbitrary  (rather than predefined) slots. Protection data MEs in a multi-PON RE may therefore be auto-created  or created by the OLT, depending on the RE's architecture, and the ANI -G pointers may be either  populated by the ONU itself (read-only) or configured by the OLT (RWSC). Likewise, the nature of  protection may be set read-only by the ONU's architecture or may be settable by the OLT.  NOTE 1 – Equipment protection is modelled with the equipment protection profile and cardholder MEs. 

---- page break ---- NOTE 2 – For ONUs that implement RE functions, this ME can be used to describe OMCI protection, RE  interface R'/S' protection, or both. For R'/S' protection, the protection type must be 1:1 without extra traffic,  because the switching is done on a link -by-link basis, and the protection link is in cold standby mode. The  instance that pertains to OMCI protection has ME ID = 0.  Relationships  One instance of this ME is associated with two instances of the ANI -G, RE ANI -G or RE  upstream amplifier. One of the ANI MEs represents the working side; the other represents the  protection side.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. If there is  more than one protection data ME, they are numbered in ascending order from  0. (R, set-by-create if applicable) (mandatory) (2 bytes)  Working ANI-G pointer: This attribute points to the ANI -G, RE ANI -G or RE upstream  amplifier ME that represents the working side of a protected PON. (R,  W if  applicable, set-by-create if applicable) (mandatory) (2 bytes)  NOTE 3 – It is possible, and indeed likely, that an ANI-G will have the same ME ID  as the RE ANI -G or even the RE upstream amplifier that supports its physical PON  interface. The ANI -G represents the embedded ONU that terminates the OMCC.  Since it is not e xpected that protection of management communications will be  implemented independently from protection of the optical layer, the ambiguity is not  expected to cause a problem.  Protection ANI-G pointer: This attribute points to the ANI-G, RE ANI-G or RE upstream  amplifier ME that represents the protection side of a protected PON.  (R, W if  applicable, set-by-create if applicable) (mandatory) (2 bytes)  Protection type: This attribute indicates the type of PON protection. Valid values are:  0 1+1 protection  1 1:1 protection without extra traffic  2 1:1 protection with ability to support extra traffic  (R, W if applicable, set-by-create if applicable) (mandatory) (1 byte)  Revertive ind: This attribute indicates whether protection is revertive (1) or non-revertive (0).  (R, W if applicable, set-by-create if applicable) (mandatory) (1 byte)  Wait to restore time: This attribute specifies the time, in seconds, to wait after a fault clears  before switching back to the working path . Upon ME instantiation, the ONU  sets this attribute to 3 s. (RWSC if applicable) (mandatory) (2 bytes)  Switching guard time : This attribute specifies the time, in milliseconds, to wait after the  detection of a fault before performing a protection switch. Specification of a  default value for this attribute is outside the scope of this Recommendation, as  it is normally handled through supplier -operator negotiations. (RWS C if  applicable) (optional) (2 bytes)  Actions  Get, set  If applicable: create, delete  Notifications  None. 

---- page break ---- 
```
