package rtp

/** This class is used by the transmission component to store the incoming RTP and RTCP data in. */
type RawPacket struct {
	packetdata    []byte
	receivetime   *RTPTime
	senderaddress Address
	isrtp         bool
}

/** Creates an instance which stores data from \c data with length \c datalen.
 *  Creates an instance which stores data from \c data with length \c datalen. Only the pointer
 *  to the data is stored, no actual copy is made! The address from which this packet originated
 *  is set to \c address and the time at which the packet was received is set to \c recvtime.
 *  The flag which indicates whether this data is RTP or RTCP data is set to \c rtp. A memory
 *  manager can be installed as well.
 */
func NewRawPacket(data []byte,
	address Address,
	recvtime *RTPTime,
	rtp bool) *RawPacket {
	this := &RawPacket{}
	this.receivetime = recvtime
	this.packetdata = data
	this.senderaddress = address
	this.isrtp = rtp
	return this
}

/** Returns the pointer to the data which is contained in this packet. */
func (this *RawPacket) GetData() []byte {
	return this.packetdata
}

/** Returns the length of the packet described by this instance. */
func (this *RawPacket) GetDataLength() int {
	return len(this.packetdata)
}

/** Returns the time at which this packet was received. */
func (this *RawPacket) GetReceiveTime() *RTPTime {
	return this.receivetime
}

/** Returns the address stored in this packet. */
func (this *RawPacket) GetSenderAddress() Address {
	return this.senderaddress
}

/** Returns \c true if this data is RTP data, \c false if it is RTCP data. */
func (this *RawPacket) IsRTP() bool {
	return this.isrtp
}

/** Sets the pointer to the data stored in this packet to zero.
 *  Sets the pointer to the data stored in this packet to zero. This will prevent
 *  a \c delete call for the actual data when the destructor of RTPRawPacket is called.
 *  This function is used by the RTPPacket and RTCPCompoundPacket classes to obtain
 *  the packet data (without having to copy it)	and to make sure the data isn't deleted
 *  when the destructor of RTPRawPacket is called.
 */
func (this *RawPacket) ZeroData() {
	this.packetdata = nil
}
