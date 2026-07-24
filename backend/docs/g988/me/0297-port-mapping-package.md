# Managed Entity

## Identity
- ME ID: 297
- ME Name: Port-mapping package
- Source Section: 9.1.8
- Source Page: 82

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Get
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: Max ports
- Size: 1 byte
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 2
- Name: Port list 1
- Size: 16 bytes
- Format: needs_review
- Access: R
- Category: mandatory

### Attribute 3
- Name: Port list 2
- Size: 16 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 4
- Name: Port list 3
- Size: 16 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 5
- Name: Port list 4
- Size: 16 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 6
- Name: Port list 5
- Size: 16 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 7
- Name: Port list 6
- Size: 16 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 8
- Name: Port list 7
- Size: 16 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 9
- Name: Port list 8
- Size: 16 bytes
- Format: needs_review
- Access: R
- Category: optional

## Extraction Status
- Status: auto_extracted
- Attributes found: 9
- Review needed: false

## Raw Source

```
9.1.8 Port-mapping package  NOTE – In [ITU-T G.984.4], this ME is called a port-mapping package-G.  This ME provides a way to map a heterogeneous set of PPTPs (ports) to a parent equipment, which  may be a cardholder or the ONU itself. It could be useful, for example, if a single plug-in circuit pack  contained a PON ANI as port 1, a video UNI as port 2, and a craft UNI as port 3. Another application 

---- page break ---- of the port-mapping package is the case where more than one UNI or ANI ME is associated with a  single physical port, for example, the RE ANI and downstream amplifier. This ME also provides an  option for an integrated ONU to represent its ports without the use of virtual cardholders and VC  packs.  If the port-mapping package is supported for the ONU as a whole, it is automatically created by the  ONU. If the port-mapping package is supported for plug-in circuit packs, it is created and destroyed  by the ONU when the corresponding circuit pack is installed or pre-provisioned in a cardholder.  The port list attributes specify ports 1..64 sequentially. Each port list is a sequence of ME types, as  defined in Table 11.2.4-1. These ME type codes define what kind of PPTP or ANI corresponds to the  specific port number. For example, for a circuit pack with 4 POTS ports, 2 xDSL ports, and 1 video  UNI port, numbered sequentially in that order, the attributes would be coded:   Max ports:  7   Port list 1  53, 53, 53, 53, 98, 98, 82, 0   Port list 2..8  All zero  Relationships  A port-mapping package may be contained by an ONU-G or a cardholder.  Attributes  Managed entity ID: This attribute uniquely identifies each instance of this ME. Through an  identical ID, this ME is implicitly linked to an instance of the ONU -G or  cardholder. (R) (mandatory) (2 bytes)  Max ports: This attribute indicates the largest port number contained in the port list  attributes. Ports are numbered from 1 to this maximum, possibly with  embedded 0 entries, but no port may exist beyond the maximum. (R)  (mandatory) (1 byte)  Each of the following attributes is a list of 8 ports, in increasing port number sequence. Each list entry  is a 2 byte field containing the ME type of the UNI or ANI corresponding to the port number. ME  types are defined in Table 11.2.4-1. Placeholders for non-existent port numbers are indicated with the  value 0.  Port list 1: (R) (mandatory) (16 bytes)  Port list 2: (R) (optional) (16 bytes)  Port list 3: (R) (optional) (16 bytes)  Port list 4: (R) (optional) (16 bytes)  Port list 5: (R) (optional) (16 bytes)  Port list 6: (R) (optional) (16 bytes)  Port list 7: (R) (optional) (16 bytes)  Port list 8: (R) (optional) (16 bytes)  Combined port table: This attribute permits the implicit linking of multiple port -level MEs  to a single physical port. For example, a single physical port may be linked to  both an RE ANI -G and an RE downstream amplifier, as illustrated in   Figure 8.2.10-4. The combination of RE ANI-G and RE downstream amplifier  cannot be directly represented in the port list attributes.  Each row of the combined port table comprises the following fields.   

---- page break ---- Field name Size,  bytes Description  Physical port 1 Duplicates are allowed.  The corresponding physical port in the port list  attribute should be 0.  Equipment type 1 2 ME type 1, from Table 11.2.4-1. The first  equipment type in the list is understood to be the  master (in the first row of the table, in the case  of duplicate physical ports). The administrative  state, operational state and ARC attributes of the  master override the corresponding attributes of  secondary MEs.  … 2 … secondary MEs that share the physical port  Equipment type 12 2 Secondary ME type 12  As many as 12 ME types can be associated with a given physical port, and  even more by duplicating the physical port field. (R) (optional) ( N rows  * 25 bytes)  Actions  Get, get next  Notifications  None.  
```
