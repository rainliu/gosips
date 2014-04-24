package parser

import (
	"gosip/core"
	"gosip/header"
)

/** Parser for AllowEvents header.
*
*@version  JAIN-SIP-1.1
*
*@author Olivier Deruelle <deruelle@nist.gov>
*@author M. Ranganathan <mranga@nist.gov>
*<a href="{@docRoot}/uncopyright.html">This code is in the public domain.</a>
*
* @version 1.0
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
 * @return SIPHeaderHeader (AllowEventsList object)
 * @throws SIPParseException if the message does not respect the spec.
 */
func (this *AllowEventsParser) Parse() (sh header.SIPHeaderHeader, ParseException error) {

	//if (debug) dbg_enter("AllowEventsParser.parse");
	allowEventsList := header.NewAllowEventsList()

	// try {
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

// finally {
//     if (debug) dbg_leave("AllowEventsParser.parse");
// }
//}
