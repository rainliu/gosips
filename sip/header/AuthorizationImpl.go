package header

import "gosips/core"

/**
* Authorization SIP header.
*
* @see ProxyAuthorization
*
* @author M. Ranganathan <mranga@nist.gov>  NIST/ITL/ANTD <br/>
*
*<a href="{@docRoot}/uncopyright.html">This code is in the public domain.</a>
*
 */
type Authorization struct {
	Authentication
}

/** Default constructor.
 */
func NewAuthorization() *Authorization {
	this := &Authorization{}
	this.Authentication.super(core.SIPHeaderNames_AUTHORIZATION)
	return this
}
