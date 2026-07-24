# Managed Entity

## Identity
- ME ID: 46
- ME Name: MAC bridge configuration data
- Source Section: 9.3.2
- Source Page: 141

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Get
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: Bridge MAC address
- Size: 6 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 2
- Name: Bridge priority
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 3
- Name: Designated root
- Size: 8 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 4
- Name: Root path cost
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 5
- Name: Bridge port count
- Size: 1 byte
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 6
- Name: Root port num
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 7
- Name: Hello time
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 8
- Name: Forward delay
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: optional

## Extraction Status
- Status: auto_extracted
- Attributes found: 8
- Review needed: false

## Raw Source

```
9.3.2 MAC bridge configuration data  This ME organizes status data associated with a MAC bridge. The ONU automatically creates or  deletes an instance of this ME upon the creation or deletion of a MAC bridge service profile.  Relationships  This ME is associated with one instance of a MAC bridge service profile.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. Through an  identical ID, this ME is implicitly linked to an instance of the MAC bridge  service profile. (R) (mandatory) (2 bytes)  Bridge MAC address : This attribute indicates the MAC address used by the bridge. The  ONU sets this attribute to a value based on criteria beyond the scope of this  Recommendation, e.g., factory settings. (R) (mandatory) (6 bytes)  Bridge priority: This attribute reports the priority of the bridge. The ONU copies this attribute  from the priority attribute of the associated MAC bridge service profile. The  value of this attribute changes with updates to the MAC bridge service profile  priority attribute. (R) (mandatory) (2 bytes)  Designated root : This attribute identifies the bridge at the root of the spanning tree. It  comprises bridge priority (2  bytes) and MAC address (6  bytes). (R)  (mandatory) (8 bytes)  Root path cost: This attribute reports the cost of the best path to the root as seen from this  bridge. Upon ME instantiation, the ONU sets this attribute to 0. (R)  (mandatory) (4 bytes)  Bridge port count : This attribute records the number of ports linked to this bridge. (R)  (mandatory) (1 byte)  Root port num : This attribute contains the port number that has the lowest cost from the  bridge to the root bridge. The value 0 means that this bridge is itself the root.  Upon ME instantiation, the ONU sets this attribute to 0. (R) (mandatory)  (2 bytes)  Hello time: This attribute is the hello time received from the designated root,  the interval  (in 256ths of a second) between HELLO packets. Its range is 0x0100 to  0x0A00 (1..10 s). (R) (optional) (2 bytes)  NOTE – [IEEE 802.1D] specifies the compatibility range for hello time to be 1..2 s.  Forward delay: This attribute is the forwarding delay time received from the designated root  (in 256ths of a second). Its range is 0x0400 to 0x1E00 (4..30 s) in accordance  with [IEEE 802.1D]. (R) (optional) (2 bytes) 

---- page break ---- Actions  Get  Notifications  None.  
```
