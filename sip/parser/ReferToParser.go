package parser

import (
	"gosips/core"
	"gosips/sip/header"
)

/** ReferTo Header parser.
 */
type ReferToParser struct {
	AddressParametersParser
}

/** Creates new ReferToParser
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
	if ParseException = this.AddressParametersParser.Parse(referTo); ParseException != nil {
		return nil, ParseException
	}
	lexer.Match('\n')
	return referTo, nil
}
