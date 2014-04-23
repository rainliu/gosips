package header

import (
	"gosip/core"
)

/**
* Allow SIPHeaderHeader.
 */
type Allow struct {
	SIPHeader // implements javax.sip.header.AllowHeader {

	/** method field
	 */
	method string
}

/** default constructor
 */
func NewAllow() *Allow {
	this := &Allow{}
	this.SIPHeader.super(core.SIPHeaderNames_ALLOW)
	return this
}

/** constructor
 * @param m String to set
 */
func NewAllowFromString(m string) *Allow {
	this := &Allow{}
	this.SIPHeader.super(core.SIPHeaderNames_ALLOW)
	this.method = m
	return this
}

/** get the method field
 * @return String
 */
func (this *Allow) GetMethod() string {
	return this.method
}

/**
 * Set the method member
 * @param method method to set.
 */
func (this *Allow) SetMethod(method string) { //throws ParseException{
	//if (method==null)
	//throw new  NullPointerException("JAIN-SIP Exception"+
	//", Allow, setMethod(), the method parameter is null.");
	this.method = method
}

func (this *Allow) String() string {
	return this.headerName + core.SIPSeparatorNames_COLON +
		core.SIPSeparatorNames_SP + this.EncodeBody() + core.SIPSeparatorNames_NEWLINE
}

/** Return body encoded in canonical form.
 * @return body encoded as a string.
 */
func (this *Allow) EncodeBody() string {
	return this.method
}
