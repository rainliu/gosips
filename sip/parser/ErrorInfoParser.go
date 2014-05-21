package parser

import (
	"gosips/core"
	"gosips/sip/address"
	"gosips/sip/header"
)

/** SIPParser for ErrorInfo header.
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
 * @return Header (ErrorInfoList object)
 * @throws SIPParseException if the message does not respect the spec.
 */
func (this *ErrorInfoParser) Parse() (sh header.Header, ParseException error) {
	errorInfoList := header.NewErrorInfoList()

	var ch byte
	var uri address.URI

	lexer := this.GetLexer()
	this.HeaderName(TokenTypes_ERROR_INFO)

	for ch, _ = lexer.LookAheadK(0); ch != '\n'; ch, _ = lexer.LookAheadK(0) {
		errorInfo := header.NewErrorInfo()
		errorInfo.SetHeaderName(core.SIPHeaderNames_ERROR_INFO)

		lexer.SPorHT()
		lexer.Match('<')
		urlParser := NewURLParserFromLexer(lexer)
		if uri, ParseException = urlParser.UriReference(); ParseException != nil {
			return nil, ParseException
		}
		errorInfo.SetErrorInfo(uri)
		lexer.Match('>')
		lexer.SPorHT()

		if ParseException = this.ParametersParser.Parse(errorInfo); ParseException != nil {
			return nil, ParseException
		}
		errorInfoList.PushBack(errorInfo)

		for ch, _ = lexer.LookAheadK(0); ch == ','; ch, _ = lexer.LookAheadK(0) {
			lexer.Match(',')
			lexer.SPorHT()

			errorInfo = header.NewErrorInfo()

			lexer.SPorHT()
			lexer.Match('<')
			urlParser = NewURLParserFromLexer(lexer)
			if uri, ParseException = urlParser.UriReference(); ParseException != nil {
				return nil, ParseException
			}
			errorInfo.SetErrorInfo(uri)
			lexer.Match('>')
			lexer.SPorHT()

			if ParseException = this.ParametersParser.Parse(errorInfo); ParseException != nil {
				return nil, ParseException
			}
			errorInfoList.PushBack(errorInfo)
		}
	}

	return errorInfoList, nil
}
