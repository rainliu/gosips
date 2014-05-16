package header

/**
 * This interface represents the Accept-Encoding request-header.
 * A client includes an AcceptEncodingHeader in a Request to tell the server
 * what coding schemes are acceptable in the Response e.g. compress, gzip.
 * <p>
 * If an AcceptEncodingHeader is present, and if the server cannot send a
 * Response which is acceptable according to the AcceptEncodingHeader, then
 * the server should return a Response with a status code of NOT_ACCEPTABLE.
 * <p>
 * An empty Accept-Encoding header field is permissible, it is equivalent to
 * <code>Accept-Encoding: identity</code>, meaning no encoding is permissible.
 * <p>
 * If no Accept-Encoding header field is present, the server SHOULD assume a
 * default value of identity.
 * <p>
 * For Example:<br>
 * <code>Accept-Encoding: gzip</code>
 *
 */
type AcceptEncodingHeader interface {
	ParametersHeader
	Encoding
	QValue
}
