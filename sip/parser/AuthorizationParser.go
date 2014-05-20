package parser

import (
	"gosips/core"
	"gosips/sip/header"
)

/** SIPParser for authorization headers.
 */
type AuthorizationParser struct {
	ChallengeParser
}

/** Constructor
 * @param String Authorization message to parse
 */
func NewAuthorizationParser(authorization string) *AuthorizationParser {
	this := &AuthorizationParser{}
	this.ChallengeParser.super(authorization)
	return this
}

/** Cosntructor
 * @param lexer to set
 */
func NewAuthorizationParserFromLexer(lexer core.Lexer) *AuthorizationParser {
	this := &AuthorizationParser{}
	this.ChallengeParser.superFromLexer(lexer)
	return this
}

/** parse the String message
 * @return Header (Authorization object)
 * @throws SIPParseException if the message does not respect the spec.
 */
func (this *AuthorizationParser) Parse() (sh header.Header, ParseException error) {
	this.HeaderName(TokenTypes_AUTHORIZATION)
	auth := header.NewAuthorization()
	ParseException = this.ChallengeParser.Parse(auth)
	return auth, ParseException
}
