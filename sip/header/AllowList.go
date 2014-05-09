package header

import (
	"container/list"
	"gosips/core"
)

/**
* List of ALLOW headers. The sip message can have multiple Allow headers
*
 */
type AllowList struct {
	SIPHeaderList
}

/** default constructor
 */
func NewAllowList() *AllowList {
	this := &AllowList{}
	this.SIPHeaderList.super(core.SIPHeaderNames_ALLOW)
	return this
}

/**
 * Gets an Iterator of all the methods of the AllowHeader. Returns an empty
 *
 * Iterator if no methods are defined in this Allow Header.
 *
 *
 *
 * @return Iterator of String objects each identifing the methods of
 *
 * AllowHeader.
 *
 * @since JAIN SIP v1.1
 *
 */

func (this *AllowList) GetMethods() *list.List {
	ll := list.New()
	for e := this.Front(); e != nil; e = e.Next() {
		ll.PushBack(e.Value.(*Allow).GetMethod())
	}
	return ll
}

/**
 * Sets the methods supported defined by this AllowHeader.
 *
 *
 *
 * @param methods - the Iterator of Strings defining the methods supported
 *
 * in this AllowHeader
 *
 * @throws ParseException which signals that an error has been reached
 *
 * unexpectedly while parsing the Strings defining the methods supported.
 *
 * @since JAIN SIP v1.1
 *
 */

func (this *AllowList) SetMethods(methods *list.List) { // throws ParseException {

	for e := methods.Front(); e != nil; e = e.Next() {
		allow := NewAllow()
		allow.SetMethod(e.Value.(string))
		this.PushBack(allow)
	}
}
