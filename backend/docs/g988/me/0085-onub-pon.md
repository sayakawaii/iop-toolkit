# Managed Entity

## Identity
- ME ID: 85
- ME Name: ONUB-PON
- Source Section: 9.1.1
- Source Page: 57

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Get, Set, Test
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: Vendor ID
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 2
- Name: Version
- Size: 14 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 3
- Name: Serial number
- Size: 8 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 4
- Name: Traffic management option
- Size: 1 byte
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 5
- Name: Deprecated
- Size: 1 byte
- Format: needs_review
- Access: R
- Category: optional

### Attribute 6
- Name: Battery backup
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 7
- Name: Operational state
- Size: 1 byte
- Format: needs_review
- Access: R
- Category: optional

### Attribute 8
- Name: ONU survival time
- Size: 1 byte
- Format: needs_review
- Access: R
- Category: optional

### Attribute 9
- Name: Logical ONU ID
- Size: 24 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 10
- Name: Logical password
- Size: 12 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 11
- Name: Credentials status
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 12
- Name: Values include
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: optional

## Extraction Status
- Status: auto_extracted
- Attributes found: 12
- Review needed: false

## Raw Source

```
9.1.1 ONU-G  This ME represents the ONU as equipment. The ONU automatically creates an instance of this ME.  It assigns values to read-only attributes according to data within the ONU itself.  This ME has evolved from the ONT-G of [ITU-T G.984.4].  Relationships  In ITU-T GTC based PON  applications, all other MEs in this Recommendation are related  directly or indirectly to the ONU-G entity.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. There is only  one instance, number 0. (R) (mandatory) (2 bytes)  Vendor ID: This attribute identifies the vendor of the ONU. It is the same as the four most  significant bytes of the ONU serial number as specified in the respective  transmission convergence (TC) layer specification. (R) (mandatory) (4 bytes)  Version: This attribute identifies the version of the ONU as defined by the vendor. The  character value 0 indicates that  version information is not available or  applicable. (R) (mandatory) (14 bytes)  Serial number: The serial number is unique for each ONU. It is defined in the respective TC  layer specification and contains the vendor ID and version number. The first  four bytes are an ASCII-encoded four-letter vendor ID. The second four bytes  are a binary encoded serial number, under the control of the ONU vendor. (R)  (mandatory) (8 bytes)  Traffic management option: This attribute identifies the upstream traffic management  function implemented in the ONU. There are three options:  0 Priority controlled and flexibly scheduled upstream traffic. The traffic  scheduler and priority queue mechanism are used for upstream traffic.  1 Rate controlled upstream traffic. The maximum upstream traffic of  each individual connection is guaranteed by shaping.  2 Priority and rate controlled. The traffic scheduler and priority queue  mechanism are used for upstream traffic. The maximum upstream  traffic of each individual connection is guaranteed by shaping.  For a further explanation, see Appendix II.  Downstream priority queues are managed via the GEM port network CTP ME.  Upon ME instantiation, the ONU sets this attribute to the value that describes  its implementation. The OLT must adapt its model to conform to the ONU 's  selection. (R) (mandatory) (1 byte)  Deprecated: This attribute is not used. If it is present, it should be set to 0. (R) (optional)  (1 byte)  Battery backup: This Boolean attribute controls whether the ONU performs backup battery  monitoring (assuming it is capable of doing so). False disables battery alarm  monitoring; true enables battery alarm monitoring. (R,  W) (mandatory)  (1 byte) 

---- page break ---- Administrative state: This attribute locks (1) and unlocks (0) the functions performed by  the ONU as an entirety. Administrative state is further described in clause  A.1.6. (R, W) (mandatory) (1 byte)  Operational state: This attribute reports whether the ME is currently capable of performing  its function. Valid values are enabled (0) and disabled (1). (R) (optional)  (1 byte)  ONU survival time : This attribute indicates the minimum guaranteed time in milliseconds  between the loss of external power and the silence of the ONU. This does not  include survival time attributable to a backup battery. The value zero implies  that the actual time is not known. (R) (optional) (1 byte)  Logical ONU ID : This attribute provides a way for the ONU to identify itself. It is a text  string, null terminated if it is shorter than 24 bytes, with a null default value.  The mechanism for creation or modification of this information is beyond the  scope of this Recom mendation, but might include , for example, a web page  displayed to a user. (R) (optional) (24 bytes)  Logical password : This attribute provides a way for the ONU to submit authentication  credentials. It is a text string, null terminated if it is shorter than 12 bytes, with  a null default value. The mechanism for creation or modification of this  information is beyond the scope of this Recommendation. (R) (optional)  (12 bytes)  Credentials status : This attribute permits the OLT to signal to the ONU whether its  credentials are valid or not. The behaviour of the ONU is not specified, but  might, for example , include displaying an error screen to the user. (R, W)  (optional) (1 byte)  Values include:  0 Initial state, status indeterminate  1 Successful authentication  2 Logical ONU ID (LOID) error  3 Password error  4 Duplicate LOID  Other values are reserved.  Extended TC-layer options: This attribute is meaningful in ITU-T G.984 systems only. It is  a bit map that defines whether the ONU supports (1) or does not support (0)  various optional TC -layer capabilities of [ITU-T G.984.3]. Bits are assigned  as follows.  Bit Meaning  1 (LSB) Annex C of [ITU-T G.984.3], PON-ID maintenance.  2 Annex D of [ ITU-T G.984.3], PLOAM channel enhancements:  swift_POPUP and Ranging_adjustment messages.  3..16 Reserved  (R) (optional) (2 bytes)  Actions  Get, set  Reboot: Reboot the ONU.  Test: Test the ONU. The test action can be used either to perform equipment  diagnostics or to measure parameters such as received optical power, video 

---- page break ---- output level, battery voltage, etc. Test and test result messages are defined in  Annex A.  Synchronize time: This action synchronizes the start time of all PM MEs of the ONU with  the reference time of the OLT. All counters of all PM MEs are cleared to 0 and  restarted. Also, the value of the interval end time attribute of the PM MEs is  set to 0 and restarted. See clause I.4 for further discussion of PM.  NOTE – This function is intended only to establish rough 15  min boundaries for PM  collection. High precision time of day synchronization is a separate function,  supported by the OLT-G ME.  Notifications  Test result: Test results are reported via a test result message if the test is invoked by a test  command from the OLT.    Attribute value change  Number Attribute value change
```
