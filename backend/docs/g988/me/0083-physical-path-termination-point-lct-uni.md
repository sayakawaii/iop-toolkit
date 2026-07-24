# Managed Entity

## Identity
- ME ID: 83
- ME Name: Physical path termination point LCT UNI
- Source Section: 9.13.3
- Source Page: 456

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Get, Set
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: Administrative state
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
9.13.3 Physical path termination point LCT UNI  This ME models debug access to the ONU from any physical or logical port, for example, via a  dedicated LCT UNI, via ordinary subscriber UNIs, or via the IP host config ME.  The ONU automatically creates an instance of this ME per port:  • when the ONU has an LCT port built into its factory configuration;  • when a cardholder is provisioned to expect a circuit pack of the LCT type;  • when a cardholder provisioned for plug-and-play is equipped with a circuit pack of the LCT  type;  NOTE – The installation of a plug -and-play card may indicate the presence of LCT ports via  equipment ID as well as its type, and indeed may cause the ONU to instantiate a port-mapping package  that specifies LCT ports.  • when the ONU supports debug access through some other physical or logical means.  The ONU automatically deletes an instance of this ME when a cardholder is neither provisioned to  expect an LCT circuit pack, nor is it equipped with an LCT circuit pack, or if the ONU is reconfigured  in such a way that it no longer supports debug access.  LCT instances are not reported during an MIB upload.  Relationships  An instance of this ME is associated with an instance of a real or virtual circuit pack ME  classified as an LCT type. An instance of this ME may also be associated with the ONU as a  whole, if the ONU supports debug access through means other than a dedicated physical LCT  port.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. This 2 byte  number indicates the physical position of the UNI. The first byte is the slot ID  (defined in clause 9.1.5). The second byte is the port ID, with the range 1..255.  If the LCT UNI is associated with the ONU as a whole, its ME ID should be  0. (R) (mandatory) (2 bytes)  Administrative state: This attribute locks (1) and unlocks (0) the functions performed by this  ME. Administrative state is described generically in clause  A.1.6. The LCT  has additional administrative state behaviour. When the administrative state is  set to lock, debug access through all physical or logical means is blocked,  except that the operation of a possible ONU remote debug ME is not affected.  Administrative lock of ME instance 0 overrides administrative lock of any  other PPTP LCT UNIs that may exist. (R, W) (mandatory) (1 byte) 

---- page break ---- Actions  Get, set  Notifications  None.  
```
