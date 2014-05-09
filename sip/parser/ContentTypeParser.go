package parser

import (
	"gosips/core"
	"gosips/sip/header"
)

/** Parser for content type header.
*
*@version  JAIN-SIP-1.1
*
*@author M. Ranganathan <mranga@nist.gov>  <br/>
*
*<a href="{@docRoot}/uncopyright.html">This code is in the public domain.</a>
*
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
	//if (debug) dbg_enter("ContentTypeParser.parse");

	//try{
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
	this.ParametersParser.Parse(contentType)
	lexer.Match('\n')
	//           }  finally {
	// if (debug) dbg_leave("ContentTypeParser.parse");
	//    }
	return contentType, nil

}
