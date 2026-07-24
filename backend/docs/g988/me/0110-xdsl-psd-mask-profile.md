# Managed Entity

## Identity
- ME ID: 110
- ME Name: xDSL PSD mask profile
- Source Section: 9.7.10
- Source Page: 271

## Classification
- Access: needs_review
- Type: needs_review
- Actions: needs_review
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: PSD mask table
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 2
- Name: Downstream RFI bands table
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: mandatory

## Extraction Status
- Status: auto_extracted
- Attributes found: 2
- Review needed: true

## Raw Source

```
9.7.10 xDSL PSD mask profile  This ME contains a PSD mask profile for an xDSL UNI. An instance of this ME is created and deleted  by the OLT.  Relationships  An instance of this ME may be associated with zero or more instances of the PPTP xDSL UNI  part 1.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. The value 0  is reserved. (R, set-by-create) (mandatory) (2 bytes)  PSD mask table: This attribute is a table that defines the PSD mask applicable at the U -C2  reference point (downstream) or the U -R2 reference point (upstream). This  mask may impose PSD restrictions in addition to the limit PSD mask defined  in the relevant Recommendation s ([ITU -T G.992.3], [ITU -T G.992.5],  [ITU-T G.993.2]).  NOTE – In [ITU-T G.997.1], this attribute is called PSDMASKds (downstream) and  PSDMASKus (upstream). In [ITU -T G.993.2], this attribute is called MIBMASKds  (downstream) and MIBMASKus (upstream). The ITU-T G.993.2 MIBMASKus does  not include breakpoints to shape US0.  The PSD mask is specified through a set of breakpoints. Each breakpoint  comprises a 2 byte subcarrier index t, with a subcarrier spacing of 4.3125 kHz,  and a 1 byte PSD mask level at that subcarrier. The set of breakpoints can then  be represented as [(t1, PSD1), (t2, PSD2), …, (tN, PSDN)]. The PSD mask level  is coded as 0 (0.0  dBm/Hz) to 190   (–95.0 dBm/Hz), in steps of 0.5 dB.  The maximum number of downstream breakpoints is 32. In the upstream  direction, the maximum number of breakpoints is 4 for [ITU -T G.992.3] and  16 for [ITU -T G.993.2]. The requirements for a valid set of breakpoints are  defined in the relevant Recommendations ([ITU-T G.992.3], [ITU-T G.992.5],  [ITU-T G.993.2]).  Each table entry in this attribute comprises:  – an entry number field (1 byte, first entry numbered 1);  – a subcarrier index field, denoted t (2 bytes);  – a PSD mask level field (1 byte). 

---- page break ---- By default, the PSD mask table is empty. Setting a subcarrier entry with a valid  PSD mask level implies insertion into the table or replacement of an existing  entry. Setting an entry 's PSD mask level to 0xFF implies deletion from the  table.  (R, W) (mandatory) (4 * N bytes where N is the number of breakpoints)  Mask valid: This Boolean attribute controls and reports the status of the PSD mask  attribute.  As a status report, the value false indicates that the PSD mask represented in  this ME has not been impressed on the DSL equipment. The value true  indicates that the PSD mask represented in this ME has been impressed on the  DSL equipment.  This attribute behaves as follows.  • If the OLT changes any of the PSD mask table entries or sets mask  valid false, then mask valid is false.  • If mask valid is false and the OLT sets mask valid true, the ONU  impresses the PSD mask data on the DSL equipment.  (R, W) (mandatory) (1 byte)  Actions  Create, delete, get, get next, set  Set table (optional)  Notifications  None.  9.7.11 xDSL downstream RFI bands profile  This ME contains the downstream RFI bands profile for an xDSL UNI. Instances of this ME are  created and deleted by the OLT.  Relationships  An instance of this ME may be associated with zero or more instances of the PPTP xDSL UNI  part 1.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. The value 0  is reserved. (R, set-by-create) (mandatory) (2 bytes)  Downstream RFI bands table : The RFIBANDS attribute is a table where each entry  comprises:  • an entry number field (1 byte, first entry numbered 1);  • subcarrier index 1 field (2 bytes);  • subcarrier index 2 field (2 bytes).  For [ITU -T G.992.5], this configuration attribute defines the subset of  downstream PSD mask breakpoints, as specified in the downstream PSD mask,  to be used to notch an RFI band. This subset consists of couples of consecutive  subcarrier indices belonging to breakpoints: [ti; ti + 1], corresponding to the low  level of the notch. Interpolation around these points is defined in [ITU -T  G.992.5]. 

---- page break ---- For [ITU-T G.993.2], this attribute defines the bands where the PSD is to be  reduced as specified in clause 7.2.1.2 of [ITU -T G.993.2]. Each band is  represented by start and stop subcarrier indices with a subcarrier spacing of  4.3125 kHz. Up to 16 bands may be specified. This attribute defines the RFI  bands for both upstream and downstream directions.  Entries have the default value 0 for both subcarrier index 1 and subcarrier  index 2. Setting an entry with a non -zero subcarrier index 1 and subcarrier  index 2 implies insertion into the table or replacement of an existing entry.   Setting an entry 's subcarrier index 1 and subcarrier index 2 to 0 implies  deletion from the table, if present.  (R, W) (mandatory for [ITU-T G.992.5], [ITU-T G.993.2]) (5 * N bytes where  N is the number of RFI bands)  Bands valid: This Boolean attribute controls and reports the operational status of the  downstream RFI bands table.  If this attribute is true, the downstream RFI bands table has been impressed on  the DSL equipment.  If this attribute is false, the downstream RFI bands table has not been  impressed on the DSL equipment. The default value is false.  This attribute can be modified by the ONU and OLT, as follows.  • If the OLT changes any of the RFI bands table entries or sets bands  valid false, then bands valid is false.  • If bands valid is false and OLT sets bands valid true, the ONU  impresses the downstream RFI bands data on to the DSL equipment.  (R, W) (mandatory) (1 byte)  Actions  Create, delete, get, get next, set  Set table (optional)  Notifications  None.  9.7.12 xDSL line inventory and status data part 1  This ME contains part 1 of the line inventory and status data for an xDSL UNI. The ONU  automatically creates or deletes an instance of this ME upon the creation or deletion of a PPTP xDSL  UNI part 1.  Relationships  An instance of this ME is associated with an xDSL UNI.  A
```
