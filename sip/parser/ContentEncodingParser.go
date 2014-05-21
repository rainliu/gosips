package parser

import (
	"gosips/core"
	"gosips/sip/header"
)

/** SIPParser for ContentLanguage header.
 */
type ContentEncodingParser struct {
	HeaderParser
}

/**
 * Creates a new instance of ContentEncodingParser
 * @param contentEncoding the header to parse
 */
func NewContentEncodingParser(contentEncoding string) *ContentEncodingParser {
	this := &ContentEncodingParser{}
	this.HeaderParser.super(contentEncoding)
	return this
}

/** Constructor
 * @param lexer the lexer to use to parse the header
 */
func NewContentEncodingParserFromLexer(lexer core.Lexer) *ContentEncodingParser {
	this := &ContentEncodingParser{}
	this.HeaderParser.superFromLexer(lexer)
	return this
}

/** parse the ContentEncodingHeader String header
 * @return Header (ContentEncodingList object)
 * @throws SIPParseException if the message does not respect the spec.
 */
func (this *ContentEncodingParser) Parse() (sh header.Header, ParseException error) {
	contentEncodingList := header.NewContentEncodingList()

	var ch byte
	lexer := this.GetLexer()
	this.HeaderName(TokenTypes_CONTENT_ENCODING)

	for ch, _ = lexer.LookAheadK(0); ch != '\n'; ch, _ = lexer.LookAheadK(0) {
		cl := header.NewContentEncoding()
		cl.SetHeaderName(core.SIPHeaderNames_CONTENT_ENCODING)

		lexer.SPorHT()
		lexer.Match(TokenTypes_ID)

		token := lexer.GetNextToken()
		cl.SetEncoding(token.GetTokenValue())

		lexer.SPorHT()
		contentEncodingList.PushBack(cl)

		for ch, _ = lexer.LookAheadK(0); ch == ','; ch, _ = lexer.LookAheadK(0) {
			cl = header.NewContentEncoding()
			lexer.Match(',')
			lexer.SPorHT()
			lexer.Match(TokenTypes_ID)
			lexer.SPorHT()
			token = lexer.GetNextToken()
			cl.SetEncoding(token.GetTokenValue())
			lexer.SPorHT()
			contentEncodingList.PushBack(cl)
		}
	}

	return contentEncodingList, nil
}
