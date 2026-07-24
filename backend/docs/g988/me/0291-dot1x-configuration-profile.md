# Managed Entity

## Identity
- ME ID: 291
- ME Name: Dot1X configuration profile
- Source Section: 9.3.15
- Source Page: 175

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Get, Set
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: Circuit ID prefix
- Size: 2 bytes
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 2
- Name: Fallback policy
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 3
- Name: Auth server 1
- Size: 2 bytes
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 4
- Name: Shared secret auth1
- Size: 25 bytes
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 5
- Name: The following two pairs of attributes are defined in the same way
- Size: 2 bytes
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 6
- Name: Shared secret auth2
- Size: 25 bytes
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 7
- Name: Auth server 3
- Size: 2 bytes
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 8
- Name: Shared secret auth3
- Size: 25 bytes
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 9
- Name: OLT proxy address
- Size: 4 bytes
- Format: needs_review
- Access: R, W
- Category: optional

## Extraction Status
- Status: auto_extracted
- Attributes found: 9
- Review needed: false

## Raw Source

```
9.3.15 Dot1X configuration profile  An instance of this ME represents a set of attributes that control an ONU 's 802.1X operation with  regard to IEEE 802 services. An instance of this ME is created by the ONU if it is capable of  supporting [IEEE 802.1X] authentication of CPE.  Relationships  One instance of this ME governs the ONU's 802.1X CPE authentication behaviour.  Attributes  Managed entity ID: This attribute provides a unique number for each instance of this ME.  There is at most one instance, number 0. (R) (mandatory) (2 bytes)  Circuit ID prefix: This attribute is a pointer to a large string ME whose content appears as  the prefix of the NAS port ID in radius access -request messages. The  remainder of the NAS port ID field is local information (for example, slot-port,  appended by the ONU itself). The default value of this attribute is the null  pointer 0. (R, W) (mandatory) (2 bytes)  Fallback policy: When set to 1 (deny), this attribute causes IEEE  802.1X conversations to  fail when no external authentication server is accessible, such that no Ethernet  service is provided. The default value 0 causes IEEE 802.1X conversations to  succeed when no external authentication server is accessible. (R,  W)  (mandatory) (1 byte)  Auth server 1 : This attribute is a pointer to a large string ME that contains the URI of the  first choice radius authentication server. The value 0 indicates that no radius  authentication server is specified. (R, W) (mandatory) (2 bytes)  Shared secret auth1 : This attribute is the shared secret for the first radius authentication  server. It is a null-terminated character string. (R, W) (mandatory) (25 bytes)  The following two pairs of attributes are defined in the same way:  Auth server 2: (R, W) (optional) (2 bytes)   Shared secret auth2: (R, W) (optional) (25 bytes)  Auth server 3: (R, W) (optional) (2 bytes)  Shared secret auth3: (R, W) (optional) (25 bytes)  OLT proxy address: This attribute indicates the IP address of a possible proxy at the OLT  for IEEE 802.1X radius messages. The default value 0.0.0.0 indicates that no  proxy is required. (R, W) (optional) (4 bytes) 

---- page break ---- Calling station ID format: Radius messages initiated by the ONU contain a calling-station- ID field that is specified to be the supplicant 's MAC address in upper -case  ASCII form, with bytes separated by a delimiter. This attribute permits  specification of the delimiter. (R, W) (optional) (2 bytes)   Value Meaning   0 ONU's internal default   1 Hyphen (-) delimiter   2 Colon (:) delimiter   3 No delimiter  0x20 – 0x7E Use this value as the delimiter  0xF0 – 0xFE Vendor-specific use  Other values are reserved.  Actions  Get, set  Notifications  None.  
```
