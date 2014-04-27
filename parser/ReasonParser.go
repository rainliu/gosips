package parser

import (
	"gosip/core"
	"gosip/header"
)

/** Parser for Reason header.
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
type ReasonParser struct {
	ParametersParser
}

/** Creates a new instance of ReasonParser
 * @param reason the header to parse
 */
func NewReasonParser(reason string) *ReasonParser {
	this := &ReasonParser{}
	this.ParametersParser.super(reason)
	return this
}

/** Constructor
 * @param lexer the lexer to use to parse the header
 */
func NewReasonParserFromLexer(lexer core.Lexer) *ReasonParser {
	this := &ReasonParser{}
	this.ParametersParser.superFromLexer(lexer)
	return this
}

/** parse the String message
 * @return SIPHeader (ReasonParserList object)
 * @throws SIPParseException if the message does not respect the spec.
 */
func (this *ReasonParser) Parse() (sh header.ISIPHeader, ParseException error) {
	reasonList := header.NewReasonList()
	//if (debug) dbg_enter("ReasonParser.parse");

	// try {
	var ch byte
	lexer := this.GetLexer()
	this.HeaderName(TokenTypes_REASON)
	lexer.SPorHT()
	for ch, _ = lexer.LookAheadK(0); ch != '\n'; ch, _ = lexer.LookAheadK(0) {
		reason := header.NewReason()
		lexer.Match(TokenTypes_ID)
		token := lexer.GetNextToken()
		value := token.GetTokenValue()

		reason.SetProtocol(value)
		this.ParametersParser.Parse(reason)
		reasonList.PushBack(reason)
		if ch, _ = lexer.LookAheadK(0); ch == ',' {
			lexer.Match(',')
			lexer.SPorHT()
		} else {
			lexer.SPorHT()
		}

	}
	// } catch (ParseException ex ) {
	// 	ex.printStackTrace();
	// 	System.out.println(lexer.getRest());
	//        } finally {
	//            if (debug) dbg_leave("ReasonParser.parse");
	//        }

	return reasonList, nil
}
