# Managed Entity

## Identity
- ME ID: 301
- ME Name: Dot1ag default MD level
- Source Section: 9.3.21
- Source Page: 183

## Classification
- Access: needs_review
- Type: needs_review
- Actions: needs_review
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: Layer 2 type
- Size: 1 byte
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 2
- Name: Catchall level
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 3
- Name: Catchall MHF creation
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 4
- Name: Catchall sender ID permission
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: mandatory

## Extraction Status
- Status: auto_extracted
- Attributes found: 4
- Review needed: true

## Raw Source

```
9.3.21 Dot1ag default MD level  The collection of the functionality called a maintenance half -function (MHF) is not explicitly  modelled as a ME by either [ IEEE 802.1ag] or the OMCI. The ONU automatically creates MHFs  according to parameters specified in a dot1ag MD or a dot1ag MA ME; the dot1ag default MD level  ME catches the corner cases not covered by other MEs, specifically VLANs not included by any  defined MA.  The dot1ag default MD level comprises a configurable table, each entry of which specifies default  MHF functionality for some set of VLANs. Once a set of VLANs is defined, operations to different  table entries or to dot1ag MAs that conflict with the set membership should be denied. In addition,  catch-all attributes are defined to specify MHF functionality when there is no match to either a table  entry or an MA.  Relationships  An ONU that supports [IEEE 802.1ag] automatically creates one instance of this ME for each  MAC bridge or IEEE 802.1p mapper, depending on the ONU's provisioning model. It should  not create an instance for an IEEE 802.1p mapper that is associated with a MAC bridge.  Attributes  Managed entity ID: This attribute uniquely identifies an instance of this ME. Through an  identical ID, this ME is implicitly linked to an instance of the MAC bridge  service profile ME or an IEEE 802.1p mapper ME. It is expected that an ONU  will implement CFM on bridges or on IEEE 802.1p mappers, but not both,  depending on its provisioning model. For precision, the reference is  disambiguated by the value of the layer 2 type pointer attribute. (R) (mandatory)  (2 bytes)  Layer 2 type: This attribute specifies whether the dot1ag default MD level ME is associated  with a MAC bridge service profile (value 0) or an IEEE 802.1p mapper  (value 1). (R) (mandatory) (1 byte)  Catchall level: This attribute ranges from 0..7 and specifies the MD level of MHFs created  when no specific match is found. (R, W) (mandatory) (1 byte)  Catchall MHF creation: This attribute determines whether, when no more specific match is  found, the bridge creates an MHF or not. This attribute is an enumeration with  the following values:  1 None. The bridge does not create any MHFs. This is the default value. 

---- page break ---- 2 Default. The bridge can create MHFs on this VID on any port through  which the VID can pass.  3 Explicit. The bridge can create MHFs on this VID on any port through  which the VID can pass, but only if a n MEP exists at some lower  maintenance level.  (R, W) (mandatory) (1 byte)  Catchall sender ID permission: This attribute determines the contents of the sender ID TLV  included in CFM messages transmitted by MPs when no more specific match  is found. This attribute is identical to that defined in the description of the  dot1ag MD ME (i.e., excluding code point 5, defer). (R,  W) (mandatory)  (1 byte)  Default MD level table: Each entry is a vector of fields, indexed by primary VLAN ID.  Primary VLAN ID (2 bytes)  Table control: This field controls the meaning of a set operation. The 1  byte  size of this field is included in get/get -next operations, but its value is  undefined under get-next and should be ignored by the OLT. (1 byte)  1 Add record to table; overwrite existing record, if any.  2 Delete record from table.  3 Clear all entries from table. This action may affect service and should  be used judiciously.  Other values are reserved.  Status: This Boolean field indicates whether this table entry is in effect (true)  or whether (false) it has been overridden by the existence of an MA for the  same VID and MD level as this table's entry, and on which an up MEP is  defined. This attribute is read-only. Space should be allocated for it during  set operations, but the value is not used. (1 byte)  Level: This field ranges from 0..7 and specifies the MD level of MHFs under  the control of this instance of the dot1ag default MD level. The additional  value 0xFF instructs the bridge to use the value in the catch -all level  attribute. (1 byte)  MHF creation: This attribute determines whether the bridge creates a n MHF  or not, under circumstances defined in clause 22.2.3 of [IEEE 802.1ag].  This attribute is an enumeration with the following values. (1 byte)  1 None. No MHFs are created on this bridge for this MA.  2 Default. The bridge can create MHFs on this VID on any port through  which the VID can pass.  3 Explicit. The bridge can create MHFs on this VID on any port through  which the VID can pass, but only if a n MEP exists at some lower  maintenance level.  4 Defer. This value causes the ONU to use the setting of the catch -all  MHF creation attribute. This is recommended to be the default value.  Sender ID permission: This attribute determines the contents of the sender ID  TLV included in CFM messages transmitted by MPs controlled by this MA.  (1 byte)  1 None: the sender ID TLV is not to be sent, default.  2 Chassis: the chassis ID length, chassis ID subtype, and chassis ID fields  of the sender ID TLV are to be sent, but not the management address  fields. 

---- page break ---- 3 Manage: the management address fields of the sender ID TLV are to  be sent, but the chassis ID length is to be transmitted with a 0 value,  and the chassis ID subtype, and chassis ID fields are not to be sent.  4 ChassisManage: all chassis ID and management address fields are to  be sent.  5 Defer: the contents of the sender ID TLV is determined by the catch - all sender ID permission attribute.  Associated VLANs list : This field comprises a list of up to 11 additional  VLAN IDs associated with the primary VLAN, 2  bytes each. Unused  placeholders, possibly including the entire field, are set to 0. (22 bytes)  (R, W) (mandatory) (29 bytes * N entries)  Actions  Get, get next, set  Set table (optional)  Notifications  None.  
```
