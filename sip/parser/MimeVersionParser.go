package parser

import (
	"gosips/core"
	"gosips/sip/header"
)

/** SIPParser for MimeVersion header.
 */
type MimeVersionParser struct {
	HeaderParser
}

/** Creates a new instance of MimeVersionParser
 * @param mimeVersion the header to parse
 */
func NewMimeVersionParser(mimeVersion string) *MimeVersionParser {
	this := &MimeVersionParser{}
	this.HeaderParser.super(mimeVersion)
	return this
}

/** Cosntructor
 * @param lexer the lexer to use to parse the header
 */
func NewMimeVersionParserFromLexer(lexer core.Lexer) *MimeVersionParser {
	this := &MimeVersionParser{}
	this.HeaderParser.superFromLexer(lexer)
	return this
}

/** parse the String message
 * @return Header (MimeVersion object)
 * @throws SIPParseException if the message does not respect the spec.
 */
func (this *MimeVersionParser) Parse() (sh header.Header, ParseException error) {
	mimeVersion := header.NewMimeVersion()

	lexer := this.GetLexer()
	this.HeaderName(TokenTypes_MIME_VERSION)

	mimeVersion.SetHeaderName(core.SIPHeaderNames_MIME_VERSION)

	var majorVersion, minorVersion int
	if majorVersion, ParseException = lexer.Number(); ParseException != nil {
		return nil, ParseException
	}
	mimeVersion.SetMajorVersion(majorVersion)
	lexer.Match('.')
	if minorVersion, ParseException = lexer.Number(); ParseException != nil {
		return nil, ParseException
	}
	mimeVersion.SetMinorVersion(minorVersion)

	lexer.SPorHT()

	lexer.Match('\n')

	return mimeVersion, nil
}
