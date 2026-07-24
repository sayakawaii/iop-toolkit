# Managed Entity

## Identity
- ME ID: 45
- ME Name: MAC bridge service profile
- Source Section: 9.3.1
- Source Page: 140

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Create, Delete, Get, Set
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: Spanning tree ind
- Size: 1 byte
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 2
- Name: Learning ind
- Size: 1 byte
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 3
- Name: Port bridging ind
- Size: 1 byte
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 4
- Name: Priority
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 5
- Name: Max age
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 6
- Name: Hello time
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 7
- Name: Forward delay
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 8
- Name: Unknown MAC address discard
- Size: 1 byte
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 9
- Name: MAC learning depth
- Size: 1 byte
- Format: needs_review
- Access: R, W, set-by-create
- Category: optional

### Attribute 10
- Name: Dynamic filtering ageing time
- Size: 4 bytes
- Format: needs_review
- Access: R, W, set- by-create
- Category: optional

## Extraction Status
- Status: auto_extracted
- Attributes found: 10
- Review needed: false

## Raw Source

```
9.3.1 MAC bridge service profile  This ME models a MAC bridge in its entirety; any number of ports may be associated with the bridge  through pointers to the MAC bridge service profile ME. Instances of this ME are created and deleted  by the OLT.  Relationships  Bridge ports are modelled by MAC bridge port configuration data MEs, any number of which  can point to a MAC bridge service profile. The real-time status of the bridge is available from  an implicitly linked MAC bridge configuration data ME.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. The first byte  is the slot ID. In an integrated ONU, this value is 0. The second byte is the  bridge group ID. (R, set-by-create) (mandatory) (2 bytes)  Spanning tree ind: The Boolean value true specifies that a spanning tree algorithm is enabled.  The value false disables (rapid) spanning tree. (R,  W, set-by-create)  (mandatory) (1 byte)  Learning ind: The Boolean value true specifies that bridge learning functions are enabled.  The value false disables bridge learning. (R,  W, set-by-create) (mandatory)  (1 byte)  Port bridging ind : The Boolean value true specifies that bridging between UNI ports is  enabled. The value false disables local bridging. (R,  W, set-by-create)  (mandatory) (1 byte)  Priority: This attribute specifies the bridge priority in the range 0..65535. The value of  this attribute is copied to the bridge priority attribute of the associated MAC  bridge configuration data ME. (R, W, set-by-create) (mandatory) (2 bytes)  Max age: This attribute specifies the maximum age (in 256ths of a second) of received  protocol information before its entry in the spanning tree listing  is discarded.  The range is 0x0600 to 0x2800 (6..40  s) in accordance with [IEEE  802.1D].  (R, W, set-by-create) (mandatory) (2 bytes)  Hello time: This attribute specifies how often (in 256ths of a second) the bridge advertises  its presence via hello packets, while acting as a root or attempting to become  a root. The range is 0x0100 to 0x0A00 (1..10  s). (R,  W, set-by-create)  (mandatory) (2 bytes)  NOTE – [IEEE 802.1D] specifies the compatibility range for hello time to be 1..2 s.  Forward delay: This attribute specifies  the forwarding delay (in 256ths of a second) when  the bridge acts as the root. The range is 0x0400 to 0x1E00 (4..30  s) in  accordance with [IEEE 802.1D]. (R, W, set-by-create) (mandatory) (2 bytes)  Unknown MAC address discard : The Boolean value true specifies  that MAC frames with  unknown DAs be discarded. The value false specifies  that such frames be  forwarded to all allowed ports. (R, W, set-by-create) (mandatory) (1 byte)  MAC learning depth: This attribute specifies the maximum number of UNI MAC addresses  to be learned by the bridge. The default value 0 specifies that there is no  administratively imposed limit. (R, W, set-by-create) (optional) (1 byte)  Dynamic filtering ageing time : This attribute specifies the age of dynamic filtering entries  in the bridge database, after which unrefreshed entries are discarded. In  accordance with clause 7.9.2 of [IEEE 802.1D] and clause 8.8.3 of [IEEE 

---- page break ---- 802.1Q], the range is 10..1 000 000 s, with a resolution of 1 s and a default of  300 s. The value 0 specifies that the ONU uses its internal default. (R, W, set- by-create) (optional) (4 bytes)  Actions  Create, delete, get, set  Notifications  None.  
```
