# Managed Entity

## Identity
- ME ID: 262
- ME Name: T-CONT
- Source Section: 9.2.2
- Source Page: 103

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Get, Set
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: Alloc-ID
- Size: 2 bytes
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 2
- Name: Deprecated
- Size: 1 byte
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 3
- Name: Policy
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: mandatory

## Extraction Status
- Status: auto_extracted
- Attributes found: 3
- Review needed: false

## Raw Source

```
9.2.2 T-CONT  An instance of the traffic container ME T-CONT represents a logical connection group associated  with a G-PON PLOAM layer alloc-ID. A T-CONT can accommodate GEM packets in priority queues  or traffic schedulers that exist in the GEM layer.  The ONU autonomously creates instances of this ME. The OLT can discover the number of T-CONT  instances via the ANI-G ME. When the ONU's MIB is reset or created for the first time, all supported  T-CONTs are created. The OLT provisions alloc -IDs to the ONU via the PLOAM channel. Via the  OMCI, the OLT must then set the alloc-ID attributes in the T-CONTs that it wants to activate for user  traffic, to create the appropriate association with the allocation ID in the PLOAM channel. There  should be a one -to-one relationship between allocation IDs and T -CONT MEs; the connection of  multiple T-CONTs to a single allocation ID is undefined.   The allocation ID that matches the ONU-ID itself is defined to be the default alloc-ID. This alloc-ID  is used to carry the OMCC. The default alloc-ID can also be used to carry user traffic, and hence can  be assigned to one of the T -CONT MEs. However, this OMCI relationship only pertains to user  traffic, and the OMCC relationship is unaffected. It can also be true that the OMCC is not contained  in any T -CONT ME construct; rather, that the OMCC remains outside of the OMCI, and that the  OMCI is not used to man age the OMCC in any way. Multiplexing of the OMCC and user data in  G-PON systems is discussed in clause B.2.4.  Relationships  One or more instances of this ME are associated with an instance of a circuit pack that supports  a PON interface function, or with the ONU-G itself. 

---- page break ---- Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. This 2 byte  number indicates the physical capability that realizes the T -CONT. It may be  represented as 0xSS BB, where SS indicates the slot ID that contains this T - CONT (0 for the ONU as a whole), and BB is the T -CONT ID, numbered by  the ONU itself. T -CONTs are numbered in ascending order, with the range  0..255 in each slot. (R) (mandatory) (2 bytes)  Alloc-ID: This attribute links the T-CONT with the alloc-ID assigned by the OLT in the  assign_alloc-ID PLOAM message. The respective TC layer specification  should be referenced for the legal values for that system. Prior to the setting of  this attribute by the OLT, this attribute has an unambiguously unusable initial  value, namely the value 0x00FF or 0xFFFF for ITU-T G.984 systems, and the  value 0xFFFF for all other ITU -T GTC based PON systems . (R, W)  (mandatory) (2 bytes)  Deprecated: The ONU should set this attribute to the value 1, and the OLT should ignore  it. (R) (mandatory) (1 byte)  Policy: This attribute indicates the T-CONT's traffic scheduling policy. Valid values:  0 Null  1 Strict priority  2 WRR – Weighted round robin  (R, W) (mandatory) (1 byte)  NOTE – This attribute is read -only, unless otherwise specified by the QoS  configuration flexibility attribute of the ONU2-G ME. If flexible configuration is not  supported, the ONU should reject an attempt to set it with a parameter error result - reason code.  Actions  Get, set  Notifications  None.  
```
