package header

import (
	"errors"
	"gosips/core"
)

/**
*Require SIP Header.
 */
type Require struct {
	SIPHeader

	/** optionTag field
	 */
	optionTag string
}

/** Default constructor
 * @param s String to set
 */
func NewRequire() *Require {
	this := &Require{}
	this.SIPHeader.super(core.SIPHeaderNames_REQUIRE)
	return this
}

/** constructor
 * @param s String to set
 */
func NewRequireFromString(s string) *Require {
	this := &Require{}
	this.SIPHeader.super(core.SIPHeaderNames_REQUIRE)
	this.optionTag = s
	return this
}

func (this *Require) String() string {
	return this.headerName + core.SIPSeparatorNames_COLON +
		core.SIPSeparatorNames_SP + this.EncodeBody() + core.SIPSeparatorNames_NEWLINE
}

/**
 * Encode in canonical form.
 * @return String
 */
func (this *Require) EncodeBody() string {
	return this.optionTag
}

/**
 * Sets the option tag value to the new supplied <var>optionTag</var>
 * parameter.
 *
 * @param optionTag - the new string value of the option tag.
 * @throws ParseException which signals that an error has been reached
 * unexpectedly while parsing the optionTag value.
 */
func (this *Require) SetOptionTag(optionTag string) (ParseException error) {
	if optionTag == "" {
		return errors.New("NullPointerException: the optionTag parameter is null")
	}
	this.optionTag = optionTag
	return nil
}

/**
 * Gets the option tag of this OptionTag class.
 *
 * @return the string that identifies the option tag value.
 */
func (this *Require) GetOptionTag() string {
	return this.optionTag
}
