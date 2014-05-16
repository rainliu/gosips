package message

import (
	"bytes"
	"container/list"
	"errors"
	"gosips/core"
	"gosips/sip/address"
	"gosips/sip/header"
	"strings"
)

/**
 * SIP Response structure.
 *
 *@version  JAIN-SIP-1.1
 *
 *@author M. Ranganathan <mranga@nist.gov>  <br/>
 *
 *<a href="{@docRoot}/uncopyright.html">This code is in the public domain.</a>
 *
 */

type SIPResponse struct {
	SIPMessage
	//implements javax.sip.message.Response {
	statusLine *header.StatusLine
}

/** Constructor.
 */
func NewSIPResponse() *SIPResponse {
	this := &SIPResponse{}
	this.SIPMessage.super()
	return this
}

func (this *SIPResponse) GetReasonPhraseFromInt(rc int) string {
	var retval string
	switch rc {
	case TRYING:
		retval = "Trying"

	case RINGING:
		retval = "Ringing"

	case CALL_IS_BEING_FORWARDED:
		retval = "Call is being forwarded"

	case QUEUED:
		retval = "Queued"

	case SESSION_PROGRESS:
		retval = "Session progress"

	case OK:
		retval = "OK"

	case ACCEPTED:
		retval = "Accepted"

	case MULTIPLE_CHOICES:
		retval = "Multiple choices"

	case MOVED_PERMANENTLY:
		retval = "Moved permanently"

	case MOVED_TEMPORARILY:
		retval = "Moved Temporarily"

	case USE_PROXY:
		retval = "Use proxy"

	case ALTERNATIVE_SERVICE:
		retval = "Alternative service"

	case BAD_REQUEST:
		retval = "Bad request"

	case UNAUTHORIZED:
		retval = "Unauthorized"

	case PAYMENT_REQUIRED:
		retval = "Payment required"

	case FORBIDDEN:
		retval = "Forbidden"

	case NOT_FOUND:
		retval = "Not found"

	case METHOD_NOT_ALLOWED:
		retval = "Method not allowed"

	case NOT_ACCEPTABLE:
		retval = "Not acceptable"

	case PROXY_AUTHENTICATION_REQUIRED:
		retval = "Proxy Authentication required"

	case REQUEST_TIMEOUT:
		retval = "Request timeout"

	case GONE:
		retval = "Gone"

	case TEMPORARILY_UNAVAILABLE:
		retval = "Temporarily Unavailable"

	case REQUEST_ENTITY_TOO_LARGE:
		retval = "Request entity too large"

	case REQUEST_URI_TOO_LONG:
		retval = "Request-URI too large"

	case UNSUPPORTED_MEDIA_TYPE:
		retval = "Unsupported media type"

	case UNSUPPORTED_URI_SCHEME:
		retval = "Unsupported URI Scheme"

	case BAD_EXTENSION:
		retval = "Bad extension"

	case EXTENSION_REQUIRED:
		retval = "Etension Required"

	case INTERVAL_TOO_BRIEF:
		retval = "Interval too brief"

	case CALL_OR_TRANSACTION_DOES_NOT_EXIST:
		retval = "Call leg/Transaction does not exist"

	case LOOP_DETECTED:
		retval = "Loop detected"

	case TOO_MANY_HOPS:
		retval = "Too many hops"

	case ADDRESS_INCOMPLETE:
		retval = "Address incomplete"

	case AMBIGUOUS:
		retval = "Ambiguous"

	case BUSY_HERE:
		retval = "Busy here"

	case REQUEST_TERMINATED:
		retval = "Request Terminated"

	case NOT_ACCEPTABLE_HERE:
		retval = "Not Accpetable here"

	case BAD_EVENT:
		retval = "Bad Event"

	case REQUEST_PENDING:
		retval = "Request Pending"

	case SERVER_INTERNAL_ERROR:
		retval = "Server Internal Error"

	case UNDECIPHERABLE:
		retval = "Undecipherable"

	case NOT_IMPLEMENTED:
		retval = "Not implemented"

	case BAD_GATEWAY:
		retval = "Bad gateway"

	case SERVICE_UNAVAILABLE:
		retval = "Service unavailable"

	case SERVER_TIMEOUT:
		retval = "Gateway timeout"

	case VERSION_NOT_SUPPORTED:
		retval = "SIP version not supported"

	case MESSAGE_TOO_LARGE:
		retval = "Message Too Large"

	case BUSY_EVERYWHERE:
		retval = "Busy everywhere"

	case DECLINE:
		retval = "Decline"

	case DOES_NOT_EXIST_ANYWHERE:
		retval = "Does not exist anywhere"

	case SESSION_NOT_ACCEPTABLE:
		retval = "Session Not acceptable"

	default:
		retval = ""

	}
	return retval

}

//    /** Set the status code.
//     *@param statusCode is the status code to Set.
//     *@throws IlegalArgumentException if invalid status code.
//     */
func (this *SIPResponse) SetStatusCode(statusCode int) { //throws ParseException {
	// if (statusCode < 100 || statusCode > 800)
	//     throw new ParseException("bad status code",0);
	if this.statusLine == nil {
		this.statusLine = header.NewStatusLine()
	}
	this.statusLine.SetStatusCode(statusCode)
}

//    /**
//     * Get the status line of the response.
//     *@return StatusLine
//     */
func (this *SIPResponse) GetStatusLine() *header.StatusLine {
	return this.statusLine
}

//    /** Get the staus code (conveniance function).
//     *@return the status code of the status line.
//     */
func (this *SIPResponse) GetStatusCode() int {
	return this.statusLine.GetStatusCode()
}

//    /** Set the reason phrase.
//     *@param reasonPhrase the reason phrase.
//     *@throws IllegalArgumentException if nil string
//     */
func (this *SIPResponse) SetReasonPhrase(reasonPhrase string) {
	//if this.reasonPhrase == nil)
	//    throw new IllegalArgumentException("Bad reason phrase");
	if this.statusLine == nil {
		this.statusLine = header.NewStatusLine()
	}
	this.statusLine.SetReasonPhrase(reasonPhrase)
}

//    /** Get the reason phrase.
//     *@return the reason phrase.
//     */
func (this *SIPResponse) GetReasonPhrase() string {
	if this.statusLine == nil || this.statusLine.GetReasonPhrase() == "" {
		return ""
	} else {
		return this.statusLine.GetReasonPhrase()
	}
}

//    /** Return true if the response is a final response.
//     *@param rc is the return code.
//     *@return true if the parameter is between the range 200 and 700.
//     */
func (this *SIPResponse) IsFinalResponseFromInt(rc int) bool {
	return rc >= 200 && rc < 700
}

//    /** Is this a final response?
//     *@return true if this is a final response.
//     */
func (this *SIPResponse) IsFinalResponse() bool {
	return this.IsFinalResponseFromInt(this.statusLine.GetStatusCode())
}

//    /**
//     * Set the status line field.
//     *@param sl Status line to Set.
//     */
func (this *SIPResponse) SetStatusLine(sl *header.StatusLine) {
	this.statusLine = sl
}

//    /**
//     * Print formatting function.
//     *Indent and parenthesize for pretty printing.
//     * Note -- use the encode method for formatting the message.
//     * Hack here to XMLize.
//     *
//     *@return a string for pretty printing.
//     */
//    public String debugDump() {
//        String superstring =  super.debugDump();
//        stringRepresentation = "";
//        sprint(MESSAGE_PACKAGE + ".SIPResponse");
//        sprint("{");
//        if (statusLine != nil) {
//            sprint(statusLine.debugDump());
//        }
//        sprint(superstring);
//        sprint("}");
//        return stringRepresentation;
//    }

//    /**
//     * Check the response structure. Must have from, to CSEQ and VIA
//     * headers.
//     */
func (this *SIPResponse) CheckHeaders() (ParseException error) {
	if this.GetCSeq() == nil {
		return errors.New("ParseException: CSeq")
	}
	if this.GetTo() == nil {
		return errors.New("ParseException: To")
	}
	if this.GetFrom() == nil {
		return errors.New("ParseException: From")
	}
	if this.GetViaHeaders() == nil {
		return errors.New("ParseException: Via")
	}
	return nil
}

//    /**
//     *  Encode the SIP Request as a string.
//     *@return The string encoded canonical form of the message.
//     */

func (this *SIPResponse) String() string {
	var retval string
	if this.statusLine != nil {
		retval = this.statusLine.String() + this.SIPMessage.String()
	} else {
		retval = this.SIPMessage.String()
	}
	return retval + core.SIPSeparatorNames_NEWLINE
}

// func (this *SIPResponse) String() string {
// 	return this.statusLine.String() + this.SIPMessage.String()
// }

//    /** Get this message as a list of encoded strings.
//     *@return LinkedList containing encoded strings for each header in
//     *   the message.
//     */

func (this *SIPResponse) GetMessageAsEncodedStrings() *list.List {
	retval := this.SIPMessage.GetMessageAsEncodedStrings()

	if this.statusLine != nil {
		retval.PushFront(this.statusLine.String())
	}
	return retval

}

//    /**
//     * Make a clone (deep copy) of this object.
//     *@return a deep copy of this object.
//     */

//    public Object clone() {
//        SIPResponse retval = (SIPResponse) super.clone();
//        retval.statusLine = (StatusLine) this.statusLine.clone();
//        return retval;
//    }
//    /**
//     * Replace a portion of this response with a new structure (given by
//     * newObj). This method finds a sub-structure that encodes to cText
//     * and has the same type as the second arguement and replaces this
//     * portion with the second argument.
//     * @param cText is the text that we want to replace.
//     * @param newObj is the new object that we want to put in place of
//     * 	cText.
//     * @param matchSubstring boolean to indicate whether to match on
//     *   substrings when searching for a replacement.
//     */
//    func (this *SIPResponse) replace(String cText, GenericObject newObj,
//    boolean matchSubstring ) {
//        if (cText == nil || newObj == nil)
//            throw new
//            IllegalArgumentException("nil args!");
//        if (newObj instanceof SIPHeader)
//            throw new
//            IllegalArgumentException("Bad replacement class " +
//            newObj.GetClass().GetName());

//        if (statusLine != nil)
//            statusLine.replace(cText,newObj,matchSubstring);
//        super.replace(cText,newObj,matchSubstring);
//    }

//    /**
//     * Compare for equality.
//     *@param other other object to compare with.
//     */
//    public boolean equals(Object other) {
//        if ( ! this.GetClass().equals(other.GetClass())) return false;
//        SIPResponse that = (SIPResponse) other;
//        return statusLine.equals(that.statusLine) &&
//        super.equals(other);
//    }

//    /**
//     * Match with a template.
//     *@param matchObj template object to match ourselves with (nil
//     * in any position in the template object matches wildcard)
//     */
//    public boolean match(Object matchObj) {
//        if (matchObj == nil) return true;
//        else if ( ! matchObj.GetClass().equals(this.GetClass()))  {
//            return false;
//        } else if (matchObj == this) return true;
//        SIPResponse that = (SIPResponse) matchObj;
//        // System.out.println("---------------------------------------");
//        // System.out.println("matching " + this.encode());
//        // System.out.println("matchObj " + that.encode());
//        StatusLine rline = that.statusLine;
//        if (this.statusLine == nil && rline != nil) return false;
//        else if (this.statusLine == rline) return super.match(matchObj);
//        else {
//            // System.out.println(statusLine.match(that.statusLine));
//            // System.out.println(super.match(matchObj));
//            // System.out.println("---------------------------------------");
//            return statusLine.match(that.statusLine) &&
//            super.match(matchObj);
//        }

//    }

//    /** Encode this into a byte array.
//     * This is used when the body has been Set as a binary array
//     * and you want to encode the body as a byte array for transmission.
//     *
//     *@return a byte array containing the SIPRequest encoded as a byte
//     *  array.
//     */

func (this *SIPResponse) EncodeAsBytes() []byte {
	var slbytes []byte
	if this.statusLine != nil {
		//try {
		slbytes = []byte(this.statusLine.String()) //.GetBytes("UTF-8");
		// } catch (UnsupportedEncodingException ex){
		//     InternalErrorHandler.handleException(ex);
		// }
	}
	superbytes := this.SIPMessage.EncodeAsBytes()
	retval := make([]byte, len(slbytes)+len(superbytes))
	i := 0
	if slbytes != nil {
		for i = 0; i < len(slbytes); i++ {
			retval[i] = slbytes[i]
		}
	}

	for j := 0; j < len(superbytes); j++ {
		retval[i] = superbytes[j]
		i++
	}
	return retval
}

/** Get the dialog identifier. Assume the incoming response
   * corresponds to a client dialog for an outgoing request.
   * Acknowledgement -- this was contributed by Lamine Brahimi.
   *
   *@return a string that can be used to identify the dialog.
  public String GetDialogId()  {
      CallID cid = (CallID)this.GetCallId();
      From from = (From) this.GetFrom();
      String retval = cid.GetCallId();
      retval += COLON + from.GetUserAtHostPort();
      retval += COLON;
      if (from.GetTag() != nil)
          retval +=  from.GetTag();

      return retval.toLowerCase();
  }
*/

//    /** Get a dialog identifier.
//     * Generates a string that can be used as a dialog identifier.
//     *
//     * @param isServer is Set to true if this is the UAS
//     * and Set to false if this is the UAC
//     */
func (this *SIPResponse) GetDialogId(isServer bool) string {
	cid := this.GetCallId()
	from := this.GetFrom().(*header.From)
	to := this.GetTo().(*header.To)
	var retval bytes.Buffer
	retval.WriteString(cid.GetCallId())
	if !isServer {
		retval.WriteString(core.SIPSeparatorNames_COLON)
		retval.WriteString(from.GetUserAtHostPort())
		if from.GetTag() != "" {
			retval.WriteString(core.SIPSeparatorNames_COLON)
			retval.WriteString(from.GetTag())
		}
		retval.WriteString(core.SIPSeparatorNames_COLON)
		retval.WriteString(to.GetUserAtHostPort())
		if to.GetTag() != "" {
			retval.WriteString(core.SIPSeparatorNames_COLON)
			retval.WriteString(to.GetTag())
		}
	} else {
		retval.WriteString(core.SIPSeparatorNames_COLON)
		retval.WriteString(to.GetUserAtHostPort())
		if to.GetTag() != "" {
			retval.WriteString(core.SIPSeparatorNames_COLON)
			retval.WriteString(to.GetTag())
		}
		retval.WriteString(core.SIPSeparatorNames_COLON)
		retval.WriteString(from.GetUserAtHostPort())
		if from.GetTag() != "" {
			retval.WriteString(core.SIPSeparatorNames_COLON)
			retval.WriteString(from.GetTag())
		}
	}
	return strings.ToLower(retval.String())
}

func (this *SIPResponse) GetDialogId2(isServer bool, toTag string) string {
	cid := this.GetCallId()
	from := this.GetFrom().(*header.From)
	to := this.GetTo().(*header.To)
	var retval bytes.Buffer
	retval.WriteString(cid.GetCallId())
	if !isServer {
		retval.WriteString(core.SIPSeparatorNames_COLON)
		retval.WriteString(from.GetUserAtHostPort())
		if from.GetTag() != "" {
			retval.WriteString(core.SIPSeparatorNames_COLON)
			retval.WriteString(from.GetTag())
		}
		retval.WriteString(core.SIPSeparatorNames_COLON)
		retval.WriteString(to.GetUserAtHostPort())
		if toTag != "" {
			retval.WriteString(core.SIPSeparatorNames_COLON)
			retval.WriteString(toTag)
		}
	} else {
		retval.WriteString(core.SIPSeparatorNames_COLON)
		retval.WriteString(to.GetUserAtHostPort())
		if toTag != "" {
			retval.WriteString(core.SIPSeparatorNames_COLON)
			retval.WriteString(toTag)
		}
		retval.WriteString(core.SIPSeparatorNames_COLON)
		retval.WriteString(from.GetUserAtHostPort())
		if from.GetTag() != "" {
			retval.WriteString(core.SIPSeparatorNames_COLON)
			retval.WriteString(from.GetTag())
		}
	}
	return strings.ToLower(retval.String())
}

//    /**
//     * Create a new SIPRequest from the given response. Note that the
//     * RecordRoute Via and CSeq headers are not copied from the response.
//     * These have to be added by the caller.
//     * This method is useful for generating ACK messages from final
//     * responses.
//     *
//     *@param requestURI is the request URI to use.
//     *@param via is the via header to use.
//     *@param cseq is the cseq header to use in the generated
//     * request.
//     */

func (this *SIPResponse) CreateRequest(requestURI address.SipURI, via *header.Via, cseq *header.CSeq) *SIPRequest {
	newRequest := NewSIPRequest()
	method := cseq.GetMethod()
	newRequest.SetMethod(method)
	newRequest.SetRequestURI(requestURI)
	if (method == "ACK" || method == "CANCEL") && this.GetTopmostVia().GetBranch() != "" {
		// Use the branch id from the OK.
		//try {
		via.SetBranch(this.GetTopmostVia().GetBranch())
		//} catch (ParseException ex) {}
	}
	newRequest.SetHeader(via)
	newRequest.SetHeader(cseq)

	for headerIterator := this.getHeaders().Front(); headerIterator != nil; headerIterator = headerIterator.Next() {
		nextHeader := headerIterator.Value.(header.Header)
		// Some headers do not belong in a Request ....
		if this.IsResponseHeader(nextHeader) {
			continue
		}
		if _, ok := nextHeader.(*header.ViaList); ok {
			continue
		}
		if _, ok := nextHeader.(*header.CSeq); ok {
			continue
		}
		if _, ok := nextHeader.(*header.ContentType); ok {
			continue
		}
		if _, ok := nextHeader.(*header.RecordRouteList); ok {
			continue
		}
		// if _, ok:=nextHeader.(*header.To); ok{
		//     nextHeader = nextHeader.clone();
		// }
		// else if (nextHeader instanceof From)
		//     nextHeader = (SIPHeader)nextHeader.clone();
		// try {
		newRequest.AttachHeader2(nextHeader, false)
		// } catch(SIPDuplicateHeaderException e){
		//     e.printStackTrace();
		// }
	}
	return newRequest
}

//    /**
//     * Get the encoded first line.
//     *
//     *@return the status line encoded.
//     *
//     */
func (this *SIPResponse) GetFirstLine() string {
	if this.statusLine == nil {
		return ""
	} else {
		return this.statusLine.String()
	}
}

func (this *SIPResponse) SetSIPVersion(sipVersion string) {
	this.statusLine.SetSipVersion(sipVersion)
}

func (this *SIPResponse) GetSIPVersion() string {
	return this.statusLine.GetSipVersion()
}
