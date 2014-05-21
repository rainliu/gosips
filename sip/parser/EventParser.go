package parser

import (
	"gosips/core"
	"gosips/sip/header"
)

/** SIPParser for Event header.
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
	lexer := this.GetLexer()
	this.HeaderName(TokenTypes_EVENT)
	lexer.SPorHT()

	event := header.NewEvent()
	lexer.Match(TokenTypes_ID)
	token := lexer.GetNextToken()
	value := token.GetTokenValue()

	event.SetEventType(value)
	if ParseException = this.ParametersParser.Parse(event); ParseException != nil {
		return nil, ParseException
	}

	lexer.SPorHT()
	lexer.Match('\n')

	return event, nil
}
