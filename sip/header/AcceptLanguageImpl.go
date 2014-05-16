package header

import (
	"bytes"
	"errors"
	"gosips/core"
	"strconv"
	"strings"
)

/**
 * Accept Language body.
 * <pre>
 * HTTP RFC 2616 Section 14.4
 * Accept-Language = "Accept-Language" ":"
 *                         1#( language-range [ ";" "q" "=" qvalue ] )
 *       language-range  = ( ( 1*8ALPHA *( "-" 1*8ALPHA ) ) | "*" )
 *
 * </pre>
 *
 * @see AcceptLanguageList
 */
type AcceptLanguage struct {
	Parameters

	/** languageRange field
	 */
	languageRange string
}

/** default constructor
 */
func NewAcceptLanguage() *AcceptLanguage {
	this := &AcceptLanguage{}
	this.Parameters.super(core.SIPHeaderNames_ACCEPT_LANGUAGE)
	return this
}

func (this *AcceptLanguage) String() string {
	return this.headerName + core.SIPSeparatorNames_COLON +
		core.SIPSeparatorNames_SP + this.EncodeBody() + core.SIPSeparatorNames_NEWLINE
}

/** Encode the value of this header to a string.
 *@return  encoded header as a string.
 */
func (this *AcceptLanguage) EncodeBody() string {
	var encoding bytes.Buffer
	if this.languageRange != "" {
		encoding.WriteString(this.languageRange)
	}
	if this.parameters.Len() > 0 {
		encoding.WriteString(core.SIPSeparatorNames_SEMICOLON)
		encoding.WriteString(this.parameters.String())
	}
	return encoding.String()
}

/** get the LanguageRange field
 * @return String
 */
func (this *AcceptLanguage) GetLanguageRange() string {
	return this.languageRange
}

/**
 * Set the languageRange.
 *
 * @param languageRange is the language range to set.
 *
 */
func (this *AcceptLanguage) SetLanguageRange(languageRange string) {
	this.languageRange = strings.TrimSpace(languageRange)
}

/**
 * Gets the language value of the AcceptLanguageHeader.
 *
 *
 *
 * @return the language Locale value of this AcceptLanguageHeader
 *
 */
func (this *AcceptLanguage) GetAcceptLanguage() string {
	if this.languageRange == "" {
		return ""
	} else {
		return this.languageRange
	}
}

/**
 * Sets the language parameter of this AcceptLanguageHeader.
 *
 *
 *
 * @param language - the new Locale value of the language of
 *
 * AcceptLanguageHeader
 *
 *
 */
func (this *AcceptLanguage) SetAcceptLanguage(language string) {
	this.languageRange = language
}


/** get the QValue field. Return -1 if the parameter has not been
 * set.
 * @return float
 */
func (this *AcceptLanguage) GetQValue() float32 {
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
func (this *AcceptLanguage) HasQValue() bool {
	return this.HasParameter(ParameterNames_Q)
}

/**
 * Remove the q value.
 */
func (this *AcceptLanguage) RemoveQValue() {
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
func (this *AcceptLanguage) SetQValue(q float32) (InvalidArgumentException error) {
	if q < 0.0 || q > 1.0 {
		return errors.New("qvalue out of range!")
	}
	this.SetParameter(ParameterNames_Q, strconv.FormatFloat(float64(q), 'f', -1, 32))
	return nil
}