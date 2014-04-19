package parser

import (
	"gosip/core"
	"gosip/header"
)

/** Parser for Authentication-Info header.
 */
type AuthenticationInfoParser struct {
	ParametersParser
}

/**
 * Creates a new instance of AuthenticationInfoParser
 * @param authenticationInfo the header to parse
 */
func NewAuthenticationInfoParser(authenticationInfo string) *AuthenticationInfoParser {
	this := &AuthenticationInfoParser{}
	this.ParametersParser.super(authenticationInfo)
	return this
}

/** Constructor
 * @param lexer the lexer to use to parse the header
 */
func NewAuthenticationInfoParserFromLexer(lexer core.Lexer) *AuthenticationInfoParser {

	this := &AuthenticationInfoParser{}
	this.ParametersParser.superFromLexer(lexer)
	return this
}

/** parse the AuthenticationInfo String header
 * @return SIPHeader (AuthenticationInfoList object)
 * @throws SIPParseException if the message does not respect the spec.
 */
func (this *AuthenticationInfoParser) Parse() (sh header.SIPHeader, ParseException error) {

	// if (debug) dbg_enter("AuthenticationInfoParser.parse");

	//try {
	var ch byte

	lexer := this.GetLexer()
	this.HeaderName(TokenTypes_AUTHENTICATION_INFO)

	authenticationInfo := header.NewAuthenticationInfo()
	authenticationInfo.SetHeaderName(core.SIPHeaderNames_AUTHENTICATION_INFO)

	lexer.SPorHT()

	//println(lexer.GetRest())

	nv := this.NameValue('=')
	authenticationInfo.SetParameter(nv.GetName(), nv.GetValue().(string))
	// println(nv.GetName())
	// println(nv.GetValue().(string))
	// println(lexer.GetRest())

	lexer.SPorHT()
	for ch, _ = lexer.LookAheadK(0); ch == ','; ch, _ = lexer.LookAheadK(0) {
		lexer.Match(',')
		lexer.SPorHT()

		nv := this.NameValue('=')
		authenticationInfo.SetParameter(nv.GetName(), nv.GetValue().(string))

		lexer.SPorHT()
	}
	lexer.SPorHT()
	//lexer.Match('\n');

	return authenticationInfo, nil
	// }
	// finally {
	//     if (debug) dbg_leave("AuthenticationInfoParser.parse");
	// }
}
