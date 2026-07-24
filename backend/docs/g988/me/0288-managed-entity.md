# Managed Entity

## Identity
- ME ID: 288
- ME Name: Managed entity
- Source Section: 9.12.9
- Source Page: 60

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Get, Set
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: Vendor product code
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 2
- Name: Security mode
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 3
- Name: Total traffic scheduler number
- Size: 1 byte
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 4
- Name: Deprecated
- Size: 1 byte
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 5
- Name: Total GEM port -ID number
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 6
- Name: SysUpTime
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 7
- Name: Connectivity capability
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 8
- Name: Current connectivity mode
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 9
- Name: MIB data sync
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: mandatory

## Extraction Status
- Status: auto_extracted
- Attributes found: 9
- Review needed: false

## Raw Source

```
9.12.9 Managed entity and 9.12.10 Attribute ME 

---- page break ---- for retrieval of ITU-T G.988 supported ME, and clause I.1.2 for discovering  the ONU hardware and service configuration of the ONU through MIB  uploads.)  NOTE 2 – 0xBz values greater than B5 are reserved to report new  ITU-T G.988 Annex A message versions (relative to ITU-T G.988 ( 2014)  Amd1).    0x80 ITU-T G.984.4 (06/04)  NOTE – For historical reasons, this code point may also appear in  ONUs that support later versions of [ITU-T G.984.4].  0x81 ITU-T G.984.4 2004 Amd.1 (06/05)  0x82 ITU-T G.984.4 2004 Amd.2 (03/06)  0x83 ITU-T G.984.4 2004 Amd.3 (12/06)  0x84 ITU-T G.984.4 2008 (02/08)  0x85 ITU-T G.984.4 2008 Amd.1 (06/09)  0x86 ITU-T G.984.4 2008 Amd.2 (2009). Baseline message set  only, without the extended message set option  0x96 ITU-T G.984.4 2008 Amd.2 (2009). Extended message set  option, in addition to the baseline message set.  0xA0 ITU-T G.988 (2010). Baseline message set only, without the  extended message set option  0xA1 ITU-T G.988 Amd.1 (2011). Baseline message set only  0xA2 ITU-T G.988 Amd.2 (2012). Baseline message set only  0xA3 ITU-T G.988 (2012). Baseline message set only  0xA4 ITU-T G.988 Amd. 1 (2014). Baseline message set only  0xB0 ITU-T G.988 (2010). Baseline and extended message set  0xB1 ITU-T G.988 Amd.1 (2011). Baseline and extended message  set  0xB2 ITU-T G.988 Amd.2 (2012). Baseline and extended message  set  0xB3 ITU-T G.988 (2012). Baseline and extended message set  0xB4 ITU-T G.988 Amd1 (2014) Baseline and extended message  set  0xB5 Do not use  (R) (mandatory) (1 byte)  Vendor product code: This attribute contains a vendor -specific product code for the ONU.  (R) (optional) (2 bytes) 

---- page break ---- Security capability : This attribute advertises the security capabilities of the ONU. The  following code points are defined:   0 Reserved   1 Advanced encryption standard -128 (AES -128) payload  encryption supported   2 Reserved   3 AES-128 and AES-256 payload encryption supported   4…6 Reserved   7 AES-128, AES -256, and Camellia -128 payload encryption  supported   8…10 Reserved   11 AES-128, AES -256, and Camellia -256 payload encryption  supported   12…14 Reserved   15 AES-128, AES-256, Camellia-128 and Camellia-256 payload  encryption supported   16-18 Reserved   19 AES-128, AES -256, and SM4( -128) payload encryption  supported   20…22 Reserved   23 AES-128, AES -256, SM4( -128), and Camel lia-128 payload  encryption supported   24…26 Reserved   27 AES-128, AES -256, SM4( -128), and Camel lia-256 payload  encryption supported   28…30 Reserved   31 AES-128, AES-256, SM4(-128), Camellia-128, and Camellia- 256 payload encryption supported   32..255 Reserved  (R) (mandatory) (1 byte)  NOTE – Reporting of value 1 is not valid for ITU-T G.9804 ONU.  Security mode: This attribute specifies the current security mode of the ONU. All secure  (X)GEM ports in an ONU must use the same security mode at any given time.  The following code points are defined:   0 Reserved   1 AES-128 algorithm   2 AES-256 algorithm   3 Camellia-128 algorithm   4 Camellia-256 algorithm   5 SM4(-128) algorithm   6..255 Reserved  Upon ME instantiation, the ONU sets this attribute to  1, AES -128. After  initiation, on request from the OLT, this attribute may be set to one of the  algorithms supported by the ONU (as indicated in Security Capability  attribute). Setting this attribute to any value supported by the ONU  does not  imply that any of the (X)GEM ports are encrypted; that process is negotiated  at the PLOAM layer. It only signifies that the  indicated encryption algorithm  set by the OLTis the security mode to be used on any of the (X)GEM ports that  the OLT may choose to encrypt. (R, W) (mandatory) (1 byte)  NOTE – Values other than 1 are only valid for ITU-T G.9804 ONUs. 

---- page break ---- Total priority queue number: This attribute reports the total number of upstream priority  queues that are not associated with a circuit pack, but with the ONU in its  entirety. Upon ME instantiation, the ONU sets this attribute to the value that  represents its capabilities. (R) (mandatory) (2 bytes)  Total traffic scheduler number: This attribute reports the total number of traffic schedulers  that are not associated with a circuit pack, but with the ONU in its entirety.  The ONU supports null function, strict priority scheduling and weighted round  robin (WRR) from the priority control and guarantee of minimum rate control  points of view, respectively. If the ONU has no global traffic schedulers, this  attribute is 0. (R) (mandatory) (1 byte)  Deprecated: This attribute should always be set to 1 by the ONU and ignored by the OLT.  (R) (mandatory) (1 byte)  Total GEM port -ID number : This attribute reports the total number of GEM port -IDs  supported by the ONU. The maximum value is specified in the corresponding  TC recommendations. Upon ME instantiation, the ONU sets this attribute to  the value that represents its capabilities. (R) (optional) (2 bytes)  SysUpTime: This attribute counts 10 ms intervals since the ONU was last initialized. It rolls  over to 0 when full (see [IETF RFC 1213]). (R) (optional) (4 bytes)  Connectivity capability: This attribute indicates the Ethernet connectivity models that the  ONU can support. The value 0 indicates that the capability is not supported; 1  signifies support. The following code points are defined.    Bit Model  1 (LSB) N:1 bridging, Figure 8.2.2-3  2 1:M mapping, Figure 8.2.2-4  3 1:P filtering, Figure 8.2.2-5  4 N:M bridge-mapping, Figure 8.2.2-6  5 1:MP map-filtering, Figure 8.2.2-7  6 N:P bridge-filtering, Figure 8.2.2-8  7 N:MP bridge-map-filtering, Figure 8.2.2-9  8…16 Reserved    NOTE 1 – It is not implied that an ONU may not support other connectivity models.  (R) (optional) (2 bytes)  Current connectivity mode: This attribute specifies the Ethernet connectivity model that the  OLT wishes to use. The following code points are defined.    Value 
```
