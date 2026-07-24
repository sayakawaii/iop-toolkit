# Managed Entity

## Identity
- ME ID: 102
- ME Name: xDSL channel downstream status data
- Source Section: 9.7.19
- Source Page: 291

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Get
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

### Attribute 4
- Name: Actual impulse noise protection
- Size: 1 byte
- Format: needs_review
- Access: R
- Category: optional

## Extraction Status
- Status: auto_extracted
- Attributes found: 4
- Review needed: false

## Raw Source

```
9.7.19 xDSL channel downstream status data  This ME contains downstream channel status data for an xDSL UNI. The ONU automatically creates  or deletes instances of this ME upon the creation or deletion of a PPTP xDSL UNI part 1.  NOTE – [ITU-T G.997.1] specifies that bit rate attributes have a granularity of 1000  bit/s. If ITU-T G.997.1  compliance is required, the ONU should only report values with this granularity.  Relationships  One or more instances of this ME are associated with an instance of an xDSL UNI.  Attributes  Managed entity  ID: This attribute uniquely identifies each instance of this ME. The two  MSBs of the first byte are the bearer channel ID. Excluding the first 2 bits of  the first byte, the remaining part of the ME ID is identical to that of this ME's  parent PPTP xDSL UNI part 1. (R) (mandatory) (2 bytes)  Actual interleaving delay: This attribute is the actual one-way interleaving delay introduced  by the PMS-TC between the alpha and beta reference points, excluding delay  in the L1 and L2 states. In the L1 and L2 states, the attribute contains the  interleaving delay in the previous L0 state. For ADSL, this attribute is derived  from the S and D attributes as cap( S*D)/4 ms, where S is the number of  symbols per codeword, D is the interleaving depth and cap() denotes rounding  to the next higher integer. For [ITU -T G.993.2], this attribut e is computed  according to the formula in clause 9.7 of [ITU -T G.993.2]. The actual  interleaving delay is coded in milliseconds, rounded to the nearest millisecond.  (R) (mandatory) (1 byte)  Actual data rate : This parameter reports the actual net data rate of the bearer channel,  excluding the rate in the L1 and L2 states. In the L1 or L2 state, the parameter  contains the net data rate in the previous L0 state. The data rate is coded in bits  per second. (R) (mandatory) (4 bytes)  Previous data rate: This parameter reports the previous net data rate of the bearer channel  just before the latest rate change event occurred, excluding transitions between  the L0 state and the L1 or L2 states. A rate change can occur at a power  management state transition, e.g., at full or short initialization, fast retrain or  power down, or at a dynamic rate adaptation. The rate is coded in bit s per  second (R) (mandatory) (4 bytes)  Actual impulse noise protection: The ACTINP attribute reports the actual INP on the bearer  channel in the L0 state. In the L1 or L2 state, the attribute contains the INP in 

---- page break ---- the previous L0 state. The value of this attribute is a number of DMT symbols,  with a granularity of 0.1 symbols. Its range is from 0 (0.0 symbols) to 254 (25.4  symbols). The special value 255 indicates an ACTINP higher than 25.4. (R)  (optional for [ITU-T G.992.1], mandatory for other xDSL Recommendations  that support this attribute) (1 byte)  Actual size of Reed -Solomon codeword : The NFEC attribute reports the actual Reed - Solomon codeword size used in the latency path in which the bearer channel  is transported. The value is coded in bytes, and ranges from 0..255. (R)  (mandatory for ITU-T G.993.2 VDSL2, optional for others) (1 byte)  Actual number of Reed-Solomon redundancy bytes: The RFEC attribute reports the actual  number of Reed-Solomon redundancy bytes per codeword used in the latency  path in which the bearer channel is transported. The value is coded in bytes,  and ranges from 0..16. The value 0 indicates no Reed -Solomon coding. (R)  (mandatory for ITU-T G.993.2 VDSL2, optional for others) (1 byte)  Actual number of bits per symbol: The LSYMB attribute reports the actual number of bits  per symbol assigned to the latency path in which the bearer channel is  transported, excluding trellis overhead. The value is coded in bits, and ranges  from 0..65535. (R) (mandatory for [TU-T G.993.2 VDSL2, optional for others)  (2 bytes)  Actual interleaving depth : The INTLVDEPTH attribute reports the actual depth of the  interleaver used in the latency path in which the bearer channel is transported.  The value ranges from 1..4096 in steps of 1. The value 1 indicates no  interleaving. (R) (mandatory for ITU -T G.993.2 VDSL2, optional for others)  (2 bytes)  Actual interleaving block length : The INTLVBLOCK attribute reports the actual block  length of the interleaver used in the latency path in which the bearer channel  is transported. The value ranges from 4..255 in steps of 1. (R) (mandatory for  ITU-T G.993.2 VDSL2, undefined for others) (1 byte)  Actual latency path : The LPATH attribute reports the index of the actual latency path in  which the bearer channel is transported. Valid values are 0..3. In  [ITU-T G.992.1], the fast path is mapped to latency index 0; the interleaved  path to index 1. (R) (mandatory for ITU -T G.993.2 VDSL2, optional for  others) (1 byte)  Actual impulse noise protection against repetitive electrical impulse noise  (ACTINP_REIN): If retransmission is used in a given transmit direction, this  parameter reports the actual INP against REIN on the bearer channel. The INP  of this attribute is equal to the integer value multiplied by 0.1 symbols. Valid  values and usage are given in clause 7.5.2.9 of [ITU-T G.997.1] (R) (optional)  (1 byte)  Actions  Get  Notifications  None.  9.7.20 xDSL channel upstream status data  This ME contains upstream channel status data for an xDSL UNI. The ONU automatically creates or  deletes instances of this ME upon the creation or deletion of a PPTP xDSL UNI part 1.
```
