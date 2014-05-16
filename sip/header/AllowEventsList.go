package header

import (
	"container/list"
	"errors"
	"gosips/core"
)

/**
* List of AllowEvents headers.
* The sip message can have multiple AllowEvents headers which are strung
* together in a list.
 */
type AllowEventsList struct {
	SIPHeaderList
}

/** default constructor
 */
func NewAllowEventsList() *AllowEventsList {
	this := &AllowEventsList{}
	this.SIPHeaderList.super(core.SIPHeaderNames_ALLOW_EVENTS)
	return this
}

/**
 * Gets an Iterator of all the methods of the AllowEventsHeader. Returns an empty
 *
 * Iterator if no methods are defined in this AllowEvents Header.
 *
 *
 *
 * @return Iterator of String objects each identifing the methods of
 *
 * AllowEventsHeader.
 *
 */
func (this *AllowEventsList) GetMethods() *list.List {
	ll := list.New()
	for e := this.Front(); e != nil; e = e.Next() {
		ll.PushBack(e.Value.(*AllowEvents).GetEventType())
	}
	return ll
}

/**
 * Sets the methods supported defined by this AllowEventsHeader.
 *
 *
 *
 * @param methods - the Iterator of Strings defining the methods supported
 *
 * in this AllowEventsHeader
 *
 * @throws ParseException which signals that an error has been reached
 *
 * unexpectedly while parsing the Strings defining the methods supported.
 *
 */

func (this *AllowEventsList) SetMethods(methods *list.List) (ParseException error) {
	for e := methods.Front(); e != nil; e = e.Next() {
		allowEvents := NewAllowEvents()
		if str, ok := e.Value.(string); ok {
			if ParseException = allowEvents.SetEventType(str); ParseException != nil {
				return ParseException
			}
		} else {
			return errors.New("ParseException: the eventType parameter is not string")
		}

		this.PushBack(allowEvents)
	}

	return nil
}
