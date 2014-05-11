package parser

import (
	"gosips/core"
	"gosips/sip/header"
	"strings"
	"time"
)

/** SIPParser for SIP Date field. Converts from SIP Date to the
 * internal storage (Calendar)
 */
type DateParser struct {
	HeaderParser
}

/** Constructor
 * @param String route message to parse to set
 */
func NewDateParser(date string) *DateParser {
	this := &DateParser{}
	this.HeaderParser.super(date)
	return this
}

func NewDateParserFromLexer(lexer core.Lexer) *DateParser {
	this := &DateParser{}
	this.HeaderParser.superFromLexer(lexer)
	return this
}

/** Parse method.
 * @throws ParseException
 * @return  the parsed Date header/
 */
func (this *DateParser) Parse() (sh header.Header, ParseException error) {
	//if (debug) dbg_enter("DateParser.parse");
	//try {
	lexer := this.GetLexer()
	this.HeaderName(TokenTypes_DATE)
	t, err := time.Parse(time.RFC1123, strings.TrimSpace(lexer.GetRest()))
	retval := header.NewDate()
	retval.SetDate(&t)
	return retval, err
	//           } finally {
	// if (debug) dbg_leave("DateParser.parse");

	//           }

}
