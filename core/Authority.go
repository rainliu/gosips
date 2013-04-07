package core

import (

)


/**
* Authority part of a URI structure. Section 3.2.2 RFC2396
*
*<a href="{@docRoot}/uncopyright.html">This code is in the public domain.</a>
*
*/
type Authority struct{
	//NetObject
    
        /** hostport field
         */    
	hostPort *HostPort;

        /** userInfo field
         */        
    userInfo *UserInfo;
}         	
	func NewAuthority() *Authority{
		return &Authority{};
	}
	
	func (this *Authority) Clone() interface{} {
		retval := &Authority{};
		
		if this.hostPort!=nil{
			retval.hostPort = this.hostPort.Clone().(*HostPort);
		}
		if this.userInfo!=nil{
			retval.userInfo = this.userInfo.Clone().(*UserInfo);
		}
		
		return retval;
	} 
        /**
         * Return the host name in encoded form.
         * @return encoded string (does the same thing as toString)
         */        
	func (this *Authority) String() string {	
		if this.userInfo != nil && this.hostPort !=nil {
	       	  return this.userInfo.String() + Separators_AT + this.hostPort.String();
		} else if this.hostPort !=nil{
			return this.hostPort.String();
		} else if this.userInfo !=nil {
			return this.userInfo.String() + Separators_AT;
		}
		
		return "";
	}
        
        /** retruns true if the two Objects are equals , false otherwise.
         * @param other Object to test.
         * @return boolean
         */        
        /*public boolean equals(Object other) {
            if (!other.getClass().getName().equals(this.getClass().getName())){
                return false;
            }
            Authority otherAuth = (Authority) other;
            if (! this.hostPort.equals(otherAuth.hostPort) ) {
                return false;
            }
	    if (this.userInfo != null && otherAuth.userInfo != null) {
		if (! this.userInfo.equals(otherAuth.userInfo)) {
		    return false;
		}
	    }
            return true;
        }*/
        
        /**
         * get the hostPort member.
         * @return HostPort
         */
		func (this *Authority) GetHostPort() *HostPort { 
            return this.hostPort ;
        } 
                
        /**
         * get the userInfo memnber.
         * @return UserInfo
         */
		func (this *Authority) GetUserInfo() *UserInfo { 
            return this.userInfo;
        }
        
	/**
         * Get password from the user info.
         * @return String
         */
		func (this *Authority) GetPassword() string { 
	    	if this.userInfo == nil {
	    		return "";
            }//else{
            	return this.userInfo.GetPassword() ;
        	//}
        }
               
        /**
         * Get the user name if it exists.
         * @return String user or null if not set.
         */
        func (this *Authority) GetUser() string{ 
        	if this.userInfo == nil {
	    		return "";
            }//else{
            	return this.userInfo.GetUser() ;
        	//}
        }
        
         /**
          * Get the host name.
          * @return Host (null if not set)
          */
         func (this *Authority) GetHost() *Host { 
	    	if this.hostPort == nil {
	    		return nil;
            }//else{
             	return this.hostPort.GetHost();
            //}
         }                      
                           
          /**
           * Get the port.
           * @return int port (-1) if port is not set.
           */
		func (this *Authority) GetPort() int { 
	    	if this.hostPort == nil{ 
	    		return -1;
            }//else{
             	return this.hostPort.GetPort();
            //}
        }
              
        /** remove the port.
         */        
        func (this *Authority) RemovePort() {
	    	if this.hostPort != nil{
	    		this.hostPort.RemovePort();
	    	}
        }
        
        /**
         * set the password.
         * @param passwd String to set
         */
        func (this *Authority) SetPassword(passwd string) {
	    	if this.userInfo == nil{
	    		this.userInfo = &UserInfo{};
	    	}
            this.userInfo.SetPassword(passwd);
        }
         
         /**
          * Set the user name of the userInfo member.
          * @param user String to set
          */
        func (this *Authority) SetUser(user string) {
	    	if this.userInfo == nil{
	    	 	this.userInfo = &UserInfo{};
            }
            this.userInfo.SetUser(user);
        }

          /**
           * set the host.
           * @param host Host to set
           */
         func (this *Authority) SetHost(host *Host) { 
	     	if this.hostPort == nil{
	     		this.hostPort = &HostPort{};
            }
            this.hostPort.SetHost(host);
         }
        
           /**
            * Set the port.
            * @param port int to set
            */
        func (this *Authority) SetPort (port int) {
	    	if this.hostPort == nil{
	     		this.hostPort = &HostPort{};
            }
            this.hostPort.SetPort(port);
        }
                    
	/**
         * Set the hostPort member
         * @param h HostPort to set
         */
		func (this *Authority) SetHostPort(h *HostPort) { 
            this.hostPort = h ;
        }
        
	/**
         * Set the userInfo member
         * @param u UserInfo to set
         */
		func (this *Authority) SetUserInfo(u *UserInfo) { 
            this.userInfo = u ;
        } 

	/** Remove the user Infor.
	*
	*/
		func (this *Authority) RemoveUserInfo() {
	   		this.userInfo = nil;
        }