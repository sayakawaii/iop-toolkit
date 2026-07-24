# Managed Entity

## Identity
- ME ID: 464
- ME Name: Synchronous Ethernet operation
- Source Section: 9.12.21
- Source Page: 446

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Get, Set
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: Administrative state
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 2
- Name: Options
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 3
- Name: ESMC QL state
- Size: 2 byte
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 4
- Name: ESMC QL generation
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 5
- Name: Tx ESMC Informational PDUs
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 6
- Name: Rx ESMC Informational PDUs
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 7
- Name: Tx ESMC Event PDUs
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 8
- Name: Rx ESMC Event PDUs
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

## Extraction Status
- Status: auto_extracted
- Attributes found: 8
- Review needed: false

## Raw Source

```
9.12.21 Synchronous Ethernet operation  This ME configures the ONU's ability to monitor and control the Synchronous Ethernet (SyncE)  operation capability as defined in [ITU G.8264].  An ONU that supports SyncE automatically creates or deletes an instance of this ME.  Relationships  A single instance of this ME is associated with the ONU-G ME.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. There is only  one instance, number 0. (R) (mandatory) (2 bytes)  Administrative state: This attribute locks (1) and unlocks (0) the functions performed by this  ME. Administrative state is further described in clause A.1.6. (R,  W)  (mandatory) (1 byte)  Options: This attribute configures the expected network QL states (please refer to  clause 11.3.2.1 of [ITU G.8264] on QL generation. The following QL types  are supported:  0 ITU-T G.8264 EEC option 1 network  1 ITU-T G.8264 EEC option 2 network  2 ITU-T G.8262.1 eEEC option 1 network  3 ITU-T G.8262.1 eEEC option 2 network  The default value is 0. (R, W) (mandatory) (1 byte)  ESMC QL state: This attribute is an enumeration that returns the received ONU SSM and  enhanced SSM code. See clause 11.3.2.2 QL Reception, Table 11-7 and Table  11-8 in [ITU G.8264]. Its values are defined as follows.  Options ITU-T G.8264 QL State Byte 1 (LSB)  SSM Code  Byte 2   Enhanced SSM  Code  EEC Option 1 QL-PRC  QL-SSU-A  QL-SSU-B  QL-EEC1  QL-DNU   Other values are reserved.  0010  0100  1000  1011  1111  0xFF  0xFF  0xFF  0xFF  0xFF  EEC Option 2 QL-STU  QL-PRS  QL-TNC  QL-ST2  QL-EEC2  QL-ST3  0000  0001  0100  0111  1010  1010  0xFF  0xFF  0xFF  0xFF  0xFF  0xFF 

---- page break ---- Options ITU-T G.8264 QL State Byte 1 (LSB)  SSM Code  Byte 2   Enhanced SSM  Code  QL-ST3E  QL-PROV  QL-DUS   Other values are reserved.  1101  1110  1111  0xFF  0xFF  0xFF  eEEC Option 1 QL-PRC   QL-SSU-A  QL-SSU-B  QL-EEC1   QL-DNU   QL-PRTC  QL-ePRTC  QL-eEEC  QL-ePRC  0010  0100  1000  1011  1111  0010  0010  1011  0010  0xFF  0xFF  0xFF  0xFF  0xFF  0x20  0x21  0x22  0x23  eEEC Option 2 QL-STU  QL-PRS  QL-TNC  QL-ST2  QL-EEC2   QL-ST3  QL-ST3E  QL-PROV  QL-DUS   QL-PRTC  QL-ePRTC  QL-eEEC  QL-ePRC  0000  0001  0100  0111  1010  1010  1101  1110  1111  0001  0001  1010  0001  0xFF  0xFF  0xFF  0xFF  0xFF  0xFF  0xFF  0xFF  0xFF  0x20  0x21  0x22  0x23    (R) (mandatory) (2 byte)  ESMC QL generation : This attribute is an enumeration that returns the current generated  ONU ESM quality level. The ESMC QL values an associated SSM codes are  same as specified for the ESMC QL state attribute. See clause 11.3.2.1 QL  Generation in [ITU G.8264].   (R) (mandatory) (2 bytes)  Tx ESMC Informational PDUs : This attribute increments when an ESMC Informational  PDU is sent by the ONU. When full, the counter rolls over to 0. (R) (mandatory)  (4 bytes)  Rx ESMC Informational PDUs : This attribute increments when an ESMC Informational  PDU is received by the ONU. When full, the counter rolls over to 0. (R)  (mandatory) (4 bytes)  Tx ESMC Event PDUs: This attribute increments when an ESMC event PDU is sent by the  ONU. When full, the counter rolls over to 0. (R) (mandatory) (4 bytes)  Rx ESMC Event PDUs: This attribute increments when an ESMC event PDU is received by  the ONU When full, the counter rolls over to 0. (R) (mandatory) (4 bytes) 

---- page break ---- QL-DNU Event counter : This attribute increments when the QL state transitions to DNU  from a non-DNU state per section 11.3.2.2 QL Reception [ITU G.8264]. When  full, the counter rolls over to 0. (R) (optional) (4 bytes)  Actions  Get, set  Notifications  None  
```
