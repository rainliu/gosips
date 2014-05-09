package header

import (
	"gosips/core"
	"strings"
)

/**
*  ContentType SIP Header
* <pre>
*14.17 Content-Type
*
*   The Content-Type entity-header field indicates the media type of the
*   entity-body sent to the recipient or, in the case of the HEAD method,
*   the media type that would have been sent had the request been a Get.
*
*   Content-Type   = "Content-Type" ":" media-type
*
*   Media types are defined in section 3.7. An example of the field is
*
*       Content-Type: text/html; charSet=ISO-8859-4
*
*   Further discussion of methods for identifying the media type of an
*   entity is provided in section 7.2.1.
*
* From  HTTP RFC 2616
* </pre>
*
*
*@version  JAIN-SIP-1.1
*
*@author M. Ranganathan <mranga@nist.gov>  <br/>
*@author Olivier Deruelle <deruelle@nist.gov><br/>
*<a href="{@docRoot}/uncopyright.html">This code is in the public domain.</a>
*
 */
type ContentType struct {
	Parameters

	//implements javax.sip.header.ContentTypeHeader {

	/** mediaRange field.
	 */
	mediaRange *MediaRange
}

/** Default constructor.
 */
func NewContentType() *ContentType {
	this := &ContentType{}
	this.Parameters.super(core.SIPHeaderNames_CONTENT_TYPE)
	return this
}

/** Constructor given a content type and subtype.
*@param contentType is the content type.
*@param contentSubtype is the content subtype
 */
func NewContentTypeFromString(contentType, contentSubtype string) *ContentType {
	this := &ContentType{}
	this.Parameters.super(core.SIPHeaderNames_CONTENT_TYPE)
	this.SetContentTypeSubType(contentType, contentSubtype)
	return this
}

/** compare two MediaRange headers.
 * @param media String to Set
 * @return int.
 */
func (this *ContentType) CompareMediaRange(media string) bool {
	return strings.ToLower(this.mediaRange.mtype+"/"+this.mediaRange.subtype) == strings.ToLower(media)
}

func (this *ContentType) String() string {
	return this.headerName + core.SIPSeparatorNames_COLON +
		core.SIPSeparatorNames_SP + this.EncodeBody() + core.SIPSeparatorNames_NEWLINE
}

/**
 * Encode into a canonical string.
 * @return String.
 */
func (this *ContentType) EncodeBody() string {
	if this.HasParameters() {
		return this.mediaRange.String() + core.SIPSeparatorNames_SEMICOLON + this.parameters.String()
	} else {
		return this.mediaRange.String()
	}
}

/** Get the mediaRange field.
 * @return MediaRange.
 */
func (this *ContentType) GetMediaRange() *MediaRange {
	return this.mediaRange
}

/** Get the Media Type.
 * @return String.
 */
func (this *ContentType) GetMediaType() string {
	return this.mediaRange.mtype
}

/** Get the MediaSubType field.
 * @return String.
 */
func (this *ContentType) GetMediaSubType() string {
	return this.mediaRange.subtype
}

/** Get the content subtype.
*@return the content subtype string (or null if not Set).
 */
func (this *ContentType) GetContentSubType() string {
	if this.mediaRange == nil {
		return ""
	} else {
		return this.mediaRange.GetSubtype()
	}
}

/** Get the content subtype.
*@return the content tyep string (or null if not Set).
 */

func (this *ContentType) GetContentType() string {
	if this.mediaRange == nil {
		return ""
	} else {
		return this.mediaRange.GetType()
	}
}

/** Get the charSet parameter.
 */
func (this *ContentType) GetCharSet() string {
	return this.GetParameter("charSet")
}

/**
 * Set the mediaRange member
 * @param m mediaRange field.
 */
func (this *ContentType) SetMediaRange(m *MediaRange) {
	this.mediaRange = m
}

/**
* Set the content type and subtype.
*@param contentType Content type string.
*@param contentSubType content subtype string
 */
func (this *ContentType) SetContentTypeSubType(contentType, contentSubType string) {
	if this.mediaRange == nil {
		this.mediaRange = NewMediaRange()
	}
	this.mediaRange.SetType(contentType)
	this.mediaRange.SetSubtype(contentSubType)
}

/**
* Set the content type.
*@param contentType Content type string.
 */

func (this *ContentType) SetContentType(contentType string) { //throws ParseException{
	//if (contentType==null) throw new
	//NullPointerException( "null arg");
	if this.mediaRange == nil {
		this.mediaRange = NewMediaRange()
	}
	this.mediaRange.SetType(contentType)
}

/** Set the content subtype.
 * @param contentType String to Set
 */
func (this *ContentType) SetContentSubType(contentType string) { //throws ParseException {
	//if (contentType==null) throw new
	//NullPointerException( "null arg");
	//if (mediaRange == null) mediaRange = new MediaRange();
	if this.mediaRange == nil {
		this.mediaRange = NewMediaRange()
	}
	this.mediaRange.SetSubtype(contentType)
}
