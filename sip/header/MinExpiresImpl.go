package header

import (
	"errors"
	"gosips/core"
	"strconv"
)

/**
* MinExpires SIP Header.
 */
type MinExpires struct {
	SIPHeader

	/** expires field
	 */
	expires int
}

/** default constructor
 */
func NewMinExpires() *MinExpires {
	this := &MinExpires{}
	this.SIPHeader.super(core.SIPHeaderNames_MIN_EXPIRES)
	return this
}

func (this *MinExpires) String() string {
	return this.headerName + core.SIPSeparatorNames_COLON +
		core.SIPSeparatorNames_SP + this.EncodeBody() + core.SIPSeparatorNames_NEWLINE
}

/**
 * Return canonical form.
 * @return String
 */
func (this *MinExpires) EncodeBody() string {
	return strconv.Itoa(this.expires)
}

/**
 * Gets the expires value of the ExpiresHeader. This expires value is
 *
 * relative time.
 *
 *
 *
 * @return the expires value of the ExpiresHeader.
 *
 */
func (this *MinExpires) GetExpires() int {
	return this.expires
}

/**
         * Sets the relative expires value of the ExpiresHeader.
	 * The expires value MUST be greater than zero and MUST be
	 * less than 2**31.
         *
         * @param expires - the new expires value of this ExpiresHeader
         *
         * @throws InvalidArgumentException if supplied value is less than zero.
         *
*/
func (this *MinExpires) SetExpires(expires int) (InvalidArgumentException error) {
	if expires < 0 {
		return errors.New("InvalidArgumentException: bad argument")
	}
	this.expires = expires
	return nil
}
