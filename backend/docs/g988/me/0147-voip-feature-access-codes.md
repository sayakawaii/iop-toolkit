# Managed Entity

## Identity
- ME ID: 147
- ME Name: VoIP feature access codes
- Source Section: 9.9.9
- Source Page: 401

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Create, Delete, Get, Set
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: Cancel call waiting
- Size: 5 bytes
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 2
- Name: Call hold
- Size: 5 bytes
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 3
- Name: Call park
- Size: 5 bytes
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 4
- Name: Caller ID activate
- Size: 5 bytes
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 5
- Name: Caller ID deactivate
- Size: 5 bytes
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 6
- Name: Do not disturb activation
- Size: 5 bytes
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 7
- Name: Do not disturb deactivation
- Size: 5 bytes
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 8
- Name: Do not disturb PIN change
- Size: 5 bytes
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 9
- Name: Emergency service number
- Size: 5 bytes
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 10
- Name: Intercom service
- Size: 5 bytes
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 11
- Name: Unattended/blind call transfer
- Size: 5 bytes
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 12
- Name: Attended call transfer
- Size: 5 bytes
- Format: needs_review
- Access: R, W
- Category: optional

## Extraction Status
- Status: auto_extracted
- Attributes found: 12
- Review needed: false

## Raw Source

```
9.9.9 VoIP feature access codes  The VoIP feature access codes ME defines administrable feature access codes for the VoIP  subscriber. It is optional for ONUs that support VoIP service s. If a non -OMCI interface is used to  manage VoIP signalling, this ME is unnecessary.  Instances of this ME are created and deleted by the OLT. A VoIP feature access codes instance is  needed for each unique set of feature access code attributes.  Relationships  An instance of this ME may be associated with one or more SIP user data MEs.  Attributes  Managed entity  ID: This attribute uniquely identifies each instance of this ME. (R)  (mandatory) (2 bytes)  The remaining attributes are access codes for the features mentioned in their names. Each attribute is  a string of characters from the set {0..9, *, #}, with trailing nulls in any unused bytes.  Cancel call waiting: (R, W) (optional) (5 bytes)  Call hold:  (R, W) (optional) (5 bytes)  Call park:  (R, W) (optional) (5 bytes)  Caller ID activate: (R, W) (optional) (5 bytes)  Caller ID deactivate: (R, W) (optional) (5 bytes)  Do not disturb activation: (R, W) (optional) (5 bytes)  Do not disturb deactivation: (R, W) (optional) (5 bytes)  Do not disturb PIN change: (R, W) (optional) (5 bytes)  Emergency service number: (R, W) (optional) (5 bytes)  Intercom service: (R, W) (optional) (5 bytes)  Unattended/blind call transfer: (R, W) (optional) (5 bytes)  Attended call transfer: (R, W) (optional) (5 bytes)  Actions  Create, delete, get, set 

---- page break ---- Notifications  None.  
```
