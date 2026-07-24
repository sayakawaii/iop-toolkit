# Managed Entity

## Identity
- ME ID: 429
- ME Name: FAST line configuration profile part 2
- Source Section: 9.7.49
- Source Page: 333

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Create, Delete, Get, Set
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

<!-- No attributes successfully parsed; see Raw Source section -->

## Extraction Status
- Status: auto_extracted
- Attributes found: 0
- Review needed: true

## Raw Source

```
Notifications  None.  9.7.49 FAST line configuration profile part 1  This ME extends the xDSL line configuration MEs with attributes that are unique to  FAST (see  [ITU-T G.9700] and [ITU-T G.9701]). The attributes of this ME are defined in [ITU-T G.997.2]. An  instance of this ME is created and deleted by the OLT.  Relationships  An instance of this ME is associated with a PPTP UNI part 3.  The overall FAST line configuration profile is modelled in several parts, all of which are  associated together through a common ME ID (The client PPTP UNI part 3 refers to the entire  set of line configuration parts).  Attributes  Managed entity ID : This attribute uniquely identifies each instance of this ME. All FAST  line configuration profiles and extensions that pertain to a given PPTP xDSL  UNI must share a common ME ID. (R, set-by-create) (mandatory) (2 bytes)  ITU-T G.9701 profiles enabling (PROFILES) : This attribute contains the ITU -T G.9701  profiles to be allowed by the xTU-C. See clause 7.1.0.1 of [ITU-T G.997.2]. It  is coded in a bit map representation (0 if not allowed, 1 if allowed) with the  following definition:  Bit Meaning  1 (LSB)  ITU-T G.9701 profile 106a  2 ITU-T G.9701 profile 106b  3  ITU-T G.9701 profile 106c  4 ITU-T G.9701 profile 212a  5 ITU-T G.9701 profile 212c  (R, W, set-by-create) (mandatory) (1 byte)  Symbol periods per TDD frame (MF) : This attribute specifies the total number of symbol  periods in a TDD frame. See clause 10.5 of [ITU-T G.9701]. Valid values are  23 and 36. (R, W, set-by-create) (mandatory) (1 byte)  Symbol periods per TDD frame dedicated for downstream transmission (Mds) : This  attribute specifies the total number of symbol positions in a TDD frame  allocated for downstream transmission. The total number of symbol positions  in a TDD frame allocated for upstream transmission is calculated as Mus  =  MF − 1 − Mds. See clause 10.5 of [ITU -T G.9701]. Valid values range from  10 to 32 (if MF  = 36) and from 6 to 19 (if MF  = 23). The default value is 28  (if MF  = 36) and 18 (if MF  = 23). See clause 7.1.1.2 of   [ITU-T G.997.2]. (R, W, set-by-create) (mandatory) (1 byte)  Downstream maximum aggregate transmit power (MAXATPds): This attribute specifies  the maximum aggregate transmit power at the U -O2 reference point in the  downstream direction during initialization and showtime (in d ecibel- milliwatts). Valid values range from 5 to 31 in steps of 1  s. See clause 7.1.2.1  of [ITU-T G.997.2]. (R, W) (mandatory) (2 bytes)  Upstream maximum aggregate transmit power (MAXATPus): This attribute specifies the  maximum aggregate transmit power at the U -R2 reference point in the  upstream direction during initialization and showtime (in d ecibel-milliwatts). 

---- page break ---- The attribute value ranges from 0 (−31.0 dBm) to 620 (+31.0 dBm) in steps of  0.1 dBm. See clause 7.1.2.2 of [ITU-T G.997.2]. (R, W) (mandatory) (2 bytes)  Downstream subcarrier masking (CARMASKds) table : This attribute specifies the  masked subcarrier bands in the downstream direction. All subcarriers within  the band, i.e., with indices higher than or equal to the start subcarrier index and  lower than or equal to the stop subcarrier index, are masked, i.e. , have a  transmit power set to zero (linear scale)  The CARMASK attribute is a table where each entry comprises:  – an entry number field (1 byte, first entry numbered 1);  – band start subcarrier index (2 bytes);  – band stop subcarrier index (2 bytes).  Subcarrier index valid values range from 0 to 4095 (subcarrier index 0 to 4095).  By default, no masked subcarriers, the table is empty. Setting a table entry with  non-zero subcarrier references implies insertion into the table. Setting an  entry's subcarrier references to 0x FFFF implies deletion from the table, if  present. See clause 7.1.2.3 of [ITU-T G.997.2].  (R, W) mandatory (5*N bytes)  Upstream subcarrier masking (CARMASKus) table : This attribute specifies the masked  subcarrier bands in the upstream direction. All subcarriers within the band, i.e.,  with indices higher than or equal to the start subcarrier index and lower than  or equal to the stop subcarrier index, have a transmit po wer set to zero (linear  scale).  The CARMASK attribute is a table where each entry comprises:  – an entry number field (1 byte, first entry numbered 1);  – band start subcarrier index (2 bytes);  – band stop subcarrier index (2 bytes).  Valid value of band subcarrier index ranges from 0 to 4095. By default, no  masked subcarriers, the table is empty. Setting a table entry with non -zero  subcarrier references implies insertion into the table. Setting an entry's  subcarrier references to 0xFFFF implies deletion from the table, if present. See  clause 7.1.2.4 of [ITU-T G.997.2].  (R, W) (mandatory) (5*N bytes)  Downstream PSD mask (PSDMASKds) table: This attribute specifies the downstream PSD  mask applicable at the U -O2 reference point. Requirements for a valid  PSDMASKds are defined in [ITU -T G.9701] clauses 7.3.1.1.2.1 and  7.3.1.1.2.2.  Each table entry in this attribute comprises:  –  an entry number field (1 byte, first entry numbered 1);  –  a subcarrier index field, denoted t (2 bytes);  –  a PSD mask level field (1 byte).  The valid value of the subcarrier index ranges from 0 to 4095. The valid values  of PSD level range from 0 (0  dBm/Hz) to 255 (−127.5  dBm/Hz), with a  granularity of −0.5  dBm/Hz. Setting a table entry with non -zero subcarrier  references implies insertion into the table. Setting an entry's subcarrier 

---- page break ---- references to 0xFFFF implies deletion from the table, if present. See  clause 7.1.2.5 of [ITU-T G.997.2].  (R, W) (mandatory) (4*N N<=32bytes)  Upstream PSD mask (PSDMASKus) table: This attribute specifies the upstream PSD mask  applicable at the U-R2 reference point. Requirements for a valid PSDMASKds  are defined in [ITU-T G.9701] clauses 7.3.1.1.2.1 and 7.3.1.1.2.2.  Each table entry in this attribute comprises:  – an e
```
