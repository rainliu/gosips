package header

import (
	"bytes"
	"gosips/core"
	"strconv"
)

/**  Retry-After SIP Header.
*
*@version  JAIN-SIP-1.1
*
*@author M. Ranganathan <mranga@nist.gov>  <br/>
*@author Olivier Deruelle <deruelle@nist.gov><br/>
*
*<a href="{@docRoot}/uncopyright.html">This code is in the public domain.</a>
*
 */
type RetryAfter struct {
	Parameters //Header implements RetryAfterHeader {

	/** constant DURATION parameter.
	 */
	//public static final String DURATION= ParameterNames.DURATION;

	/** duration field
	 */
	retryAfter int

	/** comment field
	 */
	comment string
}

/** Default constructor
 */
func NewRetryAfter() *RetryAfter {
	this := &RetryAfter{}
	this.Parameters.super(core.SIPHeaderNames_RETRY_AFTER)
	return this
}

func (this *RetryAfter) String() string {
	return this.headerName + core.SIPSeparatorNames_COLON +
		core.SIPSeparatorNames_SP + this.EncodeBody() + core.SIPSeparatorNames_NEWLINE
}

/** Encode body of this into cannonical form.
 * @return encoded body
 */
func (this *RetryAfter) EncodeBody() string {
	var encoding bytes.Buffer //  = new StringBuffer();
	if this.retryAfter != 0 {
		encoding.WriteString(strconv.Itoa(this.retryAfter))
	}
	if this.comment != "" {
		encoding.WriteString(core.SIPSeparatorNames_SP + "(")
		encoding.WriteString(this.comment)
		encoding.WriteString(")")
	}

	if this.parameters != nil && this.parameters.Len() > 0 {
		encoding.WriteString(core.SIPSeparatorNames_SEMICOLON)
		encoding.WriteString(this.parameters.String())
	}
	return encoding.String()
}

/** Boolean function
 * @return true if comment exist, false otherwise
 */
func (this *RetryAfter) HasComment() bool {
	return this.comment != ""
}

/** remove comment field
 */
func (this *RetryAfter) RemoveComment() {
	this.comment = ""
}

/** remove duration field
 */
func (this *RetryAfter) RemoveDuration() {
	this.Parameters.RemoveParameter(ParameterNames_DURATION)
}

/**
 * Sets the retry after value of the RetryAfterHeader.
 * The retry after value MUST be greater than zero and
 * MUST be less than 2**31.
 *
 * @param retryAfter - the new retry after value of this RetryAfterHeader
 * @throws InvalidArgumentException if supplied value is less than zero.
 * @since JAIN SIP v1.1
 */

func (this *RetryAfter) SetRetryAfter(retryAfter int) { //throws InvalidArgumentException {
	//if (retryAfter < 0) throw new InvalidArgumentException
	//("invalid parameter " + retryAfter);
	this.retryAfter = retryAfter
}

/**
 * Gets the retry after value of the RetryAfterHeader. This retry after
 * value is relative time.
 *
 * @return the retry after value of the RetryAfterHeader.
 * @since JAIN SIP v1.1
 */

func (this *RetryAfter) GetRetryAfter() int {
	return this.retryAfter
}

/**
 * Gets the comment of RetryAfterHeader.
 *
 * @return the comment of this RetryAfterHeader, return null if no comment
 * is available.
 */

func (this *RetryAfter) GetComment() string {
	return this.comment
}

/**
 * Sets the comment value of the RetryAfterHeader.
 *
 * @param comment - the new comment string value of the RetryAfterHeader.
 * @throws ParseException which signals that an error has been reached
 * unexpectedly while parsing the comment.
 */

func (this *RetryAfter) SetComment(comment string) { //throws ParseException {
	//if (comment==null) throw new  NullPointerException
	//  ("the comment parameter is null");
	this.comment = comment
}

/**
 * Sets the duration value of the RetryAfterHeader. The retry after value
 * MUST be greater than zero and MUST be less than 2**31.
 *
 * @param duration - the new duration value of this RetryAfterHeader
 * @throws InvalidArgumentException if supplied value is less than zero.
 * @since JAIN SIP v1.1
 */

func (this *RetryAfter) SetDuration(duration int) { //throws InvalidArgumentException {
	//if (duration < 0) throw new InvalidArgumentException
	//    ("the duration parameter is <0");
	this.Parameters.SetParameter(ParameterNames_DURATION, strconv.Itoa(duration))
}

/**
 * Gets the duration value of the RetryAfterHeader. This duration value
 * is relative time.
 *
 * @return the duration value of the RetryAfterHeader, return zero if not
 * set.
 * @since JAIN SIP v1.1
 */

func (this *RetryAfter) GetDuration() int {
	d, _ := strconv.Atoi(this.Parameters.GetParameter(ParameterNames_DURATION))

	return d
}
