package parser

import (
	"gosip/core"
	"gosip/header"
)

/** Parser for RetryAfter header.
*
*@version  JAIN-SIP-1.1
*
*@author Olivier Deruelle <deruelle@nist.gov>
*@author M. Ranganathan <mranga@nist.gov>
*<a href="{@docRoot}/uncopyright.html">This code is in the public domain.</a>
*
* @version 1.0
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
func (this *RetryAfterParser) Parse() (sh header.ISIPHeader, ParseException error) {

	//if (debug) dbg_enter("RetryAfterParser.parse");

	retryAfter := header.NewRetryAfter()
	//try {
	var ch byte
	lexer := this.GetLexer()
	this.HeaderName(TokenTypes_RETRY_AFTER)

	// mandatory delatseconds:
	ds, _ := lexer.Number()
	//try {
	//int ds= Integer.parseInt(value);
	retryAfter.SetRetryAfter(ds)
	// } catch (NumberFormatException ex) {
	//     throw createParseException(ex.getMessage());
	// } catch (InvalidArgumentException ex) {
	//     throw createParseException(ex.getMessage());
	// }

	lexer.SPorHT()
	if ch, _ = lexer.LookAheadK(0); ch == '(' {
		comment, _ := lexer.Comment()
		retryAfter.SetComment(comment)
	}
	lexer.SPorHT()

	for ch, _ = lexer.LookAheadK(0); ch == ';'; ch, _ = lexer.LookAheadK(0) {
		//while (lexer.lookAhead(0) == ';') {
		lexer.Match(';')
		lexer.SPorHT()
		lexer.Match(TokenTypes_ID)
		token := lexer.GetNextToken()
		value := token.GetTokenValue()
		if value == "duration" {
			lexer.Match('=')
			lexer.SPorHT()
			duration, _ := lexer.Number()
			//try {
			//int duration= Integer.parseInt(value);
			retryAfter.SetDuration(duration)
			// } catch (NumberFormatException ex) {
			//     throw createParseException(ex.getMessage());
			// } catch (InvalidArgumentException ex) {
			//     throw createParseException(ex.getMessage());
			// }
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
	// }
	// finally {
	//     if (debug) dbg_leave("RetryAfterParser.parse");
	// }

	return retryAfter, nil
}
