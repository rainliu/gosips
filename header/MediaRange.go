package header

import (
	"gosip/core"
)

/**
*   Media Range
* @see Accept
* @since 0.9
* @version 1.0
* <pre>
* Revisions:
*
* Version 1.0
*    1. Added encode method.
*
* media-range    = ( "STAR/STAR"
*                        | ( type "/" STAR )
*                        | ( type "/" subtype )
*                        ) *( ";" parameter )
*
* HTTP RFC 2616 Section 14.1
* </pre>
 */
type MediaRange struct {

	/** type field
	 */
	mtype string

	/** subtype field
	 */
	subtype string
}

/** Default constructor
 */
func NewMediaRange() *MediaRange {
	this := &MediaRange{}
	return this
}

/** get type field
 * @return String
 */
func (this *MediaRange) GetType() string {
	return this.mtype
}

/** get the subType field.
 * @return String
 */
func (this *MediaRange) GetSubtype() string {
	return this.subtype
}

/**
 * Set the type member
 * @param t String to set
 */
func (this *MediaRange) SetType(t string) {
	this.mtype = t
}

/**
 * Set the subtype member
 * @param s String to set
 */
func (this *MediaRange) SetSubtype(s string) {
	this.subtype = s
}

/**
 * Encode the object.
 * @return String
 */
func (this *MediaRange) String() string {
	encoding := this.mtype + core.SIPSeparatorNames_SLASH + this.subtype
	return encoding
}
