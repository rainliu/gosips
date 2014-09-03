package rtp

import "fmt"

type RTCPPacketType uint8

/** Identifies the specific kind of RTCP packet. */
const (
	SR      RTCPPacketType = iota /**< An RTCP sender report. */
	RR                            /**< An RTCP receiver report. */
	SDES                          /**< An RTCP source description packet. */
	BYE                           /**< An RTCP bye packet. */
	APP                           /**< An RTCP packet containing application specific data. */
	Unknown                       /**< The type of RTCP packet was not recognized. */
)

/** Base class for specific types of RTCP packets. */
type RTCPPacket struct {
	data        []byte
	datalen     int
	knownformat bool
	packettype  RTCPPacketType
}

func NewRTCPPacket(t RTCPPacketType, d []byte, dlen int) *RTCPPacket {
	this := &RTCPPacket{}
	this.data = make([]byte, dlen)
	this.datalen = dlen
	copy(this.data[:], d[0:dlen])
	this.packettype = t
	this.knownformat = false
	return this
}

/** Returns \c true if the subclass was able to interpret the data and \c false otherwise. */
func (this *RTCPPacket) IsKnownFormat() bool {
	return this.knownformat
}

/** Returns the actual packet type which the subclass implements. */
func (this *RTCPPacket) GetPacketType() RTCPPacketType {
	return this.packettype
}

/** Returns a pointer to the data of this RTCP packet. */
func (this *RTCPPacket) GetPacketData() []byte {
	return this.data
}

/** Returns the length of this RTCP packet. */
func (this *RTCPPacket) GetPacketLength() int {
	return this.datalen
}

func (this *RTCPPacket) Dump() {
	switch this.packettype {
	case SR:
		fmt.Printf("RTCP Sender Report      ")
	case RR:
		fmt.Printf("RTCP Receiver Report    ")
	case SDES:
		fmt.Printf("RTCP Source Description ")
	case APP:
		fmt.Printf("RTCP APP Packet         ")
	case BYE:
		fmt.Printf("RTCP Bye Packet         ")
	case Unknown:
		fmt.Printf("Unknown RTCP Packet     ")
	default:
		fmt.Printf("ERROR: Invalid packet type!")
	}
	fmt.Printf(" Length: %d\n", len(this.data))
}
