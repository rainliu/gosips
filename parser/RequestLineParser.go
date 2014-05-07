package parser

import (
	"gosip/address"
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
	var m, v string
	var err error
	var url address.URI

	retval := header.NewRequestLine()
	lexer := this.GetLexer()
	println(lexer.GetRest())
	if m, err = this.Method(); err != nil {
		return nil, err
	}
	println(lexer.GetRest())
	lexer.SPorHT()
	retval.SetMethod(m)
	lexer.SelectLexer("sip_urlLexer")
	urlParser := NewURLParserFromLexer(this.GetLexer())
	if url, err = urlParser.UriReference(); err != nil {
		return nil, err
	}
	lexer.SPorHT()
	retval.SetUri(url)
	lexer.SelectLexer("request_lineLexer")
	println(lexer.GetRest())
	if v, err = this.SipVersion(); err != nil {
		return nil, err
	}
	retval.SetSipVersion(v)
	lexer.SPorHT()
	lexer.Match('\n')
	return retval, nil
	//   } finally {
	// if (debug) dbg_leave("parse");
	//   }
}
