package parser

import (
	"gosips/core"
	"gosips/sip/header"
	"strings"
)

/** SIPParser for CALL ID header.
 */
type CallIDParser struct {
	HeaderParser
}

/** Creates new CallIDParser
 * @param String callID message to parse to set
 */
func NewCallIDParser(callId string) *CallIDParser {
	this := &CallIDParser{}

	this.HeaderParser.super(callId)

	return this
}

func (this *CallIDParser) super(callId string) {
	this.HeaderParser.super(callId)
}

/** Constructor
 * @param lexer to set
 */
func NewCallIDParserFromLexer(lexer core.Lexer) *CallIDParser {
	this := &CallIDParser{}

	this.HeaderParser.superFromLexer(lexer)

	return this
}

/** parse the String message
 * @return Header (CallID object)
 * @throws ParseException if the message does not respect the spec.
 */
func (this *CallIDParser) Parse() (sh header.Header, ParseException error) {
	lexer := this.GetLexer()
	lexer.Match(TokenTypes_CALL_ID)
	lexer.SPorHT()
	lexer.Match(':')
	lexer.SPorHT()
	rest := strings.TrimSpace(lexer.GetRest())

	callID, ParseException := header.NewCallID(rest)

	return callID, ParseException
}
