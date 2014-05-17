package header

/**
 * The Content-Type header field indicates the media type of the message-body
 * sent to the recipient. The Content-Type header field MUST be present if the
 * body is not empty.  If the body is empty, and a Content-Type header field is
 * present, it indicates that the body of the specific type has zero length
 * (for example, an empty audio file).
 * <p>
 * For Example:<br>
 * <code>Content-Type: application/sdp</code>
 *
 * @see ContentDispositionHeader
 * @see ContentLengthHeader
 * @see ContentEncodingHeader
 * @see ContentLanguageHeader
 */

type ContentTypeHeader interface {
	ParametersHeader
	MediaType
}
