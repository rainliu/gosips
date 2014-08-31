package rtp

import "errors"

/** Represents an RTP Packet.
 *  The RTPPacket class can be used to parse a RTPRawPacket instance if it represents RTP data.
 *  The class can also be used to create a new RTP packet according to the parameters specified by
 *  the user.
 */
type Packet struct {
	receivetime *RTPTime
	header      *Header
	extension   *Extension
	payload     []byte

	packet []byte
}

/** Creates an RTPPacket instance based upon the data in \c rawpack, optionally installing a memory manager.
 *  If successful, the data is moved from the raw packet to the RTPPacket instance.
 */
func NewPacketFromRawPacket(rawpack *RawPacket) *Packet {
	this := &Packet{}
	this.receivetime = rawpack.GetReceiveTime().Clone()
	if err := this.ParseRawPacket(rawpack); err != nil {
		return nil
	}
	return this
}

func (this *Packet) ParseRawPacket(rawpack *RawPacket) error {
	if !rawpack.IsRTP() { // If we didn't receive it on the RTP port, we'll ignore it
		return errors.New("ERR_RTP_PACKET_INVALIDPACKET")
	}

	this.packet = make([]byte, len(rawpack.GetData()))
	copy(this.packet, rawpack.GetData())

	this.header = NewHeader()
	if err := this.header.Parse(this.packet); err != nil {
		return err
	}

	// The version number should be correct
	if this.header.version != RTP_VERSION {
		return errors.New("ERR_RTP_PACKET_INVALIDPACKET")
	}

	// We'll check if this is possibly a RTCP packet. For this to be possible
	// the marker bit and payload type combined should be either an SR or RR
	// identifier
	if this.header.marker != 0 {
		if this.header.payloadtype == (RTP_RTCPTYPE_SR & 127) { // don't check high bit (this was the marker!!)
			return errors.New("ERR_RTP_PACKET_INVALIDPACKET")
		}
		if this.header.payloadtype == (RTP_RTCPTYPE_RR & 127) {
			return errors.New("ERR_RTP_PACKET_INVALIDPACKET")
		}
	}

	var numpadbytes, payloadoffset, payloadlength int

	payloadoffset = SIZEOF_HEADER + 4*int(this.header.csrccount)
	if this.header.extension != 0 { // got header extension
		this.extension = NewExtension()
		if err := this.extension.Parse(this.packet[payloadoffset:]); err != nil {
			return err
		}
		payloadoffset += SIZEOF_EXTENSION + 4*int(this.extension.length)
	} else {
		this.extension = nil
	}

	if this.header.padding != 0 { // adjust payload length to take padding into account
		numpadbytes = int(this.packet[len(this.packet)-1]) // last byte contains number of padding bytes
		if numpadbytes > len(this.packet)-payloadoffset {
			return errors.New("ERR_RTP_PACKET_INVALIDPACKET")
		}
	} else {
		numpadbytes = 0
	}

	payloadlength = len(this.packet) - numpadbytes - payloadoffset
	if payloadlength < 0 {
		return errors.New("ERR_RTP_PACKET_INVALIDPACKET")
	}

	return nil
}

/** Creates a new buffer for an RTP packet and fills in the fields according to the specified parameters.
 *  If \c maxpacksize is not equal to zero, an error is generated if the total packet size would exceed
 *  \c maxpacksize. The arguments of the constructor are self-explanatory. Note that the size of a header
 *  extension is specified in a number of 32-bit words. A memory manager can be installed.
 *  This constructor is similar to the other constructor, but here data is stored in an external buffer
 *  \c buffer with size \c buffersize. */
func NewPacket(payloadtype uint8,
	payloaddata []byte,
	seqnr uint16,
	timestamp uint32,
	ssrc uint32,
	gotmarker bool,
	numcsrcs uint8,
	csrcs []uint32,
	gotextension bool,
	extensionid uint16,
	extensionlen uint16,
	extensiondata []uint32) *Packet {
	this := &Packet{}

	this.receivetime = &RTPTime{0, 0}
	if err := this.BuildPacket(payloadtype,
		payloaddata,
		seqnr,
		timestamp,
		ssrc,
		gotmarker,
		numcsrcs,
		csrcs,
		gotextension,
		extensionid,
		extensionlen,
		extensiondata); err != nil {
		return nil
	}

	return this
}
func (this *Packet) BuildPacket(payloadtype uint8,
	payloaddata []byte,
	seqnr uint16,
	timestamp uint32,
	ssrc uint32,
	gotmarker bool,
	numcsrcs uint8,
	csrcs []uint32,
	gotextension bool,
	extensionid uint16,
	extensionlen uint16,
	extensiondata []uint32) error {
	if numcsrcs > RTP_MAXCSRCS {
		return errors.New("ERR_RTP_PACKET_TOOMANYCSRCS")
	}

	if payloadtype > 127 { // high bit should not be used
		return errors.New("ERR_RTP_PACKET_BADPAYLOADTYPE")
	}
	if payloadtype == 72 || payloadtype == 73 { // could cause confusion with rtcp types
		return errors.New("ERR_RTP_PACKET_BADPAYLOADTYPE")
	}

	var packetlength, packetoffset int
	packetlength = SIZEOF_HEADER
	packetlength += int(numcsrcs) * 4 //sizeof(uint32_t)*((size_t)
	if gotextension {
		packetlength += SIZEOF_EXTENSION      //(RTPExtensionHeader);
		packetlength += int(extensionlen) * 4 //sizeof(uint32_t)*((size_t)
	}
	packetlength += len(payloaddata) //payloadlen;
	this.packet = make([]byte, packetlength)

	// Ok, now we'll just fill in...
	this.header = NewHeader()
	this.header.version = RTP_VERSION
	this.header.padding = 0
	if gotextension {
		this.header.extension = 1
	} else {
		this.header.extension = 0
	}
	this.header.csrccount = numcsrcs
	if gotmarker {
		this.header.marker = 1
	} else {
		this.header.marker = 0
	}
	this.header.payloadtype = payloadtype & 127
	this.header.sequencenumber = seqnr
	this.header.timestamp = timestamp
	this.header.ssrc = ssrc
	if numcsrcs != 0 {
		this.header.csrc = make([]uint32, numcsrcs)
		for i := uint8(0); i < numcsrcs; i++ {
			this.header.csrc[i] = csrcs[i] //htonl(csrcs[i]);
		}
	}

	packetoffset = SIZEOF_HEADER + int(numcsrcs)*4
	copy(this.packet[0:packetoffset], this.header.Encode())

	if gotextension {
		this.extension = NewExtension()
		this.extension.id = extensionid
		this.extension.length = extensionlen //sizeof(uint32_t);
		if extensionlen != 0 {
			this.extension.data = make([]uint32, extensionlen)
			for i := uint16(0); i < extensionlen; i++ {
				this.extension.data[i] = extensiondata[i]
			}
		}
		copy(this.packet[packetoffset:packetoffset+SIZEOF_EXTENSION+int(extensionlen)*4], this.extension.Encode())

		packetoffset += SIZEOF_EXTENSION + int(extensionlen)*4
	} else {
		this.extension = nil
	}

	this.payload = make([]byte, len(payloaddata))
	copy(this.payload, payloaddata)
	copy(this.packet[packetoffset:packetoffset+len(payloaddata)], payloaddata)

	return nil
}

/** Returns \c true if the RTP packet has a header extension and \c false otherwise. */
func (this *Packet) HasExtension() bool {
	return this.header.extension != 0
}

/** Returns \c true if the marker bit was set and \c false otherwise. */
func (this *Packet) HasMarker() bool {
	return this.header.marker != 0
}

/** Returns the number of CSRCs contained in this packet. */
func (this *Packet) GetCSRCCount() uint8 {
	return this.header.csrccount
}

/** Returns a specific CSRC identifier.
 *  Returns a specific CSRC identifier. The parameter \c num can go from 0 to GetCSRCCount()-1.
 */
func (this *Packet) GetCSRC(num uint8) uint32 {
	if num >= this.header.csrccount {
		return 0
	}

	return this.header.csrc[num]
}

/** Returns the payload type of the packet. */
func (this *Packet) GetPayloadType() uint8 {
	return this.header.payloadtype
}

/** Returns the extended sequence number of the packet.
 *  Returns the extended sequence number of the packet. When the packet is just received,
 *  only the low $16$ bits will be set. The high 16 bits can be filled in later.
 */
// func (this *Packet) GetExtendedSequenceNumber() uint32 {
// 	return this.extseqnr
// }

/** Returns the sequence number of this packet. */
func (this *Packet) GetSequenceNumber() uint16 {
	return this.header.sequencenumber //uint16(this.extseqnr & 0x0000FFFF)
}

/** Sets the extended sequence number of this packet to \c seq. */
// func (this *Packet) SetExtendedSequenceNumber(seq uint32) {
// 	this.extseqnr = seq
// }

/** Returns the timestamp of this packet. */
func (this *Packet) GetTimestamp() uint32 {
	return this.header.timestamp
}

/** Returns the SSRC identifier stored in this packet. */
func (this *Packet) GetSSRC() uint32 {
	return this.header.ssrc
}

/** Returns a pointer to the actual payload data. */
func (this *Packet) GetPayload() []byte {
	return this.payload
}

/** If a header extension is present, this function returns the extension identifier. */
func (this *Packet) GetExtensionID() uint16 {
	return this.extension.id
}

/** Returns the length of the header extension data. */
func (this *Packet) GetExtensionLength() uint16 {
	return this.extension.length
}

/** Returns the header extension data. */
func (this *Packet) GetExtensionData() []uint32 {
	return this.extension.data
}

/** Returns the time at which this packet was received.
 *  When an RTPPacket instance is created from an RTPRawPacket instance, the raw packet's
 *  reception time is stored in the RTPPacket instance. This function then retrieves that
 *  time.
 */
func (this *Packet) GetReceiveTime() *RTPTime {
	return this.receivetime
}

/** Returns a pointer to the data of the entire packet. */
func (this *Packet) GetPacket() []byte {
	return this.packet
}

func (this *Packet) Dump() {
	/*int i;

	printf("Payload type:                %d\n",(int)GetPayloadType());
	printf("Extended sequence number:    0x%08x\n",GetExtendedSequenceNumber());
	printf("Timestamp:                   0x%08x\n",GetTimestamp());
	printf("SSRC:                        0x%08x\n",GetSSRC());
	printf("Marker:                      %s\n",HasMarker()?"yes":"no");
	printf("CSRC count:                  %d\n",GetCSRCCount());
	for (i = 0 ; i < GetCSRCCount() ; i++)
		printf("    CSRC[%02d]:                0x%08x\n",i,GetCSRC(i));
	printf("Payload:                     %s\n",GetPayloadData());
	printf("Payload length:              %d\n",GetPayloadLength());
	printf("Packet length:               %d\n",GetPacketLength());
	printf("Extension:                   %s\n",HasExtension()?"yes":"no");
	if (HasExtension())
	{
		printf("    Extension ID:            0x%04x\n",GetExtensionID());
		printf("    Extension data:          %s\n",GetExtensionData());
		printf("    Extension length:        %d\n",GetExtensionLength());
	}*/
}
