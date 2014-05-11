package parser

import (
	"gosips/core"
	"gosips/sip/header"
	"strconv"
)

/** SIPParser for Warning header.
*
*@version  JAIN-SIP-1.1
*
*@author Olivier Deruelle <deruelle@nist.gov>  <br/>
*@author M. Ranganathan <mranga@nist.gov>  <br/>
*
*<a href="{@docRoot}/uncopyright.html">This code is in the public domain.</a>
*
* @version 1.0
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
	// if (debug) dbg_enter("WarningParser.parse");

	//try {
	var ch byte
	lexer := this.GetLexer()
	this.HeaderName(TokenTypes_WARNING)

	for ch, _ = lexer.LookAheadK(0); ch != '\n'; ch, _ = lexer.LookAheadK(0) {
		//while (lexer.lookAhead(0) != '\n') {
		warning := header.NewWarning()
		warning.SetHeaderName(core.SIPHeaderNames_WARNING)

		// Parsing the 3digits code
		lexer.Match(TokenTypes_ID)
		token := lexer.GetNextToken()
		//try {
		code, _ := strconv.Atoi(token.GetTokenValue())

		warning.SetCode(code)
		// } catch (NumberFormatException ex) {
		//     throw createParseException(ex.GetMessage());
		// } catch (InvalidArgumentException ex) {
		//     throw createParseException(ex.GetMessage());
		// }
		lexer.SPorHT()

		// Parsing the agent
		lexer.Match(TokenTypes_ID)
		token = lexer.GetNextToken()
		warning.SetAgent(token.GetTokenValue())
		lexer.SPorHT()

		// Parsing the text
		text, _ := lexer.QuotedString()
		warning.SetText(text)
		lexer.SPorHT()

		warningList.PushBack(warning)

		for ch, _ = lexer.LookAheadK(0); ch == ','; ch, _ = lexer.LookAheadK(0) {
			//while (lexer.lookAhead(0) == ',') {
			lexer.Match(',')
			lexer.SPorHT()

			warning = header.NewWarning()

			// Parsing the 3digits code
			lexer.Match(TokenTypes_ID)
			tok := lexer.GetNextToken()
			//try {
			code, _ = strconv.Atoi(tok.GetTokenValue())
			warning.SetCode(code)
			// } catch (NumberFormatException ex) {
			//     throw createParseException(ex.GetMessage());
			// } catch (InvalidArgumentException ex) {
			//     throw createParseException(ex.GetMessage());
			// }
			lexer.SPorHT()

			// Parsing the agent
			lexer.Match(TokenTypes_ID)
			tok = lexer.GetNextToken()
			warning.SetAgent(tok.GetTokenValue())
			lexer.SPorHT()

			// Parsing the text
			text, _ = lexer.QuotedString()
			warning.SetText(text)
			lexer.SPorHT()

			warningList.PushBack(warning)
		}

	}
	// } finally {
	//     if (debug) dbg_leave("WarningParser.parse");
	// }

	return warningList, nil
}
