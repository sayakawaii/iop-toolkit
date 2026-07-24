# Managed Entity

## Identity
- ME ID: 111
- ME Name: xDSL downstream RFI bands profile
- Source Section: 9.7.11
- Source Page: 272

## Classification
- Access: needs_review
- Type: needs_review
- Actions: needs_review
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: Downstream RFI bands table
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: mandatory

## Extraction Status
- Status: auto_extracted
- Attributes found: 1
- Review needed: true

## Raw Source

```
9.7.11 xDSL downstream RFI bands profile  This ME contains the downstream RFI bands profile for an xDSL UNI. Instances of this ME are  created and deleted by the OLT.  Relationships  An instance of this ME may be associated with zero or more instances of the PPTP xDSL UNI  part 1.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. The value 0  is reserved. (R, set-by-create) (mandatory) (2 bytes)  Downstream RFI bands table : The RFIBANDS attribute is a table where each entry  comprises:  • an entry number field (1 byte, first entry numbered 1);  • subcarrier index 1 field (2 bytes);  • subcarrier index 2 field (2 bytes).  For [ITU -T G.992.5], this configuration attribute defines the subset of  downstream PSD mask breakpoints, as specified in the downstream PSD mask,  to be used to notch an RFI band. This subset consists of couples of consecutive  subcarrier indices belonging to breakpoints: [ti; ti + 1], corresponding to the low  level of the notch. Interpolation around these points is defined in [ITU -T  G.992.5]. 

---- page break ---- For [ITU-T G.993.2], this attribute defines the bands where the PSD is to be  reduced as specified in clause 7.2.1.2 of [ITU -T G.993.2]. Each band is  represented by start and stop subcarrier indices with a subcarrier spacing of  4.3125 kHz. Up to 16 bands may be specified. This attribute defines the RFI  bands for both upstream and downstream directions.  Entries have the default value 0 for both subcarrier index 1 and subcarrier  index 2. Setting an entry with a non -zero subcarrier index 1 and subcarrier  index 2 implies insertion into the table or replacement of an existing entry.   Setting an entry 's subcarrier index 1 and subcarrier index 2 to 0 implies  deletion from the table, if present.  (R, W) (mandatory for [ITU-T G.992.5], [ITU-T G.993.2]) (5 * N bytes where  N is the number of RFI bands)  Bands valid: This Boolean attribute controls and reports the operational status of the  downstream RFI bands table.  If this attribute is true, the downstream RFI bands table has been impressed on  the DSL equipment.  If this attribute is false, the downstream RFI bands table has not been  impressed on the DSL equipment. The default value is false.  This attribute can be modified by the ONU and OLT, as follows.  • If the OLT changes any of the RFI bands table entries or sets bands  valid false, then bands valid is false.  • If bands valid is false and OLT sets bands valid true, the ONU  impresses the downstream RFI bands data on to the DSL equipment.  (R, W) (mandatory) (1 byte)  Actions  Create, delete, get, get next, set  Set table (optional)  Notifications  None.  9.7.12 xDSL line inventory and status data part 1  This ME contains part 1 of the line inventory and status data for an xDSL UNI. The ONU  automatically creates or deletes an instance of this ME upon the creation or deletion of a PPTP xDSL  UNI part 1.  Relationships  An instance of this ME is associated with an xDSL UNI.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. Through an  identical ID, this ME is implicitly linked to an instance of the PPTP xDSL UNI  part 1. (R) (mandatory) (2 bytes)  xTU-C G.994.1 vendor ID : This is the vendor ID as inserted by the xTU -C in the  ITU-T G.994.1 CL message. It comprises 8 octets, including a country code  followed by a (regionally allocated) provider code, as defined in [ITU-T T.35].  (R) (mandatory) (8 bytes)
```
