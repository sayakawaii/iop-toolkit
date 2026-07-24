# Managed Entity

## Identity
- ME ID: 416
- ME Name: Vectoring line inventory and status data
- Source Section: 9.7.39
- Source Page: 323

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
9.7.39 Vectoring line inventory and status data  This ME contains line inventory and status attributes specific to vectoring.  Relationships  This is one of the status data MEs associated with an xDSL UNI. It is meaningful if the PPTP  supports [ITU-T G.993.5]. The ONU automatically creates or deletes an instance of this ME  upon creation and deletion of a PPTP xDSL UNI part 1 that supports these attributes.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. Through an  identical ID, this ME is implicitly linked to an instance of the PPTP xDSL UNI  part 1 ME. (R) (mandatory) (2 bytes)  Downstream XLIN scale (XLINSCds) : This parameter is the scale factor to be applied to  the downstream Xlinpsds values. Valid values are given in clause 7.5.1.39.1  of [ITU-T G.997.1] (R) (mandatory) (2 bytes)  Downstream XLIN subcarrier group size (XLINGds) : This parameter is the number of  subcarriers per group used to report Xlinpsds. Valid values are given in  clause 7.5.1.39.2 of [ITU-T G.997.1] (R) (mandatory) (1 bytes)  Downstream XLIN bandedges (XLINBANDSds) table: XLINBANDSds contains pairs of  indices (start_subcarrier_index, stop_subcarrier_index) for every band in  which XLINpsds is reported. Each index is 2 bytes. This attribute is organized  as a table, so the number of bands can be determined from the table's size. See  clause 7.5.1.39.3 of [ITU-T G.997.1] (R) (mandatory) (N bands x 4 bytes)  Downstream FEXT coupling (XLINpsds) table : For each given VCE port index k, this  parameter is a one -dimensional array of complex values in linear scale for  downstream FEXT coupling coefficients Xlinds(f) originating from the loop  connected to the VCE port k into the loop for which Xlinds(f) is being reported.  Each complex value [a(n) + j  b(n)] is represented by a 2 byte signed two's  complement value [a(n)], followed by a 2 byte signed two's complement value  [b(n)]. This attribute is organized as a table, so the number of complex values  in the array can be determined from the table's size. See clause 7.5.1.39.4 of  [ITU-T G.997.1] (R) (mandatory) (N complex values  4 bytes)  Upstream XLIN scale (XLINSCus) : This parameter is the scale factor to be applied to the  upstream X LINpsus values. Valid values are given in clause 7.5.1.39.5 of  [ITU-T G.997.1] (R) (mandatory) (2 bytes)  Upstream XLIN subcarrier group size (XLINGus) : This parameter is the number of  subcarriers per group used to report X LINpsus. Valid values are given in  clause 7.5.1.39.6 of [ITU-T G.997.1] (R) (mandatory) (1 bytes) 

---- page break ---- Upstream XLIN bandedges (XLINBANDSus) table : XLINBANDSus contains pairs of  indices (start_subcarrier_index, stop_subcarrier_index) for every band in  which XLINpsus is reported. Each index is 2 bytes. This attribute is organized  as a table, so the number of bands can be determined from the table's size. See  clause 7.5.1.39.7 of [ITU-T G.997.1] (R) (mandatory) (N bands x 4 bytes)  Upstream FEXT coupling (XLINpsus) table : For each given VCE port index k, this  parameter is a one -dimensional array of complex values in linear scale for  upstream FEXT coupling coefficients Xlinus(f) originating from the loop  connected to the VCE port k into the loop for which Xlinus(f) is being reported.  Each complex value [a(n) + j  b(n)] is represented by a 2 byte signed two's  complement value [a(n)], followed by a 2 byte signed two's complement value  [b(n)]. This attribute is organized as a table, so the number of complex values  in the array can be determined from the table's size. See clause 7.5.1.39.8 of  [ITU-T G.997.1] (R) (mandatory) (N complex values  4 bytes)  Actual vectoring mode (ACTVECTORMODE) : This parameter reports the vectoring  initialization type of the line. Valid values are given in clause 7.5.1.43.1 of  [ITU-T G.997.1] (R) (optional) (1 byte)  Actions  Get, get next  Notifications  None.  
```
