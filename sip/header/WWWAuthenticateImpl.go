package header

import "gosips/core"

/** The WWWAuthenticate SIP header.
*
*@author Olivier Deruelle <deruelle@nist.gov><br/>
*
*<a href="{@docRoot}/uncopyright.html">This code is in the public domain.</a>
*
* @see WWWAuthenticateList SIPHeader which strings these together.
 */

type WWWAuthenticate struct {
	Authentication
}

/**
 * Default Constructor.
 */
func NewWWWAuthenticate() *WWWAuthenticate {
	this := &WWWAuthenticate{}
	this.Authentication.super(core.SIPHeaderNames_WWW_AUTHENTICATE)
	return this
}
