package parser

import (
	"gosips/core"
	"gosips/sip/header"
)

/** Parser for RSeq header.
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
type RSeqParser struct {
	HeaderParser
}

/** Creates a new instance of RSeqParser
 *@param rseq the header to parse
 */
func NewRSeqParser(rseq string) *RSeqParser {
	this := &RSeqParser{}
	this.HeaderParser.super(rseq)
	return this
}

/** Constructor
 * param lexer the lexer to use to parse the header
 */
func NewRSeqParserFromLexer(lexer core.Lexer) *RSeqParser {
	this := &RSeqParser{}
	this.HeaderParser.superFromLexer(lexer)
	return this
}

/** parse the String message
 * @return SIPHeader ( RSeq object)
 * @throws SIPParseException if the message does not respect the spec.
 */
func (this *RSeqParser) Parse() (sh header.Header, ParseException error) {

	//if (debug) dbg_enter("RSeqParser.parse");
	rseq := header.NewRSeq()
	//try {
	lexer := this.GetLexer()
	this.HeaderName(TokenTypes_RSEQ)

	rseq.SetHeaderName(core.SIPHeaderNames_RSEQ)

	number, _ := lexer.Number()
	//try{
	rseq.SetSequenceNumber(number)
	// }
	// catch (InvalidArgumentException ex) {
	//         throw createParseException(ex.getMessage());
	// }
	lexer.SPorHT()

	lexer.Match('\n')

	return rseq, nil
	// }
	// finally {
	//     if (debug) dbg_leave("RSeqParser.parse");
	// }
}
