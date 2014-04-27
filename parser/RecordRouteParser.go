package parser

import (
	"gosip/core"
	"gosip/header"
)

/** Parser for a list of route headers.
*
*@version  JAIN-SIP-1.1
*
*@author Olivier Deruelle <deruelle@nist.gov>  <br/>
*@author M. Ranganathan <mranga@nist.gov>  <br/>
*<a href="{@docRoot}/uncopyright.html">This code is in the public domain.</a>
*
*@version 1.0
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
func (this *RecordRouteParser) Parse() (sh header.ISIPHeader, ParseException error) {
	recordRouteList := header.NewRecordRouteList()

	//if (debug) dbg_enter("RecordRouteParser.parse");

	// try {
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
		}
		//else throw createParseException("unexpected char");
	}
	return recordRouteList, nil
	//      } finally {
	// if (debug) dbg_leave("RecordRouteParser.parse");
	//      }

}
