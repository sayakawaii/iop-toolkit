# Managed Entity

## Identity
- ME ID: 343
- ME Name: Energy consumption performance monitoring history data
- Source Section: 9.2.14
- Source Page: 121

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
- Name: Doze time
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 3
- Name: Cyclic sleep time
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 4
- Name: Watchful sleep time
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 5
- Name: Energy consumed
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: optional

## Extraction Status
- Status: auto_extracted
- Attributes found: 5
- Review needed: false

## Raw Source

```
9.2.14 Energy consumption performance monitoring history data  This ME collects PM data associated with the ONU's energy consumption. The time spent in various  low-power states is recorded as a measure of their utility. Furthermore, the ONU may also include  the equivalent of a watt-hour meter, which can be sampled from time to time to measure actual power  consumed.   For a complete discussion of generic PM architecture, refer to clause I.4.  Relationships  An instance of this ME is associated with the ONU in its entirety.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. The ME ID  must be 0. (R, set-by-create) (mandatory) (2 bytes)  Interval end time : This attribute identifies the most recently finished 15  min interval. (R)  (mandatory) (1 byte) 

---- page break ---- Threshold data  1/2 ID: No thresholds are defined for this ME. For uniformity with other  PMs, the attribute is retained and shown as mandatory, but it should be set to  a null pointer. (R, W, set-by-create) (mandatory) (2 bytes)  Doze time : This attribute records the time during which the ONU was in doze energy  conservation mode, measured in microseconds. If watchful sleep is enabled in  the ONU dynamic power management control ME, the ONU ignores this  attribute. (R) (mandatory) (4 bytes)  Cyclic sleep time: This attribute records the time during which the ONU was in cyclic sleep  energy conservation mode, measured in microseconds. If watchful sleep is  enabled in the ONU dynamic power management control ME, the ONU  ignores this attribute. (R) (mandatory) (4 bytes)  Watchful sleep time: This attribute records the time during which the ONU was in watchful  sleep energy conservation mode, measured in microseconds. (R) (mandatory)  (4 bytes)  Energy consumed: This attribute records the energy consumed by the ONU, measured in  millijoules. (R) (optional) (4 bytes)  Actions  Create, delete, get, set  Get current data (optional)  Notifications  None.  
```
