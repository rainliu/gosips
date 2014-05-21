package parser

import (
	"bytes"
	"gosips/core"
	"gosips/sip/header"
	"strings"
)

/** SIPParser for via headers.
 */
type ViaParser struct {
	HeaderParser
}

func NewViaParser(via string) *ViaParser {
	this := &ViaParser{}
	this.HeaderParser.super(via)
	return this
}

func NewViaParserFromLexer(lexer core.Lexer) *ViaParser {
	this := &ViaParser{}
	this.HeaderParser.superFromLexer(lexer)
	return this
}

/**  a parser for the essential part of the via header.
 */
func (this *ViaParser) ParseVia(v *header.Via) (ParseException error) {
	lexer := this.GetLexer()

	var protocolName, protocolVersion *core.Token
	// The protocol
	lexer.Match(TokenTypes_ID)

	if protocolName = lexer.GetNextToken(); protocolName.GetTokenValue() != "SIP" {
		return this.CreateParseException("Protcoal Not Supported error")
	}

	lexer.SPorHT()
	// consume the "/"
	lexer.Match('/')
	lexer.SPorHT()
	lexer.Match(TokenTypes_ID)
	lexer.SPorHT()
	if protocolVersion = lexer.GetNextToken(); protocolVersion.GetTokenValue() != "2.0" {
		return this.CreateParseException("Version Not Supported error")
	}

	lexer.SPorHT()

	// We consume the "/"
	lexer.Match('/')
	lexer.SPorHT()
	lexer.Match(TokenTypes_ID)
	lexer.SPorHT()

	transport := lexer.GetNextToken()
	lexer.SPorHT()

	protocol := header.NewProtocol()
	protocol.SetProtocolName(protocolName.GetTokenValue())
	protocol.SetProtocolVersion(protocolVersion.GetTokenValue())
	protocol.SetTransport(transport.GetTokenValue())
	v.SetSentProtocol(protocol)

	// sent-By
	hnp := core.NewHostNameParserFromLexer(this.GetLexer())

	var hostPort *core.HostPort
	if hostPort, ParseException = hnp.GetHostPort(); ParseException != nil {
		return ParseException
	}
	v.SetSentBy(hostPort)

	// Ignore blanks
	lexer.SPorHT()

	var la byte
	for la, _ = lexer.LookAheadK(0); la == ';'; la, _ = lexer.LookAheadK(0) {
		lexer.Match(';')
		lexer.SPorHT()

		var nameValue *core.NameValue
		if nameValue, ParseException = this.NameValue(); ParseException != nil {
			return ParseException
		}

		name := nameValue.GetName()
		nameValue.SetName(strings.ToLower(name))
		v.SetParameterFromNameValue(nameValue)
		lexer.SPorHT()
	}

	if la, _ = lexer.LookAheadK(0); la == '(' {
		lexer.SelectLexer("charLexer")
		lexer.ConsumeK(1)

		var comment bytes.Buffer
		for {
			ch, _ := lexer.LookAheadK(0)
			if ch == ')' {
				lexer.ConsumeK(1)
				break
			} else if ch == '\\' {
				// Escaped character
				tok := lexer.GetNextToken()
				comment.WriteString(tok.GetTokenValue())
				lexer.ConsumeK(1)
				tok = lexer.GetNextToken()
				comment.WriteString(tok.GetTokenValue())
				lexer.ConsumeK(1)
			} else if ch == '\n' {
				break
			} else {
				comment.WriteByte(ch)
				lexer.ConsumeK(1)
			}
		}
		v.SetComment(comment.String())
	}

	return nil

}

/** Overrides the superclass nameValue parser because
* we have to tolerate IPV6 addresses in the received parameter.
 */
func (this *ViaParser) NameValue() (nv *core.NameValue, ParseException error) {
	lexer := this.GetLexer()
	lexer.Match(core.CORELEXER_ID)
	name := lexer.GetNextToken()

	// eat white space.
	lexer.SPorHT()

	quoted := false

	la, _ := lexer.LookAheadK(0)
	if la == '=' {
		lexer.ConsumeK(1)
		lexer.SPorHT()
		var str string
		if strings.ToLower(name.GetTokenValue()) == core.SIPParameters_RECEIVED {
			// Allow for IPV6 Addresses.
			// these could have : in them!
			str = lexer.ByteStringNoSemicolon()
		} else {
			if la, _ = lexer.LookAheadK(0); la == '"' {
				if str, ParseException = lexer.QuotedString(); ParseException != nil {
					return nil, ParseException
				}
				quoted = true
			} else {
				lexer.Match(core.CORELEXER_ID)
				value := lexer.GetNextToken()
				str = value.GetTokenValue()
			}
		}
		nv := core.NewNameValue(name.GetTokenValue(), str)
		if quoted {
			nv.SetQuotedValue()
		}
		return nv, nil
	} else {
		return core.NewNameValue(name.GetTokenValue(), nil), nil
	}
}

func (this *ViaParser) Parse() (sh header.Header, ParseException error) {
	lexer := this.GetLexer()

	viaList := header.NewViaList()

	// The first via header.
	lexer.Match(TokenTypes_VIA)
	lexer.SPorHT()   // ignore blanks
	lexer.Match(':') // expect a colon.
	lexer.SPorHT()   // ingore blanks.

	for {
		v := header.NewVia()
		if ParseException = this.ParseVia(v); ParseException != nil {
			return nil, ParseException
		}
		viaList.PushBack(v)
		lexer.SPorHT() // eat whitespace.
		if la, _ := lexer.LookAheadK(0); la == ',' {
			lexer.ConsumeK(1) // Consume the comma
			lexer.SPorHT()    // Ignore space after.
		}
		if la, _ := lexer.LookAheadK(0); la == '\n' {
			break
		}
	}
	lexer.Match('\n')

	return viaList, nil
}
