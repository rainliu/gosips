package header

import (
	"gosip/core"
	"time"
)

/**
* Date Header.
*
*@version  JAIN-SIP-1.1
*
*@author M. Ranganathan <mranga@nist.gov>  <br/>
*@author Olivier Deruelle <deruelle@nist.gov><br/>
*<a href="{@docRoot}/uncopyright.html">This code is in the public domain.</a>
*
 */
type Date struct {
	SIPHeader

	/** date field
	 */
	date *time.Time
}

/** Default constructor.
 */
func NewDate() *Date {
	this := &Date{}
	this.SIPHeader.super(core.SIPHeaderNames_DATE)
	return this
}

func (this *Date) String() string {
	return this.headerName + core.SIPSeparatorNames_COLON +
		core.SIPSeparatorNames_SP + this.EncodeBody() + core.SIPSeparatorNames_NEWLINE
}

/** Encode the header into a String.
 * @return String
 */
func (this *Date) EncodeBody() string {
	return this.date.String()
}

/**
 * Set the date member
 * @param d SIPDate to set
 */
func (this *Date) SetDate(d *time.Time) {
	this.date = d
}

/**
 * Gets the date of DateHeader. The date is repesented by the Calender
 * object.
 *
 * @return the Calendar object representing the date of DateHeader
 */
func (this *Date) GetDate() *time.Time {
	return this.date
}
