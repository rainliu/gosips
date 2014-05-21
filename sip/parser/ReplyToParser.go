package parser

import (
	"gosips/core"
	"gosips/sip/header"
)

/** SIPParser for a list of RelpyTo headers.
 */

type ReplyToParser struct {
	AddressParametersParser
}

/** Creates a new instance of ReplyToParser
 * @param replyTo the header to parse
 */
func NewReplyToParser(replyTo string) *ReplyToParser {
	this := &ReplyToParser{}
	this.AddressParametersParser.super(replyTo)
	return this
}

/** Cosntructor
 * param lexer the lexer to use to parse the header
 */
func NewReplyToParserFromLexer(lexer core.Lexer) *ReplyToParser {
	this := &ReplyToParser{}
	this.AddressParametersParser.superFromLexer(lexer)
	return this
}

/** parse the String message and generate the ReplyTo List Object
 * @return SIPHeader the ReplyTo List object
 * @throws SIPParseException if errors occur during the parsing
 */
func (this *ReplyToParser) Parse() (sh header.Header, ParseException error) {
	replyTo := header.NewReplyTo()

	this.HeaderName(TokenTypes_REPLY_TO)

	replyTo.SetHeaderName(core.SIPHeaderNames_REPLY_TO)

	ParseException = this.AddressParametersParser.Parse(replyTo)

	return replyTo, ParseException
}
