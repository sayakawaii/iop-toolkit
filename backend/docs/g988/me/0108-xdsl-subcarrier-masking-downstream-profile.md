# Managed Entity

## Identity
- ME ID: 108
- ME Name: xDSL subcarrier masking downstream profile
- Source Section: 9.7.8
- Source Page: 269

## Classification
- Access: needs_review
- Type: needs_review
- Actions: needs_review
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: Downstream subcarrier mask 1
- Size: 16 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 2
- Name: Downstream subcarrier mask 2
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 3
- Name: Upstream subcarrier mask
- Size: 8 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

## Extraction Status
- Status: auto_extracted
- Attributes found: 3
- Review needed: true

## Raw Source

```
9.7.8 xDSL subcarrier masking downstream profile  This ME contains the subcarrier masking downstream profile for an xDSL UNI. Instances of this ME  are created and deleted by the OLT.  Relationships  An instance of this ME may be associated with zero or more instances of the PPTP xDSL UNI  part 1.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. The value 0  is reserved. (R, set-by-create) (mandatory) (2 bytes)  The four following attributes are bit maps that represent downstream mask values for subcarriers  1..128 (mask 1) through 385..512 (mask 4). The MSB of the first byte corresponds to the lowest 

---- page break ---- numbered subcarrier, and the LSB of the last byte corresponds to the highest. Each bit position defines  whether the corresponding downstream subcarrier is masked (1) or not masked (0).  The number of xDSL subcarriers, downstream (NSCds) is the highest numbered subcarrier that can  be transmitted in the downstream direction. For [ITU -T G.992.3], [ITU -T G.992.4] and  [ITU-T G.992.5], it is defined in the corresponding Recommendation. For [ITU-T G.992.1], NSCds  = 256 and for [ITU-T G.992.2], NSCds = 128.  Downstream subcarrier mask 1 : Subcarriers 1 to 128. (R,  W, set-by-create) (mandatory)  (16 bytes)  Downstream subcarrier mask 2 : Subcarriers 129 to 256. (R,  W) (mandatory for modems  that support NSCds > 128) (16 bytes)  Downstream subcarrier mask 3 : Subcarriers 257 to 384. (R,  W) (mandatory for modems  that support NSCds > 256) (16 bytes)  Downstream subcarrier mask 4 : Subcarriers 385 to 512. (R,  W) (mandatory for modems  that support NSCds > 384) (16 bytes)  Mask valid: This Boolean attribute controls and reports the operational status of the  downstream subcarrier mask attributes.  If this attribute is true (1), the downstream subcarrier mask represented in this  ME has been impressed on the DSL equipment.  If this attribute is false (0), the downstream subcarrier mask represented in this  ME has not been impressed on the DSL equipment. The default value is false.  The value of this attribute can be modified by the ONU and OLT, as follows.  • If the OLT changes any of the four mask attributes or sets mask valid  false, then mask valid is false.  • If mask valid is false and the OLT sets mask valid true, the ONU  impresses the downstream subcarrier mask data on  to the DSL  equipment.  (R, W) (mandatory) (1 byte)  Actions  Create, delete, get, set  Notifications  None.  9.7.9 xDSL subcarrier masking upstream profile  This ME contains the subcarrier masking upstream profile for an xDSL UNI. An instance of this ME  is created and deleted by the OLT.  Relationships  An instance of this ME may be associated with zero or more instances of the PPTP xDSL UNI  part 1.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. The value 0  is reserved. (R, set-by-create) (mandatory) (2 bytes)  Upstream subcarrier mask: This attribute is a bit map representing upstream mask values  for subcarriers 1 to 64. The MSB of byte 1 corresponds to subcarrier 1, and the 

---- page break ---- LSB of byte 8 corresponds to subcarrier 64. Each bit position defines whether  the corresponding downstream subcarrier is masked (1) or not masked (0).  Subcarrier number 1 is the lowest, and the number of xDSL subcarriers,  upstream ( NSCus) is the highest subcarrier that can be transmitted in the  upstream direction. For [ITU -T G.992.3], [ITU -T G.992.4] and  [ITU-T G.992.5], it is defined in the corresponding Recommendation. For  Annex A of [ITU-T G.992.1] and [ITU-T G.992.2], NSCus = 32 and for Annex  B of [ITU-T G.992.1], NSCus = 64. (R, W, set-by-create) (mandatory) (8 bytes)  Actions  Create, delete, get, set  Notifications  None.  9.7.10 xDSL PSD mask profile  This ME contains a PSD mask profile for an xDSL UNI. An instance of this ME is created and deleted  by the OLT.  Relationships  An instance of this ME may be associated with zero or more instances of the PPTP xDSL UNI  part 1.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. The value 0  is reserved. (R, set-by-create) (mandatory) (2 bytes)  PSD mask table: This attribute is a table that defines the PSD mask applicable at the U -C2  reference point (downstream) or the U -R2 reference point (upstream). This  mask may impose PSD restrictions in addition to the limit PSD mask defined  in the relevant Recommendation s ([ITU -T G.992.3], [ITU -T G.992.5],  [ITU-T G.993.2]).  NOTE – In [ITU-T G.997.1], this attribute is called PSDMASKds (downstream) and  PSDMASKus (upstream). In [ITU -T G.993.2], this attribute is called MIBMASKds  (downstream) and MIBMASKus (upstream). The ITU-T G.993.2 MIBMASKus does  not include breakpoints to shape US0.  The PSD mask is specified through a set of breakpoints. Each breakpoint  comprises a 2 byte subcarrier index t, with a subcarrier spacing of 4.3125 kHz,  and a 1 byte PSD mask level at that subcarrier. The set of breakpoints can then  be represented as [(t1, PSD1), (t2, PSD2), …, (tN, PSDN)]. The PSD mask level  is coded as 0 (0.0  dBm/Hz) to 190   (–95.0 dBm/Hz), in steps of 0.5 dB.  The maximum number of downstream breakpoints is 32. In the upstream  direction, the maximum number of breakpoints is 4 for [ITU -T G.992.3] and  16 for [ITU -T G.993.2]. The requirements for a valid set of breakpoints are  defined in the relevant Recommendations ([ITU-T G.992.3], [ITU-T G.992.5],  [ITU-T G.993.2]).  Each table entry in this attribute comprises:  – an entry number field (1 byte, first entry numbered 1);  – a subcarrier index field, denoted t (2 bytes);  – a PSD mask level field (1 byte).
```
