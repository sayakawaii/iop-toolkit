# Managed Entity

## Identity
- ME ID: 268
- ME Name: GEM port network CTP
- Source Section: 9.2.3
- Source Page: 104

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Create, Delete, Get, Set
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: Port-ID
- Size: 2 bytes
- Format: needs_review
- Access: RWSC
- Category: mandatory

### Attribute 2
- Name: T-CONT pointer
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 3
- Name: Direction
- Size: 1 byte
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 4
- Name: Traffic management pointer for upstream
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 5
- Name: Traffic descriptor profile pointer for upstream
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: optional

### Attribute 6
- Name: UNI counter
- Size: 1 byte
- Format: needs_review
- Access: R
- Category: optional

### Attribute 7
- Name: Priority queue pointer for downstream
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 8
- Name: Traffic descriptor profile pointer for downstream
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: optional

### Attribute 9
- Name: Encryption key ring
- Size: 1 byte
- Format: needs_review
- Access: R, W, set-by-create
- Category: optional

## Extraction Status
- Status: auto_extracted
- Attributes found: 9
- Review needed: false

## Raw Source

```
9.2.3 GEM port network CTP  This ME represents the termination of a GEM port on an ONU. This ME aggregates connectivity  functionality from the network view and alarms from the network element view as well as artefacts  from trails.  Instances of the GEM port network CTP ME are created and deleted by the OLT. A n instance of   GEM port network CTP can be deleted only when no GEM IW TP or GEM port network CTP PM  history data are associated with it. It is the responsibility of the OLT to make sure that the ONU  configuration meets this condition.  In ITU-T G.984 systems, when a GEM port network CTP is created, its encryption state is by default  not encrypted. If the OLT wishes to configure the GEM port to use encryption, it must send the  appropriate PLOAM message. This applies equally to new CTPs and to CTPs that are re-created after  an MIB reset.  In ITU-T G.987 systems, GEM ports are dynamically encrypted. If it is intended to encrypt the GEM  port, the OLT must configure a key ring to be used, and the key must be known to the ONU at run  time. 

---- page break ---- Relationships  An instance of the GEM port network CTP ME may be associated with an instance of the T- CONT and GEM IW TP MEs.  Attributes  Managed entity  ID: This attribute uniquely identifies each instance of this ME. (R,  set-by-create) (mandatory) (2 bytes)  Port-ID: This attribute is the port -ID of the GEM port associated with this CTP.  (RWSC) (mandatory) (2 bytes)  NOTE 1 – While nothing forbids the existence of several GEM port network CTPs  with the same port-ID value, downstream traffic is modelled as being delivered to all  such GEM port network CTPs. Be aware of potential difficulties associated with  defining downstream flows and aggregating PM statistics.  T-CONT pointer : This attribute points to a T-CONT instance. (R, W, set-by-create)  (mandatory) (2 bytes)  Direction: This attribute specifies whether the GEM port is used for UNI -to-ANI (1),  ANI-to-UNI (2), or bidirectional (3) connection. (R,  W, set-by-create)  (mandatory) (1 byte)  Traffic management pointer for upstream: If the traffic management option attribute in the  ONU-G ME is 0 (priority controlled) or 2 (priority and rate controlled), this  pointer specifies the priority queue ME serving this GEM port network CTP.  If the traffic management option attribute is 1 (rate controlled), this attribute  redundantly points to the T -CONT serving this GEM port  network CTP.  (R, W, set-by-create) (mandatory) (2 bytes)  Traffic descriptor profile pointer for upstream: This attribute points to the instance of the  traffic descriptor ME that contains the upstream traffic parameters for this  GEM port network CTP. This attribute is used when the traffic management  option attribute in the ONU -G ME is 1 (rate controlled), specifying the  PIR/PBS to which the upstream traffic is shaped. This attribute is also used  when the traffic management option attribute in the ONU-G ME is 2 (priority  and rate controlled), specifying the CIR/CBS/PIR/PBS to which the upstream  traffic is policed. (R, W, set-by-create) (optional) (2 bytes)  See also Appendix II.  UNI counter: This attribute reports the number of instances of UNI -G ME associated with  this GEM port network CTP. (R) (optional) (1 byte)  Priority queue pointer for downstream: This attribute points to the instance of the priority  queue used for this GEM port network CTP in the downstream direction. It is  the responsibility of the OLT to provision the downstream pointer in a way  that is consistent with the bridge and mapper connectivity. If the pointer is null,  downstream queueing is determined by other mechanisms in the ONU. (R, W,  set-by-create) (mandatory) (2 bytes)  NOTE 2 – If the GEM port network CTP is associated with more than one UNI  (downstream multicast), the downstream priority queue pointer defines a pattern (e.g.,  queue number 3 for a given UNI) to be replicated (i.e., to queue number 3) at the other  affected UNIs. 

---- page break ---- Encryption state: This attribute indicates the current state of the GEM port network CTP 's  encryption. Legal values are defined to be the same as those of the security  mode attribute of the ONU2 -G, with the exception that attribute value 0  indicates an unencrypted GEM port. (R) (optional) (1 byte)  Traffic descriptor profile pointer for downstream : This attribute points to the instance of  the traffic descriptor ME that contains the downstream traffic parameters for  this GEM port network CTP. This attribute is used when the traffic  management option attribute in the ONU -G ME is 1 (rate controlled),  specifying the PIR/PBS to which the downstream traffic is shaped. This  attribute is also used when the traffic management option attribute in the ONU- G ME is 2 (priority and rate controlled), specifying the CIR/CBS/PIR/PBS to  which the downstream traffic is policed. (R, W, set-by-create) (optional)  (2 bytes)  See also Appendix II.  Encryption key ring : This attribute is defined in ITU -T G.987 systems only. It specifies  whether the associated GEM port is encrypted, and if so, which key ring it  uses. (R, W, set-by-create) (optional) (1 byte)  0 (default) No encryption. The downstream key index is ignored, and  upstream traffic is transmitted with key index 0.  1 Unicast payload encryption in both directions. Keys are generated by  the ONU and transmitted to the OLT via the PLOAM channel.  2 Broadcast (multicast) encryption. Keys are generated by the OLT and  distributed via the OMCI.  3 Unicast encryption, downstream only. Keys are generated by the ONU  and transmitted to the OLT via the PLOAM channel.  Other values are reserved.  Actions  Create, delete, get, set  Notifications  Alarm  Alarm  number  Alarm Description  0..4 Reserved   5 End-to-end loss of continuity Loss of continuity can be detected when the GEM port  network CTP supports a GEM interworking  termination point (optional).  6..207 Reserved   208..223 Vendor-specific alar
```
