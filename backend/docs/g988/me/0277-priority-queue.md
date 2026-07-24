# Managed Entity

## Identity
- ME ID: 277
- ME Name: Priority queue
- Source Section: 9.2.10
- Source Page: 113

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Get, Set
- Instance Type: needs_review
- Instance Value: needs_review

## Attributes

### Attribute 1
- Name: Maximum queue size
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 2
- Name: Allocated queue size
- Size: 2 bytes
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 3
- Name: Discard-block counter reset interval
- Size: 2 bytes
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 4
- Name: Threshold value for discarded blocks due to buffer overflow
- Size: 4 bytes
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 5
- Name: Traffic scheduler pointer
- Size: 2 bytes
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 6
- Name: Weight
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 7
- Name: Back pressure operation
- Size: 2 bytes
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 8
- Name: Back pressure time
- Size: 4 bytes
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 9
- Name: Back pressure occur queue threshold
- Size: 2 bytes
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 10
- Name: Back pressure clear queue threshold
- Size: 2 bytes
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 11
- Name: Packet drop queue thresholds
- Size: 8 bytes
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 12
- Name: Packet drop max_p
- Size: 2 bytes
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 13
- Name: Queue drop w_q
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: optional

## Extraction Status
- Status: auto_extracted
- Attributes found: 13
- Review needed: false

## Raw Source

```
9.2.10 Priority queue  NOTE 1 – In [ITU-T G.984.4], this is called a priority queue-G.  This ME specifies the priority queue used by a GEM port network CTP  in the upstream direction.  The upstream priority queue ME is also related to a T -CONT ME. By default, this relationship is  fixed by the ONU hardware architecture, but some ONUs may also permit the relationship to be  configured through the OMCI, as indicated by the QoS configuration flexibility attribute of the  ONU2-G ME. 

---- page break ---- In the downstream direction, priority queues are associated with UNIs. Again, the association is fixed  by default, but some ONUs may permit the association to be configured through the OMCI.  If an ONU as a whole contains priority queues, it instantiates these queues autonomously. Priority  queues may also be localized to pluggable circuit packs, in which case the ONU creates and deletes  them in accordance with circuit pack pre-provisioning and the equipped configuration.  The OLT can find all the queues by reading the priority queue ME instances. If the OLT tries to  retrieve a non-existent priority queue, the ONU denies the get action with an error indication.  See also Appendix II.  Priority queues can exist in the ONU core  and circuit packs serving both UNI and ANI functions.  Therefore, they can be indirectly created and destroyed through cardholder provisioning actions.  In the upstream direction, the weight attribute permits the configuring of an optional traffic scheduler.  Several attributes support back pressure operation, whereby a back-pressure signal is sent backwards  and causes the attached terminal to temporarily suspend sending data.  In the downstream direction, strict priority discipline among the queues serving a given UNI is the  default, with priorities established through the related port attribute. If two or more non-empty queues  have the same priority, capacity is allocated among them in proportion to their weights. Note that the  details of the downstream model differ from those of the upstream model.  The yellow packet drop thresholds specify the drop probability for a packet that has been marked  yellow (drop eligible) by a traffic descriptor or by external equipment such as a residential gateway   (RG). If the current average queue occupancy is less than the minimum threshold, the yellow packet  drop probability is zero. If the current average queue occupancy is greater than or equal to the  maximum threshold, the yellow packet drop probability is one. The yellow drop probability increases  linearly between 0 and max_p as the current average queue occupancy increases from the minimum  to the maximum threshold.  The same model can be configured for green packets, those regarded as being within the traffic  contract.  Drop precedence colour marking indicates the method by which a packet is marked as drop eligible  (yellow). For discard eligibility indicator (DEI) and priority code point (PCP) marking, a drop eligible  indicator is equivalent to yellow colour; otherwise, the colour is green. For  differentiated services  code point ( DSCP) assured forwarding (AF) marking, the lowest drop precedence is equivalent to  green; otherwise, the colour is yellow.  Relationships  One or more instances of this ME are associated with the ONU -G ME to model upstream  priority queues if the traffic management option attribute in the ONU-G ME is 0 or 2.  One or more instances of this ME are associated with a PPTP UNI ME as downstream priority  queues. Downstream priority queues may or may not be provided for a virtual Ethernet  interface point (VEIP).  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. The MSB  represents the direction (1: upstream, 0: downstream). The 15 LSBs represent  a queue ID. The queue ID is numbered in ascending order by the ONU itself.  It is strongly encouraged that the queue ID be formulated to simplify finding  related queues. One way to do this is to number the queues such that the related  port attributes are in ascending order (for the downstream and upstream queues  separately). The range of downstream queue ids is 0 to 0x7FFF and the range  of upstream queue ids is 0x8000 to 0xFFFF. (R) (mandatory) (2 bytes) 

---- page break ---- Queue configuration option : This attribute identifies the buffer partitioning policy. The  value 1 means that several queues share one buffer of maximum queue size,  while the value 0 means that each queue has an individual buffer of maximum  queue size. (R) (mandatory) (1 byte)  Maximum queue size : This attribute specifies the maximum size of the queue, in bytes,  scaled by the priority queue scale factor attribute of the ONU2 -G. (R)  (mandatory) (2 bytes)  NOTE 2 – In this and the other similar attributes of the priority queue ME, some  legacy implementations may take the queue scale factor from the GEM block length  attribute of the ANI-G ME. This option is discouraged in new implementations.  Allocated queue size: This attribute identifies the allocated size of this queue, in bytes, scaled  by the priority queue scale factor attribute of the ONU2-G. (R, W) (mandatory)  (2 bytes)  Discard-block counter reset interval: This attribute represents the interval in milliseconds  at which the counter resets itself. (R, W) (optional) (2 bytes)  Threshold value for discarded blocks due to buffer overflow : This attribute specifies the  threshold for the number of bytes (scaled by the priority queue scale factor  attribute of the ONU2 -G) discarded on this queue due to buffer overflow. Its  value controls the declaration of the block loss alarm. (R, W) (optiona l)  (2 bytes)  Related port: This attribute represents the slot, port/T -CONT and priority information  associated with the instance of priority queue ME. This attribute comprises  4 bytes.  In the upstream direction, the first 2 bytes are the ME ID of the associated T - CONT, the first byte of which is a slot number, th
```
