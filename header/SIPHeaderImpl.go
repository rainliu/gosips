package header

import (
	"gosip/core"
	"strings"
)

/**
*  Root class from which all Header objects are subclassed.
 */
type SIPHeader struct { //implements SIPHeaderNames, javax.sip.header.Header {
	// SIPObject

	/** name of this header
	 */
	headerName string
}

/** Value of the header.
 */

/** Constructor
 * @param hname String to set
 */
func NewSIPHeader(hname string) *SIPHeader {
	return &SIPHeader{headerName: hname}
}

func (this *SIPHeader) super(hname string) {
	this.headerName = hname
}

/**
 * Name of the Header
 * @return String
 */
func (this *SIPHeader) GetHeaderName() string {
	return this.headerName
}

/** Alias for getHaderName above.
*
*@return String headerName
*
 */
func (this *SIPHeader) GetName() string {
	return this.headerName
}

/**
 * Set the name of the header .
 * @param hdrname String to set
 */
func (this *SIPHeader) SetHeaderName(hdrname string) {
	this.headerName = hdrname
}

/** Get the header value (i.e. what follows the name:).
* This merely goes through and lops off the portion that follows
* the headerName:
 */
func (this *SIPHeader) GetHeaderValue() string {
	//String encodedHdr = null;
	//try {
	encodedHdr := this.String()
	//} catch (Exception ex) {
	//	return null;
	//}
	//var buffer bytes.Buffer;//new StringBuffer(encodedHdr);
	buffer := []byte(encodedHdr)
	for len(buffer) > 0 && buffer[0] != ':' {
		buffer = buffer[1:]
	}

	if len(buffer) > 0 {
		buffer = buffer[1:]
	}

	return strings.TrimSpace(string(buffer))
}

/** Return false if this is not a header list
* (SIPHeaderList overrrides this method).
*@return false
 */
func (this *SIPHeader) IsHeaderList() bool {
	return false
}

/** Encode this header into canonical form.
 */
func (this *SIPHeader) String() string {
	return this.headerName + core.SIPSeparatorNames_COLON +
		core.SIPSeparatorNames_SP + this.EncodeBody() + core.SIPSeparatorNames_NEWLINE
}

func (this *SIPHeader) Clone() interface{} {
	return &SIPHeader{headerName: this.headerName}
}

/** Encode the body of this header (the stuff that follows headerName).
* A.K.A headerValue.
 */
func (this *SIPHeader) EncodeBody() string {
	return ""
}

/** Alias for getHeaderValue.
 */
func (this *SIPHeader) GetValue() string {
	return this.GetHeaderValue()
}
