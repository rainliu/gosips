package parser

import (
	"gosips/core"
	"gosips/sip/header"
)

/** SIPParser for ProxyRequire header.
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
type ProxyRequireParser struct {
	HeaderParser
}

/** Creates a new instance of ProxyRequireParser
 *@param require the header to parse
 */
func NewProxyRequireParser(require string) *ProxyRequireParser {
	this := &ProxyRequireParser{}
	this.HeaderParser.super(require)
	return this
}

/** Constructor
 * @param lexer the lexer to use to parse the header
 */
func NewProxyRequireParserFromLexer(lexer core.Lexer) *ProxyRequireParser {
	this := &ProxyRequireParser{}
	this.HeaderParser.superFromLexer(lexer)
	return this
}

/** parse the String message
 * @return SIPHeader (ProxyRequireList object)
 * @throws SIPParseException if the message does not respect the spec.
 */
func (this *ProxyRequireParser) Parse() (sh header.Header, ParseException error) {
	proxyRequireList := header.NewProxyRequireList()
	// if (debug) dbg_enter("ProxyRequireParser.parse");

	//try {
	var ch byte
	lexer := this.GetLexer()
	this.HeaderName(TokenTypes_PROXY_REQUIRE)

	for ch, _ = lexer.LookAheadK(0); ch != '\n'; ch, _ = lexer.LookAheadK(0) {
		r := header.NewProxyRequire()
		r.SetHeaderName(core.SIPHeaderNames_PROXY_REQUIRE)

		// Parsing the option tag
		lexer.Match(TokenTypes_ID)
		token := lexer.GetNextToken()
		r.SetOptionTag(token.GetTokenValue())
		lexer.SPorHT()

		proxyRequireList.PushBack(r)

		for ch, _ = lexer.LookAheadK(0); ch == ','; ch, _ = lexer.LookAheadK(0) {
			// while (lexer.lookAhead(0) == ',') {
			lexer.Match(',')
			lexer.SPorHT()

			r = header.NewProxyRequire()

			// Parsing the option tag
			lexer.Match(TokenTypes_ID)
			token = lexer.GetNextToken()
			r.SetOptionTag(token.GetTokenValue())
			lexer.SPorHT()

			proxyRequireList.PushBack(r)
		}

	}
	// } finally {
	//     if (debug) dbg_leave("ProxyRequireParser.parse");
	// }

	return proxyRequireList, nil
}
