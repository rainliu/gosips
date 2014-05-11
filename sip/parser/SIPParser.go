package parser

import (
	"errors"
	"gosips/core"
)

const TokenTypes_START = core.CORELEXER_START

// Everything under this is reserved
const TokenTypes_END = core.CORELEXER_END

// End markder.

const TokenTypes_SIP = TokenTypes_START + 3
const TokenTypes_REGISTER = TokenTypes_START + 4
const TokenTypes_INVITE = TokenTypes_START + 5
const TokenTypes_ACK = TokenTypes_START + 6
const TokenTypes_BYE = TokenTypes_START + 7
const TokenTypes_OPTIONS = TokenTypes_START + 8
const TokenTypes_CANCEL = TokenTypes_START + 9
const TokenTypes_ERROR_INFO = TokenTypes_START + 10
const TokenTypes_IN_REPLY_TO = TokenTypes_START + 11
const TokenTypes_MIME_VERSION = TokenTypes_START + 12
const TokenTypes_ALERT_INFO = TokenTypes_START + 13
const TokenTypes_FROM = TokenTypes_START + 14
const TokenTypes_TO = TokenTypes_START + 15
const TokenTypes_VIA = TokenTypes_START + 16
const TokenTypes_USER_AGENT = TokenTypes_START + 17
const TokenTypes_SERVER = TokenTypes_START + 18
const TokenTypes_ACCEPT_ENCODING = TokenTypes_START + 19
const TokenTypes_ACCEPT = TokenTypes_START + 20
const TokenTypes_ALLOW = TokenTypes_START + 21
const TokenTypes_ROUTE = TokenTypes_START + 22
const TokenTypes_AUTHORIZATION = TokenTypes_START + 23
const TokenTypes_PROXY_AUTHORIZATION = TokenTypes_START + 24
const TokenTypes_RETRY_AFTER = TokenTypes_START + 25
const TokenTypes_PROXY_REQUIRE = TokenTypes_START + 26
const TokenTypes_CONTENT_LANGUAGE = TokenTypes_START + 27
const TokenTypes_UNSUPPORTED = TokenTypes_START + 28
const TokenTypes_SUPPORTED = TokenTypes_START + 20
const TokenTypes_WARNING = TokenTypes_START + 30
const TokenTypes_MAX_FORWARDS = TokenTypes_START + 31
const TokenTypes_DATE = TokenTypes_START + 32
const TokenTypes_PRIORITY = TokenTypes_START + 33
const TokenTypes_PROXY_AUTHENTICATE = TokenTypes_START + 34
const TokenTypes_CONTENT_ENCODING = TokenTypes_START + 35
const TokenTypes_CONTENT_LENGTH = TokenTypes_START + 36
const TokenTypes_SUBJECT = TokenTypes_START + 37
const TokenTypes_CONTENT_TYPE = TokenTypes_START + 38
const TokenTypes_CONTACT = TokenTypes_START + 39
const TokenTypes_CALL_ID = TokenTypes_START + 40
const TokenTypes_REQUIRE = TokenTypes_START + 41
const TokenTypes_EXPIRES = TokenTypes_START + 42
const TokenTypes_ENCRYPTION = TokenTypes_START + 43
const TokenTypes_RECORD_ROUTE = TokenTypes_START + 44
const TokenTypes_ORGANIZATION = TokenTypes_START + 45
const TokenTypes_CSEQ = TokenTypes_START + 46
const TokenTypes_ACCEPT_LANGUAGE = TokenTypes_START + 47
const TokenTypes_WWW_AUTHENTICATE = TokenTypes_START + 48
const TokenTypes_RESPONSE_KEY = TokenTypes_START + 49
const TokenTypes_HIDE = TokenTypes_START + 50
const TokenTypes_CALL_INFO = TokenTypes_START + 51
const TokenTypes_CONTENT_DISPOSITION = TokenTypes_START + 52
const TokenTypes_SUBSCRIBE = TokenTypes_START + 53
const TokenTypes_NOTIFY = TokenTypes_START + 54
const TokenTypes_TIMESTAMP = TokenTypes_START + 55
const TokenTypes_SUBSCRIPTION_STATE = TokenTypes_START + 56
const TokenTypes_TEL = TokenTypes_START + 57
const TokenTypes_REPLY_TO = TokenTypes_START + 58
const TokenTypes_REASON = TokenTypes_START + 59
const TokenTypes_RSEQ = TokenTypes_START + 60
const TokenTypes_RACK = TokenTypes_START + 61
const TokenTypes_MIN_EXPIRES = TokenTypes_START + 62
const TokenTypes_EVENT = TokenTypes_START + 63
const TokenTypes_AUTHENTICATION_INFO = TokenTypes_START + 64
const TokenTypes_ALLOW_EVENTS = TokenTypes_START + 65
const TokenTypes_REFER_TO = TokenTypes_START + 66
const TokenTypes_ALPHA = core.CORELEXER_ALPHA
const TokenTypes_DIGIT = core.CORELEXER_DIGIT
const TokenTypes_ID = core.CORELEXER_ID
const TokenTypes_WHITESPACE = core.CORELEXER_WHITESPACE
const TokenTypes_BACKSLASH = core.CORELEXER_BACKSLASH
const TokenTypes_QUOTE = core.CORELEXER_QUOTE
const TokenTypes_AT = core.CORELEXER_AT
const TokenTypes_SP = core.CORELEXER_SP
const TokenTypes_HT = core.CORELEXER_HT
const TokenTypes_COLON = core.CORELEXER_COLON
const TokenTypes_STAR = core.CORELEXER_STAR
const TokenTypes_DOLLAR = core.CORELEXER_DOLLAR
const TokenTypes_PLUS = core.CORELEXER_PLUS
const TokenTypes_POUND = core.CORELEXER_POUND
const TokenTypes_MINUS = core.CORELEXER_MINUS
const TokenTypes_DOUBLEQUOTE = core.CORELEXER_DOUBLEQUOTE
const TokenTypes_TILDE = core.CORELEXER_TILDE
const TokenTypes_BACK_QUOTE = core.CORELEXER_BACK_QUOTE
const TokenTypes_NULL = core.CORELEXER_NULL
const TokenTypes_EQUALS = (int)('=')
const TokenTypes_SEMICOLON = (int)(';')
const TokenTypes_SLASH = (int)('/')
const TokenTypes_L_SQUARE_BRACKET = (int)('[')
const TokenTypes_R_SQUARE_BRACKET = (int)(']')
const TokenTypes_R_CURLY = (int)('}')
const TokenTypes_L_CURLY = (int)('{')
const TokenTypes_HAT = (int)('^')
const TokenTypes_BAR = (int)('|')
const TokenTypes_DOT = (int)('.')
const TokenTypes_EXCLAMATION = (int)('!')
const TokenTypes_LPAREN = (int)('(')
const TokenTypes_RPAREN = (int)(')')
const TokenTypes_GREATER_THAN = (int)('>')
const TokenTypes_LESS_THAN = (int)('<')
const TokenTypes_PERCENT = (int)('%')
const TokenTypes_QUESTION = (int)('?')
const TokenTypes_AND = (int)('&')
const TokenTypes_UNDERSCORE = (int)('_')

/** Base parser class.
*
*@version  JAIN-SIP-1.1
*
*@author M. Ranganathan <mranga@nist.gov>  <br/>
*
*<a href="{@docRoot}/uncopyright.html">This code is in the public domain.</a>
*
 */

type SIPParser struct {
	core.CoreParser //implements TokenTypes {
}

func NewSIPParser(buffer string) *SIPParser {
	this := &SIPParser{}

	this.CoreParser.Super(buffer)
	this.CoreParser.SetLexer(NewSIPLexer("CharLexer", buffer))

	return this
}

func (this *SIPParser) super(buffer string) {
	this.CoreParser.Super(buffer)
	this.CoreParser.SetLexer(NewSIPLexer("CharLexer", buffer))
}

func (this *SIPParser) CreateParseException(exceptionString string) (ParseException error) {
	return errors.New("ParseException: " + this.GetLexer().GetBuffer() + ":" + exceptionString) // + this.GetLexer().GetPtr());
}

/*func (this *SIPParser) GetLexer() SIPLexer {
	return this.GetLexer();
}*/

func (this *SIPParser) SipVersion() (s string, ParseException error) {
	if core.Debug.ParserDebug {
		this.Dbg_enter("sipVersion")
		defer this.Dbg_leave("sipVersion")
	}

	//try {
	var tok *core.Token
	if tok, ParseException = this.GetLexer().Match(TokenTypes_SIP); ParseException != nil {
		return "", this.CreateParseException("Expecting SIP")
	}
	if tok.GetTokenValue() != "SIP" {
		return "", this.CreateParseException("Expecting SIP")
	}
	this.GetLexer().Match('/')
	if tok, ParseException = this.GetLexer().Match(TokenTypes_ID); ParseException != nil {
		return "", this.CreateParseException("Expecting SIP/2.0")
	}
	if tok.GetTokenValue() != "2.0" {
		return "", this.CreateParseException("Expecting SIP/2.0")
	}

	return "SIP/2.0", nil
	//} finally {
	//	if (debug) dbg_leave("sipVersion");
	//}
}

/** parses a method. Consumes if a valid method has been found.
 */
func (this *SIPParser) Method() (s string, ParseException error) {
	if core.Debug.ParserDebug {
		this.Dbg_enter("method")
		defer this.Dbg_leave("method")
	}
	//try {
	tokens, _ := this.GetLexer().PeekNextTokenK(1)
	//println(tokens[0].String())
	token := tokens[0]
	if token.GetTokenType() == TokenTypes_INVITE ||
		token.GetTokenType() == TokenTypes_ACK ||
		token.GetTokenType() == TokenTypes_OPTIONS ||
		token.GetTokenType() == TokenTypes_BYE ||
		token.GetTokenType() == TokenTypes_REGISTER ||
		token.GetTokenType() == TokenTypes_CANCEL ||
		token.GetTokenType() == TokenTypes_SUBSCRIBE ||
		token.GetTokenType() == TokenTypes_NOTIFY ||
		token.GetTokenType() == TokenTypes_ID {
		this.GetLexer().Consume()
		return token.GetTokenValue(), nil
	} else {
		return "", this.CreateParseException("Invalid Method")
	}
	//} finally {
	//      if (Debug.debug) dbg_leave("method");
	//}
}
