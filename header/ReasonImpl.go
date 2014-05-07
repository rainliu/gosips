package header

import (
	"bytes"
	"gosip/core"
	"strconv"
)

/**
*Definition of the Reason SIP Header.
*
*@version  JAIN-SIP-1.1
*
*@author M. Ranganathan <mranga@nist.gov>  <br/>
*
*<a href="{@docRoot}/uncopyright.html">This code is in the public domain.</a>
*
 */
type Reason struct {
	Parameters
	//implements javax.sip.header.ReasonHeader {

	//public final String TEXT = ParameterNames.TEXT;
	//public final String CAUSE = ParameterNames.CAUSE;

	protocol string
}

/** Creates a new instance of Reason */
func NewReason() *Reason {
	this := &Reason{}
	this.Parameters.super(core.SIPHeaderNames_REASON)
	return this
}

/** Get the cause token.
 *@return the cause code.
 */
func (this *Reason) GetCause() int {
	cause, _ := strconv.Atoi(this.GetParameter(ParameterNames_CAUSE))
	return cause
}

/**
 * Set the cause.
 *
 *@param cause - cause to Set.
 */
func (this *Reason) SetCause(cause int) {
	// throws javax.sip.InvalidArgumentException {
	this.parameters.AddNameAndValue("cause", strconv.Itoa(cause))
}

/** Set the protocol
 *
 *@param protocol - protocol to Set.
 */

func (this *Reason) SetProtocol(protocol string) { //throws ParseException {
	this.protocol = protocol
}

/** Return the protocol.
 *
 *@return the protocol.
 */
func (this *Reason) GetProtocol() string {
	return this.protocol
}

/** Set the text.
 *
 *@param text -- string text to Set.
 */
func (this *Reason) SetText(text string) { //throws ParseException {
	this.parameters.AddNameAndValue("text", text)
}

/** Get the text.
 *
 *@return text parameter.
 *
 */
func (this *Reason) GetText() string {
	return this.parameters.GetParameter("text")
}

/** Set the cause.


  /** Gets the unique string name of this Header. A name constant is defined in
   * each individual Header identifying each Header.
   *
   * @return the name of this specific Header
*/
func (this *Reason) GetName() string {
	return core.SIPHeaderNames_REASON

}

func (this *Reason) String() string {
	return this.headerName + core.SIPSeparatorNames_COLON +
		core.SIPSeparatorNames_SP + this.EncodeBody() + core.SIPSeparatorNames_NEWLINE
}

/** Encode the body of this header (the stuff that follows headerName).
 * A.K.A headerValue.
 */
func (this *Reason) EncodeBody() string {
	var encoding bytes.Buffer //  = new StringBuffer();
	encoding.WriteString(this.protocol)

	if this.parameters != nil && this.parameters.Len() > 0 {
		encoding.WriteString(core.SIPSeparatorNames_SEMICOLON)
		encoding.WriteString(this.parameters.String())
	}
	return encoding.String()
}

/*
func (this *Reason) SetParameter(name, value string) error {
	//throws ParseException {
	//if (name == null) throw new NullPointerException("null name");
	nv := this.parameters.GetNameValue(strings.ToLower(name))
	if nv == nil {
		nv = core.NewNameValue(name, value)
		if strings.ToLower(name) == (ParameterNames_TEXT) {
			// if (value ==
			//     throw new NullPointerException("null value");
			// if (value.startsWith(Separators.DOUBLE_QUOTE))
			//     throw new ParseException
			//     (value + " : Unexpected DOUBLE_QUOTE",0);
			nv.SetQuotedValue()
		}
		this.parameters.SetNameValue(nv)
	} else {
		nv.SetValue(value)
	}

	return nil
}*/
