package header

import (
	"gosips/core"
	"strconv"
)

/**
* RAck SIP Header implementation
*
*@version  JAIN-SIP-1.1
*
*@author M. Ranganathan <mranga@nist.gov>  <br/>
*
*<a href="{@docRoot}/uncopyright.html">This code is in the public domain.</a>
*
 */
type RAck struct {
	SIPHeader
	//implements javax.sip.header.RAckHeader{

	cSeqNumber int

	rSeqNumber int

	method string
}

/** Creates a new instance of RAck */
func NewRAck() *RAck {
	this := &RAck{}
	this.SIPHeader.super(core.SIPHeaderNames_RACK)
	return this
}

func (this *RAck) String() string {
	return this.headerName + core.SIPSeparatorNames_COLON +
		core.SIPSeparatorNames_SP + this.EncodeBody() + core.SIPSeparatorNames_NEWLINE
}

/** Encode the body of this header (the stuff that follows headerName).
 * A.K.A headerValue.
 */
func (this *RAck) EncodeBody() string {
	return strconv.Itoa(this.rSeqNumber) + core.SIPSeparatorNames_SP + strconv.Itoa(this.cSeqNumber) +
		core.SIPSeparatorNames_SP + this.method

}

/** Gets the CSeq sequence number of this RAckHeader.
 *
 * @return the integer value of the cSeq number of the RAckHeader
 */
func (this *RAck) GetCSeqNumber() int {
	return this.cSeqNumber
}

/** Gets the method of RAckHeader
 *
 * @return method of RAckHeader
 */
func (this *RAck) GetMethod() string {
	return this.method
}

/** Gets the RSeq sequence number of this RAckHeader.
 *
 * @return the integer value of the RSeq number of the RAckHeader
 */
func (this *RAck) GetRSeqNumber() int {
	return this.rSeqNumber
}

/** Sets the sequence number value of the CSeqHeader of the provisional
 * response being acknowledged. The sequence number MUST be expressible as
 * a 32-bit unsigned integer and MUST be less than 2**31.
 *
 * @param cSeqNumber - the new cSeq number of this RAckHeader
 * @throws InvalidArgumentException if supplied value is less than zero.
 */
func (this *RAck) SetCSeqNumber(cSeqNumber int) {
	//throws InvalidArgumentException {
	// if cSeqNumber <= 0
	// 	throw new InvalidArgumentException("Bad CSeq # "  + cSeqNumber);
	this.cSeqNumber = cSeqNumber
}

/** Sets the method of RAckHeader, which correlates to the method of the
 * CSeqHeader of the provisional response being acknowledged.
 *
 * @param method - the new string value of the method of the RAckHeader
 * @throws ParseException which signals that an error has been reached
 * unexpectedly while parsing the method value.
 */
func (this *RAck) SetMethod(method string) { //throws ParseException {
	this.method = method
}

/** Sets the sequence number value of the RSeqHeader of the provisional
 * response being acknowledged. The sequence number MUST be expressible as
 * a 32-bit unsigned integer and MUST be less than 2**31.
 *
 * @param rSeqNumber - the new rSeq number of this RAckHeader
 * @throws InvalidArgumentException if supplied value is less than zero.
 */
func (this *RAck) SetRSeqNumber(rSeqNumber int) { //throws InvalidArgumentException {
	//if rSeqNumber <= 0
	//	throw new InvalidArgumentException("Bad rSeq # "  + rSeqNumber);
	this.rSeqNumber = rSeqNumber
}
