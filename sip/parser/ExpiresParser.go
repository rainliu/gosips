package parser

import (
	"gosips/core"
	"gosips/sip/header"
	"strconv"
)

/** SIPParser for SIP Expires SIPParser. Converts from SIP Date to the
* internal storage (Calendar).
 */
type ExpiresParser struct {
	HeaderParser
}

/** protected constructor.
 *@param text is the text of the header to parse
 */
func NewExpiresParser(text string) *ExpiresParser {
	this := &ExpiresParser{}
	this.HeaderParser.super(text)
	return this
}

/** constructor.
 *@param lexer is the lexer passed in from the enclosing parser.
 */
func NewExpiresParserFromLexer(lexer core.Lexer) *ExpiresParser {
	this := &ExpiresParser{}
	this.HeaderParser.superFromLexer(lexer)
	return this
}

/** Parse the header.
 */
func (this *ExpiresParser) Parse() (sh header.Header, ParseException error) {
	expires := header.NewExpires()

	lexer := this.GetLexer()
	lexer.Match(TokenTypes_EXPIRES)
	lexer.SPorHT()
	lexer.Match(':')
	lexer.SPorHT()
	nextId := lexer.GetNextId()
	lexer.Match('\n')
	var delta int64
	if delta, ParseException = strconv.ParseInt(nextId, 10, 32); ParseException != nil {
		return nil, ParseException
	}
	expires.SetExpires(int(delta))
	return expires, ParseException
}
