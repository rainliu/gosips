package parser

import (
	"gosips/core"
	"gosips/sip/header"
	"strings"
)

/** SIPParser for Organization header.
 */
type OrganizationParser struct {
	HeaderParser
}

/**  Creates a new instance of OrganizationParser
 * @param organization the header to parse
 */
func NewOrganizationParser(organization string) *OrganizationParser {
	this := &OrganizationParser{}
	this.HeaderParser.super(organization)
	return this
}

/** Constructor
 * @param lexer the lexer to use to parse the header
 */
func NewOrganizationParserFromLexer(lexer core.Lexer) *OrganizationParser {
	this := &OrganizationParser{}
	this.HeaderParser.superFromLexer(lexer)
	return this
}

/** parse the String header
 * @return Header (Organization object)
 * @throws SIPParseException if the message does not respect the spec.
 */
func (this *OrganizationParser) Parse() (sh header.Header, ParseException error) {
	organization := header.NewOrganization()

	lexer := this.GetLexer()
	this.HeaderName(TokenTypes_ORGANIZATION)

	organization.SetHeaderName(core.SIPHeaderNames_ORGANIZATION)

	lexer.SPorHT()
	value := lexer.GetRest()

	organization.SetOrganization(strings.TrimSpace(value))

	return organization, nil
}
