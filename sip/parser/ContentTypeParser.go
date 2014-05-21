package parser

import (
	"gosips/core"
	"gosips/sip/header"
)

/** SIPParser for content type header.
 */
type ContentTypeParser struct {
	ParametersParser
}

func NewContentTypeParser(contentType string) *ContentTypeParser {
	this := &ContentTypeParser{}
	this.ParametersParser.super(contentType)
	return this
}

func NewContentTypeParserFromLexer(lexer core.Lexer) *ContentTypeParser {
	this := &ContentTypeParser{}
	this.ParametersParser.superFromLexer(lexer)
	return this
}

func (this *ContentTypeParser) Parse() (sh header.Header, ParseException error) {
	contentType := header.NewContentType()

	lexer := this.GetLexer()
	this.HeaderName(TokenTypes_CONTENT_TYPE)

	// The type:
	lexer.Match(TokenTypes_ID)
	t := lexer.GetNextToken()
	lexer.SPorHT()
	contentType.SetContentType(t.GetTokenValue())

	// The sub-type:
	lexer.Match('/')
	lexer.Match(TokenTypes_ID)
	subType := lexer.GetNextToken()
	lexer.SPorHT()
	contentType.SetContentSubType(subType.GetTokenValue())
	if ParseException = this.ParametersParser.Parse(contentType); ParseException != nil {
		return nil, ParseException
	}
	lexer.Match('\n')

	return contentType, nil
}
