package header

import (
	"errors"
	"gosips/core"
)

/**
*Organization SIP Header.
 */
type Organization struct {
	SIPHeader

	/** organization field
	 */
	organization string
}

/** Default constructor
 */
func NewOrganization() *Organization {
	this := &Organization{}
	this.SIPHeader.super(core.SIPHeaderNames_ORGANIZATION)
	return this
}

func (this *Organization) String() string {
	return this.headerName + core.SIPSeparatorNames_COLON +
		core.SIPSeparatorNames_SP + this.EncodeBody() + core.SIPSeparatorNames_NEWLINE
}

/**
 * Return encoding of value of the header.
 * @return String
 */
func (this *Organization) EncodeBody() string {
	return this.organization
}

/** get the organization field.
 * @return String
 */
func (this *Organization) GetOrganization() string {
	return this.organization
}

/**
 * Set the organization member
 * @param o String to set
 */
func (this *Organization) SetOrganization(o string) (ParseException error) {
	if o == "" {
		return errors.New("NullPointerException: the organization parameter is null")
	}
	this.organization = o
	return nil
}
