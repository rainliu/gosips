package header

import (
	"fmt"
	"errors"
	"bytes"
	"strings"
	"strconv"
	"gosip/core"
	"gosip/address"
)

/**
* Contact Item. 
*/
type ContactImpl struct{ //implements javax.sip.header.ContactHeader
 	AddressParametersHeaderImpl    

	// This must be private or the toString will go for a loop! 
	contactList *ContactListImpl;
   	
        /** wildCardFlag field.
         */        
    wildCardFlag bool;
}
   
        /** Default constructor.
         */        
	func NewContactImpl() *ContactImpl {
		this := &ContactImpl{};
		
		this.AddressParametersHeaderImpl.super(core.SIPHeaderNames_CONTACT);
		this.contactList = NewContactListImpl();
		
		return this;
	}
	
	func (this *ContactImpl) super(hname string){
		this.AddressParametersHeaderImpl.super(hname);
		this.contactList = NewContactListImpl();
	}

	/** Set a parameter.
	*/
	func (this *ContactImpl) SetParameter( name,  value string) (ParseException error){
	   nv := this.parameters.GetNameValue(name);
	   if nv != nil {
		 nv.SetValue(value);
	   } else {
          nv  = core.NewNameValue(name,value);
	      if strings.ToLower(name)=="methods" {
			nv.SetQuotedValue();
		  }
          this.parameters.AddNameValue(nv);
	   }
	   return nil;
	}
        
	/**
         * Encode body of the header into a cannonical String.
         * @return string encoding of the header value.
         */
	func (this *ContactImpl) EncodeBody() string {
		var encoding bytes.Buffer;//= new StringBuffer();
		if this.wildCardFlag  {
			encoding.WriteString("*")
			return encoding.String();
		}
		addr,_:=this.addr.(*address.AddressImpl);
		// Bug report by Joao Paulo
		if addr.GetAddressType() == address.NAME_ADDR {
		    encoding.WriteString(this.addr.String());
		} else  {
		    // Encoding in canonical form must have <> around address.
		    encoding.WriteString("<");
			encoding.WriteString(this.addr.String())
			encoding.WriteString(">");
		}
		if this.parameters.Len()>0 {
		    encoding.WriteString(core.SIPSeparatorNames_SEMICOLON)
		    encoding.WriteString(this.parameters.String());
		} 

		return encoding.String();
	}
		
        /** get the Contact list.
         * @return ContactList
         */        
	func (this *ContactImpl) GetContactList() *ContactListImpl {
		return this.contactList;
	}


        /** get the WildCardFlag field
         * @return boolean
         */        
	func (this *ContactImpl) GetWildCardFlag() bool {
            return this.wildCardFlag;
        } 
        
        /** get the address field.
         * @return Address
         */        
	func (this *ContactImpl) GetAddress() address.Address{
	    // JAIN-SIP stores the wild card as an address!
	    return this.addr;
    } 

        /** get the parameters List
         * @return NameValueList
         */        
	func (this *ContactImpl) GetContactParms() *core.NameValueList { 
            return this.parameters ;
        } 
		
	
	
        /** get Expires parameter.
         * @return the Expires parameter.
         */        
	func (this *ContactImpl) GetExpires() int {
		retval,_ := strconv.Atoi(this.GetParameter(core.SIPHeaderNames_EXPIRES));
		return retval;
	}
	
	/** Set the expiry time in seconds.
	*@param expiryDeltaSeconds exipry time.
	*/
	
	func (this *ContactImpl) SetExpires( expiryDeltaSeconds int) (InvalidArgumentException error){
		//Integer deltaSeconds = new Integer(expiryDeltaSeconds);
		this.parameters.AddNameValue(core.NewNameValue(core.SIPHeaderNames_EXPIRES, strconv.Itoa(expiryDeltaSeconds))) ;
		return nil;
	}
        
        /** get the Q-value
         * @return float
         */        
    func (this *ContactImpl) GetQValue() float32{
		qvalue,_ := strconv.ParseFloat(this.GetParameter(core.SIPParameters_Q), 32);
		return float32(qvalue);
	}
	
     
	
        
        /** set the Contact List
         * @param cl ContactList to set
         */        
	func (this *ContactImpl) SetContactList( cl *ContactListImpl ) {
		this.contactList = cl;
	}
        
	
	/**
         * Set the wildCardFlag member
         * @param w boolean to set
         */
	func (this *ContactImpl) SetWildCardFlag( w bool) {
		this.wildCardFlag = true;
		addr := address.NewAddressImpl();
		addr.SetWildCardFlag();
		this.SetAddress(addr);
   }
        
	/**
         * Set the address member
         *
         * @param address Address to set
         */
	func (this *ContactImpl) SetAddress(addr address.Address) {
	    // Canonical form must have <> around the address.
	    this.AddressParametersHeaderImpl.SetAddress(addr);
	    this.wildCardFlag = false;
    }
        
        
		
        
        /** set the Q-value parameter
         * @param qValue float to set
         */        
	func (this *ContactImpl) SetQValue(qValue float32) (InvalidArgumentException error){
		if qValue!=-1 && (qValue<0||qValue>1) {
			return errors.New("JAIN-SIP Exception, Contact, setQValue(), the qValue is not between 0 and 1");
		}else{
			qValueStr:= fmt.Sprintf("%f", qValue); 
			this.parameters.AddNameValue(core.NewNameValue(core.SIPParameters_Q, qValueStr)) ;
    		return nil;
    	}
    }
        
        
      