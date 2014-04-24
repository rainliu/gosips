package header

import "gosip/core"

/**
* List of Reason headers.
*@version  JAIN-SIP-1.1
*
*@author M. Ranganathan <mranga@nist.gov>  <br/>
*
*<a href="{@docRoot}/uncopyright.html">This code is in the public domain.</a>
*
*
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

/** Constructor
 * @param sip SIPObjectList to set
 */
// public ReasonList (SIPObjectList sip) {
// 	super(sip, ReasonHeader.NAME);
// }
