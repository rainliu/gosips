package parser

import (
	"gosips/core"
	"gosips/sip/header"
)

/** SIPParser for RAck header.
 */
type RAckParser struct {
	HeaderParser
}

/** Creates a new instance of RAckParser
 *@param rack the header to parse
 */
func NewRAckParser(rack string) *RAckParser {
	this := &RAckParser{}
	this.HeaderParser.super(rack)
	return this
}

/** Constructor
 * @param lexer the lexer to use to parse the header
 */
func NewRAckParserFromLexer(lexer core.Lexer) *RAckParser {
	this := &RAckParser{}
	this.HeaderParser.superFromLexer(lexer)
	return this
}

/** parse the String message
 * @return SIPHeader (RAck object)
 * @throws SIPParseException if the message does not respect the spec.
 */
func (this *RAckParser) Parse() (sh header.Header, ParseException error) {
	rack := header.NewRAck()

	lexer := this.GetLexer()
	this.HeaderName(TokenTypes_RACK)

	rack.SetHeaderName(core.SIPHeaderNames_RACK)

	var number int
	if number, ParseException = lexer.Number(); ParseException != nil {
		return nil, ParseException
	}
	rack.SetRSeqNumber(number)
	lexer.SPorHT()
	if number, ParseException = lexer.Number(); ParseException != nil {
		return nil, ParseException
	}
	rack.SetCSeqNumber(number)
	lexer.SPorHT()
	lexer.Match(TokenTypes_ID)
	token := lexer.GetNextToken()
	rack.SetMethod(token.GetTokenValue())

	lexer.SPorHT()
	lexer.Match('\n')

	return rack, nil
}
