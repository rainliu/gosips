package header

import (
	"gosip/core"
	"strconv"
)

/**

*TimeStamp SIP Header.

*

*@version  JAIN-SIP-1.1

*

*@author M. Ranganathan <mranga@nist.gov>  <br/>

*@author Olivier Deruelle <deruelle@nist.gov><br/>

*

*<a href="{@docRoot}/uncopyright.html">This code is in the public domain.</a>

*

 */

type TimeStamp struct {
	SIPHeader
	//implements TimeStampHeader{

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
	//if (delay!=-1)
	//    return new Float(timeStamp).toString()+ SP+ new Float(delay).toString();
	//  else  return new Float(timeStamp).toString();
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

/********************************************************************************/
/********************** JAIN-SIP 1.1 methods ************************************/
/********************************************************************************/

/**
 * Sets the timestamp value of this TimeStampHeader to the new timestamp
 * value passed to this method.
 *
 * @param timestamp - the new float timestamp value
 * @throws InvalidArgumentException if the timestamp value argument is a
 * negative value.
 */
func (this *TimeStamp) SetTimeStamp(timeStamp float32) { //throws InvalidArgumentException {
	//if (timeStamp<0) throw new InvalidArgumentException("JAIN-SIP Exception, TimeStamp, "+
	//      "setTimeStamp(), the timeStamp parameter is <0");
	this.timeStamp = timeStamp
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

func (this *TimeStamp) SetDelay(delay float32) { //throws InvalidArgumentException {
	//  if (delay<0 && delay!=-1) throw new InvalidArgumentException(
	//  "JAIN-SIP Exception, TimeStamp, "+"setDelay(), the delay parameter is <0");
	this.delay = delay
}
