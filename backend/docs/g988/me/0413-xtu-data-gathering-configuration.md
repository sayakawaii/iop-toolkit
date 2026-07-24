# Managed Entity

## Identity
- ME ID: 413
- ME Name: xTU data gathering configuration
- Source Section: 9.7.36
- Source Page: 318

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Get
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

<!-- No attributes successfully parsed; see Raw Source section -->

## Extraction Status
- Status: auto_extracted
- Attributes found: 0
- Review needed: true

## Raw Source

```
9.7.36 xTU data gathering configuration  This ME defines configurations specific to data gathering.  An instance of this ME is created and deleted by the OLT.  Relationships  An instance of this ME may be associated with zero or more instances of the PPTP xDSL UNI  part 1. 

---- page break ---- Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. Through an  identical ID, this ME is implicitly linked to an instance of the PPTP xDSL UNI  part 1 ME. (R, set-by-create) (mandatory) (2 bytes)  Logging depth event percentage per event – VTU-O  (LOGGING_DEPTH_EVENT_PERCENTAGE_Oi) table : This  parameter is the percentage of the data gathering event buffer assigned to event  type i at the VTU-O. See clause 7.3.6.1 of [ITU -T G.997.1]. Each element in  the table consists of 2 bytes, where the first byte is event type i, and the second  byte is the percentage of event type i defined as the integer value multiplied by  1%. (R, W) (optional) (2  N bytes for N event types)  Logging depth event percentage per event – VTU-R  (LOGGING_DEPTH_EVENT_PERCENTAGE_Ri) table: This parameter  is the percentage of the data gathering event buffer assigned to event type i at  the VTU-R. See clause 7.3.6.2 of [ITU-T G.997.1]. Each element in the table  consists of 2 bytes, where the first byte is event type i, and the second byte is  the percentage of event type i defined as the integer value multiplied by 1%.  (R, W) (optional) (2  N bytes for N event types)  Logging depth for VTU-O reporting – VTU-R (LOGGING_DEPTH_REPORTING_O):  This parameter is the logging depth that is requested for reporting the VTU-O  event trace buffer in the CO-MIB, in number of 6 byte data gathering records.  See clause 7.3.6.3 of [ITU-T G.997.1]. (R, W) (optional) (2 bytes)  Logging depth for VTU-R reporting – VTU-R (LOGGING_DEPTH_REPORTING_R):  This parameter is the logging depth that is requested for reporting the VTU-R  event trace buffer over the embedded operations channel (eoc), in number of  6 byte data gathering records. See clause 7.3.6.4 of [ITU -T G.997.1]. (R, W)  (optional) (2 bytes)  Logging data report newer events first – VTU-R  (LOGGING_REPORT_NEWER_FIRST): This parameter determines  whether the VTU -R to reports newer events first or older events first. See  clause 7.3.6.4 of [ITU-T G.997.1]. False is mapped to 0, true is mapped to 1.  (R, W) (optional) (1 byte)  Actions  Create, delete, get, get next, set  Set table (optional)  Notifications  None.  9.7.37 xDSL line inventory and status data part 8  This ME extends the attributes defined in the xDSL line inventory and status data parts 1..4.   Relationships  This is one of the status data MEs associated with an xDSL UNI. The ONU automatically  creates or deletes an instance of this ME upon creation or deletion of a PPTP xDSL UNI part  1 that supports these attributes. 

---- page break ---- Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. Through an  identical ID, this ME is implicitly linked to an instance of the PPTP xDSL UNI  part 1 ME. (R) (mandatory) (2 bytes)  Retransmission used downstream (RTX_USEDds) : This parameter specifies whether  [ITU-T G.998.4] retransmission is used (i.e., active in showtime) in the  downstream transmit direction. The valid range of values is given in  clause 7.5.1.38 of [ITU-T G.997.1]. (R) (mandatory) (1 byte)  Retransmission used upstream (RTX_USEDus) : This parameter specifies whether   [ITU-T G.998.4] retransmission is used (i.e., active in showtime) in the  upstream transmit direction. The valid range of values is given in clause  7.5.1.38 of [ITU-T G.997.1]. (R) (mandatory) (1 byte)  Date/time-stamping of near -end test parameters (STAMP -TEST-NE): This parameter  indicates the date/time when the near -end test parameters that can change  during showtime were last updated. See clause 7.5.1.36.3 of [ITU-T G.997.1].  The format of this parameter is as follows.    Year 2 bytes  Month 1 byte (1..12)  Day 1 byte (1..31)  Hour 1 byte (0..23)  Minute 1 byte (0..59)  Second 1 byte (0..59)  (R) (optional) (7 bytes)  Date/time-stamping of far -end test parameters (STAMP -TEST-FE): This parameter  indicates the date/time when the far-end test parameters that can change during  showtime were last updated. See clause 7.5.1.36.4 of [ITU -T G.997.1]. The  format of this parameter is the same as STAMP -TEST-NE. (R) (optional)  (7 bytes)  Date/time-stamping of last successful downstream OLR operation (STAMP -OLR-ds):  This parameter indicates the date/time of the last successful OLR execution in  the downstream direction that has modified the bits or gains. See  clause 7.5.1.37.1 of [ITU-T G.997.1]. The format of this parameter is the same  as STAMP-TEST-NE. (R) (optional) (7 bytes)  Date/time-stamping of last successful upstream OLR operation (STAMP-OLR-us): This  parameter indicates the date/time of the last successful OLR execution in the  upstream direction that has modified the bits or gains. See clause 7.5.1.37.2 of  [ITU-T G.997.1]. The format of this parameter is the same as STAMP-TEST- NE. (R) (optional) (7 bytes)  Actions  Get, get next  Notifications  None. 

---- page break ---- 
```
