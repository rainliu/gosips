package header

import (
	"bytes"
	"gosip/address"
	"gosip/core"
)

/**
 *Route  SIPHeader Object
 *
 *@version  JAIN-SIP-1.1
 *
 *@author M. Ranganathan <mranga@nist.gov>  <br/>
 *
 *<a href="{@docRoot}/uncopyright.html">This code is in the public domain.</a>
 *
 */

type Route struct {
	AddressParameters //Header
	//implements javax.sip.header.RouteHeader {
}

/** Default constructor
 */
func NewRoute() *Route {
	this := &Route{}
	this.AddressParameters.super(core.SIPHeaderNames_ROUTE)
	return this
}

/** Default constructor given an address.
 *
 *@param address -- address of this header.
 *
 */

func NewRouteFromAddress(addr address.Address) *Route {
	this := &Route{}
	this.AddressParameters.super(core.SIPHeaderNames_ROUTE)
	this.addr = addr
	return this
}

/**
 * Equality predicate.
 * Two routes are equal if their addresses are equal.
 *
 *@param that is the other object to compare with.
 *@return true if the route addresses are equal.
 */
// public boolean equals(Object that) {
//     if (! this.getClass().equals(that.getClass())) return false;
//     Route thatRoute = (Route) that;
//     return  this.address.getHostPort().
//     equals(thatRoute.address.getHostPort());
// }

/**
 * Hashcode so this header can be inserted into a set.
 *
 *@return the hashcode of the encoded address.
 */
func (this *Route) HashCode() int {
	//hp, _ := this.addr.GetHostPort()
	//strcon.Atoi(hp.String() //.toLowerCase().hashCode();
	panic("Route.HashCode() Not implement yet")
	return 0
}

func (this *Route) String() string {
	return this.headerName + core.SIPSeparatorNames_COLON +
		core.SIPSeparatorNames_SP + this.EncodeBody() + core.SIPSeparatorNames_NEWLINE
}

/**
 * Encode into canonical form.
 * Acknowledgement: contains a bug fix for a bug reported by
 * Laurent Schwizer
 *
 *@return a canonical encoding of the header.
 */
func (this *Route) EncodeBody() string {
	var encoding bytes.Buffer //  = new StringBuffer();
	addr, _ := this.addr.(*address.AddressImpl)
	if addr.GetAddressType() == address.NAME_ADDR {
		encoding.WriteString(core.SIPSeparatorNames_LESS_THAN)
	}
	encoding.WriteString(this.addr.String())
	if addr.GetAddressType() == address.NAME_ADDR {
		encoding.WriteString(core.SIPSeparatorNames_GREATER_THAN)
	}

	if this.parameters != nil && this.parameters.Len() > 0 {
		encoding.WriteString(core.SIPSeparatorNames_SEMICOLON)
		encoding.WriteString(this.parameters.String())
	}
	return encoding.String()
}
