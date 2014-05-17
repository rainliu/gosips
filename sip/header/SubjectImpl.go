package header

import (
	"errors"
	"gosips/core"
)

/**
*Supported SIP Header.
 */
type Subject struct {
	SIPHeader

	/** subject field
	 */
	subject string
}

/** Default Constructor.
 */
func NewSubject() *Subject {
	this := &Subject{}
	this.SIPHeader.super(core.SIPHeaderNames_SUBJECT)
	return this
}

func (this *Subject) String() string {
	return this.headerName + core.SIPSeparatorNames_COLON +
		core.SIPSeparatorNames_SP + this.EncodeBody() + core.SIPSeparatorNames_NEWLINE
}

/**
 * Generate the canonical form.
 * @return String.
 */
func (this *Subject) EncodeBody() string {
	return this.subject
}

/**

 * Sets the subject value of the SubjectHeader to the supplied string

 * subject value.

 *

 * @param subject - the new subject value of this header

 * @throws ParseException which signals that an error has been reached

 * unexpectedly while parsing the subject value.

 */

func (this *Subject) SetSubject(subject string) (ParseException error) {
	if subject == "" {
		return errors.New("NullPointerException: the subject parameter is null")
	}
	this.subject = subject
	return nil
}

/**

 * Gets the subject value of SubjectHeader

 *

 * @return subject of SubjectHeader

 */

func (this *Subject) GetSubject() string {
	return this.subject
}
