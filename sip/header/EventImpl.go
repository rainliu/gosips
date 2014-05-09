package header

import (
	"bytes"
	"gosips/core"
	"strings"
)

/**
* Event SIP Header.
*
*@version  JAIN-SIP-1.1
*
*@author M. Ranganathan <mranga@nist.gov>  <br/>
*@author Olivier Deruelle <deruelle@nist.gov><br/>
*<a href="{@docRoot}/uncopyright.html">This code is in the public domain.</a>
*
 */
type Event struct {
	Parameters // implements EventHeader{

	eventType string
}

/** Creates a new instance of Event */
func NewEvent() *Event {
	this := &Event{}
	this.Parameters.super(core.SIPHeaderNames_EVENT)
	return this
}

/**
 * Sets the eventType to the newly supplied eventType string.
 *
 * @param eventType - the  new string defining the eventType supported
 * in this EventHeader
 * @throws ParseException which signals that an error has been reached
 * unexpectedly while parsing the eventType value.
 */
func (this *Event) SetEventType(eventType string) { //throws ParseException {
	//if (eventType==null)
	//throw new  NullPointerException( " the eventType is null");
	this.eventType = eventType
}

/**
 * Gets the eventType of the EventHeader.
 *
 * @return the string object identifing the eventType of EventHeader.
 */
func (this *Event) GetEventType() string {
	return this.eventType
}

/**
 * Sets the id to the newly supplied <var>eventId</var> string.
 *
 * @param eventId - the new string defining the eventId of this EventHeader
 * @throws ParseException which signals that an error has been reached
 * unexpectedly while parsing the eventId value.
 */
func (this *Event) SetEventId(eventId string) { //throws ParseException {
	//if (eventId==null)
	//throw new  NullPointerException( " the eventId parameter is null");
	this.SetParameter(ParameterNames_ID, eventId)
}

/**
 * Gets the id of the EventHeader. This method may return null if the
 * "eventId" is not Set.
 * @return the string object identifing the eventId of EventHeader.
 */
func (this *Event) GetEventId() string {
	return this.GetParameter(ParameterNames_ID)
}

func (this *Event) String() string {
	return this.headerName + core.SIPSeparatorNames_COLON + core.SIPSeparatorNames_SP + this.EncodeBody() + core.SIPSeparatorNames_NEWLINE
}

/**
 * Encode in canonical form.
 * @return String
 */
func (this *Event) EncodeBody() string {
	var encoding bytes.Buffer

	if this.eventType != "" {
		encoding.WriteString(this.eventType)
	}

	if this.parameters != nil && this.parameters.Len() > 0 {
		encoding.WriteString(core.SIPSeparatorNames_SEMICOLON)
		encoding.WriteString(this.parameters.String())
	}

	return encoding.String()
}

/** Return true if the given event header matches the supplied one.
 *
 * @param matchTarGet -- event header to match against.
 */
func (this *Event) Match(matchTarGet *Event) bool {
	if matchTarGet == nil {
		return false
	} else {
		return strings.ToLower(matchTarGet.GetEventType()) == strings.ToLower(this.GetEventType()) &&
			strings.ToLower(this.GetEventId()) == strings.ToLower(matchTarGet.GetEventId())
	}
}
