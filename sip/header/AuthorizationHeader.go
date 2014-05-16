package header

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
 *
 */

type AuthorizationHeader interface {
	AuthenticationHeader
}
