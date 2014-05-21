package parser

import (
	"gosips/core"
	"gosips/sip/header"
)

/** SIPParser for RSeq header.
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
	rseq := header.NewRSeq()

	lexer := this.GetLexer()
	this.HeaderName(TokenTypes_RSEQ)

	rseq.SetHeaderName(core.SIPHeaderNames_RSEQ)

	var number int
	if number, ParseException = lexer.Number(); ParseException != nil {
		return nil, ParseException
	}
	rseq.SetSequenceNumber(number)

	lexer.SPorHT()
	lexer.Match('\n')

	return rseq, nil
}
