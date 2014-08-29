package rtp

import "net"

type TransmissionProtocol uint8

const (
	IPv4UDPProto     TransmissionProtocol = iota /**< Specifies the internal UDP over IPv4 transmitter. */
	IPv6UDPProto                                 /**< Specifies the internal UDP over IPv6 transmitter. */
	ExternalProto                                /**< Specifies the transmitter which can send packets using an external mechanism, and which can have received packets injected into it - see RTPExternalTransmitter for additional information. */
	UserDefinedProto                             /**< Specifies a user defined, external transmitter. */
)

/** Base class for additional information about the transmitter.
 *  This class is an abstract class which will have a specific implementation for a
 *  specific kind of transmission component. All actual implementations inherit the
 *  GetTransmissionProtocol function which identifies the component type for which
 *  these parameters are valid.
 */

/** Additional information about the UDP over IPv4/IPv6 transmitter. */
type TransmissionInfo struct {
	protocol TransmissionProtocol

	localIPlist           []net.IP
	rtpsocket, rtcpsocket int
}

func NewTransmissionInfo(iplist []net.IP, rtpsock, rtcpsock int) *TransmissionInfo {
	this := &TransmissionInfo{}
	if len(iplist[0]) == 4 {
		this.protocol = IPv4UDPProto
	} else {
		this.protocol = IPv6UDPProto
	}
	this.localIPlist = iplist
	this.rtpsocket = rtpsock
	this.rtcpsocket = rtcpsock
	return this
}

/** Returns the transmitter type for which these parameters are valid. */
func (this *TransmissionInfo) GetTransmissionProtocol() TransmissionProtocol {
	return this.protocol
}

/** Returns the list of IPv4 addresses the transmitter considers to be the local IP addresses. */
func (this *TransmissionInfo) GetLocalIPList() []net.IP {
	return this.localIPlist
}

/** Returns the socket descriptor used for receiving and transmitting RTP packets. */
func (this *TransmissionInfo) GetRTPSocket() int {
	return this.rtpsocket
}

/** Returns the socket descriptor used for receiving and transmitting RTCP packets. */
func (this *TransmissionInfo) GetRTCPSocket() int {
	return this.rtcpsocket
}
