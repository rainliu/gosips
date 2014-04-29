package parser

import (
	"errors"
	"gosip/core"
	"gosip/header"
	"strings"
	"time"
)

type IHeaderParser interface {
	Parse() (sh header.ISIPHeader, ParseException error)
}

/** Generic header parser class. The parsers for various headers extend this
* class. To create a parser for a new header, extend this class and change
* the createParser class.
 */

type HeaderParser struct {
	Parser
}

/** Creates new HeaderParser
 * @param String to parse.
 */
func NewHeaderParser(header string) *HeaderParser {
	this := &HeaderParser{}

	this.Parser.super(header)
	this.Parser.GetLexer().SetLexerName("command_keywordLexer")

	return this
}

func NewHeaderParserFromLexer(lexer core.Lexer) *HeaderParser {
	this := &HeaderParser{}

	this.Parser.SetLexer(lexer)
	this.Parser.GetLexer().SetLexerName("command_keywordLexer")

	return this
}

func (this *HeaderParser) super(header string) {
	this.Parser.super(header)
	this.Parser.GetLexer().SetLexerName("command_keywordLexer")
}

func (this *HeaderParser) superFromLexer(lexer core.Lexer) {
	this.SetLexer(lexer)
	this.Parser.GetLexer().SetLexerName("command_keywordLexer")
}

/** Parse the weekday field
 * @return an integer with the calendar content for wkday.
 */
func (this *HeaderParser) Wkday() (wk int, ParseException error) {
	this.Dbg_enter("wkday")
	defer this.Dbg_leave("wkday")
	//try {
	tok := this.GetLexer().Ttoken()
	id := strings.ToLower(tok)

	if strings.ToLower(core.SIPDateNames_MON) == (id) {
		return core.SIPCalendar_MONDAY, nil
	} else if strings.ToLower(core.SIPDateNames_TUE) == (id) {
		return core.SIPCalendar_TUESDAY, nil
	} else if strings.ToLower(core.SIPDateNames_WED) == (id) {
		return core.SIPCalendar_WEDNESDAY, nil
	} else if strings.ToLower(core.SIPDateNames_THU) == (id) {
		return core.SIPCalendar_THURSDAY, nil
	} else if strings.ToLower(core.SIPDateNames_FRI) == (id) {
		return core.SIPCalendar_FRIDAY, nil
	} else if strings.ToLower(core.SIPDateNames_SAT) == (id) {
		return core.SIPCalendar_SATURDAY, nil
	} else if strings.ToLower(core.SIPDateNames_SUN) == (id) {
		return core.SIPCalendar_SUNDAY, nil
	} else {
		return -1, this.CreateParseException("bad wkday")
	}
	//} finally {
	//	dbg_leave("wkday");
	//}

}

/** parse and return a date field.
 *@return a date structure with the parsed value.
 */
func (this *HeaderParser) Date() (t *time.Time, ParseException error) {
	//try  {
	//Calendar retval =  Calendar.getInstance(TimeZone.getTimeZone("GMT"));
	lexer := this.GetLexer()
	day, _ := lexer.Number()
	if day <= 0 || day >= 31 {
		return nil, errors.New("Bad day ")
	}
	lexer.Match(' ')
	month := strings.ToLower(lexer.Ttoken())
	var mon time.Month
	if month == "jan" {
		mon = time.January
	} else if month == "feb" {
		mon = time.February
	} else if month == "mar" {
		mon = time.March
	} else if month == "apr" {
		mon = time.April
	} else if month == "may" {
		mon = time.May
	} else if month == "jun" {
		mon = time.June
	} else if month == "jul" {
		mon = time.July
	} else if month == "aug" {
		mon = time.August
	} else if month == "sep" {
		mon = time.September
	} else if month == "oct" {
		mon = time.October
	} else if month == "nov" {
		mon = time.November
	} else if month == "dec" {
		mon = time.December
	}
	lexer.Match(' ')
	yr, _ := lexer.Number()

	t2 := time.Date(yr, mon, day, 0, 0, 0, 0, nil)
	return &t2, nil
	// } catch (Exception ex) {
	//     throw createParseException("bad date field" );
	// }

}

/** Set the time field. This has the format hour:minute:second
 */
func (this *HeaderParser) Time(t *time.Time) (c *time.Time, ParseException error) {
	//try {
	lexer := this.GetLexer()
	hour, _ := lexer.Number()
	lexer.Match(':')
	min, _ := lexer.Number()
	lexer.Match(':')
	sec, _ := lexer.Number()
	lexer.Match(' ')
	tzone := strings.ToLower(lexer.Ttoken())
	loc, _ := time.LoadLocation(tzone)

	t2 := time.Date(t.Year(), t.Month(), t.Day(), hour, min, sec, 0, loc)
	return &t2, nil
	// } catch (Exception ex) {
	//     throw createParseException ("error processing time " );
	// }
}

/** Parse the SIP header from the buffer and return a parsed
 * structure.
 *@throws ParseException if there was an error parsing.
 */
func (this *HeaderParser) Parse() (sh header.ISIPHeader, ParseException error) {
	lexer := this.GetLexer()

	name := lexer.GetNextTokenByDelim(':')
	lexer.ConsumeK(1)
	body := strings.TrimSpace(lexer.GetLine())
	// we dont set any fields because the header is
	// ok
	retval := header.NewExtension(name)
	retval.SetValue(body)
	return retval, nil

}

/** Parse the header name until the colon  and chew WS after that.
 */
func (this *HeaderParser) HeaderName(tok int) {
	this.GetLexer().Match(tok)
	this.GetLexer().SPorHT()
	this.GetLexer().Match(':')
	this.GetLexer().SPorHT()
}
