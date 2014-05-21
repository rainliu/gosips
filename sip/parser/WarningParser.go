package parser

import (
	"gosips/core"
	"gosips/sip/header"
	"strconv"
)

/** SIPParser for Warning header.
 */
type WarningParser struct {
	HeaderParser
}

/** Constructor
 * @param warning - Warning header to parse
 */
func NewWarningParser(warning string) *WarningParser {
	this := &WarningParser{}
	this.HeaderParser.super(warning)
	return this
}

/** Cosntructor
 * @param lexer - the lexer to use.
 */
func NewWarningParserFromLexer(lexer core.Lexer) *WarningParser {
	this := &WarningParser{}
	this.HeaderParser.superFromLexer(lexer)
	return this
}

/** parse the String message
 * @return SIPHeader (WarningList object)
 * @throws SIPParseException if the message does not respect the spec.
 */
func (this *WarningParser) Parse() (sh header.Header, ParseException error) {
	warningList := header.NewWarningList()

	var ch byte
	lexer := this.GetLexer()
	this.HeaderName(TokenTypes_WARNING)

	for ch, _ = lexer.LookAheadK(0); ch != '\n'; ch, _ = lexer.LookAheadK(0) {
		warning := header.NewWarning()
		warning.SetHeaderName(core.SIPHeaderNames_WARNING)

		// Parsing the 3digits code
		lexer.Match(TokenTypes_ID)
		token := lexer.GetNextToken()

		var code int
		if code, ParseException = strconv.Atoi(token.GetTokenValue()); ParseException != nil {
			return nil, ParseException
		}
		if ParseException = warning.SetCode(code); ParseException != nil {
			return nil, ParseException
		}

		lexer.SPorHT()

		// Parsing the agent
		lexer.Match(TokenTypes_ID)
		token = lexer.GetNextToken()
		warning.SetAgent(token.GetTokenValue())
		lexer.SPorHT()

		// Parsing the text
		var text string
		if text, ParseException = lexer.QuotedString(); ParseException != nil {
			return nil, ParseException
		}
		warning.SetText(text)
		lexer.SPorHT()

		warningList.PushBack(warning)

		for ch, _ = lexer.LookAheadK(0); ch == ','; ch, _ = lexer.LookAheadK(0) {
			lexer.Match(',')
			lexer.SPorHT()

			warning = header.NewWarning()

			// Parsing the 3digits code
			lexer.Match(TokenTypes_ID)
			tok := lexer.GetNextToken()

			if code, ParseException = strconv.Atoi(tok.GetTokenValue()); ParseException != nil {
				return nil, ParseException
			}
			warning.SetCode(code)

			lexer.SPorHT()

			// Parsing the agent
			lexer.Match(TokenTypes_ID)
			tok = lexer.GetNextToken()
			warning.SetAgent(tok.GetTokenValue())
			lexer.SPorHT()

			// Parsing the text
			if text, ParseException = lexer.QuotedString(); ParseException != nil {
				return nil, ParseException
			}
			warning.SetText(text)
			lexer.SPorHT()

			warningList.PushBack(warning)
		}

	}

	return warningList, nil
}
