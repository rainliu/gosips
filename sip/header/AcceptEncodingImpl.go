package header

import (
	"bytes"
	"errors"
	"gosips/core"
	"strconv"
)

/**
* Accept-Encoding SIP (HTTP) Header.
* <pre>
* From HTTP RFC 2616
*
*
*   The Accept-Encoding request-header field is similar to Accept, but
*   restricts the content-codings (section 3.5) that are acceptable in
*   the response.
*
*
*       Accept-Encoding  = "Accept-Encoding" ":"
*
*
*                          1#( codings [ ";" "q" "=" qvalue ] )
*       codings          = ( content-coding | "*" )
*
*   Examples of its use are:
*
*       Accept-Encoding: compress, gzip
*       Accept-Encoding:
*       Accept-Encoding: *
*       Accept-Encoding: compress;q=0.5, gzip;q=1.0
*       Accept-Encoding: gzip;q=1.0, identity; q=0.5, *;q=0
* </pre>
*
 */
type AcceptEncoding struct {
	Parameters

	/** contentEncoding field
	 */
	contentCoding string
}

/** default constructor
 */
func NewAcceptEncoding() *AcceptEncoding {
	this := &AcceptEncoding{}
	this.Parameters.super(core.SIPHeaderNames_ACCEPT_ENCODING)
	return this
}

func (this *AcceptEncoding) String() string {
	return this.headerName + core.SIPSeparatorNames_COLON +
		core.SIPSeparatorNames_SP + this.EncodeBody() + core.SIPSeparatorNames_NEWLINE
}

/** Encode the value of this header.
*@return the value of this header encoded into a string.
 */
func (this *AcceptEncoding) EncodeBody() string {
	var encoding bytes.Buffer
	if this.contentCoding != "" {
		encoding.WriteString(this.contentCoding)
	}
	if this.Parameters.parameters != nil && this.parameters.Len() > 0 {
		encoding.WriteString(core.SIPSeparatorNames_SEMICOLON)
		encoding.WriteString(this.parameters.String())
	}
	return encoding.String()
}

/** get ContentEncoding field
 * @return String
 */
func (this *AcceptEncoding) GetEncoding() string {
	return this.contentCoding
}

/**
 * Sets the encoding of an EncodingHeader.
 *
 * @param encoding - the new string value defining the encoding.
 * @throws ParseException which signals that an error has been reached
 * unexpectedly while parsing the encoding value.
 */
func (this *AcceptEncoding) SetEncoding(encoding string) (ParseException error) {
	if encoding == "" {
		return errors.New("encoding parameter is null")
	}
	this.contentCoding = encoding
	return nil
}


/** get the QValue field. Return -1 if the parameter has not been
 * set.
 * @return float
 */
func (this *AcceptEncoding) GetQValue() float32 {
	if !this.HasParameter(ParameterNames_Q) {
		return -1
	}
	qstr := this.GetParameterValue(ParameterNames_Q)
	q, _ := strconv.ParseFloat(qstr, 32)
	return float32(q)
}

/**
 * Return true if the q value has been set.
 * @return boolean
 */
func (this *AcceptEncoding) HasQValue() bool {
	return this.HasParameter(ParameterNames_Q)
}

/**
 * Remove the q value.
 */
func (this *AcceptEncoding) RemoveQValue() {
	this.RemoveParameter(ParameterNames_Q)
}

/**
 * Sets q-value for media-range. Q-values allow the
 *
 * user to indicate the relative degree of preference for that media-range,
 *
 * using the qvalue scale from 0 to 1. If no q-value is present, the
 *
 * media-range should be treated as having a q-value of 1.
 *
 *
 *
 * @param qValue - the new float value of the q-value
 *
 * @throws InvalidArgumentException if the q parameter value is not between <code>0 and 1</code>.
 *
 */
func (this *AcceptEncoding) SetQValue(q float32) (InvalidArgumentException error) {
	if q < 0.0 || q > 1.0 {
		return errors.New("qvalue out of range!")
	}
	this.SetParameter(ParameterNames_Q, strconv.FormatFloat(float64(q), 'f', -1, 32))
	return nil
}
