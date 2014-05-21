package parser

import (
	"gosips/core"
	"gosips/sip/header"
)

/** SIPParser for WWW authenitcate header.
 */
type WWWAuthenticateParser struct {
	ChallengeParser
}

/** Constructor
 * @param wwwAuthenticate -  message to parse
 */
func NewWWWAuthenticateParser(wwwAuthenticate string) *WWWAuthenticateParser {
	this := &WWWAuthenticateParser{}
	this.ChallengeParser.super(wwwAuthenticate)
	return this
}

/** Cosntructor
 * @param  lexer - lexer to use.
 */
func NewWWWAuthenticateParserFromLexer(lexer core.Lexer) *WWWAuthenticateParser {
	this := &WWWAuthenticateParser{}
	this.ChallengeParser.superFromLexer(lexer)
	return this
}

/** parse the String message
 * @return SIPHeader (WWWAuthenticate object)
 * @throws SIPParseException if the message does not respect the spec.
 */
func (this *WWWAuthenticateParser) Parse() (sh header.Header, ParseException error) {
	this.HeaderName(TokenTypes_WWW_AUTHENTICATE)
	wwwAuthenticate := header.NewWWWAuthenticate()
	ParseException = this.ChallengeParser.Parse(wwwAuthenticate)
	return wwwAuthenticate, ParseException
}
