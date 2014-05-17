package header

import (
	"errors"
	"gosips/core"
)

/**
* the Priority header.
 */
type Priority struct {
	SIPHeader

	/** priority field
	 */
	priority string
}

/** Default constructor
 */
func NewPriority() *Priority {
	this := &Priority{}
	this.SIPHeader.super(core.SIPHeaderNames_PRIORITY)
	return this
}

func (this *Priority) String() string {
	return this.headerName + core.SIPSeparatorNames_COLON +
		core.SIPSeparatorNames_SP + this.EncodeBody() + core.SIPSeparatorNames_NEWLINE
}

/**
 * Encode into canonical form.
 * @return String
 */
func (this *Priority) EncodeBody() string {
	return this.priority
}

/**
 * get the priority value.
 * @return String
 */
func (this *Priority) GetPriority() string {
	return this.priority
}

/**
 * Set the priority member
 * @param p String to set
 */
func (this *Priority) SetPriority(p string) (ParseException error) {
	if p == "" {
		return errors.New("NullPointerException: the priority parameter is null")
	}
	this.priority = p
	return nil
}
