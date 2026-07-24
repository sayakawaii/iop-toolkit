# Managed Entity

## Identity
- ME ID: 290
- ME Name: Dot1X port extension package
- Source Section: 9.3.14
- Source Page: 173

## Classification
- Access: needs_review
- Type: needs_review
- Actions: Get, Set
- Instance Type: per-instance
- Instance Value: 2 bytes

## Attributes

### Attribute 1
- Name: Dot1x enable
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: mandatory

### Attribute 2
- Name: Action register
- Size: 1 byte
- Format: needs_review
- Access: W
- Category: mandatory

### Attribute 3
- Name: Authenticator PAE state
- Size: 1 byte
- Format: needs_review
- Access: R
- Category: optional

### Attribute 4
- Name: Admin controlled directions
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 5
- Name: Operational controlled directions
- Size: 1 byte
- Format: needs_review
- Access: R
- Category: optional

### Attribute 6
- Name: Authenticator controlled port status
- Size: 1 byte
- Format: needs_review
- Access: R
- Category: optional

### Attribute 7
- Name: Quiet period
- Size: 2 bytes
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 8
- Name: Server timeout period
- Size: 2 bytes
- Format: needs_review
- Access: R, W
- Category: optional

### Attribute 9
- Name: Re-authentication period
- Size: 2 bytes
- Format: needs_review
- Access: R
- Category: optional

### Attribute 10
- Name: Re-authentication enabled
- Size: 1 byte
- Format: needs_review
- Access: R
- Category: optional

### Attribute 11
- Name: Key transmission enabled
- Size: 1 byte
- Format: needs_review
- Access: R, W
- Category: optional

## Extraction Status
- Status: auto_extracted
- Attributes found: 11
- Review needed: false

## Raw Source

```
9.3.14 Dot1X port extension package  An instance of this ME represents a set of attributes that control a port's IEEE 802.1X operation. It is  created and deleted autonomously by the ONU upon the creation or deletion of a PPTP that supports  [IEEE 802.1X] authentication of customer premises equipment (CPE).  Relationships  An instance of this ME is associated with a PPTP that performs IEEE 802.1X authentication  of CPE (e.g., Ethernet or DSL).  Attributes  Managed entity ID: This attribute provides a unique number for each instance of this ME.  Its value is the same as that of its associated PPTP (i.e., slot and port number).  (R) (mandatory) (2 bytes)  Dot1x enable: If true, this Boolean attribute forces the associated port to authenticate via  [IEEE 802.1X] as a precondition of normal service. The default value false  does not impose IEEE 802.1X authentication on the associated port. (R,  W)  (mandatory) (1 byte)  Action register: This attribute defines a set of actions that can be performed on the associated  port. The act of writing to the register causes the specified action.  1 Force re -authentication – this opcode initiates an IEEE  802.1X  re-authentication conversation with the associated port. The port  remains in its current authorization state until the conversation  concludes.  2 Force unauthenticated – this opcode initiates an IEEE  802.1X  authentication conversation whose outcome is predestined to fail,  thereby disabling normal Ethernet service on the port. The port 's  provisioning is not changed, such that upon re -initialization, a new  IEEE 802.1X conversation may restore service without prejudice.  3 Force authenticated – this opcode initiates an IEEE  802.1X  authentication conversation whose outcome is predestined to succeed,  thereby unconditionally enabling normal Ethernet service on the port.  The port's provisioning is not changed, such that upon re-initialization,  a new IEEE 802.1X conversation is required.  (W) (mandatory) (1 byte)  Authenticator PAE state: This attribute returns the value of the port 's PAE state. States are  further described in [IEEE 802.1X]. Values are coded as follows.  0 Initialize  1 Disconnected  2 Connecting  3 Authenticating  4 Authenticated  5 Aborting  6 Held  7 Force auth  8 Force unauth  9 Restart  (R) (optional) (1 byte) 

---- page break ---- Backend authentication state : This attribute returns the value of the port 's back -end  authentication state. States are further described in [IEEE 802.1X]. Values are  coded as follows.  0 Request  1 Response  2 Success  3 Fail  4 Timeout  5 Idle  6 Initialize  7 Ignore  (R) (optional) (1 byte)  Admin controlled directions : This attribute controls the directionality of the port 's  authentication requirement. The default value 0 indicates that control is  imposed in both directions. The value 1 indicates that control is imposed only  on traffic from the subscriber towards the network. (R, W) (optional) (1 byte)  Operational controlled directions : This attribute indicates the directionality of the port 's  current authentication state. The value 0 indicates that control is imposed in  both directions. The value 1 indicates that control is imposed only on traffic  from the subscriber towards the network. (R) (optional) (1 byte)  Authenticator controlled port status: This attribute indicates whether the controlled port is  currently authorized (1) or unauthorized (2). (R) (optional) (1 byte)  Quiet period: This attribute specifies the interval between EAP request/identity invitations  sent to the peer. Other events such as carrier present or EAPOL start frames  from the peer may trigger an EAP request/identity frame from the ONU at any  time; this attribute c ontrols the ONU 's periodic behaviour in the absence of  these other inputs. It is expressed in seconds. (R, W) (optional) (2 bytes)  Server timeout period : This attribute specifies the time the ONU will wait for a response  from the radius server before timing out. Within this maximum interval, the  ONU may initiate several retransmissions with exponentially increasing delay.  Upon timeout, the ONU may try another radius server if there is one, or invoke  the fallback policy, if no alternate radius servers are available. Server timeout  is expressed in seconds, with a default value of 30 and a maximum value of  65535. (R, W) (optional) (2 bytes)  Re-authentication period: This attribute records the re -authentication interval specified by  the radius authentication server. It is expressed in seconds. The attribute is only  meaningful after a port has been authenticated. (R) (optional) (2 bytes)  Re-authentication enabled: This Boolean attribute records whether the radius authentication  server has enabled re -authentication on this service (true) or not (false). The  attribute is only meaningful after a port has been authenticated. (R) (optional)  (1 byte)  Key transmission enabled : This Boolean attribute indicates whether key transmission is  enabled (true) or not (false). This feature is not required; the parameter is listed  here for completeness vis-à-vis [IEEE 802.1X]. (R, W) (optional) (1 byte)  Actions  Get, set 

---- page break ---- Notifications  Alarm  Alarm  number Alarm Description  0 dot1x local  authentication – allowed  No radius authentication server was accessible. In  accordance with local policy, the port was allowed access  without authentication.  1 dot1x local  authentication – denied  No radius authentication server was accessible. In  accordance with local policy, the port was denied access.  2..207 Reserved   208..223 Vendor-specific alarms Not to be standardized  
```
