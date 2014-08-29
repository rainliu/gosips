package rtp

import (
	"time"
)

type RTPTime struct {
	sec      uint32
	microsec uint32
}

func CurrentRTPTime() *RTPTime {
	t := time.Now()
	u := t.UnixNano() / 1000
	s := uint32(u / 1000000)
	m := uint32(u - int64(s)*1000000)
	return &RTPTime{sec: s, microsec: m}
}

func WaitRTPTime(delay *RTPTime) {
	time.Sleep(time.Duration(delay.sec)*time.Second + time.Duration(delay.microsec)*time.Microsecond)
}

func NewRTPTimeFromFloat64(t float64) *RTPTime {
	s := uint32(t)
	t2 := t - (float64(s))
	t2 *= 1000000.0
	m := uint32(t2)

	this := &RTPTime{sec: s, microsec: m}
	return this
}

func NewRTPTimeFromNTPTime(ntptime *NTPTime) *RTPTime {
	var s, m uint32

	if ntptime.GetMSW() < NTPTIMEOFFSET {
		s = 0
		m = 0
	} else {
		s = ntptime.GetMSW() - NTPTIMEOFFSET

		x := float64(ntptime.GetLSW())
		x /= (65536.0 * 65536.0)
		x *= 1000000.0
		m = uint32(x)
	}

	this := &RTPTime{sec: s, microsec: m}
	return this
}

func (this *RTPTime) Clone() *RTPTime {
	return &RTPTime{sec: this.sec, microsec: this.microsec}
}

func (this *RTPTime) GetNTPTime() *NTPTime {
	msw := this.sec + NTPTIMEOFFSET
	x := float64(this.microsec) / float64(1000000.0)
	x *= (65536.0 * 65536.0)
	lsw := uint32(x)

	return NewNTPTime(msw, lsw)
}

func (this *RTPTime) Sub(that *RTPTime) {
	this.sec -= that.sec
	if that.microsec > this.microsec {
		this.sec--
		this.microsec += 1000000
	}
	this.microsec -= that.microsec
}

func (this *RTPTime) Add(that *RTPTime) {
	this.sec += that.sec
	this.microsec += that.microsec
	if this.microsec >= 1000000 {
		this.sec++
		this.microsec -= 1000000
	}
}

func (this *RTPTime) LT(that *RTPTime) bool { //<
	if this.sec < that.sec {
		return true
	}
	if this.sec > that.sec {
		return false
	}
	if this.microsec < that.microsec {
		return true
	}
	return false
}

func (this *RTPTime) GT(that *RTPTime) bool { //>
	if this.sec > that.sec {
		return true
	}
	if this.sec < that.sec {
		return false
	}
	if this.microsec > that.microsec {
		return true
	}
	return false
}

func (this *RTPTime) ELT(that *RTPTime) bool { //<=
	if this.sec < that.sec {
		return true
	}
	if this.sec > that.sec {
		return false
	}
	if this.microsec <= that.microsec {
		return true
	}
	return false
}

func (this *RTPTime) EGT(that *RTPTime) bool { //>=
	if this.sec > that.sec {
		return true
	}
	if this.sec < that.sec {
		return false
	}
	if this.microsec >= that.microsec {
		return true
	}
	return false
}
