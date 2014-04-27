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

type RouteParser struct {
	AddressParametersParser
}

/** Constructor
 * @param String route message to parse to set
 */
func NewRouteParser(route string) *RouteParser {
	this := &RouteParser{}
	this.AddressParametersParser.super(route)
	return this
}

func NewRouteParserFromLexer(lexer core.Lexer) *RouteParser {
	this := &RouteParser{}
	this.AddressParametersParser.superFromLexer(lexer)
	return this
}

/** parse the String message and generate the Route List Object
 * @return SIPHeader the Route List object
 * @throws SIPParseException if errors occur during the parsing
 */
func (this *RouteParser) Parse() (sh header.ISIPHeader, ParseException error) {
	routeList := header.NewRouteList()
	//if (debug) dbg_enter("parse");

	//try {
	var ch byte
	lexer := this.GetLexer()
	lexer.Match(TokenTypes_ROUTE)
	lexer.SPorHT()
	lexer.Match(':')
	lexer.SPorHT()
	for {
		route := header.NewRoute()
		this.AddressParametersParser.Parse(route)
		routeList.PushBack(route)
		lexer.SPorHT()
		if ch, _ = lexer.LookAheadK(0); ch == ',' {
			lexer.Match(',')
			lexer.SPorHT()
		} else if ch, _ = lexer.LookAheadK(0); ch == '\n' {
			break
		} //else throw createParseException("unexpected char");
	}
	return routeList, nil
	//      } finally {
	// if (debug) dbg_leave("parse");
	//      }

}
