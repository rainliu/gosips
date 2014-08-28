package rtp

import (
	"fmt"
	"net"
)

/** Represents an IPv4 IP address and port.
 *  This class is used by the UDP over IPv4 transmission component.
 *  When an RTPIPv4Address is used in one of the multicast functions of the transmitter, the port
 *  number is ignored. When an instance is used in one of the accept or ignore functions of the
 *  transmitter, a zero port number represents all ports for the specified IP address.
 */
type IPAddress struct {
	ip   net.IP
	port uint16
}

/** Creates an instance with IP address \c ip and port number \c port (\c port is interpreted in host byte order). */
func NewIPAddress(ip net.IP, port uint16) *IPAddress {
	return &IPAddress{ip: ip, port: port}
}

/** Sets the IP address of this instance to \c ip. */
func (this *IPAddress) SetIP(ip net.IP) {
	this.ip = ip
}

/** Sets the port number for this instance to \c port which is interpreted in host byte order. */
func (this *IPAddress) SetPort(port uint16) {
	this.port = port
}

/** Returns the IP address contained in this instance in host byte order. */
func (this *IPAddress) GetIP() net.IP {
	return this.ip
}

/** Returns the port number of this instance in host byte order. */
func (this *IPAddress) GetPort() uint16 {
	return this.port
}

/////////////////////////////////
func (this *IPAddress) GetAddressType() AddressType {
	if len(this.ip) == 4 {
		return IPv4AddressType
	} else {
		return IPv6AddressType
	}
}

func (this *IPAddress) IsSameAddress(addr RTPAddress) bool {
	if addr2, ok := addr.(*IPAddress); ok {
		if this.ip.Equal(addr2.ip) && addr2.port == this.port {
			return true
		}
	}
	return false
}

func (this *IPAddress) IsFromSameHost(addr RTPAddress) bool {
	if addr2, ok := addr.(*IPAddress); ok {
		if this.ip.Equal(addr2.ip) {
			return true
		}
	}
	return false
}

func (this *IPAddress) Clone() RTPAddress {
	return &IPAddress{ip: this.ip, port: this.port}
}

func (this *IPAddress) String() string {
	return fmt.Sprintf("%s:%d", this.ip.String(), this.port)
}
