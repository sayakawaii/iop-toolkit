# Managed Entity

## Identity
- ME ID: 417
- ME Name: Data gathering line test, diagnostic and status
- Source Section: 9.7.40
- Source Page: 324

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
9.7.40 Data gathering line test, diagnostic and status  This ME contains xDSL data gathering line test, diagnostic and status parameters.  An instance of this ME is created and deleted by the OLT.  Relationships  An instance of this ME may be associated with zero or more instances of the PPTP xDSL UNI  part 1.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. Through an  identical ID, this ME is implicitly linked to an instance of the PPTP xDSL UNI  part 1 ME. (R, set-by-create) (mandatory) (2 bytes)  Logging depth – VTU-O (LOGGING_DEPTH_O): This parameter is the maximum depth  of the entire data gathering event buffer at the VTU -O, in number of 6  byte  data gathering records. See clause 7.5.3.1 of [ITU -T G.997.1] (R) (optional)  (4 bytes)  Logging depth – VTU-R (LOGGING_DEPTH_R): This parameter is the maximum depth  of the entire data gathering event buffer at the VTU -R, in number of 6  byte  data gathering records. See clause 7.5.3.2 of [ITU -T G.997.1] (R) (optional)  (4 bytes)  Actual logging depth for reporting – VTU-O (ACT_logging_depth_reporting_O): This  parameter is the actual logging depth that is used for reporting the VTU -O  event trace buffer in the CO-MIB, in number of 6 byte data gathering records.  See clause 7.5.3.3 of [ITU-T G.997.1] (R) (optional) (4 bytes) 

---- page break ---- Actual logging depth for reporting – VTU-R (ACT_logging_depth_reporting_R): This  parameter is actual logging depth that is used for reporting the VTU -R event  trace buffer over the eoc, in number of 6  byte data gathering records. See  clause 7.5.3.4 of [ITU-T G.997.1] (R) (optional) (4 bytes)  Event trace buffer – VTU-O (EVENT_TRACE_BUFFER_O) table: This parameter is the  event trace buffer containing the event records that originated at the VTU -O.  See clause 7.5.3.5 of [ITU-T G.997.1] (R) (optional) (N bytes)  Event trace buffer – VTU-R (EVENT_TRACE_BUFFER_R) table: This parameter is the  event trace buffer containing the event records that originated at the VTU -R.  See clause 7.5.3.6 of [ITU-T G.997.1] (R) (optional) (N bytes)  Actions  Get, get next  Notifications  None.  
```
