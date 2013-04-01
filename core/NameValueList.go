package core

import (
	"container/list"
)

/**
* Implements a simple NameValue association with a quick lookup 
* function (via a hash table) this class is not thread safe 
* because it uses HashTables.
*
*@version  JAIN-SIP-1.1
*
*@author M. Ranganathan <mranga@nist.gov>  <br/>
*
*<a href="{@docRoot}/uncopyright.html">This code is in the public domain.</a>
*
*/

type NameValueList struct{
	GenericObjectListImpl
}
	
	func NewNameValueList( listName string ) *NameValueList{
	    this := &NameValueList{};
		
		this.GenericObjectListImpl.listName = listName;
		
		return this;
	}
	
	func (this *NameValueList) Add1(nv *NameValue) {
		if nv == nil {
		 	//throw new NullPointerException("null nv");
		 	return;
		}
		this.PushBack(nv);
	}

	/**
	* Add a name value record to this list.
	*/
	func (this *NameValueList) Add2(name string, value interface{}) {
	 	nv := NewNameValue(name,value);
		this.Add1(nv);
	}

	/**
	* Set a namevalue object in this list.
	*/
	func (this *NameValueList) Set1(nv *NameValue) {
		this.Delete(nv.name);
		this.Add1(nv);
	}

	/**
	* Set a namevalue object in this list.
	*/
	func (this *NameValueList) Set2(name string, value interface{}) {
		nv := NewNameValue(name,value);
		this.Set1(nv);
	}      

	/**
         *  Compare if two NameValue lists are equal.
	 *@param otherObject  is the object to compare to.
	 *@return true if the two objects compare for equality.
         */
    /*public boolean equals(Object otherObject) {
            if (!otherObject.getClass().equals
                (this.getClass())) {
                return false;
            }
            NameValueList other = (NameValueList) otherObject;

            if (this.size() != other.size()) {
		return false;
	    }
	    ListIterator li = this.listIterator();
	    
	    while (li.hasNext()) {
		NameValue nv = (NameValue) li.next();
		boolean found = false;
	        ListIterator li1 = other.listIterator();
		while (li1.hasNext()) {
			NameValue nv1  = (NameValue) li1.next();
			// found a match so break;
			if (nv.equals(nv1))   {
			   found = true;
			   break;
			}
		}
		if (! found ) return false;
	    }
	    return true;
	}*/
	

	/**
	*  Do a lookup on a given name and return value associated with it.
	*/
	func (this *NameValueList) GetValue(name string) interface{} {
	 	nv := this.GetNameValue(name);
		if nv != nil{
			return nv.value;
		}
		
		return nil;
	}

	/**
	* Get the NameValue record given a name.
	* @since 1.0
	*/
	func (this *NameValueList) GetNameValue (name string) *NameValue{
		for e:=this.Front(); e!=nil; e=e.Next(){
			nv := e.Value.(*NameValue);
			if nv.GetName() == name {
				return nv;
			}
		}
		
		return nil;
	}
	
	/**
	* Returns a boolean telling if this NameValueList
	* has a record with this name  
	* @since 1.0
	*/
	func (this *NameValueList) HasNameValue (name string) bool {
		return  this.GetNameValue(name) != nil;
	}

	/**
	* Remove the element corresponding to this name.
	* @since 1.0
	*/
	func (this *NameValueList) Delete(name string) bool{
	    for e:=this.Front(); e!=nil; e=e.Next(){
			nv := e.Value.(*NameValue);
			if nv.GetName() == name {
				this.Remove(e);
				return true;
			}
		}
		
		return false;
	}
        
	 /**
	  *Get a list of parameter names.
	  *@return a list iterator that has the names of the parameters.
	  */
	 func (this *NameValueList) GetNames() *list.List {
		ll := list.New();
		for e:=this.Front(); e!=nil; e=e.Next(){
			nv := e.Value.(*NameValue);
			ll.PushBack(nv.GetName());
		}
		return ll;
	 }


        
        func (this *NameValueList) Clone() interface{}  {
            retval := &NameValueList{};
            retval.indentation = this.indentation;
        	retval.listName = this.listName;
        	retval.stringRep = this.stringRep;
        	retval.separator = this.separator;

            li := list.New();
            for e:=this.Front(); e!=nil; e=e.Next(){
				nv := e.Value.(*NameValue);
				nnv := nv.Clone().(*NameValue);
				li.PushBack(nnv);
			}
            
            return retval;
        }
        
        /** Get the parameter as a String.
         *@return the parameter as a string.
         */
        func (this *NameValueList) GetParameter(name string) string{
            val := this.GetValue(name);
            if val == nil {
            	return "";
            }
            if gv, ok := val.(GenericObject); ok { 
                return gv.Encode();
            } 
            
            return val.(string);
        }