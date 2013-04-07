package header

import (
	"bytes"
	"container/list"
	"gosip/core"
)


/**
 *  This is the root class for all lists of SIP headers.
 *  It imbeds a SIPObjectList object and extends SIPHeader
 *  Lists of ContactSIPObjects etc. derive from this class.
 *  This supports homogeneous  lists (all elements in the list are of
 *  the same class). We use this for building type homogeneous lists of
 *  SIPObjects that appear in SIPHeaders
 *
 *
 */
type SIPHeaderListImpl struct{//extends 
	SIPHeaderImpl 
	list.List
}    
    /** Constructor
     * @param hl SIPObjectList to set
     * @param hname String to set
     */
    func NewSIPHeaderListImpl(hname string) *SIPHeaderListImpl{
		this := &SIPHeaderListImpl{};
		this.SIPHeaderImpl.super(hname);
		
		return this;
	}
	
	func (this *SIPHeaderListImpl) super(hname string) {
		this.SIPHeaderImpl.super(hname);
	}
   
   /**
 * Implement the clone method.
 */
func (this *SIPHeaderListImpl) Clone() interface{} {
    retval := &SIPHeaderListImpl{}

    retval.headerName = this.headerName;

    return retval
}


func (this *SIPHeaderListImpl) Concatenate(shl *SIPHeaderListImpl){//, topFlag bool) {
    if shl == nil {
        return
    }

    //if !topFlag {
        for e := shl.Front(); e != nil; e = e.Next() {
            this.PushBack(e)
        }
    /*} else {
        // add given items to the end of the list.
        first := this.Front()
        for e := objList.Front(); e != nil; e = e.Next() {
            this.InsertBefore(e, first)
        }
    }*/
}

/**
     * Encode a list of sip headers.
     * Headers are returned in cannonical form.
     * @return String encoded string representation of this list of
     * 	 headers. (Contains string append of each encoded header).
     */
    func (this *SIPHeaderListImpl) String() string{
        if this.Len()==0 { 
        	return this.headerName + ":" + core.Separators_NEWLINE;
        }
        
        var encoding bytes.Buffer;//= new StringBuffer();
        // The following headers do not have comma separated forms for
        // multiple headers. Thus, they must be encoded separately.
        if  this.headerName==SIPHeaderNames_WWW_AUTHENTICATE ||
        	this.headerName==SIPHeaderNames_PROXY_AUTHENTICATE ||
        	this.headerName==SIPHeaderNames_AUTHORIZATION ||
        	this.headerName==SIPHeaderNames_PROXY_AUTHORIZATION { //||
			//this instanceof ExtensionHeaderList ) {
            //ListIterator li = hlist.listIterator();
            for e := this.Front(); e != nil; e = e.Next() {
		        if sh, ok := e.Value.(SIPHeader); ok {
		            encoding.WriteString(sh.String());
		        } else {
		            encoding.WriteString(e.Value.(string));
		        }
		    }
            
            return encoding.String();
        } else {
	    // These can be concatenated together in an comma separated
	    // list.
            return this.headerName + core.Separators_COLON + core.Separators_SP + this.EncodeBody() + core.Separators_NEWLINE;
        }
    }
    
    /** Encode the body of this header (the stuff that follows headerName).
     * A.K.A headerValue. This will not give a reasonable result for
     *WWW-Authenticate, Authorization, Proxy-Authenticate and
     *Proxy-Authorization and hence this is protected.
     */
    func (this *SIPHeaderListImpl) EncodeBody() string{
        var encoding bytes.Buffer;// = new StringBuffer();
        //ListIterator iterator = this.listIterator();
        for e := this.Front(); e != nil; e = e.Next() {
	        if sh, ok := e.Value.(SIPHeader); ok {
	            encoding.WriteString(sh.EncodeBody());
	        } else {
	            encoding.WriteString(e.Value.(string));
	        }
	        if e.Next()!=nil{
	        	encoding.WriteString(core.Separators_COMMA);
	        }
	    }
       
        return encoding.String();
    }
    
    func (this *SIPHeaderListImpl) IsHeaderList() bool{ 
		return true; 
	}
    