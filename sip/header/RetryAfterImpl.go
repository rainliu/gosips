package header

import (
	"bytes"
	"errors"
	"gosips/core"
	"strconv"
)

/**  Retry-After SIP Header.
 */
type RetryAfter struct {
	Parameters

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
	var encoding bytes.Buffer
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

func (this *RetryAfter) SetRetryAfter(retryAfter int) (InvalidArgumentException error) {
	if retryAfter < 0 {
		return errors.New("InvalidArgumentException: invalid parameter")
	}
	this.retryAfter = retryAfter
	return nil
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

func (this *RetryAfter) SetComment(comment string) (ParseException error) {
	if comment == "" {
		return errors.New("NullPointerException: the comment parameter is null")
	}
	this.comment = comment
	return nil
}

/**
 * Sets the duration value of the RetryAfterHeader. The retry after value
 * MUST be greater than zero and MUST be less than 2**31.
 *
 * @param duration - the new duration value of this RetryAfterHeader
 * @throws InvalidArgumentException if supplied value is less than zero.
 * @since JAIN SIP v1.1
 */

func (this *RetryAfter) SetDuration(duration int) (InvalidArgumentException error) {
	if duration < 0 {
		return errors.New("InvalidArgumentException: the duration parameter is <0")
	}
	this.Parameters.SetParameter(ParameterNames_DURATION, strconv.Itoa(duration))
	return nil
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
