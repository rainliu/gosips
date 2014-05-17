package header

import "gosips/core"

/**
* ProxyAuthorization SIP header.
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
