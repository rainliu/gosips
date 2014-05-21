package parser

import (
	"gosips/core"
	"gosips/sip/header"
	"strings"
)

/** SIPParser for the SIP status line.
 */
type StatusLineParser struct {
	SIPParser
}

func NewStatusLineParser(statusLine string) *StatusLineParser {
	this := &StatusLineParser{}
	this.CoreParser.Super(statusLine)
	this.SetLexer(NewSIPLexer("status_lineLexer", statusLine))
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
	return scode, ParseException
}

func (this *StatusLineParser) ReasonPhrase() string {
	return strings.TrimSpace(this.GetLexer().GetRest())
}

func (this *StatusLineParser) Parse() (sl *header.StatusLine, ParseException error) {
	retval := header.NewStatusLine()
	lexer := this.GetLexer()

	var version string
	if version, ParseException = this.SipVersion(); ParseException != nil {
		return nil, ParseException
	}
	retval.SetSipVersion(version)
	lexer.SPorHT()

	var scode int
	if scode, ParseException = this.StatusCode(); ParseException != nil {
		return nil, ParseException
	}
	retval.SetStatusCode(scode)
	lexer.SPorHT()
	rp := this.ReasonPhrase()
	retval.SetReasonPhrase(rp)
	lexer.SPorHT()
	return retval, nil
}
