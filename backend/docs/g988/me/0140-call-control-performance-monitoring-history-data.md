# Managed Entity

## Identity
- ME ID: 140
- ME Name: Call control performance monitoring history data
- Source Section: 9.9.12
- Source Page: 405

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Create, Delete, Get, Set
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: Interval end time
- Size: 1 byte
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 2
- Name: Threshold data 1/2 ID
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 3
- Name: Call setup failures
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 4
- Name: Call setup timer
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 5
- Name: Call terminate failures
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 6
- Name: Analog port off -hook timer
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

## Extraction Status
- Status: auto_extracted
- Attributes found: 6
- Review needed: false

## Raw Source

```
9.9.12 Call control performance monitoring history data  This ME collects PM data related to the call control channel. Instances of this ME are created and  deleted by the OLT.  For a complete discussion of generic PM architecture, refer to clause I.4.  Relationships  An instance of this ME is associated with an instance of the PPTP POTS UNI ME.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. Through an  identical ID, this ME is implicitly linked to an instance of the PPTP POTS  UNI. (R, set-by-create) (mandatory) (2 bytes)  Interval end time : This attribute identifies the most recently finished 15  min interval. (R)  (mandatory) (1 byte)  Threshold data 1/2 ID: This attribute points to an instance of the threshold data  1 ME that  contains PM threshold values. Since no threshold value attribute number  exceeds 7, a threshold data 2 ME is optional. (R, W, set-by-create) (mandatory)  (2 bytes)  Call setup failures: This attribute counts call set-up failures. (R) (mandatory) (4 bytes)  Call setup timer : This attribute is a high water -mark that records the longest duration of a  single call set -up detected during this interval. Time is measured in  milliseconds from the time an initial set -up was requested by the subscriber  until the time at which a response was provided to the subscriber in the form  of busy tone, audible ring tone, etc. (R) (mandatory) (4 bytes)  Call terminate failures: This attribute counts the number of calls that were terminated with  cause. (R) (mandatory) (4 bytes) 

---- page break ---- Analog port releases : This attribute counts the number of analogue port releases without  dialling detected (abandoned calls). (R) (mandatory) (4 bytes)  Analog port off -hook timer: This attribute is a high water -mark that records the longest  period of a single off-hook detected on the analogue port. Time is measured in  milliseconds. (R) (mandatory) (4 bytes)  Actions  Create, delete, get, set  Get current data (optional)  Notifications  Threshold crossing alert  Alarm  number Threshold crossing alert Threshold value attribute No. (Note)  0 CCPM call set-up fail 1  1 CCPM set-up timeout 2  2 CCPM call terminate 3  3 CCPM port release with no  dialling  4  4 CCPM port offhook timeout 5  NOTE – This number associates the TCA with the specified threshold value attribute of the  threshold data 1 managed entity.  
```
