package header

import (
	"bytes"
	"gosips/core"
	"strconv"
)

/**
* Status Line (for SIPReply) messages.
 */

type StatusLine struct {
	matchStatusClass bool

	/** sipVersion field
	 */
	sipVersion string

	/** status code field
	 */
	statusCode int

	/** reasonPhrase field
	 */
	reasonPhrase string
}

/** Default Constructor
 */
func NewStatusLine() *StatusLine {
	this := &StatusLine{}
	this.sipVersion = SIPConstants_SIP_VERSION_STRING
	return this
}

/** Match with a template.
 * Match only the response class if the last two digits of the
 * match templates are 0's
 */

//    public boolean match( Object matchObj) {
//        if (! (matchObj instanceof StatusLine)) return false;
//        StatusLine sl = (StatusLine) matchObj;
// // A pattern matcher has been registered.
// if ( sl.matchExpression != null )
// 	return sl.matchExpression.match(this.encode());
// // no patter matcher has been registered..
//        if (sl.sipVersion != null && ! sl.sipVersion.equals(sipVersion))
//            return false;
//        if ( sl.statusCode != 0) {
//            if (matchStatusClass) {
//                int hiscode = sl.statusCode;
//                String codeString = new Integer(sl.statusCode).toString();
//                String mycode = new Integer(statusCode).toString();
//                if (codeString.charAt(0) != mycode.charAt(0)) return false;
//            } else {
//                if (statusCode != sl.statusCode) return false;
//            }
//        }
//        if (sl.reasonPhrase == null ||
//        reasonPhrase == sl.reasonPhrase) return true;
//        return reasonPhrase.equals(sl.reasonPhrase);

//    }

/** Set the flag on a match template.
 *If this Set to true, then the whole status code is matched (default
 * behavior) else only the class of the response is matched.
 */
func (this *StatusLine) SetMatchStatusClass(flag bool) {
	this.matchStatusClass = flag
}

/**
 * Encode into a canonical form.
 * @return String
 */
func (this *StatusLine) String() string {
	var encoding bytes.Buffer
	encoding.WriteString(SIPConstants_SIP_VERSION_STRING + core.SIPSeparatorNames_SP + strconv.Itoa(this.statusCode) + core.SIPSeparatorNames_SP)
	if this.reasonPhrase != "" {
		encoding.WriteString(this.reasonPhrase)
	}
	encoding.WriteString(core.SIPSeparatorNames_NEWLINE)
	return encoding.String()
}

/** Get the Sip Version
 * @return SipVersion
 */
func (this *StatusLine) GetSipVersion() string {
	return this.sipVersion
}

/** Get the Status Code
 * @return StatusCode
 */
func (this *StatusLine) GetStatusCode() int {
	return this.statusCode
}

/** Get the ReasonPhrase field
 * @return  ReasonPhrase field
 */
func (this *StatusLine) GetReasonPhrase() string {
	return this.reasonPhrase
}

/**
 * Set the sipVersion member
 * @param s String to Set
 */
func (this *StatusLine) SetSipVersion(s string) {
	this.sipVersion = s
}

/**
 * Set the statusCode member
 * @param statusCode int to Set
 */
func (this *StatusLine) SetStatusCode(statusCode int) {
	this.statusCode = statusCode
}

/**
 * Set the reasonPhrase member
 * @param reasonPhrase String to Set
 */
func (this *StatusLine) SetReasonPhrase(reasonPhrase string) {
	this.reasonPhrase = reasonPhrase
}

/**
 * Get the major version number.
 *@return String major version number
 */
func (this *StatusLine) GetVersionMajor() string {
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
 *@return String minor version number
 */
func (this *StatusLine) GetVersionMinor() string {
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
