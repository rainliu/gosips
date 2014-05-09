package header

import (
	"gosips/core"
)

/**
* In-Reply-To SIP header. Keeps a list of CallIdentifiers
*
*@version  JAIN-SIP-1.1
*
*@author M. Ranganathan <mranga@nist.gov>  <br/>
*
*<a href="{@docRoot}/uncopyright.html">This code is in the public domain.</a>
*
 */
type InReplyToList struct {
	SIPHeaderList
}

/** Default constructor
 */
func NewInReplyToList() *InReplyToList {
	this := &InReplyToList{}
	this.SIPHeaderList.super(core.SIPHeaderNames_IN_REPLY_TO)
	return this
}
