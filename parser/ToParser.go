package parser

import (
	"gosip/address"
	"gosip/core"
	"gosip/header"
)

/* To Header parser.*/
type ToParser struct {
	AddressParametersParser
}

/** Creates new ToParser
 * @param String to set
 */
func NewToParser(to string) *ToParser {
	this := &ToParser{}
	this.AddressParametersParser.super(to)
	return this
}

func NewToParserFromLexer(lexer core.Lexer) *ToParser {
	this := &ToParser{}
	this.AddressParametersParser.superFromLexer(lexer)
	return this
}

func (this *ToParser) Parse() (sh header.SIPHeaderHeader, ParseException error) {
	to := header.NewTo()
	this.HeaderName(TokenTypes_TO)
	this.AddressParametersParser.Parse(to)
	this.GetLexer().Match('\n')
	addr, _ := to.GetAddress().(*address.AddressImpl)
	if addr.GetAddressType() == address.ADDRESS_SPEC {
		// the parameters are header parameters.
		if to.GetAddress().GetURI().IsSipURI() {
			sipUri, _ := to.GetAddress().GetURI().(*address.SipUri)
			parms := sipUri.GetUriParms()
			if parms != nil && parms.Len() > 0 {
				to.SetParameters(parms)
				sipUri.RemoveUriParms()
			}
		}
	}
	return to, nil
}
