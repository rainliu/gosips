package parser

import (
	"errors"
	"gosips/core"
	"strings"
)

/** A factory class that does a name lookup on a registered parser and
* returns a header parser for the given name.
 */

/** create a parser for a header. This is the parser factory.
 */
func CreateParser(line string) (parser Parser, ParseException error) {
	var lexer SIPLexer
	headerName := strings.TrimSpace(strings.ToLower(lexer.GetHeaderName(line)))
	headerValue := lexer.GetHeaderValue(line)
	if headerName == "" || headerValue == "" {
		return nil, errors.New("ParseException: The header name or value is null")
	}

	switch headerName {
	case strings.ToLower(core.SIPHeaderNames_REPLY_TO):
		parser = NewReplyToParser(line)
	case strings.ToLower(core.SIPHeaderNames_IN_REPLY_TO):
		parser = NewInReplyToParser(line)
	case strings.ToLower(core.SIPHeaderNames_ACCEPT_ENCODING):
		parser = NewAcceptEncodingParser(line)
	case strings.ToLower(core.SIPHeaderNames_ACCEPT_LANGUAGE):
		parser = NewAcceptLanguageParser(line)
	case "t":
		parser = NewToParser(line)
	case strings.ToLower(core.SIPHeaderNames_TO):
		parser = NewToParser(line)
	case strings.ToLower(core.SIPHeaderNames_FROM):
		parser = NewFromParser(line)
	case "f":
		parser = NewFromParser(line)
	case strings.ToLower(core.SIPHeaderNames_CSEQ):
		parser = NewCSeqParser(line)
	case strings.ToLower(core.SIPHeaderNames_VIA):
		parser = NewViaParser(line)
	case "v":
		parser = NewViaParser(line)
	case strings.ToLower(core.SIPHeaderNames_CONTACT):
		parser = NewContactParser(line)
	case "m":
		parser = NewContactParser(line)
	case strings.ToLower(core.SIPHeaderNames_CONTENT_TYPE):
		parser = NewContentTypeParser(line)
	case "c":
		parser = NewContentTypeParser(line)
	case strings.ToLower(core.SIPHeaderNames_CONTENT_LENGTH):
		parser = NewContentLengthParser(line)
	case "l":
		parser = NewContentLengthParser(line)
	case strings.ToLower(core.SIPHeaderNames_AUTHORIZATION):
		parser = NewAuthorizationParser(line)
	case strings.ToLower(core.SIPHeaderNames_WWW_AUTHENTICATE):
		parser = NewWWWAuthenticateParser(line)
	case strings.ToLower(core.SIPHeaderNames_CALL_ID):
		parser = NewCallIDParser(line)
	case "i":
		parser = NewCallIDParser(line)
	case strings.ToLower(core.SIPHeaderNames_ROUTE):
		parser = NewRouteParser(line)
	case strings.ToLower(core.SIPHeaderNames_RECORD_ROUTE):
		parser = NewRecordRouteParser(line)
	case strings.ToLower(core.SIPHeaderNames_DATE):
		parser = NewDateParser(line)
	case strings.ToLower(core.SIPHeaderNames_PROXY_AUTHORIZATION):
		parser = NewProxyAuthorizationParser(line)
	case strings.ToLower(core.SIPHeaderNames_PROXY_AUTHENTICATE):
		parser = NewProxyAuthenticateParser(line)
	case strings.ToLower(core.SIPHeaderNames_RETRY_AFTER):
		parser = NewRetryAfterParser(line)
	case strings.ToLower(core.SIPHeaderNames_REQUIRE):
		parser = NewRequireParser(line)
	case strings.ToLower(core.SIPHeaderNames_PROXY_REQUIRE):
		parser = NewProxyRequireParser(line)
	case strings.ToLower(core.SIPHeaderNames_TIMESTAMP):
		parser = NewTimeStampParser(line)
	case strings.ToLower(core.SIPHeaderNames_UNSUPPORTED):
		parser = NewUnsupportedParser(line)
	case strings.ToLower(core.SIPHeaderNames_USER_AGENT):
		parser = NewUserAgentParser(line)
	case strings.ToLower(core.SIPHeaderNames_SUPPORTED):
		parser = NewSupportedParser(line)
	case "k":
		parser = NewSupportedParser(line)
	case strings.ToLower(core.SIPHeaderNames_SERVER):
		parser = NewServerParser(line)
	case strings.ToLower(core.SIPHeaderNames_SUBJECT):
		parser = NewSubjectParser(line)
	case "s":
		parser = NewSubjectParser(line)
	case strings.ToLower(core.SIPHeaderNames_SUBSCRIPTION_STATE):
		parser = NewSubscriptionStateParser(line)
	case strings.ToLower(core.SIPHeaderNames_MAX_FORWARDS):
		parser = NewMaxForwardsParser(line)
	case strings.ToLower(core.SIPHeaderNames_MIME_VERSION):
		parser = NewMimeVersionParser(line)
	case strings.ToLower(core.SIPHeaderNames_MIN_EXPIRES):
		parser = NewMinExpiresParser(line)
	case strings.ToLower(core.SIPHeaderNames_ORGANIZATION):
		parser = NewOrganizationParser(line)
	case strings.ToLower(core.SIPHeaderNames_PRIORITY):
		parser = NewPriorityParser(line)
	case strings.ToLower(core.SIPHeaderNames_RACK):
		parser = NewRAckParser(line)
	case strings.ToLower(core.SIPHeaderNames_RSEQ):
		parser = NewRSeqParser(line)
	case strings.ToLower(core.SIPHeaderNames_REASON):
		parser = NewReasonParser(line)
	case strings.ToLower(core.SIPHeaderNames_WARNING):
		parser = NewWarningParser(line)
	case strings.ToLower(core.SIPHeaderNames_EXPIRES):
		parser = NewExpiresParser(line)
	case strings.ToLower(core.SIPHeaderNames_EVENT):
		parser = NewEventParser(line)
	case "o":
		parser = NewEventParser(line)
	case strings.ToLower(core.SIPHeaderNames_ERROR_INFO):
		parser = NewErrorInfoParser(line)
	case strings.ToLower(core.SIPHeaderNames_CONTENT_LANGUAGE):
		parser = NewContentLanguageParser(line)
	case strings.ToLower(core.SIPHeaderNames_CONTENT_ENCODING):
		parser = NewContentEncodingParser(line)
	case "e":
		parser = NewContentEncodingParser(line)
	case strings.ToLower(core.SIPHeaderNames_CONTENT_DISPOSITION):
		parser = NewContentDispositionParser(line)
	case strings.ToLower(core.SIPHeaderNames_CALL_INFO):
		parser = NewCallInfoParser(line)
	case strings.ToLower(core.SIPHeaderNames_AUTHENTICATION_INFO):
		parser = NewAuthenticationInfoParser(line)
	case strings.ToLower(core.SIPHeaderNames_ALLOW):
		parser = NewAllowParser(line)
	case strings.ToLower(core.SIPHeaderNames_ALLOW_EVENTS):
		parser = NewAllowEventsParser(line)
	case "u":
		parser = NewAllowEventsParser(line)
	case strings.ToLower(core.SIPHeaderNames_ALERT_INFO):
		parser = NewAlertInfoParser(line)
	case strings.ToLower(core.SIPHeaderNames_ACCEPT):
		parser = NewAcceptParser(line)
	case strings.ToLower(core.SIPHeaderNames_REFER_TO):
		parser = NewReferToParser(line)
	default:
		// Just generate a generic SIPHeader. We define
		// parsers only for the above.
		parser = NewHeaderParser(line)
	}

	return parser, nil
}
