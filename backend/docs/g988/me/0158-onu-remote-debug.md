# Managed Entity

## Identity
- ME ID: 158
- ME Name: ONU remote debug
- Source Section: 9.1.12
- Source Page: 88

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Get, Set
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: Command format
- Size: 1 byte
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 2
- Name: Command
- Size: 25 bytes
- Format: needs_review
- Access: W
- Category: mandatory

## Extraction Status
- Status: auto_extracted
- Attributes found: 2
- Review needed: false

## Raw Source

```
9.1.12 ONU remote debug  This ME is used to send vendor -specific debug commands to the ONU and receive vendor -specific  replies back for processing on the OLT. This allows for the remote debugging of an ONU that may  not be accessible by other means. The command format may have two modes, one being text and the  other free format. In text format, both the command and reply are ASCII strings, but are otherwise  unconstrained. In free format, the content and format of command and reply are vendor-specific.  An ONU that supports remote debugging automatically creates an instance of this ME. It is not  reported during an MIB upload.  Relationships  One instance of this ME is associated with the ONU ME.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. There is only  one instance, number 0. (R) (mandatory) (2 bytes)  Command format: This attribute defines the format of the command and reply attributes.  The value 0 defines ASCII string format, while 1 specifies free format. (R)  (mandatory) (1 byte)  Command: This attribute is used to send a command to the ONU. The format of the  command is defined by the command format. If the format is ASCII string, the  command should be null terminated unless the string is exactly 25 bytes long.  The action of setting this attribute should trigger the ONU to discard any  previous command reply information and execute the current debugging  command. (W) (mandatory) (25 bytes)  Reply table: This attribute is used to pass reply information back to the OLT. Its format is  defined by the command format attribute. The get, get next action sequence  must be used with this attribute, since its size is unspecified. (R) (mandatory)  (N bytes)  Actions  Get, get next, set 

---- page break ---- Notifications  None.  
```
