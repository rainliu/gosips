/**
 * ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
 * Module Name   : GoSIP Specification
 * File Name     : Router.go
 * Author        : Rain Liu
 *~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
 */

//TODO move back to address package

package message

import (
	"container/list"
	"gosips/sip/address"
)

/**
 * The Router interface defines accessor methods to retrieve the default Route
 * and Outbound Proxy of this SipStack. The Outbound Proxy and default Route are
 * made up one or more {@link Hop}'s. This Router information is user-defined in
 * an object that implements this interface. The location of the user-defined
 * Router object is supplied in the Properties object passed to the
 * {@link javax.sip.SipFactory#createSipStack(Properties)} method upon creation
 * of the SIP Stack object.
 * The Router object must accept a SipStack as an argument to the constructor in
 * order for the Router to access attributes of the SipStack such as IP Address.
 * The constructor of an object implementing the Router interface must be
 * <code>RouterImpl(SipStack sipStack, String outboundProxy) {}</code>
 * <p>
 * The user may define a routing policy dependent on the operation of the
 * SipStack i.e. user agent or proxy, however this routing policy can not be
 * changed dynamically, i.e. the SipStack needs to be deleted and re-created.
 *
 * @author Sun Microsystems
 * @since 1.1
 */

type Router interface {

	/**
	 * Gets the Outbound Proxy parameter of this Router, this method may return
	 * null if no outbound proxy is defined.
	 *
	 * @return the Outbound Proxy of this Router.
	 * @see Hop
	 */
	GetOutboundProxy() address.Hop

	/**
	 * Gets the ListIterator of the hops of the default Route. This method may
	 * return null if a default route is not defined.
	 *
	 * @param request - the Request message that determines the default route.
	 * @return the ListIterator over all the hops of this Router.
	 * @see Hop
	 */
	GetNextHops(request Request) *list.List //ListIterator

}
