package header

import (
	"bytes"
	"gosips/core"
	"gosips/sip/address"
)

/**
* AlertInfo SIP Header.
 */
type AlertInfo struct {
	Parameters

	/** URI field
	 */
	uri *address.URIImpl
}

/** Constructor
 */
func NewAlertInfo() *AlertInfo {
	this := &AlertInfo{}
	this.Parameters.super(core.SIPHeaderNames_ALERT_INFO)
	return this
}

func (this *AlertInfo) String() string {
	return this.headerName + core.SIPSeparatorNames_COLON +
		core.SIPSeparatorNames_SP + this.EncodeBody() + core.SIPSeparatorNames_NEWLINE
}

/**
 * Return value encoding in canonical form.
 * @return The value of the header in canonical encoding.
 */
func (this *AlertInfo) EncodeBody() string {
	var encoding bytes.Buffer
	encoding.WriteString(core.SIPSeparatorNames_LESS_THAN)
	encoding.WriteString(this.uri.String())
	encoding.WriteString(core.SIPSeparatorNames_GREATER_THAN)

	if this.parameters != nil && this.parameters.Len() > 0 {
		encoding.WriteString(core.SIPSeparatorNames_SEMICOLON)
		encoding.WriteString(this.parameters.String())
	}
	return encoding.String()

}

/**
 * Set the uri member
 * @param u URI to set
 */
func (this *AlertInfo) SetAlertInfo(uri address.URI) {
	this.uri = uri.(*address.URIImpl)
}

/**
 * Returns the AlertInfo value of this AlertInfoHeader.
 *
 *
 *
 * @return the URI representing the AlertInfo.
 *
 */
func (this *AlertInfo) GetAlertInfo() address.URI {
	return this.uri
}
