package header

import (
	"errors"
	"gosips/core"
	"strings"
)

/**
* The call identifer that goes into a callID header and a in-reply-to header.
* @see CallID
* @see InReplyTo
 */
type CallIdentifier struct {
	/** localId field
	 */
	localId string

	/** host field
	 */
	host string
}

/** Constructor
 *@param local id is the local id.
 *@param host is the host.
 */
/*func NewCallIdentifier( localId,  host string) *CallIdentifier{
    this := &CallIdentifier{};
    this.localId = localId;
    this.host = host;
    return this;
}*/

/** constructor
 * @param cid String to set
 * @throws IllegalArgumentException if cid is null or is not a token,
 * or token@token
 */
func NewCallIdentifier(cid string) (*CallIdentifier, error) {
	this := &CallIdentifier{}
	IllegalArgumentException := this.SetCallID(cid)
	if IllegalArgumentException != nil {
		return nil, IllegalArgumentException
	} else {
		return this, nil
	}
}

/**
 * Get the encoded version of this id.
 * @return String to set
 */
func (this *CallIdentifier) String() string {
	if this.host != "" {
		return this.localId + core.SIPSeparatorNames_AT + this.host
	} else {
		return this.localId
	}
}

/**
 * Compare two call identifiers for equality.
 * @param other Object to set
 * @return true if the two call identifiers are equals, false
 * otherwise
 */
/*public boolean equals( Object other) {
    if (! other.getClass().equals(this.getClass())) {
        return false;
    }
    CallIdentifier that = (CallIdentifier) other;
    if (this.localId.compareTo(that.localId) != 0) {
        return false;
    }
    if (this.host == that.host) return true;
    if ( (this.host == null && that.host != null) ||
         (this.host != null && that.host == null) ) return false;
    if (host.compareToIgnoreCase(that.host) != 0 ) {
        return false;
    }
    return true;
}*/

/** get the LocalId field
 * @return String
 */
func (this *CallIdentifier) GetLocalId() string {
	return this.localId
}

/** get the host field
 * @return host member String
 */
func (this *CallIdentifier) GetHost() string {
	return this.host
}

/**
 * Set the localId member
 * @param localId String to set
 */
func (this *CallIdentifier) SetLocalId(localId string) {
	this.localId = localId
}

/** set the callId field
 * @param cid Strimg to set
 * @throws IllegalArgumentException if cid is null or is not a token or
 * token@token
 */
func (this *CallIdentifier) SetCallID(cid string) (IllegalArgumentException error) {
	if cid == "" {
		return errors.New("IllegalArgumentException: NULL!")
	}
	index := strings.Index(cid, "@")
	if index == -1 {
		this.localId = cid
		this.host = ""
	} else {
		this.localId = cid[0:index]
		this.host = cid[index+1 : len(cid)]
		if this.localId == "" || this.host == "" {
			return errors.New("IllegalArgumentException: CallID  must be token@token or token")
		}
	}

	return nil
}

/**
 * Set the host member
 * @param host String to set
 */
func (this *CallIdentifier) SetHost(host string) {
	this.host = host
}
