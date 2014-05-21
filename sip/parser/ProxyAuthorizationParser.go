package parser

import (
	"gosips/core"
	"gosips/sip/header"
)

/** SIPParser for ProxyAuthorization headers.
 */
type ProxyAuthorizationParser struct {
	ChallengeParser
}

/** Constructor
 * @param proxyAuthorization --  header to parse
 */
func NewProxyAuthorizationParser(proxyAuthorization string) *ProxyAuthorizationParser {
	this := &ProxyAuthorizationParser{}
	this.ChallengeParser.super(proxyAuthorization)
	return this
}

/** Cosntructor
 * @param lexer to set
 */
func NewProxyAuthorizationParserFromLexer(lexer core.Lexer) *ProxyAuthorizationParser {
	this := &ProxyAuthorizationParser{}
	this.ChallengeParser.superFromLexer(lexer)
	return this
}

/** parse the String message
 * @return SIPHeader (ProxyAuthenticate object)
 * @throws ParseException if the message does not respect the spec.
 */
func (this *ProxyAuthorizationParser) Parse() (sh header.Header, ParseException error) {
	this.HeaderName(TokenTypes_PROXY_AUTHORIZATION)
	proxyAuth := header.NewProxyAuthorization()
	ParseException = this.ChallengeParser.Parse(proxyAuth)
	return proxyAuth, ParseException
}
