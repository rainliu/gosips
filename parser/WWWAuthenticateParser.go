package parser

import (
	"gosip/core"
	"gosip/header"
)

/** Parser for WWW authenitcate header.
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
func (this *WWWAuthenticateParser) Parse() (sh header.SIPHeaderHeader, ParseException error) {
	//if (debug) dbg_enter("parse");
	//try {
	this.HeaderName(TokenTypes_WWW_AUTHENTICATE)
	wwwAuthenticate := header.NewWWWAuthenticate()
	this.ChallengeParser.Parse(wwwAuthenticate)
	return wwwAuthenticate, nil
	// } finally {
	//    if (debug) dbg_leave("parse");
	// }
}
