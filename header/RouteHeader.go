package header

/**
 * The Route header field is used to force routing for a request through the
 * listed set of proxies. Each host removes the first entry and then proxies
 * the Request to the host listed in that entry using it as the RequestURI.
 * <p>
 * Explicit Route assignment (if needed) for the initial dialog establishment
 * is the applications responsibility, but once established Routes are
 * maintained by the dialog layer and should not be manupulated by the
 * application. For example the SipProvider queries the dialog for Route
 * assignment and adds these to the outgoing message as needed. The
 * {@link javax.sip.address.Router} may be used by the application to determine
 * the initial Route of the message.
 *
 * @see RecordRouteHeader
 * @see HeaderAddress
 * @see Parameters
 *
 * @version 1.1
 * @author Sun Microsystems
 *
 */
type RouteHeader interface {
	HeaderAddress
	ParametersHeader
	Header

	/**
	 * Name of RouteHeader
	 */
	// public final static String NAME = "Route";
}
