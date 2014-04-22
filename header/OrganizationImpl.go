package header

import "gosip/core"

/**
*Organization SIP Header.
*
*@version  JAIN-SIP-1.1
*
*@author M. Ranganathan <mranga@nist.gov>  <br/>
*@author Olivier Deruelle <deruelle@nist.gov><br/>
*
*<a href="{@docRoot}/uncopyright.html">This code is in the public domain.</a>
*
 */
type Organization struct {
	SIPHeaderImpl
	//implements OrganizationHeader {

	/** organization field
	 */
	organization string
}

/** Default constructor
 */
func NewOrganization() *Organization {
	this := &Organization{}
	this.SIPHeaderImpl.super(core.SIPHeaderNames_ORGANIZATION)
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
func (this *Organization) SetOrganization(o string) { //throws ParseException{
	// if (o==null) throw new  NullPointerException("JAIN-SIP Exception,"+
	// " Organization, setOrganization(), the organization parameter is null");
	this.organization = o
}
