package header

/**
 * A Contact header field value provides a URI whose meaning depends on
 * the type of request or response it is in. A Contact header field value
 * can contain a display name, a URI with URI parameters, and header
 * parameters.
 * <p>
 * The Contact header field provides a SIP or SIPS URI that can be used
 * to contact that specific instance of the User Agent for subsequent requests.
 * The Contact header field MUST be present and contain exactly one SIP
 * or SIPS URI in any request that can result in the establishment of a
 * dialog.  For the methods defined in this specification, that includes
 * only the INVITE request.  For these requests, the scope of the
 * Contact is global.  That is, the Contact header field value contains
 * the URI at which the User Agent would like to receive requests, and this URI
 * MUST be valid even if used in subsequent requests outside of any
 * dialogs.
 * <p>
 * If the Request-URI or top Route header field value contains a SIPS URI, the
 * Contact header field MUST contain a SIPS URI as well.
 * <p>
 * <b>Messages and Contact Headers</b>
 * <ul>
 * <li>Requests: A contact header is mandatory for INVITE's and optional for
 * ACK, OPTIONS and REGISTER requests. This allows the callee to send future
 * Requests, such as BYE Requests, directly to the caller instead of through a
 * series of proxies.
 * <li>Informational Responses - A contact header is optional in a Informational
 * Response to an INVITE request. It has the same semantics in an Informational
 * Response as a Success Response.
 * <li>Success Responses - A contact header is mandatory in response to INVITE's
 * and optional in response to OPTIONS and REGISTER requests. An user agent
 * server sending a Success Response to an INIVTE must insert a ContactHeader
 * in the Response indicating the SIP address under which it is reachable most
 * directly for future SIP Requests.
 * <li>Redirection Responses - A contact header is optional in response to
 * INVITE's, OPTIONS, REGISTER and BYE requests. A proxy may also delete the
 * contact header.
 * <li>Ambiguous Header: - A contact header is optional in response to
 * INVITE, OPTIONS, REGISTER and BYE requests.
 * </ul>
 *
 * The ContactHeader defines the Contact parameters "q" and "expires".
 * The <code>q-value</code> value is used to prioritize addresses in a
 * list of contact addresses. The <code>expires</code> value suggests an
 * expiration interval that indicates how long the client would like a
 * registration to be valid for a specific address. These parameters are only
 * used when the Contact is present in a:
 * <ul>
 * <li>REGISTER request
 * <li>REGISTER response
 * <li>3xx response
 * </ul>
 *
 * For Example:<br>
 * <code> Contact: "Mr. Watson" sip:watson@worcester.jcp.org;
 * q=0.7; expires=3600, "Mr. Watson" mailto:watson@jcp.org.com; q=0.1
 *
 * @see AddressHeader
 * @see Parameters
 */

type ContactHeader interface {
	AddressHeader
	ParametersHeader
	QValue

	/**
	 * Returns the value of the <code>expires</code> parameter as delta-seconds.
	 * When a client sends a REGISTER request, it MAY suggest an expiration
	 * interval that indicates how long the client would like the registration
	 * to be valid for a specific address. There are two ways in which a client
	 * can suggest an expiration interval for a binding:
	 * <ul>
	 * <li>through an Expires header field
	 * <li>an "expires" Contact header parameter.
	 * </ul>
	 * The latter allows expiration intervals to be suggested on a per-binding
	 * basis when more than one binding is given in a single REGISTER request,
	 * whereas the former suggests an expiration interval for all Contact
	 * header field values that do not contain the "expires" parameter. If
	 * neither mechanism for expressing a suggested expiration time is present
	 * in a REGISTER, the client is indicating its desire for the server to
	 * choose.
	 * <p>
	 * A User Agent requests the immediate removal of a binding by specifying an
	 * expiration interval of "0" for that contact address in a REGISTER
	 * request.  User Agents SHOULD support this mechanism so that bindings can be
	 * removed before their expiration interval has passed. The
	 * REGISTER-specific Contact header field value of "*" applies to all
	 * registrations, but it MUST NOT be used unless the Expires header
	 * field is present with a value of "0". The "*" value can be determined
	 * if "this.getNameAddress().isWildcard() = = true".
	 *
	 * @param seconds - new relative value of the expires parameter.
	 * 0 implies removal of Registration specified in Contact Header.
	 * @throws InvalidArgumentException if supplied value is less than zero.
	 */

	SetExpires(expires int) (InvalidArgumentException error)

	/**
	 * Returns the value of the <code>expires</code> parameter.
	 *
	 * @return value of the <code>expires</code> parameter measured in
	 * delta-seconds, O implies removal of Registration specified in Contact
	 * Header.
	 */

	GetExpires() int
}
