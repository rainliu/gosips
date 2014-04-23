package parser

import (
	"gosip/core"
	"gosip/header"
)

/** Parser for ContentLanguage header.
*
*@version  JAIN-SIP-1.1
*
*@author Olivier Deruelle <deruelle@nist.gov>  <br/>
*@author M. Ranganathan <mranga@nist.gov>  <br/>
*
*<a href="{@docRoot}/uncopyright.html">This code is in the public domain.</a>
*
* @version 1.0
 */
type ContentEncodingParser struct {
	HeaderParserImpl
}

/**
 * Creates a new instance of ContentEncodingParser
 * @param contentEncoding the header to parse
 */
func NewContentEncodingParser(contentEncoding string) *ContentEncodingParser {
	this := &ContentEncodingParser{}
	this.HeaderParserImpl.super(contentEncoding)
	return this
}

/** Constructor
 * @param lexer the lexer to use to parse the header
 */
func NewContentEncodingParserFromLexer(lexer core.Lexer) *ContentEncodingParser {
	this := &ContentEncodingParser{}
	this.HeaderParserImpl.superFromLexer(lexer)
	return this
}

/** parse the ContentEncodingHeader String header
 * @return SIPHeaderHeader (ContentEncodingList object)
 * @throws SIPParseException if the message does not respect the spec.
 */
func (this *ContentEncodingParser) Parse() (sh header.SIPHeaderHeader, ParseException error) {

	//if (debug) dbg_enter("ContentEncodingParser.parse");
	contentEncodingList := header.NewContentEncodingList()

	var ch byte
	//try {
	lexer := this.GetLexer()
	this.HeaderName(TokenTypes_CONTENT_ENCODING)

	for ch, _ = lexer.LookAheadK(0); ch != '\n'; ch, _ = lexer.LookAheadK(0) {
		//while (lexer.lookAhead(0) != '\n') {
		cl := header.NewContentEncoding()
		cl.SetHeaderName(core.SIPHeaderNames_CONTENT_ENCODING)

		lexer.SPorHT()
		lexer.Match(TokenTypes_ID)

		token := lexer.GetNextToken()
		cl.SetEncoding(token.GetTokenValue())

		lexer.SPorHT()
		contentEncodingList.PushBack(cl)

		for ch, _ = lexer.LookAheadK(0); ch == ','; ch, _ = lexer.LookAheadK(0) {
			//while (lexer.lookAhead(0) == ',') {
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
	// } catch (ParseException ex ) {
	//     throw createParseException(ex.getMessage());
	// } finally {
	//     if (debug) dbg_leave("ContentEncodingParser.parse");
	// }
}
