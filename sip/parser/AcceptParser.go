package parser

import (
	"gosips/core"
	"gosips/sip/header"
)

/** SIPParser for Accept header.
*
*<a href="{@docRoot}/uncopyright.html">This code is in the public domain.</a>
*
 */
type AcceptParser struct {
	ParametersParser
}

/**
 * Creates a new instance of Accept SIPParser
 * @param accept  the header to parse
 */
func NewAcceptParser(accept string) *AcceptParser {
	this := &AcceptParser{}
	this.ParametersParser.super(accept)
	return this
}

/** Constructor
 * @param lexer the lexer to use to parse the header
 */
func NewAcceptParserFromLexer(lexer core.Lexer) *AcceptParser {
	this := &AcceptParser{}
	this.ParametersParser.superFromLexer(lexer)
	return this
}

func (this *AcceptParser) super(buffer string) {
	this.ParametersParser.super(buffer)
}

func (this *AcceptParser) superFromLexer(lexer core.Lexer) {
	this.ParametersParser.superFromLexer(lexer)
}

/** parse the Accept  String header
 * @return Header (AcceptList  object)
 * @throws SIPParseException if the message does not respect the spec.
 */
func (this *AcceptParser) Parse() (sh header.Header, ParseException error) {

	//if (debug) dbg_enter("AcceptParser.parse");
	acceptList := header.NewAcceptList()

	//try {
	var ch byte
	lexer := this.GetLexer()
	this.HeaderName(TokenTypes_ACCEPT)

	accept := header.NewAccept()
	accept.SetHeaderName(core.SIPHeaderNames_ACCEPT)

	lexer.SPorHT()
	lexer.Match(TokenTypes_ID)
	token := lexer.GetNextToken()
	accept.SetContentType(token.GetTokenValue())
	lexer.Match('/')
	lexer.Match(TokenTypes_ID)
	token = lexer.GetNextToken()
	accept.SetContentSubType(token.GetTokenValue())
	lexer.SPorHT()

	this.ParametersParser.Parse(accept)
	acceptList.PushBack(accept)

	for ch, _ = lexer.LookAheadK(0); ch == ','; ch, _ = lexer.LookAheadK(0) {
		lexer.Match(',')
		lexer.SPorHT()

		accept = header.NewAccept()

		lexer.Match(TokenTypes_ID)
		token = lexer.GetNextToken()
		accept.SetContentType(token.GetTokenValue())
		lexer.Match('/')
		lexer.Match(TokenTypes_ID)
		token = lexer.GetNextToken()
		accept.SetContentSubType(token.GetTokenValue())
		lexer.SPorHT()
		this.ParametersParser.Parse(accept)
		acceptList.PushBack(accept)

	}
	return acceptList, nil
	//        }
	//        finally {
	//            if (debug) dbg_leave("AcceptParser.parse");
	//        }
}
