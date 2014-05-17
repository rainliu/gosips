package header

import (
	"errors"
	"gosips/core"
	"strconv"
)

/**
* the WarningValue SIPObject.
*
* @see WarningList SIPHeader which strings these toGether.
 */

type Warning struct {
	SIPHeader

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
func (this *Warning) SetCode(code int) (InvalidArgumentException error) {
	if code >= 300 && code < 400 {
		this.code = code
		return nil
	} else {
		return errors.New("InvalidArgumentException: Code parameter in the Warning header is invalid")
	}
}

/**
 * Sets host of WarningHeader
 * @param host String to Set
 * @throws ParseException if host is not accepted by implementation
 */
func (this *Warning) SetAgent(host string) (ParseException error) {
	if host == "" {
		return errors.New("NullPointerException: the host parameter in the Warning header is null")
	} else {
		this.agent = host
		return nil
	}
}

/**
 * Sets text of WarningHeader
 * @param text String to Set
 * @throws ParseException if text is not accepted by implementation
 */
func (this *Warning) SetText(text string) (ParseException error) {
	if text == "" {
		return errors.New("ParseException: The text parameter in the Warning header is null")
	}
	this.text = text
	return nil
}
