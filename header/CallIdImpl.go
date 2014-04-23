package header

import (
	"gosip/core"
)

/**
* Call ID SIPHeaderHeader.
 */

type CallID struct { // implements javax.sip.header.CallIdHeader {
	SIPHeader

	/** callIdentifier field
	 */
	callIdentifier *CallIdentifier
}

/** Default constructor
 */
/*func NewCallID () *CallID {
			this := &CallID{};
			this.SIPHeaderHeader.headerName = CALL_ID;
			return this;
        }*/

/** Constructor given the call Identifier.
*@param callId string call identifier (should be localid@host)
*@throws IllegalArgumentException if call identifier is bad.
 */
func NewCallID(callId string) (this *CallID, IllegalArgumentException error) {
	this = &CallID{}
	this.SIPHeader.super(core.SIPHeaderNames_CALL_ID)
	this.callIdentifier, IllegalArgumentException = NewCallIdentifier(callId)
	if IllegalArgumentException != nil {
		return nil, IllegalArgumentException
	} else {
		return this, nil
	}
}

func (this *CallID) super(callId string) (IllegalArgumentException error) {
	this.SIPHeader.super(core.SIPHeaderNames_CALL_ID)
	this.callIdentifier, IllegalArgumentException = NewCallIdentifier(callId)
	return IllegalArgumentException
}

/**
 * Compare two call ids for equality.
 * @param other Object to set
 * @return true if the two call ids are equals, false otherwise
 */
/*public boolean equals(Object other) {
    if (! this.getClass().equals(other.getClass())) {
        return false;
    }
    CallID that = (CallID) other;
    return this.callIdentifier.equals(that.callIdentifier);
}*/
func (this *CallID) String() string {
	return this.headerName + core.SIPSeparatorNames_COLON +
		core.SIPSeparatorNames_SP + this.EncodeBody() + core.SIPSeparatorNames_NEWLINE
}

/** Encode the body part of this header (i.e. leave out the hdrName).
*@return String encoded body part of the header.
 */
func (this *CallID) EncodeBody() string {
	if this.callIdentifier == nil {
		return ""
	} else {
		return this.callIdentifier.String()
	}
}

/** get the CallId field. This does the same thing as
	 * encodeBody
         * @return String the encoded body part of the
*/
func (this *CallID) GetCallId() string {
	return this.EncodeBody()
}

/**
 * get the call Identifer member.
 * @return CallIdentifier
 */
func (this *CallID) GetCallIdentifer() *CallIdentifier {
	return this.callIdentifier
}

/** set the CallId field
         * @param cid String to set. This is the body part of the Call-Id
	  *  header. It must have the form localId@host or localId.
         * @throws IllegalArgumentException if cid is null, not a token, or is
         * not a token@token.
*/
func (this *CallID) SetCallId(cid string) (ParseException error) {
	//try {
	this.callIdentifier, ParseException = NewCallIdentifier(cid)
	//} catch (IllegalArgumentException ex) {
	//throw new ParseException(cid,0);
	//}
	return ParseException
}

/**
 * Set the callIdentifier member.
 * @param cid CallIdentifier to set (localId@host).
 */
func (this *CallID) SetCallIdentifier(cid *CallIdentifier) {
	this.callIdentifier = cid
}
