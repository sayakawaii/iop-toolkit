# Managed Entity

## Identity
- ME ID: 63
- ME Name: Traffic schedulerB-PON
- Source Section: 9.2.11
- Source Page: 117

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Get, Set
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: T-CONT pointer
- Size: 2 bytes
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 2
- Name: Traffic scheduler pointer
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 3
- Name: Policy
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 4
- Name: Priority/weight
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: mandatory

## Extraction Status
- Status: auto_extracted
- Attributes found: 4
- Review needed: false

## Raw Source

```
9.2.11 Traffic scheduler  NOTE 1 – In [ITU-T G.984.4], this ME is called a traffic scheduler-G.  An instance of this ME represents a logical object that can control upstream GEM packets. A traffic  scheduler can accommodate GEM packets after a priority queue or other traffic scheduler and transfer  them towards the next traffic scheduler or T-CONT. Because T -CONTs and traffic schedulers are  created autonomously by the ONU, the ONU vendor predetermines the most complex traffic handling  model it is prepared to support; the OLT may use less than the ONU's full capabilities, but cannot ask  for more. See Appendix II for more details.  After the ONU creates instances of the T-CONT ME, it then autonomously  creates instances of the  traffic scheduler ME.  Relationships  The traffic scheduler ME may be related to a T -CONT or other traffic schedulers through  pointer attributes.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. This 2 byte  number indicates the physical capability that realizes the traffic scheduler. The  first byte is the slot ID of the circuit pack with which this traffic scheduler is  associated. For a traffic scheduler that is not associated with a circuit pack, the  first byte is 0xFF. The second byte is the traffic scheduler id, assigned by the  ONU itself. Traffic schedulers are numbered in ascending order with the range  0..0xFF in each circuit pack or in the ONU core. (R) (mandatory) (2 bytes)  T-CONT pointer: This attribute points to the T -CONT ME instance associated with this  traffic scheduler. This pointer is used when this traffic scheduler is connected  to the T-CONT directly; It is null (0) otherwise. (R, W) (mandatory) (2 bytes) 

---- page break ---- NOTE 2  – This attribute is read -only unless otherwise specified by the QoS  configuration flexibility attribute of the ONU2-G ME. If flexible configuration is not  supported, the ONU should reject an attempt to set the T-CONT pointer attribute with  a parameter error result-reason code.  Traffic scheduler pointer: This attribute points to another traffic scheduler ME instance that  may serve this traffic scheduler. This pointer is used when this traffic scheduler  is connected to another traffic s cheduler; it is null (0) otherwise.  (R)  (mandatory) (2 bytes)  Policy: This attribute represents scheduling policy. Valid values include:  0 Null  1 Strict priority  2 WRR (weighted round robin)  The traffic scheduler derives priority or weight values for its tributary traffic  schedulers or priority queues from the tributary MEs themselves.  (R, W) (mandatory) (1 byte)  NOTE 3 – This attribute is read -only unless otherwise specified by the QoS  configuration flexibility attribute of the ONU2-G ME. If flexible configuration is not  supported, the ONU should reject an attempt to set the policy attribute with a  parameter error result-reason code.  Priority/weight: This attribute represents the priority for strict priority scheduling or the  weight for WRR scheduling. This value is used by the next upstream ME, as  indicated by the T-CONT pointer attribute or traffic scheduler pointer attribute.  If the indicated pointer has policy  = strict priority, this value is interpreted as  a priority (0 is the highest priority, 255 the lowest).  If the indicated pointer has policy = WRR, this value is interpreted as a weight.  Higher values receive more bandwidth.  Upon ME instantiation, the ONU sets this attribute to  0. (R, W) (mandatory)  (1 byte)  Actions  Get, set  Notifications  None.  
```
