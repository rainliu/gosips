package header

import "gosips/core"

/**
* InReplyTo SIP Header.
 */
type InReplyTo struct {
	SIPHeader

	callId *CallIdentifier
}

/** Default constructor
 */
func NewInReplyTo() *InReplyTo {
	this := &InReplyTo{}
	this.SIPHeader.super(core.SIPHeaderNames_IN_REPLY_TO)
	return this
}

/** constructor
 * @param cid CallIdentifier to Set
 */
func NewInReplyToFromCallIdentifier(cid *CallIdentifier) *InReplyTo {
	this := &InReplyTo{}
	this.SIPHeader.super(core.SIPHeaderNames_IN_REPLY_TO)
	this.callId = cid
	return this
}

/**
 * Sets the Call-Id of the InReplyToHeader. The CallId parameter uniquely
 * identifies a serious of messages within a dialogue.
 *
 * @param callId - the string value of the Call-Id of this InReplyToHeader.
 * @throws ParseException which signals that an error has been reached
 * unexpectedly while parsing the callId value.
 */
func (this *InReplyTo) SetCallId(callId string) (ParseException error) {
	this.callId, ParseException = NewCallIdentifier(callId)
	return ParseException
}

/**
 * Returns the Call-Id of InReplyToHeader. The CallId parameter uniquely
 * identifies a series of messages within a dialogue.
 *
 * @return the String value of the Call-Id of this InReplyToHeader
 */
func (this *InReplyTo) GetCallId() string {
	if this.callId == nil {
		return ""
	} else {
		return this.callId.String()
	}
}

func (this *InReplyTo) String() string {
	return this.headerName + core.SIPSeparatorNames_COLON +
		core.SIPSeparatorNames_SP + this.EncodeBody() + core.SIPSeparatorNames_NEWLINE
}

/**
 * Generate canonical form of the header.
 * @return String
 */
func (this *InReplyTo) EncodeBody() string {
	return this.callId.String()
}
