package header

import "gosips/core"

/**
* Proxy Require SIPSIPObject (list of option tags)
*
*@version  JAIN-SIP-1.1
*
*@author M. Ranganathan <mranga@nist.gov>  <br/>
*
*<a href="{@docRoot}/uncopyright.html">This code is in the public domain.</a>
*
 */
type ProxyRequireList struct {
	SIPHeaderList
}

/** Default Constructor
 */
func NewProxyRequireList() *ProxyRequireList {
	this := &ProxyRequireList{}
	this.SIPHeaderList.super(core.SIPHeaderNames_PROXY_REQUIRE)
	return this
}

/** Constructor
 * @param sip SIPObjectList to set
 */
// public ProxyRequireList (SIPObjectList sip) {
// 	super(sip, PROXY_REQUIRE);
// }
