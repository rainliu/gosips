package parser

import (
	"gosips/core"
	"gosips/sip/header"
	"strconv"
)

/**
* Accept-Encoding SIP (HTTP) Header parser.
*
* <pre>
*
*   The Accept-Encoding request-header field is similar to Accept, but
*   restricts the content-codings (section 3.5) that are acceptable in
*   the response.
*
*
*       Accept-Encoding  = "Accept-Encoding" ":"
*                      ( encoding *( "," encoding) )
*       encoding         = ( codings *[ ";" "q" "=" qvalue ] )
*       codings          = ( content-coding | "*" )
*
*   Examples of its use are:
*
*       Accept-Encoding: compress, gzip
*       Accept-Encoding:
*       Accept-Encoding: *
*       Accept-Encoding: compress;q=0.5, gzip;q=1.0
*       Accept-Encoding: gzip;q=1.0, identity; q=0.5, *;q=0
* </pre>
*
 */
type AcceptEncodingParser struct {
	HeaderParser
}

/** Constructor
 * @param String AcceptEncoding message to parse
 */
func NewAcceptEncodingParser(acceptEncoding string) *AcceptEncodingParser {
	this := &AcceptEncodingParser{}
	this.HeaderParser.super(acceptEncoding)
	return this
}

/** Cosntructor
 * @param lexer to set
 */
func NewAcceptEncodingParserFromLexer(lexer core.Lexer) *AcceptEncodingParser {
	this := &AcceptEncodingParser{}
	this.HeaderParser.superFromLexer(lexer)
	return this
}

/** parse the String message
 * @return Header (AcceptEncoding object)
 * @throws ParseException if the message does not respect the spec.
 */
func (this *AcceptEncodingParser) Parse() (sh header.Header, ParseException error) {
	acceptEncodingList := header.NewAcceptEncodingList()

	var ch byte
	lexer := this.GetLexer()
	this.HeaderName(TokenTypes_ACCEPT_ENCODING)

	// empty body is fine for this header.
	if ch, _ = lexer.LookAheadK(0); ch == '\n' {
		acceptEncoding := header.NewAcceptEncoding()
		acceptEncodingList.PushBack(acceptEncoding)
	} else {
		for ch, _ = lexer.LookAheadK(0); ch != '\n'; ch, _ = lexer.LookAheadK(0) {
			acceptEncoding := header.NewAcceptEncoding()
			if ch, _ = lexer.LookAheadK(0); ch != ';' { // Content-Coding:
				lexer.Match(TokenTypes_ID)
				value := lexer.GetNextToken()
				acceptEncoding.SetEncoding(value.GetTokenValue())
			}

			for ch, _ = lexer.LookAheadK(0); ch == ';'; ch, _ = lexer.LookAheadK(0) {
				lexer.Match(';')
				lexer.SPorHT()
				lexer.Match('q')
				lexer.SPorHT()
				lexer.Match('=')
				lexer.SPorHT()
				lexer.Match(TokenTypes_ID)
				value := lexer.GetNextToken()

				var qv float64
				if qv, ParseException = strconv.ParseFloat(value.GetTokenValue(), 32); ParseException != nil {
					return nil, ParseException
				}
				if ParseException = acceptEncoding.SetQValue(float32(qv)); ParseException != nil {
					return nil, ParseException
				}
				lexer.SPorHT()
			}

			acceptEncodingList.PushBack(acceptEncoding)
			if ch, _ = lexer.LookAheadK(0); ch == ',' {
				lexer.Match(',')
				lexer.SPorHT()
			}

		}
	}
	return acceptEncodingList, nil
}
