package header

import "gosip/core"

/**
* ProxyAuthorization SIP header.
*
* @see ProxyAuthorization
*
* @author M. Ranganathan <mranga@nist.gov>  NIST/ITL/ANTD <br/>
*
*<a href="{@docRoot}/uncopyright.html">This code is in the public domain.</a>
*
 */
type ProxyAuthorization struct {
	Authentication
}

/** Default constructor.
 */
func NewProxyAuthorization() *ProxyAuthorization {
	this := &ProxyAuthorization{}
	this.Authentication.super(core.SIPHeaderNames_PROXY_AUTHORIZATION)
	return this
}
