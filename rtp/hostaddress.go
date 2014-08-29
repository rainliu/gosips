package rtp

import "fmt"

/** A very general kind of address consisting of a port number and a number of bytes describing the host address. */
type HostAddress struct {
	host string
	port uint16
}

/** Creates an instance of the class using \c addrlen bytes of \c hostaddress as host identification,
 *  and using \c port as the port number. */
func NewHostAddress(host string, port uint16) *HostAddress {
	return &HostAddress{host: host, port: port}
}

/** Sets the host address to the first \c addrlen bytes of \c hostaddress. */
func (this *HostAddress) SetHost(host string) {
	this.host = host
}

/** Sets the port number to \c port. */
func (this *HostAddress) SetPort(port uint16) {
	this.port = port
}

/** Returns a pointer to the stored host address. */
func (this *HostAddress) GetHost() string {
	return this.host
}

/** Returns the port number stored in this instance. */
func (this *HostAddress) GetPort() uint16 {
	return this.port
}

/////////////////////////////////
func (this *HostAddress) GetAddressType() AddressType {
	return HostAddressType
}

func (this *HostAddress) IsSameAddress(addr Address) bool {
	if addr2, ok := addr.(*HostAddress); ok {
		if addr2.host == this.host && addr2.port == this.port {
			return true
		}
	}
	return false
}

func (this *HostAddress) IsFromSameHost(addr Address) bool {
	if addr2, ok := addr.(*HostAddress); ok {
		if addr2.host == this.host {
			return true
		}
	}
	return false
}

func (this *HostAddress) Clone() Address {
	return &HostAddress{host: this.host, port: this.port}
}

func (this *HostAddress) String() string {
	return fmt.Sprintf("%s:%d", this.host, this.port)
}
