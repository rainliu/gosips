package header

import (
	"bytes"
	"errors"
	"gosips/core"
)

/**
* Content Dispositon SIP Header.
 */
type ContentDisposition struct {
	Parameters

	/** dispositionType field.
	 */
	dispositionType string
}

/** Default constructor.
 */
func NewContentDisposition() *ContentDisposition {
	this := &ContentDisposition{}
	this.Parameters.super(core.SIPHeaderNames_CONTENT_DISPOSITION)
	return this
}

func (this *ContentDisposition) String() string {
	return this.headerName + core.SIPSeparatorNames_COLON +
		core.SIPSeparatorNames_SP + this.EncodeBody() + core.SIPSeparatorNames_NEWLINE
}

/**
 * Encode value of header into canonical string.
 * @return encoded value of header.
 *
 */
func (this *ContentDisposition) EncodeBody() string {
	var encoding bytes.Buffer

	if this.dispositionType != "" {
		encoding.WriteString(this.dispositionType)
	}

	if this.Parameters.parameters != nil && this.parameters.Len() > 0 {
		encoding.WriteString(core.SIPSeparatorNames_SEMICOLON)
		encoding.WriteString(this.parameters.String())
	}
	return encoding.String()
}

/** Set the disposition type.
 *@param disposition type.
 */
func (this *ContentDisposition) SetDispositionType(dispositionType string) (ParseException error) {
	if dispositionType == "" {
		return errors.New("NullPointerException: the dispositionType parameter is null")
	}
	this.dispositionType = dispositionType
	return nil
}

/** Get the disposition type.
 *@param GetDispositionType
 */
func (this *ContentDisposition) GetDispositionType() string {
	return this.dispositionType
}

/** Get the dispositionType field.
 * @return String
 */
func (this *ContentDisposition) GetHandling() string {
	return this.GetParameter("handling")
}

/** Set the dispositionType field.
 * @param type String to Set.
 */
func (this *ContentDisposition) SetHandling(handling string) (ParseException error) {
	if handling == "" {
		return errors.New("NullPointerException: the handling parameter is null")
	}
	this.SetParameter("handling", handling)
	return nil
}

/**
 * Gets the interpretation of the message body or message body part of
 *
 * this ContentDispositionHeader.
 *
 * @return interpretation of the message body or message body part
 *
 */
func (this *ContentDisposition) GetContentDisposition() string {
	return this.EncodeBody()
}
