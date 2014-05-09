package parser

import (
	"gosips/core"
	"gosips/sip/header"
)

/** Parser for Event header.
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
type EventParser struct {
	ParametersParser
}

/**
 * Creates a new instance of EventParser
 * @param event the header to parse
 */
func NewEventParser(event string) *EventParser {
	this := &EventParser{}
	this.ParametersParser.super(event)
	return this
}

/** Cosntructor
 * @param lexer the lexer to use to parse the header
 */
func NewEventParserFromLexer(lexer core.Lexer) *EventParser {
	this := &EventParser{}
	this.ParametersParser.superFromLexer(lexer)
	return this
}

/** parse the String message
 * @return Header (Event object)
 * @throws SIPParseException if the message does not respect the spec.
 */
func (this *EventParser) Parse() (sh header.Header, ParseException error) {

	//if (debug) dbg_enter("EventParser.parse");

	//try {
	lexer := this.GetLexer()
	this.HeaderName(TokenTypes_EVENT)
	lexer.SPorHT()

	event := header.NewEvent()
	lexer.Match(TokenTypes_ID)
	token := lexer.GetNextToken()
	value := token.GetTokenValue()

	event.SetEventType(value)
	this.ParametersParser.Parse(event)

	lexer.SPorHT()
	lexer.Match('\n')

	return event, nil

	// } catch (ParseException ex ) {
	//      throw createParseException(ex.getMessage());
	// } finally {
	//     if (debug) dbg_leave("EventParser.parse");
	// }
}
