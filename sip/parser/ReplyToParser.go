package parser

import (
	"gosips/core"
	"gosips/sip/header"
)

/** Parser for a list of RelpyTo headers.
*
*@version  JAIN-SIP-1.1
*
*@author Olivier Deruelle <deruelle@nist.gov>  <br/>
*@author M. Ranganathan <mranga@nist.gov>  <br/>
*<a href="{@docRoot}/uncopyright.html">This code is in the public domain.</a>
*
*@version 1.0
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
	//if (debug) dbg_enter("ReplyTo.parse");

	// try {
	this.HeaderName(TokenTypes_REPLY_TO)

	replyTo.SetHeaderName(core.SIPHeaderNames_REPLY_TO)

	this.AddressParametersParser.Parse(replyTo)

	return replyTo, nil
	// } finally {
	//     if (debug) dbg_leave("ReplyTo.parse");
	// }

}
