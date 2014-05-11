package parser

import (
	"gosips/core"
	"gosips/sip/header"
)

/** SIPParser for Allow header.
 */
type AllowParser struct {
	HeaderParser
}

/**
 * Creates a new instance of AllowParser
 * @param allow the header to parse
 */
func NewAllowParser(allow string) *AllowParser {
	this := &AllowParser{}
	this.HeaderParser.super(allow)
	return this
}

/** Constructor
 * @param lexer the lexer to use to parse the header
 */
func NewAllowParserFromLexer(lexer core.Lexer) *AllowParser {
	this := &AllowParser{}
	this.HeaderParser.superFromLexer(lexer)
	return this
}

/** parse the Allow String header
 * @return Header (AllowList object)
 * @throws SIPParseException if the message does not respect the spec.
 */
func (this *AllowParser) Parse() (sh header.Header, ParseException error) {

	//if (debug) dbg_enter("AllowParser.parse");
	allowList := header.NewAllowList()

	//try {
	var ch byte

	lexer := this.GetLexer()
	this.HeaderName(TokenTypes_ALLOW)

	allow := header.NewAllow()
	allow.SetHeaderName(core.SIPHeaderNames_ALLOW)

	lexer.SPorHT()
	lexer.Match(TokenTypes_ID)
	token := lexer.GetNextToken()
	allow.SetMethod(token.GetTokenValue())

	allowList.PushBack(allow)
	lexer.SPorHT()
	for ch, _ = lexer.LookAheadK(0); ch == ','; ch, _ = lexer.LookAheadK(0) {
		lexer.Match(',')
		lexer.SPorHT()

		allow = header.NewAllow()
		lexer.Match(TokenTypes_ID)
		token = lexer.GetNextToken()
		allow.SetMethod(token.GetTokenValue())

		allowList.PushBack(allow)
		lexer.SPorHT()
	}
	lexer.SPorHT()
	lexer.Match('\n')

	return allowList, nil
	// }
	// finally {
	//     if (debug) dbg_leave("AllowParser.parse");
	// }
}
