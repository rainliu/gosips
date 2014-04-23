package header

import "gosip/core"

/**
* WWWAuthenticate SIPHeader (of which there can be several?)
*
*@version  JAIN-SIP-1.1
*
*@author M. Ranganathan <mranga@nist.gov>  <br/>
*
*<a href="{@docRoot}/uncopyright.html">This code is in the public domain.</a>
*
 */
type WWWAuthenticateList struct {
	SIPHeaderList
}

/**
 * constructor.
 */
func NewWWWAuthenticateList() *WWWAuthenticateList {
	this := &WWWAuthenticateList{}
	this.SIPHeaderList.super(core.SIPHeaderNames_WWW_AUTHENTICATE)
	return this
}
