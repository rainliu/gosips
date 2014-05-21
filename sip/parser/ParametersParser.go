package parser

import (
	"gosips/core"
	"gosips/sip/header"
)

/** parameters parser header.
 */
type ParametersParser struct {
	HeaderParser
}

func NewParametersParserFromLexer(lexer core.Lexer) *ParametersParser {
	this := &ParametersParser{}
	this.HeaderParser.superFromLexer(lexer)
	return this
}

func NewParametersParser(buffer string) *ParametersParser {
	this := &ParametersParser{}
	this.HeaderParser.super(buffer)
	return this
}

func (this *ParametersParser) super(buffer string) {
	this.HeaderParser.super(buffer)
}

func (this *ParametersParser) superFromLexer(lexer core.Lexer) {
	this.HeaderParser.superFromLexer(lexer)
}

func (this *ParametersParser) Parse(parametersHeader header.ParametersHeader) (ParseException error) {
	var ch byte
	var nv *core.NameValue

	lexer := this.GetLexer()

	lexer.SPorHT()
	if ch, ParseException = lexer.LookAheadK(0); ParseException != nil {
		return ParseException
	}

	for ch == ';' {
		lexer.ConsumeK(1)
		// eat white space
		lexer.SPorHT()
		//println(lexer.GetRest())
		if nv, ParseException = this.NameValue('='); ParseException != nil {
			return ParseException
		}

		//println(lexer.GetRest())
		if nv.IsValueQuoted() {
			parametersHeader.SetParameter(nv.GetName(), "\""+nv.GetValue().(string)+"\"")
		} else {
			parametersHeader.SetParameter(nv.GetName(), nv.GetValue().(string))
		}
		// eat white space
		lexer.SPorHT()

		if ch, ParseException = lexer.LookAheadK(0); ParseException != nil {
			return ParseException
		}
	}
	return nil
}
