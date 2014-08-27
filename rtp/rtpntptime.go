package rtp

const RTP_NTPTIMEOFFSET = uint32(2208988800)

/**
 * This is a simple wrapper for the most significant word (MSW) and least
 * significant word (LSW) of an NTP timestamp.
 */
type RTPNTPTime struct {
	msw uint32
	lsw uint32
}

func NewRTPNTPTime(m, l uint32) *RTPNTPTime {
	this := &RTPNTPTime{msw: m, lsw: l}
	return this
}

/** Returns the most significant word. */
func (this *RTPNTPTime) GetMSW() uint32 {
	return this.msw
}

/** Returns the least significant word. */
func (this *RTPNTPTime) GetLSW() uint32 {
	return this.lsw
}
