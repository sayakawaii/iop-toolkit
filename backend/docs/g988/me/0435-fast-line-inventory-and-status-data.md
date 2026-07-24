# Managed Entity

## Identity
- ME ID: 435
- ME Name: FAST line inventory and status data
- Source Section: 9.7.56
- Source Page: 347

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Get
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
9.7.56 FAST line inventory and status data, part 1  This ME extends the FAST line inventory and status data MEs with attributes specific to [ITU -T  G.997.2]. The ONU automatically creates or deletes an instance of this ME upon the creation or  deletion of a PPTP xDSL UNI part 1.  Relationships  This is one of the status data MEs associated with an xDSL UNI  part 1. It is required only if  FAST is supported by the PPTP. The ONU automatically creates or deletes an instance of this  ME upon creation or deletion of a PPTP xDSL UNI part 1 that supports these attributes.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. Through an  identical ID, this ME is implicitly linked to an instance of the PPTP xDSL UNI  part 1. (R) (mandatory) (2 bytes)  ITU-T G.9701 profile (PROFILE): This attribute reports for each profile whether operation  according to that profile is enabled (0) or disabled (1). Only one profile can be  enabled. See clause 7.10.1.1 of [ITU-T G.997.2] (R) (mandatory) (1 byte)  Upstream gamma data rate (GDRus): This attribute reports the upstream net data rate as  defined in clause  7.11.1.1, lowered by any throughput capability limitations  remaining in the DRA or L2+ functions, assuming no user data are transmitted 

---- page break ---- over all the other lines. Valid values range from 0 (0  kbit/s) to 4294967295  (232–1 kbit/s). See clause 7.11.1.3 of [ITU-T G.997.2] (R) (mandatory) (bytes)  Upstream attainable gamma data rate (ATTGDRus): This attribute reports the attainable  upstream net data rate (as defined in clause 7.11.2.1), lowered by any  throughput capability limitations remaining in the DRA or L2+ functions,  assuming no user data are transmitted over all the other Lines, and assuming  MAXGDR (as defined in clause 7.2.1.3) is configured to its maximum valid  value. Valid values range from 0 (0 kbit/s) to 4294967295 (232 – 1 kbit/s). See  clause 7.11.2.3 of [ITU-T G.997.2] (R) (mandatory) (bytes)  DPU system vendor ID (DPU_SYSTEM_VENDOR) : This attribute reports the DPU  system vendor ID as inserted by the FTU -O in the embedded operations  channel (see clause 11.2.2.10 of [ITU -T G.9701]) and as defined in  clause 9.3.3.1 of [ITU-T G.994.1]. See clause 7.13.2.1 of [ITU-T G.997.2] (R)  (optional) (8 bytes)  NT system vendor ID (NT_SYSTEM_VENDOR) : This attribute reports the NT system  vendor ID as inserted by the FTU-R in the embedded operations channel (see  clause 11.2.2.10 of [ITU -T G.9701]) and as defined in clause 9.3.3.1 of  [ITU-T G.994.1]. See clause 7.13.2.2 of [ITU -T G.997.2] (R) (optional)  (8 bytes)  DPU serial number (DPU_SYSTEM_SERIALNR) : This attribute reports the DPU serial  number as inserted by the FTU -O in the embedded operations channel. See  clause 11.2.2.10 of [ITU -T G.9701]. It is vendor -specific information. The  combination of DPU system vendor ID and DPU system serial number creates  a unique number for each DPU. See clause 7.13.2.3 of [ITU -T G.997.2] (R)  (optional) (32 bytes)  NT serial number (NT_SYSTEM_SERIALNR): This attribute reports the NT system serial  number as inserted by the FTU -R in the embedded operations channel. See  clause 11.2.2.10 of [ITU -T G.9701]. It shall contain the NT system serial  number, the NT model and the NT firmware version. All shall be en coded in  this order and separated by space characters, i.e., "<NT serial  number><space><NT model><space><NT firmware version>". The  combination of NT system vendor ID and NT system serial number creates a  unique number  for each NT. See clause 7.13.2.4 of [ITU -T G.997.2] (R)  (optional) (32 bytes)  Downstream gamma data rate (GDRds): This attribute reports the downstream net data rate  as defined in clause  7.11.1.1 of [ITU-T G.997.2], lowered by any throughput  capability limitations remaining in the DRA or L2+ functions, assuming no  user data are transmitted over all the other lines. Valid values range from 0  (0 kbit/s) to 4294967295 ( 232 – 1 kbit/s). See clause 7.11.1.3 of [ITU -T  G.997.2] (R) (optional) (4 bytes)  Downstream a ttainable gamma data rate (ATTGDR ds): This attribute reports the  attainable downstream net data rate (as defined in clause 7.11.2.1  of [ITU-T  G.997.2]), lowered by any throughput capability limitations remaining in the  DRA or L2+ functions, assuming no user data are transmitted over all the other  Lines, and assuming MAXGDR (as defined in clause 7.2.1.3) is configured to  its maximum valid value. Valid values range from 0 (0 kbit/s) to 4294967295  (232 – 1 kbit/s). See clause 7.11.2.3 of [ITU-T G.997.2] (R) (optional) (4 bytes)  Downstream signal-to-noise ratio margin (SNRMds): This attribute reports the signal-to- noise ratio margin (as defined in clauses 9 .8.3.2 of [ITU-T G.9701] and 

---- page break ---- 11.4.1.1.10 of [ITU-T G.9701]) in the downstream direction during the L0 link  state. See clause 7.10.3. 1 of [ITU -T G.997.2] for detailed specification.   (R) (optional) (2 bytes).  Upstream signal-to-noise ratio margin (SNRMus): This attribute reports the signal-to-noise  ratio margin (as defined in clauses 9.8.3.2 of [ITU-T G.9701] and 11.4.1.1.10  of [ITU-T G.9701]) in the upstream direction during the L0 link state.  See  clause 7.10.3.2 of [ITU-T G.997.2] for detailed specification. (R) (optional) (2  bytes).  Downstream net data rate (NDRds): This attribute reports the downstream NDR as defined  in clause 11.4.1.1.1 of [ITU-T G.9701]. A special value 0xFFFFFFFF indicates  that the NDR is undetermined.  See clause 7.11.1.1 of [ITU -T G.997.2]  for  detailed specification. (R) (optional) (4 bytes).  Upstream net data rate (NDRus): This attribute reports the downstream NDR as defined in  clause 11.4.1.1.1 of [ITU-T G.9701]. A special value of 0xFFFFFFFF (232–1)  indicates that the NDR is undetermined. See clause 7.11.1.1 of [ITU-T G.997.2]  for detailed specification. (R) (optional) (4 bytes).  Downstream attainable net data rate (ATTNDRds): This attribute reports the downstr
```
