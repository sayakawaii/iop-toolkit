# Managed Entity

## Identity
- ME ID: 305
- ME Name: Dot1ag CFM stack
- Source Section: 9.3.25
- Source Page: 191

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Get
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: Layer 2 type
- Size: 1 byte
- Format: needs_review
- Access: R
- Category: mandatory

## Extraction Status
- Status: auto_extracted
- Attributes found: 1
- Review needed: false

## Raw Source

```
9.3.25 Dot1ag CFM stack  This ME reports the maintenance status of a bridge port at any given time. An ONU that supports  [IEEE 802.1ag] functionality automatically creates an instance of the dot1ag CFM stack ME for each  MAC bridge or IEEE 802.1p mapper, depending on its provisioning model.  The dot1ag CFM stack also lists any VLANs and bridge ports against which configuration errors are  currently identified. The ONU should reject operations that create configuration errors. However,  these errors can arise because of operations on other MEs th at are not necessarily possible to detect  during CFM configuration. 

---- page break ---- Relationships  An ONU that supports [IEEE 802.1ag] creates one instance of this ME for each MAC bridge  or IEEE 802.1p mapper, depending on its provisioning model. It should not create an instance  for an IEEE 802.1p mapper that is associated with a MAC bridge.  Attributes  Managed entity ID: This attribute uniquely identifies an instance of this ME. Through an  identical ID, this ME is implicitly linked to an instance of the MAC bridge  service profile ME or an IEEE 802.1p mapper ME. It is expected that an ONU  will implement CFM on bridges or on IEEE 802.1p mappers, but not both. For  precision, the reference is disambiguated by the value of the layer 2 type  pointer attribute. (R) (mandatory) (2 bytes)  Layer 2 type: This attribute specifies whether the dot1ag CFM stack is associated with a  MAC bridge service profile (value 0) or an IEEE 802.1p mapper (value 1). (R)  (mandatory) (1 byte)  MP status table : This attribute is a list of entries, each entry reporting one aspect of the  maintenance status of one port. If a port is associated with more than one CFM  maintenance entity, each is represented as a separate item in this table attribute;  a port that has no current maintenance functions is not represented in the table  (so the table may be empty). Each entry is defined as follows.  Port ID: The ME ID of the MAC bridge port config data whose information  is reported in this entry. If the layer 2 parent is an IEEE 802.1p mapper, a  null pointer. (2 bytes)  Level: The level at which the reported maintenance function exists, 0..7.  (1 byte)  Direction: The value 1 (down) or 2 (up). (1 byte)  VLAN ID: If this table entry reports a maintenance function associated with a  VLAN, this field contains the value of the primary VLAN ID. If no VLAN  is associated with this entry, this field contains the value 0. (2 bytes)  MD: A pointer to the associated dot1ag maintenance domain ME. If no MD is  associated with this entry, a null pointer. (2 bytes)  MA: A pointer to the associated dot1ag maintenance association ME. If no  MA is associated with this entry, a null pointer. (2 bytes)  MEP ID: If this table entry reports an MEP, this field contains the value of its  MEP ID (range 1..8191). If this table entry reports an MHF, this field  contains the value 0. (2 bytes)  MAC address: The MAC address of the MP. (6 bytes)  (R) (mandatory) (18N bytes)  Configuration error list table: This attribute is based on the [ IEEE 802.1ag] configuration  error list. It is a list of entries, each entry reporting a VLAN and a bridge port  against which a configuration error has been detected. The table may be empty  at any given time. Entries are defined as follows:  VLAN ID: If this table entry reports a maintenance function associated with a  VLAN, this field contains the value of the VLAN ID in error. If no VLAN  is associated with this entry, this field contains the value 0. (2 bytes) 

---- page break ---- Port ID: A pointer to the MAC bridge port config data whose information is  reported in this entry. If the layer 2 parent is an IEEE 802.1p mapper, a null  pointer. (2 bytes)  Detected configuration error: A bit mask with the following meanings. A list  entry exists if and only if at least one of these bits is set. Definitions appear  in clause 22.2.4 of [IEEE 802.1ag]: (1 byte)  0x01 CFM leak. MA x is associated with a specific VID list, one or more  of the VIDs in MA x can pass through the bridge port, no up MEP  is configured for MA x on the bridge port, no down MEP is  configured on any bridge port for MA x, and another MA y, at a  higher MD level than MA x, and associated with at least one of the  VID(s) also in MA x, does have a n MEP configured on the bridge  port.  0x02 Conflicting VIDs. MA x is associated with a specific VID list, an  up MEP is configured on MA x on the bridge port, and another MA  y, associated with at least one of the VID(s) also in MA x, and at  the same MD level as MA x, also has an up MEP configured on  some bridge port.  0x04 Excessive levels. The number of different MD levels at which  maintenance domain intermediate points ( MIPs) are to be created  on this port exceeds the bridge's capabilities.  0x08 Overlapped levels. An MEP is created for one VID at one MD level,  but an MEP is also configured on another VID at that MD level or  higher, exceeding the bridge's capabilities.  (R) (mandatory) (5N bytes)  Actions  Get, get next  Notifications  Attribute value change  Number Attribute value  change  Description  1..2 N/A   3 Config error list table This AVC indicates that an entry in the configuration error  list table has been added or removed. It may be advisable for  the OLT to audit the configuration of related MEs.  4..16 Reserved   
```
