package header

import (
	"gosip/core"
	"strconv"
)

/**
* MimeVersion SIP Header.
*
*@version  JAIN-SIP-1.1
*
*@author M. Ranganathan <mranga@nist.gov>  <br/>
*@author Olivier Deruelle <deruelle@nist.gov><br/>
*
*<a href="{@docRoot}/uncopyright.html">This code is in the public domain.</a>
*
 */
type MimeVersion struct {
	SIPHeader
	//implements MimeVersionHeader{

	/** mimeVersion field
	 */
	minorVersion int

	/** majorVersion field
	 */
	majorVersion int
}

/** Default constructor
 */
func NewMimeVersion() *MimeVersion {
	this := &MimeVersion{}
	this.SIPHeader.super(core.SIPHeaderNames_MIME_VERSION)
	return this
}

/**
 * Gets the Minor version value of this MimeVersionHeader.
 *
 * @return the Minor version of this MimeVersionHeader
 */
func (this *MimeVersion) GetMinorVersion() int {
	return this.minorVersion
}

/**
 * Gets the Major version value of this MimeVersionHeader.
 *
 * @return the Major version of this MimeVersionHeader
 */
func (this *MimeVersion) GetMajorVersion() int {
	return this.majorVersion
}

/**
 * Sets the Minor-Version argument of this MimeVersionHeader to the supplied
 * <var>minorVersion</var> value.
 *
 * @param minorVersion - the new integer Minor version
 * @throws InvalidArgumentException
 */
func (this *MimeVersion) SetMinorVersion(minorVersion int) { //throws InvalidArgumentException {
	// if (minorVersion<0) throw new InvalidArgumentException("JAIN-SIP Exception"+
	//", MimeVersion, SetMinorVersion(), the minorVersion parameter is null");
	this.minorVersion = minorVersion
}

/**
 * Sets the Major-Version argument of this MimeVersionHeader to the supplied
 * <var>majorVersion</var> value.
 *
 * @param majorVersion - the new integer Major version
 * @throws InvalidArgumentException
 */
func (this *MimeVersion) SetMajorVersion(majorVersion int) { //throws InvalidArgumentException {
	//if (majorVersion<0) throw new InvalidArgumentException("JAIN-SIP Exception"+
	//", MimeVersion, SetMajorVersion(), the majorVersion parameter is null");
	this.majorVersion = majorVersion
}

func (this *MimeVersion) String() string {
	return this.headerName + core.SIPSeparatorNames_COLON +
		core.SIPSeparatorNames_SP + this.EncodeBody() + core.SIPSeparatorNames_NEWLINE
}

/**
 * Return canonical form.
 * @return String
 */
func (this *MimeVersion) EncodeBody() string {
	return strconv.Itoa(this.majorVersion) + core.SIPSeparatorNames_DOT +
		strconv.Itoa(this.minorVersion)
}
