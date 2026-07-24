# Managed Entity

## Identity
- ME ID: 138
- ME Name: VoIP config data
- Source Section: 9.9.18
- Source Page: 413

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Get, Set
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: Available signalling protocols
- Size: 1 byte
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 2
- Name: Available VoIP configuration methods
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 3
- Name: VoIP configuration method used
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 4
- Name: VoIP configuration address pointer
- Size: 2 bytes
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 5
- Name: Retrieve profile
- Size: 1 byte
- Format: needs_review
- Access: W
- Category: mandatory

### Attribute 6
- Name: Profile version
- Size: 25 bytes
- Format: needs_review
- Access: R
- Category: mandatory

## Extraction Status
- Status: auto_extracted
- Attributes found: 6
- Review needed: false

## Raw Source

```
9.9.18 VoIP config data  The VoIP configuration data ME defines the configuration for VoIP in the ONU. The OLT uses this  ME to discover the VoIP signalling protocols and configuration methods supported by this ONU. The  OLT then uses this ME to select the desired signalling protocol and configuration method. The entity  is conditionally required for ONUs that offer VoIP services.  An ONU that supports VoIP services automatically creates an instance of this ME.  Relationships  One instance of this ME is associated with the ONU.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. There is only  one instance, number 0. (R) (mandatory) (2 bytes)  Available signalling protocols : This attribute is a bit map that defines the VoIP signalling  protocols supported in the ONU. The bit value 1 specifies that the ONU  supports the associated protocol.  1 (LSB) SIP  2 ITU-T H.248  3 MGCP  (R) (mandatory) (1 byte) 

---- page break ---- Signalling protocol used: This attribute specifies the VoIP signalling protocol to use. Only  one type of protocol is allowed at a time. Valid values are:  0 None  1 SIP  2 ITU-T H.248  3 MGCP  0xFF Selected by non-OMCI management interface  (R, W) (mandatory) (1 byte)  Available VoIP configuration methods : This attribute is a bit map that indicates the  capabilities of the ONU with regard to VoIP service configuration. The bit  value 1 specifies that the ONU supports the associated capability.  1 (LSB) ONU capable of using the OMCI to configure its VoIP services.  2 ONU capable of working with configuration file retrieval to  configure its VoIP services.  3 ONU capable of working with [BBF TR -069] to configure its  VoIP services.  4 ONU capable of working with IETF sipping config framework to  configure its VoIP services.  5 ONU capable of working with [BBF TR -369] to configure its  VoIP services.  Bits 6..24 are reserved by ITU -T. Bits 25..32 are reserved for proprietary  vendor configuration capabilities. (R) (mandatory) (4 bytes)  VoIP configuration method used : Specifies which method is used to configure the ONU 's  VoIP service.  0 Do not configure – ONU default  1 OMCI  2 Configuration file retrieval  3 BBF TR-069  4 IETF sipping config framework  5 BBF TR-369  6..240 Reserved by ITU-T  241..255 Reserved for proprietary vendor configuration methods  (R, W) (mandatory) (1 byte)  VoIP configuration address pointer : If this attribute is set to any value other than a null  pointer, it points to a network address ME, which indicates the address of the  server to contact using the method indicated in the VoIP configuration method  used attribute. This attribute is only relevant for non -OMCI configuration  methods.  If this attribute is set to a null pointer, no address is defined by this attribute.  However, the address may be defined by other methods, such as deriving it  from the ONU identifier attribute of the IP host config data ME and using a  well-known URI schema.  The default value is 0xFFFF (R, W) (mandatory) (2 bytes) 

---- page break ---- VoIP configuration state: Indicates the status of the ONU VoIP service.  0 Inactive: configuration retrieval has not been attempted  1 Active: configuration was retrieved  2 Initializing: configuration is now being retrieved  3 Fault: configuration retrieval process failed  Other values are reserved. At ME instantiation, the ONU sets this attribute to  0. (R) (mandatory) (1 byte)  Retrieve profile: This attribute provides a means by which the ONU may be notified that a  new VoIP profile should be retrieved. By setting this attribute, the OLT  triggers the ONU to retrieve a new profile. The actual value in the set action is  ignored because it is the ac tion of setting that is important. (W) (mandatory)  (1 byte)  Profile version : This attribute is a character string that identifies the version of the last  retrieved profile. (R) (mandatory) (25 bytes)  Actions  Get, set  Notifications  Attribute value change  Number Attribute value change Description  1..7 N/A   8 Profile version Version of last retrieved profile  9..16 Reserved     Alarm  Alarm  number Alarm Description  0 VCD config server name Failed to resolve the configuration server name.   1 VCD config server reach Cannot reach configuration server (the port cannot be  reached, ICMP errors)  2 VCD config server connect Cannot connect to the configuration server (due to bad  credentials or other faults after the port has responded)  3 VCD config server validate Cannot validate the configuration server  4 VCD config server auth Cannot authenticate the configuration session (e.g.,  missing credentials)  5 VCD config server timeout Timeout waiting for response from configuration server  6 VCD config server fail Failure response received from configuration server  7 VCD config file error Configuration file received has an error   8 VCD subscription name Failed to resolve the subscription server name  9 VCD subscription reach Cannot reach subscription server (the port cannot be  reached, ICMP errors)  10 VCD subscription connect Cannot connect to subscription server (due to bad  credentials or other faults after the port has responded)  11 VCD subscription validate Cannot validate subscription server 

---- page break ---- Alarm  Alarm  number Alarm Description  12 VCD subscription auth Cannot authenticate subscription session (e.g., missing  credentials)  13 VCD subscription timeout Timeout waiting for response from subscription server  14 VCD subscription fail Failure response received from subscription server  15 VCD reboot request A non-OMCI management interface has requested a  reboot of the ONU.  NOTE – This alarm is used only to indicate the request  and not to indicate that a reboot has actually taken place.  16  VCD Notify timeout Failure to receive the NOTIFY that the server is required  to send following acceptance of a SUBSCRIBE request.  17  VCD Notify malformed Malformed NOTIFY request  18  VCD Notify Reject
```
