package header

import (
	"bytes"
	"container/list"
	"gosip/core"
)

/**
* the UserAgent SIPObject.
*
*@author Olivier Deruelle <deruelle@nist.gov><br/>
*
*<a href="{@docRoot}/uncopyright.html">This code is in the public domain.</a>
*
 */
type UserAgent struct {
	SIPHeader
	//implements  UserAgentHeader {

	/** Product tokens.
	 */
	productTokens *list.List
}

/**
 * Constructor.
 */
func NewUserAgent() *UserAgent {
	this := &UserAgent{}
	this.SIPHeader.super(core.SIPHeaderNames_USER_AGENT)
	this.productTokens = list.New()
	return this
}

/**
 * Return canonical form.
 * @return String
 */
func (this *UserAgent) EncodeProduct() string {
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
func (this *UserAgent) AddProductToken(pt string) {
	this.productTokens.PushBack(pt)
}

func (this *UserAgent) String() string {
	return this.headerName + core.SIPSeparatorNames_COLON +
		core.SIPSeparatorNames_SP + this.EncodeBody() + core.SIPSeparatorNames_NEWLINE
}

/** Encode only the body of this header.
*@return encoded value of the header.
 */
func (this *UserAgent) EncodeBody() string {
	return this.EncodeProduct()
}

/**
 * Returns the list value of the product parameter.
 *
 * @return the software of this UserAgentHeader
 */
func (this *UserAgent) GetProduct() *list.List {
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
func (this *UserAgent) SetProduct(product *list.List) { //throws ParseException {
	//      if (product==null) throw new  NullPointerException
	// ("JAIN-SIP Exception, UserAgent, "+
	//      		"setProduct(), the "+
	//           	" product parameter is null");
	this.productTokens = product
}
