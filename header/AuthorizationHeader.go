package header

import "gosip/address"

/**
 * The Authorization header is used when a user agent wishes to authenticate
 * itself with a server - usually, but not necessarily, after receiving an
 * UNAUTHORIZED Response - by including an AuthorizationHeader with the
 * Request. The AuthorizationHeader consists of credentials
 * containing the authentication information of the user agent for the
 * realm of the resource being requested.
 * <p>
 * This header field, along with Proxy-Authorization, breaks the general
 * rules about multiple header field values.  Although not a comma-
 * separated list, this header field name may be present multiple times,
 * and MUST NOT be combined into a single header line.
 * <p>
 * For Example:<br>
 * <code>Authorization: Digest username="Alice", realm="atlanta.com",<br>
 *      nonce="84a4cc6f3082121f32b42a2187831a9e",<br>
 *      response="7587245234b3434cc3412213e5f113a5432"</code>
 *
 * @see Parameters
 * @see WWWAuthenticateHeader
 * @see ProxyAuthorizationHeader
 * @version 1.1
 * @author Sun Microsystems
 *
 */

type AuthorizationHeader interface {
	ParametersHeader
	Header

	/**
	 * Sets the scheme of the Response information for this AuthorizationHeader.
	 * For example, Digest.
	 *
	 * @param scheme - the new string value that identifies the response
	 * information scheme.
	 * @since v1.1
	 */
	SetScheme(scheme string)

	/**
	 * Returns the scheme of the Response information for this AuthorizationHeader.
	 *
	 * @return the string value of the response information.
	 * @since v1.1
	 */
	GetScheme() string

	/**
	 * Sets the Realm of the AuthorizationHeader to the <var>realm</var>
	 * parameter value. Realm strings MUST be globally unique.  It is
	 * RECOMMENDED that a realm string contain a hostname or domain name.
	 * Realm strings SHOULD present a human-readable identifier that can be
	 * rendered to a user.
	 *
	 * @param realm the new Realm String of this AuthorizationHeader.
	 * @throws ParseException which signals that an error has been reached
	 * unexpectedly while parsing the realm.
	 * @since v1.1
	 */
	SetRealm(realm string) // throws ParseException;

	/**
	 * Returns the Realm value of this AuthorizationHeader. This convenience
	 * method returns only the realm of the complete Response.
	 *
	 * @return the String representing the Realm information, null if value is
	 * not Set.
	 * @since v1.1
	 */
	GetRealm() string

	/**
	 * Sets the Username of the AuthorizationHeader to the <var>username</var>
	 * parameter value.
	 *
	 * @param username the new Username String of this AuthorizationHeader.
	 * @throws ParseException which signals that an error has been reached
	 * unexpectedly while parsing the username.
	 * @since v1.1
	 */
	SetUsername(username string) //throws ParseException;

	/**
	 * Returns the Username value of this AuthorizationHeader. This convenience
	 * method returns only the username of the complete Response.
	 *
	 * @return the String representing the Username information, null if value is
	 * not Set.
	 * @since v1.1
	 */
	GetUsername() string

	/**
	 * Sets the Nonce of the AuthorizationHeader to the <var>nonce</var>
	 * parameter value.
	 *
	 * @param nonce - the new nonce String of this AuthorizationHeader.
	 * @throws ParseException which signals that an error has been reached
	 * unexpectedly while parsing the nonce value.
	 * @since v1.1
	 */
	SetNonce(nonce string) //throws ParseException;

	/**
	 * Returns the Nonce value of this AuthorizationHeader.
	 *
	 * @return the String representing the nonce information, null if value is
	 * not Set.
	 * @since v1.1
	 */
	GetNonce() string

	/**
	 * Sets the URI of the AuthorizationHeader to the <var>uri</var>
	 * parameter value.
	 *
	 * @param uri - the new URI of this AuthorizationHeader.
	 * @since v1.1
	 */
	SetURI(uri address.URI)

	/**
	 * Returns the URI value of this AuthorizationHeader, for example DigestURI.
	 *
	 * @return the URI representing the URI information, null if value is
	 * not Set.
	 * @since v1.1
	 */
	GetURI() address.URI

	/**
	 * Sets the Response of the AuthorizationHeader to the new <var>response</var>
	 * parameter value.
	 *
	 * @param response - the new response String of this AuthorizationHeader.
	 * @throws ParseException which signals that an error has been reached
	 * unexpectedly while parsing the Response.
	 * @since v1.1
	 */
	SetResponse(response string) //throws ParseException;

	/**
	 * Returns the Response value of this AuthorizationHeader.
	 *
	 * @return the String representing the Response information.
	 * @since v1.1
	 */
	GetResponse() string

	/**
	 * Sets the Algorithm of the AuthorizationHeader to the new
	 * <var>algorithm</var> parameter value.
	 *
	 * @param algorithm - the new algorithm String of this AuthorizationHeader.
	 * @throws ParseException which signals that an error has been reached
	 * unexpectedly while parsing the algorithm value.
	 * @since v1.1
	 */
	SetAlgorithm(algorithm string) //throws ParseException;

	/**
	 * Returns the Algorithm value of this AuthorizationHeader.
	 *
	 * @return the String representing the Algorithm information, null if the
	 * value is not Set.
	 * @since v1.1
	 */
	GetAlgorithm() string

	/**
	 * Sets the CNonce of the AuthorizationHeader to the <var>cNonce</var>
	 * parameter value.
	 *
	 * @param cNonce - the new cNonce String of this AuthorizationHeader.
	 * @throws ParseException which signals that an error has been reached
	 * unexpectedly while parsing the cNonce value.
	 * @since v1.1
	 */
	SetCNonce(cNonce string) //throws ParseException;

	/**
	 * Returns the CNonce value of this AuthorizationHeader.
	 *
	 * @return the String representing the cNonce information, null if value is
	 * not Set.
	 * @since v1.1
	 */
	GetCNonce() string

	/**
	 * Sets the Opaque value of the AuthorizationHeader to the new
	 * <var>opaque</var> parameter value.
	 *
	 * @param opaque - the new Opaque string of this AuthorizationHeader.
	 * @throws ParseException which signals that an error has been reached
	 * unexpectedly while parsing the opaque value.
	 * @since v1.1
	 */
	SetOpaque(opaque string) // throws ParseException;

	/**
	 * Returns the Opaque value of this AuthorizationHeader.
	 *
	 * @return the String representing the Opaque information, null if the
	 * value is not Set.
	 * @since v1.1
	 */
	GetOpaque() string

	/**
	 * Sets the MessageQop value of the AuthorizationHeader to the new
	 * <var>qop</var> parameter value.
	 *
	 * @param qop - the new Qop string of this AuthorizationHeader.
	 * @throws ParseException which signals that an error has been reached
	 * unexpectedly while parsing the Qop value.
	 * @since v1.1
	 */
	SetQop(qop string) //throws ParseException;

	/**
	 * Returns the Qop value of this AuthorizationHeader.
	 *
	 * @return the string representing the Qop information, null if the
	 * value is not Set.
	 * @since v1.1
	 */
	GetQop() string

	/**
	 * Sets the Nonce Count of the AuthorizationHeader to the <var>nonceCount</var>
	 * parameter value.
	 *
	 * @param nonceCount - the new nonceCount integer of this AuthorizationHeader.
	 * @throws ParseException which signals that an error has been reached
	 * unexpectedly while parsing the nonceCount value.
	 * @since v1.1
	 */
	SetNonceCount(nonceCount int) //throws ParseException;

	/**
	 * Returns the Nonce Count value of this AuthorizationHeader.
	 *
	 * @return the integer representing the nonceCount information, -1 if value is
	 * not Set.
	 * @since v1.1
	 */
	GetNonceCount() int

	/**
	 * Name of AuthorizationHeader
	 */
	//public final static String NAME = "Authorization";

}
