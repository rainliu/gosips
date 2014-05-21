package parser

import (
	"gosips/core"
	"gosips/sip/address"
	"gosips/sip/header"
)

/** SIPParser for the SIP request line.
 */
type RequestLineParser struct {
	SIPParser
}

func NewRequestLineParser(requestLine string) *RequestLineParser {
	this := &RequestLineParser{}
	this.CoreParser.Super(requestLine)
	this.SetLexer(NewSIPLexer("method_keywordLexer", requestLine))
	return this
}
func NewRequestLineParserFromLexer(lexer core.Lexer) *RequestLineParser {
	this := &RequestLineParser{}
	this.SetLexer(lexer)
	this.GetLexer().SelectLexer("method_keywordLexer")
	return this
}

func (this *RequestLineParser) Parse() (rl *header.RequestLine, ParseException error) {
	var m, v string
	var url address.URI

	retval := header.NewRequestLine()
	lexer := this.GetLexer()

	if m, ParseException = this.Method(); ParseException != nil {
		return nil, ParseException
	}

	lexer.SPorHT()
	retval.SetMethod(m)
	lexer.SelectLexer("sip_urlLexer")
	urlParser := NewURLParserFromLexer(this.GetLexer())
	if url, ParseException = urlParser.UriReference(); ParseException != nil {
		return nil, ParseException
	}
	lexer.SPorHT()
	retval.SetUri(url)
	lexer.SelectLexer("request_lineLexer")

	if v, ParseException = this.SipVersion(); ParseException != nil {
		return nil, ParseException
	}
	retval.SetSipVersion(v)
	lexer.SPorHT()
	lexer.Match('\n')

	return retval, nil
}
