package header

import (
	"gosips/core"
)

/**
* In-Reply-To SIP header. Keeps a list of CallIdentifiers
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
