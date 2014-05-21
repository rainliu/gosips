package parser

import (
	"gosips/core"
	"gosips/sip/header"
	"strings"
)

/** SIPParser for Subject  header.
 */
type SubjectParser struct {
	HeaderParser
}

/** Creates a new instance of SubjectParser
 * @param subject the header to parse
 */
func NewSubjectParser(subject string) *SubjectParser {
	this := &SubjectParser{}
	this.HeaderParser.super(subject)
	return this
}

/** Cosntructor
 * @param lexer the lexer to use to parse the header
 */
func NewSubjectParserFromLexer(lexer core.Lexer) *SubjectParser {
	this := &SubjectParser{}
	this.HeaderParser.superFromLexer(lexer)
	return this
}

/** parse the String message
 * @return SIPHeader (Subject object)
 * @throws SIPParseException if the message does not respect the spec.
 */
func (this *SubjectParser) Parse() (sh header.Header, ParseException error) {
	subject := header.NewSubject()

	lexer := this.GetLexer()
	this.HeaderName(TokenTypes_SUBJECT)

	lexer.SPorHT()

	s := lexer.GetRest()
	subject.SetSubject(strings.TrimSpace(s))

	return subject, nil
}
