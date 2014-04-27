package header

import (
	"gosip/core"
	"strconv"
)

/**
* ContentLength ISIPHeader (of which there can be only one in a SIPMessage).
*<pre>
*Fielding, et al.            Standards Track                   [Page 119]
*RFC 2616                        HTTP/1.1                       June 1999
*
*
*      14.13 Content-Length
*
*   The Content-Length entity-header field indicates the size of the
*   entity-body, in decimal number of OCTETs, sent to the recipient or,
*   in the case of the HEAD method, the size of the entity-body that
*   would have been sent had the request been a Get.
*
*       Content-Length    = "Content-Length" ":" 1*DIGIT
*
*   An example is
*
*       Content-Length: 3495
*
*   Applications SHOULD use this field to indicate the transfer-length of
*   the message-body, unless this is prohibited by the rules in section
*   4.4.
*
*   Any Content-Length greater than or equal to zero is a valid value.
*   Section 4.4 describes how to determine the length of a message-body
*   if a Content-Length is not given.
*
*   Note that the meaning of this field is significantly different from
*   the corresponding definition in MIME, where it is an optional field
*   used within the "message/external-body" content-type. In HTTP, it
*   SHOULD be sent whenever the message's length can be determined prior
*   to being transferred, unless this is prohibited by the rules in
*   section 4.4.
* </pre>
*
*@author M. Ranganathan <mranga@nist.gov>  <br/>
*@author Olivier Deruelle <deruelle@nist.gov><br/>
 */
type ContentLength struct {
	SIPHeader //implements
	//javax.sip.header.ContentLengthHeader {

	/** contentLength field.
	 */
	contentLength int
}

/** Default constructor.
 */
func NewContentLength() *ContentLength {
	this := &ContentLength{}
	this.SIPHeader.super(core.SIPHeaderNames_CONTENT_LENGTH)
	return this
}

/**
 *Constructor given a length.
 */
func NewContentLengthFromInt(length int) *ContentLength {
	this := &ContentLength{}
	this.SIPHeader.super(core.SIPHeaderNames_CONTENT_LENGTH)
	this.contentLength = length
	return this
}

/** Get the ContentLength field.
 * @return int
 */
func (this *ContentLength) GetContentLength() int {
	return this.contentLength
}

/**
 * Set the contentLength member
 * @param contentLength int to Set
 */
func (this *ContentLength) SetContentLength(contentLength int) error { //throws InvalidArgumentException{
	//if (contentLength<0) throw new InvalidArgumentException("JAIN-SIP Exception"+
	//", ContentLength, SetContentLength(), the contentLength parameter is <0");
	this.contentLength = contentLength
	return nil
}

func (this *ContentLength) String() string {
	return this.headerName + core.SIPSeparatorNames_COLON +
		core.SIPSeparatorNames_SP + this.EncodeBody() + core.SIPSeparatorNames_NEWLINE
}

/**
 * Encode into a canonical string.
 * @return String
 */
func (this *ContentLength) EncodeBody() string {
	//if (contentLength == null) return "0";
	return strconv.Itoa(this.contentLength)
}

/** Pattern matcher ignores content length.
 */
func (this *ContentLength) Match(other interface{}) bool {
	_, ok := other.(*ContentLength)
	return ok
}
