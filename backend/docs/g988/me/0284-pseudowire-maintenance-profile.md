# Managed Entity

## Identity
- ME ID: 284
- ME Name: Pseudowire maintenance profile
- Source Section: 9.8.7
- Source Page: 368

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Create, Delete, Get, Set
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: Jitter buffer maximum depth
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: optional

### Attribute 2
- Name: Jitter buffer desired depth
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: optional

### Attribute 3
- Name: Misconnected packets declaration policy
- Size: 1 byte
- Format: needs_review
- Access: R, W, set-by-create
- Category: optional

### Attribute 4
- Name: Misconnected packets clear policy
- Size: 1 byte
- Format: needs_review
- Access: R, W, set-by-create
- Category: optional

### Attribute 5
- Name: Loss of packets declaration policy
- Size: 1 byte
- Format: needs_review
- Access: R, W, set-by-create
- Category: optional

### Attribute 6
- Name: Loss of packets clear policy
- Size: 1 byte
- Format: needs_review
- Access: R, W, set-by-create
- Category: optional

### Attribute 7
- Name: Buffer overrun/underrun declaration policy
- Size: 1 byte
- Format: needs_review
- Access: R, W, set-by-create
- Category: optional

### Attribute 8
- Name: Buffer overrun/underrun clear policy
- Size: 1 byte
- Format: needs_review
- Access: R, W, set-by-create
- Category: optional

### Attribute 9
- Name: Malformed packets declaration policy
- Size: 1 byte
- Format: needs_review
- Access: R, W, set-by-create
- Category: optional

### Attribute 10
- Name: Malformed packets clear policy
- Size: 1 byte
- Format: needs_review
- Access: R, W, set-by-create
- Category: optional

### Attribute 11
- Name: R-bit transmit set policy
- Size: 1 byte
- Format: needs_review
- Access: R, W, set-by-create
- Category: optional

### Attribute 12
- Name: R-bit transmit clear policy
- Size: 1 byte
- Format: needs_review
- Access: R, W, set-by-create
- Category: optional

### Attribute 13
- Name: L bit receive policy
- Size: 1 byte
- Format: needs_review
- Access: R, W, set-by-create
- Category: optional

### Attribute 14
- Name: SES threshold
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: optional

## Extraction Status
- Status: auto_extracted
- Attributes found: 14
- Review needed: false

## Raw Source

```
9.8.7 Pseudowire maintenance profile  The pseudowire maintenance profile permits the configuration of pseudowire service exception  handling. It is created and deleted by the OLT.  The settings, and indeed existence, of a pseudowire maintenance profile affect the behaviour of the  pseudowire PM history data ME only in establishing criteria for counting SESs, but in no other way.  The pseudowire maintenance profile primarily affects the alarms declared by the subscribing  pseudowire TP.  Relationships  One or more instances of the pseudowire TP may point to an instance of the pseudowire  maintenance profile. If the pseudowire TP does not refer to a pseudowire maintenance profile,  the ONU's default exception handling is implied.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. The value 0  is reserved. (R, set-by-create) (mandatory) (2 bytes)  Jitter buffer maximum depth : This attribute specifies the desired maximum depth of the  playout buffer in the PSN to the TDM direction. The value is expressed as a  multiple of the 125  μs frame rate. The default value 0 selects the ONU 's  internal policy. (R, W, set-by-create) (optional) (2 bytes)  Jitter buffer desired depth : This attribute specifies the desired nominal fill depth of the  playout buffer in the PSN to the TDM direction. The value is expressed as a  multiple of the 125  μs frame rate. The default value 0 selects the ONU 's  internal policy. (R, W, set-by-create) (optional) (2 bytes) 

---- page break ---- Fill policy: This attribute defines the payload bit pattern to be applied toward s the TDM  service if no payload packet is available to play out. The default value 0  specifies that the ONU apply its internal policy.  0 ONU default, vendor -specific (recommended: AIS for  unstructured service, all 1s for structured service)  1 Play out AIS according to the service definition (for example,  DS3 AIS)  2 Play out all 1s  3 Play out all 0s  4 Repeat the previous data  5 Play out DS1 idle (Appendix C of [b-ATIS-0600403])  6..15 Reserved for future standardization  16..255 Vendor-specific, not to be standardized  (R, W, set-by-create) (optional) (1 byte)  Four pairs of alarm-related policy attributes, defined in the following, share common behaviour.  The alarm declaration policy attribute defines the anomaly rate that causes the corresponding alarm  to be declared. It is an integer percentage between 1..100. If this density of anomalies occurs during  the alarm onset soak interval, the alarm is declared.  The default value 0 selects the ONU 's internal  policy.  The alarm clear policy attribute defines the anomaly rate that causes the corresponding alarm to be  cleared. It is an integer percentage between 0..99. If no more than this density of anomalies occurs  during the alarm clear soak interval, the alarm is clea red. The default value 255 selects the ONU 's  internal policy.  Misconnected packets declaration policy: (R, W, set-by-create) (optional) (1 byte)  Misconnected packets clear policy: (R, W, set-by-create) (optional) (1 byte)  Loss of packets declaration policy: (R, W, set-by-create) (optional) (1 byte)  Loss of packets clear policy: (R, W, set-by-create) (optional) (1 byte)  Buffer overrun/underrun declaration policy: (R, W, set-by-create) (optional) (1 byte)  Buffer overrun/underrun clear policy: (R, W, set-by-create) (optional) (1 byte)  Malformed packets declaration policy: (R, W, set-by-create) (optional) (1 byte)  Malformed packets clear policy: (R, W, set-by-create) (optional) (1 byte)  R-bit transmit set policy: This attribute defines the number of consecutive lost packets that  causes the transmitted R bit to be set in the TDM to the PSN direction,  indicating lost packets to the far end. The default value 0 selects the ONU 's  internal policy. (R, W, set-by-create) (optional) (1 byte)  R-bit transmit clear policy : This attribute defines the number of consecutive valid packets  that causes the transmitted R bit to be cleared in the TDM to the PSN direction,  removing the remote failure indication to the far end. The default value 0  selects the ONU's internal policy. (R, W, set-by-create) (optional) (1 byte) 

---- page break ---- R-bit receive policy : This attribute defines the action toward s the N  × 64 TDM interface  when remote failure is indicated on packets received from the PSN (either  R-bit set or M = 0b10 while the L bit is cleared).  0 Do nothing (recommended to be the default)  1 Play out service-specific RAI/REI/RDI code  2 Send channel idle signalling and idle channel payload to all DS0s  comprising the service  (R, W, set-by-create) (optional) (1 byte)  L bit receive policy: This attribute defines the action towards the TDM interface when far-end  TDM failure is indicated on packets received from the PSN (L bit set).  0 Play out service-specific AIS (recommended to be the default)  1 Repeat last received packet  2 Send channel idle signalling and idle channel payload to all DS0s  comprising the service  (R, W, set-by-create) (optional) (1 byte)  SES threshold: Number of lost, malformed or otherwise unusable packets expected in the  PSN to the TDM direction within a 1 s interval that causes a n SES to be  counted. Stray packets do not count toward s an SES, nor do packets whose L  bit is set at the far end. The value 0 specifies that the ONU use s its internal  default, which is not necessarily the same as the recommended default value  3. (R, W, set-by-create) (optional) (2 bytes)  Actions  Create, delete, get, set  Notifications  None.  
```
