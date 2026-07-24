# Managed Entity

## Identity
- ME ID: 160
- ME Name: Equipment extension package
- Source Section: 9.1.9
- Source Page: 84

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Get, Set
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: Environmental sense
- Size: 2 bytes
- Format: needs_review
- Access: R, W
- Category: optional

## Extraction Status
- Status: auto_extracted
- Attributes found: 1
- Review needed: false

## Raw Source

```
9.1.9 Equipment extension package  This ME supports optional extensions to circuit pack MEs. If the circuit pack supports these features,  the ONU creates and deletes this ME along with its associated real or virtual circuit pack.  Relationships  An equipment extension package may be contained by an ONU-G or cardholder.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. Through an  identical ID, this ME is implicitly linked to an instance of the ONU -G or  cardholder. (R) (mandatory) (2 bytes)  Environmental sense: This attribute provisions an ONU that supports external sense points,  e.g., physical security detectors at an enclosure. Each pair of bits is defined as  follows.  00 Sense point disabled (default)  01 Report contact closure  10 Report contact open  11 Sense point disabled (same as 00)  If the byte is represented in binary as 0B hhgg ffee ddcc bbaa, bits hh  correspond to sense point 1, while bits aa correspond to sense point 8. (R,  W)  (optional) (2 bytes)  NOTE – Some specific sense point applications are already defined on the ONU -G  ME. It is the vendor's choice how to configure and report sense points that appear both  generically and specifically. 

---- page break ---- Contact closure output : This attribute provisions an ONU that supports external contact  closure outputs, e.g., sump pump or air conditioner activation at an ONU  enclosure. A contact point is said to be released when it is not energized.  Whether this corresponds to an open or a closed electrical circuit depends on  the ONU's wiring options. Upon ONU initialization, all contact points should  go to the released state.  If the byte is represented in binary as 0B hhgg ffee ddcc bbaa, bits hh  correspond to contact output point 1, while bits aa correspond to contact output  point 8.  On write, the bits of this attribute have the following meaning.  0x No change to contact output point state  10 Release contact output point  11 Operate contact output point  On read, the left bit in each pair should be set to 0 at the ONU and ignored at  the OLT. The right bit indicates a released output point with 0 and an operated  contact point with 1. (R, W) (optional) (2 bytes)  Actions  Get, set  Notifications  Alarm  Alarm  number  Alarm Description  0 Reserved   1 Sense point 1 Environmental sense point 1 active  2 Sense point 2 Environmental sense point 2 active  3 Sense point 3 Environmental sense point 3 active  4 Sense point 4 Environmental sense point 4 active  5 Sense point 5 Environmental sense point 5 active  6 Sense point 6 Environmental sense point 6 active  7 Sense point 7 Environmental sense point 7 active  8 Sense point 8 Environmental sense point 8 active  9..207 Reserved   208..223 Vendor-specific alarms Not to be standardized  
```
