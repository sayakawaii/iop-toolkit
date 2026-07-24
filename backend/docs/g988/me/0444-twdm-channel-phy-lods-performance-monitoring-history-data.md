# Managed Entity

## Identity
- ME ID: 444
- ME Name: TWDM channel PHY/LODS performance monitoring history data
- Source Section: 9.16.3
- Source Page: 489

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
- Name: BIP-32 bit error count
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 4
- Name: Corrected PSBd HEC error count
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 5
- Name: Uncorrectable PSBd HEC error count
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 6
- Name: Corrected downstream FS header HEC error count
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 7
- Name: Uncorrectable downstream FS header HEC error count
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 8
- Name: Total number of LODS events
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 9
- Name: LODS events restored in operating TWDM channel
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 10
- Name: LODS events restored in protection TWDM channel
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 11
- Name: LODS events resulting in reactivation
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 12
- Name: LODS events resulting in reactivation after retuning to protection TWDM channel
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

## Extraction Status
- Status: auto_extracted
- Attributes found: 12
- Review needed: false

## Raw Source

```
9.16.3 TWDM channel PHY/LODS performance monitoring history data  This ME collects certain PM data associated with the slot/circuit pack, hosting on e or more ANI-G  MEs, and a specific TWDM channel. Instances of this ME are created and deleted by the OLT.  For a complete discussion of generic PM architecture, refer to clause I.4.  Relationships  An instance of this ME is associated with an instance of TWDM channel ME.   Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. Through an  identical ID, this ME is implicitly linked to an instance of the TWDM channel  ME. (R, set-by-create) (mandatory) (2 bytes)  Interval end time : This attribute identifies the most recently finished 15  min interval. (R)  (mandatory) (1 byte)  Threshold data 1/2 ID: This attribute points to an instance of the threshold data 1 and 2 MEs  that contains PM threshold values. (R, W, set-by-create) (mandatory) (2 bytes)  Total received words protected by bit-interleaved parity-32 (BIP-32): The count of 4 byte  words included in BIP -32 check. This is a product of the number of  downstream FS frames received by the size of the downstream FS frame after  the FEC parity byte, if any, have been removed. The count applies to the entire  downstream data flow, whether or not addressed to that ONT. (R) (mandatory)  (8 bytes)  BIP-32 bit error count : Count of the bit errors in the received downstream FS frames as  measured using BIP-32. If FEC is supported in the downstream direction, the  BIP-32 count applies to the downstream FS frame after the FEC correction has  been applied and the FEC parity bytes have been removed. (R) (mandatory)  (4 bytes)  Corrected PSBd HEC error count: The count of the errors in either CFC or OCS fields of  the PSBd block that have been corrected using the HEC technique. (R)  (mandatory) (4 bytes)  Uncorrectable PSBd HEC error count: The count of the errors in either CFC or OCS fields  of the PSBd block that could not be corrected using the HEC technique. (R)  (mandatory) (4 bytes)  Corrected downstream FS header HEC error count : The count of the errors in the  downstream FS header that have been corrected using the HEC technique. (R)  (mandatory) (4 bytes)  Uncorrectable downstream FS header HEC error count : The count of the errors in the  downstream FS header that could not be corrected using the HEC technique.  (R) (mandatory) (4 bytes)  Total number of LODS events : The count of the state transitions from O5.1/O5.2 to O6,  referring to the ONU activation cycle state machine, clause 12  of [ITU - T G.989.3]. (R) (mandatory) (4 bytes)  LODS events restored in operating TWDM channel : The count of LODS events cleared  automatically without retuning. (R) (mandatory) (4 bytes)  LODS events restored in protection TWDM channel: The count of LODS events resolved  by retuning to a pre -configured protection TWDM channel. The event is  counted against the original operating channel. (R) (mandatory) (4 bytes) 

---- page break ---- LODS events restored in discretionary TWDM channel : The count of LODS events  resolved by retuning to a TWDM channel chosen by the ONU, without  retuning. Implies that the wavelength channel protection for the operating  channel is not active. The event is counted against the original operating  channel (R) (mandatory) (4 bytes)  LODS events resulting in reactivation: The count of LODS events resolved through ONU  reactivation; that is, either TO2 (without WLCP) or TO3 + TO4 (with WLCP)  expires before the downstream channel is reacquired, referring to the ONU  activation cycle state machine, clause 12  of [ITU -T G.989.3]. The event is  counted against the original operating channel (R) (mandatory) (4 bytes)  LODS events resulting in reactivation after retuning to protection TWDM channel: The  count of LODS events resolved through ONU reactivation after attempted  protection switching, which turns unsuccessful due to a handshake failure. (R)  (mandatory) (4 bytes)  LODS events resulting in reactivation after retuning to discretionary TWDM channel :  The count of LODS events resolved through ONU reactivation after attempted  retuning to a discretionary channel, which turns unsuccessful due to a  handshake failure. (R) (mandatory) (4 bytes)  Actions  Create, delete, get, set  Get current data (optional)  Notifications  Threshold crossing alert  Alarm  number Threshold crossing alert  Threshold value  attribute No.  (Note)  0 N/A   1 BIP-32 bit error count 2  2 PSBd HEC errors – corrected 3  3 PSBd HEC errors – uncorrectable 4  4 FS header errors – corrected 5  5 FS header errors – uncorrectable 6  6 Total LODS event count 7  7 LODS – restored in operating TWDM channel 8  8 LODS – restored in protection TWDM channel 9  9 LODS – restored in discretionary TWDM channel 10  10 LODS – reactivations 11  11 LODS – handshake failure in protection channel 12  12 LODS – handshake failure in discretionary channel 13  NOTE – This number associates the TCA with the specified threshold value attribute  of the threshold data 1/2 managed entities. 

---- page break ---- 
```
