package core

import (
<<<<<<< HEAD
	"strings"
	"net"
)


=======
    "net"
    "strings"
)

>>>>>>> update code
/**
 * Stores hostname.
 *<a href="{@docRoot}/uncopyright.html">This code is in the public domain.</a>
 *
 */
<<<<<<< HEAD
 
 const (
    HOSTNAME = iota //0:	1;
    IPV4ADDRESS 	//1: 	2;
    IPV6ADDRESS 	//2:	3;
 )
 
type Host struct{
	GenericObjectImpl
	
    /** hostName field
     */
    hostname string;
=======

const (
    HOSTNAME    = iota //0:	1;
    IPV4ADDRESS        //1: 	2;
    IPV6ADDRESS        //2:	3;
)

type Host struct {
    //GenericObjectImpl

    /** hostName field
     */
    hostname string
>>>>>>> update code

    /** address field
     */

<<<<<<< HEAD
    addressType int;

    inetAddress net.IP;
}

    /** default constructor
     */
    /*public Host() {
        addressType = HOSTNAME;
    }*/

    /** Constructor given host name or IP address.
     */
    func (this *Host) NewHost(hostName string) *Host {
        if hostName == ""{
            return nil;//throw new IllegalArgumentException("null host name");
        }
        
        this.hostname = hostName;
        if this.isIPv6Address(hostName) {
            this.addressType = IPV6ADDRESS;
        }
        this.addressType = IPV4ADDRESS;
    }

    /** constructor
     * @param name String to set
     * @param addrType int to set
     */
    func (this *Host) NewHost2(name string, addrType int) *Host {
        this.addressType = addrType;
        this.hostname = strings.ToLower(strings.TrimSpace(name));
    }

    /**
     * Return the host name in encoded form.
     * @return String
     */
    func (this *Host) Encode() string {
        if this.addressType == IPV6ADDRESS && !this.isIPv6Reference(this.hostname) {
            return "[" + this.hostname + "]";
        }
        return this.hostname;
    }

    /**
     * Compare for equality of hosts.
     * Host names are compared by textual equality. No dns lookup
     * is performed.
     * @param obj Object to set
     * @return boolean
     */
    /*func (this *Host) equals(Object obj) bool{
        if (!this.getClass().equals(obj.getClass())) {
            return false;
        }
        Host otherHost = (Host) obj;
        return otherHost.hostname.equals(hostname);

    }*/

    /** get the HostName field
     * @return String
     */
    func (this *Host) GetHostname() string {
        return this.hostname;
    }

    /** get the Address field
     * @return String
     */
    func (this *Host) GetAddress() string{
        return this.hostname;
    }

    /**
     * Convenience function to get the raw IP destination address
     * of a SIP message as a String.
     * @return String
     */
    func (this *Host) GetIpAddress() string {
        var rawIpAddress string;
        if this.hostname == nil {
        	return null;
        }
        
        if this.addressType == HOSTNAME {
            //try {
                if this.inetAddress == nil {
                    this.inetAddress = net.ParseIP(this.hostname);
                }
                rawIpAddress = this.inetAddress.String();//getHostAddress();
            //} catch (UnknownHostException ex) {
            //    dbgPrint("Could not resolve hostname " + ex);
            //}
        } else {
            rawIpAddress = hostname;
        }
        return rawIpAddress;
    }

    /**
     * Set the hostname member.
     * @param h String to set
     */
    func (this *Host) SetHostname(h string) {
        inetAddress = nil;
        if this.isIPv6Address(h){
            this.addressType = IPV6ADDRESS;
        }else{
            this.addressType = HOSTNAME;
        }    
// Null check bug fix sent in by jpaulo@ipb.pt
        if h != ""{
        	this.hostname = strings.ToLower(strings.TrimSpace(h));
		}
    }

    /** Set the IP Address.
     *@param address is the address string to set.
     */
    func (this *Host) SetHostAddress(address string) {
        this.inetAddress = nil;
        if this.isIPv6Address(address){
            this.addressType = IPV6ADDRESS;
        }else{
            this.addressType = IPV4ADDRESS;
        if address != "" {
        	this.hostname = strings.TrimSpace(address);
    	}
    }

    /**
     * Set the address member
     * @param address address String to set
     */
    func (this *Host) SetAddress(address string){
        this.SetHostAddress(address);
    }

    /** Return true if the address is a DNS host name
     *  (and not an IPV4 address)
     *@return true if the hostname is a DNS name
     */
    func (this *Host) IsHostname() bool {
        return addressType == HOSTNAME;
    }

    /** Return true if the address is a DNS host name
     *  (and not an IPV4 address)
     *@return true if the hostname is host address.
     */
    func (this *Host) IsIPAddress() bool{
        return addressType != HOSTNAME;
    }

    /** Get the inet address from this host.
     * Caches the inet address returned from dns lookup to avoid
     * lookup delays.
     *
     *@throws UnkownHostexception when the host name cannot be resolved.
     */
    func (this *Host) GetInetAddress() net.IP {
        if this.hostname == "" {
        	return nil;
        }
        if this.inetAddress != nil {
        	return this.inetAddress;
        }
        inetAddress = net.ParseIP(this.hostname);
        return inetAddress;

    }

    //----- IPv6
    /**
     * Verifies whether the <code>address</code> could
     * be an IPv6 address
     */
    func (this *Host) isIPv6Address(address string) bool {
        return address != "" && strings.Index(address, ":") != -1;
    }

    /**
     * Verifies whether the ipv6reference, i.e. whether it enclosed in
     * square brackets
     */
    func (this *Host) isIPv6Reference(address string) bool{
        return address[0] == '[' && address[len(address)-1] == ']';
    }

    func (this *Host) Clone() interface{} {
		retval := &Host{};
		retval.addressType = this.addressType;
		retval.hostname = this.hostname;
		return retval;
    }
=======
    addressType int

    inetAddress net.IP
}

/** default constructor
 */
/*public Host() {
    addressType = HOSTNAME;
}*/

/** Constructor given host name or IP address.
 */
func NewHost(hname string) *Host {
    if hname == "" {
        return nil //throw new IllegalArgumentException("null host name");
    }

    this := &Host{};

    this.hostname = hname
    if this.isIPv6Address(hname) {
        this.addressType = IPV6ADDRESS
    }
    this.addressType = IPV4ADDRESS

    return this;
}

/** constructor
 * @param name String to set
 * @param addrType int to set
 */
/*func NewHost2(name string, addrType int) *Host {
    this := &Host{};
    this.addressType = addrType
    this.hostname = strings.ToLower(strings.TrimSpace(name))
    return this;
}*/

/**
 * Return the host name in encoded form.
 * @return String
 */
func (this *Host) String() string {
    if this.addressType == IPV6ADDRESS && !this.isIPv6Reference(this.hostname) {
        return "[" + this.hostname + "]"
    }
    return this.hostname
}

/**
 * Compare for equality of hosts.
 * Host names are compared by textual equality. No dns lookup
 * is performed.
 * @param obj Object to set
 * @return boolean
 */
/*func (this *Host) equals(Object obj) bool{
    if (!this.getClass().equals(obj.getClass())) {
        return false;
    }
    Host otherHost = (Host) obj;
    return otherHost.hostname.equals(hostname);

}*/

/** get the HostName field
 * @return String
 */
func (this *Host) GetHostName() string {
    return this.hostname
}

/** get the Address field
 * @return String
 */
func (this *Host) GetAddress() string {
    return this.hostname
}

/**
 * Convenience function to get the raw IP destination address
 * of a SIP message as a String.
 * @return String
 */
func (this *Host) GetIpAddress() string {
    var rawIpAddress string
    if this.hostname == "" {
        return ""
    }

    if this.addressType == HOSTNAME {
        //try {
        if this.inetAddress == nil {
            this.inetAddress = net.ParseIP(this.hostname)
        }
        rawIpAddress = this.inetAddress.String() //getHostAddress();
        //} catch (UnknownHostException ex) {
        //    dbgPrint("Could not resolve hostname " + ex);
        //}
    } else {
        rawIpAddress = this.hostname
    }
    return rawIpAddress
}

/**
 * Set the hostname member.
 * @param h String to set
 */
func (this *Host) SetHostName(hname string) {
    this.inetAddress = nil
    if this.isIPv6Address(hname) {
        this.addressType = IPV6ADDRESS
    } else {
        this.addressType = HOSTNAME
    }
    // Null check bug fix sent in by jpaulo@ipb.pt
    if hname != "" {
        this.hostname = strings.ToLower(strings.TrimSpace(hname))
    }
}

/** Set the IP Address.
 *@param address is the address string to set.
 */
func (this *Host) SetHostAddress(address string) {
    this.inetAddress = nil
    if this.isIPv6Address(address) {
        this.addressType = IPV6ADDRESS
    } else {
        this.addressType = IPV4ADDRESS
    }
    if address != "" {
        this.hostname = strings.TrimSpace(address)
    }
}

/**
 * Set the address member
 * @param address address String to set
 */
func (this *Host) SetAddress(address string) {
    this.SetHostAddress(address)
}

/** Return true if the address is a DNS host name
 *  (and not an IPV4 address)
 *@return true if the hostname is a DNS name
 */
func (this *Host) IsHostName() bool {
    return this.addressType == HOSTNAME
}

/** Return true if the address is a DNS host name
 *  (and not an IPV4 address)
 *@return true if the hostname is host address.
 */
func (this *Host) IsIPAddress() bool {
    return this.addressType != HOSTNAME
}

/** Get the inet address from this host.
 * Caches the inet address returned from dns lookup to avoid
 * lookup delays.
 *
 *@throws UnkownHostexception when the host name cannot be resolved.
 */
func (this *Host) GetInetAddress() net.IP {
    if this.hostname == "" {
        return nil
    }
    if this.inetAddress != nil {
        return this.inetAddress
    }
    this.inetAddress = net.ParseIP(this.hostname)
    return this.inetAddress

}

//----- IPv6
/**
 * Verifies whether the <code>address</code> could
 * be an IPv6 address
 */
func (this *Host) isIPv6Address(address string) bool {
    return address != "" && strings.Index(address, ":") != -1
}

/**
 * Verifies whether the ipv6reference, i.e. whether it enclosed in
 * square brackets
 */
func (this *Host) isIPv6Reference(address string) bool {
    return address[0] == '[' && address[len(address)-1] == ']'
}

func (this *Host) Clone() interface{} {
    retval := &Host{}
    retval.addressType = this.addressType
    retval.hostname = this.hostname
    return retval
}
>>>>>>> update code
