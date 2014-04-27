package parser

import (
	"gosip/core"
	"gosip/header"
	"strconv"
)

/** Parser for TimeStamp header.
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
type TimeStampParser struct {
	HeaderParser
}

/** Creates a new instance of TimeStampParser
 * @param timeStamp the header to parse
 */
func NewTimeStampParser(timeStamp string) *TimeStampParser {
	this := &TimeStampParser{}
	this.HeaderParser.super(timeStamp)
	return this
}

/** Constructor
 * @param lexer the lexer to use to parse the header
 */
func NewTimeStampParserFromLexer(lexer core.Lexer) *TimeStampParser {
	this := &TimeStampParser{}
	this.HeaderParser.superFromLexer(lexer)
	return this
}

/** parse the String message
 * @return SIPHeader (TimeStamp object)
 * @throws SIPParseException if the message does not respect the spec.
 */
func (this *TimeStampParser) Parse() (sh header.ISIPHeader, ParseException error) {

	//if (debug) dbg_enter("TimeStampParser.parse");
	timeStamp := header.NewTimeStamp()
	//try {
	var ch byte
	lexer := this.GetLexer()
	this.HeaderName(TokenTypes_TIMESTAMP)

	timeStamp.SetHeaderName(core.SIPHeaderNames_TIMESTAMP)

	lexer.SPorHT()
	firstNumber, _ := lexer.Number()

	//try {
	var ts float64
	if ch, _ = lexer.LookAheadK(0); ch == '.' {
		lexer.Match('.')
		secondNumber, _ := lexer.Number()

		s := strconv.Itoa(firstNumber) + "." + strconv.Itoa(secondNumber)
		ts, _ = strconv.ParseFloat(s, 32)
	} else {
		ts = float64(firstNumber)
	}

	timeStamp.SetTimeStamp(float32(ts))
	// }
	// catch (NumberFormatException ex) {
	//      throw createParseException(ex.getMessage());
	// } catch (InvalidArgumentException ex) {
	//     throw createParseException(ex.getMessage());
	// }

	lexer.SPorHT()
	if ch, _ = lexer.LookAheadK(0); ch != '\n' {
		firstNumber, _ = lexer.Number()

		//try {
		//float ts;

		if ch, _ = lexer.LookAheadK(0); ch == '.' {
			lexer.Match('.')
			secondNumber, _ := lexer.Number()

			s := strconv.Itoa(firstNumber) + "." + strconv.Itoa(secondNumber)
			ts, _ = strconv.ParseFloat(s, 32)
		} else {
			ts = float64(firstNumber)
		}

		timeStamp.SetDelay(float32(ts))
		// }
		// catch (NumberFormatException ex) {
		//     throw createParseException(ex.getMessage());
		// } catch (InvalidArgumentException ex) {
		//     throw createParseException(ex.getMessage());
		// }
	}

	// }
	// finally {
	//     if (debug) dbg_leave("TimeStampParser.parse");
	// }

	return timeStamp, nil
}
