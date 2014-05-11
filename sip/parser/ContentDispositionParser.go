package parser

import (
	"gosips/core"
	"gosips/sip/header"
)

/** SIPParser for ContentLanguage header.
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
type ContentDispositionParser struct {
	ParametersParser
}

/**
 * Creates a new instance of ContentDispositionParser
 * @param contentDisposition the header to parse
 */
func NewContentDispositionParser(contentDisposition string) *ContentDispositionParser {
	this := &ContentDispositionParser{}
	this.ParametersParser.super(contentDisposition)
	return this
}

/** Constructor
 * @param lexer the lexer to use to parse the header
 */
func NewContentDispositionParserFromLexer(lexer core.Lexer) *ContentDispositionParser {
	this := &ContentDispositionParser{}
	this.ParametersParser.superFromLexer(lexer)
	return this
}

/** parse the ContentDispositionHeader String header
 * @return Header (ContentDispositionList object)
 * @throws SIPParseException if the message does not respect the spec.
 */
func (this *ContentDispositionParser) Parse() (sh header.Header, ParseException error) {

	//if (debug) dbg_enter("ContentDispositionParser.parse");

	//try {
	//var ch byte
	//try {
	lexer := this.GetLexer()
	this.HeaderName(TokenTypes_CONTENT_DISPOSITION)

	cd := header.NewContentDisposition()
	cd.SetHeaderName(core.SIPHeaderNames_CONTENT_DISPOSITION)

	lexer.SPorHT()
	lexer.Match(TokenTypes_ID)

	token := lexer.GetNextToken()
	//println(token.GetTokenValue())
	cd.SetDispositionType(token.GetTokenValue())
	lexer.SPorHT()
	this.ParametersParser.Parse(cd)

	lexer.SPorHT()
	lexer.Match('\n')

	return cd, nil
	// } catch (ParseException ex ) {
	//     throw createParseException(ex.getMessage());
	// } finally {
	//     if (debug) dbg_leave("ContentDispositionParser.parse");
	// }
}
