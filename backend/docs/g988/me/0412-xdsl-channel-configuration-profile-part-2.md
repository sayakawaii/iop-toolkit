# Managed Entity

## Identity
- ME ID: 412
- ME Name: xDSL channel configuration profile part 2
- Source Section: 9.7.35
- Source Page: 317

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
Notifications  None.  9.7.35 xDSL channel configuration profile part 2  This ME contains the channel configuration profile for an xDSL UNI. An instance of this ME is  created and deleted by the OLT.  NOTE – If [ITU -T G.997.1] compatibility is required, bit rates should only be set to integer multiples of  1000 bits/s. The ONU may reject attempts to set other values for bit rate attributes.  Relationships  An instance of this ME may be associated with zero or more instances of the PPTP xDSL UNI  part 1.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. Through an  identical ID, this ME is implicitly linked to an instance of the xDSL channel  configuration profile. (R, set-by-create) (mandatory) (2 bytes)  Minimum expected throughput for retransmission (MINETR_RTX): If retransmission is  used in a given transmit direction, this attribute specifies the minimum  expected throughput for the bearer channel, in bits  per second. See clause  7.3.2.1.8 of [ITU-T G.997.1]. (R, W) (mandatory) (4 bytes)  Maximum expected throughput for retransmission (MAXETR_RTX) : If retransmission  is used in a given transmit direction, this parameter specifies the maximum  expected throughput for the bearer channel, in bits  per second. See clause  7.3.2.1.9 of [ITU-T G.997.1]. (R, W) (mandatory) (4 bytes)  Maximum net data rate for retransmission (MAXNDR_RTX) : If retransmission is used  in a given transmit direction, this parameter specifies the maximum net data  rate for the bearer channel, in bits  per second. See clause 7.3.2.1.10 of  [ITU-T G.997.1]. (R, W) (mandatory) (4 bytes)  Maximum delay for retransmission (DELAYMAX_RTX) : If retransmission is used in a  given transmit direction, this parameter specifies the maximum for the  instantaneous delay due to the effect of retransmission only. This delay is  defined as the integer value of this attribute multiplied by 1 ms. The valid delay  values are given in clause 7.3.2.11 of [ITU -T G.997.1]. (R,  W) (mandatory)  (1 bytes)  Minimum delay for retransmission (DELAYMIN_RTX) : If retransmission is used in a  given transmit direction, this parameter specifies the minimum for the  instantaneous delay due to the effect of retransmission only. This delay is  defined as the integer value of this attribute multiplied by 1 ms. The valid delay  values are given in clause 7.3.2.12 of [ITU -T G.997.1]. (R,  W) (mandatory)  (1 bytes)  Minimum impulse noise protection against single high impulse noise event (SHINE) for  retransmission (INPMIN_SHINE_RTX) : If retransmission is used in a  given transmit direction, this parameter specifies the minimum INP against a  SHINE for the bearer channel if it is transported over DMT symbols with a  subcarrier spacing of 4.3125  kHz. The valid range of values is given in  clause 7.3.2.13 of [ITU-T G.997.1]. (R, W) (mandatory) (1 bytes) 

---- page break ---- Minimum impulse noise protection against SHINE for retransmission for systems using  8.625 kHz subcarrier spacing (INPMIN8_SHINE_RTX): If retransmission  is used in a given transmit direction, this parameter specifies the minimum INP  against SHINE for the bearer channel if it is transported over DMT symbols  with a subcarrier spacing of 8.625  kHz. The valid range of values is given in  clause 7.3.2.14 of [ITU-T G.997.1]. (R, W) (mandatory) (1 bytes)  SHINERATIO_RTX: If retransmission is used in a given transmit direction, this parameter  specifies the SHINE ratio. This ratio is defined as the integer value of this  attribute multiplied by 0.001. The valid range of values is given in  clause 7.3.2.15 of [ITU-T G.997.1]. (R, W) (mandatory) (1 bytes)  Minimum impulse noise protection against REIN for retransmission  (INPMIN_REIN_RTX): If retransmission is used in a given transmit  direction, this parameter specifies the minimum INP against REIN for the  bearer channel if it is transported over DMT symbols with a subcarrier spacing  of 4.3125 kHz. The valid range of values is given in clause 7.3.2.16 of [ITU-T  G.997.1]. (R, W) (mandatory) (1 bytes)  Minimum impulse noise protection against REIN for retransmission for systems using  8.625 kHz subcarrier spacing (INPMIN8_REIN_RTX): If retransmission is  used in a given transmit direction, this parameter specifies the minimum INP  against REIN for the bearer channel if it is transported over DMT symbols with  a subcarrier spacing of 8.625 kHz. The valid range of values is given in clause  7.3.2.17 of [ITU-T G.997.1]. (R, W) (mandatory) (1 bytes)  REIN inter-arrival time for retransmission (IAT_REIN_RTX): If retransmission is used  in a given transmit direction, this parameter specifies the IAT that shall be  assumed for REIN protection. The valid range of values is given in clause  7.3.2.18 of [ITU-T G.997.1]. (R, W) (mandatory) (1 bytes)  Target net data rate (TARGET_NDR) : If retransmission is not used in a given  transmit  direction, this parameter specifies the target net data of the bearer channel, in  bits per second. See clause 7.3.2.19.1 of [ITU-T G.997.1]. (R, W) (mandatory)  (4 bytes)  Target expected throughput for retransmission (TARGET_ETR) : If retransmission is  used in a given transmit direction, this parameter specifies the target expected  throughput for the bearer channel, in bits per second. See clause 7.3.2.19.2 of  [ITU-T G.997.1]. (R, W) (mandatory) (4 bytes)  Actions  Create, delete, get, set  Notifications  None.  9.7.36 xTU data gathering configuration  This ME defines configurations specific to data gathering.  An instance of this ME is created and deleted by the OLT.  Relationships  An instance of this ME may be associated with zero or more instances of the PPTP xDSL UNI  part 1. 

---- page break ---- Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. Through an  identical ID, this ME is implicitly linked to an instance of the PPTP xDSL UNI  part 1 ME. (R, set-by-create) (mandatory) (2 bytes)  Logging depth event percenta
```
