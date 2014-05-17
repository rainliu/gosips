package header

import (
	"gosips/core"
)

/**
* List of contact headers.ContactLists are also maintained in a hashtable for quick lookup.
 */
type ContactList struct {
	SIPHeaderList
}

/**
* Constructor.
 */
func NewContactList() *ContactList {
	this := &ContactList{}
	this.SIPHeaderList.super(core.SIPHeaderNames_CONTACT)
	return this
}

func (this *ContactList) super(hname string) {
	this.SIPHeaderList.super(hname)
}

/**
 * add a new contact header. Store it in the hashtable also
 * @param contact -- contact to add to this list.
 * @throws IllegalArgumentException if Duplicate Contact for same addr
 */
func (this *ContactList) AddContact(contact ContactHeader) (IllegalArgumentException error) {
	// Concatenate my lists.
	this.SIPHeaderList.PushBack(contact)
	return nil
}

/**
 * make a clone of this contact list.
 * @return Object cloned list.
 */
func (this *ContactList) Clone() interface{} {
	retval := NewContactList()
	for c := this.Front(); c != nil; c = c.Next() {
		newc := c.Value.(ContactHeader).Clone().(*Contact)
		retval.PushBack(newc)
	}
	return retval
}

/**
         * Get an array of contact addresses.
  	 *
         * @return  array of contacts.
	 *
*/
func (this *ContactList) GetContacts() []ContactHeader {
	retval := make([]ContactHeader, this.Len())

	i := 0
	for e := this.Front(); e != nil; e = e.Next() {
		nextContact := e.Value.(ContactHeader)
		retval[i] = nextContact
		i++
	}
	return retval
}
