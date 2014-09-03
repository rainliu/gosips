package rtp

import "errors"

const RTP_VERSION = 2
const RTP_MAXCSRCS = 15
const RTP_MINPACKETSIZE = 600
const RTP_DEFAULTPACKETSIZE = 1400
const RTP_PROBATIONCOUNT = 2
const RTP_MAXPRIVITEMS = 256
const RTP_SENDERTIMEOUTMULTIPLIER = 2
const RTP_BYETIMEOUTMULTIPLIER = 1
const RTP_MEMBERTIMEOUTMULTIPLIER = 5
const RTP_COLLISIONTIMEOUTMULTIPLIER = 10
const RTP_NOTETTIMEOUTMULTIPLIER = 25
const RTP_DEFAULTSESSIONBANDWIDTH = 10000.0

const RTP_RTCPTYPE_SR = 200
const RTP_RTCPTYPE_RR = 201
const RTP_RTCPTYPE_SDES = 202
const RTP_RTCPTYPE_BYE = 203
const RTP_RTCPTYPE_APP = 204

const RTP_HEADER_V_MSK = 0x3
const RTP_HEADER_V_POS = 0
const RTP_HEADER_P_MSK = 0x1
const RTP_HEADER_P_POS = 2
const RTP_HEADER_X_MSK = 0x1
const RTP_HEADER_X_POS = 3
const RTP_HEADER_CC_MSK = 0xF
const RTP_HEADER_CC_POS = 4
const RTP_HEADER_M_MSK = 0x1
const RTP_HEADER_M_POS = 8
const RTP_HEADER_PT_MSK = 0x7F
const RTP_HEADER_PT_POS = 9

const RTCP_HEADER_C_MSK = 0x1F
const RTCP_HEADER_C_POS = 0
const RTCP_HEADER_P_MSK = 0x1
const RTCP_HEADER_P_POS = 5
const RTCP_HEADER_V_MSK = 0x3
const RTCP_HEADER_V_POS = 6

const SIZEOF_RTPHEADER = 12   //12 bytes or 3 dwords
const SIZEOF_RTPEXTENSION = 4 //4 bytes or 1 dwords
const SIZEOF_RTCPHEADER = 4   //4 bytes or 1 dwords

type RTPHeader struct {
	version     uint8 //:2;
	padding     uint8 //:1;
	extension   uint8 //:1;
	csrccount   uint8 //:4;
	marker      uint8 //:1;
	payloadtype uint8 //:7;

	sequencenumber uint16
	timestamp      uint32
	ssrc           uint32
	csrc           []uint32
}

func NewRTPHeader() *RTPHeader {
	this := &RTPHeader{}

	return this
}

func NewRTPHeaderFromBytes(packetbytes []byte) *RTPHeader {
	this := &RTPHeader{}
	if err := this.Parse(packetbytes); err != nil {
		return nil
	}
	return this
}

func (this *RTPHeader) Parse(packetbytes []byte) error {
	if len(packetbytes) < SIZEOF_RTPHEADER {
		return errors.New("ERR_RTP_PACKET_INVALIDPACKET")
	}

	this.version = (packetbytes[0] >> RTP_HEADER_V_POS) & RTP_HEADER_V_MSK
	this.padding = (packetbytes[0] >> RTP_HEADER_P_POS) & RTP_HEADER_P_MSK
	this.extension = (packetbytes[0] >> RTP_HEADER_X_POS) & RTP_HEADER_X_MSK
	this.csrccount = (packetbytes[0] >> RTP_HEADER_CC_POS) & RTP_HEADER_CC_MSK
	this.marker = (packetbytes[1] >> RTP_HEADER_M_POS) & RTP_HEADER_M_MSK
	this.payloadtype = (packetbytes[1] >> RTP_HEADER_PT_POS) & RTP_HEADER_PT_MSK

	this.sequencenumber = uint16(packetbytes[2])<<8 | uint16(packetbytes[3])
	this.timestamp = uint32(packetbytes[4])<<24 | uint32(packetbytes[5])<<16 | uint32(packetbytes[6])<<8 | uint32(packetbytes[7])
	this.ssrc = uint32(packetbytes[8])<<24 | uint32(packetbytes[9])<<16 | uint32(packetbytes[10])<<8 | uint32(packetbytes[11])

	if len(packetbytes) < SIZEOF_RTPHEADER+int(this.csrccount)*4 {
		return errors.New("ERR_RTP_PACKET_INVALIDPACKET")
	}

	this.csrc = make([]uint32, this.csrccount)
	for i := uint8(0); i < this.csrccount; i++ {
		this.csrc[i] = uint32(packetbytes[SIZEOF_RTPHEADER+i*4+0])<<24 |
			uint32(packetbytes[SIZEOF_RTPHEADER+i*4+1])<<16 |
			uint32(packetbytes[SIZEOF_RTPHEADER+i*4+2])<<8 |
			uint32(packetbytes[SIZEOF_RTPHEADER+i*4+3])
	}
	return nil
}

func (this *RTPHeader) Encode() []byte {
	var packetbytes []byte
	packetbytes = make([]byte, SIZEOF_RTPHEADER+int(this.csrccount)*4)

	packetbytes[0] = ((this.version & RTP_HEADER_V_MSK) << RTP_HEADER_V_POS) |
		((this.padding & RTP_HEADER_P_MSK) << RTP_HEADER_P_POS) |
		((this.extension & RTP_HEADER_X_MSK) << RTP_HEADER_X_POS) |
		((this.csrccount & RTP_HEADER_CC_MSK) << RTP_HEADER_CC_POS)

	packetbytes[1] = ((this.marker & RTP_HEADER_M_MSK) << RTP_HEADER_M_POS) |
		((this.payloadtype & RTP_HEADER_PT_MSK) << RTP_HEADER_PT_POS)

	packetbytes[2] = byte((this.sequencenumber >> 8) & 0xFF)
	packetbytes[3] = byte(this.sequencenumber & 0xFF)

	packetbytes[4] = byte((this.timestamp >> 24) & 0xFF)
	packetbytes[5] = byte((this.timestamp >> 16) & 0xFF)
	packetbytes[6] = byte((this.timestamp >> 8) & 0xFF)
	packetbytes[7] = byte((this.timestamp >> 0) & 0xFF)

	packetbytes[8] = byte((this.ssrc >> 24) & 0xFF)
	packetbytes[9] = byte((this.ssrc >> 16) & 0xFF)
	packetbytes[10] = byte((this.ssrc >> 8) & 0xFF)
	packetbytes[11] = byte((this.ssrc >> 0) & 0xFF)

	for i := uint8(0); i < this.csrccount; i++ {
		packetbytes[SIZEOF_RTPHEADER+i*4+0] = byte((this.csrc[i] >> 24) & 0xFF)
		packetbytes[SIZEOF_RTPHEADER+i*4+0] = byte((this.csrc[i] >> 16) & 0xFF)
		packetbytes[SIZEOF_RTPHEADER+i*4+0] = byte((this.csrc[i] >> 8) & 0xFF)
		packetbytes[SIZEOF_RTPHEADER+i*4+0] = byte((this.csrc[i] >> 0) & 0xFF)
	}

	return packetbytes
}

type RTPExtension struct {
	id     uint16
	length uint16
	data   []uint32
}

func NewRTPExtension() *RTPExtension {
	this := &RTPExtension{}

	return this
}

func NewRTPExtensionFromBytes(packetbytes []byte) *RTPExtension {
	this := &RTPExtension{}
	if err := this.Parse(packetbytes); err != nil {
		return nil
	}
	return this
}

func (this *RTPExtension) Parse(packetbytes []byte) error {
	if len(packetbytes) < SIZEOF_RTPEXTENSION {
		return errors.New("ERR_RTP_PACKET_INVALIDPACKET")
	}

	this.id = uint16(packetbytes[0])<<8 | uint16(packetbytes[1])
	this.length = uint16(packetbytes[2])<<8 | uint16(packetbytes[3])

	if len(packetbytes) < SIZEOF_RTPEXTENSION+int(this.length)*4 {
		return errors.New("ERR_RTP_PACKET_INVALIDPACKET")
	}

	for i := uint16(0); i < this.length; i++ {
		this.data[i] = uint32(packetbytes[SIZEOF_RTPEXTENSION+i*4+0])<<24 |
			uint32(packetbytes[SIZEOF_RTPEXTENSION+i*4+1])<<16 |
			uint32(packetbytes[SIZEOF_RTPEXTENSION+i*4+2])<<8 |
			uint32(packetbytes[SIZEOF_RTPEXTENSION+i*4+3])
	}

	return nil
}

func (this *RTPExtension) Encode() []byte {
	var packetbytes []byte
	packetbytes = make([]byte, SIZEOF_RTPEXTENSION+this.length*4)

	packetbytes[0] = byte((this.id >> 8) & 0xFF)
	packetbytes[1] = byte((this.id >> 0) & 0xFF)
	packetbytes[2] = byte((this.length >> 8) & 0xFF)
	packetbytes[3] = byte((this.length >> 0) & 0xFF)

	for i := uint16(0); i < this.length; i++ {
		packetbytes[SIZEOF_RTPEXTENSION+i*4+0] = byte(this.data[i]>>24) & 0xFF
		packetbytes[SIZEOF_RTPEXTENSION+i*4+1] = byte(this.data[i]>>16) & 0xFF
		packetbytes[SIZEOF_RTPEXTENSION+i*4+2] = byte(this.data[i]>>8) & 0xFF
		packetbytes[SIZEOF_RTPEXTENSION+i*4+3] = byte(this.data[i]>>0) & 0xFF
	}

	return packetbytes
}

type SourceIdentifier struct {
	ssrc uint32
}

type RTCPCommonHeader struct {
	count   uint8 //:5;
	padding uint8 //:1;
	version uint8 //:2;

	packettype uint8
	length     uint16
}

type RTCPSenderReport struct {
	ntptime_msw  uint32
	ntptime_lsw  uint32
	rtptimestamp uint32
	packetcount  uint32
	octetcount   uint32
}

type RTCPReceiverReport struct {
	ssrc         uint32 // Identifies about which SSRC's data this report is...
	fractionlost uint8
	packetslost  [3]uint8
	exthighseqnr uint32
	jitter       uint32
	lsr          uint32
	dlsr         uint32
}

type RTCPSDESHeader struct {
	sdesid uint8
	length uint8
}
