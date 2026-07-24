# Managed Entity

## Identity
- ME ID: 450
- ME Name: TWDM channel tuning performance monitoring history data part 2
- Source Section: 9.16.8
- Source Page: 496

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
- Name: Tuning control requests for Rx only or Rx and Tx
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 4
- Name: Tuning control requests for Tx only
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 5
- Name: Tuning control requests rejected/INT_SFC
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 6
- Name: Tuning control requests rejected/DS_xxx
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 7
- Name: Tuning control requests fulfilled with ONU reacquired at target channel
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 8
- Name: Tuning control requests failed due to target DS wavelength channel not found
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 9
- Name: Tuning control requests resolved with ONU reacquired at discretionary channel
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 10
- Name: Tuning control requests Rollback/COM_DS
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 11
- Name: Tuning control requests Rollback/DS_xxx
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 12
- Name: Tuning control requests Rollback/US_xxx
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 13
- Name: Tuning control requests failed with ONU reactivation
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

## Extraction Status
- Status: auto_extracted
- Attributes found: 13
- Review needed: false

## Raw Source

```
9.16.8 TWDM channel tuning performance monitoring history data part 1  This ME collects certain tuning-control-related PM data associated with the slot/circuit pack, hosting  one or more ANI-G MEs, for a specific TWDM channel. Instances of this ME are created and deleted  by the OLT.  The relevant events this ME is concerned with are counted towards the PM statistics associated with  the source TWDM channel. The attribute descriptions refer to the ONU activation cycle states and  timers specified in clause 12 of [ITU-T G.989.3]. This ME contains the counters characterized as  mandatory in clause 14 of [ITU-T G.989.3].  For a complete discussion of generic PM architecture, refer to clause I.4.  Relationships  An instance of this ME is associated with an instance of TWDM channel ME.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. Through an  identical ID, this ME is implicitly linked to an instance of the TWDM channel  ME. (R, set-by-create) (mandatory) (2 bytes)  Interval end time : This attribute identifies the most recently finished 15  min interval. (R)  (mandatory) (1 byte)  Threshold data 1/2 ID: This attribute points to an instance of the threshold data 1 and 2 MEs  that contains PM threshold values. (R, W, set-by-create) (mandatory) (2 bytes)  Tuning control requests for Rx only or Rx and Tx: The counter of received Tuning_Control  PLOAM messages with Request operation code that contain tuning  instructions either for receiver only or for both receive r and transmitter. (R)  (mandatory) (4 bytes)  Tuning control requests for Tx only: The counter of received Tuning_Control PLOAM  messages with Request operation code that contain tuning instructions for  transmitter only. (R) (mandatory) (4 bytes)  Tuning control requests rejected/INT_SFC: The counter of transmitted Tuning_Response  PLOAM messages with NACK operation code and INT_SFC response code,  indicating inability to start transceiver tuning by the specified time (SFC). (R)  (mandatory) (4 bytes)  Tuning control requests rejected/DS_xxx : The aggregate counter of transmitted  Tuning_Response PLOAM messages with NACK operation code and any  DS_xxx response code, indicating target downstream wavelength channel  inconsistency. (R) (mandatory) (4 bytes) 

---- page break ---- Tuning control requests rejected/US_xxx : The aggregate counter of transmitted  Tuning_Response PLOAM messages with NACK operation code and any  US_xxx response code, indicating target upstream wavelength channel  inconsistency. (R) (mandatory) (4 bytes)  Tuning control requests fulfilled with ONU reacquired at target channel: The counter of  controlled tuning attempts for which an upstream tuning confirmation has been  obtained in the target channel. (R) (mandatory) (4 bytes)  Tuning control requests failed due to target DS wavelength channel not found : The  counter of controlled tuning attempts that failed due to timer TO4 expiration  in the DS Tuning state (O8) in the target channel. (R) (mandatory) (4 bytes)  Tuning control requests failed due to no feedback in target DS wavelength channel: The  counter of controlled tuning attempts that failed due to timer TO5 expiration  in the US Tuning state (O9) in the target channel. (R) (mandatory) (4 bytes)  Tuning control requests resolved with ONU reacquired at discretionary channel : The  counter of controlled tuning attempts for which an upstream tuning  confirmation has been obtained in the discretionary channel. (R) (mandatory)  (4 bytes)  Tuning control requests Rollback/COM_DS: The counter of controlled tuning attempts that  failed due to communication condition in the target channel, as indicated by  the Tuning_Response PLOAM message with Rollback operation code and  COM_DS response code. (R) (mandatory) (4 bytes)  Tuning control requests Rollback/DS_xxx : The aggregate counter of controlled tuning  attempts that failed due to target downstream wavelength channel  inconsistency, as indicated by the Tuning_Response PLOAM message with  Rollback operation code and any DS_xxx response code. (R) (mandatory)  (4 bytes)  Tuning control requests Rollback/US_xxx : The aggregate counter of controlled tuning  attempts that failed due to target upstream wavelength channel parameter  violation, as indicated by the Tuning_Response PLOAM message with  Rollback operation code and US_xxx response code. (R) (mandatory) (4 bytes)  Tuning control requests failed with ONU reactivation : The counter of controlled tuning  attempts that failed on any reason, with expiration of timers TO4 or TO5  causing the ONU transition into state O1. (R) (mandatory) (4 bytes)  Actions  Create, delete, get, set  Get current data (optional)  Notifications  Threshold crossing alert  Alarm  number Threshold crossing alert Threshold value  attribute No. (Note)  0 Tuning control requests rejected/INT_SFC 1  1 Tuning control requests rejected/DS_xxx 2  2 Tuning control requests rejected/US_xxx 3  3 Tuning control requests failed/TO4 exp. 4 

---- page break ---- Threshold crossing alert  Alarm  number Threshold crossing alert Threshold value  attribute No. (Note)  4 Tuning control requests failed/TO5 exp. 5  5 Tuning control requests Rollback/COM_DS 6  6 Tuning control requests Rollback/DS_xxx 7  7 Tuning control requests Rollback/US_xxx 8  8 Tuning control requests failed/Reactivation 9  NOTE – This number associates the TCA with the specified threshold value attribute of the threshold data  1/2 managed entities.  
```
