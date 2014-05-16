package header

/**
 * The Authentication-Info header field provides for mutual
 * authentication with HTTP Digest. A UAS MAY include this header field
 * in a 2xx response to a request that was successfully authenticated
 * using digest based on the Authorization header field.
 * <p>
 * For Example:<br>
 * <code>Authentication-Info: nextnonce="47364c23432d2e131a5fb210812c"</code>
 *
 * @version 1.1
 * @author Sun Microsystems
 */
type AuthenticationInfoHeader interface {
	ParametersHeader
	//Header

	/**
	 * Sets the NextNonce of the AuthenticationInfoHeader to the <var>nextNonce</var>
	 * parameter value.
	 *
	 * @param nextNonce - the new nextNonce String of this AuthenticationInfoHeader.
	 * @throws ParseException which signals that an error has been reached
	 * unexpectedly while parsing the nextNonce value.
	 *
	 */
	SetNextNonce(nextNonce string) //throws ParseException;

	/**
	 * Returns the nextNonce value of this AuthenticationInfoHeader.
	 *
	 * @return the String representing the nextNonce information, null if value is
	 * not set.
	 *
	 */
	GetNextNonce() string

	/**
	 * Sets the Qop value of the AuthenticationInfoHeader to the new
	 * <var>qop</var> parameter value.
	 *
	 * @param qop - the new Qop string of this AuthenticationInfoHeader.
	 * @throws ParseException which signals that an error has been reached
	 * unexpectedly while parsing the Qop value.
	 *
	 */
	SetQop(qop string) // throws ParseException;

	/**
	 * Returns the messageQop value of this AuthenticationInfoHeader.
	 *
	 * @return the string representing the messageQop information, null if the
	 * value is not set.
	 *
	 */
	GetQop() string

	/**
	 * Sets the CNonce of the AuthenticationInfoHeader to the <var>cNonce</var>
	 * parameter value.
	 *
	 * @param cNonce - the new cNonce String of this AuthenticationInfoHeader.
	 * @throws ParseException which signals that an error has been reached
	 * unexpectedly while parsing the cNonce value.
	 *
	 */
	SetCNonce(cNonce string) //throws ParseException;

	/**
	 * Returns the CNonce value of this AuthenticationInfoHeader.
	 *
	 * @return the String representing the cNonce information, null if value is
	 * not set.
	 *
	 */
	GetCNonce() string

	/**
	 * Sets the Nonce Count of the AuthenticationInfoHeader to the <var>nonceCount</var>
	 * parameter value.
	 *
	 * @param nonceCount - the new nonceCount integer of this AuthenticationInfoHeader.
	 * @throws ParseException which signals that an error has been reached
	 * unexpectedly while parsing the nonceCount value.
	 *
	 */
	SetNonceCount(nonceCount int) //throws ParseException;

	/**
	 * Returns the Nonce Count value of this AuthenticationInfoHeader.
	 *
	 * @return the integer representing the nonceCount information, -1 if value is
	 * not set.
	 *
	 */
	GetNonceCount() int

	/**
	 * Sets the Response of the AuthenticationInfoHeader to the new <var>response</var>
	 * parameter value.
	 *
	 * @param response - the new response String of this AuthenticationInfoHeader.
	 * @throws ParseException which signals that an error has been reached
	 * unexpectedly while parsing the Response.
	 *
	 */
	SetResponse(response string) // throws ParseException;

	/**
	 * Returns the Response value of this AuthenticationInfoHeader.
	 *
	 * @return the String representing the Response information.
	 *
	 */
	GetResponse() string

	/**
	 * Name of the AlertInfoHeader
	 */
	//public final static String NAME = "Authentication-Info";

}
