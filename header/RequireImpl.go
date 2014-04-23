package header

import "gosip/core"

/**
*Require SIP Header.
*
*@version  JAIN-SIP-1.1
*
*@author M. Ranganathan <mranga@nist.gov>  <br/>
*@author Olivier Deruelle <deruelle@nist.gov><br/>
*
*<a href="{@docRoot}/uncopyright.html">This code is in the public domain.</a>
*
 */
type Require struct {
	SIPHeader // implements RequireHeader{

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
func (this *Require) SetOptionTag(optionTag string) { //throws ParseException {
	//if (optionTag==null) throw new  NullPointerException(
	//  "JAIN-SIP Exception, Require, "+"setOptionTag(), the optionTag parameter is null");
	this.optionTag = optionTag
}

/**
 * Gets the option tag of this OptionTag class.
 *
 * @return the string that identifies the option tag value.
 */
func (this *Require) GetOptionTag() string {
	return this.optionTag
}
