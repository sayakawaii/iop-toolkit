# Managed Entity

## Identity
- ME ID: 107
- ME Name: xDSL channel configuration profile
- Source Section: 9.7.7
- Source Page: 267

## Classification
- Access: needs_review
- Type: needs_review
- Actions: needs_review
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: Minimum data rate
- Size: 4 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 2
- Name: Maximum data rate
- Size: 4 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 3
- Name: Rate adaptation ratio
- Size: 1 byte
- Format: needs_review
- Access: R, W, set-by-create
- Category: optional

### Attribute 4
- Name: Maximum interleaving delay
- Size: 1 byte
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 5
- Name: Data rate threshold upshift
- Size: 4 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: optional

### Attribute 6
- Name: Minimum data rate in low-power state
- Size: 4 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 7
- Name: Minimum impulse noise protection
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 8
- Name: Minimum SOS bit rate downstream
- Size: 4 bytes
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 9
- Name: Minimum SOS bit rate upstream
- Size: 4 bytes
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 10
- Name: Downstream subcarrier mask 1
- Size: 16 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 11
- Name: Downstream subcarrier mask 2
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 12
- Name: Upstream subcarrier mask
- Size: 8 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

## Extraction Status
- Status: auto_extracted
- Attributes found: 12
- Review needed: true

## Raw Source

```
9.7.7 xDSL channel configuration profile  This ME contains the channel configuration profile for an xDSL UNI. An instance of this ME is  created and deleted by the OLT.  NOTE – If [ITU -T G.997.1] compatibility is required, bit rates should only be set to integer multiples of  1000 bits/s. The ONU may reject attempts to set other values for bit rate attributes.  Relationships  An instance of this ME may be associated with zero or more instances of the PPTP xDSL UNI  part 1.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. The value 0  is reserved. (R, set-by-create) (mandatory) (2 bytes)  Minimum data rate : This parameter specifies the minimum desired net data rate for the  bearer channel. It is coded in bit s per second. (R,  W, set-by-create)  (mandatory) (4 bytes)  Maximum data rate : This parameter specifies the maximum desired net data rate for the  bearer channel. It is coded in bit s per second. (R,  W, set-by-create)  (mandatory) (4 bytes)  Rate adaptation ratio: This attribute specifies the weight that should be taken into account   when performing rate adaptation in the direction of the bearer channel. The  attribute is defined as a percentage. The value 20, for example, means that 20%  of the available data rate (in excess of the minimum data rate summed over all  bearer channels) is assigned to this bearer channel and 80% to the other bearer  channels. The OLT must ensure that the sum of rate adaptation ratios over all  bearers in one direction is 100%. (R, W, set-by-create) (optional) (1 byte)  Maximum interleaving delay : This attribute is the maximum one -way interleaving delay  introduced by the PMS-TC between the alpha and the beta reference points, in  the direction of the bearer channel. The one-way interleaving delay is defined 

---- page break ---- in individual xDSL Recommendations as cap( S*D) /4 ms, where S is the S  factor, D is the interleaving depth, and cap() denotes rounding to the next  higher integer. xTUs choose S and D values such that the actual one -way  interleaving delay does not exceed the configured maximum interleaving  delay.  The delay is coded in milliseconds, varying from 2 to 63, with special meaning  assigned to values 0, 1 and 255. The value 0 indicates that no delay bound is  imposed. The value 1 indicates the fast latency path is to be used in the ITU-T  G.992.1 operating mode and S and D are to be selected such that S  1 and D =  1 in ITU -T G.992.2, ITU-T G.992.3, ITU-T G.992.4, ITU-T G.992.5 and  ITU-T G.993.2 operating modes. The value 255 indicates a delay bound of  1 ms in ITU-T G.993.2 operation. (R, W, set-by-create) (mandatory) (1 byte)  Data rate threshold upshift: This attribute is a threshold on the cumulative data rate upshift  achieved over one or more bearer channel data rate adaptations. An upshift rate  change (DRT up) notification is issued by the PPTP xDSL UNI part 1 when  the actual data rate exceeds the data rate at the last entry into showtime by more  than the threshold. The data rate threshold is coded in bit s per second. (R, W,  set-by-create) (mandatory for xDSL standards that use this attribute) (4 bytes)  Data rate threshold downshift : This attribute is a threshold on the cumulative data rate  downshift achieved over one or more bearer channel data rate adaptations. A  downshift rate change (DRT down) notification is issued by the PPTP xDSL  UNI part 1 when the actual data rate is below the data rate at the last entry into  showtime by more than the threshold. The data rate threshold is coded in bit s  per second. (R, W, set-by-create) (mandatory for xDSL standards that use this  attribute) (4 bytes)  Minimum reserved data rate: This attribute specifies the desired minimum reserved net data  rate for the bearer channel. The rate is coded in bit s per second. This attribute  is needed only if the rate adaptation mode is set to dynamic in the xDSL line  configuration profile part 1. (R, W, set-by-create) (optional) (4 bytes)  Minimum data rate in low-power state: This parameter specifies the minimum desired net  data rate for the bearer channel during the low-power state (L1/L2). The power  management low-power states L1 and L2 are defined in [ITU-T G.992.2] and  [ITU-T G.992.3], respectively. The data rate is coded in bits per second. (R, W,  set-by-create) (mandatory) (4 bytes)  Minimum impulse noise protection: The INPmin attribute specifies the minimum INP for the  bearer channel if it is transported over DMT symbols with a subcarrier spacing  of 4.3125 kHz. INP is expressed in DMT symbols with a subcarrier spacing of  4.3125 kHz. It can be ½ symbol or any integer number of symbols from 0 to  16, inclusive.  If the xTU does not support the configured INP min value, it uses the nearest  supported INP value greater than INPmin.  Value INPmin  1 0 symbols  2 ½ symbol  N (N – 2) symbols, 3 ≤ N ≤ 18  (R, W, set-by-create) (optional for [ITU -T G.992.1], mandatory for other  xDSL standards that use this attribute) (1 byte) 

---- page break ---- Maximum bit error ratio : This attribute specifies the desired maximum bit error ratio for  the bearer channel. It is only valid for [ITU -T G.992.3], [ITU-T G.992.4] and  [ITU-T G.992.5]. The bit error ratio is specified via the following values:  1 10–3  2 10–5  3 10–7  (R, W, set-by-create) (mandatory for standards that use this attribute) (1 byte)  Minimum impulse noise protection 8 kHz: The INPmin8 attribute specifies the minimum INP  for the bearer channel if it is transported over DMT symbols with a subcarrier  spacing of 8.625 kHz. It is only valid for [ITU-T G.993.2]. INP is expressed in  DMT symbols with a subcarrier spacing of 8.625 kHz. It can take any integer  value from 0 (default) to 16, inclusive. (R, W) (mandatory for  [ITU-T G.993.2]) (1 byte)  Maximum delay variation: The DVMAX attribute specifies the maximum value for delay  variation allowed in an OLR procedure. Its value ranges from 1 (0.1 ms) to 254  (25
```
