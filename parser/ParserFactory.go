package parser

import (
	"gosip/core"
	"strings"
)

/** A factory class that does a name lookup on a registered parser and
* returns a header parser for the given name.
*
*@version  JAIN-SIP-1.1
*
*@author M. Ranganathan <mranga@nist.gov>  <br/>
*
*<a href="{@docRoot}/uncopyright.html">This code is in the public domain.</a>
*
 */

//    // Was missing (bug noticed by Steve Crossley)
//    parserTable.put("r",ReferToParser()

//   }
/** create a parser for a header. This is the parser factory.
 */
func CreateParser(line string) IHeaderParser { //, ParseException error) {
	var lexer Lexer
	headerName := strings.ToLower(lexer.GetHeaderName(line))
	headerValue := lexer.GetHeaderValue(line)
	if headerName == "" || headerValue == "" {
		return nil //errors.New("ParseException: The header name or value is null")
	}

	//Class parserClass =(Class) parserTable.get(headerName.toLowerCase());

	switch headerName {
	case core.SIPHeaderNames_REPLY_TO:
		return NewReplyToParser(line)
	case core.SIPHeaderNames_IN_REPLY_TO:
		return NewInReplyToParser(line)
	case core.SIPHeaderNames_ACCEPT_ENCODING:
		return NewAcceptEncodingParser(line)
	case core.SIPHeaderNames_ACCEPT_LANGUAGE:
		return NewAcceptLanguageParser(line)
	case "t":
		return NewToParser(line)
	case core.SIPHeaderNames_TO:
		return NewToParser(line)
	case core.SIPHeaderNames_FROM:
		return NewFromParser(line)
	case "f":
		return NewFromParser(line)
	case core.SIPHeaderNames_CSEQ:
		return NewCSeqParser(line)
	case core.SIPHeaderNames_VIA:
		return NewViaParser(line)
	case "v":
		return NewViaParser(line)
	case core.SIPHeaderNames_CONTACT:
		return NewContactParser(line)
	case "m":
		return NewContactParser(line)
	case core.SIPHeaderNames_CONTENT_TYPE:
		return NewContentTypeParser(line)
	case "c":
		return NewContentTypeParser(line)
	case core.SIPHeaderNames_CONTENT_LENGTH:
		return NewContentLengthParser(line)
	case "l":
		return NewContentLengthParser(line)
	case core.SIPHeaderNames_AUTHORIZATION:
		return NewAuthorizationParser(line)
	case core.SIPHeaderNames_WWW_AUTHENTICATE:
		return NewWWWAuthenticateParser(line)
	case core.SIPHeaderNames_CALL_ID:
		return NewCallIDParser(line)
	case "i":
		return NewCallIDParser(line)
	case core.SIPHeaderNames_ROUTE:
		return NewRouteParser(line)
	case core.SIPHeaderNames_RECORD_ROUTE:
		return NewRecordRouteParser(line)
	case core.SIPHeaderNames_DATE:
		return NewDateParser(line)
	case core.SIPHeaderNames_PROXY_AUTHORIZATION:
		return NewProxyAuthorizationParser(line)
	case core.SIPHeaderNames_PROXY_AUTHENTICATE:
		return NewProxyAuthenticateParser(line)
	case core.SIPHeaderNames_RETRY_AFTER:
		return NewRetryAfterParser(line)
	case core.SIPHeaderNames_REQUIRE:
		return NewRequireParser(line)
	case core.SIPHeaderNames_PROXY_REQUIRE:
		return NewProxyRequireParser(line)
	case core.SIPHeaderNames_TIMESTAMP:
		return NewTimeStampParser(line)
	case core.SIPHeaderNames_UNSUPPORTED:
		return NewUnsupportedParser(line)
	case core.SIPHeaderNames_USER_AGENT:
		return NewUserAgentParser(line)
	case core.SIPHeaderNames_SUPPORTED:
		return NewSupportedParser(line)
	case "k":
		return NewSupportedParser(line)
	case core.SIPHeaderNames_SERVER:
		return NewServerParser(line)
	case core.SIPHeaderNames_SUBJECT:
		return NewSubjectParser(line)
	case core.SIPHeaderNames_SUBSCRIPTION_STATE:
		return NewSubscriptionStateParser(line)
	case core.SIPHeaderNames_MAX_FORWARDS:
		return NewMaxForwardsParser(line)
	case core.SIPHeaderNames_MIME_VERSION:
		return NewMimeVersionParser(line)
	case core.SIPHeaderNames_MIN_EXPIRES:
		return NewMinExpiresParser(line)
	case core.SIPHeaderNames_ORGANIZATION:
		return NewOrganizationParser(line)
	case core.SIPHeaderNames_PRIORITY:
		return NewPriorityParser(line)
	case core.SIPHeaderNames_RACK:
		return NewRAckParser(line)
	case core.SIPHeaderNames_RSEQ:
		return NewRSeqParser(line)
	case core.SIPHeaderNames_REASON:
		return NewReasonParser(line)
	case core.SIPHeaderNames_WARNING:
		return NewWarningParser(line)
	case core.SIPHeaderNames_EXPIRES:
		return NewExpiresParser(line)
	case core.SIPHeaderNames_EVENT:
		return NewEventParser(line)
	case "o":
		return NewEventParser(line)
	case core.SIPHeaderNames_ERROR_INFO:
		return NewErrorInfoParser(line)
	case core.SIPHeaderNames_CONTENT_LANGUAGE:
		return NewContentLanguageParser(line)
	case core.SIPHeaderNames_CONTENT_ENCODING:
		return NewContentEncodingParser(line)
	case "e":
		return NewContentEncodingParser(line)
	case core.SIPHeaderNames_CONTENT_DISPOSITION:
		return NewContentDispositionParser(line)
	case core.SIPHeaderNames_CALL_INFO:
		return NewCallInfoParser(line)
	case core.SIPHeaderNames_AUTHENTICATION_INFO:
		return NewAuthenticationInfoParser(line)
	case core.SIPHeaderNames_ALLOW:
		return NewAllowParser(line)
	case core.SIPHeaderNames_ALLOW_EVENTS:
		return NewAllowEventsParser(line)
	case "u":
		return NewAllowEventsParser(line)
	case core.SIPHeaderNames_ALERT_INFO:
		return NewAlertInfoParser(line)
	case core.SIPHeaderNames_ACCEPT:
		return NewAcceptParser(line)
	case core.SIPHeaderNames_REFER_TO:
		return NewReferToParser(line)
	default:
		// Just generate a generic SIPHeader. We define
		// parsers only for the above.
		return NewHeaderParser(line)
	}

	return nil
}
