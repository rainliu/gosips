package parser

import (
	"errors"
	"gosips/core"
	"gosips/sip/header"
	"strings"
)

/** SIPParser for Server header.
 */
type ServerParser struct {
	HeaderParser
}

/** Creates a new instance of ServerParser
 * @param server the header to parse
 */
func NewServerParser(server string) *ServerParser {
	this := &ServerParser{}
	this.HeaderParser.super(server)
	return this
}

/** Constructor
 * @param lexer the lexer to use to parse the header
 */
func NewServerParserFromLexer(lexer core.Lexer) *ServerParser {
	this := &ServerParser{}
	this.HeaderParser.superFromLexer(lexer)
	return this
}

/** parse the String server
 * @return SIPHeader (Server object)
 * @throws SIPParseException if the message does not respect the spec.
 */
func (this *ServerParser) Parse() (sh header.Header, ParseException error) {
	server := header.NewServer()

	var ch byte
	lexer := this.GetLexer()
	this.HeaderName(TokenTypes_SERVER)
	if ch, _ = lexer.LookAheadK(0); ch == '\n' {
		return nil, errors.New("empty header")
	}

	//  mandatory token: product[/product-version] | (comment)
	for ch, ParseException = lexer.LookAheadK(0); ch != '\n' && ParseException == nil; ch, ParseException = lexer.LookAheadK(0) {
		if ch == '(' {
			comment, _ := lexer.Comment()
			server.AddProductToken("(" + comment + ")")
		} else {
			var tok string
			tok, ParseException = lexer.GetString('/')
			if ParseException != nil {
				tok = lexer.GetRest()
				server.AddProductToken(tok)
				break
			} else {
				if tok[len(tok)-1] == '\n' {
					tok = strings.TrimSpace(tok)
				}
				server.AddProductToken(tok)
			}
		}
	}

	return server, nil
}
