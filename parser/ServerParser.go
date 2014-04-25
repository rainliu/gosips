package parser

import (
	"errors"
	"gosip/core"
	"gosip/header"
	"strings"
)

/** Parser for Server header.
*
*@version  JAIN-SIP-1.1
*
*@author Olivier Deruelle <deruelle@nist.gov>  <br/>
*@author M. Ranganathan <mranga@nist.gov>  <br/>
*
*<a href="{@docRoot}/uncopyright.html">This code is in the public domain.</a>
*
* @version 1.0
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
func (this *ServerParser) Parse() (sh header.SIPHeaderHeader, ParseException error) {

	// if (debug) dbg_enter("ServerParser.parse");
	server := header.NewServer()
	// try {
	var ch byte
	lexer := this.GetLexer()
	this.HeaderName(TokenTypes_SERVER)
	if ch, _ = lexer.LookAheadK(0); ch == '\n' {
		return nil, errors.New("empty header")
	}
	//  mandatory token: product[/product-version] | (comment)
	for ch, _ = lexer.LookAheadK(0); ch != '\n'; ch, _ = lexer.LookAheadK(0) {
		//while (this.lexer.lookAhead(0) != '\n'
		//&& this.lexer.lookAhead(0) != '\0') {
		if ch, _ = lexer.LookAheadK(0); ch == '(' {
			comment, _ := lexer.Comment()
			server.AddProductToken("(" + comment + ")")
		} else {
			// String tok;
			//try {
			tok, err := lexer.GetString('/')
			if err != nil {
				tok = lexer.GetRest()
				server.AddProductToken(tok)
				break
			} else {
				if tok[len(tok)-1] == '\n' {
					tok = strings.TrimSpace(tok)
				}
				server.AddProductToken(tok)
			}
			//    } catch (ParseException ex) {
			// tok = this.lexer.getRest();
			// server.addProductToken(tok);
			// break;
			//    }
		}
	}

	// }
	// finally {
	//     if (debug) dbg_leave("ServerParser.parse");
	// }

	return server, nil
}
