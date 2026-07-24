# Managed Entity

## Identity
- ME ID: 103
- ME Name: xDSL channel upstream status data
- Source Section: 9.7.20
- Source Page: 292

## Classification
- Access: needs_review
- Type: needs_review
- Actions: needs_review
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: Actual interleaving delay
- Size: 1 byte
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 2
- Name: Actual data rate
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 3
- Name: Previous data rate
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

## Extraction Status
- Status: auto_extracted
- Attributes found: 3
- Review needed: true

## Raw Source

```
9.7.20 xDSL channel upstream status data  This ME contains upstream channel status data for an xDSL UNI. The ONU automatically creates or  deletes instances of this ME upon the creation or deletion of a PPTP xDSL UNI part 1. 

---- page break ---- Relationships  One or more instances of this ME are associated with an instance of an xDSL UNI.  Attributes  Managed entity  ID: This attribute uniquely identifies each instance of this ME. The two  MSBs of the first byte are the bearer channel ID. Excluding the first 2 bits of  the first byte, the remaining part of the ME ID is identical to that of this ME's  parent PPTP xDSL UNI part 1. (R) (mandatory) (2 bytes)  Actual interleaving delay: This attribute is the actual one-way interleaving delay introduced  by the PMS-TC between the alpha and beta reference points, excluding the L1  and L2 states. In the L1 and L2 states, this attribute contains the interleaving  delay in the previous L0 state . For ADSL, this attribute is derived from the S  and D attributes as cap( S*D)/4 ms, where S is the number of symbols per  codeword, D is the interleaving depth and cap() denotes rounding to the next  higher integer. For [ITU -T G.993.2], this attribute is computed according to  the formula in clause 9.7 of [ITU-T G.993.2]. The actual interleaving delay is  coded in m illiseconds, rounded to the nearest m illisecond. (R) (mandatory)  (1 byte)  Actual data rate : This parameter reports the actual net data rate of the bearer channel,  excluding the L1 and L2 states. In the L1 or L2 state, the parameter contains  the net data rate in the previous L0 state. The data rate is coded in bit s per  second. (R) (mandatory) (4 bytes)  Previous data rate: This parameter reports the previous net data rate of the bearer channel  just before the latest rate change event occurred, excluding transitions between  the L0 state and the L1 or L2 state. A rate change can occur at a power  management state transition, e.g., at full or short initialization, fast retrain or  power down, or at a dynamic rate adaptation. The rate is coded in bit s per  second. (R) (mandatory) (4 bytes)  Actual impulse noise protection: The ACTINP attribute reports the actual INP on the bearer  channel in the L0 state. In the L1 or L2 state, the attribute contains the INP in  the previous L0 state. The value is coded in fractions of DMT symbols with a  granularity of 0.1 symbols. The range  is from 0 (0.0 symbols) to 254 (25.4  symbols). The special value 255 indicates an ACTINP higher than 25.4. (R)  (mandatory for ITU -T G.993.2 VDSL2, optional for other xDSL  Recommendations that support it) (1 byte)  Impulse noise protection reporting mode : The INPREPORT attribute reports the method  used to compute the ACTINP. If set to 0, the ACTINP is computed according  to the INP_no_erasure formula (clause 9.6 of [ITU -T G.993.2]). If set to 1,  ACTINP is the value estimated by the xTU receiver. (R) (manda tory for   ITU-T G.993.2 VDSL2) (1 byte)  Actual size of Reed -Solomon codeword : The NFEC attribute reports the actual Reed - Solomon codeword size used in the latency path in which the bear er channel  is transported. Its value is coded in bytes in the range 0..255. (R) (mandatory  for ITU-T G.993.2 VDSL2, optional for others) (1 byte)  Actual number of Reed-Solomon redundancy bytes: The RFEC attribute reports the actual  number of Reed-Solomon redundancy bytes per codeword used in the latency  path in which the bearer channel is transported. Its value is co ded in bytes in  the range 0..16. The value 0 indicates no Reed -Solomon coding. (R)  (mandatory for ITU-T G.993.2 VDSL2, optional for others) (1 byte)
```
