package parser

import (
	"gosip/core"
	"gosip/header"
)

/** Parser for Unsupported header.
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
func (this *UnsupportedParser) Parse() (sh header.ISIPHeader, ParseException error) {
	unsupportedList := header.NewUnsupportedList()
	//if (debug) dbg_enter("UnsupportedParser.parse");

	//try {
	var ch byte
	lexer := this.GetLexer()
	this.HeaderName(TokenTypes_UNSUPPORTED)

	for ch, _ = lexer.LookAheadK(0); ch != '\n'; ch, _ = lexer.LookAheadK(0) {
		//while (lexer.lookAhead(0) != '\n') {
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
			// while (lexer.lookAhead(0) == ',') {
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
	// } finally {
	//     if (debug) dbg_leave("UnsupportedParser.parse");
	// }

	return unsupportedList, nil
}
