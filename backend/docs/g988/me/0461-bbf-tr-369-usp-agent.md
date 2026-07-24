# Managed Entity

## Identity
- ME ID: 461
- ME Name: BBF TR-369 USP agent
- Source Section: 9.12.20
- Source Page: 444

## Classification
- Access: needs_review
- Type: needs_review
- Actions: needs_review
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
9.12.20 BBF TR-369 USP agent  The BBF TR -369 USP agent ME models a USP Agent and the configuration needed to establish  communications between the USP Agent and a USP Controller.  [BBF TR -369] supports other means to discover its controller (e.g., pre -configured in firmware,  DHCP/DNS/mDNS discovery, and through another known trusted controller). This ME is therefore  optional (an ONU vendor may choose to implement it, and an operator may choose not to configure  the USP controller through this ME if supported by the ONU vendor).  An ONU that supports OMCI configuration of USP information automatically creates instances of  this ME   Relationships  An instance of the BBF TR-369 USP agent ME points to the virtual Ethernet interface point  (VEIP) ME. This ME is used for associating USP as a non-OMCI management domain. 

---- page break ---- Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME . Through an  identical ID, this ME is implicitly linked to an instance of a VEIP . (R)  (mandatory) (2 bytes)  Available USP Message Transfer Protocols: This table is a list of IANA registered service- name for the available MTPs:  Service name: The service name of the  USP MTP. It is a character string,  padded with trailing nulls if it is shorter than 15 bytes (15 bytes)  (R) (mandatory) (15N bytes)  Controller table : Each record in the table  (indexed by the controller network address)  is  comprised of the following fields:   Controller network address: This attribute points to an instance of a network  address ME that contains an FQDN or absolute URL identifying the  controller (e.g., "<scheme>://example.com:<port>[/<path>]). (2 bytes)  Table control: This field controls the meani ng of a set operation . The 1 byte  size of this field is included in get/get -next operations, but its value is  undefined under get-next and should be ignored by the OLT. (1 byte)  1 Add record to table; overwrite existing record, if any.  2 Delete record from table.  3 Clear all entries from table. This action may affect service and should  be used judiciously.  Other values are reserved.    Associated tag: This field is a TCI value for BBF TR -369 management  traffic passing through the VEIP. A TCI, comprising user priority, CFI  and VID, is represented by 2  bytes. The value 0xFFFF specifies that  BBF TR-369 management traffic passes through the VEIP with neither  a VLAN nor a priority tag. (2 bytes)  Provisioning Code: This attribute points to a large string ME that contains  the provisioning code (64 character UTF -8 string) that identifies the  primary service provider and other provisioning information, which  may be used by the Controller to determine service provider specific  customization and provisioning parameters. A 0xFFFF null pointer  indicates the absence of a provisioning code. (2 bytes)  USP retry minimum wait interval: This attribute configures the first retry  wait interval in seconds as specified in TR -369 "Failure Handling in  the session Context". The default value of 5 seconds is used. (2 bytes)  USP retry interval multiplier: This attribute configures the retry interval  multiplier as specified in TR -369 "Failure Handling in the session  Context". This value must be expressed in units of 0.001. Values range  between 1000 and 65535. The default value of 2000 is used. (2 bytes)  Message Transfer Protocol (MTP) Used : The service name of the USP  MTP. It is a character string, padded with trailing nulls if it is shorter  than 15 bytes (15 bytes)  (R,W) (mandatory) (26N bytes) 

---- page break ---- Actions  Get, get next, set  Set table (optional)  Notifications  None.  
```
