package parser

import (
	"gosips/core"
	"gosips/sip/header"
)

/** To Header parser.
*
*@version  JAIN-SIP-1.1
*
*@author Olivier Deruelle <deruelle@nist.gov>  <br/>
*
*<a href="{@docRoot}/uncopyright.html">This code is in the public domain.</a>
*
 */
type ReferToParser struct {
	AddressParametersParser
}

/** Creates new ToParser
 * @param String to set
 */
func NewReferToParser(referTo string) *ReferToParser {
	this := &ReferToParser{}
	this.AddressParametersParser.super(referTo)
	return this
}

func NewReferToParserFromLexer(lexer core.Lexer) *ReferToParser {
	this := &ReferToParser{}
	this.AddressParametersParser.superFromLexer(lexer)
	return this
}

func (this *ReferToParser) Parse() (sh header.Header, ParseException error) {
	lexer := this.GetLexer()
	this.HeaderName(TokenTypes_REFER_TO)
	referTo := header.NewReferTo()
	this.AddressParametersParser.Parse(referTo)
	lexer.Match('\n')
	return referTo, nil
}
