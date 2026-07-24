# Managed Entity

## Identity
- ME ID: 8
- ME Name: UNIB-PON
- Source Section: 9.12.1
- Source Page: 424

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Get, Set
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: Deprecated
- Size: 2 bytes
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 2
- Name: Administrative state
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 3
- Name: It is  recommended that this attribute not be used
- Size: 1 byte
- Format: needs_review
- Access: R
- Category: optional

### Attribute 4
- Name: Non-OMCI management identifier
- Size: 2 bytes
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 5
- Name: Relay agent options
- Size: 2 bytes
- Format: needs_review
- Access: R, W
- Category: optional

## Extraction Status
- Status: auto_extracted
- Attributes found: 5
- Review needed: false

## Raw Source

```
9.12.1 UNI-G  This ME organizes data associated with UNIs supported by GEM. One instance of the UNI -G ME  exists for each UNI supported by the ONU.  The ONU automatically creates or deletes instances of this ME upon the creation or deletion of a real  or virtual circuit pack ME, one per port.  Relationships  An instance of the UNI-G ME exists for each instance of a PPTP ME.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. Through an  identical ID, this ME is implicitly linked to an instance of a PPTP. (R)  (mandatory) (2 bytes)  Deprecated: This attribute is not used. It should be set to 0 by the OLT and ignored by the  ONU. (R, W) (mandatory) (2 bytes)  Administrative state: This attribute locks (1) and unlocks (0) the functions performed by this  ME. Administrative state is further described in clause A.1.6. (R,  W)  (mandatory) (1 byte)  NOTE – PPTP MEs also have an administrative state attribute. The user port is  unlocked only if both administrative state attributes are set to unlocked. It is  recommended that this attribute not be used: that the OLT set it to 0 and that the ONU  ignore it.  Management capability: An ONU may support the ability for some or all of its PPTPs to be  managed either directly by the OMCI or from a non-OMCI management  environment such as [BBF TR -069]/[BBF TR-369]. This attribute advertises  the ONU's capabilities for each PPTP.  This attribute is an enumeration with the following code points:  0 OMCI only  1 Non-OMCI only. In this case, the PPTP may be visible to the OMCI,  but only in a read-only sense, e.g., for PM collection.  2 Both OMCI and non-OMCI  (R) (optional) (1 byte)  Non-OMCI management identifier: If a PPTP can be managed either directly by the OMCI  or a non-OMCI management environment, this attribute specifies how it is in  fact to be managed. This attribute is either 0 (default  = OMCI management), 

---- page break ---- or it is a pointer to a VEIP, which in turn links to a non -OMCI management  environment. (R, W) (optional) (2 bytes)  Relay agent options: This attribute is a pointer to a large string ME whose content specifies  one or more DHCP relay agent options. (R, W) (optional) (2 bytes)  The contents of the large string are parsed by the ONU and converted into text  strings. Variable substitution is based on defined three-character groups, each  of which begins with the '%' character. The string '%%' is an escape mechanism  whose output is a single '%' character. When the ONU cannot perform variable  substitution on a substring of the large string, it generates the specified option  as an exact quotation of the provisioned substring value.  Provisioning of the large string is separate from the operation of setting the  pointer in this attribute. It is the responsibility of the OLT to ensure that the  large string contents are correct and meaningful.  Three-character variable definitions are as follows. The first variable in the  large string must specify one of the option types. Both options for a given IP  version may be present if desired, each introduced by its option identifier.  Terminology is taken from clause 3.9.3 of [b-BBF TR-101].  %01, %18  Specifies that the following string is for option 82 sub-option 1,  agent circuit-ID (IPv4) or option 18, interface-ID (IPv6). The  equivalence permits the same large string to be used in both IP  environments.  %02, %37  Specifies that the following string is for option 82 sub-option 2,  relay agent remote-ID (IPv4) or option 37, relay agent  remote-ID (IPv6). The equivalence permits the same large string  to be used in both IP environments.  %SL In TR-101, this is called a slot. In an ONU, this variable refers to  a shelf. It is meaningful if the ONU has multiple shelves  internally or is daisy-chained to multiple equipment modules. The  range of this variable is "0".. "99"  %SU In TR -101, this is called a sub -slot. In fact, it represents a  cardholder. The range of this variable is "0".. "99"  %PO UNI port number. The range of this variable is "0".. "999"  %AE ATM or Ethernet. This variable can take on the values "atm" or  "eth".  %SV S-VID for Ethernet UNI, or ATM VPI for ATM UNI, as it exists  on the DHCP request received upstream across the UNI. Range  "0".. "4096" for S -VID; range "0".. "255" for VPI. The value  "4096" indicates no S-VID tag.  %CV C-VID (Q-VID) for Ethernet UNI or ATM VCI for ATM UNI,  as it exists on the DHCP request received upstream across the  UNI. Range "0".. "4096" for C-VID; range "0".."65535" for VCI.  The value "4096" indicates no C-VID tag.  Spaces in the provisioned string are significant.  Example: if the large string were provisioned with the value   %01%SL/%SU/%PO:%AE/%SV.%CV<null>, 

---- page break ---- then the ONU would generate the following DHCP option 82 agent circuit-ID  string for an Ethernet UNI that sent a DHCP request with no S tag and C tag =  3210 on shelf 2, slot 3, port 4.   2/3/4:eth/4096.3210  With the same provisioning, the ONU would generate the following DHCP  option 82 agent circuit -ID string for an ATM UNI that sent a DHCP request  on VPI = 123 and VCI = 4567 on shelf 2, slot 3, port 4.  2/3/4:atm/123.4567  Actions  Get, set  Notifications  None.  
```
