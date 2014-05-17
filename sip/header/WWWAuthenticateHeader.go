package header

import "gosips/sip/address"

/**
 * This interface represents the WWW-Authenticate response-header. A
 * WWW-Authenticate header field value contains an authentication challenge.
 * When a UAS receives a request from a UAC, the UAS MAY authenticate
 * the originator before the request is processed.  If no credentials
 * (in the Authorization header field) are provided in the request, the
 * UAS can challenge the originator to provide credentials by rejecting
 * the request with a 401 (Unauthorized) status code. The WWW-Authenticate
 * response-header field MUST be included in 401 (Unauthorized) response
 * messages.  The field value consists of at least one challenge that indicates
 * the authentication scheme(s) and parameters applicable to the realm.
 * <p>
 * For Example:<br>
 * <code>WWW-Authenticate: Digest realm="atlanta.com", domain="sip:boxesbybob.com",
 * qop="auth", nonce="f84f1cec41e6cbe5aea9c8e88d359", opaque="", stale=FALSE,
 * algorithm=MD5</code>
 *
 * @see Parameters
 */

type WWWAuthenticateHeader interface {
	ParametersHeader

	/**
	 * Sets the scheme of the challenge information for this WWWAuthenticateHeader.
	 * For example, Digest.
	 *
	 * @param scheme - the new string value that identifies the challenge
	 * information scheme.
	 *
	 */
	SetScheme(scheme string)

	/**
	 * Returns the scheme of the challenge information for this WWWAuthenticateHeader.
	 *
	 * @return the string value of the challenge information.
	 *
	 */
	GetScheme() string

	/**
	 * Sets the Realm of the WWWAuthenicateHeader to the realm
	 * parameter value. Realm strings MUST be globally unique.  It is
	 * RECOMMENDED that a realm string contain a hostname or domain name.
	 * Realm strings SHOULD present a human-readable identifier that can be
	 * rendered to a user.
	 *
	 * @param realm the new Realm string of this WWWAuthenicateHeader.
	 * @throws ParseException which signals that an error has been reached
	 * unexpectedly while parsing the realm.
	 *
	 */
	SetRealm(realm string) (ParseException error)

	/**
	 * Returns the Realm value of this WWWAuthenicateHeader. This convenience
	 * method returns only the realm of the complete Challenge.
	 *
	 * @return the string representing the Realm information, null if value is
	 * not Set.
	 *
	 */
	GetRealm() string

	/**
	 * Sets the Nonce of the WWWAuthenicateHeader to the nonce
	 * parameter value.
	 *
	 * @param nonce - the new nonce string of this WWWAuthenicateHeader.
	 * @throws ParseException which signals that an error has been reached
	 * unexpectedly while parsing the nonce value.
	 *
	 */
	SetNonce(nonce string) (ParseException error)

	/**
	 * Returns the Nonce value of this WWWAuthenicateHeader.
	 *
	 * @return the string representing the nonce information, null if value is
	 * not Set.
	 *
	 */
	GetNonce() string

	/**
	 * Sets the URI of the WWWAuthenicateHeader to the uri
	 * parameter value.
	 *
	 * @param uri - the new URI of this WWWAuthenicateHeader.
	 *
	 */
	SetURI(uri address.URI)

	/**
	 * Returns the URI value of this WWWAuthenicateHeader, for example DigestURI.
	 *
	 * @return the URI representing the URI information, null if value is
	 * not Set.
	 *
	 */
	GetURI() address.URI

	/**
	 * Sets the Algorithm of the WWWAuthenicateHeader to the new
	 * algorithm parameter value.
	 *
	 * @param algorithm - the new algorithm string of this WWWAuthenicateHeader.
	 * @throws ParseException which signals that an error has been reached
	 * unexpectedly while parsing the algorithm value.
	 *
	 */
	SetAlgorithm(algorithm string) (ParseException error)

	/**
	 * Returns the Algorithm value of this WWWAuthenicateHeader.
	 *
	 * @return the string representing the Algorithm information, null if the
	 * value is not Set.
	 *
	 */
	GetAlgorithm() string

	/**
	 * Sets the Qop value of the WWWAuthenicateHeader to the new
	 * qop parameter value.
	 *
	 * @param qop - the new Qop string of this WWWAuthenicateHeader.
	 * @throws ParseException which signals that an error has been reached
	 * unexpectedly while parsing the Qop value.
	 *
	 */
	SetQop(qop string) (ParseException error)

	/**
	 * Returns the Qop value of this WWWAuthenicateHeader.
	 *
	 * @return the string representing the Qop information, null if the
	 * value is not Set.
	 *
	 */
	GetQop() string

	/**
	 * Sets the Opaque value of the WWWAuthenicateHeader to the new
	 * opaque parameter value.
	 *
	 * @param opaque - the new Opaque string of this WWWAuthenicateHeader.
	 * @throws ParseException which signals that an error has been reached
	 * unexpectedly while parsing the opaque value.
	 *
	 */
	SetOpaque(opaque string) (ParseException error)

	/**
	 * Returns the Opaque value of this WWWAuthenicateHeader.
	 *
	 * @return the string representing the Opaque information, null if the
	 * value is not Set.
	 *
	 */
	GetOpaque() string

	/**
	 * Sets the Domain of the WWWAuthenicateHeader to the domain
	 * parameter value.
	 *
	 * @param domain - the new Domain string of this WWWAuthenicateHeader.
	 * @throws ParseException which signals that an error has been reached
	 * unexpectedly while parsing the domain.
	 *
	 */
	SetDomain(domain string) (ParseException error)

	/**
	 * Returns the Domain value of this WWWAuthenicateHeader.
	 *
	 * @return the string representing the Domain information, null if value is
	 * not Set.
	 *
	 */
	GetDomain() string

	/**
	 * Sets the value of the stale parameter of the WWWAuthenicateHeader to the
	 * stale parameter value.
	 *
	 * @param stale - the new boolean value of the stale parameter.
	 *
	 */
	SetStale(stale bool)

	/**
	 * Returns the boolean value of the state paramater of this
	 * WWWAuthenicateHeader.
	 *
	 * @return the boolean representing if the challenge is stale.
	 *
	 */
	IsStale() bool
}
