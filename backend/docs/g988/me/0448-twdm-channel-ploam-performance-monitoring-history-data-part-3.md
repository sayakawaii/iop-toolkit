# Managed Entity

## Identity
- ME ID: 448
- ME Name: TWDM channel PLOAM performance monitoring history data part 3
- Source Section: 9.16.5
- Source Page: 492

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
- Name: PLOAM MIC errors
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 4
- Name: Downstream PLOAM message count
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 5
- Name: Ranging_Time message count
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 6
- Name: Protection_Control message count
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 7
- Name: Adjust_Tx_Wavelength message count
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

## Extraction Status
- Status: auto_extracted
- Attributes found: 7
- Review needed: false

## Raw Source

```
9.16.5 TWDM channel PLOAM performance monitoring history data part 1  This ME collects certain PLOAM-related PM data associated with the slot/circuit pack, hosting on e  or more ANI-G MEs, for a specific TWDM channel. Instances of this ME are created and deleted by  the OLT.  The downstream PLOAM message counts of this ME include only the received PLOAM messages  pertaining to the given ONU, i.e.:  – unicast PLOAM messages, addressed by ONU-ID;  – broadcast PLOAM messages, addressed by serial number;  – broadcast PLOAM messages, addressed to all ONUs on the PON.  This ME includes all PLOAM PM counters characterized as mandatory in clause 14  of [ITU - T G.989.3].  For a complete discussion of generic PM architecture, refer to clause I.4.  Relationships  An instance of this ME is associated with an instance of TWDM channel ME.   Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. Through an  identical ID, this ME is implicitly linked to an instance of the TWDM channel  ME. (R, set-by-create) (mandatory) (2 bytes)  Interval end time : This attribute identifies the most recently finished 15  min interval. (R)  (mandatory) (1 byte)  Threshold data 1/2 ID: This attribute points to an instance of the threshold data 1 and 2 MEs  that contains PM threshold values. (R, W, set-by-create) (mandatory) (2 bytes)  PLOAM MIC errors : The counter of received PLOAM messages that remain unparsable  due to MIC error. (R) (mandatory) (4 bytes)  Downstream PLOAM message count : The counter of received broadcast and unicast  PLOAM messages pertaining to the given ONU. (R) (mandatory) (4 bytes)  Ranging_Time message count: The counter of received Ranging_Time PLOAM messages.  (R) (mandatory) (4 bytes)  Protection_Control message count : The counter of received Protection_Control PLOAM  messages. (R) (mandatory) (4 bytes)  Adjust_Tx_Wavelength message count : The counter of received Adjust_Tx_Wavelength  PLOAM messages. (R) (mandatory) (4 bytes) 

---- page break ---- Adjust_Tx_Wavelength adjustment amplitude: An estimator of the absolute value of the  transmission wavelength adjustment. (R) (mandatory) (4 bytes)  Actions  Create, delete, get, set  Get current data (optional)  Notifications  Threshold crossing alert  Alarm  number Threshold crossing alert Threshold value  attribute No. (Note)  0 PLOAM MIC errors 1  NOTE – This number associates the TCA with the specified threshold value attribute of the threshold data  1/2 managed entities.  
```
