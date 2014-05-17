package header

import (
	"errors"
	"gosips/core"
	"strconv"
)

/**
*TimeStamp SIP Header.
 */

type TimeStamp struct {
	SIPHeader

	/** timeStamp field
	 */
	timeStamp float32

	/** delay field
	 */
	delay float32
}

/** Default Constructor
 */
func NewTimeStamp() *TimeStamp {
	this := &TimeStamp{}
	this.SIPHeader.super(core.SIPHeaderNames_TIMESTAMP)
	this.delay = -1
	return this
}

func (this *TimeStamp) String() string {
	return this.headerName + core.SIPSeparatorNames_COLON +
		core.SIPSeparatorNames_SP + this.EncodeBody() + core.SIPSeparatorNames_NEWLINE
}

/**
 * Return canonical form of the header.
 * @return String
 */
func (this *TimeStamp) EncodeBody() string {
	if this.delay != -1 {
		return strconv.FormatFloat(float64(this.timeStamp), 'f', -1, 32) + core.SIPSeparatorNames_SP + strconv.FormatFloat(float64(this.delay), 'f', -1, 32)
	} else {
		return strconv.FormatFloat(float64(this.timeStamp), 'f', -1, 32)
	}
}

/** return true if delay exists
 * @return boolean
 */
func (this *TimeStamp) HasDelay() bool {
	return this.delay != -1
}

/* remove the Delay field
 */
func (this *TimeStamp) RemoveDelay() {
	this.delay = -1
}

/**
 * Sets the timestamp value of this TimeStampHeader to the new timestamp
 * value passed to this method.
 *
 * @param timestamp - the new float timestamp value
 * @throws InvalidArgumentException if the timestamp value argument is a
 * negative value.
 */
func (this *TimeStamp) SetTimeStamp(timeStamp float32) (InvalidArgumentException error) {
	if timeStamp < 0 {
		return errors.New("InvalidArgumentException: the timeStamp parameter is <0")
	}
	this.timeStamp = timeStamp
	return nil
}

/**
 * Gets the timestamp value of this TimeStampHeader.
 *
 * @return the timestamp value of this TimeStampHeader
 */
func (this *TimeStamp) GetTimeStamp() float32 {
	return this.timeStamp
}

/**
 * Gets delay of TimeStampHeader. This method return <code>-1</code> if the
 * delay paramater is not set.
 *
 * @return the delay value of this TimeStampHeader
 */
func (this *TimeStamp) GetDelay() float32 {
	return this.delay
}

/**
 * Sets the new delay value of the TimestampHeader to the delay paramter
 * passed to this method
 *
 * @param delay - the new float delay value
 * @throws InvalidArgumentException if the delay value argumenmt is a
 * negative value other than <code>-1</code>.
 */

func (this *TimeStamp) SetDelay(delay float32) (InvalidArgumentException error) {
	if delay < 0 && delay != -1 {
		return errors.New("InvalidArgumentException: the delay parameter is <0")
	}
	this.delay = delay
	return nil
}
