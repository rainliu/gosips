package parser

import (
	"gosip/core"
	"gosip/header"
)

/** Parser for authorization headers.
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
 * @param Lexer lexer to set
 */
func NewAuthorizationParserFromLexer(lexer core.Lexer) *AuthorizationParser {
	this := &AuthorizationParser{}
	this.ChallengeParser.superFromLexer(lexer)
	return this
}

/** parse the String message
 * @return SIPHeaderHeader (Authorization object)
 * @throws SIPParseException if the message does not respect the spec.
 */
func (this *AuthorizationParser) Parse() (sh header.SIPHeaderHeader, ParseException error) {
	//dbg_enter("parse");
	// try {
	this.HeaderName(TokenTypes_AUTHORIZATION)
	auth := header.NewAuthorization()
	this.ChallengeParser.Parse(auth)
	return auth, nil
	//   } finally {
	// dbg_leave("parse");

	//   }

}
