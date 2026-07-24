# Managed Entity

## Identity
- ME ID: 168
- ME Name: VDSL2 line inventory and status data part 1
- Source Section: 9.7.16
- Source Page: 285

## Classification
- Access: needs_review
- Type: needs_review
- Actions: needs_review
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: VDSL2 transmission system capability xTU -C
- Size: 1 byte
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 2
- Name: VDSL2 transmission system
- Size: 1 byte
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 3
- Name: VDSL2 profile
- Size: 1 byte
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 4
- Name: VDSL2 US0 PSD mask
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 5
- Name: ACTSNRMODEds
- Size: 1 byte
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 6
- Name: HLINGds
- Size: 1 byte
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 7
- Name: HLOGGds
- Size: 1 byte
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 8
- Name: QLNGds
- Size: 1 byte
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 9
- Name: SNRGds
- Size: 1 byte
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 10
- Name: MREFPSDds table
- Size: 1 byte
- Format: needs_review
- Access: R
- Category: optional

## Extraction Status
- Status: auto_extracted
- Attributes found: 10
- Review needed: true

## Raw Source

```
Actions  Get, get next  Notifications  None.  9.7.16 VDSL2 line inventory and status data part 1  This ME extends the xDSL line configuration MEs. The ME name was chosen because its attributes  were initially unique to ITU -T G.993.2 VDSL2. Due to continuing standards development, some  attributes – and therefore this ME – have also become applicable to other Recommendations,  specifically [ITU-T G.992.3] and [ITU-T G.992.5].  This ME contains general and downstream attributes.  Relationships  This is one of the status data MEs associated with an xDSL UNI. It is meaningful if the PPTP  supports [ITU -T G.992.3], [ITU -T G.992.5] or [ITU -T G.993.2]. The ONU automatically  creates or deletes an instance of this ME upon creation and deletion of a PPTP xDSL UNI  part 1 that supports these attributes.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. Through an  identical ID, this ME is implicitly linked to an instance of the PPTP xDSL UNI  part 1 ME. (R) (mandatory) (2 bytes)  VDSL2 transmission system capability xTU -C: This attribute extends the xTU -C  transmission system capability attribute of the xDSL line inventory and status  data part 1 to include xTU -C VDSL2 capabilities. It is defined by bits 57..64  of Table 9.7.12-1. (R) (mandatory) (1 byte)  VDSL2 transmission system: This attribute reports the transmission system in use. It extends  the xDSL transmission system attribute of the xDSL line inventory and status  data part 2 ME with a byte that includes VDSL2 capabilities currently in use.  It is defined by bits 57..64 of Table 9.7.12-1. (R) (mandatory) (1 byte)  VDSL2 profile: This attribute identifies the profile in use. It is a bit map (0 if not allowed, 1  if allowed) with the following definition:  Bit Meaning  1 (LSB) ITU-T G.993.2 profile 8a  2 ITU-T G.993.2 profile 8b  3 ITU-T G.993.2 profile 8c  4 ITU-T G.993.2 profile 8d  5 ITU-T G.993.2 profile 12a  6 ITU-T G.993.2 profile 12b  7 ITU-T G.993.2 profile 17a  8 ITU-T G.993.2 profile 30a  (R) (mandatory) (1 byte) 

---- page break ---- VDSL2 limit PSD mask and bandplan: This attribute defines the limit PSD mask and band  plan in use. It is a bit map as defined by Table 9.7.6-1. (R) (mandatory)  (8 bytes)  VDSL2 US0 PSD mask: This attribute defines the US0 PSD mask in use. It is a bit map as  defined by Table 9.7.6-2. (R) (mandatory) (4 bytes)  ACTSNRMODEds: This attribute indicates whether transmitter -referred virtual noise is  active on the line in the downstream direction.  1 Virtual noise inactive  2 Virtual noise active  (R) (mandatory) (1 byte)  The following four attributes have similar definitions. In each case, valid attribute values are 1, 2, 4,  8. In ADSL applications, the corresponding value is fixed at 1, and therefore need not be specified.  For VDSL2, it is equal to the size of the subcarr ier group used to compute these attributes (see  clause 11.4.1 of [ITU-T G.993.2]).  HLINGds: This attribute contains the number of subcarriers per group used to report  HLINpsds. (R) (mandatory) (1 byte)  HLOGGds: This attribute contains the number of subcarriers per group used to report  HLOGpsds. (R) (mandatory) (1 byte)  QLNGds: This attribute contains the number of subcarriers per group used to report  QLNpsds. (R) (mandatory) (1 byte)  SNRGds: This attribute contains the number of subcarriers per group used to report  SNRpsds. (R) (mandatory) (1 byte)  MREFPSDds table : The downstream medley reference PSD table contains the set of  breakpoints exchanged in the MREFPSDds fields of the O -PRM message of  [ITU-T G.993.2].  The format is similar to that specified f or the PSD descriptor in  [ITU-T G.993.2]. In [ITU-T G.993.2], the first byte gives the size of the table,  each entry of which is 3 bytes. In the OMCI definition, the first byte is omitted  because the size of the table is known from the response to the get command.  (R) (mandatory) (3 * N bytes, where N is the number of breakpoints)  TRELLISds: This attribute reports whether trellis coding is in use in the downstream  direction.  0 Trellis not used  1 Trellis used  (R) (mandatory for ITU-T G.993.2 VDSL2, optional for others) (1 byte)  Actual rate adaptation mode downstream : The ACT-RA-MODEds attribute indicates the  actual active RA mode in the downstream direction.  1 MANUAL  2 AT_INIT  3 DYNAMIC  4 DYNAMIC with SOS ([ITU-T G.993.2] only)  (R) (optional) (1 byte)  Actual impulse noise protection robust operations channel ( ROC) downstream: The  ACTINP-ROC-ds attribute reports the actual INP of the ROC in the  downstream direction expressed in multiples of T4k. The INP of this attribute
```
