package header

import (
	"bytes"
	"errors"
	"gosips/core"
	"strconv"
)

/**
*Accept header : The top level header is actually AcceptList which is a list of
*Accept headers.
 */
type Accept struct {
	Parameters

	/** mediaRange field
	 */
	mediaRange *MediaRange
}

/** default constructor
 */
func NewAccept() *Accept {
	this := &Accept{}
	this.Parameters.super(core.SIPHeaderNames_ACCEPT)
	return this
}

/** returns true if this header allows all ContentTypes,
 * false otherwise.
 * @return Boolean
 */
func (this *Accept) AllowsAllContentTypes() bool {
	if this.mediaRange == nil {
		return false
	} else {
		return this.mediaRange.GetType() == core.SIPSeparatorNames_STAR
	}
}

/**
 * returns true if this header allows all ContentSubTypes,
 * false otherwise.
 * @return boolean
 */
func (this *Accept) AllowsAllContentSubTypes() bool {
	if this.mediaRange == nil {
		return false
	} else {
		return this.mediaRange.GetSubtype() == core.SIPSeparatorNames_STAR
	}
}

func (this *Accept) String() string {
	return this.headerName + core.SIPSeparatorNames_COLON +
		core.SIPSeparatorNames_SP + this.EncodeBody() + core.SIPSeparatorNames_NEWLINE
}

/** Encode the value of this header into cannonical form.
*@return encoded value of the header as a string.
 */
func (this *Accept) EncodeBody() string {
	var encoding bytes.Buffer
	if this.mediaRange != nil {
		encoding.WriteString(this.mediaRange.String())
	}
	if this.parameters != nil && this.parameters.Len() > 0 {
		encoding.WriteString(core.SIPSeparatorNames_SEMICOLON)
		encoding.WriteString(this.parameters.String())
	}
	return encoding.String()
}

/** get the MediaRange field
 * @return MediaRange
 */
func (this *Accept) GetMediaRange() *MediaRange {
	return this.mediaRange
}

/** get the contentType field
 * @return String
 */
func (this *Accept) GetContentType() string {
	if this.mediaRange == nil {
		return ""
	} else {
		return this.mediaRange.GetType()
	}
}

/** get the ContentSubType fiels
 * @return String
 */
func (this *Accept) GetContentSubType() string {
	if this.mediaRange == nil {
		return ""
	} else {
		return this.mediaRange.GetSubtype()
	}
}

/** set the ContentSubType field
 * @param subtype String to set
 */
func (this *Accept) SetContentSubType(subtype string) {
	if this.mediaRange == nil {
		this.mediaRange = NewMediaRange()
	}
	this.mediaRange.SetSubtype(subtype)
}

/** set the ContentType field
 * @param type String to set
 */
func (this *Accept) SetContentType(mtype string) {
	if this.mediaRange == nil {
		this.mediaRange = NewMediaRange()
	}
	this.mediaRange.SetType(mtype)
}

/**
 * Set the mediaRange member
 * @param m MediaRange field
 */
func (this *Accept) SetMediaRange(m *MediaRange) {
	this.mediaRange = m
}

/** get the QValue field. Return -1 if the parameter has not been
 * set.
 * @return float
 */
func (this *Accept) GetQValue() float32 {
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
func (this *Accept) HasQValue() bool {
	return this.HasParameter(ParameterNames_Q)
}

/**
 * Remove the q value.
 */
func (this *Accept) RemoveQValue() {
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
func (this *Accept) SetQValue(q float32) (InvalidArgumentException error) {
	if q < 0.0 || q > 1.0 {
		return errors.New("qvalue out of range!")
	}
	this.SetParameter(ParameterNames_Q, strconv.FormatFloat(float64(q), 'f', -1, 32))
	return nil
}
