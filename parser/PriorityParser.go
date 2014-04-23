package parser

import (
	"gosip/core"
	"gosip/header"
)

/** Parser for Priority header.
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
type PriorityParser struct {
	HeaderParserImpl
}

/** Creates a new instance of PriorityParser
 * @param priority the header to parse
 */
func NewPriorityParser(priority string) *PriorityParser {
	this := &PriorityParser{}
	this.HeaderParserImpl.super(priority)
	return this
}

/** Constructor
 * @param lexer the lexer to use to parse the header
 */
func NewPriorityParserFromLexer(lexer core.Lexer) *PriorityParser {
	this := &PriorityParser{}
	this.HeaderParserImpl.superFromLexer(lexer)
	return this
}

/** parse the String header
 * @return SIPHeader (Priority object)
 * @throws SIPParseException if the message does not respect the spec.
 */
func (this *PriorityParser) Parse() (sh header.SIPHeaderHeader, ParseException error) {

	//if (debug) dbg_enter("PriorityParser.parse");
	priority := header.NewPriority()
	// try {
	lexer := this.GetLexer()
	this.HeaderName(TokenTypes_PRIORITY)

	priority.SetHeaderName(core.SIPHeaderNames_PRIORITY)

	lexer.SPorHT()
	lexer.Match(TokenTypes_ID)
	token := lexer.GetNextToken()

	priority.SetPriority(token.GetTokenValue())

	lexer.SPorHT()
	lexer.Match('\n')

	return priority, nil
	// }
	// finally {
	//     if (debug) dbg_leave("PriorityParser.parse");
	// }
}
