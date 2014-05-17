package header

import (
	"bytes"
	"errors"
	"gosips/core"
	"gosips/sip/address"
)

/**
* CallInfo Header.
 */

type CallInfo struct {
	Parameters

	info *address.URIImpl
}

/** Default constructor
 */
func NewCallInfo() *CallInfo {
	this := &CallInfo{}
	this.Parameters.super(core.SIPHeaderNames_CALL_INFO)
	return this
}

func (this *CallInfo) String() string {
	return this.headerName + core.SIPSeparatorNames_COLON +
		core.SIPSeparatorNames_SP + this.EncodeBody() + core.SIPSeparatorNames_NEWLINE
}

/**
 * Return canonical representation.
 * @return String
 */
func (this *CallInfo) EncodeBody() string {
	var encoding bytes.Buffer

	encoding.WriteString(core.SIPSeparatorNames_LESS_THAN)

	encoding.WriteString(this.info.String())

	encoding.WriteString(core.SIPSeparatorNames_GREATER_THAN)

	if this.parameters != nil && this.parameters.Len() > 0 {
		encoding.WriteString(core.SIPSeparatorNames_SEMICOLON)
		encoding.WriteString(this.parameters.String())
	}

	return encoding.String()
}

/** get the purpose field
 * @return String
 */
func (this *CallInfo) GetPurpose() string {
	return this.GetParameter("purpose")
}

/** get the URI field
 * @return URI
 */
func (this *CallInfo) GetInfo() address.URI {
	return this.info
}

/** set the purpose field
 * @param purpose is the purpose field.
 */
func (this *CallInfo) SetPurpose(purpose string) error {
	if purpose == "" {
		return errors.New("NullPointerException: null arg")
	}
	this.SetParameter("purpose", purpose)
	return nil
}

/** set the URI field
 * @param info is the URI to set.
 */
func (this *CallInfo) SetInfo(info address.URI) {
	this.info = info.(*address.URIImpl)
}
