package parser

import (
	"gosips/core"
	"gosips/sip/header"
)

/** SIPParser for MinExpires header.
 */
type MinExpiresParser struct {
	HeaderParser
}

/** protected constructor.
 *@param text is the text of the header to parse
 */
func NewMinExpiresParser(minExpires string) *MinExpiresParser {
	this := &MinExpiresParser{}
	this.HeaderParser.super(minExpires)
	return this
}

/** constructor.
 *@param lexer is the lexer passed in from the enclosing parser.
 */
func NewMinExpiresParserFromLexer(lexer core.Lexer) *MinExpiresParser {
	this := &MinExpiresParser{}
	this.HeaderParser.superFromLexer(lexer)
	return this
}

/** parse the String message
 * @return Header (MinExpiresParser)
 * @throws SIPParseException if the message does not respect the spec.
 */
func (this *MinExpiresParser) Parse() (sh header.Header, ParseException error) {
	minExpires := header.NewMinExpires()

	lexer := this.GetLexer()
	this.HeaderName(TokenTypes_MIN_EXPIRES)

	minExpires.SetHeaderName(core.SIPHeaderNames_MIN_EXPIRES)

	var number int
	if number, ParseException = lexer.Number(); ParseException != nil {
		return nil, ParseException
	}

	minExpires.SetExpires(number)

	lexer.SPorHT()

	lexer.Match('\n')

	return minExpires, nil
}
