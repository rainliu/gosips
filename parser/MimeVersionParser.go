package parser

import (
	"gosip/core"
	"gosip/header"
)

/** Parser for MimeVersion header.
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
type MimeVersionParser struct {
	HeaderParserImpl
}

/** Creates a new instance of MimeVersionParser
 * @param mimeVersion the header to parse
 */
func NewMimeVersionParser(mimeVersion string) *MimeVersionParser {
	this := &MimeVersionParser{}
	this.HeaderParserImpl.super(mimeVersion)
	return this
}

/** Cosntructor
 * @param lexer the lexer to use to parse the header
 */
func NewMimeVersionParserFromLexer(lexer core.Lexer) *MimeVersionParser {
	this := &MimeVersionParser{}
	this.HeaderParserImpl.superFromLexer(lexer)
	return this
}

/** parse the String message
 * @return SIPHeaderHeader (MimeVersion object)
 * @throws SIPParseException if the message does not respect the spec.
 */
func (this *MimeVersionParser) Parse() (sh header.SIPHeaderHeader, ParseException error) {

	// if (debug) dbg_enter("MimeVersionParser.parse");
	mimeVersion := header.NewMimeVersion()
	// try {
	lexer := this.GetLexer()
	this.HeaderName(TokenTypes_MIME_VERSION)

	mimeVersion.SetHeaderName(core.SIPHeaderNames_MIME_VERSION)

	// try{
	majorVersion, _ := lexer.Number()
	mimeVersion.SetMajorVersion(majorVersion)
	lexer.Match('.')
	minorVersion, _ := lexer.Number()
	mimeVersion.SetMinorVersion(minorVersion)

	// }
	// catch (InvalidArgumentException ex) {
	//         throw createParseException(ex.getMessage());
	// }
	lexer.SPorHT()

	lexer.Match('\n')

	return mimeVersion, nil
	// }
	// finally {
	//     if (debug) dbg_leave("MimeVersionParser.parse");
	// }
}
