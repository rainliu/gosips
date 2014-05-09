package header

import "gosips/core"

/**
* RecordRoute List of SIP headers (a collection of Addresses)
*
*@version  JAIN-SIP-1.1
*
*@author M. Ranganathan <mranga@nist.gov>  <br/>
*
*<a href="{@docRoot}/uncopyright.html">This code is in the public domain.</a>
*
 */
type RecordRouteList struct {
	SIPHeaderList
}

/** Default constructor
 */
func NewRecordRouteList() *RecordRouteList {
	this := &RecordRouteList{}
	this.SIPHeaderList.super(core.SIPHeaderNames_RECORD_ROUTE)
	return this
}
