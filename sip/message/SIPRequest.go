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

const DEFAULT_USER = "ip"
const DEFAULT_TTL = 1
const DEFAULT_TRANSPORT = "udp"
const DEFAULT_METHOD = INVITE

/**
* The SIP Request structure-- this belongs to the parser who fills it up.
*  Acknowledgements: Mark Bednarek made a few fixes to this code.
*   Jeff Keyser added two methods that create responses and generate
*   cancel requests from incoming orignial  requests without
*   the additional overhead  of encoding and decoding messages.
*
*@version  JAIN-SIP-1.1
*
*@author M. Ranganathan <mranga@nist.gov>  <br/>
*
*<a href="{@docRoot}/uncopyright.html">This code is in the public domain.</a>
*
 */
type SIPRequest struct {
	SIPMessage
	//implements javax.sip.message.Request  {

	//transactionPointer core.GenericObject

	requestLine *header.RequestLine
}

/** Constructor.
 */
func NewSIPRequest() *SIPRequest {
	this := &SIPRequest{}
	this.SIPMessage.super()
	return this
}

/**
* Replace a portion of this response with a new structure (given by
* newObj). This method finds a sub-structure that encodes to cText
* and has the same type as the second arguement and replaces this
* portion with the second argument.
* @param cText is the text that we want to replace.
* @param newObj is the new object that we want to put in place of
* 	cText.
 */

//        public void replace
// 	(String ctext, GenericObject newObject,
// 	boolean matchSubstring) {
// 	if (ctext == nil || newObject == nil)  {
// 		throw new IllegalArgumentException
// 			("Illegal argument nil");
// 	}

// 	requestLine.replace(ctext,newObject,matchSubstring);
// 	super.replace(ctext,newObject,matchSubstring);

// }

// /** Get the Request Line of the SIPRequest.
// *@return the request line of the SIP Request.
// */

func (this *SIPRequest) GetRequestLine() *header.RequestLine {
	return this.requestLine
}

// /** Set the request line of the SIP Request.
// *@param requestLine is the request line to Set in the SIP Request.
// */

func (this *SIPRequest) SetRequestLine(requestLine *header.RequestLine) {
	this.requestLine = requestLine
}

// /** Convert to a formatted string for pretty printing. Note that
// * the encode method converts this into a sip message that is suitable
// * for transmission. Note hack here if you want to convert the nice
// * curly brackets into some grotesque XML tag.
// *
// *@return a string which can be used to examine the message contents.
// *
// */
// public String debugDump() {
// 	String superstring =  super.debugDump();
// 	stringRepresentation = "";
// 	sprint(PackageNames.MESSAGE_PACKAGE+ ".SIPRequest");
// 	sprint("{");
// 	if(requestLine != nil) sprint(requestLine.debugDump());
// 	sprint(superstring);
// 	sprint("}");
// 	return stringRepresentation;
// }

// /**
// * Check header for constraints.
// * (1) Invite options and bye requests can only have SIP URIs in the
// * contact headers.
// * (2) Request must have cseq, to and from and via headers.
// * (3) Method in request URI must match that in CSEQ.
// */
func (this *SIPRequest) CheckHeaders() (ParseException error) {
	prefix := "Missing Header "

	/* Check for required headers */

	if this.GetCSeq() == nil {
		errors.New("ParseException:" + prefix + "CSeq")
	}
	if this.GetTo() == nil {
		errors.New("ParseException:" + prefix + "To")
	}
	if this.GetFrom() == nil {
		errors.New("ParseException:" + prefix + "From")
	}
	if this.GetViaHeaders() == nil {
		errors.New("ParseException:" + prefix + "Via")
	}
	if this.GetMaxForwards() == nil {
		errors.New("ParseException:" + prefix + "MaxForwards")
	}

	/*  BUGBUG
		* Need to revisit this check later...
	               * for now we just leave this to the
			* application to catch.
	*/

	if this.requestLine != nil && this.requestLine.GetMethod() != "" &&
		this.GetCSeq().GetMethod() != "" &&
		strings.ToLower(this.requestLine.GetMethod()) != strings.ToLower(this.GetCSeq().GetMethod()) {
		errors.New("ParseException: CSEQ method mismatch with  Request-Line ")
	}

	return nil
}

// /**
// * Set the default values in the request URI if necessary.
// */
func (this *SIPRequest) SetDefaults() {
	// The request line may be unparseable (Set to nil by the
	// exception handler.
	if this.requestLine == nil {
		return
	}
	method := this.requestLine.GetMethod()
	// The requestLine may be malformed!
	if method == "" {
		return
	}
	u := this.requestLine.GetUri()
	if u == nil {
		return
	}
	if method == REGISTER || method == INVITE {
		if sipUri, ok := u.(address.SipURI); ok {
			//SipURIImpl sipUri = (SipURIImpl)  u;
			sipUri.SetUserParam(DEFAULT_USER)
			// try {
			sipUri.SetTransportParam(DEFAULT_TRANSPORT)
			// } catch (ParseException ex) {}
		}
	}
}

// /**
// * Patch up the request line as necessary.
// */
func (this *SIPRequest) SetRequestLineDefaults() {
	method := this.requestLine.GetMethod()
	if method == "" {
		cseq := this.GetCSeq().(*header.CSeq)
		if cseq != nil {
			method = cseq.GetMethod()
			this.requestLine.SetMethod(method)
		}
	}
}

// /**
// * A conveniance function to access the Request URI.
// *@return the requestURI if it exists.
// */
func (this *SIPRequest) GetRequestURI() address.URI {
	if this.requestLine == nil {
		return nil
	} else {
		return this.requestLine.GetUri()
	}
}

//        /** Sets the RequestURI of Request. The Request-URI is a SIP or
//         * SIPS URI or a general URI. It indicates the user or service to which
//         * this request  is being addressed. SIP elements MAY support
//         * Request-URIs with schemes  other than "sip" and "sips", for
//         * example the "tel" URI scheme. SIP  elements MAY translate
//         * non-SIP URIs using any mechanism at their disposal,  resulting
//         * in SIP URI, SIPS URI, or some other scheme.
//         *
//         * @param requestURI - the new Request URI of this request message
//         */
func (this *SIPRequest) SetRequestURI(uri address.URI) {
	if this.requestLine == nil {
		this.requestLine = header.NewRequestLine()
	}
	this.requestLine.SetUri(uri)
}

// /** Set the method.
// *@param method is the method to Set.
// *@throws IllegalArgumentException if the method is nil
// */
func (this *SIPRequest) SetMethod(method string) {
	//if method == nil
	//  throw new IllegalArgumentException("nil method");
	if this.requestLine == nil {
		this.requestLine = header.NewRequestLine()
	}
	this.requestLine.SetMethod(method)
	if this.cSeqHeader != nil {
		//try{
		this.cSeqHeader.SetMethod(method)
		//}catch(ParseException e){}
	}
}

// /** Get the method from the request line.
// *@return the method from the request line if the method exits and
// * nil if the request line or the method does not exist.
// */
func (this *SIPRequest) GetMethod() string {
	if this.requestLine == nil {
		return ""
	} else {
		return this.requestLine.GetMethod()
	}
}

// /**
// *  Encode the SIP Request as a string.
// *
// *@return an encoded String containing the encoded SIP Message.
// */

func (this *SIPRequest) String() string {
	var retval string
	if this.requestLine != nil {
		this.SetRequestLineDefaults()
		retval = this.requestLine.String() + this.SIPMessage.String()
	} else {
		retval = this.SIPMessage.String()
	}
	return retval + core.SIPSeparatorNames_NEWLINE
}

// /** ALias for encode above.
// */
// public String toString() { return this.encode(); }

// /**
// * Make a clone (deep copy) of this object.
// * You can use this if you
// * want to modify a request while preserving the original
// *
// *@return a deep copy of this object.
// */

//        public Object clone() {

//            SIPRequest retval = (SIPRequest) super.clone();
//     if (this.requestLine != nil) {
//              retval.requestLine = (RequestLine) this.requestLine.clone();
//       retval.SetRequestLineDefaults();
//     }
//            return retval;
//        }

// /**
// * Compare for equality.
// *
// *@param other object to compare ourselves with.
// */
// public boolean equals(Object other) {
//     if ( ! this.GetClass().equals(other.GetClass())) return false;
//     SIPRequest that = (SIPRequest) other;

//     return  requestLine.equals(that.requestLine)
//      && super.equals(other);
// }

// /**
// * Get the message as a linked list of strings.
// * Use this if you want to iterate through the message.
// *
// *@return a linked list containing the request line and
// * headers encoded as strings.
// */
func (this *SIPRequest) GetMessageAsEncodedStrings() *list.List {
	retval := this.SIPMessage.GetMessageAsEncodedStrings()
	if this.requestLine != nil {
		this.SetRequestLineDefaults()
		retval.PushFront(this.requestLine.String())
	}
	return retval
}

// /**
// * Match with a template. You can use this if you want to match
// * incoming messages with a pattern and do something when you find
// * a match. This is useful for building filters/pattern matching
// * responders etc.
// *
// *@param matchObj object to match ourselves with (nil matches wildcard)
// *
// */
// public boolean match(Object matchObj) {
//    if (matchObj == nil) return true;
//    else if ( ! matchObj.GetClass().equals(this.GetClass()))
// 			return false;
//    else if (matchObj == this) return true;
//    SIPRequest that = (SIPRequest) matchObj;
//    RequestLine rline = that.requestLine;
//    if (this.requestLine == nil && rline != nil) return false;
//    else if (this.requestLine == rline) return super.match(matchObj);
//    return requestLine.match(that.requestLine) && super.match(matchObj);

// }

// /** Get a dialog identifier.
// * Generates a string that can be used as a dialog identifier.
//        *
// *@param isServer is Set to true if this is the UAS
// *	and Set to false if this is the UAC
// */
func (this *SIPRequest) GetDialogId(isServer bool) string {
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

// /** Get a dialog id given the remote tag.
// */
func (this *SIPRequest) GetDialogId2(isServer bool, toTag string) string {
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

// /** Encode this into a byte array.
// * This is used when the body has been Set as a binary array
// * and you want to encode the body as a byte array for transmission.
// *
// *@return a byte array containing the SIPRequest encoded as a byte
// *  array.
// */

func (this *SIPRequest) EncodeAsBytes() []byte {
	var rlbytes []byte
	if this.requestLine != nil {
		// try {
		rlbytes = []byte(this.requestLine.String()) //.GetBytes("UTF-8");
		//   } catch (UnsupportedEncodingException ex) {
		// InternalErrorHandler.handleException(ex);
		//   }
	}
	superbytes := this.SIPMessage.EncodeAsBytes()
	retval := make([]byte, len(rlbytes)+len(superbytes))
	i := 0
	if rlbytes != nil {
		for i = 0; i < len(rlbytes); i++ {
			retval[i] = rlbytes[i]
		}
	}

	for j := 0; j < len(superbytes); j++ {
		retval[i] = superbytes[j]
		i++
	}
	return retval
}

// /** Creates a default SIPResponse message for this request. Note
//        * You must add the necessary tags to outgoing responses if need
// * be. For efficiency, this method does not clone the incoming
// * request. If you want to modify the outgoing response, be sure
// * to clone the incoming request as the headers are shared and
// * any modification to the headers of the outgoing response will
// * result in a modification of the incoming request.
// * Tag fields are just copied from the incoming request.
// * Contact headers are removed from the incoming request.
// * Added by Jeff Keyser.
// *
// *@param statusCode Status code for the response.
// * Reason phrase is generated.
// *
// *@return A SIPResponse with the status and reason supplied, and a copy
// *of all the original headers from this request.
// */

func (this *SIPRequest) CreateResponse(statusCode int) *SIPResponse {
	var newResponse SIPResponse
	//Iterator	headerIterator;
	//SIPHeader	nextHeader;

	reasonPhrase := newResponse.GetReasonPhraseFromInt(statusCode)
	return this.CreateResponse2(statusCode, reasonPhrase)

}

// * Creates a default SIPResponse message for this request. Note
//        * You must add the necessary tags to outgoing responses if need
// * be. For efficiency, this method does not clone the incoming
// * request. If you want to modify the outgoing response, be sure
// * to clone the incoming request as the headers are shared and
// * any modification to the headers of the outgoing response will
// * result in a modification of the incoming request.
// * Tag fields are just copied from the incoming request.
// * Contact headers are removed from the incoming request.
// * Added by Jeff Keyser. Route headers are not added to the
// * response.
// *
// *@param statusCode Status code for the response.
// *@param reasonPhrase Reason phrase for this response.
// *
// *@return A SIPResponse with the status and reason supplied, and a copy
// *of all the original headers from this request.

func (this *SIPRequest) CreateResponse2(statusCode int, reasonPhrase string) *SIPResponse {
	//var	newResponse SIPResponse;
	//Iterator	headerIterator;
	//SIPHeader	nextHeader;

	newResponse := NewSIPResponse()
	//try {
	newResponse.SetStatusCode(statusCode)
	// } catch (ParseException ex) {
	//    throw new IllegalArgumentException("Bad code "+statusCode);
	// }
	if reasonPhrase != "" {
		newResponse.SetReasonPhrase(reasonPhrase)
	} else {
		newResponse.SetReasonPhrase(newResponse.GetReasonPhraseFromInt(statusCode))
	}

	for headerIterator := this.getHeaders().Front(); headerIterator != nil; headerIterator = headerIterator.Next() {
		nextHeader := headerIterator.Value.(header.Header)
		if _, ok := nextHeader.(*header.From); ok {
			newResponse.AttachHeader2(nextHeader, false)
		}
		if _, ok := nextHeader.(*header.To); ok {
			newResponse.AttachHeader2(nextHeader, false)
		}
		if _, ok := nextHeader.(*header.ViaList); ok {
			newResponse.AttachHeader2(nextHeader, false)
		}
		if _, ok := nextHeader.(*header.CallID); ok {
			newResponse.AttachHeader2(nextHeader, false)
		}
		if _, ok := nextHeader.(*header.RecordRouteList); ok {
			newResponse.AttachHeader2(nextHeader, false)
		}
		if _, ok := nextHeader.(*header.CSeq); ok {
			newResponse.AttachHeader2(nextHeader, false)
		}
		if _, ok := nextHeader.(*header.MaxForwards); ok {
			newResponse.AttachHeader2(nextHeader, false)
		}
		// nextHeader instanceof ExtensionHeader 	||
		if _, ok := nextHeader.(*header.TimeStamp); ok {
			/**
			 if (SIPMessage.isRequestHeader(nextHeader)) {
				continue;
			 } else if (nextHeader instanceof ContentLength) {
				// content length added when content is
				// added...
				continue;
			 } else if ( nextHeader instanceof ContactList)  {
			        // contacts are stripped from the response.
				continue;
			 }
			 **/
			//try {
			newResponse.AttachHeader2(nextHeader, false)
			// } catch(SIPDuplicateHeaderException e){
			// e.printStackTrace();
			// }
		}
	}
	return newResponse
}

// /** Creates a default SIPResquest message that would cancel
// * this request. Note that tag assignment and removal of
// * is left to the caller (we use whatever tags are present in the
// * original request).  Acknowledgement: Added by Jeff Keyser.
// * Incorporates a bug report from Andreas Byström.
// *
// *@return A CANCEL SIPRequest with a copy all the original headers
// * from this request except for Require, ProxyRequire.
// */

func (this *SIPRequest) CreateCancelRequest() *SIPRequest {
	//SIPRequest	newRequest;
	//Iterator	headerIterator;
	//SIPHeader	nextHeader;

	newRequest := NewSIPRequest()
	newRequest.SetRequestLine(this.requestLine)
	newRequest.SetMethod(CANCEL)

	for headerIterator := this.getHeaders().Front(); headerIterator != nil; headerIterator = headerIterator.Next() {
		nextHeader := headerIterator.Value.(header.Header)
		if _, ok := nextHeader.(*header.RequireList); ok {
			continue
		} else if _, ok := nextHeader.(*header.ProxyRequireList); ok {
			continue
		} else if _, ok := nextHeader.(*header.ContentLength); ok {
			continue
		} else if _, ok := nextHeader.(*header.ContentType); ok {
			continue
		} else if _, ok := nextHeader.(*header.ViaList); ok {
			/**
			   SIPHeader sipHeader =
				(SIPHeader)
			       ((ViaList) nextHeader).GetFirst().clone() ;
			   nextHeader = new ViaList();
			   ((ViaList)nextHeader).add(sipHeader);
			 **/
			//nextHeader = (ViaList) ((ViaList) nextHeader).clone();
		} else if cseq, ok := nextHeader.(*header.CSeq); ok { // CSeq method for a cancel request must be cancel.
			//CSeq cseq = (CSeq) nextHeader.clone();
			//try{
			cseq.SetMethod(CANCEL)
			//}
			//catch(ParseException e){}
			nextHeader = cseq
		}
		//try {
		newRequest.AttachHeader2(nextHeader, false)
		// } catch(SIPDuplicateHeaderException e){
		// 	e.printStackTrace();
		// }
	}
	return newRequest
}

// /** Creates a default ACK SIPRequest message for this original request.
// * Note that the defaultACK SIPRequest does not include the
// * content of the original SIPRequest. If responSetoHeader
// * is nil then the toHeader of this request is used to
// * construct the ACK.  Note that tag fields are just copied
// * from the original SIP Request.  Added by Jeff Keyser.
// *
// *@param responSetoHeader To header to use for this request.
// *
// *@return A SIPRequest with an ACK method.
// */
func (this *SIPRequest) CreateAckRequest(responseToHeader *header.To) *SIPRequest {
	// SIPRequest	newRequest;
	// Iterator	headerIterator;
	// SIPHeader	nextHeader;

	newRequest := NewSIPRequest()
	newRequest.SetRequestLine(this.requestLine)
	newRequest.SetMethod(ACK)
	for headerIterator := this.getHeaders().Front(); headerIterator != nil; headerIterator = headerIterator.Next() {
		nextHeader := headerIterator.Value.(header.Header)
		if _, ok := nextHeader.(*header.RouteList); ok {
			// Ack and cancel do not Get ROUTE headers.
			// Route header for ACK is assigned by the
			// Dialog if necessary.
			continue
		} else if _, ok := nextHeader.(*header.ProxyAuthorization); ok {
			// Remove proxy auth header.
			// Assigned by the Dialog if necessary.
			continue
		} else if cl, ok := nextHeader.(*header.ContentLength); ok {
			// Adding content is responsibility of user.
			//nextHeader = nextHeader.(*header.ContentLength)
			//try{
			cl.SetContentLength(0)
			nextHeader = cl
			//}
			//catch(InvalidArgumentException e){}
		} else if _, ok := nextHeader.(*header.ContentType); ok {
			// Content type header is removed since
			// content length is 0. Bug fix from
			// Antonis Kyardis.
			continue
		} else if cseq, ok := nextHeader.(*header.CSeq); ok {
			// The CSeq header field in the
			// ACK MUST contain the same value for the
			// sequence number as was present in the
			// original request, but the method parameter
			// MUST be equal to "ACK".
			//nextHeader = nextHeader.(*header.CSeq)
			//try{
			cseq.SetMethod(ACK)
			nextHeader = cseq
			//catch(ParseException e){}
			// = cseq;
		} else if _, ok := nextHeader.(*header.To); ok {
			if responseToHeader != nil {
				nextHeader = responseToHeader
			} else {
				//nextHeader = (SIPHeader) nextHeader.clone();
			}
		} else if vl, ok := nextHeader.(*header.ViaList); ok {
			// Bug reported by Gianluca Martinello
			//The ACK MUST contain a single Via header field,
			// and this MUST be equal to the top Via header
			// field of the original
			// request.

			nextHeader = vl.Front().Value.(header.Header)
		} else {
			//nextHeader = (SIPHeader) nextHeader.clone();
		}

		//try {
		newRequest.AttachHeader2(nextHeader, false)
		// } catch(SIPDuplicateHeaderException e){
		// 	e.printStackTrace();
		// }
	}
	return newRequest
}

// /** Create a new default SIPRequest from the original request. Warning:
// * the newly created SIPRequest, shares the headers of
// * this request but we generate any new headers that we need to modify
// * so  the original request is umodified. However, if you modify the
// * shared headers after this request is created, then the newly
// * created request will also be modified.
// * If you want to modify the original request
// * without affecting the returned Request
// * make sure you clone it before calling this method.
// *
// * Only required headers are copied.
// * <ul>
// * <li>
// * Contact headers are not included in the newly created request.
// * Setting the appropriate sequence number is the responsibility of
// * the caller. </li>
// * <li> RouteList is not copied for ACK and CANCEL </li>
// * <li> Note that we DO NOT copy the body of the
// * argument into the returned header. We do not copy the content
// * type header from the original request either. These have to be
// * added seperately and the content length has to be correctly Set
// * if necessary the content length is Set to 0 in the returned header.
// * </li>
// * <li>Contact List is not copied from the original request.</li>
// * <li>RecordRoute List is not included from original request. </li>
// * <li>Via header is not included from the original request. </li>
// * </ul>
// *
// *@param requestLine is the new request line.
// *
// *@param switchHeaders is a boolean flag that causes to and from
// * 	headers to switch (Set this to true if you are the
// *	server of the transaction and are generating a BYE
// *	request). If the headers are switched, we generate
// *	new From and To headers otherwise we just use the
// *	incoming headers.
// *
// *@return a new Default SIP Request which has the requestLine specified.
// *
// */
func (this *SIPRequest) CreateSIPRequest(requestLine *header.RequestLine, switchHeaders bool) *SIPRequest {
	newRequest := NewSIPRequest()
	newRequest.requestLine = this.requestLine
	for headerIterator := this.getHeaders().Front(); headerIterator != nil; headerIterator = headerIterator.Next() {
		nextHeader := headerIterator.Value.(header.Header)
		// For BYE and cancel Set the CSeq header to the
		// appropriate method.
		if newCseq, ok := nextHeader.(*header.CSeq); ok {
			// CSeq newCseq = (CSeq) nextHeader.clone();
			nextHeader = newCseq
			//try{
			newCseq.SetMethod(this.requestLine.GetMethod())
			//}
			//catch(ParseException e){}
		} else if vl, ok := nextHeader.(*header.ViaList); ok {
			via := vl.Front().Value.(*header.Via)
			via.RemoveParameter("branch")
			nextHeader = via
			// Cancel and ACK preserve the branch ID.
		} else if to, ok := nextHeader.(*header.To); ok {
			if switchHeaders {
				from := header.NewFrom()
				from.CloneTo(to)
				from.RemoveTag()
				nextHeader = from
			} else {
				//to2 := to.Clone()
				to.RemoveTag()
				nextHeader = to
			}
		} else if from, ok := nextHeader.(*header.From); ok {
			if switchHeaders {
				to := header.NewTo()
				to.CloneFrom(from)
				to.RemoveTag()
				nextHeader = to
			} else {
				//from2 := from.Clone()
				from.RemoveTag()
				nextHeader = from
			}
		} else if cl, ok := nextHeader.(*header.ContentLength); ok {
			//ContentLength cl  =
			//	(ContentLength) nextHeader.clone();
			//try{
			cl.SetContentLength(0)
			//}
			//catch(InvalidArgumentException e){}
			nextHeader = cl
		} else {
			_, ok1 := nextHeader.(*header.CallID)
			_, ok2 := nextHeader.(*header.MaxForwards)
			if !ok1 && !ok2 {
				// Route is kept by dialog.
				// RR is added by the caller.
				// Contact is added by the Caller
				// Any extension headers must be added
				// by the caller.
				continue
			}
		}
		//try {
		newRequest.AttachHeader2(nextHeader, false)
		//} catch(SIPDuplicateHeaderException e){
		//	e.printStackTrace();
		//}
	}
	return newRequest

}

/** Create a BYE request from this request.
 *
 *@param switchHeaders is a boolean flag that causes from and
 *	isServerTransaction to headers to be swapped. Set this
 *	to true if you are the server of the dialog and are generating
 *      a BYE request for the dialog.
 *@return a new default BYE request.
 */
func (this *SIPRequest) CreateBYERequest(switchHeaders bool) *SIPRequest {
	requestLine := this.requestLine //.clone();
	requestLine.SetMethod("BYE")
	return this.CreateSIPRequest(requestLine, switchHeaders)
}

/** Create an ACK request from this request. This is suitable for
* generating an ACK for an INVITE  client transaction.
*
*@return an ACK request that is generated from this request.
*
 */
func (this *SIPRequest) CreateACKRequest() *SIPRequest {
	requestLine := this.requestLine //.clone();
	requestLine.SetMethod(ACK)
	return this.CreateSIPRequest(requestLine, false)
}

/**
* Get the host from the topmost via header.
*
*@return the string representation of the host from the topmost via
* header.
 */
func (this *SIPRequest) GetViaHost() string {
	via := this.GetViaHeaders().Front().Value.(*header.Via)
	return via.GetHost()
}

/**
* Get the port from the topmost via header.
*
*@return the port from the topmost via header (5060 if there is
*  no port indicated).
 */
func (this *SIPRequest) GetViaPort() int {
	via := this.GetViaHeaders().Front().Value.(*header.Via)
	if via.HasPort() {
		return via.GetPort()
	} else {
		return 5060
	}
}

/**
* Get the first line encoded.
*
*@return a string containing the encoded request line.
 */
func (this *SIPRequest) GetFirstLine() string {
	if this.requestLine == nil {
		return ""
	} else {
		return this.requestLine.String()
	}
}

/** Set the sip version.
*
*@param sipVerison -- the sip version to Set.
 */

func (this *SIPRequest) SetSIPVersion(sipVersion string) {
	//throws ParseException {
	//if (sipVersion == nil || !sipVersion.equals("SIP/2.0"))
	//	throw new ParseException ("sipVersion" , 0);
	this.requestLine.SetSIPVersion(sipVersion)
}

/** Get the SIP version.
*
*@return the SIP version from the request line.
 */
func (this *SIPRequest) GetSIPVersion() string {
	return this.requestLine.GetSipVersion()
}

//func (this *SIPRequest) GetTransaction() core.GenericObject {
//	// Return an opaque pointer to the transaction object.
//	// This is for consistency checking and quick lookup.
//	return this.transactionPointer
//}

//func (this *SIPRequest) SetTransaction(transaction core.GenericObject) {
//	this.transactionPointer = transaction
//}
