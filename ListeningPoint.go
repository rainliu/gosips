/**
 * ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
 * Module Name   : GoSIP Specification
 * File Name     : ListeningPoint.go
 * Author        : Rain Liu 
 *~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
 */

package gosip

import (

)


/**
 * This interface represents a unique IP network listening point,
 * which consists of port and transport. A ListeningPoint is a Java
 * representation of the port that a SipProvider messaging entity uses to send
 * and receive messages.
 * <p>
 * For any address and port that a server listens on for UDP, it MUST listen on 
 * that same port and address for TCP.  This is because a message may need to 
 * be sent using TCP, rather than UDP, if it is too large. To handle this  
 * a Listening point with the same port but with TCP transport would be 
 * created and attached to a new SipProvider, upon which the SipListener is 
 * registered. However the converse is not true, a server need not listen for 
 * UDP on a particular address and port just because it is listening on that 
 * same address and port for TCP.
 * <p>
 * ListeningPoints can be created from the 
 * {@link SipStack#createListeningPoint(int, String)}. A SipStack object may 
 * have multiple ListeningPoints, while a SipProvider as a messaging entity 
 * may only have a single ListeningPoint.
 *
 * @author Sun Microsystems
 * @version 1.1
 */

//Constants
    
    /**
     * Transport constant: TCP
     */
    const TCP = "TCP";

    /**
     * Transport constant: UDP
     */
    const UDP = "UDP";

    /**
     * Transport constant: SCTP
     *
     * @since v1.1
     */
    const SCTP = "SCTP";

    /**
     * Transport constant: TLS over TCP 
     *
     * @since v1.1
     */
    const TLS = "TLS";

    /**
     * Port Constant: Default port 5060. This constant should only be used
     * when the transport of the ListeningPoint is set to UDP, TCP or SCTP.
     *
     * @since v1.1
     */
    const PORT_5060 = 5060;

    /**
     * Port Constant: Default port 5061. This constant should only be used
     * when the transport of the Listening Point is set to TLS over TCP or TCP 
     * assuming the scheme is "sips".
     *
     * @since v1.1
     */
    const PORT_5061 = 5061;


type ListeningPoint interface {// extends Cloneable, Serializable {

    /**
     * Gets the port of the ListeningPoint. The default port of a ListeningPoint
     * is dependent on the scheme and transport.  For example:
     * <ul>
     * <li>The default port is 5060, if the transport UDP and the scheme is 
     * <i>sip:</i>.
     * <li>The default port is 5060, if the transport is TCP and the scheme 
     * is <i>sip:</i>.
     * <li>The default port is 5060, if the transport is SCTP and the scheme 
     * is <i>sip:</i>.
     * <li>The default port is 5061, if the transport is TLS over TCP and the 
     * scheme is <i>sip:</i>.
     * <li>The default port is 5061, if the transport is TCP and the scheme 
     * is <i>sips:</i>.
     * </ul>
     *
     * @return the integer value of the port.
     */
    GetPort() int;

    /**
     * Gets the transport of the ListeningPoint.
     *
     * @return the string value of the transport.
     */
    GetTransport() string;

    /**
     * This method indicates whether the specified object is equal to this 
     * Listening Point. The specified object is equal to this ListeningPoint 
     * if the specified object is a ListeningPoint and the transport and port 
     * in the specified Listening Point is the same as this Listening Point.
     *
     * @param obj - the object with which to compare this ListeningPoint.
     * @return true if this ListeningPoint is "equal to" the obj argument;
     * false otherwise.
     */
    Equals(obj interface{}) bool;
}

