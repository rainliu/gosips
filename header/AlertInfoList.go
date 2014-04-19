package header

import (
	"gosip/core"
)

/**
* AlertInfo SIPHeader - there can be several AlertInfo headers.
*
*@author M. Ranganathan <mranga@nist.gov>  <br/>
*
*<a href="{@docRoot}/uncopyright.html">This code is in the public domain.</a>
*
 */
type AlertInfoList struct {
	SIPHeaderList
}

/** default constructor
 */
func NewAlertInfoList() *AlertInfoList {
	this := &AlertInfoList{}
	this.SIPHeaderList.super(core.SIPHeaderNames_ALERT_INFO)
	return this
}
