# Managed Entity

## Identity
- ME ID: 105
- ME Name: xDSL line configuration profile part 2
- Source Section: 9.7.3
- Source Page: 247

## Classification
- Access: needs_review
- Type: needs_review
- Actions: needs_review
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: Power management state forced
- Size: 1 byte
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 2
- Name: Power management state enabling
- Size: 1 byte
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 3
- Name: Downstream target noise margin
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 4
- Name: Upstream target noise margin
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 5
- Name: Upstream maximum noise margin
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 6
- Name: Downstream minimum noise margin
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 7
- Name: Upstream minimum noise margin
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 8
- Name: Downstream rate adaptation mode
- Size: 1 byte
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 9
- Name: Upstream rate adaptation mode
- Size: 1 byte
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 10
- Name: Downstream upshift noise margin
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: optional

### Attribute 11
- Name: Upstream upshift noise margin
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: optional

### Attribute 12
- Name: Minimum overhead rate upstream
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set- by-create
- Category: optional

### Attribute 13
- Name: Minimum overhead rate downstream
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: optional

### Attribute 14
- Name: Downstream minimum time interval for upshift rate adaptation
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: optional

### Attribute 15
- Name: Upstream minimum time interval for upshift rate adaptation
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: optional

### Attribute 16
- Name: Downstream downshift noise margin
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: optional

### Attribute 17
- Name: Upstream downshift noise margin
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: optional

### Attribute 18
- Name: Downstream minimum time interval for downshift rate adaptation
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: optional

### Attribute 19
- Name: Upstream minimum time interval for downshift rate adaptation
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: optional

### Attribute 20
- Name: xTU impedance state forced
- Size: 1 byte
- Format: needs_review
- Access: R, W, set-by-create
- Category: optional

### Attribute 21
- Name: L2-time
- Size: 1 byte
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 22
- Name: Downstream maximum nominal power spectral density
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 23
- Name: Upstream maximum nominal power spectral density
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 24
- Name: Downstream maximum nominal aggregate transmit power
- Size: 1 byte
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 25
- Name: Upstream maximum nominal aggregate transmit power
- Size: 1 byte
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 26
- Name: Upstream maximum aggregate  receive power
- Size: 2 bytes
- Format: needs_review
- Access: R, W set-by-create
- Category: mandatory

### Attribute 27
- Name: VDSL2 transmission system enabling
- Size: 1 byte
- Format: needs_review
- Access: R, W, set-by-create
- Category: optional

### Attribute 28
- Name: Automode cold start forced
- Size: 1 byte
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 29
- Name: Force INP downstream
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 30
- Name: Update request flag for far-end test parameters
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 31
- Name: INM inter-arrival time offset upstream
- Size: 2 bytes
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 32
- Name: INM inter-arrival time step upstream
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: optional

## Extraction Status
- Status: auto_extracted
- Attributes found: 32
- Review needed: true

## Raw Source

```
9.7.3 xDSL line configuration profile part 1  The overall xDSL line configuration profile is modelled in several parts, all of which are associated  together through a common ME ID (the client PPTP xDSL UNI part 1 has a single pointer, which  refers to the entire set of line configuration profile parts).  It is worth noting that attributes in the line configuration profile family affect the real -time service  delivery of an xDSL UNI, e.g., by triggering diagnostics. Despite the fact that they are called profiles,  it may be advisable to instantiate a complete set of these MEs for each PPTP xDSL UNI.  Relationships  An instance of this ME may be associated with zero or more instances of an xDSL UNI.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. The value 0  is reserved. All xDSL and VDSL2 line configuration profiles and extensions 

---- page break ---- that pertain to a given PPTP xDSL UNI must share a common ME ID. (R,  set-by-create) (mandatory)  xTU transmission system enabling (xTSE) : This configuration attribute specifies the  transmission system coding types to be allowed by the near -end xTU. It is a  bit map as defined in Table 9.7.12-1. (R,  W, set-by-create) (mandatory)  (7 bytes)  NOTE 1 – This attribute is only 7  bytes long. An eighth byte enabling VDSL2  capabilities is defined in the VDSL2 transmission system enabling attribute of the  xDSL line configuration profile part 2 ME.  Power management state forced : This configuration parameter forces the line state of the  near-end xTU. It is coded as an integer value with the following definition.  0 Force the line from the L3 idle state to the L0 full-on state. This  transition requires the short initialization procedures. After reaching  the L0 state, the line may enter into or exit from the L2 low-power  state if the L2 state is enabled. If the L0 state is not reached after a  vendor-discretionary number of retries or within a  vendor-discretionary timeout, an initialization failure occurs.  Whenever the line is in the L3 state, it attempts to transition to the L0  state until it is forced into another state through this configuration  parameter.  2 Force the line from the L0 full-on to the L2 low-power state. This is  an out-of-service test value for triggering the L2 mode.  3 Force the line from the L0 full-on or L2 low-power state to the L3  idle state. This transition requires the orderly shutdown procedure.  After reaching the L3 state, the line remains there until it is forced  into another state through this configuration parameter.  (R, W, set-by-create) (mandatory) (1 byte)  Power management state enabling : The PMMode attribute specifies the line states into  which the xTU-C or x digital subscriber line transceiver unit at the remote end  (xTU-R) may autonomously go. It is a bit map (0 if not allowed, 1 if allowed)  with the following definition.  Bit 1 (LSB): L3 idle state  Bit 2: L1/L2 low-power state  (R, W, set-by-create) (mandatory) (1 byte)  Downstream target noise margin : This attribute specifies the noise margin the xTU -R  receiver must achieve, relative to the BER requirement for each of the  downstream bearer channels, to successfully complete initialization. Its value  ranges from 0 (0.0 dB) to 310 (31.0 dB). (R,  W, set-by-create) (mandatory)  (2 bytes)  Upstream target noise margin: This attribute specifies the noise margin the xTU-C receiver  must achieve, relative to the BER requirement for each of the upstream bearer  channels, to successfully complete initialization. Its value ranges from 0 (0.0  dB) to 310 (31.0 dB). (R, W, set-by-create) (mandatory) (2 bytes) 

---- page break ---- Downstream maximum noise margin: The MAXSNRMds attribute specifies the maximum  noise margin the xTU-R receiver tries to sustain. If the noise margin is above  this level, the xTU-R requests the xTU-C to reduce its transmit power, if this  functionality is supported by the applicable xDSL Recommendation. Its value  ranges from 0 (0.0 dB) to 310 (31.0 dB). The special value 0xFFFF indicates  that the maximum noise margin limit is unbounded. (R,  W, set-by-create)  (mandatory) (2 bytes)  Upstream maximum noise margin : The MAXSNRMus attribute specifies the maximum  noise margin the xTU-C receiver tries to sustain. If the noise margin is above  this level, the xTU-C requests the xTU-R to reduce its transmit power, if this  functionality is supported by the applicable xDSL Recommendation. Its value  ranges from 0 (0.0 dB) to 310 (31.0 dB). The special value 0xFFFF indicates  that the maximum noise margin limit is unbounded. (R,  W, set-by-create)  (mandatory) (2 bytes)  Downstream minimum noise margin : This attribute specifies the minimum noise margin  the xTU-R receiver must tolerate. If the noise margin falls below this level, the  xTU-R requests the xTU -C to increase its transmit power. If an increase in  xTU-C transmit power is not possible, a loss-of-margin (LOM) defect occurs,  the xTU -R fails and attempts to re -initialize, and the PPTP declares a line  initialization failure (LINIT) alarm. Its value ranges from 0 (0.0 dB) to 310  (31.0 dB). (R, W, set-by-create) (mandatory) (2 bytes)  Upstream minimum noise margin : This attribute specifies the minimum noise margin the  xTU-C receiver must tolerate. If the noise margin falls below this level, the  xTU-C requests the xTU -R to increase its transmit power. If an increase in  xTU-R transmit power is not possible, an LOM defect occurs, the xTU-C fails  and attempts to re -initialize, and the PPTP declares a LINIT alarm. Its value  ranges from 0 (0.0 dB) to 310 (31.0 dB). (R,  W, set-by-create) (mandatory)  (2 bytes)  Downstream rate adaptation mode : The RA -MODEds attribute specifies the mode of  operation of a rate -adaptive xTU-C in the transmit direction. The parameter  can take four values.  1 Mode 1: MANUAL – Rate changed manually.  At start-up  The minimum data rate attribute of the associated xDSL channel  
```
