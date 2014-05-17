package header

import (
	"errors"
	"gosips/core"
)

/**
* the Unsupported header.
 */
type Unsupported struct {
	SIPHeader

	/** option-Tag field.
	 */
	optionTag string
}

/** Default Constructor.
 */
func NewUnsupported() *Unsupported {
	this := &Unsupported{}
	this.SIPHeader.super(core.SIPHeaderNames_UNSUPPORTED)
	return this
}

/** Constructor
 * @param ot String to set
 */
func NewUnsupportedFromString(ot string) *Unsupported {
	this := &Unsupported{}
	this.SIPHeader.super(core.SIPHeaderNames_UNSUPPORTED)
	this.optionTag = ot
	return this
}

func (this *Unsupported) String() string {
	return this.headerName + core.SIPSeparatorNames_COLON +
		core.SIPSeparatorNames_SP + this.EncodeBody() + core.SIPSeparatorNames_NEWLINE
}

/**
 * Return a canonical value.
 * @return String.
 */
func (this *Unsupported) EncodeBody() string {
	return this.optionTag
}

/** get the option tag field
 * @return option Tag field
 */
func (this *Unsupported) GetOptionTag() string {
	return this.optionTag
}

/**
 * Set the option member
 * @param o String to set
 */
func (this *Unsupported) SetOptionTag(o string) (ParseException error) {
	if o == "" {
		return errors.New("NullPointerException: The option tag parameter is null")
	}
	this.optionTag = o
	return nil
}
