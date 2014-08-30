package rtp

import "errors"

/** Represents an RTP Packet.
 *  The RTPPacket class can be used to parse a RTPRawPacket instance if it represents RTP data.
 *  The class can also be used to create a new RTP packet according to the parameters specified by
 *  the user.
 */
type Packet struct {
	err error

	hasextension, hasmarker bool
	numcsrcs                int

	payloadtype                 uint8
	extseqnr, timestamp, ssrc   uint32
	packet                      []byte
	payload                     []byte
	packetlength, payloadlength int

	extid           uint16
	extension       []byte
	extensionlength int

	externalbuffer bool

	receivetime *RTPTime
}

/** Creates an RTPPacket instance based upon the data in \c rawpack, optionally installing a memory manager.
 *  If successful, the data is moved from the raw packet to the RTPPacket instance.
 */
func NewPacketFromRawPacket(rawpack *RawPacket) *Packet {
	this := &Packet{}
	this.receivetime = rawpack.GetReceiveTime().Clone()
	this.err = this.ParseRawPacket(rawpack)
	return this
}

/** Creates a new buffer for an RTP packet and fills in the fields according to the specified parameters.
 *  If \c maxpacksize is not equal to zero, an error is generated if the total packet size would exceed
 *  \c maxpacksize. The arguments of the constructor are self-explanatory. Note that the size of a header
 *  extension is specified in a number of 32-bit words. A memory manager can be installed.
 *  This constructor is similar to the other constructor, but here data is stored in an external buffer
 *  \c buffer with size \c buffersize. */
// func NewPacket(payloadtype uint8,
// 	payloaddata []byte,
// 	seqnr uint16,
// 	timestamp uint32,
// 	ssrc uint32,
// 	gotmarker bool,
// 	numcsrcs uint8,
// 	csrcs []uint32,
// 	gotextension bool,
// 	extensionid uint16,
// 	extensionlen_numwords uint16,
// 	extensiondata []byte,
// 	buffer []byte) *Packet {
// 	this := &Packet{}

// 	this.receivetime = &RTPTime{0, 0}
// 	this.err = this.BuildPacket(payloadtype,
// 		payloaddata,
// 		seqnr,
// 		timestamp,
// 		ssrc,
// 		gotmarker,
// 		numcsrcs,
// 		csrcs,
// 		gotextension,
// 		extensionid,
// 		extensionlen_numwords,
// 		extensiondata,
// 		buffer)

// 	return this
// }

func (this *Packet) ParseRawPacket(rawpack *RawPacket) error {
	var packetbytes []byte
	var packetlen int
	var payloadtype uint8
	var rtpheader *Header
	var marker bool
	var csrccount int
	var hasextension bool
	var payloadoffset, payloadlength int
	var numpadbytes int
	var rtpextheader *ExtensionHeader
	var exthdrlen uint16

	if !rawpack.IsRTP() { // If we didn't receive it on the RTP port, we'll ignore it
		return errors.New("ERR_RTP_PACKET_INVALIDPACKET")
	}

	// The length should be at least the size of the RTP header
	packetlen = rawpack.GetDataLength()
	if packetlen < SIZEOF_HEADER {
		return errors.New("ERR_RTP_PACKET_INVALIDPACKET")
	}

	packetbytes = rawpack.GetData()
	rtpheader = NewHeader(packetbytes)

	// The version number should be correct
	if rtpheader.version != RTP_VERSION {
		return errors.New("ERR_RTP_PACKET_INVALIDPACKET")
	}

	// We'll check if this is possibly a RTCP packet. For this to be possible
	// the marker bit and payload type combined should be either an SR or RR
	// identifier
	marker = rtpheader.marker != 0
	payloadtype = rtpheader.payloadtype
	if marker {
		if payloadtype == (RTP_RTCPTYPE_SR & 127) { // don't check high bit (this was the marker!!)
			return errors.New("ERR_RTP_PACKET_INVALIDPACKET")
		}
		if payloadtype == (RTP_RTCPTYPE_RR & 127) {
			return errors.New("ERR_RTP_PACKET_INVALIDPACKET")
		}
	}

	csrccount = int(rtpheader.csrccount)
	payloadoffset = SIZEOF_HEADER + csrccount*4 //sizeof(uint32_t));

	if rtpheader.padding != 0 { // adjust payload length to take padding into account
		numpadbytes = int(packetbytes[packetlen-1]) // last byte contains number of padding bytes
		if numpadbytes <= 0 {
			return errors.New("ERR_RTP_PACKET_INVALIDPACKET")
		}
	} else {
		numpadbytes = 0
	}

	hasextension = rtpheader.extension != 0
	if hasextension { // got header extension
		rtpextheader = NewExtensionHeader(packetbytes[payloadoffset:])
		payloadoffset += SIZEOF_EXTENSIONHEADER
		exthdrlen = rtpextheader.length     //ntohs(rtpextheader->length);
		payloadoffset += int(exthdrlen) * 4 //sizeof(uint32_t);
	} else {
		rtpextheader = nil
		exthdrlen = 0
	}

	payloadlength = packetlen - numpadbytes - payloadoffset
	if payloadlength < 0 {
		return errors.New("ERR_RTP_PACKET_INVALIDPACKET")
	}

	// Now, we've got a valid packet, so we can create a new instance of RTPPacket
	// and fill in the members
	this.hasextension = hasextension
	if hasextension {
		this.extid = rtpextheader.extid                                                 //ntohs(rtpextheader->extid);
		this.extensionlength = int(rtpextheader.length) * 4                             //((int)ntohs(rtpextheader->length))*sizeof(uint32_t);
		this.extension = packetbytes[SIZEOF_HEADER+csrccount*4+SIZEOF_EXTENSIONHEADER:] //((uint8_t *)rtpextheader)+sizeof(RTPExtensionHeader);
	}

	this.hasmarker = marker
	this.numcsrcs = csrccount
	this.payloadtype = payloadtype

	// Note: we don't fill in the EXTENDED sequence number here, since we
	// don't have information about the source here. We just fill in the low
	// 16 bits
	this.extseqnr = uint32(rtpheader.sequencenumber) //(uint32_t)ntohs(rtpheader->sequencenumber);

	this.timestamp = rtpheader.timestamp //ntohl(rtpheader->timestamp);
	this.ssrc = rtpheader.ssrc           //ntohl(rtpheader->ssrc);
	this.packet = packetbytes
	this.payload = packetbytes[payloadoffset:]
	this.packetlength = packetlen
	this.payloadlength = payloadlength

	// We'll zero the data of the raw packet, since we're using it here now!
	rawpack.ZeroData()

	return nil
}

// uint32_t RTPPacket::GetCSRC(int num) const
// {
// 	if (num >= numcsrcs)
// 		return 0;

// 	uint8_t *csrcpos;
// 	uint32_t *csrcval_nbo;
// 	uint32_t csrcval_hbo;

// 	csrcpos = packet+sizeof(RTPHeader)+num*sizeof(uint32_t);
// 	csrcval_nbo = (uint32_t *)csrcpos;
// 	csrcval_hbo = ntohl(*csrcval_nbo);
// 	return csrcval_hbo;
// }

// int RTPPacket::BuildPacket(uint8_t payloadtype,const void *payloaddata,size_t payloadlen,uint16_t seqnr,
// 		  uint32_t timestamp,uint32_t ssrc,bool gotmarker,uint8_t numcsrcs,const uint32_t *csrcs,
// 		  bool gotextension,uint16_t extensionid,uint16_t extensionlen_numwords,const void *extensiondata,
// 		  void *buffer,size_t maxsize)
// {
// 	if (numcsrcs > RTP_MAXCSRCS)
// 		return ERR_RTP_PACKET_TOOMANYCSRCS;

// 	if (payloadtype > 127) // high bit should not be used
// 		return ERR_RTP_PACKET_BADPAYLOADTYPE;
// 	if (payloadtype == 72 || payloadtype == 73) // could cause confusion with rtcp types
// 		return ERR_RTP_PACKET_BADPAYLOADTYPE;

// 	packetlength = sizeof(RTPHeader);
// 	packetlength += sizeof(uint32_t)*((size_t)numcsrcs);
// 	if (gotextension)
// 	{
// 		packetlength += sizeof(RTPExtensionHeader);
// 		packetlength += sizeof(uint32_t)*((size_t)extensionlen_numwords);
// 	}
// 	packetlength += payloadlen;

// 	if (maxsize > 0 && packetlength > maxsize)
// 	{
// 		packetlength = 0;
// 		return ERR_RTP_PACKET_DATAEXCEEDSMAXSIZE;
// 	}

// 	// Ok, now we'll just fill in...

// 	RTPHeader *rtphdr;

// 	if (buffer == 0)
// 	{
// 		packet = RTPNew(GetMemoryManager(),RTPMEM_TYPE_BUFFER_RTPPACKET) uint8_t [packetlength];
// 		if (packet == 0)
// 		{
// 			packetlength = 0;
// 			return ERR_RTP_OUTOFMEM;
// 		}
// 		externalbuffer = false;
// 	}
// 	else
// 	{
// 		packet = (uint8_t *)buffer;
// 		externalbuffer = true;
// 	}

// 	RTPPacket::hasmarker = gotmarker;
// 	RTPPacket::hasextension = gotextension;
// 	RTPPacket::numcsrcs = numcsrcs;
// 	RTPPacket::payloadtype = payloadtype;
// 	RTPPacket::extseqnr = (uint32_t)seqnr;
// 	RTPPacket::timestamp = timestamp;
// 	RTPPacket::ssrc = ssrc;
// 	RTPPacket::payloadlength = payloadlen;
// 	RTPPacket::extid = extensionid;
// 	RTPPacket::extensionlength = ((size_t)extensionlen_numwords)*sizeof(uint32_t);

// 	rtphdr = (RTPHeader *)packet;
// 	rtphdr->version = RTP_VERSION;
// 	rtphdr->padding = 0;
// 	if (gotmarker)
// 		rtphdr->marker = 1;
// 	else
// 		rtphdr->marker = 0;
// 	if (gotextension)
// 		rtphdr->extension = 1;
// 	else
// 		rtphdr->extension = 0;
// 	rtphdr->csrccount = numcsrcs;
// 	rtphdr->payloadtype = payloadtype&127; // make sure high bit isn't set
// 	rtphdr->sequencenumber = htons(seqnr);
// 	rtphdr->timestamp = htonl(timestamp);
// 	rtphdr->ssrc = htonl(ssrc);

// 	uint32_t *curcsrc;
// 	int i;

// 	curcsrc = (uint32_t *)(packet+sizeof(RTPHeader));
// 	for (i = 0 ; i < numcsrcs ; i++,curcsrc++)
// 		*curcsrc = htonl(csrcs[i]);

// 	payload = packet+sizeof(RTPHeader)+((size_t)numcsrcs)*sizeof(uint32_t);
// 	if (gotextension)
// 	{
// 		RTPExtensionHeader *rtpexthdr = (RTPExtensionHeader *)payload;

// 		rtpexthdr->extid = htons(extensionid);
// 		rtpexthdr->length = htons((uint16_t)extensionlen_numwords);

// 		payload += sizeof(RTPExtensionHeader);
// 		memcpy(payload,extensiondata,RTPPacket::extensionlength);

// 		payload += RTPPacket::extensionlength;
// 	}
// 	memcpy(payload,payloaddata,payloadlen);
// 	return 0;
// }

// 	/** If an error occurred in one of the constructors, this function returns the error code. */
// 	int GetCreationError() const														{ return error; }

// 	/** Returns \c true if the RTP packet has a header extension and \c false otherwise. */
// 	bool HasExtension() const															{ return hasextension; }

// 	/** Returns \c true if the marker bit was set and \c false otherwise. */
// 	bool HasMarker() const																{ return hasmarker; }

// 	/** Returns the number of CSRCs contained in this packet. */
// 	int GetCSRCCount() const															{ return numcsrcs; }

// 	/** Returns a specific CSRC identifier.
// 	 *  Returns a specific CSRC identifier. The parameter \c num can go from 0 to GetCSRCCount()-1.
// 	 */
// 	uint32_t GetCSRC(int num) const;

// 	/** Returns the payload type of the packet. */
// 	uint8_t GetPayloadType() const														{ return payloadtype; }

// 	/** Returns the extended sequence number of the packet.
// 	 *  Returns the extended sequence number of the packet. When the packet is just received,
// 	 *  only the low $16$ bits will be set. The high 16 bits can be filled in later.
// 	 */
// 	uint32_t GetExtendedSequenceNumber() const											{ return extseqnr; }

// 	/** Returns the sequence number of this packet. */
// 	uint16_t GetSequenceNumber() const													{ return (uint16_t)(extseqnr&0x0000FFFF); }

// 	/** Sets the extended sequence number of this packet to \c seq. */
// 	void SetExtendedSequenceNumber(uint32_t seq)										{ extseqnr = seq; }

// 	/** Returns the timestamp of this packet. */
// 	uint32_t GetTimestamp() const														{ return timestamp; }

// 	/** Returns the SSRC identifier stored in this packet. */
// 	uint32_t GetSSRC() const															{ return ssrc; }

// 	/** Returns a pointer to the data of the entire packet. */
// 	uint8_t *GetPacketData() const														{ return packet; }

// 	/** Returns a pointer to the actual payload data. */
// 	uint8_t *GetPayloadData() const														{ return payload; }

// 	/** Returns the length of the entire packet. */
// 	size_t GetPacketLength() const														{ return packetlength; }

// 	/** Returns the payload length. */
// 	size_t GetPayloadLength() const														{ return payloadlength; }

// 	/** If a header extension is present, this function returns the extension identifier. */
// 	uint16_t GetExtensionID() const														{ return extid; }

// 	/** Returns the length of the header extension data. */
// 	uint8_t *GetExtensionData() const													{ return extension; }

// 	/** Returns the length of the header extension data. */
// 	size_t GetExtensionLength() const													{ return extensionlength; }

// 	/** Returns the time at which this packet was received.
// 	 *  When an RTPPacket instance is created from an RTPRawPacket instance, the raw packet's
// 	 *  reception time is stored in the RTPPacket instance. This function then retrieves that
// 	 *  time.
// 	 */
// 	RTPTime GetReceiveTime() const														{ return receivetime; }
