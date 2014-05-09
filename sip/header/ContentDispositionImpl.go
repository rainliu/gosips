package header

import (
	"bytes"
	"gosips/core"
)

/**
* Content Dispositon SIP Header.
*
*@author M. Ranganathan <mranga@nist.gov>  <br/>
*
*<a href="${docRoot}/uncopyright.html">This code is in the public domain.</a>
*
*@version JAIN-SIP-1.1
*
 */
type ContentDisposition struct {
	Parameters
	//implements javax.sip.header.ContentDispositionHeader {

	/** dispositionType field.
	 */
	dispositionType string
}

/** Default constructor.
 */
func NewContentDisposition() *ContentDisposition {
	this := &ContentDisposition{}
	this.Parameters.super(core.SIPHeaderNames_CONTENT_DISPOSITION)
	return this
}

func (this *ContentDisposition) String() string {
	return this.headerName + core.SIPSeparatorNames_COLON +
		core.SIPSeparatorNames_SP + this.EncodeBody() + core.SIPSeparatorNames_NEWLINE
}

/**
 * Encode value of header into canonical string.
 * @return encoded value of header.
 *
 */
func (this *ContentDisposition) EncodeBody() string {
	var encoding bytes.Buffer //= new StringBuffer();

	if this.dispositionType != "" {
		encoding.WriteString(this.dispositionType)
	}

	if this.Parameters.parameters != nil && this.parameters.Len() > 0 {
		encoding.WriteString(core.SIPSeparatorNames_SEMICOLON)
		encoding.WriteString(this.parameters.String())
	}
	return encoding.String()
}

/** Set the disposition type.
 *@param disposition type.
 */
func (this *ContentDisposition) SetDispositionType(dispositionType string) { //throws ParseException {
	//if (dispositionType==null) throw new  NullPointerException("JAIN-SIP Exception"+
	//, ContentDisposition, SetDispositionType(), the dispositionType parameter is null");
	this.dispositionType = dispositionType
}

/** Get the disposition type.
 *@param GetDispositionType
 */
func (this *ContentDisposition) GetDispositionType() string {
	return this.dispositionType
}

/** Get the dispositionType field.
 * @return String
 */
func (this *ContentDisposition) GetHandling() string {
	return this.GetParameter("handling")
}

/** Set the dispositionType field.
 * @param type String to Set.
 */
func (this *ContentDisposition) SetHandling(handling string) {
	//throws ParseException {
	// if (handling==null) throw new  NullPointerException("JAIN-SIP Exception"+
	//", ContentDisposition, SetHandling(), the handling parameter is null");
	this.SetParameter("handling", handling)
}

/**
 * Gets the interpretation of the message body or message body part of
 *
 * this ContentDispositionHeader.
 *
 *
 *
 * @return interpretation of the message body or message body part
 *
 */
func (this *ContentDisposition) GetContentDisposition() string {
	return this.EncodeBody()
}
