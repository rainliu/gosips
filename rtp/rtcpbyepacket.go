package rtp

import "fmt"

/** Describes an RTCP BYE packet. */
type RTCPBYEPacket struct {
	RTCPPacket
	reasonoffset int
}

/** Creates an instance based on the data in \c data with length \c datalen.
 *  Creates an instance based on the data in \c data with length \c datalen. Since the \c data pointer
 *  is referenced inside the class (no copy of the data is made) one must make sure that the memory it
 *  points to is valid as long as the class instance exists.
 */
func NewRTCPBYEPacket(data []byte, datalen int) *RTCPBYEPacket {
	this := &RTCPBYEPacket{}
	this.data = make([]byte, datalen)
	this.datalen = datalen
	copy(this.data[:], data[0:datalen])
	this.packettype = BYE
	this.knownformat = false

	this.reasonoffset = 0
	byelen := datalen
	if ((this.data[0] >> RTCP_HEADER_P_POS) & RTCP_HEADER_P_MSK) != 0 {
		padcount := data[datalen-1]
		if (padcount & 0x03) != 0 { // not a multiple of four! (see rfc 3550 p 37)
			return this
		}
		if int(padcount) >= byelen {
			return this
		}
		byelen -= int(padcount)
	}

	count := (this.data[0] >> RTCP_HEADER_C_POS) & RTCP_HEADER_C_MSK
	ssrclen := SIZEOF_RTCPHEADER + int(count)*4 //*sizeof(uint32_t) + sizeof(RTCPCommonHeader);
	if ssrclen > byelen {
		return this
	}
	if ssrclen < byelen { // there's probably a reason for leaving
		reaslen := int(data[ssrclen])
		if reaslen > byelen-ssrclen-1 {
			return this
		}
		this.reasonoffset = ssrclen
	}
	this.knownformat = true

	return this
}

/** Returns the number of SSRC identifiers present in this BYE packet. */
func (this *RTCPBYEPacket) GetSSRCCount() int {
	if !this.knownformat {
		return 0
	}

	count := (this.data[0] >> RTCP_HEADER_C_POS) & RTCP_HEADER_C_MSK

	//RTCPCommonHeader *hdr = (RTCPCommonHeader *)data;
	return int(count) //(int)(hdr->count);
}

/** Returns the SSRC described by \c index which may have a value from 0 to GetSSRCCount()-1
 *  (note that no check is performed to see if \c index is valid).
 */
func (this *RTCPBYEPacket) GetSSRC(index int) uint32 {
	if !this.knownformat {
		return 0
	}
	ssrc := uint32(this.data[SIZEOF_RTCPHEADER+4*index+0])<<24 |
		uint32(this.data[SIZEOF_RTCPHEADER+4*index+1])<<16 |
		uint32(this.data[SIZEOF_RTCPHEADER+4*index+2])<<8 |
		uint32(this.data[SIZEOF_RTCPHEADER+4*index+3])<<0
	//uint32_t *ssrc = (uint32_t *)(data+sizeof(RTCPCommonHeader)+sizeof(uint32_t)*index);
	return ssrc //ntohl(*ssrc);
}

/** Returns true if the BYE packet contains a reason for leaving. */
func (this *RTCPBYEPacket) HasReasonForLeaving() bool {
	if !this.knownformat {
		return false
	}
	if this.reasonoffset == 0 {
		return false
	}
	return true
}

/** Returns the length of the string which describes why the source(s) left. */
func (this *RTCPBYEPacket) GetReasonLength() int {
	if !this.knownformat {
		return 0
	}
	if this.reasonoffset == 0 {
		return 0
	}
	reasonlen := int(this.data[this.reasonoffset])
	return reasonlen
}

/** Returns the actual reason for leaving data. */
func (this *RTCPBYEPacket) GetReasonData() []byte {
	if !this.knownformat {
		return nil
	}
	if this.reasonoffset == 0 {
		return nil
	}
	reasonlen := this.data[this.reasonoffset]
	if reasonlen == 0 {
		return nil
	}
	return this.data[this.reasonoffset+1:]
}

func (this *RTCPBYEPacket) Dump() {
	this.RTCPPacket.Dump()
	if !this.IsKnownFormat() {
		fmt.Printf("    Unknown format")
		return
	}

	num := this.GetSSRCCount()

	for i := 0; i < num; i++ {
		fmt.Printf("    SSRC: %d", this.GetSSRC(i))
	}
	if this.HasReasonForLeaving() {
		fmt.Printf("    Reason: %s", this.GetReasonData())
	}
}
