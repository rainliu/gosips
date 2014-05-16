package header

import (
	"errors"
	"gosips/core"
)

/**
* AllowEvents Header.
 */
type AllowEvents struct {
	SIPHeader

	/** method field
	 */
	eventType string
}

/** default constructor
 */
func NewAllowEvents() *AllowEvents {
	this := &AllowEvents{}
	this.SIPHeader.super(core.SIPHeaderNames_ALLOW_EVENTS)
	return this
}

/** constructor
 * @param m String to set
 */
func NewAllowEventsFromString(m string) *AllowEvents {
	this := &AllowEvents{}
	this.SIPHeader.super(core.SIPHeaderNames_ALLOW_EVENTS)
	this.eventType = m
	return this
}

/**
 * Sets the eventType defined in this AllowEventsHeader.
 *
 * @param eventType - the String defining the method supported
 * in this AllowEventsHeader
 * @throws ParseException which signals that an error has been reached
 * unexpectedly while parsing the Strings defining the eventType supported
 */
func (this *AllowEvents) SetEventType(eventType string) (ParseException error) {
	if eventType == "" {
		errors.New("NullPointerException: the eventType parameter is null")
	}
	this.eventType = eventType
	return nil
}

/**
 * Gets the eventType of the AllowEventsHeader.
 *
 * @return the String object identifing the eventTypes of AllowEventsHeader.
 */
func (this *AllowEvents) GetEventType() string {
	return this.eventType
}

/** Return body encoded in canonical form.
 * @return body encoded as a string.
 */
func (this *AllowEvents) EncodeBody() string {
	return this.eventType
}
