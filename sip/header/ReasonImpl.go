package header

import (
	"bytes"
	"gosips/core"
	"strconv"
)

/**
*Definition of the Reason SIP Header.
 */
type Reason struct {
	Parameters

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
func (this *Reason) SetCause(cause int) (InvalidArgumentException error) {
	this.parameters.AddNameAndValue("cause", strconv.Itoa(cause))
	return nil
}

/** Set the protocol
 *
 *@param protocol - protocol to Set.
 */

func (this *Reason) SetProtocol(protocol string) (ParseException error) {
	this.protocol = protocol
	return nil
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
func (this *Reason) SetText(text string) (ParseException error) {
	this.parameters.AddNameAndValue("text", text)
	return nil
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
	var encoding bytes.Buffer
	encoding.WriteString(this.protocol)

	if this.parameters != nil && this.parameters.Len() > 0 {
		encoding.WriteString(core.SIPSeparatorNames_SEMICOLON)
		encoding.WriteString(this.parameters.String())
	}
	return encoding.String()
}
