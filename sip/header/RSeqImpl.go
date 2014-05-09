package header

import (
	"gosips/core"
	"strconv"
)

/**
*
*@version  JAIN-SIP-1.1
*
*@author M. Ranganathan <mranga@nist.gov>  <br/>
*
*<a href="{@docRoot}/uncopyright.html">This code is in the public domain.</a>
*
 */

type RSeq struct {
	SIPHeader
	//implements javax.sip.header.RSeqHeader {
	sequenceNumber int
}

/** Creates a new instance of RSeq */
func NewRSeq() *RSeq {
	this := &RSeq{}
	this.SIPHeader.super(core.SIPHeaderNames_RSEQ)
	return this
}

/** Gets the sequence number of this RSeqHeader.
 *
 * @return the integer value of the Sequence number of the RSeqHeader
 */
func (this *RSeq) GetSequenceNumber() int {
	return this.sequenceNumber
}

/** Sets the sequence number value of the RSeqHeader of the provisional
 * response. The sequence number MUST be expressible as a 32-bit unsigned
 * integer and MUST be less than 2**31.
 *
 * @param sequenceNumber - the new Sequence number of this RSeqHeader
 * @throws InvalidArgumentException if supplied value is less than zero.
 */
func (this *RSeq) SetSequenceNumber(sequenceNumber int) {
	// throws InvalidArgumentException {
	//     if (sequenceNumber <= 0)
	//             throw new InvalidArgumentException
	//                 ("Bad seq number " + sequenceNumber);
	this.sequenceNumber = sequenceNumber
}

func (this *RSeq) String() string {
	return this.headerName + core.SIPSeparatorNames_COLON +
		core.SIPSeparatorNames_SP + this.EncodeBody() + core.SIPSeparatorNames_NEWLINE
}

/** Encode the body of this header (the stuff that follows headerName).
 * A.K.A headerValue.
 */
func (this *RSeq) EncodeBody() string {
	return strconv.Itoa(this.sequenceNumber)
}
