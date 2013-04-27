package header

import (
	"bytes"
	"errors"
	"strings"
	"gosip/core"
	"gosip/address"
)



/**  
*To SIP Header.
*/

type To struct{ //  implements javax.sip.header.ToHeader {
	AddressParametersHeaderImpl
}
        /** default Constructor.
         */
    func NewTo() *To{
    	this := &To{};
		this.AddressParametersHeaderImpl.super(core.SIPHeaderNames_TO);
		return this;
    }
    
    func (this *To) super(name string) {
    	this.AddressParametersHeaderImpl.super(name);
    }

	/** Generate a TO header from a FROM header
	*/
     func (this *To) CloneFrom (from *From) {
		this.super(core.SIPHeaderNames_TO);
		this.SetAddress(from.addr);
		this.SetParameters(from.parameters);
     }
    
    /**
     * Compare two To headers for equality.
     * @param otherHeader Object to set
     * @return true if the two headers are the same.
     */
    /*public boolean equals(Object otherHeader) {
	try {
          if (address==null) return false;
          if (!otherHeader.getClass().equals(this.getClass())){
	      return false;
           }

          To otherTo = (To) otherHeader;
          if (! otherTo.getAddress().equals( address )) {
	      return false;
          }
	  return true;
	  // exitpoint = 3;
	} finally {
	    // System.out.println("equals " + retval + exitpoint);
	}
    }*/

   /**
    * Encode the header into a String.
    * @since 1.0
    * @return String
    */
    func (this *To) String() string {
        return this.headerName + core.SIPSeparatorNames_COLON + core.SIPSeparatorNames_SP + this.EncodeBody() + core.SIPSeparatorNames_NEWLINE;
    }

   /**
    * Encode the header content into a String.
    * @return String
    */
    func (this *To) EncodeBody() string {
        if this.addr==nil{
         	return "";
		}
		var retval bytes.Buffer;
		addr,_:=this.addr.(*address.AddressImpl);
        if addr.GetAddressType() == address.ADDRESS_SPEC {
            retval.WriteString(core.SIPSeparatorNames_LESS_THAN);
        }
        retval.WriteString(this.addr.String());
        if addr.GetAddressType() == address.ADDRESS_SPEC {
            retval.WriteString(core.SIPSeparatorNames_GREATER_THAN);
        }

        if this.parameters.Len()>0 {
            retval.WriteString(core.SIPSeparatorNames_SEMICOLON);
            retval.WriteString(this.parameters.String());
        }
        return retval.String();
    }

   /**
    * Conveniance accessor function to get the hostPort field from the address.
    * Warning -- this assumes that the embedded URI is a SipURL.
    *
    * @return hostport field
    */
    func (this *To) GetHostPort() (*core.HostPort, error) {
        if this.addr == nil { 
        	return nil, errors.New("Address is nil");
        }
        addr,_:=this.addr.(*address.AddressImpl);
        return addr.GetHostPort();
    }

   /**
    * Get the display name from the address.
    * @return Display name
    */
    func (this *To) GetDisplayName() string {
        if this.addr == nil { 
        	return "";
        }
        return this.addr.GetDisplayName();
    }

   /**
    * Get the tag parameter from the address parm list.
    * @return tag field
    */
    func (this *To) GetTag() string {
        if this.parameters == nil{
         	return "";
        }
        return this.GetParameter(ParameterNames_TAG);
    }

    /** Boolean function
     * @return true if the Tag exist
     */
    func (this *To) HasTag() bool {
    	if this.parameters == nil{
         	return false;
        }
        return this.HasParameter(ParameterNames_TAG);
    }
    
      /** remove Tag member
       */
    func (this *To) RemoveTag() {
        if this.parameters != nil{
        	this.parameters.Delete(ParameterNames_TAG);
    	}
    }   

   /**
    * Set the tag member. This should be set to null for the initial request
    * in a dialog.
    * @param t tag String to set.
    */
    func (this *To) SetTag( t string) (ParseException error) {
        if t == "" {
         	return errors.New("NullPointerException: null tag ");
		}else if strings.TrimSpace(t)==""{
		 	return errors.New("ParseException: bad tag");
		}
		/*if (LogWriter.needsLogging) {
			LogWriter.logMessage("To:setTag " + t);
			LogWriter.logStackTrace();
		}*/
        this.SetParameter(ParameterNames_TAG,t);
        return nil;
    }    

   /** Get the user@host port string.
    */
    func (this *To) GetUserAtHostPort() string{
        if this.addr==nil{ 
        	return "";
        }
        addr,_:=this.addr.(*address.AddressImpl);
		return addr.GetUserAtHostPort();
    }       

    /** Gets a string representation of the Header. This method overrides the
     * toString method in java.lang.Object.
     *
     * @return string representation of Header
     */
    /*public String String() {
        return this.encode();
    }*/