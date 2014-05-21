package parser

import (
	"gosips/core"
	"strings"
)

/** SIPLexer class for the parser.
 */
type SIPLexer struct {
	core.CoreLexer
}

func NewSIPLexer(lexerName, buffer string) *SIPLexer {
	this := &SIPLexer{}
	this.CoreLexer.Super(lexerName, buffer)
	this.SelectLexer(lexerName)

	return this
}

func (this *SIPLexer) super(lexerName, buffer string) {
	this.CoreLexer.Super(lexerName, buffer)
	this.SelectLexer(lexerName)
}

/** get the header name of the line
 *  @return String
 */
func (this *SIPLexer) GetHeaderName(line string) string {
	if line == "" {
		return ""
	}
	var headerName string
	begin := strings.Index(line, ":")
	if begin >= 1 {
		headerName = line[0:begin]
	}

	return headerName
}

/** get the header value of the line
 *  @return String
 */
func (this *SIPLexer) GetHeaderValue(line string) string {
	if line == "" {
		return ""
	}
	var headerValue string
	begin := strings.Index(line, ":")
	if begin != -1 {
		headerValue = line[begin+1:]
	}

	return headerValue
}
func (this *SIPLexer) SelectLexer(lexerName string) {
	this.CoreLexer.SelectLexer(lexerName)
	if this.CurrentLexer() == nil {
		this.AddLexer(lexerName)
		if lexerName == "method_keywordLexer" {
			this.AddKeyword(strings.ToUpper(core.SIPTransportNames_SIP), TokenTypes_SIP)
			this.AddKeyword(strings.ToUpper(core.SIPMethodNames_REGISTER), TokenTypes_REGISTER)
			this.AddKeyword(strings.ToUpper(core.SIPMethodNames_ACK), TokenTypes_ACK)
			this.AddKeyword(strings.ToUpper(core.SIPMethodNames_OPTIONS), TokenTypes_OPTIONS)
			this.AddKeyword(strings.ToUpper(core.SIPMethodNames_BYE), TokenTypes_BYE)
			this.AddKeyword(strings.ToUpper(core.SIPMethodNames_INVITE), TokenTypes_INVITE)
			this.AddKeyword(strings.ToUpper(core.SIPMethodNames_SUBSCRIBE), TokenTypes_SUBSCRIBE)
			this.AddKeyword(strings.ToUpper(core.SIPMethodNames_NOTIFY), TokenTypes_NOTIFY)
		} else if lexerName == "command_keywordLexer" {
			this.AddKeyword(strings.ToUpper(core.SIPHeaderNames_ERROR_INFO), TokenTypes_ERROR_INFO)
			this.AddKeyword(strings.ToUpper(core.SIPHeaderNames_ALLOW_EVENTS), TokenTypes_ALLOW_EVENTS)
			this.AddKeyword(strings.ToUpper(core.SIPHeaderNames_AUTHENTICATION_INFO), TokenTypes_AUTHENTICATION_INFO)
			this.AddKeyword(strings.ToUpper(core.SIPHeaderNames_EVENT), TokenTypes_EVENT)
			this.AddKeyword(strings.ToUpper(core.SIPHeaderNames_MIN_EXPIRES), TokenTypes_MIN_EXPIRES)
			this.AddKeyword(strings.ToUpper(core.SIPHeaderNames_RSEQ), TokenTypes_RSEQ)
			this.AddKeyword(strings.ToUpper(core.SIPHeaderNames_RACK), TokenTypes_RACK)
			this.AddKeyword(strings.ToUpper(core.SIPHeaderNames_REASON), TokenTypes_REASON)
			this.AddKeyword(strings.ToUpper(core.SIPHeaderNames_REPLY_TO), TokenTypes_REPLY_TO)
			this.AddKeyword(strings.ToUpper(core.SIPHeaderNames_SUBSCRIPTION_STATE), TokenTypes_SUBSCRIPTION_STATE)
			this.AddKeyword(strings.ToUpper(core.SIPHeaderNames_TIMESTAMP), TokenTypes_TIMESTAMP)
			this.AddKeyword(strings.ToUpper(core.SIPHeaderNames_IN_REPLY_TO), TokenTypes_IN_REPLY_TO)
			this.AddKeyword(strings.ToUpper(core.SIPHeaderNames_MIME_VERSION), TokenTypes_MIME_VERSION)
			this.AddKeyword(strings.ToUpper(core.SIPHeaderNames_ALERT_INFO), TokenTypes_ALERT_INFO)
			this.AddKeyword(strings.ToUpper(core.SIPHeaderNames_FROM), TokenTypes_FROM)
			this.AddKeyword(strings.ToUpper(core.SIPHeaderNames_TO), TokenTypes_TO)
			this.AddKeyword(strings.ToUpper(core.SIPHeaderNames_REFER_TO), TokenTypes_REFER_TO)
			this.AddKeyword(strings.ToUpper(core.SIPHeaderNames_VIA), TokenTypes_VIA)
			this.AddKeyword(strings.ToUpper(core.SIPHeaderNames_USER_AGENT), TokenTypes_USER_AGENT)
			this.AddKeyword(strings.ToUpper(core.SIPHeaderNames_SERVER), TokenTypes_SERVER)
			this.AddKeyword(strings.ToUpper(core.SIPHeaderNames_ACCEPT_ENCODING), TokenTypes_ACCEPT_ENCODING)
			this.AddKeyword(strings.ToUpper(core.SIPHeaderNames_ACCEPT), TokenTypes_ACCEPT)
			this.AddKeyword(strings.ToUpper(core.SIPHeaderNames_ALLOW), TokenTypes_ALLOW)
			this.AddKeyword(strings.ToUpper(core.SIPHeaderNames_ROUTE), TokenTypes_ROUTE)
			this.AddKeyword(strings.ToUpper(core.SIPHeaderNames_AUTHORIZATION), TokenTypes_AUTHORIZATION)
			this.AddKeyword(strings.ToUpper(core.SIPHeaderNames_PROXY_AUTHORIZATION), TokenTypes_PROXY_AUTHORIZATION)
			this.AddKeyword(strings.ToUpper(core.SIPHeaderNames_RETRY_AFTER), TokenTypes_RETRY_AFTER)
			this.AddKeyword(strings.ToUpper(core.SIPHeaderNames_PROXY_REQUIRE), TokenTypes_PROXY_REQUIRE)
			this.AddKeyword(strings.ToUpper(core.SIPHeaderNames_CONTENT_LANGUAGE), TokenTypes_CONTENT_LANGUAGE)
			this.AddKeyword(strings.ToUpper(core.SIPHeaderNames_UNSUPPORTED), TokenTypes_UNSUPPORTED)
			this.AddKeyword(strings.ToUpper(core.SIPHeaderNames_SUPPORTED), TokenTypes_SUPPORTED)
			this.AddKeyword(strings.ToUpper(core.SIPHeaderNames_WARNING), TokenTypes_WARNING)
			this.AddKeyword(strings.ToUpper(core.SIPHeaderNames_MAX_FORWARDS), TokenTypes_MAX_FORWARDS)
			this.AddKeyword(strings.ToUpper(core.SIPHeaderNames_DATE), TokenTypes_DATE)
			this.AddKeyword(strings.ToUpper(core.SIPHeaderNames_PRIORITY), TokenTypes_PRIORITY)
			this.AddKeyword(strings.ToUpper(core.SIPHeaderNames_PROXY_AUTHENTICATE), TokenTypes_PROXY_AUTHENTICATE)
			this.AddKeyword(strings.ToUpper(core.SIPHeaderNames_CONTENT_ENCODING), TokenTypes_CONTENT_ENCODING)
			this.AddKeyword(strings.ToUpper(core.SIPHeaderNames_CONTENT_LENGTH), TokenTypes_CONTENT_LENGTH)
			this.AddKeyword(strings.ToUpper(core.SIPHeaderNames_SUBJECT), TokenTypes_SUBJECT)
			this.AddKeyword(strings.ToUpper(core.SIPHeaderNames_CONTENT_TYPE), TokenTypes_CONTENT_TYPE)
			this.AddKeyword(strings.ToUpper(core.SIPHeaderNames_CONTACT), TokenTypes_CONTACT)
			this.AddKeyword(strings.ToUpper(core.SIPHeaderNames_CALL_ID), TokenTypes_CALL_ID)
			this.AddKeyword(strings.ToUpper(core.SIPHeaderNames_REQUIRE), TokenTypes_REQUIRE)
			this.AddKeyword(strings.ToUpper(core.SIPHeaderNames_EXPIRES), TokenTypes_EXPIRES)
			this.AddKeyword(strings.ToUpper(core.SIPHeaderNames_RECORD_ROUTE), TokenTypes_RECORD_ROUTE)
			this.AddKeyword(strings.ToUpper(core.SIPHeaderNames_ORGANIZATION), TokenTypes_ORGANIZATION)
			this.AddKeyword(strings.ToUpper(core.SIPHeaderNames_CSEQ), TokenTypes_CSEQ)
			this.AddKeyword(strings.ToUpper(core.SIPHeaderNames_ACCEPT_LANGUAGE), TokenTypes_ACCEPT_LANGUAGE)
			this.AddKeyword(strings.ToUpper(core.SIPHeaderNames_WWW_AUTHENTICATE), TokenTypes_WWW_AUTHENTICATE)
			this.AddKeyword(strings.ToUpper(core.SIPHeaderNames_CALL_INFO), TokenTypes_CALL_INFO)
			this.AddKeyword(strings.ToUpper(core.SIPHeaderNames_CONTENT_DISPOSITION), TokenTypes_CONTENT_DISPOSITION)
			// And now the dreaded short forms....
			this.AddKeyword(strings.ToUpper(core.SIPHeaderNames_K), TokenTypes_SUPPORTED)
			this.AddKeyword(strings.ToUpper(core.SIPHeaderNames_C), TokenTypes_CONTENT_TYPE)
			this.AddKeyword(strings.ToUpper(core.SIPHeaderNames_E), TokenTypes_CONTENT_ENCODING)
			this.AddKeyword(strings.ToUpper(core.SIPHeaderNames_F), TokenTypes_FROM)
			this.AddKeyword(strings.ToUpper(core.SIPHeaderNames_I), TokenTypes_CALL_ID)
			this.AddKeyword(strings.ToUpper(core.SIPHeaderNames_M), TokenTypes_CONTACT)
			this.AddKeyword(strings.ToUpper(core.SIPHeaderNames_L), TokenTypes_CONTENT_LENGTH)
			this.AddKeyword(strings.ToUpper(core.SIPHeaderNames_S), TokenTypes_SUBJECT)
			this.AddKeyword(strings.ToUpper(core.SIPHeaderNames_T), TokenTypes_TO)
			this.AddKeyword(strings.ToUpper(core.SIPHeaderNames_V), TokenTypes_VIA)
			this.AddKeyword(strings.ToUpper(core.SIPHeaderNames_R), TokenTypes_REFER_TO)
		} else if lexerName == "status_lineLexer" {
			this.AddKeyword(strings.ToUpper(core.SIPTransportNames_SIP), TokenTypes_SIP)
		} else if lexerName == "request_lineLexer" {
			this.AddKeyword(strings.ToUpper(core.SIPTransportNames_SIP), TokenTypes_SIP)
		} else if lexerName == "sip_urlLexer" {
			this.AddKeyword(strings.ToUpper(core.SIPTransportNames_TEL), TokenTypes_TEL)
			this.AddKeyword(strings.ToUpper(core.SIPTransportNames_SIP), TokenTypes_SIP)
		}
	} /*else{
		println("this.CurrentLexer() != nil");
	}*/
}
