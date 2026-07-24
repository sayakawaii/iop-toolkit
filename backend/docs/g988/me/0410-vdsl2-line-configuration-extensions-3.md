# Managed Entity

## Identity
- ME ID: 410
- ME Name: VDSL2 line configuration extensions 3
- Source Section: 9.7.6
- Source Page: 257

## Classification
- Access: needs_review
- Type: needs_review
- Actions: needs_review
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: VDSL2 profiles enabling
- Size: 1 byte
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 2
- Name: VDSL2 limit PSD masks
- Size: 8 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 3
- Name: VDSL2 US0 disabling
- Size: 8 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 4
- Name: VDSL2 US0 PSD masks
- Size: 4 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

## Extraction Status
- Status: auto_extracted
- Attributes found: 4
- Review needed: true

## Raw Source

```
9.7.6 VDSL2 line configuration extensions  This ME extends the xDSL line configuration MEs with attributes that were originally unique to ITU- T G.993.2 VDSL2. Due to continuing standards development, some attributes – and therefore this  ME – have also become applicable to other Recommendations, specifically [ITU -T G.992.3] and  [ITU-T G.992.5]. The attributes of this ME are further defined in [ITU -T G.997.1]. An instance of  this ME is created and deleted by the OLT.  Relationships  An instance of this ME may be associated with zero or more instances of an xDSL UNI.  The overall xDSL line configuration profile is modelled in several parts, all of which are  associated together through a common ME ID (the client PPTP xDSL UNI part 1 has a single  pointer, which refers to the entire set of line configuration parts). 

---- page break ---- Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. All xDSL  and VDSL2 line configuration profiles and extensions that pertain to a given  PPTP xDSL UNI must share a common ME ID. (R, set-by-create) (mandatory)  (2 bytes)  VDSL2 profiles enabling: The PROFILES attribute contains the ITU -T G.993.2 profiles to  be allowed by the xTU -C. It is coded in a bit map representation (0 if not  allowed, 1 if allowed) with the following definition.  Bit Meaning  1 (LSB) ITU-T G.993.2 profile 8a  2 ITU-T G.993.2 profile 8b  3 ITU-T G.993.2 profile 8c  4 ITU-T G.993.2 profile 8d  5 ITU-T G.993.2 profile 12a  6 ITU-T G.993.2 profile 12b  7 ITU-T G.993.2 profile 17a  8 ITU-T G.993.2 profile 30a  (R, W, set-by-create) (mandatory for ITU-T G.993.2) (1 byte)  VDSL2 PSD mask class selection (CLASSMASK): To reduce the number of configuration  possibilities, the limit PSD masks are grouped in the following PSD mask  classes.  – Class 998 Annex A of [ITU-T G.993.2]: D-32, D-48, D-64, D-128  – Class 997-M1c Annex B of [ITU-T G.993.2]: 997-M1c-A-7  – Class 997-M1x Annex B of [ITU-T G.993.2]: 997-M1x-M  – Class 997 -M2x Annex B of [ITU -T G.993.2]: 997E17-M2x-NUS0,  997E30-M2x-NUS0  – Class 998-M2x Annex B of [ITU-T G.993.2]: 998-M2x-A,  998-M2x-M, 998-M2x-B, 998-M2x-NUS0, 998E17-M2x-NUS0,  998E17-M2x-NUS0-M, 998E30-M2x-NUS0,  998E30-M2x-NUS0-M, 998E17-M2x-A  – Class 998ADE-M2x Annex B of [ITU-T G.993.2]: 998-M2x-A,  998-M2x-M, 998-M2x-B, 998-M2x-NUS0, 998ADE17-M2x-A,  998ADE17-M2x-B, 998ADE17-M2x-M, 998ADE17-M2x-NUS0-M,  998ADE30-M2x-NUS0-A, 998ADE30-M2x-NUS0-M,  HPEADE1230, HPEADE1730  – Class 998 -B Annex C: POTS -138b, POTS -276b (clause C.2.1.1 of  [ITU-T G.993.2]), TCM-ISDN (clause C.2.1.2 of [ITU-T G.993.2])  – Class 998 -CO Annex C of [ITU -T G.993.2]: POTS -138co,  POTS-276co (clause C.2.1.1 of [ITU-T G.993.2])  – Class HPE -M1 Annex B of [ITU -T G.993.2]: HPE17 -M1-NUS0,  HPE30-M1-NUS0, HPE1230-M1-NUS0, HPE1730-M1-NUS0  Each class is designed such that the PSD levels of each limit PSD mask of a  specific class are equal in their respective passbands above 552 kHz. 

---- page break ---- The CLASSMASK attribute is defined per annex of [ITU-T G.993.2] enabled  in the xTSE table (see Table 9.7.12 -1). It selects a single PSD mask class per  annex of [ITU-T G.993.2] to be activated at the very high -speed digital  subscriber line t ransceiver unit, operator end  (VTU-O). The coding is as  follows:  Attribute value Annex A of [ITU- T G.993.2]   Annex B of [ITU- T G.993.2]  Annex C of [ITU- T G.993.2]  1 998 997-M1c 998-B  2  997-M1x 998-CO  3  997-M2x   4  Deprecated   5  998-M2x   6  998ADE-M2x   7  HPE   NOTE 1 – A single PSD mask class may be selected per annex of [ITU-T G.993.2].  NOTE 2 – It is expected that only a single annex will be enabled at any given time,  such that the CLASSMASK attribute, as well as the LIMITMASK and US0DISABLE  attributes below, need not be vectors of values.   NOTE 3 – Attribute value 4 was formerly defined in [ITU -T G.997.1], and is no  longer used.  (R, W, set-by-create) (mandatory) (1 byte)  VDSL2 limit PSD masks: The LIMITMASK attribute contains the ITU-T G.993.2 limit PSD  masks of the selected PSD mask class, enabled by the near -end xTU for each  class of profiles. One LIMITMASK parameter is defined per annex enabled in  the xTSE (see Table 9.7.12-1).  The profiles are grouped in the following profile classes:  – Class 8: Profiles 8a, 8b, 8c, 8d  – Class 12: Profiles 12a, 12b  – Class 17: Profile 17a  – Class 30: Profile 30a  For each profile class, several limit PSD masks of the selected PSD mask class  (CLASSMASK) may be enabled. The enabling attribute is coded in a bit map  representation (0 if the associated mask is not allowed, 1 if it is allowed). The  bit mask is defined in  Table 9.7.6-1. (R,  W, set-by-create) (mandatory)  (8 bytes)  VDSL2 US0 disabling : The US0DISABLE attribute specifies whether channel US0 is  disabled for each limit PSD mask enabled in the LIMITMASK attribute.  For each limit PSD mask enabled in the LIMITMASK attribute, one bit  indicates if US0 is disabled. The disabling attribute is a bit map where the value  1 specifies that US0 is disabled for the associated limit mask. The bit map has  the same structure as the  LIMITMASK attribute. (R,  W, set-by-create)  (mandatory) (8 bytes)  VDSL2 US0 PSD masks : The US0MASK attribute contains the US0 PSD masks to be  allowed by the xTU -C. This attribute is only defined for Annex  A of  [ITU-T G.993.2]. It is represented as a bit map (0 if not allowed, 1 if allowed)  with the definitions of Table 9.7.6-2. (R, W, set-by-create) (mandatory)  (4 bytes) 

---- page break ---- VDSL2-CARMASK table: This attribute specifies restrictions, additional to the band plan,  that determine the set of subcarriers allowed for transmission in both upstream  and downstream directions.  The VDSL2 -CARMASK attribute describes the not -masked subcarriers in  terms of one or more frequency bands. Each band is represented by start and  stop subcarrier indices with a subcarrier spacing of 4.3125  kHz. The valid  range of subcarrier indices i
```
