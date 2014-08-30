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

const HEADER_V_MSK = 0x3
const HEADER_V_POS = 0
const HEADER_P_MSK = 0x1
const HEADER_P_POS = 2
const HEADER_X_MSK = 0x1
const HEADER_X_POS = 3
const HEADER_CC_MSK = 0xF
const HEADER_CC_POS = 4
const HEADER_M_MSK = 0x1
const HEADER_M_POS = 8
const HEADER_PT_MSK = 0x7F
const HEADER_PT_POS = 9

const SIZEOF_HEADER = 12         //12 bytes or 3 dwords
const SIZEOF_EXTENSIONHEADER = 4 //4 bytes or 1 dwords

type Header struct {
	version     uint8 //:2;
	padding     uint8 //:1;
	extension   uint8 //:1;
	csrccount   uint8 //:4;
	marker      uint8 //:1;
	payloadtype uint8 //:7;

	sequencenumber uint16
	timestamp      uint32
	ssrc           uint32
}

func NewHeader(packetbytes []byte) *Header {
	this := &Header{}
	if err := this.Parse(packetbytes); err != nil {
		return nil
	}
	return this
}

func (this *Header) Parse(packetbytes []byte) error {
	if len(packetbytes) < SIZEOF_HEADER {
		return errors.New("ERR_RTP_PACKET_INVALIDPACKET")
	}

	this.version = (packetbytes[0] >> HEADER_V_POS) & HEADER_V_MSK
	this.padding = (packetbytes[0] >> HEADER_P_POS) & HEADER_P_MSK
	this.extension = (packetbytes[0] >> HEADER_X_POS) & HEADER_X_MSK
	this.csrccount = (packetbytes[0] >> HEADER_CC_POS) & HEADER_CC_MSK
	this.marker = (packetbytes[1] >> HEADER_M_POS) & HEADER_M_MSK
	this.payloadtype = (packetbytes[1] >> HEADER_PT_POS) & HEADER_PT_MSK

	this.sequencenumber = uint16(packetbytes[2])<<8 | uint16(packetbytes[3])
	this.timestamp = uint32(packetbytes[4])<<24 | uint32(packetbytes[5])<<16 | uint32(packetbytes[6])<<8 | uint32(packetbytes[7])
	this.ssrc = uint32(packetbytes[8])<<24 | uint32(packetbytes[9])<<16 | uint32(packetbytes[10])<<8 | uint32(packetbytes[11])

	return nil
}

func (this *Header) Encode() []byte {
	var packetbytes []byte
	packetbytes = make([]byte, SIZEOF_HEADER)

	packetbytes[0] = ((this.version & HEADER_V_MSK) << HEADER_V_POS) |
		((this.padding & HEADER_P_MSK) << HEADER_P_POS) |
		((this.extension & HEADER_X_MSK) << HEADER_X_POS) |
		((this.csrccount & HEADER_CC_MSK) << HEADER_CC_POS)

	packetbytes[1] = ((this.marker & HEADER_M_MSK) << HEADER_M_POS) |
		((this.payloadtype & HEADER_PT_MSK) << HEADER_PT_POS)

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

	return packetbytes
}

type ExtensionHeader struct {
	extid  uint16
	length uint16
}

func NewExtensionHeader(packetbytes []byte) *ExtensionHeader {
	this := &ExtensionHeader{}
	if err := this.Parse(packetbytes); err != nil {
		return nil
	}
	return this
}

func (this *ExtensionHeader) Parse(packetbytes []byte) error {
	if len(packetbytes) < 4 {
		return errors.New("ERR_RTP_PACKET_INVALIDPACKET")
	}

	this.extid = uint16(packetbytes[0])<<8 | uint16(packetbytes[1])
	this.length = uint16(packetbytes[2])<<8 | uint16(packetbytes[3])

	return nil
}

func (this *ExtensionHeader) Encode() []byte {
	var packetbytes []byte
	packetbytes = make([]byte, SIZEOF_EXTENSIONHEADER)

	packetbytes[0] = byte((this.extid >> 8) & 0xFF)
	packetbytes[1] = byte((this.extid >> 0) & 0xFF)
	packetbytes[2] = byte((this.length >> 8) & 0xFF)
	packetbytes[3] = byte((this.length >> 0) & 0xFF)

	return packetbytes
}

type SourceIdentifier struct {
	ssrc uint32
}
