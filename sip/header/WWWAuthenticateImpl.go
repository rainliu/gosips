package header

import "gosips/core"

/** The WWWAuthenticate SIP header.
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
