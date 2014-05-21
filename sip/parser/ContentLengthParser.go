package parser

import (
	"gosips/core"
	"gosips/sip/header"
)

/** SIPParser for Content-Length Header.
 */
type ContentLengthParser struct {
	HeaderParser
}

func NewContentLengthParser(contentLength string) *ContentLengthParser {
	this := &ContentLengthParser{}
	this.HeaderParser.super(contentLength)
	return this
}

func NewContentLengthParserFromLexer(lexer core.Lexer) *ContentLengthParser {
	this := &ContentLengthParser{}
	this.HeaderParser.superFromLexer(lexer)
	return this
}

func (this *ContentLengthParser) Parse() (sh header.Header, ParseException error) {
	var number int
	contentLength := header.NewContentLength()
	lexer := this.GetLexer()
	this.HeaderName(TokenTypes_CONTENT_LENGTH)
	if number, ParseException = lexer.Number(); ParseException != nil {
		return nil, ParseException
	}
	if ParseException = contentLength.SetContentLength(number); ParseException != nil {
		return nil, ParseException
	}
	lexer.SPorHT()
	lexer.Match('\n')
	return contentLength, nil
}
