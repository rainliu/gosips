package parser

import (
	"gosips/core"
	"gosips/sip/header"
)

/** SIPParser for a list of route headers.
 */
type RecordRouteParser struct {
	AddressParametersParser
}

/** Constructor
 * @param String recordRoute message to parse to set
 */
func NewRecordRouteParser(recordRoute string) *RecordRouteParser {
	this := &RecordRouteParser{}
	this.AddressParametersParser.super(recordRoute)
	return this
}

func NewRecordRouteParserFromLexer(lexer core.Lexer) *RecordRouteParser {
	this := &RecordRouteParser{}
	this.AddressParametersParser.superFromLexer(lexer)
	return this
}

/** parse the String message and generate the RecordRoute List Object
 * @return SIPHeader the RecordRoute List object
 * @throws ParseException if errors occur during the parsing
 */
func (this *RecordRouteParser) Parse() (sh header.Header, ParseException error) {
	recordRouteList := header.NewRecordRouteList()

	var ch byte
	lexer := this.GetLexer()
	lexer.Match(TokenTypes_RECORD_ROUTE)
	lexer.SPorHT()
	lexer.Match(':')
	lexer.SPorHT()
	for {
		recordRoute := header.NewRecordRoute()
		this.AddressParametersParser.Parse(recordRoute)
		recordRouteList.PushBack(recordRoute)
		lexer.SPorHT()
		if ch, _ = lexer.LookAheadK(0); ch == ',' {
			lexer.Match(',')
			lexer.SPorHT()
		} else if ch, _ = lexer.LookAheadK(0); ch == '\n' {
			break
		} else {
			return nil, this.CreateParseException("unexpected char")
		}
	}

	return recordRouteList, nil
}
