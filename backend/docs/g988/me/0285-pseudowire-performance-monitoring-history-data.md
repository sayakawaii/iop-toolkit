# Managed Entity

## Identity
- ME ID: 285
- ME Name: Pseudowire performance monitoring history data
- Source Section: 9.8.8
- Source Page: 370

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
- Name: Transmitted packets
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 4
- Name: Missing packets
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 5
- Name: Misordered packets, usable
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 6
- Name: Misordered packets dropped
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 7
- Name: Playout buffer underruns/overruns
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 8
- Name: Malformed packets
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 9
- Name: Stray packets
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 10
- Name: Remote packet loss
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 11
- Name: TDM L-bit packets transmitted
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 12
- Name: ES
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 13
- Name: SES
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 14
- Name: UAS
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

## Extraction Status
- Status: auto_extracted
- Attributes found: 14
- Review needed: false

## Raw Source

```
9.8.8 Pseudowire performance monitoring history data  This ME collects PM for a pseudowire TP. Most of the attributes monitor packets received from the  PSN, and may therefore be considered egress PM. For the most part, ingress PM is collected at the  CES PPTP ME.  NOTE – The pseudowire PM history data ME collects data similar, but not identical, to that available from the  MAC bridge port PM history data ME associated with a MAC bridge. When the pseudowire is bridge -based,  it may not be necessary to collect both.  Instances of this ME are created and deleted by the OLT.  For a complete discussion of generic PM architecture, refer to clause I.4.  Relationships  An instance of this ME is associated with an instance of the pseudowire TP.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. Through an  identical ID, this ME is implicitly linked to an instance of the pseudowire TP.  (R, set-by-create) (mandatory) (2 bytes)  Interval end time : This attribute identifies the most recently finished 15  min interval. (R)  (mandatory) (1 byte)  Threshold data 1/2 ID: This attribute points to an instance of the threshold data 1 and 2 MEs  that contains PM threshold values. (R, W, set-by-create) (mandatory) (2 bytes) 

---- page break ---- Received packets : This attribute counts the total number of packets, both payload and  signalling, received in the PSN to the TDM direction. (R) (mandatory)  (4 bytes)  Transmitted packets: This attribute counts the total number of packets, both payload and  signalling, transmitted in the TDM to the PSN direction. The count includes  packets whose L bit is set and which may therefore not contain a payload. (R)  (mandatory) (4 bytes)  Missing packets: This attribute counts the number of lost packets, as indicated by gaps in the  control word numbering sequence. Both payload and signalling packets, if any,  contribute to this count. (R) (mandatory) (4 bytes)  Misordered packets, usable : This attribute counts the number of packets received out of  order, but which were able to be successfully re-ordered and played out. Both  payload and signalling packets, if any, contribute to this count. (R)  (mandatory) (4 bytes)  Misordered packets dropped: This attribute counts the number of packets received out of  sequence that were discarded, either because the ONU did not support  reordering or because it was too late to reorder them. Both payload and  signalling packets, if any, contribute to this count. (R) (mandatory) (4 bytes)  Playout buffer underruns/overruns: This attribute counts the number of packets that were  discarded because they arrived too late or too early to be played out. Both  payload and signalling packets, if any, contribute to this count. (R)  (mandatory) (4 bytes)  Malformed packets: This attribute counts the number of malformed packets, e.g., because  the packet length was not as expected or because of an unexpected RTP  payload type. Both payload and signalling packets, if any, contribute to this  count. (R) (mandatory) (4 bytes)  Stray packets: This attribute counts the number of packets whose ECID or RTP SSRC failed  to match the expected value, or which are otherwise known to have been  misdelivered. Stray packets are discarded without affecting any of the other  PM counters. Both payload and si gnalling packets, if any, contribute to this  count. (R) (mandatory) (4 bytes)  Remote packet loss: This attribute counts received packets whose R bit is set, indicating the  loss of packets at the far end. Both payload and signalling packets, if any,  contribute to this count. (R) (mandatory) (4 bytes)  TDM L-bit packets transmitted : This attribute counts the number of packets transmitted  with the L bit set, indicating a near -end TDM fault. Both payload and  signalling packets, if any, contribute to this count. (R) (mandatory) (4 bytes)  ES: This attribute counts errored seconds. Any discarded, lost, malformed or  unusable packet received from the PSN during a given second causes this  counter to increment. Both payload and signalling packets, if any, contribute  to this count. (R) (mandatory) (4 bytes)  SES: This attribute counts severely errored seconds. The criterion for a n SES may  be configured through the pseudowire maintenance profile ME. Both payload  and signalling packets, if any, contribute to this count. (R) (mandatory)  (4 bytes)  UAS: This attribute counts unavailable seconds. An unavailable second begins at the  onset of 10 consecutive SES and ends at the onset of 10 consecutive seconds 

---- page break ---- that are not severely errored. A service is unavailable if either its payload or its  signalling, if any, are unavailable. During unavailable time, only UAS should  be counted; other anomalies should not be counted. (R) (mandatory) (4 bytes)  Actions  Create, delete, get, set  Get current data (optional)  Notifications  Threshold crossing alert  Alarm  number Threshold crossing alert Threshold value attribute No. (Note)  0 Missing packets 1  1 Misordered packets, usable 2  2 Misordered packets dropped 3  3 Playout buffer underruns/overruns 4  4 Malformed packets 5  5 Stray packets 6  6 Remote packet loss 7  7 ES 8  8 SES 9  9 UAS 10  NOTE – This number associates the TCA with the specified threshold value attribute of the  threshold data 1/2 managed entities.  
```
