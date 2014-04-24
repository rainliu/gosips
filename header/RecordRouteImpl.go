package header

import (
	"bytes"
	"gosip/address"
	"gosip/core"
)

/** The Request-Route header is added to a request by any proxy that insists on
 * being in the path of subsequent requests for the same call leg.
 *
 *@version  JAIN-SIP-1.1
 *
 *@author M. Ranganathan <mranga@nist.gov>  <br/>
 *
 *<a href="{@docRoot}/uncopyright.html">This code is in the public domain.</a>
 *
 */
type RecordRoute struct {
	AddressParameters
	//javax.sip.header.RecordRouteHeader {
}

/**  constructor
 * @param addr address to set
 */
func NewRecordRouteFromAddress(addr address.Address) *RecordRoute {
	this := &RecordRoute{}
	this.AddressParameters.super(core.SIPHeaderNames_RECORD_ROUTE)
	this.addr = addr
	return this
}

/** default constructor
 */
func NewRecordRoute() *RecordRoute {
	this := &RecordRoute{}
	this.AddressParameters.super(core.SIPHeaderNames_RECORD_ROUTE)
	return this
}

func (this *RecordRoute) String() string {
	return this.headerName + core.SIPSeparatorNames_COLON +
		core.SIPSeparatorNames_SP + this.EncodeBody() + core.SIPSeparatorNames_NEWLINE
}

/** Encode into canonical form.
 *@return String containing the canonicaly encoded header.
 */
func (this *RecordRoute) EncodeBody() string {
	var encoding bytes.Buffer //  = new StringBuffer();
	addr, _ := this.addr.(*address.AddressImpl)
	if addr.GetAddressType() == address.ADDRESS_SPEC {
		encoding.WriteString(core.SIPSeparatorNames_LESS_THAN)
	}
	encoding.WriteString(addr.String())
	if addr.GetAddressType() == address.ADDRESS_SPEC {
		encoding.WriteString(core.SIPSeparatorNames_GREATER_THAN)
	}

	if this.parameters != nil && this.parameters.Len() > 0 {
		encoding.WriteString(core.SIPSeparatorNames_SEMICOLON)
		encoding.WriteString(this.parameters.String())
	}
	return encoding.String()
}
