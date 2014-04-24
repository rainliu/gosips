package header

import "gosip/core"

/**
* the Unsupported header.
*
*@author Olivier Deruelle <deruelle@nist.gov><br/>
*
*<a href="{@docRoot}/uncopyright.html">This code is in the public domain.</a>
*
 */
type Unsupported struct {
	SIPHeader
	//implements javax.sip.header.UnsupportedHeader {

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
func (this *Unsupported) SetOptionTag(o string) { //throws ParseException{
	//if (o==null) throw new  NullPointerException("JAIN-SIP Exception, "+
	//" Unsupported, setOptionTag(), The option tag parameter is null");
	this.optionTag = o
}
