package parser

import (
	"gosips/core"
	"gosips/sip/header"
)

/** SIPParser for a list of route headers.
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
func (this *RouteParser) Parse() (sh header.Header, ParseException error) {
	routeList := header.NewRouteList()

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
		} else {
			return nil, this.CreateParseException("unexpected char")
		}
	}

	return routeList, nil
}
