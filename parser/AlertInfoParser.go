package parser

import (
	"gosip/core"
	"gosip/header"
)

/** Parser for AlertInfo header.
*
* @version 1.0
 */
type AlertInfoParser struct {
	ParametersParser
}

/**
 * Creates a new instance of AlertInfo Parser
 * @param alertInfo  the header to parse
 */
func NewAlertInfoParser(alertInfo string) *AlertInfoParser {
	this := &AlertInfoParser{}
	this.ParametersParser.super(alertInfo)
	return this
}

/** Constructor
 * @param lexer the lexer to use to parse the header
 */
func NewAlertInfoParserFromLexer(lexer core.Lexer) *AlertInfoParser {
	this := &AlertInfoParser{}
	this.ParametersParser.superFromLexer(lexer)
	return this
}

func (this *AlertInfoParser) super(buffer string) {
	this.ParametersParser.super(buffer)
}

func (this *AlertInfoParser) superFromLexer(lexer core.Lexer) {
	this.ParametersParser.superFromLexer(lexer)
}

/** parse the AlertInfo  String header
 * @return SIPHeader (AlertInfoList  object)
 * @throws SIPParseException if the message does not respect the spec.
 */
func (this *AlertInfoParser) Parse() (sh header.SIPHeader, ParseException error) {

	//if (debug) dbg_enter("AlertInfoParser.parse");
	alertInfoList := header.NewAlertInfoList()

	// try {
	var ch byte
	lexer := this.GetLexer()
	this.HeaderName(TokenTypes_ALERT_INFO)

	for ch, _ = lexer.LookAheadK(0); ch != '\n'; ch, _ = lexer.LookAheadK(0) {
		alertInfo := header.NewAlertInfo()
		alertInfo.SetHeaderName(core.SIPHeaderNames_ALERT_INFO)

		lexer.SPorHT()
		lexer.Match('<')
		urlParser := NewURLParserFromLexer(lexer)
		uri,_ := urlParser.UriReference()
		alertInfo.SetAlertInfo(uri)
		lexer.Match('>')
		lexer.SPorHT()

		this.ParametersParser.Parse(alertInfo)
		alertInfoList.PushBack(alertInfo)

		for ch, _ = lexer.LookAheadK(0); ch == ','; ch, _ = lexer.LookAheadK(0) {
			lexer.Match(',')
			lexer.SPorHT()

			alertInfo = header.NewAlertInfo()

			lexer.SPorHT()
			lexer.Match('<')
			urlParser = NewURLParserFromLexer(lexer)
			uri,_ = urlParser.UriReference()
			alertInfo.SetAlertInfo(uri)
			lexer.Match('>')
			lexer.SPorHT()

			this.ParametersParser.Parse(alertInfo)
			alertInfoList.PushBack(alertInfo)
		}
	}

	return alertInfoList, nil
	//        }
	//        finally {
	//            if (debug) dbg_leave("AlertInfoParser.parse");
	//        }
}
