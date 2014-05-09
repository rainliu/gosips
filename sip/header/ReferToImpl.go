package header

import (
	"bytes"
	"gosips/core"
	"gosips/sip/address"
)

/**
*ReferTo SIP Header.
*
*@version  JAIN-SIP-1.1
*
*@author M. Ranganathan <mranga@nist.gov>  <br/>
*@author Olivier Deruelle <deruelle@nist.gov><br/>
*
*<a href="{@docRoot}/uncopyright.html">This code is in the public domain.</a>
*
 */

type ReferTo struct {
	AddressParameters //Header implements
	//javax.sip.header.ReferToHeader {
}

/** default Constructor.
 */
func NewReferTo() *ReferTo {
	this := &ReferTo{}
	this.AddressParameters.super(core.SIPHeaderNames_REFER_TO)
	return this
}

func (this *ReferTo) String() string {
	return this.headerName + core.SIPSeparatorNames_COLON +
		core.SIPSeparatorNames_SP + this.EncodeBody() + core.SIPSeparatorNames_NEWLINE
}

/**
 * Encode the header content into a String.
 * @return String
 */
func (this *ReferTo) EncodeBody() string {
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
