package header

import (
	"bytes"
	"gosip/address"
	"gosip/core"
)

/**
* ErrorInfo SIP Header.
*
*@version  JAIN-SIP-1.1
*
*@author M. Ranganathan <mranga@nist.gov>  <br/>
*@author Olivier Deruelle <deruelle@nist.gov><br/>
*<a href="{@docRoot}/uncopyright.html">This code is in the public domain.</a>
*
 */
type ErrorInfo struct {
	Parameters
	//implements ErrorInfoHeader{

	errorInfo address.URI
}

/** Default constructor.
 */
func NewErrorInfo() *ErrorInfo {
	this := &ErrorInfo{}
	this.Parameters.super(core.SIPHeaderNames_ERROR_INFO)
	return this
}

/** Constructor given the error info
*@param errorInfo -- the error information to set.
 */
func NewErrorInfoFromURI(errorInfo address.URI) *ErrorInfo {
	this := &ErrorInfo{}
	this.Parameters.super(core.SIPHeaderNames_ERROR_INFO)
	this.errorInfo = errorInfo
	return this
}

func (this *ErrorInfo) String() string {
	return this.headerName + core.SIPSeparatorNames_COLON + core.SIPSeparatorNames_SP + this.EncodeBody() + core.SIPSeparatorNames_NEWLINE
}

/**
 * Encode into canonical form.
 * @return String
 */
func (this *ErrorInfo) EncodeBody() string {
	var encoding bytes.Buffer

	encoding.WriteString(core.SIPSeparatorNames_LESS_THAN)

	encoding.WriteString(this.errorInfo.String())

	encoding.WriteString(core.SIPSeparatorNames_GREATER_THAN)

	if this.parameters != nil && this.parameters.Len() > 0 {
		encoding.WriteString(core.SIPSeparatorNames_SEMICOLON)
		encoding.WriteString(this.parameters.String())
	}

	return encoding.String()
}

/**
 * Sets the ErrorInfo of the ErrorInfoHeader to the <var>errorInfo</var>
 * parameter value.
 *
 * @param errorInfo the new ErrorInfo of this ErrorInfoHeader.
 */
func (this *ErrorInfo) SetErrorInfo(errorInfo address.URI) {
	this.errorInfo = errorInfo
}

/**
 * Returns the ErrorInfo value of this ErrorInfoHeader. This message
 * may return null if a String message identifies the ErrorInfo.
 *
 * @return the URI representing the ErrorInfo.
 */
func (this *ErrorInfo) GetErrorInfo() address.URI {
	return this.errorInfo
}

/**
 * Sets the Error information message to the new <var>message</var> value
 * supplied to this method.
 *
 * @param message - the new string value that represents the error message.
 * @throws ParseException which signals that an error has been reached
 * unexpectedly while parsing the error message.
 */
func (this *ErrorInfo) SetErrorMessage(message string) { //throws ParseException {
	//if (message==null) throw new  NullPointerException("JAIN-SIP Exception "+
	//", ErrorInfoHeader, setErrorMessage(), the message parameter is null");
	this.SetParameter("message", message)
}

/**
 * Get the Error information message of this ErrorInfoHeader.
 *
 * @return the stringified version of the ErrorInfo header.
 */
func (this *ErrorInfo) GetErrorMessage() string {
	return this.GetParameter("message")
}
