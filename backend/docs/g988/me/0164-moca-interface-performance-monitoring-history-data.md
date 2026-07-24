# Managed Entity

## Identity
- ME ID: 164
- ME Name: MoCA interface performance monitoring history data
- Source Section: 9.10.3
- Source Page: 422

## Classification
- Access: needs_review
- Type: needs_review
- Actions: needs_review
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: Threshold data 1/2 ID
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 2
- Name: PHY Tx broadcast rate
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 3
- Name: Node table
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 4
- Name: Deprecated
- Size: 2 bytes
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 5
- Name: Administrative state
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 6
- Name: It is  recommended that this attribute not be used
- Size: 1 byte
- Format: needs_review
- Access: R
- Category: optional

## Extraction Status
- Status: auto_extracted
- Attributes found: 6
- Review needed: true

## Raw Source

```
9.10.3 MoCA interface performance monitoring history data  This ME collects PM data for an MoCA interface. Instances of this ME are created and deleted by  the OLT.  NOTE – The structure of this ME is an exception to the normal definition of PM MEs and normal PM  behaviour (clause I.4). It should not be used as a guide for the definition of future MEs. Among other  exceptions, this ME contains only current values, which are retrievable by get and get next operations; no  history is retained.  Relationships  An instance of this ME is associated with an instance of the PPTP MoCA UNI ME.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. Through an  identical ID, this ME is implicitly linked to an instance of the PPTP MoCA  UNI. (R, set-by-create) (mandatory) (2 bytes) 

---- page break ---- Interval end time : This attribute identifies the most recently finished 15  min interval. (R)  (mandatory) (1 byte)  Threshold data 1/2 ID: This attribute points to an instance of the threshold data  1 ME that  contains PM threshold values. Since no threshold value attribute number  exceeds 7, a threshold data 2 ME is optional. (R, W, set-by-create) (mandatory)  (2 bytes)  PHY Tx broadcast rate : This attribute indicates the MoCA PHY broadcast transmit rate  from the ONU MoCA interface to all the nodes in bit s per second. (R)  (optional) (4 bytes)  Node table: This attribute lists current nodes in the node table. The table contains MAC  addresses and statistics for those nodes. These table attributes are further  described in the following . Space for non -supported optional fields must be  allocated in table records, and filled with zero bytes.  MAC address: A unique identifier of a node within the table. (6 bytes)  PHY Tx rate: MoCA PHY unicast transmit rate from the ONU MoCA  interface to the node identified by the MAC address, in bits per second.  (4 bytes)  Tx power control reduction : The reduction in transmitter level due to  power control, in d ecibels. Valid values range from 0 (full power) to  60. (1 byte)  PHY Rx rate: MoCA PHY unicast receive rate to the ONU MoCA  interface from the node identified by the MAC address, in bit s per  second. (optional) (4 bytes)  Rx power level : The power level received at the ONU MoCA interface  from the node identified by the MAC address, in d ecibel-milliwatts,  represented as a 2s complement integer. Valid values range from +10  (0x0A) to –80 (0xB0). (1 byte)  PHY Rx broadcast rate: MoCA PHY broadcast receive rate to the ONU  MoCA interface from the node identified by MAC address, in bits  per  second. (optional) (4 bytes)  Rx broadcast power level: The power level received at the ONU MoCA  interface from the node identified by the MAC address, in d ecibel- milliwatts, represented as a 2s complement integer. Valid values range  from +10 (0x0A) to –80 (0xB0). (1 byte)  Tx packet: Number of packets transmitted to the node. (4 bytes)  Rx packet: Number of packets received from the node. (4 bytes)  Rx errored and missed: Number of errored and missed packets received  from the node. The sum of this field across all entries in the node table  contributes to the Rx errored and missed TCA. This field is reset to 0  on 15 min boundaries. (4 bytes)  Rx errored: Number of errored packets received from the node. The sum  of this field across all entries in the node table contributes to the Rx  errored TCA. This field is reset to 0 on 15  min boundaries. (optional)  (4 bytes)  (R) (mandatory) (37 * N bytes, where N is the number of nodes in the node  table)  Actions  Create, delete, get, get next, set 

---- page break ---- Notifications  Threshold crossing alert  Alarm  number Threshold crossing alert Threshold value attribute No. (Note)  0 Total rx errored and missed 1  1 Total rx errored 2  NOTE – This number associates the TCA with the specified threshold value attribute of the  threshold data 1 managed entity.  9.11 This clause is intentionally left blank  9.12 General purpose managed entities  9.12.1 UNI-G  This ME organizes data associated with UNIs supported by GEM. One instance of the UNI -G ME  exists for each UNI supported by the ONU.  The ONU automatically creates or deletes instances of this ME upon the creation or deletion of a real  or virtual circuit pack ME, one per port.  Relationships  An instance of the UNI-G ME exists for each instance of a PPTP ME.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. Through an  identical ID, this ME is implicitly linked to an instance of a PPTP. (R)  (mandatory) (2 bytes)  Deprecated: This attribute is not used. It should be set to 0 by the OLT and ignored by the  ONU. (R, W) (mandatory) (2 bytes)  Administrative state: This attribute locks (1) and unlocks (0) the functions performed by this  ME. Administrative state is further described in clause A.1.6. (R,  W)  (mandatory) (1 byte)  NOTE – PPTP MEs also have an administrative state attribute. The user port is  unlocked only if both administrative state attributes are set to unlocked. It is  recommended that this attribute not be used: that the OLT set it to 0 and that the ONU  ignore it.  Management capability: An ONU may support the ability for some or all of its PPTPs to be  managed either directly by the OMCI or from a non-OMCI management  environment such as [BBF TR -069]/[BBF TR-369]. This attribute advertises  the ONU's capabilities for each PPTP.  This attribute is an enumeration with the following code points:  0 OMCI only  1 Non-OMCI only. In this case, the PPTP may be visible to the OMCI,  but only in a read-only sense, e.g., for PM collection.  2 Both OMCI and non-OMCI  (R) (optional) (1 byte)  Non-OMCI management identifier: If a PPTP can be managed either directly by the OMCI  or a non-OMCI management environment, this attribute specifies how it is in  fact to be managed. This attribute is either 0 (default  = OMCI management),
```
