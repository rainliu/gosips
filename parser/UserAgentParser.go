package parser

import (
	"errors"
	"gosip/core"
	"gosip/header"
	"strings"
)

/** Parser for UserAgent header.
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
func (this *UserAgentParser) Parse() (sh header.SIPHeaderHeader, ParseException error) {

	// if (debug) dbg_enter("UserAgentParser.parse");
	userAgent := header.NewUserAgent()
	//try {
	var ch byte
	lexer := this.GetLexer()
	this.HeaderName(TokenTypes_USER_AGENT)
	if ch, _ = lexer.LookAheadK(0); ch == '\n' {
		return nil, errors.New("empty header")
	}

	//  mandatory token: product[/product-version] | (comment)
	for ch, _ = lexer.LookAheadK(0); ch != '\n'; ch, _ = lexer.LookAheadK(0) {
		//while (this.lexer.lookAhead(0) != '\n'
		//	&& this.lexer.lookAhead(0) != '\0') {
		if ch, _ = lexer.LookAheadK(0); ch == '(' {
			comment, _ := lexer.Comment()
			userAgent.AddProductToken("(" + comment + ")")
		} else {
			//String tok;
			//try {
			tok, err := lexer.GetString('/')
			if err != nil {
				tok = lexer.GetRest()
				userAgent.AddProductToken(strings.TrimSpace(tok))
				break
			} else {
				if tok[len(tok)-1] == '\n' {
					tok = strings.TrimSpace(tok)
				}
				userAgent.AddProductToken(tok)
			}

			//    } catch (ParseException ex) {
			// tok = this.lexer.getRest();
			// userAgent.addProductToken(tok.trim());
			// break;
			//    }
		}
	}
	// }
	// finally {
	//     if (debug) dbg_leave("UserAgentParser.parse");
	// }

	return userAgent, nil
}
