package parser

import (
	"gosip/core"
	"gosip/header"
)

/** Parser for the SIP request line.
*
*@version  JAIN-SIP-1.1
*
*@author M. Ranganathan <mranga@nist.gov>  <br/>
*
*<a href="{@docRoot}/uncopyright.html">This code is in the public domain.</a>
 */

type RequestLineParser struct {
	Parser
}

func NewRequestLineParser(requestLine string) *RequestLineParser {
	this := &RequestLineParser{}
	this.ParserCore.Super(requestLine)
	this.SetLexer(NewLexer("method_keywordLexer", requestLine))
	return this
}
func NewRequestLineParserFromLexer(lexer core.Lexer) *RequestLineParser {
	this := &RequestLineParser{}
	this.SetLexer(lexer)
	this.GetLexer().SelectLexer("method_keywordLexer")
	return this
}

func (this *RequestLineParser) Parse() (rl *header.RequestLine, ParseException error) {
	// if (debug) dbg_enter("parse");
	// try {
	retval := header.NewRequestLine()
	lexer := this.GetLexer()
	m, _ := this.Method()
	lexer.SPorHT()
	retval.SetMethod(m)
	lexer.SelectLexer("sip_urlLexer")
	urlParser := NewURLParserFromLexer(this.GetLexer())
	url, _ := urlParser.UriReference()
	lexer.SPorHT()
	retval.SetUri(url)
	lexer.SelectLexer("request_lineLexer")
	v, _ := this.SipVersion()
	retval.SetSipVersion(v)
	lexer.SPorHT()
	lexer.Match('\n')
	return retval, nil
	//   } finally {
	// if (debug) dbg_leave("parse");
	//   }
}
