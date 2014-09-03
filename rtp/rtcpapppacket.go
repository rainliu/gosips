package rtp

import "fmt"

/** Describes an RTCP APP packet. */
type RTCPAPPPacket struct {
	RTCPPacket
	appdatalen int
}

/** Creates an instance based on the data in \c data with length \c datalen.
 *  Creates an instance based on the data in \c data with length \c datalen. Since the \c data pointer
 *  is referenced inside the class (no copy of the data is made) one must make sure that the memory it
 *  points to is valid as long as the class instance exists.
 */
func NewRTCPAPPPacket(data []byte, datalen int) *RTCPAPPPacket {
	this := &RTCPAPPPacket{}
	this.data = make([]byte, datalen)
	this.datalen = datalen
	copy(this.data[:], data[0:datalen])
	this.packettype = APP
	this.knownformat = false

	this.appdatalen = datalen
	if ((this.data[0] >> RTCP_HEADER_P_POS) & RTCP_HEADER_P_MSK) != 0 {
		padcount := this.data[datalen-1]
		if (padcount & 0x03) != 0 { // not a multiple of four! (see rfc 3550 p 37)
			return this
		}
		if int(padcount) >= this.appdatalen {
			return this
		}
		this.appdatalen -= int(padcount)
	}

	if this.appdatalen < SIZEOF_RTCPHEADER+4*2 {
		return this
	}
	this.appdatalen -= SIZEOF_RTCPHEADER + 4*2
	this.knownformat = true

	return this
}

/** Returns the subtype contained in the APP packet. */
func (this *RTCPAPPPacket) GetSubType() uint8 {
	if !this.knownformat {
		return 0
	}

	count := (this.data[0] >> RTCP_HEADER_C_POS) & RTCP_HEADER_C_MSK
	//RTCPCommonHeader *hdr = (RTCPCommonHeader *)data;
	return count
}

/** Returns the SSRC of the source which sent this packet. */
func (this *RTCPAPPPacket) GetSSRC() uint32 {
	if !this.knownformat {
		return 0
	}

	ssrc := uint32(this.data[SIZEOF_RTCPHEADER+0])<<24 |
		uint32(this.data[SIZEOF_RTCPHEADER+1])<<16 |
		uint32(this.data[SIZEOF_RTCPHEADER+2])<<8 |
		uint32(this.data[SIZEOF_RTCPHEADER+3])<<0
	return ssrc //ntohl(*ssrc);
}

/** Returns the name contained in the APP packet.
 *  Returns the name contained in the APP packet. This alway consists of four bytes and is not NULL-terminated.
 */
func (this *RTCPAPPPacket) GetName() []byte {
	if !this.knownformat {
		return nil
	}

	return this.data[SIZEOF_RTCPHEADER+4:] //(data+sizeof(RTCPCommonHeader)+sizeof(uint32_t));
}

/** Returns a pointer to the actual data. */
func (this *RTCPAPPPacket) GetAPPData() []byte {
	if !this.knownformat {
		return nil
	}
	if this.appdatalen == 0 {
		return nil
	}
	return this.data[SIZEOF_RTCPHEADER+4*2:]
}

/** Returns the length of the actual data. */
func (this *RTCPAPPPacket) GetAPPDataLength() int {
	if !this.knownformat {
		return 0
	}
	return this.appdatalen
}

func (this *RTCPAPPPacket) Dump() {
	this.RTCPPacket.Dump()
	if !this.IsKnownFormat() {
		fmt.Printf("    Unknown format!")
	} else {
		fmt.Printf("    SSRC:   %d", this.GetSSRC())
		fmt.Printf("    Name:   %s", this.GetName())
		fmt.Printf("    Length: %d", this.GetAPPDataLength())
	}
}
