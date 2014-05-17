package header

import (
	"errors"
	"gosips/core"
	"strconv"
)

/**
* RAck SIP Header implementation
 */
type RAck struct {
	SIPHeader

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
func (this *RAck) SetCSeqNumber(cSeqNumber int) (InvalidArgumentException error) {
	if cSeqNumber <= 0 {
		return errors.New("InvalidArgumentException: Bad CSeq")
	}
	this.cSeqNumber = cSeqNumber
	return nil
}

/** Sets the method of RAckHeader, which correlates to the method of the
 * CSeqHeader of the provisional response being acknowledged.
 *
 * @param method - the new string value of the method of the RAckHeader
 * @throws ParseException which signals that an error has been reached
 * unexpectedly while parsing the method value.
 */
func (this *RAck) SetMethod(method string) (ParseException error) {
	this.method = method
	return nil
}

/** Sets the sequence number value of the RSeqHeader of the provisional
 * response being acknowledged. The sequence number MUST be expressible as
 * a 32-bit unsigned integer and MUST be less than 2**31.
 *
 * @param rSeqNumber - the new rSeq number of this RAckHeader
 * @throws InvalidArgumentException if supplied value is less than zero.
 */
func (this *RAck) SetRSeqNumber(rSeqNumber int) (InvalidArgumentException error) {
	if rSeqNumber <= 0 {
		return errors.New("InvalidArgumentException: Bad rSeq")
	}
	this.rSeqNumber = rSeqNumber
	return nil
}
