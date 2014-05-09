package header

import "gosips/core"

/**
* List of ProxyAuthenticate headers.
*@version  JAIN-SIP-1.1
*
*@author M. Ranganathan <mranga@nist.gov>  <br/>
*
*<a href="{@docRoot}/uncopyright.html">This code is in the public domain.</a>
*
 */
type ProxyAuthenticateList struct {
	SIPHeaderList
}

/** Default constructor
 */
func NewProxyAuthenticateList() *ProxyAuthenticateList {
	this := &ProxyAuthenticateList{}
	this.SIPHeaderList.super(core.SIPHeaderNames_PROXY_AUTHENTICATE)
	return this
}
