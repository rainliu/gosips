package parser

import (
	"gosips/core"
	"gosips/sip/header"
)

/** SIPParser for Unsupported header.
 */
type UnsupportedParser struct {
	HeaderParser
}

/** Creates a new instance of UnsupportedParser
 * @param unsupported - Unsupported header to parse
 */
func NewUnsupportedParser(unsupported string) *UnsupportedParser {
	this := &UnsupportedParser{}
	this.HeaderParser.super(unsupported)
	return this
}

/** Constructor
 * @param lexer - the lexer to use to parse the header
 */
func NewUnsupportedParserFromLexer(lexer core.Lexer) *UnsupportedParser {
	this := &UnsupportedParser{}
	this.HeaderParser.superFromLexer(lexer)
	return this
}

/** parse the String message
 * @return SIPHeader (Unsupported object)
 * @throws SIPParseException if the message does not respect the spec.
 */
func (this *UnsupportedParser) Parse() (sh header.Header, ParseException error) {
	unsupportedList := header.NewUnsupportedList()

	var ch byte
	lexer := this.GetLexer()
	this.HeaderName(TokenTypes_UNSUPPORTED)

	for ch, _ = lexer.LookAheadK(0); ch != '\n'; ch, _ = lexer.LookAheadK(0) {
		lexer.SPorHT()
		unsupported := header.NewUnsupported()
		unsupported.SetHeaderName(core.SIPHeaderNames_UNSUPPORTED)

		// Parsing the option tag
		lexer.Match(TokenTypes_ID)
		token := lexer.GetNextToken()
		unsupported.SetOptionTag(token.GetTokenValue())
		lexer.SPorHT()

		unsupportedList.PushBack(unsupported)

		for ch, _ = lexer.LookAheadK(0); ch == ','; ch, _ = lexer.LookAheadK(0) {
			lexer.Match(',')
			lexer.SPorHT()

			unsupported = header.NewUnsupported()

			// Parsing the option tag
			lexer.Match(TokenTypes_ID)
			token = lexer.GetNextToken()
			unsupported.SetOptionTag(token.GetTokenValue())
			lexer.SPorHT()

			unsupportedList.PushBack(unsupported)
		}

	}

	return unsupportedList, nil
}
