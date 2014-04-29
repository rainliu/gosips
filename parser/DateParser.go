package parser

import (
	"gosip/core"
	"gosip/header"
)

/** Parser for SIP Date field. Converts from SIP Date to the
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
func (this *DateParser) Parse() (sh header.ISIPHeader, ParseException error) {
	//if (debug) dbg_enter("DateParser.parse");
	//try {
	lexer := this.GetLexer()
	this.HeaderName(TokenTypes_DATE)
	this.Wkday()
	lexer.Match(',')
	lexer.Match(' ')
	cal, _ := this.Date()
	lexer.Match(' ')
	this.Time(cal)
	lexer.Match('\n')
	retval := header.NewDate()
	retval.SetDate(cal)
	return retval, nil
	//           } finally {
	// if (debug) dbg_leave("DateParser.parse");

	//           }

}

/**
        public static void main(String args[]) throws ParseException {
		String date[] = {
			"Date: Sun, 07 Jan 2001 19:05:06 GMT\n",
			"Date: Mon, 08 Jan 2001 19:05:06 GMT\n" };

		for (int i = 0; i < date.length; i++ ) {
		    System.out.println("Parsing " + date[i]);
		    DateParser dp =
			  new DateParser(date[i]);
		    SIPDateHeader d = (SIPDateHeader) dp.parse();
		    System.out.println("encoded = " +d.encode());
		}

	}
**/
