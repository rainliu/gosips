package parser

import (
	"gosip/core"
	"gosip/header"
)

/** Parser for Content-Length Header.
*@version  JAIN-SIP-1.1
*
*@author Olivier Deruelle  <br/>
*@author M. Ranganathan <mranga@nist.gov>  <br/>
*
*<a href="{@docRoot}/uncopyright.html">This code is in the public domain.</a>
*
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

func (this *ContentLengthParser) Parse() (sh header.SIPHeaderHeader, ParseException error) {
	// if (debug) dbg_enter("ContentLengthParser.enter");
	//      try {
	contentLength := header.NewContentLength()
	lexer := this.GetLexer()
	this.HeaderName(TokenTypes_CONTENT_LENGTH)
	number, _ := lexer.Number()
	contentLength.SetContentLength(number)
	lexer.SPorHT()
	lexer.Match('\n')
	return contentLength, nil
	//            } catch (InvalidArgumentException ex) {
	//   throw createParseException(ex.getMessage());
	//            } catch (NumberFormatException ex) {
	//   throw createParseException(ex.getMessage());
	//            }  finally {
	// if (debug) dbg_leave("ContentLengthParser.leave");
	//     }
}
