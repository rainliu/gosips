package parser

import (
	"gosip/core"
	"gosip/header"
)

/** Parser for ProxyAuthorization headers.
*
*@version  JAIN-SIP-1.1
*
*@author M. Ranganathan <mranga@nist.gov>  <br/>
*
*<a href="{@docRoot}/uncopyright.html">This code is in the public domain.</a>
*
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
 * @param Lexer lexer to set
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
func (this *ProxyAuthorizationParser) Parse() (sh header.ISIPHeader, ParseException error) {
	this.HeaderName(TokenTypes_PROXY_AUTHORIZATION)
	proxyAuth := header.NewProxyAuthorization()
	this.ChallengeParser.Parse(proxyAuth)
	return proxyAuth, nil
}
