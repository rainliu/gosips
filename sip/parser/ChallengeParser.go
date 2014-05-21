package parser

import (
	"gosips/core"
	"gosips/sip/header"
)

/** SIPParser for the challenge portion of the authentication header.
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
func (this *ChallengeParser) ParseParameter(h header.AuthorizationHeader) (ParseException error) {
	var nv *core.NameValue
	if nv, ParseException = this.NameValue('='); ParseException != nil {
		return ParseException
	}
	ParseException = h.SetParameter(nv.GetName(), nv.GetValue().(string))
	return ParseException
}

/** parser the String message
 * @return Challenge object
 * @throws ParseException if the message does not respect the spec.
 */
func (this *ChallengeParser) Parse(h header.AuthorizationHeader) (ParseException error) {
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
		if ParseException = this.ParseParameter(h); ParseException != nil {
			return ParseException
		}
		lexer.SPorHT()
		if ch, ParseException = lexer.LookAheadK(0); ch == '\n' || ParseException != nil { //||ch=='\0'
			break
		}
		lexer.Match(',')
		lexer.SPorHT()
	}
	return nil
}
