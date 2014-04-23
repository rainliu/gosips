package header

import (
	"gosip/core"
	"strings"
)

/**
* A generic extension header for the stack.
* The input text of the header gets recorded here.
 */

type Extension struct { //implements  javax.sip.header.ExtensionHeader  {
	SIPHeader

	value string
}

/** This was added to allow for automatic cloning of headers.
 */
//public Extension() {}

func NewExtension(headerName string) *Extension {
	this := &Extension{}

	this.SIPHeader.super(headerName)

	return this
}

func (this *Extension) super(headerName string) {
	this.SIPHeader.super(headerName)
}

/** Set the name of the header.
*@param headerName is the name of the header to set.
 */

func (this *Extension) SetName(headerName string) {
	this.headerName = headerName
}

/** Set the value of the header.
 */
func (this *Extension) SetValue(value string) {
	this.value = value
}

/** Get the value of the extension header.
*@return the value of the extension header.
 */
func (this *Extension) GetHeaderValue() string {
	if this.value != "" {
		return this.value
	} else {
		var encodedHdr string
		//try {
		// Bug fix submitted by Lamine Brahimi
		encodedHdr = this.String()
		//} catch (Exception ex) {
		//	return null;
		//}
		buffer := []byte(encodedHdr)
		for len(buffer) > 0 && buffer[0] != ':' {
			buffer = buffer[1:]
		}
		buffer = buffer[1:]
		this.value = strings.TrimSpace(string(buffer))
		return this.value
	}
}

/** Return the canonical encoding of this header.
 */
func (this *Extension) String() string {
	return this.headerName + core.SIPSeparatorNames_COLON +
		core.SIPSeparatorNames_SP + this.value +
		core.SIPSeparatorNames_NEWLINE
}

/** Return just the body of this header encoded (leaving out the
* name and the CRLF at the end).
 */

func (this *Extension) EncodeBody() string {
	return this.GetHeaderValue()
}
