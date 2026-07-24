# Managed Entity

## Identity
- ME ID: 99
- ME Name: Physical path termination point xDSL UNI part 2
- Source Section: 9.7.1
- Source Page: 243

## Classification
- Access: needs_review
- Type: needs_review
- Actions: needs_review
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: Loopback configuration
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 2
- Name: Administrative state
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 3
- Name: Operational state
- Size: 1 byte
- Format: needs_review
- Access: R
- Category: optional

### Attribute 4
- Name: xDSL line configuration profile
- Size: 2 bytes
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 5
- Name: xDSL subcarrier masking downstream profile
- Size: 2 bytes
- Format: needs_review
- Access: R, W
- Category: mandatory

## Extraction Status
- Status: auto_extracted
- Attributes found: 5
- Review needed: true

## Raw Source

```
9.7.1 Physical path termination point xDSL UNI part 1  This ME represents the point where physical paths terminate on an xDSL CO modem (xTU-C). The 

---- page break ---- xDSL ME family2 is used for ADSL VDSL2 and FAST services. A legacy family of VDSL MEs  remains valid for ITU-T G.993.1 VDSL, if needed. It is documented in [ITU-T G.983.2].  The ONU automatically creates an instance of this ME per port:  • when the ONU has xDSL ports built into its factory configuration;  • when a cardholder is provisioned to expect a circuit pack of the xDSL type;  • when a cardholder provisioned for plug-and-play is equipped with a circuit pack of the xDSL  type. Note that the installation of a plug -and-play card may indicate the presence of xDSL  ports via equipment ID as well as its type, and indeed may cause the ONU to instantiate a  port-mapping package that specifies xDSL ports.  The ONU automatically deletes instances of this ME when a cardholder is neither provisioned to  expect an xDSL circuit pack, nor is it equipped with an xDSL circuit pack.  Relationships  An instance of this ME is associated with each instance of a real or pre -provisioned xDSL  port.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. This 2 byte  number indicates the physical position of the UNI. The six LSBs of the first  byte are the slot ID, defined in clause 9.1.5. The two MSBs indicate the channel  number in some of the implicitly linked MEs, and must be 0 in the PPTP itself.  This reduces the possible number of physical slots to 64. The second byte is  the port ID, with the range 1..255. (R) (mandatory) (2 bytes)  Loopback configuration : This attribute represents the loopback configuration of this  physical interface.  0 No loopback  1 Loopback2 – a loopback at the ONU toward s the OLT. The OLT can  execute a physical level loopback test after loopback2 is set.  Upon ME instantiation, the ONU sets this attribute to 0. (R,  W) (mandatory)  (1 byte)  Administrative state: This attribute locks (1) and unlocks (0) the functions performed by this  ME. Administrative state is further described in clause A.1.6. (R,  W)  (mandatory) (1 byte)  Operational state : This attribute indicates whether the ME is capable of performing its  function. Valid values are enabled (0) and disabled (1). (R) (optional) (1 byte)  xDSL line configuration profile : This attribute points to an instance of the xDSL line  configuration profiles (part 1, 2 and 3) MEs, and if necessary, also to VDSL2  line configuration extensions (1 and 2) MEs, also to vectoring line  configuration extension MEs. Upon ME instantiation, the ONU sets this  attribute to 0, a null pointer. (R, W) (mandatory) (2 bytes)  xDSL subcarrier masking downstream profile : This attribute points to an instance of the  xDSL subcarrier masking downstream profile ME. Upon ME instantiation, the  ONU sets this attribute to 0, a null pointer. (R, W) (mandatory) (2 bytes)  ____________________  2  The xDSL MEs include the ITU-T G.992 family as well as ITU-T G.993.2 VDSL2, but not ITU-T G.993.1 VDSL.
```
