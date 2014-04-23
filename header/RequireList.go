package header

import "gosip/core"

/**
* List of Require headers.
* <pre>
* Require  =  "Require" ":" 1#option-tag
* </pre>
*
*@version  JAIN-SIP-1.1
*
*@author M. Ranganathan <mranga@nist.gov>  <br/>
*
*<a href="{@docRoot}/uncopyright.html">This code is in the public domain.</a>
*
*
 */
type RequireList struct {
	SIPHeaderList
}

/** Default constructor
 */
func NewRequireList() *RequireList {
	this := &RequireList{}
	this.SIPHeaderList.super(core.SIPHeaderNames_REQUIRE)
	return this
}

/** Constructor
 * @param sip SIPObjectList to set
 */
// public RequireList (SIPObjectList sip) {
// 	super(sip, RequireHeader.NAME);
// }
