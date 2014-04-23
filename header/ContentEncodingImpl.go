package header

import (
	"gosip/core"
)

/**
* Content encoding part of a content encoding header list.
*@see ContentEncodingList
*<pre>
* From HTTP RFC 2616
*14.11 Content-Encoding
*
*   The Content-Encoding entity-header field is used as a modifier to the
*   media-type. When present, its value indicates what additional content
*   codings have been applied to the entity-body, and thus what decoding
*   mechanisms must be applied in order to obtain the media-type
*   referenced by the Content-Type header field. Content-Encoding is
*   primarily used to allow a document to be compressed without losing
*   the identity of its underlying media type.
*
*       Content-Encoding  = "Content-Encoding" ":" 1#content-coding
*
*   Content codings are defined in section 3.5. An example of its use is
*
*       Content-Encoding: gzip
*
*   The content-coding is a characteristic of the entity identified by
*   the Request-URI. Typically, the entity-body is stored with this
*   encoding and is only decoded before rendering or analogous usage.
*   However, a non-transparent proxy MAY modify the content-coding if the
*   new coding is known to be acceptable to the recipient, unless the
*   "no-transform" cache-control directive is present in the message.
*
*   If the content-coding of an entity is not "identity", then the
*   response MUST include a Content-Encoding entity-header (section
*   14.11) that lists the non-identity content-coding(s) used.
*
*   If the content-coding of an entity in a request message is not
*   acceptable to the origin server, the server SHOULD respond with a
*   status code of 415 (Unsupported Media Type).
*
*   If multiple encodings have been applied to an entity, the content
*   codings MUST be listed in the order in which they were applied.
*   Additional information about the encoding parameters MAY be provided
*   by other entity-header fields not defined by this specification.
*</pre>
*
*@author M. Ranganathan <mranga@nist.gov>  <br/>
*@author Olivier Deruelle <deruelle@nist.gov><br/>
 */
type ContentEncoding struct {
	SIPHeader
	//implements 	javax.sip.header.ContentEncodingHeader {

	/** contentEncoding field.
	 */
	contentEncoding string
}

/** Default constructor.
 */
func NewContentEncoding() *ContentEncoding {
	this := &ContentEncoding{}
	this.SIPHeader.super(core.SIPHeaderNames_CONTENT_ENCODING)
	return this
}

/** constructor.
 * @param enc String to set.
 */
func NewContentEncodingFromString(enc string) *ContentEncoding {
	this := &ContentEncoding{}
	this.SIPHeader.super(core.SIPHeaderNames_CONTENT_ENCODING)
	this.contentEncoding = enc
	return this
}

func (this *ContentEncoding) String() string {
	return this.headerName + core.SIPSeparatorNames_COLON +
		core.SIPSeparatorNames_SP + this.EncodeBody() + core.SIPSeparatorNames_NEWLINE
}

/**
 * Canonical encoding of body of the header.
 * @return  encoded body of the header.
 */
func (this *ContentEncoding) EncodeBody() string {
	return this.contentEncoding
}

/** get the ContentEncoding field.
 * @return String
 */
func (this *ContentEncoding) GetEncoding() string {
	return this.contentEncoding
}

/** set the ConentEncoding field.
 * @param encoding String to set
 */
func (this *ContentEncoding) SetEncoding(encoding string) { //throws ParseException {
	// if (encoding==null) throw new  NullPointerException("JAIN-SIP Exception, "+
	// " encoding is null");
	this.contentEncoding = encoding
}
