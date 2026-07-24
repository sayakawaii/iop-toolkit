# Managed Entity

## Identity
- ME ID: 310
- ME Name: Multicast subscriber config info
- Source Section: 9.3.28
- Source Page: 202

## Classification
- Access: needs_review
- Type: needs_review
- Actions: needs_review
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: ME type
- Size: 1 byte
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 2
- Name: Max simultaneous groups
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: optional

### Attribute 3
- Name: Max multicast bandwidth
- Size: 4 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: optional

### Attribute 4
- Name: Bandwidth enforcement
- Size: 1 byte
- Format: needs_review
- Access: R, W, set-by-create
- Category: optional

## Extraction Status
- Status: auto_extracted
- Attributes found: 4
- Review needed: true

## Raw Source

```
9.3.28 Multicast subscriber config info  This ME organizes data associated with multicast management at subscriber ports of IEEE  802.1  bridges, including IEEE 802.1p mappers when the provisioning model is mapper -based rather than  bridge-based. Instances of this ME are created and deleted by the OLT. Because of backward  compatibility considerations, a subscriber port without an associated multicast subscriber config info  ME would be expected to support unrestricted multicast access; this ME may therefore be viewed as  restrictive, rather than permissive.  Through separate attributes, this ME supports either a single multicast operations profile in its  backward compatible form, or a list of multicast operations profiles instead (the list may of course  contain a single entry). The OLT can determine whether th e ONU supports the multiple profile  capability by performing a get operation on the optional multicast service package table attribute,  which exists only on ONUs that are prepared to support the feature.  Relationships  An instance of this ME is associated with one instance of the MAC bridge port configuration  data or the IEEE 802.1p mapper service profile.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. Through an  identical ID, this ME is implicitly linked to an instance of the MAC bridge port  configuration data or IEEE 802.1p mapper ME. (R, set-by-create) (mandatory)  (2 bytes)  ME type: This attribute indicates the type of the ME implicitly linked by the ME ID  attribute.  0 MAC bridge port config data  1 IEEE 802.1p mapper service profile  (R, W, set-by-create) (mandatory) (1 byte) 

---- page break ---- Multicast operations profile pointer : This attribute points to an instance of the multicast  operations profile. This attribute is ignored by the ONU if a non -empty  multicast service package table attribute is present. (R,W, set-by-create)  (mandatory) (2 bytes)  Max simultaneous groups : This attribute specifies the maximum number of dynamic  multicast groups that may be replicated to the client port at any one time. The  recommended default value 0 specifies that no administrative limit is to be  imposed. (R, W, set-by-create) (optional) (2 bytes)  Max multicast bandwidth : This attribute specifies the maximum imputed dynamic  bandwidth, in bytes per second, that may be delivered to the client port at any  one time. The recommended default value 0 specifies that no administrative  limit is to be imposed. (R, W, set-by-create) (optional) (4 bytes)  Bandwidth enforcement: The recommended default value of this Boolean attribute is false,  and specifies that attempts to exceed the max multicast bandwidth be counted  but honoured. The value true specifies that such attempts be counted and  denied. The imputed bandwidth value is taken from the dynamic access control  list table, both for a new join request and for pre -existing groups. (R,  W,  set-by-create) (optional) (1 byte)  Multicast service package table: This attribute is a list that specifies one or more multicast  service packages. When the ONU receives an IGMP/MLD join request, it  searches the multicast service package table in row key order, matching the  VID (UNI) field (several rows can share the same VID). For each VID (UNI)  match, the multicast operations profile pointer is used to access the ME that  contains the attributes associated with the service package. The search stops  when all requested multicast groups have been found and dealt with.  Each list entry is a vector of six components as follow.  – Table control (2 bytes)  The first 2 bytes of each entry contain a key into the table. It is the  responsibility of the OLT to assign and track table keys and content.  Since row keys are created by the OLT, they may be densely or sparsely  packed.  The two MSBs of this field determine the meaning of a set operation.  These bits are returned as 00 during get next operations.      Bits 16..15 Meaning  00 Reserved  01 Write this entry into the table. Overwrite any  existing entry with the same row key.  10 Delete this entry from the table. The remaining  fields are not meaningful.  11 Clear all entries from the table. The remaining fields  are not meaningful.  Bits 14..11 are reserved. Bits 10..1 are the row key itself.  – VID (UNI) . The value in this field is compared with the VID of  upstream IGMP/MLD messages, and is used to decide whether to  honour a join request. (2 bytes) 

---- page break ---- Values:  0..4095 – Matched against the VID of the IGMP/MLD message. 0  indicates a priority-tagged message, whose P bits are ignored.  4096 – Matches untagged IGMP/MLD messages only.  4097 – Matches tagged messages only, but ignores the value of the  VID.  0xFFFF – Unspecified.  The VID (UNI) comparison occurs prior to any action defined by the  upstream IGMP tag control attribute in an associated multicast  operations profile (or alternatively, before any modification by a  possible (extended) VLAN tagging operation configuration data ME).  – Max simultaneous groups. This field specifies the maximum number  of dynamic multicast groups that may be replicated to the client port at  any one time, for the multicast service package that is associated with  this row. The value 0 specifies that no administrative limit is to be   imposed. (2 bytes)  – Max multicast bandwidth. This field specifies the maximum imputed  dynamic bandwidth, in bytes per second, that may be delivered to the  client port at any one time, for the multicast service package that is  associated with this row. The value 0 specifies that no administrative  limit is to be imposed. (4 bytes)  NOTE – The port is also constrained by the global max simultaneous groups  and max multicast bandwidth attributes of the multicast subscriber config info  ME.  – Multicast operations profile pointer . This field contains the ME ID  of the multicast operations profile ME associated with this service  package. (2
```
