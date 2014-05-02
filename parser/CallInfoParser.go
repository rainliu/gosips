package parser

import (
	"gosip/core"
	"gosip/header"
)

/** Parser for CallInfo header.
*
*@version  JAIN-SIP-1.1
*
*@author Olivier Deruelle <deruelle@nist.gov>
*@author M. Ranganathan <mranga@nist.gov>
*<a href="{@docRoot}/uncopyright.html">This code is in the public domain.</a>
*
* @version 1.0
 */
type CallInfoParser struct {
	ParametersParser
}

/**
 * Creates a new instance of CallInfoParser
 * @param callInfo the header to parse
 */
func NewCallInfoParser(callInfo string) *CallInfoParser {
	this := &CallInfoParser{}
	this.ParametersParser.super(callInfo)
	return this
}

/** Constructor
 * @param lexer the lexer to use to parse the header
 */
func NewCallInfoParserFromLexer(lexer core.Lexer) *CallInfoParser {
	this := &CallInfoParser{}
	this.ParametersParser.superFromLexer(lexer)
	return this
}

/** parse the CallInfo String header
 * @return Header (CallInfoList object)
 * @throws SIPParseException if the message does not respect the spec.
 */
func (this *CallInfoParser) Parse() (sh header.Header, ParseException error) {

	// if (debug) dbg_enter("CallInfoParser.parse");
	callInfoList := header.NewCallInfoList()

	//try {
	var ch byte
	lexer := this.GetLexer()
	this.HeaderName(TokenTypes_CALL_INFO)

	for ch, _ = lexer.LookAheadK(0); ch != '\n'; ch, _ = lexer.LookAheadK(0) {
		callInfo := header.NewCallInfo()
		callInfo.SetHeaderName(core.SIPHeaderNames_CALL_INFO)

		lexer.SPorHT()
		lexer.Match('<')
		urlParser := NewURLParserFromLexer(lexer)
		uri, _ := urlParser.UriReference()
		callInfo.SetInfo(uri)
		lexer.Match('>')
		lexer.SPorHT()

		this.ParametersParser.Parse(callInfo)
		callInfoList.PushBack(callInfo)

		for ch, _ = lexer.LookAheadK(0); ch == ','; ch, _ = lexer.LookAheadK(0) {
			lexer.Match(',')
			lexer.SPorHT()

			callInfo = header.NewCallInfo()

			lexer.SPorHT()
			lexer.Match('<')
			urlParser = NewURLParserFromLexer(lexer)
			uri, _ = urlParser.UriReference()
			callInfo.SetInfo(uri)
			lexer.Match('>')
			lexer.SPorHT()

			this.ParametersParser.Parse(callInfo)
			callInfoList.PushBack(callInfo)
		}
	}

	return callInfoList, nil
	// }
	// finally {
	//     if (debug) dbg_leave("CallInfoParser.parse");
	// }
}
