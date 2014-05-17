package header

import (
	"gosips/core"
)

/**
* Error Info sip header.
*@see ErrorInfoList
*<pre>
*
* 6.24 Error-Info
*
*   The Error-Info response header provides a pointer to additional
*   information about the error status response. This header field is
*   only contained in 3xx, 4xx, 5xx and 6xx responses.
*
*
*
*       Error-Info  =  "Error-Info" ":" # ( "<" URI ">" *( ";" generic-param ))
*</pre>
*
 */
type ErrorInfoList struct {
	SIPHeaderList
}

/** Default constructor.
 */
func NewErrorInfoList() *ErrorInfoList {
	this := &ErrorInfoList{}
	this.SIPHeaderList.super(core.SIPHeaderNames_ERROR_INFO)
	return this
}
