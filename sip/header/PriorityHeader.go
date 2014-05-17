package header

/**

 * The Priority header field indicates the urgency of the request as perceived

 * by the client.  The Priority header field describes the priority that the

 * SIP request should have to the receiving human or its agent. For example,

 * it may be factored into decisions about call routing and acceptance. For

 * these decisions, a message containing no Priority header field SHOULD be

 * treated as if it specified a Priority of <var>"Normal"</var>.

 * <p>

 * The Priority header field does not influence the use of communications

 * resources such as packet forwarding priority in routers or access to

 * circuits in PSTN gateways.

 * <p>

 * The currently defined priority values are:

 * <ul>

 * <li> EMERGENCY

 * <li> URGENT

 * <li> NORMAL

 * <li> NON_URGENT

 * </ul>

 * For Example:<br>

 * <code>Subject: Weekend plans<br>

 * Priority: non-urgent</code>

 */

type PriorityHeader interface {
	Header

	/**

	 * Set priority of PriorityHeader

	 *

	 * @param priority - the new string priority value

	 * @throws ParseException which signals that an error has been reached

	 * unexpectedly while parsing the priority value.

	 */

	SetPriority(priority string) (ParseException error)

	/**

	 * Gets the string priority value of the PriorityHeader.

	 *

	 * @return the string priority value of the PriorityHeader

	 */

	GetPriority() string
}
