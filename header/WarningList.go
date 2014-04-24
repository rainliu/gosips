package header

import "gosip/core"

/**
* A Warning SIPObject. (A list of Warning headers).
*
*@version  JAIN-SIP-1.1
*
*@author M. Ranganathan <mranga@nist.gov>  <br/>
*
*<a href="{@docRoot}/uncopyright.html">This code is in the public domain.</a>
*
 */
type WarningList struct {
	SIPHeaderList
}

/** Constructor.
 *
 */
func NewWarningList() *WarningList {
	this := &WarningList{}
	this.SIPHeaderList.super(core.SIPHeaderNames_WARNING)
	return this
}
