package parser

import (
	"errors"
	"gosips/core"
	"gosips/sip/header"
)

/** SIPParser for Content-Length Header.
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

func (this *ContentLengthParser) Parse() (sh header.Header, ParseException error) {
	// if (debug) dbg_enter("ContentLengthParser.enter");
	//      try {
	var number int
	contentLength := header.NewContentLength()
	lexer := this.GetLexer()
	this.HeaderName(TokenTypes_CONTENT_LENGTH)
	if number, ParseException = lexer.Number(); ParseException != nil {
		return nil, errors.New("ContentLengthParseException: Number " + ParseException.Error())
	}
	if err := contentLength.SetContentLength(number); err != nil {
		return nil, err
	}
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
