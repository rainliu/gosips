GoSIP
=====

SIP stack in Go

Objective
---------
The objective of GoSIP is to develop a stack interface to the Session Initiation Protocol (SIP) that can be used independently or by higher level programming entities and environments. GoSIP was designed to provide a developer with a standardized interface for SIP services which are functionally compatible with the RFC3261 specification. This specification is a general purpose transaction based Java interface to the SIP protocol. It is rich both semantically and in definition to the SIP protocol. 

Design approach
---------------
GoSIP supports RFC 3261 functionality and the following SIP extensions; the INFO method (RFC 2976), Reliability of provisional responses (RFC 3262), Event Notification Framework (RFC 3265), the UPDATE method (RFC 3311), the Reason Header (RFC 3326), the Message method (RFC 3428) defined for instant messaging and the REFER method (RFC 3515).

GoSIP standardizes the interface to the generic transactional model defined by the SIP protocol, providing access to dialog functionality from the transaction interface. The architecture is developed for the J2SE environment therefore is event based utilizing the Listener/Provider event model. The specification is asynchronous in nature using transactional identifiers to correlate messages. It defines various factory classes for creating Request and Response messages and SIP headers.  JAIN SIP defines an interface for each Header supported, which can be added to Request or Response messages respectively. These messages are passed to the SipProvider with a transaction to be sent onto the network, while the SipListener listens for incoming Events that encapsulate messages that may be responses to initiated dialogs or new incoming dialogs.

GoSIP is extensible by design. It defines a generic extension header interface that can be used by applications that utilize headers that are not supported directly by GoSIP. The design also defines a mechanism to support future dialog creation methods in a GoSIP environment. GoSIP can be managed statically with regards to IP addresses and router function, and dynamically specific to ports and transports. 

The default handling of message retransmissions in GoSIP is dependent on the application. Stateful proxy applications need not be concerned with retransmissions as these are handled by GoSIP. Typically User Agent applications must handle retransmissions of ACK’s and 2xx Responses, however GoSIP provides a convenience function that ensures all retransmissions are handled by the GoSIP implementation, reducing the complexity for applications acting as user agents.
