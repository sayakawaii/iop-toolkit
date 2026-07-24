# Managed Entity

## Identity
- ME ID: 53
- ME Name: Physical path termination point POTS UNI
- Source Section: 9.9.1
- Source Page: 383

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Get, Set, Test
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: Administrative state
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 2
- Name: Deprecated
- Size: 2 bytes
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 3
- Name: ARC
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 4
- Name: ARC interval
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 5
- Name: Impedance
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 6
- Name: Rx gain
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 7
- Name: Tx gain
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 8
- Name: Operational state
- Size: 1 byte
- Format: needs_review
- Access: R
- Category: optional

### Attribute 9
- Name: Hook state
- Size: 1 byte
- Format: needs_review
- Access: R
- Category: optional

### Attribute 10
- Name: POTS holdover time
- Size: 2 bytes
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 11
- Name: Loss of softswitch
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: optional

## Extraction Status
- Status: auto_extracted
- Attributes found: 11
- Review needed: false

## Raw Source

```
9.9.1 Physical path termination point POTS UNI  This ME represents a POTS UNI in the ONU, where a physical path terminates and physical path  level functions (analogue telephony) are performed.  The ONU automatically creates an instance of this ME per port as follows.  • When the ONU has POTS ports built into its factory configuration.  • When a cardholder is provisioned to expect a circuit pack of the POTS type. 

---- page break ---- • When a cardholder provisioned for plug-and-play is equipped with a circuit pack of the POTS  type. Note that the installation of a plug -and-play card may indicate the presence of POTS  ports via equipment ID as well as type, and indeed may cause the ONU to instantiate a port- mapping package that specifies POTS ports.  The ONU automatically deletes instances of this ME when a cardholder is neither provisioned to  expect a POTS circuit pack, nor is it equipped with a POTS circuit pack.  Relationships  An instance of this ME is associated with each real or pre -provisioned POTS port. Either a  SIP or a VoIP voice CTP links to the POTS UNI. Status is available from a VoIP line status  ME, and RTP and call control PM may be collected on this point.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. This 2 byte  number indicates the physical position of the UNI. The first byte is the slot ID  (defined in clause 9.1.5). The second byte is the port ID, with the range 1..255.  (R) (mandatory) (2 bytes)  Administrative state: This attribute shuts down (2), locks (1) and unlocks (0) the functions  performed by this ME. If the administrative state is set to shut down while the  POTS UNI line state is non -idle, no action is taken until the POTS UNI line  state changes to idle, whereupon the administrative state changes to locked. If  the administrative state is set to shut down and the POTS UNI line state is  already idle, the administrative state is immediately set to locked. In both cases,  the transition from shutting down to locked state is signalled with an AVC.  When the administrative state is set to lock, all user functions of this UNI are  blocked, and alarms, TCAs and AVCs for this ME and all dependent MEs are  no longer generated. Selection of a default value for this attribute is outside the  scope of this Recommendation. (R, W) (mandatory) (1 byte)  Deprecated: This attribute is not used and should not be supported. (R,  W) (optional)  (2 bytes)  ARC: See clause A.1.4.3. (R, W) (optional) (1 byte)  ARC interval: See clause A.1.4.3. (R, W) (optional) (1 byte)  Impedance: This attribute specifies the impedance for the POTS UNI. Valid values include  the following.  0 600 Ohm  1 900 Ohm  The following parameter sets from Annex C of [ETSI TS 101 270-1] are also  defined:  2 C1=150 nF, R1=750 Ohm, R2=270 Ohm  3 C1=115 nF, R1=820 Ohm, R2=220 Ohm  4 C1=230 nF, R1=1050 Ohm, R2=320 Ohm  where C1, R1, and R2 are related as shown in  Figure 9.9.1-1. Upon ME  instantiation, the ONU sets this attribute to 0. (R, W) (optional) (1 byte) 

---- page break ---- G.988(12)_F9.9.1-1   Figure 9.9.1-1 – Impedance model for POTS UNI    Transmission path: This attribute allows setting the POTS UNI either to full -time on-hook  transmission (0) or part-time on-hook transmission (1). Upon ME instantiation,  the ONU sets this attribute to 0. (R, W) (optional) (1 byte)  Rx gain: This attribute specifies a gain value for the received signal in the form of a 2s  complement number. Valid values are –120 (12.0 dB) to 60 (+6.0  dB). The  direction of the affected signal is in the D to A direction, towards the telephone  set. Upon ME instantiation, the ONU sets this attribute to 0. (R, W) (optional)  (1 byte)  Tx gain: This attribute specifies a gain value for the transmit signal in the form of a 2s  complement number. Valid values are –120 (12.0 dB) to 60 (+6.0  dB). The  direction of the affected signal is in the A to D direction, away from the  telephone set. Upon ME instantiation, the ONU sets this attribute to 0. (R, W)  (optional) (1 byte)  Operational state : This attribute indicates whether the ME is capable of performing its  function. Valid values are enabled (0) and disabled (1). (R) (optional) (1 byte)  Hook state: This attribute indicates the current state of the subscriber line: 0 = on hook, 1 =  off hook (R) (optional) (1 byte)  POTS holdover time: This attribute determines the time during which the POTS loop voltage  is held up when a LOS or softswitch connectivity is detected (please refer to  the following table for description of behaviours) . After the specified time  elapses, the ONU drops the loop voltage, and may thereby cause premises  intrusion alarm or fire panel circuits to go active. When the ONU ranges  successfully on the PON or softswitch connectivity is restored , it restores the  POTS loop voltage immediately and resets t he timer to zero. The attribute is  expressed in seconds. The default value 0 selects the vendor's factory policy.  (R, W) (optional) (2 bytes)    POTS holdover  time  Loss of  softswitch Behaviour  0 Don't care Vendor-specific  POTS holdover  time > 0   False T/R will be brought down on expiration of the holdover timer. The  holdover timer is started upon detection of LOS. T/R is restored  immediately upon ONU ranging. This setting is recommended for  burglar alarms.   POTS holdover  time > 0  True T/R will be brought down on expiration of the holdover timer. The  holdover timer is started on detection of softswitch connectivity keep  alive signal. T/R is restored immediately upon softswitch connectivity.  This setting is recommended for fire panels. 

---- page break ---- Nominal feed voltage: This attribute indicates the designed nominal feed voltage of the POTS  loop. It is an absolute value with resolution 1  V. This attribute does not  represent the actual voltage measured on the loop, which is available through  the test command. (R, W) (optional) (1 byte)  Loss of softswitch : This Boo
```
