package header

import (
	"strings"
	"gosip/core"
)


        
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
		return this.headerName + core.SIPSeparatorNames_COLON + 
			core.SIPSeparatorNames_SP + this.EncodeBody() + core.SIPSeparatorNames_NEWLINE;
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
		