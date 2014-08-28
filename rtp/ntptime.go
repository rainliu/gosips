package rtp

const NTPTIMEOFFSET = 2208988800

/**
 * This is a simple wrapper for the most significant word (MSW) and least
 * significant word (LSW) of an NTP timestamp.
 */
type NTPTime struct {
	msw uint32
	lsw uint32
}

func NewNTPTime(m, l uint32) *NTPTime {
	this := &NTPTime{msw: m, lsw: l}
	return this
}

/** Returns the most significant word. */
func (this *NTPTime) GetMSW() uint32 {
	return this.msw
}

/** Returns the least significant word. */
func (this *NTPTime) GetLSW() uint32 {
	return this.lsw
}
