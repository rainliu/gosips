package core

import (
<<<<<<< HEAD
	"net"
)


=======
    "bytes"
    "net"
    "strconv"
)

>>>>>>> update code
/**
* Holds the hostname:port.
*<a href="{@docRoot}/uncopyright.html">This code is in the public domain.</a>
*
<<<<<<< HEAD
*/

type HostPort struct{
	GenericObjectImpl
    
	// host / ipv4/ ipv6/ 
        /** host field
         */    
	host Host;

        /** port field
         *
         */        
    port int;
}
	     
        /** Default constructor
         */        
	/*public HostPort() {
           
		host = null;
		port = null; // marker for not set.
	}*/
        
        /**
         * Encode this hostport into its string representation.
         * Note that this could be different from the string that has
         * been parsed if something has been edited.
         * @return String
         */
	func (this *HostPort) Encode() string {
		var retval string;//= new StringBuffer();
		retval.append(this.host.Encode());
		if this.port != -1{ 
		    retval.append(COLON).append(port.toString());
		}
		return retval;
	}
        
        /** returns true if the two objects are equals, false otherwise.
         * @param other Object to set
         * @return boolean
         */        
        /*public boolean equals(Object other) {
=======
 */

type HostPort struct {
    //GenericObjectImpl

    // host / ipv4/ ipv6/
    /** host field
     */
    host *Host

    /** port field
     *
     */
    port int
}

/** Default constructor
 */
func NewHostPort() *HostPort {
	return &HostPort{}; // marker for not set.
}

/**
 * Encode this hostport into its string representation.
 * Note that this could be different from the string that has
 * been parsed if something has been edited.
 * @return String
 */
func (this *HostPort) String() string {
    var retval bytes.Buffer //= new StringBuffer();
    if this.host!=nil {
        retval.WriteString(this.host.String())
        if this.port != -1 {
            retval.WriteString(Separators_COLON+strconv.Itoa(this.port))
        }
    }
    return retval.String()
}

/** returns true if the two objects are equals, false otherwise.
 * @param other Object to set
 * @return boolean
 */
/*public boolean equals(Object other) {
>>>>>>> update code
            if (! this.getClass().equals(other.getClass())) {
                return false;
            }
            HostPort that = (HostPort) other;
	    if ( (this.port == null && that.port != null) ||
		 (this.port != null && that.port == null) ) return false;
<<<<<<< HEAD
	    else if (this.port == that.port && this.host.equals(that.host)) 
		return true;
	    else 
              return this.host.equals(that.host) && this.port.equals(that.port);
        }*/
   
        /** get the Host field
         * @return host field
         */        
		func (this *HostPort) GetHost() Host{ 
            return this.host ;
        } 
                
        /** get the port field
         * @return int
         */        
		func (this *HostPort) GetPort() int {
			return this.port;
        }
        
       /**
        * Returns boolean value indicating if Header has port
        * @return boolean value indicating if Header has port
        */
        func (this *HostPort) HasPort() bool {
            return this.port != -1;
        }
        
        /** remove port.
         */        
        func (this *HostPort) RemovePort() {
            this.port = -1;
        }
        
	/**
         * Set the host member
         * @param h Host to set
         */
		func (this *HostPort) SetHost(h Host) {
            this.host = h;
        }
        
	/**
         * Set the port member
         * @param p int to set
         */
		func (this *HostPort) SetPort(p int) {
	   		// -1 is same as remove port.
			this.port = p
        } 
	
        /** Return the internet address corresponding to the host.
         *@throws java.net.UnkownHostException if host name cannot be resolved.
         *@return the inet address for the host.
         */
        func (this *HostPort) GetInetAddress() net.IP{
            if this.host == "" {
            	return null;
            }
            return net.ParseIP(this.host);
        }


	func (this *HostPort) Clone() interface{}{
	 	retval := &HostPort{};
		retval.host = this.host.Clone().(*Host);
		retval.port = this.port;
		return retval;
	}
=======
	    else if (this.port == that.port && this.host.equals(that.host))
		return true;
	    else
              return this.host.equals(that.host) && this.port.equals(that.port);
        }*/

/** get the Host field
 * @return host field
 */
func (this *HostPort) GetHost() *Host {
    return this.host
}

/** get the port field
 * @return int
 */
func (this *HostPort) GetPort() int {
    return this.port
}

/**
 * Returns boolean value indicating if Header has port
 * @return boolean value indicating if Header has port
 */
func (this *HostPort) HasPort() bool {
    return this.port != -1
}

/** remove port.
 */
func (this *HostPort) RemovePort() {
    this.port = -1
}

/**
 * Set the host member
 * @param h Host to set
 */
func (this *HostPort) SetHost(h *Host) {
    this.host = h
}

/**
 * Set the port member
 * @param p int to set
 */
func (this *HostPort) SetPort(p int) {
    // -1 is same as remove port.
    this.port = p
}

/** Return the internet address corresponding to the host.
 *@throws java.net.UnkownHostException if host name cannot be resolved.
 *@return the inet address for the host.
 */
func (this *HostPort) GetInetAddress() net.IP {
    if this.host == nil {
        return nil
    }
    return net.ParseIP(this.host.GetHostName())
}

func (this *HostPort) Clone() interface{} {
    retval := &HostPort{}
    if this.host!=nil {
        retval.host = this.host.Clone().(*Host)
    }else{
        retval.host = nil;
    }
    retval.port = this.port
    return retval
}
>>>>>>> update code
