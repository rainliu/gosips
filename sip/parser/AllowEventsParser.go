package parser

import (
	"gosips/core"
	"gosips/sip/header"
)

/** SIPParser for AllowEvents header.
 */
type AllowEventsParser struct {
	HeaderParser
}

/**
 * Creates a new instance of AllowEventsParser
 * @param allowEvents the header to parse
 */
func NewAllowEventsParser(allowEvents string) *AllowEventsParser {
	this := &AllowEventsParser{}
	this.HeaderParser.super(allowEvents)
	return this
}

/** Constructor
 * @param lexer the lexer to use to parse the header
 */
func NewAllowEventsParserFromLexer(lexer core.Lexer) *AllowEventsParser {
	this := &AllowEventsParser{}
	this.HeaderParser.superFromLexer(lexer)
	return this
}

/** parse the AllowEvents String header
 * @return Header (AllowEventsList object)
 * @throws SIPParseException if the message does not respect the spec.
 */
func (this *AllowEventsParser) Parse() (sh header.Header, ParseException error) {
	allowEventsList := header.NewAllowEventsList()

	var ch byte

	lexer := this.GetLexer()
	this.HeaderName(TokenTypes_ALLOW_EVENTS)

	allowEvents := header.NewAllowEvents()
	allowEvents.SetHeaderName(core.SIPHeaderNames_ALLOW_EVENTS)

	lexer.SPorHT()
	lexer.Match(TokenTypes_ID)
	token := lexer.GetNextToken()
	allowEvents.SetEventType(token.GetTokenValue())

	allowEventsList.PushBack(allowEvents)
	lexer.SPorHT()

	for ch, _ = lexer.LookAheadK(0); ch == ','; ch, _ = lexer.LookAheadK(0) {
		lexer.Match(',')
		lexer.SPorHT()

		allowEvents = header.NewAllowEvents()
		lexer.Match(TokenTypes_ID)
		token = lexer.GetNextToken()
		allowEvents.SetEventType(token.GetTokenValue())

		allowEventsList.PushBack(allowEvents)
		lexer.SPorHT()
	}
	lexer.SPorHT()
	lexer.Match('\n')

	return allowEventsList, nil
}
