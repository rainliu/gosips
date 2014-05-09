package parser

import (
	"gosips/core"
	"gosips/sip/header"
	"strconv"
	"strings"
)

/** Parser for SubscriptionState header.
*
*@version  JAIN-SIP-1.1
*
*@author Olivier Deruelle <deruelle@nist.gov>
*@author M. Ranganathan <mranga@nist.gov>
*<a href="{@docRoot}/uncopyright.html">This code is in the public domain.</a>
*
* @version 1.0
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

	// if (debug) dbg_enter("SubscriptionStateParser.parse");

	subscriptionState := header.NewSubscriptionState()
	//  try {
	var ch byte
	lexer := this.GetLexer()
	this.HeaderName(TokenTypes_SUBSCRIPTION_STATE)

	subscriptionState.SetHeaderName(core.SIPHeaderNames_SUBSCRIPTION_STATE)

	// State:
	lexer.Match(TokenTypes_ID)
	token := lexer.GetNextToken()
	subscriptionState.SetState(token.GetTokenValue())

	for ch, _ = lexer.LookAheadK(0); ch == ';'; ch, _ = lexer.LookAheadK(0) {
		//while (lexer.lookAhead(0) == ';') {
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
			// try {
			expires, _ := strconv.Atoi(value)
			subscriptionState.SetExpires(expires)
			// } catch (NumberFormatException ex) {
			//     throw createParseException(ex.GetMessage());
			// } catch (InvalidArgumentException ex) {
			//     throw createParseException(ex.GetMessage());
			// }
		} else if strings.ToLower(value) == "retry-after" {
			lexer.Match('=')
			lexer.SPorHT()
			lexer.Match(TokenTypes_ID)
			token = lexer.GetNextToken()
			value = token.GetTokenValue()
			//try {
			retryAfter, _ := strconv.Atoi(value)
			subscriptionState.SetRetryAfter(retryAfter)
			// } catch (NumberFormatException ex) {
			//     throw createParseException(ex.GetMessage());
			// } catch (InvalidArgumentException ex) {
			//     throw createParseException(ex.GetMessage());
			// }
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
	// }
	// finally {
	//     if (debug) dbg_leave("SubscriptionStateParser.parse");
	// }

	return subscriptionState, nil
}
