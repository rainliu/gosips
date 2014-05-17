package header

import "gosips/core"

/**
* RecordRoute List of SIP headers (a collection of Addresses)
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
