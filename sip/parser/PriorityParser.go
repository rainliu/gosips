package parser

import (
	"gosips/core"
	"gosips/sip/header"
)

/** SIPParser for Priority header.
 */
type PriorityParser struct {
	HeaderParser
}

/** Creates a new instance of PriorityParser
 * @param priority the header to parse
 */
func NewPriorityParser(priority string) *PriorityParser {
	this := &PriorityParser{}
	this.HeaderParser.super(priority)
	return this
}

/** Constructor
 * @param lexer the lexer to use to parse the header
 */
func NewPriorityParserFromLexer(lexer core.Lexer) *PriorityParser {
	this := &PriorityParser{}
	this.HeaderParser.superFromLexer(lexer)
	return this
}

/** parse the String header
 * @return SIPHeader (Priority object)
 * @throws SIPParseException if the message does not respect the spec.
 */
func (this *PriorityParser) Parse() (sh header.Header, ParseException error) {
	priority := header.NewPriority()

	lexer := this.GetLexer()
	this.HeaderName(TokenTypes_PRIORITY)

	priority.SetHeaderName(core.SIPHeaderNames_PRIORITY)

	lexer.SPorHT()
	lexer.Match(TokenTypes_ID)
	token := lexer.GetNextToken()

	priority.SetPriority(token.GetTokenValue())

	lexer.SPorHT()
	lexer.Match('\n')

	return priority, nil
}
