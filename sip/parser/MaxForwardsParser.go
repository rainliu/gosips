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
	//if (debug) dbg_enter("MaxForwardsParser.enter");
	//    try {
	contentLength := header.NewMaxForwards()
	this.HeaderName(TokenTypes_MAX_FORWARDS)
	lexer := this.GetLexer()
	number, _ := lexer.Number()
	if err := contentLength.SetMaxForwards(number); err != nil {
		return nil, err
	}
	lexer.SPorHT()
	lexer.Match('\n')
	return contentLength, nil
	/* } catch (InvalidArgumentException ex) {
			   throw createParseException(ex.getMessage());
	              } catch (NumberFormatException ex) {
			   throw createParseException(ex.getMessage());
	              }  finally {
				if (debug) dbg_leave("MaxForwardsParser.leave");
		      }*/
}
