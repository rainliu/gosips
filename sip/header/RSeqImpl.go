package header

import (
	"errors"
	"gosips/core"
	"strconv"
)

type RSeq struct {
	SIPHeader

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
func (this *RSeq) SetSequenceNumber(sequenceNumber int) (InvalidArgumentException error) {
	if sequenceNumber <= 0 {
		return errors.New("InvalidArgumentException: Bad seq number")
	}
	this.sequenceNumber = sequenceNumber
	return nil
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
