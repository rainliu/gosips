package parser

import (
	"gosips/core"
	"gosips/sip/header"
)

/** Parser for Supported header.
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
type SupportedParser struct {
	HeaderParser
}

/** Creates a new instance of SupportedParser
 * @param supported the header to parse
 */
func NewSupportedParser(supported string) *SupportedParser {
	this := &SupportedParser{}
	this.HeaderParser.super(supported)
	return this
}

/** Constructor
 * @param lexer the lexer to use to parse the header
 */
func NewSupportedParserFromLexer(lexer core.Lexer) *SupportedParser {
	this := &SupportedParser{}
	this.HeaderParser.superFromLexer(lexer)
	return this
}

/** parse the String message
 * @return SIPHeader (Supported object)
 * @throws SIPParseException if the message does not respect the spec.
 */
func (this *SupportedParser) Parse() (sh header.Header, ParseException error) {
	supportedList := header.NewSupportedList()
	//if (debug) dbg_enter("SupportedParser.parse");

	//try {
	var ch byte
	lexer := this.GetLexer()
	this.HeaderName(TokenTypes_SUPPORTED)

	for ch, _ = lexer.LookAheadK(0); ch != '\n'; ch, _ = lexer.LookAheadK(0) {
		// while (lexer.lookAhead(0) != '\n') {
		lexer.SPorHT()
		supported := header.NewSupported()
		supported.SetHeaderName(core.SIPHeaderNames_SUPPORTED)

		// Parsing the option tag
		lexer.Match(TokenTypes_ID)
		token := lexer.GetNextToken()
		supported.SetOptionTag(token.GetTokenValue())
		lexer.SPorHT()

		supportedList.PushBack(supported)

		for ch, _ = lexer.LookAheadK(0); ch == ','; ch, _ = lexer.LookAheadK(0) {
			//while (lexer.lookAhead(0) == ',') {
			lexer.Match(',')
			lexer.SPorHT()

			supported = header.NewSupported()

			// Parsing the option tag
			lexer.Match(TokenTypes_ID)
			token = lexer.GetNextToken()
			supported.SetOptionTag(token.GetTokenValue())
			lexer.SPorHT()

			supportedList.PushBack(supported)
		}

	}
	// } finally {
	//     if (debug) dbg_leave("SupportedParser.parse");
	// }

	return supportedList, nil
}
