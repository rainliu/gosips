package parser

import (
	"gosips/core"
	"gosips/sip/header"
)

/** SIPParser for ProxyAuthenticate headers.
 */
type ProxyAuthenticateParser struct {
	ChallengeParser
}

/** Constructor
 * @param String paAuthenticate message to parse
 */
func NewProxyAuthenticateParser(proxyAuthenticate string) *ProxyAuthenticateParser {
	this := &ProxyAuthenticateParser{}
	this.ChallengeParser.super(proxyAuthenticate)
	return this
}

/** Cosntructor
 * @param lexer to set
 */
func NewProxyAuthenticateParserFromLexer(lexer core.Lexer) *ProxyAuthenticateParser {
	this := &ProxyAuthenticateParser{}
	this.ChallengeParser.superFromLexer(lexer)
	return this
}

/** parse the String message
 * @return SIPHeader (ProxyAuthenticate object)
 * @throws ParseException if the message does not respect the spec.
 */
func (this *ProxyAuthenticateParser) Parse() (sh header.Header, ParseException error) {
	this.HeaderName(TokenTypes_PROXY_AUTHENTICATE)
	proxyAuthenticate := header.NewProxyAuthenticate()
	ParseException = this.ChallengeParser.Parse(proxyAuthenticate)
	return proxyAuthenticate, ParseException
}
