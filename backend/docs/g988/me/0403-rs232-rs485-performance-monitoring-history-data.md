# Managed Entity

## Identity
- ME ID: 403
- ME Name: RS232/RS485 performance monitoring history data
- Source Section: 9.15.3
- Source Page: 485

## Classification
- Access: needs_review
- Type: needs_review
- Actions: needs_review
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
- Name: Threshold data 1/2 id
- Size: 2 bytes
- Format: needs_review
- Access: R, W, set-by-create
- Category: mandatory

### Attribute 3
- Name: Outgoing bytes from PON port
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 4
- Name: Incoming bytes fr om RS232/RS485 controller chipset
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 5
- Name: Outgoing bytes f rom RS232/RS485 controller chipset
- Size: 4 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 6
- Name: Total TWDM channel number
- Size: 1 byte
- Format: needs_review
- Access: R
- Category: mandatory

## Extraction Status
- Status: auto_extracted
- Attributes found: 6
- Review needed: true

## Raw Source

```
9.15.3 RS232/RS485 performance monitoring history data  This ME collects PM data for a RS232/RS485 interface. Instances of this ME are created and deleted  by the OLT.  Relationships  An instance of this ME is associated with an instance of the PPTP RS232/RS485 UNI ME.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. Through an  identical ID, this ME is implicitly linked to an instance of the PPTP  RS232/RS485 UNI. (R, set-by-create) (mandatory) (2 bytes)  Interval end time : This attribute identifies the most recently finished 15  min interval. (R)  (mandatory) (1 byte)  Threshold data 1/2 id: This attribute points to an instance of the threshold data 1 and 2 MEs  that contains PM threshold values. (R, W, set-by-create) (mandatory) (2 bytes) 

---- page break ---- Incoming bytes from PON port: This attribute counts the bytes received on the PON port.   (R) (optional) (4 bytes)  Outgoing bytes from PON port: This attribute counts the bytes transmitted on the PON port.  (R) (optional) (4 bytes)  Incoming bytes fr om RS232/RS485 controller chipset : This attribute counts the bytes  received on the RS232/RS485 chipset.  (R) (optional)  (4 bytes)  Outgoing bytes f rom RS232/RS485 controller chipset : This attribute counts the bytes  transmitted on the RS232/RS485 chipset. (R) (optional)  (4 bytes)  Actions  Create, delete, get, set  Notifications  Threshold crossing alert  Number Threshold crossing alert Threshold value attribute No.  (Note)  1 Incoming packets 1  2 Incoming bits 2  3 Outgoing packets 3  4 Outgoing bits 4  NOTE – This number associates the TCA with the specified threshold value  attribute of the threshold data 1/2 managed entities.  9.16 ITU-T G.989.3 TWDM and ITU-T G.9804.2 TWDM/TDM PON  This clause defines MEs associated with TWDM/TDM PON management.  9.16.1 TWDM/TDM System Profile managed entity  This ME models the TWDM subsystem of NG-PON2 system and TWDM/TDM subsystem of ITU-T  G.9804.2 based PON system. An instance of this ME corresponds to a physical or virtual slot of the  ONU housing one or more access network interfaces. The instances of this ME are instantiated  autonomously by the ONU.  Relationships  An instance of this ME is associated with an instance of a circuit pack that supports a PON  interface function. It is, therefore, implicitly associated with all ANI -G MEs whose ME ID  refers the specific Slot ID.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. This 2 byte  number is represented as 0xSS 00, where SS indicates the slot ID (as defined  in clause 9.1.5 and referenced in clause 9.2.1 of G.988.  (R) (mandatory)  (2 bytes)  Total TWDM channel number : This attribute indicates the number of distinct TWDM  channels the optical network termination  (ONT) supports in given slot. (R)  (mandatory) (1 byte)
```
