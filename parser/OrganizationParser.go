package parser

import (
	"gosip/core"
	"gosip/header"
	"strings"
)

/** Parser for Organization header.
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
type OrganizationParser struct {
	HeaderParserImpl
}

/**  Creates a new instance of OrganizationParser
 * @param organization the header to parse
 */
func NewOrganizationParser(organization string) *OrganizationParser {
	this := &OrganizationParser{}
	this.HeaderParserImpl.super(organization)
	return this
}

/** Constructor
 * @param lexer the lexer to use to parse the header
 */
func NewOrganizationParserFromLexer(lexer core.Lexer) *OrganizationParser {
	this := &OrganizationParser{}
	this.HeaderParserImpl.superFromLexer(lexer)
	return this
}

/** parse the String header
 * @return SIPHeader (Organization object)
 * @throws SIPParseException if the message does not respect the spec.
 */
func (this *OrganizationParser) Parse() (sh header.SIPHeader, ParseException error) {

	//if (debug) dbg_enter("OrganizationParser.parse");
	organization := header.NewOrganization()
	// try {
	lexer := this.GetLexer()
	this.HeaderName(TokenTypes_ORGANIZATION)

	organization.SetHeaderName(core.SIPHeaderNames_ORGANIZATION)

	lexer.SPorHT()
	value := lexer.GetRest()

	organization.SetOrganization(strings.TrimSpace(value))

	return organization, nil
	// }
	// finally {
	//     if (debug) dbg_leave("OrganizationParser.parse");
	// }
}
