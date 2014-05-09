package header

import (
	"bytes"
	"gosips/core"
	"strconv"
)

/**
*SubscriptionState header
*
*@version  JAIN-SIP-1.1
*
*@author Olivier Deruelle <deruelle@nist.gov><br/>
*<a href="{@docRoot}/uncopyright.html">This code is in the public domain.</a>
*
 */
type SubscriptionState struct {
	Parameters //Header
	//implements SubscriptionStateHeader{

	expires    int
	retryAfter int
	reasonCode string
	state      string
}

/** Creates a new instance of SubscriptionState */
func NewSubscriptionState() *SubscriptionState {
	this := &SubscriptionState{}
	this.Parameters.super(core.SIPHeaderNames_SUBSCRIPTION_STATE)
	this.expires = -1
	this.retryAfter = -1
	return this
}

/**
 * Sets the relative expires value of the SubscriptionStateHeader. The
 * expires value MUST be greater than zero and MUST be less than 2**31.
 *
 * @param expires - the new expires value of this SubscriptionStateHeader.
 * @throws InvalidArgumentException if supplied value is less than zero.
 */
func (this *SubscriptionState) SetExpires(expires int) { //throws InvalidArgumentException {
	// if (expires<=0) throw new InvalidArgumentException("JAIN-SIP "+
	// "Exception, SubscriptionState, SetExpires(), the expires parameter is <=0");
	this.expires = expires
}

/**
 * Gets the expires value of the SubscriptionStateHeader. This expires value is
 * relative time.
 *
 * @return the expires value of the SubscriptionStateHeader.
 */
func (this *SubscriptionState) GetExpires() int {
	return this.expires
}

/**
 * Sets the retry after value of the SubscriptionStateHeader. The retry after value
 * MUST be greater than zero and MUST be less than 2**31.
 *
 * @param retryAfter - the new retry after value of this SubscriptionStateHeader
 * @throws InvalidArgumentException if supplied value is less than zero.
 */
func (this *SubscriptionState) SetRetryAfter(retryAfter int) { //throws InvalidArgumentException {
	// if (retryAfter<=0) throw new InvalidArgumentException("JAIN-SIP "+
	// "Exception, SubscriptionState, SetRetryAfter(), the retryAfter parameter is <=0");
	this.retryAfter = retryAfter
}

/**
 * Gets the retry after value of the SubscriptionStateHeader. This retry after
 * value is relative time.
 *
 * @return the retry after value of the SubscriptionStateHeader.
 */
func (this *SubscriptionState) GetRetryAfter() int {
	return this.retryAfter
}

/**
 * Gets the reason code of SubscriptionStateHeader.
 *
 * @return the comment of this SubscriptionStateHeader, return null if no reason code
 * is available.
 */
func (this *SubscriptionState) GetReasonCode() string {
	return this.reasonCode
}

/**
 * Sets the reason code value of the SubscriptionStateHeader.
 *
 * @param reasonCode - the new reason code string value of the SubscriptionStateHeader.
 * @throws ParseException which signals that an error has been reached
 * unexpectedly while parsing the reason code.
 */
func (this *SubscriptionState) SetReasonCode(reasonCode string) { //throws ParseException {
	// if (reasonCode==null) throw new  NullPointerException("JAIN-SIP "+
	// "Exception, SubscriptionState, SetReasonCode(), the reasonCode parameter is null");
	this.reasonCode = reasonCode
}

/**
 * Gets the state of SubscriptionStateHeader.
 *
 * @return the state of this SubscriptionStateHeader.
 */
func (this *SubscriptionState) GetState() string {
	return this.state
}

/**
 * Sets the state value of the SubscriptionStateHeader.
 *
 * @param state - the new state string value of the SubscriptionStateHeader.
 * @throws ParseException which signals that an error has been reached
 * unexpectedly while parsing the state.
 */
func (this *SubscriptionState) SetState(state string) { //throws ParseException{
	// if (state==null) throw new  NullPointerException("JAIN-SIP "+
	// "Exception, SubscriptionState, SetState(), the state parameter is null");
	this.state = state
}

func (this *SubscriptionState) String() string {
	return this.headerName + core.SIPSeparatorNames_COLON +
		core.SIPSeparatorNames_SP + this.EncodeBody() + core.SIPSeparatorNames_NEWLINE
}

/** Just the encoded body of the header.
 * @return the string encoded header body.
 */
func (this *SubscriptionState) EncodeBody() string {
	var encoding bytes.Buffer //  = new StringBuffer();
	if this.state != "" {
		encoding.WriteString(this.state)
	}
	if this.reasonCode != "" {
		encoding.WriteString(";reason=" + this.reasonCode)
	}
	if this.retryAfter != -1 {
		encoding.WriteString(";retry-after=" + strconv.Itoa(this.retryAfter))
	}
	if this.expires != -1 {
		encoding.WriteString(";expires=" + strconv.Itoa(this.expires))
	}

	if this.parameters != nil && this.parameters.Len() > 0 {
		encoding.WriteString(core.SIPSeparatorNames_SEMICOLON)
		encoding.WriteString(this.parameters.String())
	}
	return encoding.String()
}
