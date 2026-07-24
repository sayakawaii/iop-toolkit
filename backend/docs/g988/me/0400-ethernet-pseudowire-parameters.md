# Managed Entity

## Identity
- ME ID: 400
- ME Name: Ethernet pseudowire parameters
- Source Section: 9.8.18
- Source Page: 382

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Create, Delete, Get, Set
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: MTU
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

## Extraction Status
- Status: auto_extracted
- Attributes found: 1
- Review needed: false

## Raw Source

```
9.8.18 Ethernet pseudowire parameters  This ME contains the Ethernet pseudowire parameters. Instances of this ME are created and deleted  by the OLT.  Relationships  An instance of this ME is associated with an instance of the  PW Ethernet configuration data  ME.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. Through an  identical ID, this ME is implicitly linked to an instance of the  PW Ethernet  configuration data ME. (R, set-by-create) (mandatory) (2 bytes)  MTU: This attribute identifies the maximum transmission unit (bytes) that can be  received from the CPE in the upstream direction. Larger frames are discarded.  (R, W, set-by-create) (mandatory) (2 bytes) 

---- page break ---- Actions  Create, delete, get, set  Notifications  None.  9.9 Voice services  This clause defines MEs associated with a POTS (VoIP service), as shown in Figure 9.9-1.  G.988(12)_F9.9-1 SIP configuration portal 9.9.19 Points to: (3x) Network address Points to: Authentication security method Large string (AoR) Network address N N N N N N N N Points to: (3x) Large string Network address TCP/UDP configuration data Points to: (2x) Network address TCP/UDP configuration data Points to: Network address 9.9.13 RTP PM history data 9.9.10 Network dial plan table 9.9.12 Call control PM history data 9.9.11 V oIP line status 9.9.1 PPTP POTS UNI 9.9.4 V oIP voice CTP 9.9.2 SIP user data 9.9.5 V oIP media profile 9.9.9 V oIP feature access codes 9.9.18 V oIP config data  RTP  profile data 9.9.7  V oice  service profile 9.9.6 SIP call initiation PM history data 9.9.15 V oIP application service profile 9.9.8 SIP agent PM history data 9.9.14 SIP agent configuration data 9.9.3 MGC PM history data 9.9.17  MGC  configuration data 9.9.16 MGC configuration portal 9.9.20   Figure 9.9-1 – Managed entities associated with a POTS (VoIP service)  9.9.1 Physical path termination point POTS UNI  This ME represents a POTS UNI in the ONU, where a physical path terminates and physical path  level functions (analogue telephony) are performed.  The ONU automatically creates an instance of this ME per port as follows.  • When the ONU has POTS ports built into its factory configuration.  • When a cardholder is provisioned to expect a circuit pack of the POTS type.
```
