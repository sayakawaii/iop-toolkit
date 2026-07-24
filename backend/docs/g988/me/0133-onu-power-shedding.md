# Managed Entity

## Identity
- ME ID: 133
- ME Name: ONU power shedding
- Source Section: 9.1.7
- Source Page: 80

## Classification
- Access: needs_review
- Type: needs_review
- Actions: needs_review
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: Restore power timer reset interval
- Size: 2 bytes
- Format: needs_review
- Access: R, W
- Category: mandatory

## Extraction Status
- Status: auto_extracted
- Attributes found: 1
- Review needed: true

## Raw Source

```
9.1.7 ONU power shedding  This ME models the ONU's ability to shed services when the ONU goes into battery operation mode  after AC power failure. Shedding classes are defined in the following table, which may span multiple  circuit pack types. This feature works in conjunction with the power shed override attribute of the  circuit pack ME, which can selectively prevent power shedding of priority ports.  An ONU that supports power shedding automatically creates an instance of this ME.  The following table defines the binding of shedding class and PPTP type. The coding is taken from  Table 9.1.5-1. In the case of hybrid circuit pack types, multiple shedding classes may affect a circuit  pack if the hardware is capable of partial power shedding.  An ONU may choose to model its ports with the port -mapping package of clause 9.1.8, rather than  with real or virtual circuit packs. In this case, power shedding pertains to individual PPTPs (listed in  column 2 of the table).   

---- page break ---- Shedding class PPTP type Coding Content  ATM ATM PPTP 1..12 Various ATM UNIs  CES CES PPTP 13 C1.5 (DS1)  14 C2.0 (E1)  15 C6.3 (J2)  16 C-DS1/E1  17 C-DS1/E1/J1  Data Ethernet PPTP 22 10BASE-T  23 100BASE-T  24 10/100 BASE-T  Frame Unspecified 25..27 Non-Ethernet LANs  CES CES PPTP 28 C1.5 (J1)  Sdh-sonet Sdh-sonet 29..31 ATM sdh-sonet interfaces  Voice POTS PPTP 32 POTS  ISDN PPTP 33 ISDN BRI (deprecated)  Data Ethernet PPTP 34 Gigabit optical Ethernet  DSL xDSL PPTP 35 xDSL  SHDSL 36 SHDSL  VDSL PPTP 37 ITU-T G.993.1 VDSL  N/A Video UNI 38 Radio frequency (RF) video service  N/A LCT PPTP 39 Local craft terminal   Data IEEE 802.11 PPTP 40 Wireless  Voice (DSL may also apply) xDSL + POTS 41 xDSL/POTS  VDSL + POTS 42 ITU-T G.993.1 VDSL/POTS  N/A Unspecified 43 Common equipment  Unspecified 44 Combined video, PON  Unspecified 45 Mixed services (Power shedding  based on port type)  Data MoCA PPTP 46 MoCA  Data Ethernet PPTP 47 10/100/1000 BASE-T  49 10G Ethernet  N/A PON PPTP 237..238 XG-PON ANIs  Video overlay Video ANI PPTP  Video return Video RPD  Relationships  One instance of this ME is associated with the ONU ME.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. There is only  one instance, number 0. (R) (mandatory) (2 bytes)  Restore power timer reset interval: The time delay, in seconds, before resetting the power- shedding timers after full power restoration. Upon ME instantiation, the ONU  sets this attribute to 0. (R, W) (mandatory) (2 bytes)
```
