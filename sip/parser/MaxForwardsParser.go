package parser

import (
	"gosips/core"
	"gosips/sip/header"
)

/** SIPParser for Max Forwards Header.*/
type MaxForwardsParser struct {
	HeaderParser
}

func NewMaxForwardsParser(contentLength string) *MaxForwardsParser {
	this := &MaxForwardsParser{}
	this.HeaderParser.super(contentLength)
	return this
}

func NewMaxForwardsParserFromLexer(lexer core.Lexer) *MaxForwardsParser {
	this := &MaxForwardsParser{}
	this.HeaderParser.superFromLexer(lexer)
	return this
}

func (this *MaxForwardsParser) super(hname string) {
	this.HeaderParser.super(hname)
}

func (this *MaxForwardsParser) Parse() (sh header.Header, ParseException error) {
	contentLength := header.NewMaxForwards()
	this.HeaderName(TokenTypes_MAX_FORWARDS)
	lexer := this.GetLexer()

	var number int
	if number, ParseException = lexer.Number(); ParseException != nil {
		return nil, ParseException
	}
	if ParseException = contentLength.SetMaxForwards(number); ParseException != nil {
		return nil, ParseException
	}
	lexer.SPorHT()
	lexer.Match('\n')

	return contentLength, nil
}
