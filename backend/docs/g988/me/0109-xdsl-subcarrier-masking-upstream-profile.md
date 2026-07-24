# Managed Entity

## Identity
- ME ID: 109
- ME Name: xDSL subcarrier masking upstream profile
- Source Section: 9.7.9
- Source Page: 270

## Classification
- Access: needs_review
- Type: needs_review
- Actions: needs_review
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: Upstream subcarrier mask
- Size: 8 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

## Extraction Status
- Status: auto_extracted
- Attributes found: 1
- Review needed: true

## Raw Source

```
9.7.9 xDSL subcarrier masking upstream profile  This ME contains the subcarrier masking upstream profile for an xDSL UNI. An instance of this ME  is created and deleted by the OLT.  Relationships  An instance of this ME may be associated with zero or more instances of the PPTP xDSL UNI  part 1.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. The value 0  is reserved. (R, set-by-create) (mandatory) (2 bytes)  Upstream subcarrier mask: This attribute is a bit map representing upstream mask values  for subcarriers 1 to 64. The MSB of byte 1 corresponds to subcarrier 1, and the 

---- page break ---- LSB of byte 8 corresponds to subcarrier 64. Each bit position defines whether  the corresponding downstream subcarrier is masked (1) or not masked (0).  Subcarrier number 1 is the lowest, and the number of xDSL subcarriers,  upstream ( NSCus) is the highest subcarrier that can be transmitted in the  upstream direction. For [ITU -T G.992.3], [ITU -T G.992.4] and  [ITU-T G.992.5], it is defined in the corresponding Recommendation. For  Annex A of [ITU-T G.992.1] and [ITU-T G.992.2], NSCus = 32 and for Annex  B of [ITU-T G.992.1], NSCus = 64. (R, W, set-by-create) (mandatory) (8 bytes)  Actions  Create, delete, get, set  Notifications  None.  9.7.10 xDSL PSD mask profile  This ME contains a PSD mask profile for an xDSL UNI. An instance of this ME is created and deleted  by the OLT.  Relationships  An instance of this ME may be associated with zero or more instances of the PPTP xDSL UNI  part 1.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. The value 0  is reserved. (R, set-by-create) (mandatory) (2 bytes)  PSD mask table: This attribute is a table that defines the PSD mask applicable at the U -C2  reference point (downstream) or the U -R2 reference point (upstream). This  mask may impose PSD restrictions in addition to the limit PSD mask defined  in the relevant Recommendation s ([ITU -T G.992.3], [ITU -T G.992.5],  [ITU-T G.993.2]).  NOTE – In [ITU-T G.997.1], this attribute is called PSDMASKds (downstream) and  PSDMASKus (upstream). In [ITU -T G.993.2], this attribute is called MIBMASKds  (downstream) and MIBMASKus (upstream). The ITU-T G.993.2 MIBMASKus does  not include breakpoints to shape US0.  The PSD mask is specified through a set of breakpoints. Each breakpoint  comprises a 2 byte subcarrier index t, with a subcarrier spacing of 4.3125 kHz,  and a 1 byte PSD mask level at that subcarrier. The set of breakpoints can then  be represented as [(t1, PSD1), (t2, PSD2), …, (tN, PSDN)]. The PSD mask level  is coded as 0 (0.0  dBm/Hz) to 190   (–95.0 dBm/Hz), in steps of 0.5 dB.  The maximum number of downstream breakpoints is 32. In the upstream  direction, the maximum number of breakpoints is 4 for [ITU -T G.992.3] and  16 for [ITU -T G.993.2]. The requirements for a valid set of breakpoints are  defined in the relevant Recommendations ([ITU-T G.992.3], [ITU-T G.992.5],  [ITU-T G.993.2]).  Each table entry in this attribute comprises:  – an entry number field (1 byte, first entry numbered 1);  – a subcarrier index field, denoted t (2 bytes);  – a PSD mask level field (1 byte).
```
