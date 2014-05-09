package header

import (
	"bytes"
	"container/list"
	"gosips/core"
)

/**

*Supported SIP Header.

*

*@version  JAIN-SIP-1.1

*

*@author M. Ranganathan <mranga@nist.gov>  <br/>

*@author Olivier Deruelle <deruelle@nist.gov><br/>

*

*<a href="{@docRoot}/uncopyright.html">This code is in the public domain.</a>

*

 */
type Server struct {
	SIPHeader
	//implements ServerHeader{

	/** Product tokens.
	 */
	productTokens *list.List
}

/**
 * Constructor.
 */
func NewServer() *Server {
	this := &Server{}
	this.SIPHeader.super(core.SIPHeaderNames_SERVER)
	this.productTokens = list.New()
	return this
}

/**
 * Return canonical form.
 * @return String
 */
func (this *Server) EncodeProduct() string {
	var encoding bytes.Buffer //  = new StringBuffer();

	for e := this.productTokens.Front(); e != nil; e = e.Next() {
		encoding.WriteString(e.Value.(string))

		if e.Next() != nil {
			encoding.WriteString("/")
		}
	}
	return encoding.String()
}

/** set the productToken field
 * @param pt String to set
 */
func (this *Server) AddProductToken(pt string) {
	this.productTokens.PushBack(pt)
}

func (this *Server) String() string {
	return this.headerName + core.SIPSeparatorNames_COLON +
		core.SIPSeparatorNames_SP + this.EncodeBody() + core.SIPSeparatorNames_NEWLINE
}

/** Encode only the body of this header.
*@return encoded value of the header.
 */
func (this *Server) EncodeBody() string {
	return this.EncodeProduct()
}

/**
 * Returns the list value of the product parameter.
 *
 * @return the software of this UserAgentHeader
 */
func (this *Server) GetProduct() *list.List {
	if this.productTokens == nil || this.productTokens.Len() == 0 {
		return nil
	} else {
		return this.productTokens
	}
}

/**
 * Sets the product value of the UserAgentHeader.
 *
 * @param product - a List specifying the product value
 * @throws ParseException which signals that an error has been reached
 * unexpectedly while parsing the product value.
 */
func (this *Server) SetProduct(product *list.List) { //throws ParseException {
	//      if (product==null) throw new  NullPointerException
	// ("JAIN-SIP Exception, UserAgent, "+
	//      		"setProduct(), the "+
	//           	" product parameter is null");
	this.productTokens = product
}
