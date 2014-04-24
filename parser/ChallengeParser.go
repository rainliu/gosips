package parser

import (
	"gosip/core"
	"gosip/header"
)

/** Parser for the challenge portion of the authentication header.
 *
 *@version  JAIN-SIP-1.1
 *
 *@author Olivier Deruelle  <deruelle@nist.gov>  <br/>
 *
 *<a href="{@docRoot}/uncopyright.html">This code is in the public domain.</a>
 *
 * @version 1.0
 */

type ChallengeParser struct {
	HeaderParser
}

/** Constructor
 * @param String challenge  message to parse to set
 */
func NewChallengeParser(challenge string) *ChallengeParser {
	this := &ChallengeParser{}
	this.HeaderParser.super(challenge)
	return this
}

/** Constructor
 * @param String challenge  message to parse to set
 */
func NewChallengeParserFromLexer(lexer core.Lexer) *ChallengeParser {
	this := &ChallengeParser{}
	this.HeaderParser.superFromLexer(lexer)
	return this
}

func (this *ChallengeParser) super(buffer string) {
	this.HeaderParser.super(buffer)
}

func (this *ChallengeParser) superFromLexer(lexer core.Lexer) {
	this.HeaderParser.superFromLexer(lexer)
}

/** Get the parameter of the challenge string
 * @return NameValue containing the parameter
 */
func (this *ChallengeParser) ParseParameter(h header.AuthorizationHeader) { //} error{
	//if (debug) dbg_enter("parseParameter");
	//try {
	nv := this.NameValue('=')
	h.SetParameter(nv.GetName(), nv.GetValue().(string))
	// } finally {
	//if (debug) dbg_leave("parseParameter");
	// }
}

/** parser the String message
 * @return Challenge object
 * @throws ParseException if the message does not respect the spec.
 */
func (this *ChallengeParser) Parse(h header.AuthorizationHeader) error { //throws ParseException {
	var ch byte
	lexer := this.GetLexer()
	// the Scheme:
	lexer.SPorHT()
	lexer.Match(TokenTypes_ID)
	t := lexer.GetNextToken()
	lexer.SPorHT()
	h.SetScheme(t.GetTokenValue())

	// The parameters:
	for ch, _ = lexer.LookAheadK(0); ch != '\n'; ch, _ = lexer.LookAheadK(0) {
		this.ParseParameter(h)
		lexer.SPorHT()
		if ch, _ = lexer.LookAheadK(0); ch == '\n' { //||ch=='\0'
			break
		}
		lexer.Match(',')
		lexer.SPorHT()
	}
	return nil
}
