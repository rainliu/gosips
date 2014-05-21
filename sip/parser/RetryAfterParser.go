package parser

import (
	"gosips/core"
	"gosips/sip/header"
)

/** SIPParser for RetryAfter header.
 */
type RetryAfterParser struct {
	HeaderParser
}

/** Creates a new instance of RetryAfterParser
 * @param retryAfter the header to parse
 */
func NewRetryAfterParser(retryAfter string) *RetryAfterParser {
	this := &RetryAfterParser{}
	this.HeaderParser.super(retryAfter)
	return this
}

/** Constructor
 * @param lexer the lexer to use to parse the header
 */
func NewRetryAfterParserFromLexer(lexer core.Lexer) *RetryAfterParser {
	this := &RetryAfterParser{}
	this.HeaderParser.superFromLexer(lexer)
	return this
}

/** parse the String message
 * @return SIPHeader (RetryAfter object)
 * @throws SIPParseException if the message does not respect the spec.
 */
func (this *RetryAfterParser) Parse() (sh header.Header, ParseException error) {
	retryAfter := header.NewRetryAfter()

	var ch byte
	lexer := this.GetLexer()
	this.HeaderName(TokenTypes_RETRY_AFTER)

	// mandatory delatseconds:
	var ds int
	if ds, ParseException = lexer.Number(); ParseException != nil {
		return nil, ParseException
	}
	retryAfter.SetRetryAfter(ds)

	lexer.SPorHT()
	if ch, _ = lexer.LookAheadK(0); ch == '(' {
		comment, _ := lexer.Comment()
		retryAfter.SetComment(comment)
	}
	lexer.SPorHT()

	for ch, _ = lexer.LookAheadK(0); ch == ';'; ch, _ = lexer.LookAheadK(0) {
		lexer.Match(';')
		lexer.SPorHT()
		lexer.Match(TokenTypes_ID)
		token := lexer.GetNextToken()
		value := token.GetTokenValue()
		if value == "duration" {
			lexer.Match('=')
			lexer.SPorHT()

			var duration int
			if duration, ParseException = lexer.Number(); ParseException != nil {
				return nil, ParseException
			}
			retryAfter.SetDuration(duration)
		} else {
			lexer.SPorHT()
			lexer.Match('=')
			lexer.SPorHT()
			lexer.Match(TokenTypes_ID)
			secondToken := lexer.GetNextToken()
			secondValue := secondToken.GetTokenValue()
			retryAfter.SetParameter(value, secondValue)
		}
		lexer.SPorHT()
	}

	return retryAfter, nil
}
