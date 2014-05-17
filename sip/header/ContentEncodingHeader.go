package header

/**
 * A ContentEncodingHeader is used as a modifier to the "media-type". When
 * present, its value indicates what additional content codings have been
 * applied to the entity-body, and thus what decoding mechanisms must be
 * applied in order to obtain the media-type referenced by the
 * ContentTypeHeader. The ContentEncodingHeader is primarily used to allow a
 * body to be compressed without losing the identity of its underlying media
 * type.
 * <p>
 * If multiple encodings have been applied to an entity, the ContentEncodings
 * must be listed in the order in which they were applied. All content-coding
 * values are case-insensitive. Clients MAY apply content encodings to the body
 * in requests. A server MAY
 * apply content encodings to the bodies in responses. The server MUST only
 * use encodings listed in the Accept-Encoding header field in the request.
 * If the server is not capable of decoding the body, or does not recognize any
 * of the content-coding values, it must send a UNSUPPORTED_MEDIA_TYPE
 * Response, listing acceptable encodings in an AcceptEncodingHeader.
 *
 * @see ContentDispositionHeader
 * @see ContentLengthHeader
 * @see ContentTypeHeader
 * @see ContentLanguageHeader
 */

type ContentEncodingHeader interface {
	Header
	Encoding
}
