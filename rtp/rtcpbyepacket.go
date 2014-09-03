package rtp

/** Describes an RTCP BYE packet. */
type RTCPBYEPacket struct {
	RTCPPacket
	reasonoffset int
}

// 	/** Creates an instance based on the data in \c data with length \c datalen.
// 	 *  Creates an instance based on the data in \c data with length \c datalen. Since the \c data pointer
// 	 *  is referenced inside the class (no copy of the data is made) one must make sure that the memory it
// 	 *  points to is valid as long as the class instance exists.
// 	 */
// RTCPBYEPacket::RTCPBYEPacket(uint8_t *data,size_t datalength)
// 	: RTCPPacket(BYE,data,datalength)
// {
// 	knownformat = false;
// 	reasonoffset = 0;

// 	RTCPCommonHeader *hdr;
// 	size_t len = datalength;

// 	hdr = (RTCPCommonHeader *)data;
// 	if (hdr->padding)
// 	{
// 		uint8_t padcount = data[datalength-1];
// 		if ((padcount & 0x03) != 0) // not a multiple of four! (see rfc 3550 p 37)
// 			return;
// 		if (((size_t)padcount) >= len)
// 			return;
// 		len -= (size_t)padcount;
// 	}

// 	size_t ssrclen = ((size_t)(hdr->count))*sizeof(uint32_t) + sizeof(RTCPCommonHeader);
// 	if (ssrclen > len)
// 		return;
// 	if (ssrclen < len) // there's probably a reason for leaving
// 	{
// 		uint8_t *reasonlength = (data+ssrclen);
// 		size_t reaslen = (size_t)(*reasonlength);
// 		if (reaslen > (len-ssrclen-1))
// 			return;
// 		reasonoffset = ssrclen;
// 	}
// 	knownformat = true;
// }
// 	/** Returns the number of SSRC identifiers present in this BYE packet. */
// 	int GetSSRCCount() const;

// 	/** Returns the SSRC described by \c index which may have a value from 0 to GetSSRCCount()-1
// 	 *  (note that no check is performed to see if \c index is valid).
// 	 */
// 	uint32_t GetSSRC(int index) const; // note: no check is performed to see if index is valid!

// 	/** Returns true if the BYE packet contains a reason for leaving. */
// 	bool HasReasonForLeaving() const;

// 	/** Returns the length of the string which describes why the source(s) left. */
// 	size_t GetReasonLength() const;

// 	/** Returns the actual reason for leaving data. */
// 	uint8_t *GetReasonData();

// inline int RTCPBYEPacket::GetSSRCCount() const
// {
// 	if (!knownformat)
// 		return 0;

// 	RTCPCommonHeader *hdr = (RTCPCommonHeader *)data;
// 	return (int)(hdr->count);
// }

// inline uint32_t RTCPBYEPacket::GetSSRC(int index) const
// {
// 	if (!knownformat)
// 		return 0;
// 	uint32_t *ssrc = (uint32_t *)(data+sizeof(RTCPCommonHeader)+sizeof(uint32_t)*index);
// 	return ntohl(*ssrc);
// }

// inline bool RTCPBYEPacket::HasReasonForLeaving() const
// {
// 	if (!knownformat)
// 		return false;
// 	if (reasonoffset == 0)
// 		return false;
// 	return true;
// }

// inline size_t RTCPBYEPacket::GetReasonLength() const
// {
// 	if (!knownformat)
// 		return 0;
// 	if (reasonoffset == 0)
// 		return 0;
// 	uint8_t *reasonlen = (data+reasonoffset);
// 	return (size_t)(*reasonlen);
// }

// inline uint8_t *RTCPBYEPacket::GetReasonData()
// {
// 	if (!knownformat)
// 		return 0;
// 	if (reasonoffset == 0)
// 		return 0;
// 	uint8_t *reasonlen = (data+reasonoffset);
// 	if ((*reasonlen) == 0)
// 		return 0;
// 	return (data+reasonoffset+1);
// }

// void RTCPBYEPacket::Dump()
// {
// 	RTCPPacket::Dump();
// 	if (!IsKnownFormat())
// 	{
// 		std::cout << "    Unknown format" << std::endl;
// 		return;
// 	}

// 	int num = GetSSRCCount();
// 	int i;

// 	for (i = 0 ; i < num ; i++)
// 		std::cout << "    SSRC: " << GetSSRC(i) << std::endl;
// 	if (HasReasonForLeaving())
// 	{
// 		char str[1024];
// 		memcpy(str,GetReasonData(),GetReasonLength());
// 		str[GetReasonLength()] = 0;
// 		std::cout << "    Reason: " << str << std::endl;
// 	}
// }
