package parser

import (
	"gosip/core"
	"gosip/header"
)

/** Parser for MinExpires header.
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
type MinExpiresParser struct {
	HeaderParserImpl
}

/** protected constructor.
 *@param text is the text of the header to parse
 */
func NewMinExpiresParser(minExpires string) *MinExpiresParser {
	this := &MinExpiresParser{}
	this.HeaderParserImpl.super(minExpires)
	return this
}

/** constructor.
 *@param lexer is the lexer passed in from the enclosing parser.
 */
func NewMinExpiresParserFromLexer(lexer core.Lexer) *MinExpiresParser {
	this := &MinExpiresParser{}
	this.HeaderParserImpl.superFromLexer(lexer)
	return this
}

/** parse the String message
 * @return SIPHeaderHeader (MinExpiresParser)
 * @throws SIPParseException if the message does not respect the spec.
 */
func (this *MinExpiresParser) Parse() (sh header.SIPHeaderHeader, ParseException error) {

	// if (debug) dbg_enter("MinExpiresParser.parse");
	minExpires := header.NewMinExpires()
	// try {
	lexer := this.GetLexer()
	this.HeaderName(TokenTypes_MIN_EXPIRES)

	minExpires.SetHeaderName(core.SIPHeaderNames_MIN_EXPIRES)

	number, _ := lexer.Number()
	//try{
	minExpires.SetExpires(number)
	// }
	// catch (InvalidArgumentException ex) {
	//         throw createParseException(ex.getMessage());
	// }
	lexer.SPorHT()

	lexer.Match('\n')

	return minExpires, nil
	// }
	// finally {
	//     if (debug) dbg_leave("MinExpiresParser.parse");
	// }
}
