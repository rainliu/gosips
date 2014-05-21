package parser

import (
	"gosips/core"
	"gosips/sip/header"
)

/** SIPParser for CSeq headers.
 */
type CSeqParser struct {
	HeaderParser
}

func NewCSeqParser(cseq string) *CSeqParser {
	this := &CSeqParser{}
	this.HeaderParser.super(cseq)
	return this
}

func (this *CSeqParser) super(cseq string) {
	this.HeaderParser.super(cseq)
}

func NewCSeqParserFromLexer(lexer core.Lexer) *CSeqParser {
	this := &CSeqParser{}
	this.HeaderParser.superFromLexer(lexer)
	return this
}

func (this *CSeqParser) Parse() (sh header.Header, ParseException error) {
	lexer := this.GetLexer()
	lexer.Match(TokenTypes_CSEQ)
	lexer.SPorHT()
	lexer.Match(':')
	lexer.SPorHT()

	var number int
	var method string
	if number, ParseException = lexer.Number(); ParseException != nil {
		return nil, ParseException
	}
	lexer.SPorHT()
	if method, ParseException = this.Method(); ParseException != nil {
		return nil, ParseException
	}
	//c.SetMethod(m);
	lexer.SPorHT()
	lexer.Match('\n')

	c := header.NewCSeq(number, method)

	return c, nil
}
