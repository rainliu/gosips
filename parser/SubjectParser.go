package parser

import (
	"gosip/core"
	"gosip/header"
	"strings"
)

/** Parser for Subject  header.
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
	//if (debug) dbg_enter("SubjectParser.parse");

	// try {
	lexer := this.GetLexer()
	this.HeaderName(TokenTypes_SUBJECT)

	lexer.SPorHT()

	s := lexer.GetRest()
	subject.SetSubject(strings.TrimSpace(s))

	// } finally {
	//     if (debug) dbg_leave("SubjectParser.parse");
	// }

	return subject, nil
}
