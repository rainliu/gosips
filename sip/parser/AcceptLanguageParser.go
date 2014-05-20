package parser

import (
	"gosips/core"
	"gosips/sip/header"
	"strconv"
)

/**
* Accept Language body.
* <pre>
*
* Accept-Language = "Accept-Language" ":"
*                         1#( language-range [ ";" "q" "=" qvalue ] )
*       language-range  = ( ( 1*8ALPHA *( "-" 1*8ALPHA ) ) | "*" )
*
* HTTP RFC 2616 Section 14.4
* </pre>
*
*  Accept-Language: da, en-gb;q=0.8, en;q=0.7
*
* @see AcceptLanguageList
 */

/** SIPParser for Accept Language Headers.
 */

type AcceptLanguageParser struct {
	HeaderParser
}

/** Constructor
 * @param String AcceptLanguage message to parse
 */
func NewAcceptLanguageParser(acceptLanguage string) *AcceptLanguageParser {
	this := &AcceptLanguageParser{}
	this.HeaderParser.super(acceptLanguage)
	return this
}

/** Cosntructor
 * @param lexer to set
 */
func NewAcceptLanguageParserFromLexer(lexer core.Lexer) *AcceptLanguageParser {
	this := &AcceptLanguageParser{}
	this.HeaderParser.superFromLexer(lexer)
	return this
}

/** parse the String message
 * @return Header (AcceptLanguage object)
 * @throws ParseException if the message does not respect the spec.
 */
func (this *AcceptLanguageParser) Parse() (sh header.Header, ParseException error) {
	acceptLanguageList := header.NewAcceptLanguageList()

	var ch byte
	lexer := this.GetLexer()
	this.HeaderName(TokenTypes_ACCEPT_LANGUAGE)

	for ch, _ = lexer.LookAheadK(0); ch != '\n'; ch, _ = lexer.LookAheadK(0) {
		acceptLanguage := header.NewAcceptLanguage()
		acceptLanguage.SetHeaderName(core.SIPHeaderNames_ACCEPT_LANGUAGE)
		if ch, _ = lexer.LookAheadK(0); ch != ';' { // Content-Coding:
			lexer.Match(TokenTypes_ID)
			value := lexer.GetNextToken()
			acceptLanguage.SetLanguageRange(value.GetTokenValue())
		}

		for ch, _ = lexer.LookAheadK(0); ch == ';'; ch, _ = lexer.LookAheadK(0) {
			lexer.Match(';')
			lexer.SPorHT()
			lexer.Match('q')
			lexer.SPorHT()
			lexer.Match('=')
			lexer.SPorHT()
			lexer.Match(TokenTypes_ID)
			value := lexer.GetNextToken()

			var qv float64
			if qv, ParseException = strconv.ParseFloat(value.GetTokenValue(), 32); ParseException != nil {
				return nil, ParseException
			}

			if ParseException = acceptLanguage.SetQValue(float32(qv)); ParseException != nil {
				return nil, ParseException
			}

			lexer.SPorHT()
		}

		acceptLanguageList.PushBack(acceptLanguage)
		if ch, _ = lexer.LookAheadK(0); ch == ',' {
			lexer.Match(',')
			lexer.SPorHT()
		} else {
			lexer.SPorHT()
		}
	}

	return acceptLanguageList, nil
}
