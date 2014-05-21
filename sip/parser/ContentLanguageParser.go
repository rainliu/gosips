package parser

import (
	"gosips/core"
	"gosips/sip/header"
)

/** SIPParser for ContentLanguage header.
 */
type ContentLanguageParser struct {
	HeaderParser
}

/**
 * Creates a new instance of ContentLanguageParser
 * @param contentLanguage the header to parse
 */
func NewContentLanguageParser(contentLanguage string) *ContentLanguageParser {
	this := &ContentLanguageParser{}
	this.HeaderParser.super(contentLanguage)
	return this
}

/** Constructor
 * @param lexer the lexer to use to parse the header
 */
func NewContentLanguageParserFromLexer(lexer core.Lexer) *ContentLanguageParser {
	this := &ContentLanguageParser{}
	this.HeaderParser.superFromLexer(lexer)
	return this
}

/** parse the ContentLanguageHeader String header
 * @return Header (ContentLanguageList object)
 * @throws SIPParseException if the message does not respect the spec.
 */
func (this *ContentLanguageParser) Parse() (sh header.Header, ParseException error) {
	contentLanguageList := header.NewContentLanguageList()

	var ch byte
	lexer := this.GetLexer()
	this.HeaderName(TokenTypes_CONTENT_LANGUAGE)

	for ch, _ = lexer.LookAheadK(0); ch != '\n'; ch, _ = lexer.LookAheadK(0) {
		cl := header.NewContentLanguage()
		cl.SetHeaderName(core.SIPHeaderNames_CONTENT_LANGUAGE)

		lexer.SPorHT()
		lexer.Match(TokenTypes_ID)

		token := lexer.GetNextToken()
		cl.SetContentLanguage(token.GetTokenValue())

		lexer.SPorHT()
		contentLanguageList.PushBack(cl)

		for ch, _ = lexer.LookAheadK(0); ch == ','; ch, _ = lexer.LookAheadK(0) {
			cl = header.NewContentLanguage()
			lexer.Match(',')
			lexer.SPorHT()
			lexer.Match(TokenTypes_ID)
			lexer.SPorHT()
			token = lexer.GetNextToken()
			cl.SetContentLanguage(token.GetTokenValue())

			lexer.SPorHT()
			contentLanguageList.PushBack(cl)
		}
	}

	return contentLanguageList, nil
}
