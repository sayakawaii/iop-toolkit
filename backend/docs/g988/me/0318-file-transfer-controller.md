# Managed Entity

## Identity
- ME ID: 318
- ME Name: File transfer controller
- Source Section: 9.12.13
- Source Page: 436

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Get, Set
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: Supported transfer protocols
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 2
- Name: File instance
- Size: 2 bytes
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 3
- Name: Local file name pointer
- Size: 2 bytes
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 4
- Name: File transfer trigger
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 5
- Name: VLAN
- Size: 2 bytes
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 6
- Name: File size
- Size: 4 bytes
- Format: needs_review
- Access: R, W
- Category: optional

## Extraction Status
- Status: auto_extracted
- Attributes found: 6
- Review needed: false

## Raw Source

```
9.12.13 File transfer controller  This optional ME allows file transfers to be conducted out of band. It is intended to facilitate software  image download, but may be used for other file transfer applications as well.  Relationships  One instance of this ME exists in an ONU that supports out-of-band file transfer.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. There is one  instance, whose value is 0. (R) (mandatory) (2 bytes)  Supported transfer protocols : This attribute is a bit map. Each bit indicates that the  corresponding protocol is (1) or is not (0) supported for file transfers. TFTP is  mandatory; the other protocols are optional. (R) (mandatory) (2 bytes)  Bit Protocol  1 (LSB) FTP  2 TFTP  3 SFTP  4 HTTP  5 HTTPS  6 FLUTE ([b-IETF RFC 3926]; multicast, download only)  7 DSM-CC (multicast, download only)  8..12 Reserved  13..16 Vendor-specific use; not to be standardized  File type: This attribute specifies the owner ME type of the file to be transferred. It is a  value from Table 11.2.4-1. For example, for software image download, this  attribute has the value 7. File systems on circuit packs are identified with the  value 6, and on the ONU-G as a whole with the value 256. (R, W) (mandatory)  (2 bytes)  File instance: This attribute specifies the instance of the file to be transferred. This attribute  is the same as the ME ID attribute of the owner ME. (R,  W) (mandatory)  (2 bytes)  Local file name pointer: This attribute designates a large string ME that specifies a local file  name. Since naming is not specified for any OMCI files, this attribute should  be a null pointer in standard OMCI usage, e.g., in the software image case. The  use of this attribute for named files is vendor -specific. (R,  W) (mandatory)  (2 bytes) 

---- page break ---- Network address pointer: This attribute is a pointer to a network address ME that specifies  optional authentication information, along with the URI to be used for the file  transfer. The URI should specify the protocol, one from the list of protocols  supported by the ONU, an IP a ddress or a string that can be resolved into an  IP address, and optionally a port. For unidirectional multicast download (e.g.,  DSM-CC), the URI should specify a multicast IP source address. (R,  W)  (mandatory) (2 bytes)  File transfer trigger: This attribute causes the file transfer to begin. If a given set operation  writes values to several attributes of this ME, the ONU should apply the file  transfer trigger after updating all other attributes. Some operations may not be  applicable to some files; the ONU should deny commands that request  unsupported actions. (R, W) (mandatory) (1 byte)  Value Meaning  0 Reserved  1 Initiate file download (to the ONU)  2 Initiate file upload (from the ONU)  3 Abort current file transfer  4 Delete target file (on the ONU)  5 Perform a directory listing operation. The scope of the directory is  not specified; at the vendor 's option, the listing may be filtered by  matching some or all of file type, file instance and local file name  attributes.  6..255 Reserved  File transfer status : This attribute reports the status of a file transfer. (R) (mandatory)  (1 byte)  Value Meaning  0 File transfer completed successfully  1 File transfer aborted successfully  2 File deleted  3 URL undefined or unreachable  4 Failure to authenticate  5 File transfer in progress  6 Remote failure  7 Local failure  8..255 Reserved  GEM IWTP pointer: This attribute is a pointer that specifies a unicast or multicast GEM IW  TP, depending on whether the transfer protocol to be used is unicast or  multicast. (R, W) (optional) (2 bytes)  VLAN: This attribute specifies the VLAN to be used for the transfer, assuming  multicast protocol. The default value 0 indicates that no VLAN is specified.  (R, W) (optional) (2 bytes)  File size: This attribute allows the OLT to specify the size of a file to be downloaded, in  bytes. The ONU may use this value to reserve memory or to deny the download  command if it has insufficient space. The default value 0 does not specify a  file size. (R, W) (optional) (4 bytes)  Directory listing table: When a directory listing is complete, this attribute contains the result  of a directory listing operation. The content and format of the table is not  specified. (R) (optional) (N bytes) 

---- page break ---- Actions  Get, set  Notifications  Attribute value change  Number Attribute value change Description  1..6 N/A   7 File transfer status   8..10 N/A   11 Directory listing table This AVC signals to the OLT that a directory listing  operation is complete and may be retrieved with a get, get  next sequence.   12..16 Reserved   
```
