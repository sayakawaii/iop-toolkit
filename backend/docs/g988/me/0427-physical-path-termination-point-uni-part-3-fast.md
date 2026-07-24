# Managed Entity

## Identity
- ME ID: 427
- ME Name: Physical path termination point UNI part 3 (FAST)
- Source Section: 9.7.48
- Source Page: 332

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Get, Set
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: FAST line configuration profile
- Size: 2 bytes
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 2
- Name: FAST data path configuration profile
- Size: 2 bytes
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 3
- Name: FAST channel configuration profile for bearer channel 0 downstream
- Size: 2 bytes
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 4
- Name: FAST channel configuration profile for bearer channel 0 upstream
- Size: 2 bytes
- Format: needs_review
- Access: R, W
- Category: optional

## Extraction Status
- Status: auto_extracted
- Attributes found: 4
- Review needed: false

## Raw Source

```
9.7.48 Physical path termination point UNI part 3 (FAST)  This ME represents the point in the ONU where physical paths terminate on a FAST [ITU-T G.9700,  ITU-T G.9701] DPU modem (FTU-O).   The ONU creates or deletes an instance of this ME at the same time it creates or deletes the  corresponding PPTP xDSL UNI part 1 intended to serve a FAST connection, as per [ITU-T G.997.2].  Relationships  An instance of this ME is associated with each instance of a real or preprovisioned FAST port  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. Through an  identical ID, this ME is implicitly linked to an instance of a PPTP xDSL UNI  part 1. (R) (mandatory) (2 bytes)  FAST line configuration profile : This attribute points to an instance of the FAST line  configuration profiles (part 1, 2, 3 and 4) MEs, also to FAST vectoring line  configuration extension MEs. Upon ME instantiation, the ONU sets this  attribute to 0, a null pointer. (R, W) (mandatory) (2 bytes)  FAST data path configuration profile: This attribute points to an instance of the FAST data  configuration profile that defines data path parameters. Upon ME instantiation,  the ONU sets this attribute to 0, a null pointer. (R, W) (optional) (2 bytes)  FAST channel configuration profile for bearer channel 0 downstream : This attribute  points to an instance of the FAST channel configuration profile that defines  channel parameters. Upon ME instantiation, the ONU sets this attribute to 0, a  null pointer. (R, W) (optional) (2 bytes) (R, W) (optional) (2 bytes)  FAST channel configuration profile for bearer channel 0 upstream: This attribute points  to an instance of the FAST channel configuration profile that defines channel  parameters. Upon ME instantiation, the ONU sets this attribute to 0, a null  pointer (R, W) (optional) (2 bytes)  Actions  Get, set 

---- page break ---- Notifications  None.  
```
