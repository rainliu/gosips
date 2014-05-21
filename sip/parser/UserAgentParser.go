package parser

import (
	"errors"
	"gosips/core"
	"gosips/sip/header"
	"strings"
)

/** SIPParser for UserAgent header.
 */
type UserAgentParser struct {
	HeaderParser
}

/** Constructor
 * @param userAgent - UserAgent header to parse
 *
 */
func NewUserAgentParser(userAgent string) *UserAgentParser {
	this := &UserAgentParser{}
	this.HeaderParser.super(userAgent)
	return this
}

/** Constructor
 * @param lexer - the lexer to use.
 */
func NewUserAgentParserFromLexer(lexer core.Lexer) *UserAgentParser {
	this := &UserAgentParser{}
	this.HeaderParser.superFromLexer(lexer)
	return this
}

/** parse the String message
 * @return SIPHeader (UserAgent object)
 * @throws SIPParseException if the message does not respect the spec.
 */
func (this *UserAgentParser) Parse() (sh header.Header, ParseException error) {
	userAgent := header.NewUserAgent()

	var ch byte
	lexer := this.GetLexer()
	this.HeaderName(TokenTypes_USER_AGENT)
	if ch, _ = lexer.LookAheadK(0); ch == '\n' {
		return nil, errors.New("empty header")
	}

	//  mandatory token: product[/product-version] | (comment)
	for ch, ParseException = lexer.LookAheadK(0); ch != '\n' && ParseException == nil; ch, ParseException = lexer.LookAheadK(0) {
		if ch == '(' {
			comment, _ := lexer.Comment()
			userAgent.AddProductToken("(" + comment + ")")
		} else {
			var tok string
			tok, ParseException = lexer.GetString('/')
			if ParseException != nil {
				tok = lexer.GetRest()
				userAgent.AddProductToken(strings.TrimSpace(tok))
				break
			} else {
				if tok[len(tok)-1] == '\n' {
					tok = strings.TrimSpace(tok)
				}
				userAgent.AddProductToken(tok)
			}
		}
	}

	return userAgent, nil
}
