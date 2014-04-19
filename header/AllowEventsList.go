package header

import (
	"container/list"
	"gosip/core"
)

/**
* List of AllowEvents headers.
* The sip message can have multiple AllowEvents headers which are strung
* together in a list.
*
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
 * @since JAIN SIP v1.1
 *
 */

func (this *AllowEventsList) GetMethods() *list.List {
	//li := this.super.hlist.listIterator();
	ll := list.New()
	for e := this.Front(); e != nil; e = e.Next() {
		ll.PushBack(e.Value.(*AllowEvents).GetEventType())
	}
	/* while (li.hasNext()) {
	       AllowEvents allowEvents = (AllowEvents) li.next();
	       ll.add(allowEvents.getEventType());
	   }
	   return ll.listIterator();*/
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
 * @since JAIN SIP v1.1
 *
 */

func (this *AllowEventsList) SetMethods(methods *list.List) { // throws ParseException {
	/*ListIterator it = methods.listIterator();
	  while (it.hasNext()) {
	      AllowEvents allowEvents = new AllowEvents();
	      allowEvents.setEventType((String) it.next());
	      this.add(allowEvents);
	  }*/
	for e := methods.Front(); e != nil; e = e.Next() {
		allowEvents := NewAllowEvents()
		allowEvents.SetEventType(e.Value.(string))
		this.PushBack(allowEvents)
	}
}
