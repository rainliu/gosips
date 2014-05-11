package parser

import (
	"gosips/core"
	"gosips/sip/header"
)

/** SIPParser for Require header.
*
*@version  JAIN-SIP-1.1
*
*@author Olivier Deruelle <deruelle@nist.gov>
*@author M. Ranganathan <mranga@nist.gov>  <br/>
*
*<a href="{@docRoot}/uncopyright.html">This code is in the public domain.</a>
*
* @version 1.0
 */
type RequireParser struct {
	HeaderParser
}

/** Creates a new instance of RequireParser
 * @param require the header to parse
 */
func NewRequireParser(require string) *RequireParser {
	this := &RequireParser{}
	this.HeaderParser.super(require)
	return this
}

/** Constructor
 * @param lexer the lexer to use to parse the header
 */
func NewRequireParserFromLexer(lexer core.Lexer) *RequireParser {
	this := &RequireParser{}
	this.HeaderParser.superFromLexer(lexer)
	return this
}

/** parse the String message
 * @return SIPHeader (RequireList object)
 * @throws SIPParseException if the message does not respect the spec.
 */
func (this *RequireParser) Parse() (sh header.Header, ParseException error) {
	requireList := header.NewRequireList()
	//if (debug) dbg_enter("RequireParser.parse");

	// try {
	var ch byte
	lexer := this.GetLexer()
	this.HeaderName(TokenTypes_REQUIRE)

	for ch, _ = lexer.LookAheadK(0); ch != '\n'; ch, _ = lexer.LookAheadK(0) {
		r := header.NewRequire()
		r.SetHeaderName(core.SIPHeaderNames_REQUIRE)

		// Parsing the option tag
		lexer.Match(TokenTypes_ID)
		token := lexer.GetNextToken()
		r.SetOptionTag(token.GetTokenValue())
		lexer.SPorHT()

		requireList.PushBack(r)

		for ch, _ = lexer.LookAheadK(0); ch == ','; ch, _ = lexer.LookAheadK(0) {
			lexer.Match(',')
			lexer.SPorHT()

			r = header.NewRequire()

			// Parsing the option tag
			lexer.Match(TokenTypes_ID)
			token = lexer.GetNextToken()
			r.SetOptionTag(token.GetTokenValue())
			lexer.SPorHT()

			requireList.PushBack(r)
		}

	}
	// } finally {
	//     if (debug) dbg_leave("RequireParser.parse");
	// }

	return requireList, nil
}
