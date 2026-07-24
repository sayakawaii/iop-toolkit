# Managed Entity

## Identity
- ME ID: 324
- ME Name: xDSL impulse noise monitor performance monitoring history data
- Source Section: 9.7.27
- Source Page: 304

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Create, Delete, Get, Set
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: Interval end time
- Size: 1 byte
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 2
- Name: Threshold data 1/2 ID
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 3
- Name: INM INPEQ histogram table
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 4
- Name: INM IAT histogram
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 5
- Name: INM IAT histogram LFE
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 6
- Name: FEXT downstream SNR margin
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 7
- Name: NEXT downstream SNR margin
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 8
- Name: FEXT upstream SNR margin
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 9
- Name: NEXT upstream SNR margin
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 10
- Name: FEXT downstream maximum attainable data rate
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 11
- Name: NEXT downstream maximum attainable data rate
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 12
- Name: FEXT upstream maximum attainable data rate
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 13
- Name: NEXT upstream maximum attainable data rate
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 14
- Name: FEXT downstream actual power spectral density
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 15
- Name: NEXT downstream actual power spectral density
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 16
- Name: FEXT upstream actual power spectral density
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 17
- Name: NEXT upstream actual power spectral density
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 18
- Name: FEXT downstream actual aggregate transmit power
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 19
- Name: NEXT downstream actual aggregate transmit power
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 20
- Name: FEXT upstream actual aggregate transmit power
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 21
- Name: NEXT upstream actual aggregate transmit power
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 22
- Name: FEXT downstream quiet line noise PSD measurement time
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 23
- Name: NEXT downstream quiet line noise PSD measurement time
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 24
- Name: NEXT upstream quiet line noise PSD measurement time
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 25
- Name: NEXT downstream SNR measurement time
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 26
- Name: NEXT upstream SNR measurement time
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 27
- Name: FEXT downstream bits allocation table
- Size: 2 bytes
- Format: needs_review
- Access: R, set-by-create
- Category: mandatory

### Attribute 28
- Name: Interval end time
- Size: 1 byte
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 29
- Name: Threshold data 1/2 ID
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 30
- Name: Error-free bits counter
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

## Extraction Status
- Status: auto_extracted
- Attributes found: 30
- Review needed: false

## Raw Source

```
9.7.27 xDSL impulse noise monitor performance monitoring history data  This ME collects PM data from the impulse noise monitor function at both near and far ends. Instances  of this ME are created and deleted by the OLT. Note that, unlike most xDSL PM, [ITU -T G.997.1]  only requires current and previous 15 min interval storage; a longer view of this PM is not expected  at 15 min granularity.  For a complete discussion of generic PM architecture, refer to clause I.4.  Relationships  An instance of this ME may be associated with an xDSL UNI. This ME is meaningful only  for ITU-T G.993.2 VDSL2, [ITU-T G.992.3] and [ITU-T G.992.5].  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. The ME ID  is identical to that of this ME 's parent PPTP xDSL UNI part 1. (R,  set-by-create) (mandatory) (2 bytes)  Interval end time : This attribute identifies the most recently finished 15  min interval. (R)  (mandatory) (1 byte)  Threshold data 1/2 ID: No thresholds are defined for this ME. For uniformity with other PM,  the attribute is retained and shown as mandatory, but it should be set to a null  pointer. (R, W, set-by-create) (mandatory) (2 bytes)  INM INPEQ histogram table: INMINPEQ1..17-L is a count of the near-end INMAINPEQi  anomalies occurring on the line during the accumulation period. This  parameter is subject to inhibiting – see clause 7.2.7.13 of [ITU-T G.997.1]. (R)  (optional) (2 bytes * 17 entries = 34 bytes)  INM total measurement : INMME -L is a count of the near -end INMAME anomalies  occurring on the line during the accumulation period. This parameter is subject  to inhibiting – see clause 7.2.7.13 of [ITU-T G.997.1]. (R) (optional) (2 bytes)  INM IAT histogram : INMIAT0..7 -L is a count of the near -end INMAIATi anomalies  occurring on the line during the accumulation period. This parameter is subject  to inhibiting – see clause 7.2.7.13 of [ITU-T G.997.1]. (R) (optional) (2 bytes *  8 entries = 16 bytes)  INM INPEQ histogram LFE table : INMINPEQ1..17 -LFE is a count of the far -end  INMAINPEQi anomalies occurring on the line during the accumulation  period. This parameter is subject to inhibiting – see clause 7.2.7.13 of  [ITU-T G.997.1]. (R) (optional) (2 bytes * 17 entries = 34 bytes) 

---- page break ---- INM total measurement LFE: INMME-LFE is a count of the far-end INMAME anomalies  occurring on the line during the accumulation period. This parameter is subject  to inhibiting – see clause 7.2.7.13 of [ITU-T G.997.1]. (R) (optional) (2 bytes)  INM IAT histogram LFE: INMIAT0..7-LFE is a count of the far-end INMAIATi anomalies  occurring on the line during the accumulation period. This parameter is subject  to inhibiting – see clause 7.2.7.13 of [ITU-T G.997.1]. (R) (optional) (2 bytes *  8 entries = 16 bytes)  Actions  Create, delete, get, get next, set  Get current data (optional)  Notifications  None.  9.7.28 xDSL line inventory and status data part 5  This ME extends the attributes defined in the xDSL line inventory and status data parts 1..4. This ME  reports FEXT and NEXT attributes, and pertains to Annex C of [ITU -T G.992.3] (ADSL2) and  Annex C of [ITU-T G.992.5] (ADSL2plus).  Relationships  This is one of the status data MEs associated with an xDSL UNI. The ONU automatically  creates or deletes an instance of this ME upon creation or deletion of a PPTP xDSL UNI part  1 that supports these attributes.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. Through an  identical ID, this ME is implicitly linked to an instance of the PPTP xDSL UNI  part 1 ME. (R) (mandatory) (2 bytes)  FEXT downstream SNR margin: The FEXT SNRMds attribute is the downstream SNR  margin measured during FEXT R duration at the ATU -R. The attribute value  ranges from 0 ( –64.0 dB) to 1270 (+63.0  dB). The special value 0xFFFF  indicates that the attribute is out of range. (R) (mandatory) (2 bytes)  NEXT downstream SNR margin: The NEXT SNRMds attribute is the downstream SNR  margin measured during NEXT R duration at the ATU -R. The attribute value  ranges from 0 ( –64.0 dB) to 1270 (+63.0  dB). The special value 0xFFFF  indicates that the attribute is out of range. (R) (mandatory) (2 bytes)  FEXT upstream SNR margin: The FEXT SNRMus attribute is the upstream SNR margin  (see clause 7.5.1.16 of [ITU-T G.997.1]) measured during FEXTC duration at  the ATU-C. The attribute value ranges from 0 (–64.0 dB) to 1270 (+63.0 dB).  The special value 0xFFFF indicates that the attribute is out of range. (R)  (mandatory) (2 bytes)  NEXT upstream SNR margin: The NEXT SNRMus attribute is the upstream SNR margin  (see clause 7.5.1.16 of [ITU-T G.997.1]) measured during NEXTC duration at  the ATU-C. The attribute value ranges from 0 (–64.0 dB) to 1270 (+63.0 dB).  The special value 0xFFFF indicates that the attribute is out of range. (R)  (mandatory) (2 bytes)  FEXT downstream maximum attainable data rate: The FEXT ATTNDRds attribute is the  maximum downstream net data rate calculated from FEXT downstream 

---- page break ---- SNR(f) (see clause 7.5.1.28.3.1 of [ITU-T G.997.1]). The rate is coded in bits  per second. (R) (mandatory) (4 bytes)  NEXT downstream maximum attainable data rate : The NEXT ATTNDRds attribute is  the maximum downstream net data rate calculated from NEXT downstream  SNR(f) (see clause 7.5.1.28.3.2 of [ITU-T G.997.1]). The rate is coded in bits  per second. (R) (mandatory) (4 bytes)  FEXT upstream maximum attainable data rate : The FEXT ATTNDRus attribute is the  maximum upstream net data rate calculated from FEXT upstream SNR(f) (see  clause 7.5.1.28.6.1 of [ITU-T G.997.1]). The rate is coded in bit s per second.  (R) (mandatory) (4 bytes)  NEXT upstream maximum attainable data rate : The NEXT ATTNDRus attribute is the  maximum upstream net data rate calculated from NEXT upstream SNR(f) (see  clause 7.5.1.28.6.2 of [ITU-T G.997.1]). The rate is coded in bit s per second.  (R) (mandatory) (4 bytes)  FEXT downstr
```
