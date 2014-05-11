package parser

import (
	"gosips/core"
	"gosips/sip/header"
)

/** SIPParser for ProxyAuthenticate headers.
*
*@version  JAIN-SIP-1.1
*
*@author M. Ranganathan <mranga@nist.gov>  <br/>
*
*<a href="{@docRoot}/uncopyright.html">This code is in the public domain.</a>
*
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
	this.ChallengeParser.Parse(proxyAuthenticate)
	return proxyAuthenticate, nil
}
