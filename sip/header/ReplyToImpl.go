package header

import (
	"bytes"
	"gosips/core"
	"gosips/sip/address"
)

/**
* ReplyTo Header.
*
*@version  JAIN-SIP-1.1
*
*@author M. Ranganathan <mranga@nist.gov>  <br/>
*@author Olivier Deruelle <deruelle@nist.gov><br/>
*
*<a href="{@docRoot}/uncopyright.html">This code is in the public domain.</a>
*
 */
type ReplyTo struct {
	AddressParameters //Header
	//implements ReplyToHeader {
}

/** Default constructor
 */
func NewReplyTo() *ReplyTo {
	this := &ReplyTo{}
	this.AddressParameters.super(core.SIPHeaderNames_REPLY_TO)
	return this
}

/** Default constructor given an address.
 *
 *@param address -- address of this header.
 *
 */
func NewReplyToFromAddress(addr address.Address) *ReplyTo {
	this := &ReplyTo{}
	this.AddressParameters.super(core.SIPHeaderNames_REPLY_TO)
	this.addr = addr
	return this
}

func (this *ReplyTo) String() string {
	return this.headerName + core.SIPSeparatorNames_COLON +
		core.SIPSeparatorNames_SP + this.EncodeBody() + core.SIPSeparatorNames_NEWLINE
}

/**
 * Encode the header content into a String.
 * @return String
 */
func (this *ReplyTo) EncodeBody() string {
	var encoding bytes.Buffer //  = new StringBuffer();
	addr, _ := this.addr.(*address.AddressImpl)
	if addr.GetAddressType() == address.ADDRESS_SPEC {
		encoding.WriteString(core.SIPSeparatorNames_LESS_THAN)
	}
	encoding.WriteString(this.addr.String())
	if addr.GetAddressType() == address.ADDRESS_SPEC {
		encoding.WriteString(core.SIPSeparatorNames_GREATER_THAN)
	}

	if this.parameters != nil && this.parameters.Len() > 0 {
		encoding.WriteString(core.SIPSeparatorNames_SEMICOLON)
		encoding.WriteString(this.parameters.String())
	}
	return encoding.String()
}

/**
 * Conveniance accessor function to get the hostPort field from the address
 * @return HostPort
 */
func (this *ReplyTo) GetHostPort() *core.HostPort {
	addr, _ := this.addr.(*address.AddressImpl)
	hp, _ := addr.GetHostPort()
	return hp
}

/**
 * Get the display name from the address.
 * @return String
 */
func (this *ReplyTo) GetDisplayName() string {
	return this.addr.GetDisplayName()
}
