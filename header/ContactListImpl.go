package header

import (
	"gosip/core"
	//"gosip/address"
)

/**
* List of contact headers.ContactLists are also maintained in a hashtable for quick lookup.
*/
type ContactListImpl struct{
 	SIPHeaderListImpl
} 	
    
        /** constructor
         * @param hl SIPObjectList
         */        
	/*public ContactList(SIPObjectList hl) {
		super(hl,CONTACT);
	}*/

	/**
	* Constructor. 
	*/
	func NewContactListImpl() *ContactListImpl {
		this := &ContactListImpl{}
	    this.SIPHeaderListImpl.super( core.SIPHeaderNames_CONTACT);
	    return this;
		// Set the headerlist field in our superclass.
	}
	
	func (this *ContactListImpl) super(hname string) {
		this.SIPHeaderListImpl.super(hname);
	}

	/**
         * add a new contact header. Store it in the hashtable also
         * @param contact -- contact to add to this list.
         * @throws IllegalArgumentException if Duplicate Contact for same addr
         */
	func (this *ContactListImpl) AddContact(contact ContactHeader) (IllegalArgumentException error){
		// Concatenate my lists.
		this.SIPHeaderListImpl.PushBack(contact);
		return nil;
	}


	/**
        * make a clone of this contact list.
        * @return Object cloned list.
        */
	func (this *ContactListImpl) Clone() interface{} {
		retval := NewContactListImpl();
		for c:= this.Front(); c != nil; c = c.Next() {
		 	newc := c.Value.(ContactHeader).Clone();
			retval.PushBack(newc);
		}
		return retval;
	}

	
	
	/**
         * Get an array of contact addresses.
  	 *
         * @return  array of contacts.
	 *
         */
	func (this *ContactListImpl) GetContacts() []ContactHeader{
	   retval := make([]ContactHeader, this.Len());

	   i := 0;
	   for e:=this.Front(); e!=nil; e=e.Next() {
		nextContact := e.Value.(ContactHeader);
		retval[i] = nextContact;
		i ++;
	   }
	   return retval;
	}
	
