package parser

import (
	"gosips/core"
	"gosips/sip/header"
	"strconv"
)

/** SIPParser for TimeStamp header.
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
func (this *TimeStampParser) Parse() (sh header.Header, ParseException error) {
	timeStamp := header.NewTimeStamp()

	var ch byte
	var firstNumber, secondNumber int

	lexer := this.GetLexer()
	this.HeaderName(TokenTypes_TIMESTAMP)

	timeStamp.SetHeaderName(core.SIPHeaderNames_TIMESTAMP)
	lexer.SPorHT()

	if firstNumber, ParseException = lexer.Number(); ParseException != nil {
		return nil, ParseException
	}

	var ts float64
	if ch, _ = lexer.LookAheadK(0); ch == '.' {
		lexer.Match('.')

		if secondNumber, ParseException = lexer.Number(); ParseException != nil {
			return nil, ParseException
		}

		s := strconv.Itoa(firstNumber) + "." + strconv.Itoa(secondNumber)
		ts, _ = strconv.ParseFloat(s, 32)
	} else {
		ts = float64(firstNumber)
	}

	timeStamp.SetTimeStamp(float32(ts))

	lexer.SPorHT()
	if ch, _ = lexer.LookAheadK(0); ch != '\n' {
		if firstNumber, ParseException = lexer.Number(); ParseException != nil {
			return nil, ParseException
		}

		if ch, _ = lexer.LookAheadK(0); ch == '.' {
			lexer.Match('.')
			if secondNumber, ParseException = lexer.Number(); ParseException != nil {
				return nil, ParseException
			}

			s := strconv.Itoa(firstNumber) + "." + strconv.Itoa(secondNumber)
			ts, _ = strconv.ParseFloat(s, 32)
		} else {
			ts = float64(firstNumber)
		}

		timeStamp.SetDelay(float32(ts))
	}

	return timeStamp, nil
}
