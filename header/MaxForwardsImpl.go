package header

import (
	"errors"
	"gosip/core"
	"strconv"
)

/**
* MaxForwards SIPHeaderHeader
 */
type MaxForwards struct { //implements MaxForwardsHeader {
	SIPHeader

	/** maxForwards field.
	 */
	maxForwards int
}

/** Default constructor.
 */
func NewMaxForwards() *MaxForwards {
	this := &MaxForwards{}
	this.SIPHeader.super(core.SIPHeaderNames_MAX_FORWARDS)
	return this
}

func (this *MaxForwards) super(maxForwards int) {
	this.SIPHeader.super(core.SIPHeaderNames_MAX_FORWARDS)
	this.maxForwards = maxForwards
}

/** get the MaxForwards field.
 * @return the maxForwards member.
 */
func (this *MaxForwards) GetMaxForwards() int {
	return this.maxForwards
}

/**
 * Set the maxForwards member
 * @param maxForwards maxForwards parameter to set
 */
func (this *MaxForwards) SetMaxForwards(maxForwards int) (InvalidArgumentException error) {
	if maxForwards < 0 || maxForwards > 255 {
		return errors.New("bad max forwards value " + strconv.Itoa(maxForwards))
	}
	this.maxForwards = maxForwards
	return nil
}

func (this *MaxForwards) String() string {
	return this.headerName + core.SIPSeparatorNames_COLON + core.SIPSeparatorNames_SP + this.EncodeBody() + core.SIPSeparatorNames_NEWLINE
}

/**
 * Encode into a string.
 * @return encoded string.
 *
 */
func (this *MaxForwards) EncodeBody() string {
	return strconv.Itoa(this.maxForwards)
}

/** Boolean function
 * @return true if MaxForwards field reached zero.
 */
func (this *MaxForwards) HasReachedZero() bool {
	return this.maxForwards == 0
}

/** decrement MaxForwards field one by one.
 */
func (this *MaxForwards) DecrementMaxForwards() {
	if this.maxForwards >= 0 {
		this.maxForwards--
	}
}
