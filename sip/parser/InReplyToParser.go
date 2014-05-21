package parser

import (
	"gosips/core"
	"gosips/sip/header"
)

/** SIPParser for InReplyTo header.
 */
type InReplyToParser struct {
	HeaderParser
}

/** Creates a new instance of InReplyToParser
 * @param inReplyTo the header to parse
 */
func NewInReplyToParser(inReplyTo string) *InReplyToParser {
	this := &InReplyToParser{}
	this.HeaderParser.super(inReplyTo)
	return this
}

/** Constructor
 * @param lexer the lexer to use to parse the header
 */
func NewInReplyToParserFromLexer(lexer core.Lexer) *InReplyToParser {
	this := &InReplyToParser{}
	this.HeaderParser.superFromLexer(lexer)
	return this
}

/** parse the String message
 * @return Header (InReplyToList object)
 * @throws SIPParseException if the message does not respect the spec.
 */
func (this *InReplyToParser) Parse() (sh header.Header, ParseException error) {
	inReplyToList := header.NewInReplyToList()

	var ch byte
	lexer := this.GetLexer()
	this.HeaderName(TokenTypes_IN_REPLY_TO)

	for ch, _ = lexer.LookAheadK(0); ch != '\n'; ch, _ = lexer.LookAheadK(0) {
		inReplyTo := header.NewInReplyTo()
		inReplyTo.SetHeaderName(core.SIPHeaderNames_IN_REPLY_TO)

		lexer.Match(TokenTypes_ID)
		token := lexer.GetNextToken()
		if ch, _ = lexer.LookAheadK(0); ch == '@' {
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
			lexer.Match(',')
			lexer.SPorHT()

			inReplyTo = header.NewInReplyTo()

			lexer.Match(TokenTypes_ID)
			token = lexer.GetNextToken()
			if ch, _ = lexer.LookAheadK(0); ch == '@' {
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
}
