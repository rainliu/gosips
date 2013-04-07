package core

import (
	"strings"
)


/**
* User information part of a URL. 
*<a href="{@docRoot}/uncopyright.html">This code is in the public domain.</a>
*/
/** Constant field
         */        
const (
	TELEPHONE_SUBSCRIBER = iota //1 ;
	USER //= 2; 
)

type UserInfo struct{
	//NetObject
    
        /** user field
         */    
		user string;

        /** password field
         */        
        password string;
        
		/** userType field
         */        
        userType int;
}

        /** Default constructor
         */        
        func NewUserInfo() *UserInfo { 
        	return &UserInfo{};
        }
              
        /**
         * Compare for equality.
         * @param obj Object to set
         * @return true if the two headers are equals, false otherwise.
         */
        /*public boolean equals(Object obj) {
            if (! getClass().getName().equals(obj.getClass().getName())) {
                return false;
            }
            UserInfo other = (UserInfo) obj;
            if (this.userType != other.userType) {
                return false;
            }
            if (! this.user.equalsIgnoreCase(other.user)) {
                return false;
            }
            if (this.password != null &&
                other.password == null)  return false;
            
            if (other.password != null && this.password == null) return false;
            
	    if (this.password == other.password ) return true;

            return (this.password.equals(other.password));
        }*/
    
    func (this *UserInfo) Clone() interface{} {
    	retval := &UserInfo{}
    	
    	retval.user = this.user;
        retval.password = this.password;        
        retval.userType = this.userType;
        
    	return retval;
    }
        
        /**
         * Encode the user information as a string.
         * @return String
         */
	func (this *UserInfo) String() string {
		if this.password != ""{
			return this.user + Separators_COLON + this.password;
		}//else{
		 	return this.user;
		//}
	}

	/** Clear the password field.
	*/
	func (this *UserInfo) ClearPassword() {
		this.password = "";
	}
        
        /**
         * Gets the user type (which can be set to TELEPHONE_SUBSCRIBER or USER)
         * @return the type of user.
         */
	func (this *UserInfo) GetUserType() int {
		return this.userType;
	}

        /** get the user field.
         * @return String
         */        
	 	func (this *UserInfo) GetUser() string { 
            return this.user ;
        } 

        /** get the password field.
         * @return String
         */        
		func (this *UserInfo) GetPassword() string { 
            return this.password;
        } 

	/**
         * Set the user member
         * @param user String to set
         */
		func (this *UserInfo) SetUser(user string) { 
            this.user = user ;
     	   // BUG Fix submitted by Lamine Brahimi 
	   // add this (taken form sip_messageParser)
           // otherwise comparison of two SipUrl will fail because this
           // parameter is not set (whereas it is set in sip_messageParser).
            if 	user!= "" && 
            	strings.Index(user, Separators_POUND) >= 0 || 
				strings.Index(user, Separators_SEMICOLON) >= 0  {
                this.SetUserType(TELEPHONE_SUBSCRIBER);
            } else {
            	this.SetUserType(USER);
			}
        } 

	/**
         * Set the password member
         * @param p String to set
         */
		func (this *UserInfo) SetPassword(p string) { 
            this.password = p ;
        }      
	
	/**
         * Set the user type (to TELEPHONE_SUBSCRIBER or USER).
         * @param type int to set
         * @throws IllegalArgumentException if type is not in range.
         */
	func (this *UserInfo) SetUserType(t int){
		if t != TELEPHONE_SUBSCRIBER && t != USER {
		   println("Parameter not in range");
		   return
		}
		this.userType = t;
	}