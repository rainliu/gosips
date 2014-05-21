package parser

import (
	"gosips/core"
	"gosips/sip/header"
	"strconv"
	"strings"
)

/** SIPParser for SubscriptionState header.
 */
type SubscriptionStateParser struct {
	HeaderParser
}

/** Creates a new instance of SubscriptionStateParser
 * @param subscriptionState the header to parse
 */
func NewSubscriptionStateParser(subscriptionState string) *SubscriptionStateParser {
	this := &SubscriptionStateParser{}
	this.HeaderParser.super(subscriptionState)
	return this
}

/** Constructor
 * @param lexer the lexer to use to parse the header
 */
func NewSubscriptionStateParserFromLexer(lexer core.Lexer) *SubscriptionStateParser {
	this := &SubscriptionStateParser{}
	this.HeaderParser.superFromLexer(lexer)
	return this
}

/** parse the String message
 * @return SIPHeader (SubscriptionState  object)
 * @throws SIPParseException if the message does not respect the spec.
 */
func (this *SubscriptionStateParser) Parse() (sh header.Header, ParseException error) {
	subscriptionState := header.NewSubscriptionState()

	var ch byte
	lexer := this.GetLexer()
	this.HeaderName(TokenTypes_SUBSCRIPTION_STATE)

	subscriptionState.SetHeaderName(core.SIPHeaderNames_SUBSCRIPTION_STATE)

	// State:
	lexer.Match(TokenTypes_ID)
	token := lexer.GetNextToken()
	subscriptionState.SetState(token.GetTokenValue())

	for ch, _ = lexer.LookAheadK(0); ch == ';'; ch, _ = lexer.LookAheadK(0) {
		lexer.Match(';')
		lexer.SPorHT()
		lexer.Match(TokenTypes_ID)
		token = lexer.GetNextToken()
		value := token.GetTokenValue()
		if strings.ToLower(value) == "reason" {
			lexer.Match('=')
			lexer.SPorHT()
			lexer.Match(TokenTypes_ID)
			token = lexer.GetNextToken()
			value = token.GetTokenValue()
			subscriptionState.SetReasonCode(value)
		} else if strings.ToLower(value) == "expires" {
			lexer.Match('=')
			lexer.SPorHT()
			lexer.Match(TokenTypes_ID)
			token = lexer.GetNextToken()
			value = token.GetTokenValue()

			var expires int
			if expires, ParseException = strconv.Atoi(value); ParseException != nil {
				return nil, ParseException
			}
			subscriptionState.SetExpires(expires)
		} else if strings.ToLower(value) == "retry-after" {
			lexer.Match('=')
			lexer.SPorHT()
			lexer.Match(TokenTypes_ID)
			token = lexer.GetNextToken()
			value = token.GetTokenValue()

			var retryAfter int
			if retryAfter, ParseException = strconv.Atoi(value); ParseException != nil {
				return nil, ParseException
			}
			subscriptionState.SetRetryAfter(retryAfter)
		} else {
			lexer.Match('=')
			lexer.SPorHT()
			lexer.Match(TokenTypes_ID)
			secondToken := lexer.GetNextToken()
			secondValue := secondToken.GetTokenValue()
			subscriptionState.SetParameter(value, secondValue)
		}

		lexer.SPorHT()
	}

	return subscriptionState, nil
}
