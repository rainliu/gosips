package header

import (
	"gosips/core"
	"strconv"
)

/**
* the WarningValue SIPObject.
*
*@author M. Ranganathan <mranga@nist.gov>  <br/>
*@author Olivier Deruelle <deruelle@nist.gov><br/>
*
*<a href="{@docRoot}/uncopyright.html">This code is in the public domain.</a>
*
* @see WarningList SIPHeader which strings these toGether.
 */

type Warning struct {
	SIPHeader //implements  WarningHeader {

	/** warn code field, the warn code consists of three digits.
	 */
	code int

	/** the name or pseudonym of the server adding
	 * the Warning header, for use in debugging
	 */
	agent string

	/** warn-text field
	 */
	text string
}

/**
 * constructor.
 */
func NewWarning() *Warning {
	this := &Warning{}
	this.SIPHeader.super(core.SIPHeaderNames_WARNING)
	return this
}

func (this *Warning) String() string {
	return this.headerName + core.SIPSeparatorNames_COLON +
		core.SIPSeparatorNames_SP + this.EncodeBody() + core.SIPSeparatorNames_NEWLINE
}

/** Encode the body of the header (return the stuff following name:).
 *@return the string encoding of the header value.
 */
func (this *Warning) EncodeBody() string {
	//print(this.code)
	if this.text != "" {
		return strconv.Itoa(this.code) + core.SIPSeparatorNames_SP + this.agent +
			core.SIPSeparatorNames_SP + core.SIPSeparatorNames_DOUBLE_QUOTE + this.text + core.SIPSeparatorNames_DOUBLE_QUOTE
	} else {
		return strconv.Itoa(this.code) + core.SIPSeparatorNames_SP + this.agent
	}
}

/**
 * Gets code of WarningHeader
 * @return code of WarningHeader
 */
func (this *Warning) GetCode() int {
	return this.code
}

/**
 * Gets agent host of WarningHeader
 * @return agent host of WarningHeader
 */
func (this *Warning) GetAgent() string {
	return this.agent
}

/**
 * Gets text of WarningHeader
 * @return text of WarningHeader
 */
func (this *Warning) GetText() string {
	return this.text
}

/**
 * Sets code of WarningHeader
 * @param code int to Set
 * @throws SipParseException if code is not accepted by implementation
 */
func (this *Warning) SetCode(code int) { //throws InvalidArgumentException {
	//println(code)
	if code >= 300 && code < 400 {
		this.code = code
	}
	//println(this.code)
	// else throw new InvalidArgumentException
	// ("Code parameter in the Warning header is invalid: code="+code);
}

/**
 * Sets host of WarningHeader
 * @param host String to Set
 * @throws ParseException if host is not accepted by implementation
 */
func (this *Warning) SetAgent(host string) { //throws ParseException {
	// if (host==null)
	//       throw new NullPointerException
	//      ("the host parameter in the Warning header is null");
	// else {
	this.agent = host
	//}
}

/**
 * Sets text of WarningHeader
 * @param text String to Set
 * @throws ParseException if text is not accepted by implementation
 */
func (this *Warning) SetText(text string) { //throws ParseException {
	// if (text==null) {
	//      throw new ParseException
	//        ("The text parameter in the Warning header is null",0);
	// }
	this.text = text
}
