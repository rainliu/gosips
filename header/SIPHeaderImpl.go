package header

import (
	"strings"
	"gosip/core"
)

const SIPHeaderNames_MIN_EXPIRES = "Min-Expires";  //1
const SIPHeaderNames_ERROR_INFO = "Error-Info";  //2
const SIPHeaderNames_MIME_VERSION = "MIME-Version"; //3
const SIPHeaderNames_IN_REPLY_TO= "In-Reply-To"; //4
const SIPHeaderNames_ALLOW= "Allow"; //5
const SIPHeaderNames_CONTENT_LANGUAGE= "Content-Language"; //6
const SIPHeaderNames_CALL_INFO= "Call-Info"; //7
const SIPHeaderNames_CSEQ="CSeq"; //8
const SIPHeaderNames_ALERT_INFO="Alert-Info"; //9
const SIPHeaderNames_ACCEPT_ENCODING="Accept-Encoding"; //10
const SIPHeaderNames_ACCEPT= "Accept"; //11
const SIPHeaderNames_ACCEPT_LANGUAGE= "Accept-Language"; //12
const SIPHeaderNames_RECORD_ROUTE= "Record-Route"; //13
const SIPHeaderNames_TIMESTAMP="Timestamp"; //14
const SIPHeaderNames_TO="To"; //15
const SIPHeaderNames_VIA="Via"; //16
const SIPHeaderNames_FROM="From"; //17
const SIPHeaderNames_CALL_ID="Call-ID"; //18
const SIPHeaderNames_AUTHORIZATION="Authorization"; //19
const SIPHeaderNames_PROXY_AUTHENTICATE="Proxy-Authenticate"; //20
const SIPHeaderNames_SERVER="Server"; //21
const SIPHeaderNames_UNSUPPORTED="Unsupported"; //22
const SIPHeaderNames_RETRY_AFTER="Retry-After"; //23
const SIPHeaderNames_CONTENT_TYPE="Content-Type"; //24
const SIPHeaderNames_CONTENT_ENCODING="Content-Encoding"; //25
const SIPHeaderNames_CONTENT_LENGTH="Content-Length"; //26
const SIPHeaderNames_ROUTE="Route"; //27
const SIPHeaderNames_CONTACT="Contact"; //28
const SIPHeaderNames_WWW_AUTHENTICATE="WWW-Authenticate"; //29
const SIPHeaderNames_MAX_FORWARDS="Max-Forwards"; //30
const SIPHeaderNames_ORGANIZATION="Organization"; //31
const SIPHeaderNames_PROXY_AUTHORIZATION="Proxy-Authorization"; //32
const SIPHeaderNames_PROXY_REQUIRE="Proxy-Require"; //33
const SIPHeaderNames_REQUIRE="Require"; //34
const SIPHeaderNames_CONTENT_DISPOSITION="Content-Disposition";//35
const SIPHeaderNames_SUBJECT="Subject";//36
const SIPHeaderNames_USER_AGENT="User-Agent";//37
const SIPHeaderNames_WARNING="Warning"; //38
const SIPHeaderNames_PRIORITY="Priority"; //39
const SIPHeaderNames_DATE="Date"; //40
const SIPHeaderNames_EXPIRES="Expires"; //41
const SIPHeaderNames_SUPPORTED = "Supported";//42
const SIPHeaderNames_AUTHENTICATION_INFO="Authentication-Info";//43
const SIPHeaderNames_REPLY_TO = "Reply-To"; //44
const SIPHeaderNames_RACK	= "RAck";//45
const SIPHeaderNames_RSEQ	= "RSeq";//46
const SIPHeaderNames_REASON = "Reason";//47
const SIPHeaderNames_SUBSCRIPTION_STATE = "Subscription-State";//48
const SIPHeaderNames_EVENT = "Event";//44
const SIPHeaderNames_ALLOW_EVENTS= "Allow-Events";//45
        
/**
*  Root class from which all SIPHeader objects are subclassed.
*/
type SIPHeaderImpl struct{ //implements SIPHeaderNames, javax.sip.header.Header {
	// SIPObject 
	
        
        /** name of this header
         */        
	headerName string;
}
	/** Value of the header.
	*/

        /** Constructor
         * @param hname String to set
         */        
	func NewSIPHeaderImpl( hname string) *SIPHeaderImpl {
		return &SIPHeaderImpl{headerName : hname};
	}
        
    func (this *SIPHeaderImpl) super (hname string) {
    	this.headerName = hname;
    }
        
        /**
         * Name of the SIPHeader
         * @return String
         */
	func (this *SIPHeaderImpl) GetHeaderName() string {
		return this.headerName;
	}	

	/** Alias for getHaderName above.
	*
	*@return String headerName
	*
	*/
	func (this *SIPHeaderImpl) GetName() string { 
		return this.headerName; 
	}
        




	/**
         * Set the name of the header .
         * @param hdrname String to set
         */
	func (this *SIPHeaderImpl) SetHeaderName( hdrname string) {
		this.headerName = hdrname;
	}


	/** Get the header value (i.e. what follows the name:).
	* This merely goes through and lops off the portion that follows
	* the headerName:
	*/
	func (this *SIPHeaderImpl) GetHeaderValue() string {
		//String encodedHdr = null;
		//try {
		   encodedHdr := this.String();
		//} catch (Exception ex) {
		//	return null;
		//}
		//var buffer bytes.Buffer;//new StringBuffer(encodedHdr);
		buffer:= []byte(encodedHdr);
		for len(buffer) > 0 && buffer[0] != ':' {
			buffer=buffer[1:];
		}
		
		if len(buffer) > 0 {
			buffer=buffer[1:];
		}
		
		return strings.TrimSpace(string(buffer));
	}

	/** Return false if this is not a header list 
	* (SIPHeaderList overrrides this method).
	*@return false
	*/
	func (this *SIPHeaderImpl) IsHeaderList() bool{ 
		return false; 
	}

	/** Encode this header into canonical form.
	*/
	func (this *SIPHeaderImpl) String() string{
		return this.headerName + core.Separators_COLON + 
			core.Separators_SP + this.EncodeBody() + core.Separators_NEWLINE;
	}
	
	func (this *SIPHeaderImpl) Clone() interface{} {
		return &SIPHeaderImpl{headerName : this.headerName};
	}

	/** Encode the body of this header (the stuff that follows headerName).
	* A.K.A headerValue.
	*/
	func (this *SIPHeaderImpl) EncodeBody() string{
		return "";
	}
        
        
        /** Alias for getHeaderValue.
         */
    func (this *SIPHeaderImpl) GetValue() string{ 
        return this.GetHeaderValue(); 
    }
		