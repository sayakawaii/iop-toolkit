# Managed Entity

## Identity
- ME ID: 432
- ME Name: FAST channel configuration profile, part 1
- Source Section: 9.7.53
- Source Page: 342

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
9.7.53 FAST channel configuration profile, part 1  This ME contains the FAST channel configuration profile for an xDSL UNI. An instance of this ME  is created and deleted by the OLT.  Relationships  An instance of this ME is associated with a PPTP UNI part 3. 

---- page break ---- The overall FAST line configuration profile is modelled in several parts, all of which are  associated together through a common ME ID . (The client PPTP xDSL UNI part 3 refers to  the entire set of line configuration parts).    Attributes  Managed entity ID : This attribute uniquely identifies each instance of this ME . All FAST  line configuration profiles and extensions that pertain to a given PPTP xDSL  UNI must share a common ME ID. (R, set-by-create) (mandatory) (2 bytes)  Upstream maximum net data rate (MAXNDRus): This attribute specifies the value of the  maximum upstream net data rate. See clause 11.4.2.2 of [ITU -T G.9701].  Valid values range from 0 (0  kbit/s) to 4294967295 (2 32–1 kbit/s). See clause  7.2.1.1 of [ITU-T G.997.2]. (R, W) (mandatory) (4 bytes)  Upstream minimum expected throughput (MINETRus): This attribute specifies the value  of the minimum upstream expected throughput. See clause 11.4.2.1 of [ITU-T  G.9701]. Valid values range from 0 (0  kbit/s) to 4294967295 (2 32–1 kbit/s).  See clause 7.2.1.2 of [ITU-T G.997.2]. (R, W) (mandatory) (4 bytes)  Upstream m aximum gamma data rate (MAXGDR us): This attribute specifies the  maximum upstream value of the GDR (see clause 7.11.1.3). The GDR shall  not exceed MAXGDR at the start of showtime and during showtime. Valid  values range from 0 (0 kbit/s) to 4294967295 (232–1 kbit/s). See clause 7.2.1.3  of [ITU-T G.997.2]. (R, W) (mandatory) (4 bytes)  Upstream minimum gamma data rate (MINGDRus): This attribute specifies the minimum  upstream value of the GDR (see clause 7.11.1.3). The GDR may be lower than  MINGDR. If the GDR is lower than MINGDR at initialization or when GDR  becomes lower than MINGDR during showtime, a TCA occurs. Valid values  range from 0 (0  kbit/s) to 4294967295 (2 32–1 kbit/s). See clause 7.2.1.4 of  [ITU-T G.997.2]. (R, W) (mandatory) (4 bytes)  Upstream m aximum delay (DELAYMAX us): This attribute specifies the maximum  allowed delay for upstream retransmission. See clause 9.8 of [ITU-T G.9701].  The ITU-T G.9701 control parameter delay_max is set to the same value as the  maximum delay. See clause 11.4.2.3 of [ITU -T G.9701]. Valid values range  from 4 (1  ms) to 252 (63  ms) in steps of 0.25  ms. See clause 7.2.2.1 of  [ITU-T G.997.2]. (R, W) (mandatory) (4 bytes)  Upstream minimum impulse noise protection against SHINE (INPMIN_SHINEus): This  attribute specifies the minimum upstream INP against SHINE. See clause 9.8  of [ITU-T G.9701]. The ITU -T G.9701 control parameter INP_min_shine is  set to the same value as the minimum INP against SHINE. See clause 11.4.2.4  of [ITU-T G.9701]. Valid values range from 0 to 520 (520 symbol periods).  See clause 7.2.2.2 of [ITUT G.997.2]. (R, W) (mandatory) (2 bytes)  Upstream SHINE ratio (SHINERATIO us): This attribute specifies the upstream SHINE  ratio that is used in the definition of the expected throughput rate (ETR). See  clause 9.8 of [ITU -T G.9701]. The ITU -T G.9701 control parameter  SHINEratio is set to the same value as the SHINE ratio. See clause 11.4.2.5 of  [ITU-T G.9701]. The value is expressed in units of 0.001, Valid values range  from 0 to 100 (0.01) in steps of 0.001. See clause 7.2.2.3 of [ITU-T G.997.2].  (R, W) (mandatory) (1 byte) 

---- page break ---- Upstream minimum impulse noise protection against REIN (INPMIN_REIN us): This  attribute specifies the minimum upstream INP against REIN. See clause 9.8 of  [ITU-T G.9701]. The ITU-T G.9701 control parameter INP_min_rein is set to  the same value as the minimum INP against REIN. See clause 11.4.2.6 of  [ITU-T G.9701]. Valid values range from 0 to 63 (63  symbol periods). See  clause 7.2.2.4 of [ITU-T G.997.2]. (R, W) (mandatory) (1 byte)  Upstream REIN Inter-arrival time (IAT_REIN us): This attribute specifies the upstream  REIN IAT. See clause 9.8 of [ITU -T G.9701]. The ITU -T G.9701 control  parameter iat_rein_flag is set to the same value as the REIN IAT. See clause  11.4.2.7 of [ITU-T G.9701].  The REIN IAT is specified via the following values:  1 100 Hz;  2 120 Hz;  3 360 Hz.  See clause 7.2.2.5 of [ITU-T G.997.2].  (R, W) (mandatory) (1 byte)  Upstream m inimum Reed-Solomon RFEC/NFEC ratio (RNRATIO us): This attribute  specifies the minimal required ratio, RFEC/NFEC, of Reed -Solomon code  parameters in the upstream direction . The ITU -T G.9701 control parameter  rnratio is set to the same value as the minimum Reed-Solomon RFEC/NFEC  ratio. See clause 11.4.2.8 of [ITU -T G.9701]. The value is expressed in units  of 1/32, Valid values rang e from 0 to 8 (1/4). See clause 7.2.2.6 of [ITU -T  G.997.2]. (R, W) (mandatory) (1 byte)  RTX-TC testmode (RTX_TESTMODE) : This Boolean attribute specifies whether the  retransmission test mode defined in clause 9.8.3.1.2 [ITU -T G.9701] is  enabled (true) or disabled (disabled). See clause 7.2.2.7 of [ITU -T G.997.2].  (R, W) (optional) (1 byte)  Actions  Create, delete, get, set  Notifications  None.  9.7.53.1  FAST channel configuration profile, part 2  This ME contains the FAST channel configuration profile for an xDSL UNI. An instance of this ME  is created and deleted by the OLT.  Relationships  An instance of this ME is associated with a PPTP UNI part 3.  The overall FAST line configuration profile is modelled in several parts, all of which are  associated together through a common ME ID. (The client PPTP UNI part 3 refers to the entire  set of line configuration parts).  Attributes   Managed entity ID : This attribute uniquely identifies each instance of this ME. All FAST  line configuration profiles and extensions that pertain to a given PPTP xDSL  UNI must share a common ME ID. (R, set-by-create) (mandatory) (2 bytes) 

---- page break ---- Downstrea
```
