package parser

import (
	"gosips/core"
	"gosips/sip/header"
	"strings"
)

/** Parser for the SIP status line.
*
*@version  JAIN-SIP-1.1
*
*@author M. Ranganathan <mranga@nist.gov>  <br/>
*
*<a href="{@docRoot}/uncopyright.html">This code is in the public domain.</a>
*
 */

type StatusLineParser struct {
	Parser
}

func NewStatusLineParser(statusLine string) *StatusLineParser {
	this := &StatusLineParser{}
	this.ParserCore.Super(statusLine)
	this.SetLexer(NewLexer("status_lineLexer", statusLine))
	return this
}

func NewStatusLineParserFromLexer(lexer core.Lexer) *StatusLineParser {
	this := &StatusLineParser{}
	this.SetLexer(lexer)
	this.GetLexer().SelectLexer("status_lineLexer")
	return this
}

func (this *StatusLineParser) StatusCode() (scode int, ParseException error) {
	scode, ParseException = this.GetLexer().Number()
	//if (debug) dbg_enter("statusCode");
	//try {
	//int retval = Integer.parseInt(scode);
	return scode, ParseException
	// } catch (NumberFormatException ex) {
	//     throw new ParseException(lexer.getBuffer() +
	// 		":" + ex.getMessage(), lexer.getPtr());
	// } finally {
	// 	if (debug) dbg_leave("statusCode");
	// }

}

func (this *StatusLineParser) ReasonPhrase() string { //, ParseException error){
	return strings.TrimSpace(this.GetLexer().GetRest()) //.trim();
}

func (this *StatusLineParser) Parse() (sl *header.StatusLine, ParseException error) {
	//try {
	//if (debug) dbg_enter("parse");
	retval := header.NewStatusLine()
	lexer := this.GetLexer()
	version, _ := this.SipVersion()
	retval.SetSipVersion(version)
	lexer.SPorHT()
	scode, _ := this.StatusCode()
	retval.SetStatusCode(scode)
	lexer.SPorHT()
	rp := this.ReasonPhrase()
	retval.SetReasonPhrase(rp)
	lexer.SPorHT()
	return retval, nil
	// } finally {
	//    if (debug) dbg_leave("parse");
	// }
}
