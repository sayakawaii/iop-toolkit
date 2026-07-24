# Managed Entity

## Identity
- ME ID: 463
- ME Name: FAST line failures performance monitoring data
- Source Section: 9.7.60
- Source Page: 354

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

## Extraction Status
- Status: auto_extracted
- Attributes found: 2
- Review needed: false

## Raw Source

```
9.7.60 FAST line failures performance monitoring data  This ME collects data on the line failures of the FTU-O and the FTU-R for the previous or the current  15 minutes interval (depending on whether "get" or "get current data" command is used, as defined  in clause I.4). Instances of this ME are created and deleted by the OLT.   Relationships  An instance of this ME is associated with an xDSL UNI part 1. It is required only if FAST is  supported by the PPTP. The ONU automatically creates or deletes an instance of this ME  upon creation or deletion of a PPTP xDSL UNI part 1 that supports these attributes.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. Through an  identical ID, this ME is implicitly linked to an instance of a PPTP xDSL UNI  part 1. (R, set-by-create) (mandatory) (2 bytes)  Interval end time : This attribute identifies the most recently finished 15  min interval. (R)  (mandatory) (1 byte)  Threshold data 1/2 ID: This attribute points to an instance of the threshold data 1 and 2 MEs  that contain PM threshold values. (R, W, set-by-create) (mandatory) (2 bytes)  Near-end LOS failure counter  (CURR/PREV_NE_15_LOS): This attribute reports the  count of LOS failures at the XTU -C, as defined in clause 7.4.1.1 of [ITU-T  G.997.2]. See clause 7.7. 4 of [ITU-T G.997.2] for detailed specification. (R)  (mandatory) (4 bytes).  Near-end LOR failure counter (CURR/PREV_NE_15_LOR): This attribute reports the  count of LOR failures at the XTU -C, as defined in clause 7.4.1.2 of  [ITU-T G.997.2]. See clause 7.7. 5 of [ITU -T G.997.2] for detailed  specification. (R) (mandatory) (4 bytes).  Near-end LOM failure counter (CURR/PREV_NE_15_LOM): This attribute reports the  count of LOM failures at the XTU -C, as defined in clause 7.4.1.3 of  [ITU-T G.997.2]. See clause 7.7. 6 of [ITU -T G.997.2] for detailed  specification. (R) (mandatory) (4 bytes).  Near-end LPR failure count er ( CURR/PREV_NE_15_LPR): This attribute reports the  count of LPR failures at the XTU -C, as defined in clause 7.4.1.4 of  [ITU-T G.997.2]. See clause 7.7. 7 of [ITU -T G.997.2] for detailed  specification. (R) (mandatory) (4 bytes). 

---- page break ---- Far-end LOS failure counter  (CURR/PREV_FE_15_LOS): This attribute reports the  count of LOS failures at the XTU -C, as defined in clause 7.4.2.1 of  [ITU-T G.997.2]. See clause 7.7. 4 of [ITU -T G.997.2] for detailed  specification. (R) (mandatory) (4 bytes).  Far-end LO R failure counter  (CURR/PREV_FE_15_LOR): This attribute reports the  count of LOR failures at the XTU -C, as defined in clause 7.4.2.2 of  [ITU-T G.997.2]. See clause 7.7. 5 of [ITU -T G.997.2] for detailed  specification. (R) (mandatory) (4 bytes).  Far-end LO M failure counter  (CURR/PREV_FE_15_LOM): This attribute reports the  count of LOM failures at the XTU -C, as defined in clause 7.4.2.3 of  [ITU-T G.997.2]. See clause 7.7. 6 of [ITU -T G.997.2] for detailed  specification. (R) (mandatory) (4 bytes).  Far-end LPR failure counter  (CURR/PREV_FE_15_LPR): This attribute reports the  count of LPR failures at the XTU -C, as defined in clause 7.4.2.4 of  [ITU-T G.997.2]. See clause 7.7. 7 of [ITU-T G.997.2] for detailed  specification. (R) (mandatory) (4 bytes).  Actions  Create, delete, get, set  Get current data (optional)  Notifications  None  9.8 Time division multiplex services  This clause defines MEs associated with CES UNIs, as shown in Figure 9.8-1.
```
