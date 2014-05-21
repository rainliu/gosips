package parser

import (
	"gosips/core"
	"strings"
)

/** A factory class that does a name lookup on a registered parser and
* returns a header parser for the given name.
 */

/** create a parser for a header. This is the parser factory.
 */
func CreateParser(line string) Parser { //, ParseException error) {
	var lexer SIPLexer
	headerName := strings.TrimSpace(strings.ToLower(lexer.GetHeaderName(line)))
	headerValue := lexer.GetHeaderValue(line)
	if headerName == "" || headerValue == "" {
		return nil //errors.New("ParseException: The header name or value is null")
	}

	switch headerName {
	case strings.ToLower(core.SIPHeaderNames_REPLY_TO):
		return NewReplyToParser(line)
	case strings.ToLower(core.SIPHeaderNames_IN_REPLY_TO):
		return NewInReplyToParser(line)
	case strings.ToLower(core.SIPHeaderNames_ACCEPT_ENCODING):
		return NewAcceptEncodingParser(line)
	case strings.ToLower(core.SIPHeaderNames_ACCEPT_LANGUAGE):
		return NewAcceptLanguageParser(line)
	case "t":
		return NewToParser(line)
	case strings.ToLower(core.SIPHeaderNames_TO):
		return NewToParser(line)
	case strings.ToLower(core.SIPHeaderNames_FROM):
		return NewFromParser(line)
	case "f":
		return NewFromParser(line)
	case strings.ToLower(core.SIPHeaderNames_CSEQ):
		return NewCSeqParser(line)
	case strings.ToLower(core.SIPHeaderNames_VIA):
		return NewViaParser(line)
	case "v":
		return NewViaParser(line)
	case strings.ToLower(core.SIPHeaderNames_CONTACT):
		return NewContactParser(line)
	case "m":
		return NewContactParser(line)
	case strings.ToLower(core.SIPHeaderNames_CONTENT_TYPE):
		return NewContentTypeParser(line)
	case "c":
		return NewContentTypeParser(line)
	case strings.ToLower(core.SIPHeaderNames_CONTENT_LENGTH):
		return NewContentLengthParser(line)
	case "l":
		return NewContentLengthParser(line)
	case strings.ToLower(core.SIPHeaderNames_AUTHORIZATION):
		return NewAuthorizationParser(line)
	case strings.ToLower(core.SIPHeaderNames_WWW_AUTHENTICATE):
		return NewWWWAuthenticateParser(line)
	case strings.ToLower(core.SIPHeaderNames_CALL_ID):
		return NewCallIDParser(line)
	case "i":
		return NewCallIDParser(line)
	case strings.ToLower(core.SIPHeaderNames_ROUTE):
		return NewRouteParser(line)
	case strings.ToLower(core.SIPHeaderNames_RECORD_ROUTE):
		return NewRecordRouteParser(line)
	case strings.ToLower(core.SIPHeaderNames_DATE):
		return NewDateParser(line)
	case strings.ToLower(core.SIPHeaderNames_PROXY_AUTHORIZATION):
		return NewProxyAuthorizationParser(line)
	case strings.ToLower(core.SIPHeaderNames_PROXY_AUTHENTICATE):
		return NewProxyAuthenticateParser(line)
	case strings.ToLower(core.SIPHeaderNames_RETRY_AFTER):
		return NewRetryAfterParser(line)
	case strings.ToLower(core.SIPHeaderNames_REQUIRE):
		return NewRequireParser(line)
	case strings.ToLower(core.SIPHeaderNames_PROXY_REQUIRE):
		return NewProxyRequireParser(line)
	case strings.ToLower(core.SIPHeaderNames_TIMESTAMP):
		return NewTimeStampParser(line)
	case strings.ToLower(core.SIPHeaderNames_UNSUPPORTED):
		return NewUnsupportedParser(line)
	case strings.ToLower(core.SIPHeaderNames_USER_AGENT):
		return NewUserAgentParser(line)
	case strings.ToLower(core.SIPHeaderNames_SUPPORTED):
		return NewSupportedParser(line)
	case "k":
		return NewSupportedParser(line)
	case strings.ToLower(core.SIPHeaderNames_SERVER):
		return NewServerParser(line)
	case strings.ToLower(core.SIPHeaderNames_SUBJECT):
		return NewSubjectParser(line)
	case "s":
		return NewSubjectParser(line)
	case strings.ToLower(core.SIPHeaderNames_SUBSCRIPTION_STATE):
		return NewSubscriptionStateParser(line)
	case strings.ToLower(core.SIPHeaderNames_MAX_FORWARDS):
		return NewMaxForwardsParser(line)
	case strings.ToLower(core.SIPHeaderNames_MIME_VERSION):
		return NewMimeVersionParser(line)
	case strings.ToLower(core.SIPHeaderNames_MIN_EXPIRES):
		return NewMinExpiresParser(line)
	case strings.ToLower(core.SIPHeaderNames_ORGANIZATION):
		return NewOrganizationParser(line)
	case strings.ToLower(core.SIPHeaderNames_PRIORITY):
		return NewPriorityParser(line)
	case strings.ToLower(core.SIPHeaderNames_RACK):
		return NewRAckParser(line)
	case strings.ToLower(core.SIPHeaderNames_RSEQ):
		return NewRSeqParser(line)
	case strings.ToLower(core.SIPHeaderNames_REASON):
		return NewReasonParser(line)
	case strings.ToLower(core.SIPHeaderNames_WARNING):
		return NewWarningParser(line)
	case strings.ToLower(core.SIPHeaderNames_EXPIRES):
		return NewExpiresParser(line)
	case strings.ToLower(core.SIPHeaderNames_EVENT):
		return NewEventParser(line)
	case "o":
		return NewEventParser(line)
	case strings.ToLower(core.SIPHeaderNames_ERROR_INFO):
		return NewErrorInfoParser(line)
	case strings.ToLower(core.SIPHeaderNames_CONTENT_LANGUAGE):
		return NewContentLanguageParser(line)
	case strings.ToLower(core.SIPHeaderNames_CONTENT_ENCODING):
		return NewContentEncodingParser(line)
	case "e":
		return NewContentEncodingParser(line)
	case strings.ToLower(core.SIPHeaderNames_CONTENT_DISPOSITION):
		return NewContentDispositionParser(line)
	case strings.ToLower(core.SIPHeaderNames_CALL_INFO):
		return NewCallInfoParser(line)
	case strings.ToLower(core.SIPHeaderNames_AUTHENTICATION_INFO):
		return NewAuthenticationInfoParser(line)
	case strings.ToLower(core.SIPHeaderNames_ALLOW):
		return NewAllowParser(line)
	case strings.ToLower(core.SIPHeaderNames_ALLOW_EVENTS):
		return NewAllowEventsParser(line)
	case "u":
		return NewAllowEventsParser(line)
	case strings.ToLower(core.SIPHeaderNames_ALERT_INFO):
		return NewAlertInfoParser(line)
	case strings.ToLower(core.SIPHeaderNames_ACCEPT):
		return NewAcceptParser(line)
	case strings.ToLower(core.SIPHeaderNames_REFER_TO):
		return NewReferToParser(line)
	default:
		// Just generate a generic SIPHeader. We define
		// parsers only for the above.
		return NewHeaderParser(line)
	}

	return nil
}
