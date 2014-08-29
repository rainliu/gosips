package rtp

import "net"

const TRANS_HASHSIZE = 8317
const TRANS_DEFAULTPORTBASE = 5000
const TRANS_RTPRECEIVEBUFFER = 32768
const TRANS_RTCPRECEIVEBUFFER = 32768
const TRANS_RTPTRANSMITBUFFER = 32768
const TRANS_RTCPTRANSMITBUFFER = 32768

/** Base class for transmission parameters.
 *  This class is an abstract class which will have a specific implementation for a
 *  specific kind of transmission component. All actual implementations inherit the
 *  GetTransmissionProtocol function which identifies the component type for which
 *  these parameters are valid.
 */

/** Parameters for the UDP over IPv4 transmitter. */
type TransmissionParams struct {
	protocol TransmissionProtocol

	portbase                 uint16
	bindIP                   net.IP
	mcastifidx               uint32
	localIPs                 []net.IP
	multicastTTL             uint8
	rtpsendbuf, rtprecvbuf   int
	rtcpsendbuf, rtcprecvbuf int
}

func NewTransmissionParams() *TransmissionParams {
	this := &TransmissionParams{}
	this.multicastTTL = 1
	this.mcastifidx = 0
	this.portbase = TRANS_DEFAULTPORTBASE
	this.rtpsendbuf = TRANS_RTPTRANSMITBUFFER
	this.rtprecvbuf = TRANS_RTPRECEIVEBUFFER
	this.rtcpsendbuf = TRANS_RTCPTRANSMITBUFFER
	this.rtcprecvbuf = TRANS_RTCPRECEIVEBUFFER
	return this
}
func (this *TransmissionParams) SetTransmissionProtocol(protocol TransmissionProtocol) {
	this.protocol = protocol
}

func (this *TransmissionParams) GetTransmissionProtocol() TransmissionProtocol {
	return this.protocol
}

/** Sets the IP address which is used to bind the sockets to \c ip. */
func (this *TransmissionParams) SetBindIP(ip net.IP) {
	this.bindIP = ip
}

/** Sets the multicast interface index. */
func (this *TransmissionParams) SetMulticastInterfaceIndex(idx uint32) {
	this.mcastifidx = idx

}

/** Sets the RTP portbase to \c pbase. This has to be an even number. */
func (this *TransmissionParams) SetPortbase(pbase uint16) {
	this.portbase = pbase
}

/** Sets the multicast TTL to be used to \c mcastTTL. */
func (this *TransmissionParams) SetMulticastTTL(mcastTTL uint8) {
	this.multicastTTL = mcastTTL
}

/** Passes a list of IP addresses which will be used as the local IP addresses. */
func (this *TransmissionParams) SetLocalIPList(iplist []net.IP) {
	this.localIPs = iplist
}

/** Sets the RTP socket's send buffer size. */
func (this *TransmissionParams) SetRTPSendBuffer(s int) {
	this.rtpsendbuf = s
}

/** Sets the RTP socket's receive buffer size. */
func (this *TransmissionParams) SetRTPReceiveBuffer(s int) {
	this.rtprecvbuf = s
}

/** Sets the RTCP socket's send buffer size. */
func (this *TransmissionParams) SetRTCPSendBuffer(s int) {
	this.rtcpsendbuf = s
}

/** Sets the RTCP socket's receive buffer size. */
func (this *TransmissionParams) SetRTCPReceiveBuffer(s int) {
	this.rtcprecvbuf = s
}

/** Clears the list of local IP addresses.
 *  Clears the list of local IP addresses. An empty list will make the transmission component
 *  itself determine the local IP addresses.
 */
// func (this *TransmissionParams)  ClearLocalIPList()												{
// 	this.localIPs.clear();
// }

/** Returns the IP address which will be used to bind the sockets. */
func (this *TransmissionParams) GetBindIP() net.IP {
	return this.bindIP
}

/** Returns the multicast interface index. */
func (this *TransmissionParams) GetMulticastInterfaceIndex() uint32 {
	return this.mcastifidx
}

/** Returns the RTP portbase which will be used (default is 5000). */
func (this *TransmissionParams) GetPortbase() uint16 {
	return this.portbase
}

/** Returns the multicast TTL which will be used (default is 1). */
func (this *TransmissionParams) GetMulticastTTL() uint8 {
	return this.multicastTTL
}

/** Returns the list of local IP addresses. */
func (this *TransmissionParams) GetLocalIPList() []net.IP {
	return this.localIPs
}

/** Returns the RTP socket's send buffer size. */
func (this *TransmissionParams) GetRTPSendBuffer() int {
	return this.rtpsendbuf
}

/** Returns the RTP socket's receive buffer size. */
func (this *TransmissionParams) GetRTPReceiveBuffer() int {
	return this.rtprecvbuf
}

/** Returns the RTCP socket's send buffer size. */
func (this *TransmissionParams) GetRTCPSendBuffer() int {
	return this.rtcpsendbuf
}

/** Returns the RTCP socket's receive buffer size. */
func (this *TransmissionParams) GetRTCPReceiveBuffer() int {
	return this.rtcprecvbuf
}
