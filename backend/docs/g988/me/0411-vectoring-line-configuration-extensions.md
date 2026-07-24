# Managed Entity

## Identity
- ME ID: 411
- ME Name: Vectoring line configuration extensions
- Source Section: 9.7.34
- Source Page: 315

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Get
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: SHINERATIO_RTX
- Size: 1 bytes
- Format: needs_review
- Access: R, W
- Category: mandatory

## Extraction Status
- Status: auto_extracted
- Attributes found: 1
- Review needed: false

## Raw Source

```
9.7.34 Vectoring line configuration extensions  This ME extends the xDSL line configuration MEs with attributes that are specific to vectoring. An  instance of this ME is created and deleted by the OLT.  Relationships  An instance of this ME may be associated with zero or more instances of an xDSL UNI.  The overall xDSL line configuration profile is modelled in several parts, all of which are  associated together through a common ME ID (the client PPTP xDSL UNI part 1 has a single  pointer, which refers to the entire set of line configuration parts).  Attributes  Managed entity ID : This attribute uniquely identifies each instance of this ME. All xDSL  and VDSL2 and vectoring line configuration profiles and extensions that  pertain to a given PPTP xDSL UNI must share a common ME ID. (R,  set-by-create) (mandatory) (2 bytes)  Vectoring frequency-band control upstream (VECTOR_BAND_CONTROLus) table:  This configuration parameter is an array of pairs of sub-carrier indices [a(i),  b(i)]. Up to eight frequency bands may be configured. The same value of this  parameter shall be set for all lines of the same vector group. See  clause 7.3.1.13.1 of [ITU-T G.997.1].  This attribute is a table where each entry comprises:  – band number field, i (1 byte, range 1-8);  – band start subcarrier index, a(i) (2 bytes);  – band stop subcarrier index, b(i) (2 bytes).  The band number field is the table index. By default, the table is empty. Setting  a table entry with non -zero subcarrier indices implies insertion into the table.  Setting an entry's subcarrier indices to zero implies deletion from the table, if  present.  The maximum number of bands is eight, so the maximum size of the table is  40 bytes. (R, W) (mandatory) (N  5 bytes) 

---- page break ---- Vectoring frequency -band control downstream (VECTOR_BAND_CONTROLds)  table: This configuration parameter is an array of pairs of sub -carrier indices  [a(i), b(i)]. Up to eight frequency bands may be configured. The same value of  this parameter shall be set for all lines of the same vector group. See  clause 7.3.1.13.2 of [ITU-T G.997.1]. The format of this attribute is the same  as VECTOR_BAND_CONTROLus. The maximum number of bands is eight,  so the maximum size of the table is 40 bytes. (R,  W) (mandatory) ( N   5 bytes)  FEXT cancellation line priorities upstream (FEXT_CANCEL_PRIORITYus) : This  attribute specifies line priority for the line in the vectored group in the upstream  direction. Allowed values are 0 (LOW) and 1 (HIGH). See clause  7.3.1.13.3  of [ITU-T G.997.1]. (R, W) (optional) (1 byte)  FEXT cancellation line priorities downstream (FEXT_CANCEL_PRIORITYds) : This  attribute specifies line priority for the line in the vectored group in the  downstream direction. Allowed values are 0 (LOW) and 1 (HIGH). See clause  7.3.1.13.4 of [ITU-T G.997.1]. (R, W) (optional) (1 byte)  FEXT cancellation enabling/disabling upstream (FEXT_CANCEL_ENABLEus) : A  value of 1 enables and a value of 0 disables FEXT cancellation in the upstream  direction from all the other vectored lines into the line in the vectored group.  See clause 7.3.1.13.5 of [ITU-T G.997.1]. (R, W) (mandatory) (1 byte)  FEXT cancellation enabling/disabling downstream (FEXT_CANCEL_ENABLEds) : A  value of 1 enables and a value of 0 disables FEXT cancellation in the  downstream direction from all the other vectored lines into the line in the  vectored group. See clause 7.3.1.13.6 of [ITU-T G.997.1]. (R, W) (mandatory)  (1 byte)  Downstream requested XLIN subcarrier group size (XLINGREQds): This attribute is the  requested value of XLINGds. Valid values are given in clause 7.3.1.13.7 of  [ITU-T G.997.1]. (R, W) (mandatory) (1 byte)  Upstream requested XLIN subcarrier group size (XLINGREQus) : This attribute is the  requested value of XLINGus. Valid values are given in clause  7.3.1.13.8 of  [ITU-T G.997.1]. (R, W) (mandatory) (1 byte)  Vectoring mode enable (VECTORMODE_ENABLE): This attribute defines the vectoring  initialization type to be allowed by the VTU-O on the line. It is coded in a bit- map representation as defined in clause 7.3.1.13.9 of [ITU-T G.997.1]. (R, W)  (optional) (1 byte)  VCE ID (VCE_ID) : For the line in a vectored group, the VCE ID uniquely identifies the  VCE that manages and controls the vectored group to which the line belongs.  The valid range of values is given in clause 7.4.13.1 of [ITU -T G.997.1]. (R)  (mandatory) (1 byte)  VCE port index (VCE_port_index): For the line in a vectored group, the VCE port index is  the physical index that uniquely identifies the VCE port to which the line is  connected. The valid range of values is given in clause 7.4.13.2 of  [ITU-T G.997.1]. (R) (mandatory) (2 bytes)  Actions  Create, delete, get, get next, set  Set table (optional) 

---- page break ---- Notifications  None.  9.7.35 xDSL channel configuration profile part 2  This ME contains the channel configuration profile for an xDSL UNI. An instance of this ME is  created and deleted by the OLT.  NOTE – If [ITU -T G.997.1] compatibility is required, bit rates should only be set to integer multiples of  1000 bits/s. The ONU may reject attempts to set other values for bit rate attributes.  Relationships  An instance of this ME may be associated with zero or more instances of the PPTP xDSL UNI  part 1.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. Through an  identical ID, this ME is implicitly linked to an instance of the xDSL channel  configuration profile. (R, set-by-create) (mandatory) (2 bytes)  Minimum expected throughput for retransmission (MINETR_RTX): If retransmission is  used in a given transmit direction, this attribute specifies the minimum  expected throughput for the bearer channel, in bits  per second. See clause  7.3.2.1.8 of [ITU-T G.997.1]. (R, W) (mandatory) (4 bytes)  Maximum expected throughput for retransmission (MAXETR_RTX) : If retransmission  is used in a given transmit direction, this parameter
```
