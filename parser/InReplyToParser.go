package parser

import (
	"gosip/core"
	"gosip/header"
)

/** Parser for InReplyTo header.
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
type InReplyToParser struct {
	HeaderParserImpl
}

/** Creates a new instance of InReplyToParser
 * @param inReplyTo the header to parse
 */
func NewInReplyToParser(inReplyTo string) *InReplyToParser {
	this := &InReplyToParser{}
	this.HeaderParserImpl.super(inReplyTo)
	return this
}

/** Constructor
 * @param lexer the lexer to use to parse the header
 */
func NewInReplyToParserFromLexer(lexer core.Lexer) *InReplyToParser {
	this := &InReplyToParser{}
	this.HeaderParserImpl.superFromLexer(lexer)
	return this
}

/** parse the String message
 * @return SIPHeaderHeader (InReplyToList object)
 * @throws SIPParseException if the message does not respect the spec.
 */
func (this *InReplyToParser) Parse() (sh header.SIPHeaderHeader, ParseException error) {

	// if (debug) dbg_enter("InReplyToParser.parse");
	inReplyToList := header.NewInReplyToList()

	// try {
	var ch byte
	lexer := this.GetLexer()
	this.HeaderName(TokenTypes_IN_REPLY_TO)

	for ch, _ = lexer.LookAheadK(0); ch != '\n'; ch, _ = lexer.LookAheadK(0) {
		//while (lexer.lookAhead(0) != '\n') {
		inReplyTo := header.NewInReplyTo()
		inReplyTo.SetHeaderName(core.SIPHeaderNames_IN_REPLY_TO)

		lexer.Match(TokenTypes_ID)
		token := lexer.GetNextToken()
		if ch, _ = lexer.LookAheadK(0); ch == '@' {
			//if (lexer.lookAhead(0)=='@') {
			lexer.Match('@')
			lexer.Match(TokenTypes_ID)
			secToken := lexer.GetNextToken()
			inReplyTo.SetCallId(token.GetTokenValue() + "@" +
				secToken.GetTokenValue())
		} else {
			inReplyTo.SetCallId(token.GetTokenValue())
		}

		lexer.SPorHT()

		inReplyToList.PushBack(inReplyTo)

		for ch, _ = lexer.LookAheadK(0); ch == ','; ch, _ = lexer.LookAheadK(0) {
			//while (lexer.lookAhead(0) == ',') {
			lexer.Match(',')
			lexer.SPorHT()

			inReplyTo = header.NewInReplyTo()

			lexer.Match(TokenTypes_ID)
			token = lexer.GetNextToken()
			if ch, _ = lexer.LookAheadK(0); ch == '@' {
				//if (lexer.lookAhead(0)=='@') {
				lexer.Match('@')
				lexer.Match(TokenTypes_ID)
				secToken := lexer.GetNextToken()
				inReplyTo.SetCallId(token.GetTokenValue() + "@" +
					secToken.GetTokenValue())
			} else {
				inReplyTo.SetCallId(token.GetTokenValue())
			}

			inReplyToList.PushBack(inReplyTo)
		}
	}

	return inReplyToList, nil
	// }
	// finally {
	//     if (debug) dbg_leave("InReplyToParser.parse");
	// }
}
