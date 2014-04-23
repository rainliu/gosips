package parser

import (
	"gosip/core"
	"gosip/header"
)

/** Parser for ErrorInfo header.
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
type ErrorInfoParser struct {
	ParametersParser
}

/**
 * Creates a new instance of ErrorInfoParser
 * @param errorInfo the header to parse
 */
func NewErrorInfoParser(errorInfo string) *ErrorInfoParser {

	this := &ErrorInfoParser{}
	this.ParametersParser.super(errorInfo)
	return this
}

/** Constructor
 * @param lexer the lexer to use to parse the header
 */
func NewErrorInfoParserFromLexer(lexer core.Lexer) *ErrorInfoParser {

	this := &ErrorInfoParser{}
	this.ParametersParser.superFromLexer(lexer)
	return this
}

/** parse the ErrorInfo String header
 * @return SIPHeaderHeader (ErrorInfoList object)
 * @throws SIPParseException if the message does not respect the spec.
 */
func (this *ErrorInfoParser) Parse() (sh header.SIPHeaderHeader, ParseException error) {

	// if (debug) dbg_enter("ErrorInfoParser.parse");
	errorInfoList := header.NewErrorInfoList()

	//try {
	var ch byte

	lexer := this.GetLexer()
	this.HeaderName(TokenTypes_ERROR_INFO)

	for ch, _ = lexer.LookAheadK(0); ch != '\n'; ch, _ = lexer.LookAheadK(0) {
		//while (lexer.lookAhead(0) != '\n') {
		errorInfo := header.NewErrorInfo()
		errorInfo.SetHeaderName(core.SIPHeaderNames_ERROR_INFO)

		lexer.SPorHT()
		lexer.Match('<')
		urlParser := NewURLParserFromLexer(lexer)
		uri, _ := urlParser.UriReference()
		errorInfo.SetErrorInfo(uri)
		lexer.Match('>')
		lexer.SPorHT()

		this.ParametersParser.Parse(errorInfo)
		errorInfoList.PushBack(errorInfo)

		for ch, _ = lexer.LookAheadK(0); ch == ','; ch, _ = lexer.LookAheadK(0) {
			//while (lexer.lookAhead(0) == ',') {
			lexer.Match(',')
			lexer.SPorHT()

			errorInfo = header.NewErrorInfo()

			lexer.SPorHT()
			lexer.Match('<')
			urlParser = NewURLParserFromLexer(lexer)
			uri, _ = urlParser.UriReference()
			errorInfo.SetErrorInfo(uri)
			lexer.Match('>')
			lexer.SPorHT()

			this.ParametersParser.Parse(errorInfo)
			errorInfoList.PushBack(errorInfo)
		}
	}

	return errorInfoList, nil
	// }
	// finally {
	//     if (debug) dbg_leave("ErrorInfoParser.parse");
	// }
}
