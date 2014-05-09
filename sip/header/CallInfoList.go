package header

import (
	"gosips/core"
)

/**
* A list of CallInfo headers (there can be multiple in a message).
*
*@author M. Ranganathan <mranga@nist.gov>  <br/>
*
*<a href="${docRoot}/uncopyright.html">This code is in the public domain.</a>
 */
type CallInfoList struct {
	SIPHeaderList
}

/** Default constructor
 */

func NewCallInfoList() *CallInfoList {
	this := &CallInfoList{}
	this.SIPHeaderList.super(core.SIPHeaderNames_CALL_INFO)
	return this
}
