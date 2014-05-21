package parser

import (
	"gosips/core"
	"gosips/sip/header"
)

/** SIPParser for Reason header.
 */
type ReasonParser struct {
	ParametersParser
}

/** Creates a new instance of ReasonParser
 * @param reason the header to parse
 */
func NewReasonParser(reason string) *ReasonParser {
	this := &ReasonParser{}
	this.ParametersParser.super(reason)
	return this
}

/** Constructor
 * @param lexer the lexer to use to parse the header
 */
func NewReasonParserFromLexer(lexer core.Lexer) *ReasonParser {
	this := &ReasonParser{}
	this.ParametersParser.superFromLexer(lexer)
	return this
}

/** parse the String message
 * @return SIPHeader (ReasonParserList object)
 * @throws SIPParseException if the message does not respect the spec.
 */
func (this *ReasonParser) Parse() (sh header.Header, ParseException error) {
	reasonList := header.NewReasonList()

	var ch byte
	lexer := this.GetLexer()
	this.HeaderName(TokenTypes_REASON)
	lexer.SPorHT()
	for ch, _ = lexer.LookAheadK(0); ch != '\n'; ch, _ = lexer.LookAheadK(0) {
		reason := header.NewReason()
		lexer.Match(TokenTypes_ID)
		token := lexer.GetNextToken()
		value := token.GetTokenValue()

		reason.SetProtocol(value)
		if ParseException = this.ParametersParser.Parse(reason); ParseException != nil {
			return nil, ParseException
		}
		reasonList.PushBack(reason)
		if ch, _ = lexer.LookAheadK(0); ch == ',' {
			lexer.Match(',')
			lexer.SPorHT()
		} else {
			lexer.SPorHT()
		}

	}

	return reasonList, nil
}
