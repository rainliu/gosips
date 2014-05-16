package header

/**
 * The Allow header field lists the set of methods supported by the User Agent
 * generating the message. All methods, including ACK and CANCEL, understood
 * by the User Agent MUST be included in the list of methods in the Allow header
 * field, when present.
 * The absence of an Allow header field MUST NOT be interpreted to mean that
 * the User Agent sending the message supports no methods. Rather, it implies
 * that the User Agent is not providing any information on what methods it
 * supports. Supplying an Allow header field in responses to methods other than
 * OPTIONS reduces the number of messages needed.
 * <p>
 * For Example:<br>
 * <code>Allow: INVITE, ACK, OPTIONS, CANCEL, BYE</code>
 */
type AllowHeader interface {
	Header

	/**
	 * Sets the method supported by this AllowHeader.
	 *
	 * @param method - the String defining the method supported
	 * in this AllowHeader
	 * @throws ParseException which signals that an error has been reached
	 * unexpectedly while parsing the method supported.
	 */
	SetMethod(method string) (ParseException error)

	/**
	 * Gets the method of the AllowHeader. Returns null if no method is
	 * defined in this Allow Header.
	 *
	 * @return the string identifing the method of AllowHeader.
	 */
	GetMethod() string
}
