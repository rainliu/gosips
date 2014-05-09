package header

import (
	"gosips/core"
)

/**
* Accept List of  SIP headers.
*@see Accept
 */
type AcceptList struct {
	SIPHeaderList
}

/** default constructor
 */
func NewAcceptList() *AcceptList {
	this := &AcceptList{}
	this.SIPHeaderList.super(core.SIPHeaderNames_ACCEPT)
	return this
}
