package header

import (
	"bytes"
	"gosip/core"
	"strconv"
)

/**
*Accept header : The top level header is actually AcceptList which is a list of
*Accept headers.
*
*@version JAIN-SIP-1.1
*
*@author M. Ranganathan <mranga@nist.gov>  <br/>
*
*<a href="{@docRoot}/uncopyright.html">This code is in the public domain.</a>
*
 */
type Accept struct {
	Parameters
	//implements javax.sip.header.AcceptHeader{

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

/** Encode the value of this header into cannonical form.
*@return encoded value of the header as a string.
 */
func (this *Accept) EncodeBody() string {
	//String s="";
	var encoding bytes.Buffer
	if this.mediaRange != nil {
		encoding.WriteString(this.mediaRange.Encode())
	}
	if this.parameters != nil && this.parameters.Len() > 0 {
		encoding.WriteString(core.SIPSeparatorNames_SEMICOLON)
		encoding.WriteString(this.parameters.String())
		//core.SIPSeparatorNames_SP + ";" + 
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

/**
 * Get the q value.
 * @return float
 */
func (this *Accept) GetQValue() float32 {
	q, err := strconv.ParseFloat(this.GetParameter(ParameterNames_Q), 32)
	if err != nil {
		return -1
	} else {
		return float32(q)
	}
}

/**
 * Return true if the q value has been set.
 * @return boolean
 */
func (this *Accept) HasQValue() bool {
	return this.HasParameter(ParameterNames_Q)

}

/**
 *Remove the q value.
 */
func (this *Accept) RemoveQValue() {
	this.RemoveParameter(ParameterNames_Q)
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
 * Set the q value
 * @param qValue float to set
 * @throws IllegalArgumentException if qValue is <0.0 or >1.0
 */
func (this *Accept) SetQValue(qValue float32) {
	//throws InvalidArgumentException {
	if qValue == -1 {
		this.RemoveParameter(ParameterNames_Q)
	}
	s := strconv.FormatFloat(float64(qValue), 'f', 4, 32)
	this.SetParameter(ParameterNames_Q, s)

}

/**
 * Set the mediaRange member
 * @param m MediaRange field
 */
func (this *Accept) SetMediaRange(m *MediaRange) {
	this.mediaRange = m
}
