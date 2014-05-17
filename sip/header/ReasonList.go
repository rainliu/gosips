package header

import "gosips/core"

/**
* List of Reason headers.
 */
type ReasonList struct {
	SIPHeaderList
}

/** Default constructor
 */
func NewReasonList() *ReasonList {
	this := &ReasonList{}
	this.SIPHeaderList.super(core.SIPHeaderNames_REASON)
	return this
}
