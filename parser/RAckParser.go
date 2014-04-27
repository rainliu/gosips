package parser

import (
	"gosip/core"
	"gosip/header"
)

/** Parser for RAck header.
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
type RAckParser struct {
	HeaderParser
}

/** Creates a new instance of RAckParser
 *@param rack the header to parse
 */
func NewRAckParser(rack string) *RAckParser {
	this := &RAckParser{}
	this.HeaderParser.super(rack)
	return this
}

/** Constructor
 * @param lexer the lexer to use to parse the header
 */
func NewRAckParserFromLexer(lexer core.Lexer) *RAckParser {
	this := &RAckParser{}
	this.HeaderParser.superFromLexer(lexer)
	return this
}

/** parse the String message
 * @return SIPHeader (RAck object)
 * @throws SIPParseException if the message does not respect the spec.
 */
func (this *RAckParser) Parse() (sh header.ISIPHeader, ParseException error) {

	//if (debug) dbg_enter("RAckParser.parse");
	rack := header.NewRAck()
	// try {
	lexer := this.GetLexer()
	this.HeaderName(TokenTypes_RACK)

	rack.SetHeaderName(core.SIPHeaderNames_RACK)

	//try{
	number, _ := lexer.Number()
	rack.SetRSeqNumber(number)
	lexer.SPorHT()
	number, _ = lexer.Number()
	rack.SetCSeqNumber(number)
	lexer.SPorHT()
	lexer.Match(TokenTypes_ID)
	token := lexer.GetNextToken()
	rack.SetMethod(token.GetTokenValue())

	// }
	// catch (InvalidArgumentException ex) {
	//         throw createParseException(ex.getMessage());
	// }
	lexer.SPorHT()
	lexer.Match('\n')

	return rack, nil
	// }
	// finally {
	//     if (debug) dbg_leave("RAckParser.parse");
	// }
}
