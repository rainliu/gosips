package header

import (
	"bytes"
	"gosips/core"
	"gosips/sip/address"
)

/**
* RequestLine of SIP Request.
 */
type RequestLine struct {
	/** uri field. Note that this can be a SIP URI or a generic URI
	 * like tel URI.
	 */
	uri address.URI

	/** method field.
	 */
	method string

	/** sipVersion field
	 */
	sipVersion string
}

/** Default constructor
 */
func NewRequestLine() *RequestLine {
	this := &RequestLine{}
	this.sipVersion = "SIP/2.0"
	return this
}

/** Constructor given the request URI and the method.
 */
func NewRequestLineFromString(requestURI address.URI, method string) *RequestLine {
	this := &RequestLine{}
	this.uri = requestURI
	this.method = method
	this.sipVersion = "SIP/2.0"
	return this
}

/** Set the SIP version.
*@param sipVersion -- the SIP version to Set.
 */
func (this *RequestLine) SetSIPVersion(sipVersion string) {
	this.sipVersion = sipVersion
}

/** Encode the request line as a String.
	 *
         * @return requestLine encoded as a string.
*/
func (this *RequestLine) String() string {
	var encoding bytes.Buffer

	if this.method != "" {
		encoding.WriteString(this.method)
		encoding.WriteString(core.SIPSeparatorNames_SP)
	}
	if this.uri != nil {
		encoding.WriteString(this.uri.String())
		encoding.WriteString(core.SIPSeparatorNames_SP)
	}
	encoding.WriteString(this.sipVersion + core.SIPSeparatorNames_NEWLINE)
	return encoding.String()
}

/** Get the Request-URI.
	 *
         * @return the request URI
*/
func (this *RequestLine) GetUri() address.URI {
	return this.uri
}

/**
         * Get the Method
	 *
         * @return method string.
*/
func (this *RequestLine) GetMethod() string {
	return this.method
}

/**
         * Get the SIP version.
	 *
         * @return String
*/
func (this *RequestLine) GetSipVersion() string {
	return this.sipVersion
}

/**
 * Set the uri member.
 * @param uri URI to Set.
 */
func (this *RequestLine) SetUri(uri address.URI) {
	this.uri = uri
}

/**
         * Set the method member
	 *
         * @param method String to Set
*/
func (this *RequestLine) SetMethod(method string) {
	this.method = method
}

/**
         * Set the sipVersion member
	 *
         * @param s String to Set
*/
func (this *RequestLine) SetSipVersion(s string) {
	this.sipVersion = s
}

/**
* Get the major verrsion number.
*
*@return String major version number
 */
func (this *RequestLine) GetVersionMajor() string {
	if this.sipVersion == "" {
		return ""
	}
	var major string
	slash := false
	for i := 0; i < len(this.sipVersion); i++ {
		if this.sipVersion[i] == '.' {
			break
		}
		if slash {
			if major == "" {
				major = "" + string(this.sipVersion[i])
			} else {
				major += string(this.sipVersion[i])
			}
		}
		if this.sipVersion[i] == '/' {
			slash = true
		}
	}
	return major
}

/**
         * Get the minor version number.
	 *
	 *@return String minor version number
	 *
*/
func (this *RequestLine) GetVersionMinor() string {
	if this.sipVersion == "" {
		return ""
	}
	var minor string
	dot := false
	for i := 0; i < len(this.sipVersion); i++ {
		if dot {
			if minor == "" {
				minor = "" + string(this.sipVersion[i])
			} else {
				minor += string(this.sipVersion[i])
			}
		}
		if this.sipVersion[i] == '.' {
			dot = true
		}
	}
	return minor
}

/**
* Compare for equality.
*
*@param other object to compare with. We assume that all fields
* are Set.
 */
// public boolean equals(Object other)  {
//     boolean retval;
//     if ( ! other.GetClass().equals(this.GetClass()) ) {
// 		return false;
//     }
//     RequestLine that = (RequestLine) other;
//     try {
//        retval =  this.method.equals(that.method)
// 	&& this.uri.equals(that.uri)
// 	&& this.sipVersion.equals(that.sipVersion);
//     } catch (NullPointerException ex) {
// 	retval =  false;
//     }
//     return retval;

// }
