package parser

import (
	"bytes"
	"gosip/core"
	"gosip/header"
	"strings"
)

/** Parser for via headers.
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
	// The protocol
	lexer.Match(TokenTypes_ID)
	protocolName := lexer.GetNextToken()

	lexer.SPorHT()
	// consume the "/"
	lexer.Match('/')
	lexer.SPorHT()
	lexer.Match(TokenTypes_ID)
	lexer.SPorHT()
	protocolVersion := lexer.GetNextToken()

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
	hostPort, _ := hnp.GetHostPort()
	v.SetSentBy(hostPort)

	// Ignore blanks
	lexer.SPorHT()

	var la byte
	// parameters
	la, _ = lexer.LookAheadK(0)
	for la == ';' {
		lexer.Match(';')
		lexer.SPorHT()
		nameValue, _ := this.NameValue()
		name := nameValue.GetName()
		nameValue.SetName(strings.ToLower(name))
		v.SetParameter(nameValue.GetName(), nameValue.GetValue().(string))
		lexer.SPorHT()
		la, _ = lexer.LookAheadK(0)
	}

	if la, _ = lexer.LookAheadK(0); la == '(' {
		lexer.SelectLexer("charLexer")
		lexer.ConsumeK(1)
		var comment bytes.Buffer //=new StringBuffer();
		//cond:=true;
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
	//if (debug) dbg_enter("nameValue");
	//try {
	lexer := this.GetLexer()
	lexer.Match(core.LexerCore_ID)
	name := lexer.GetNextToken()
	// eat white space.
	lexer.SPorHT()
	//try {

	quoted := false

	la, _ := lexer.LookAheadK(0)

	if la == '=' {
		lexer.ConsumeK(1)
		lexer.SPorHT()
		var str string
		if strings.ToLower(name.GetTokenValue()) != core.SIPParameters_RECEIVED { //bug?
			// Allow for IPV6 Addresses.
			// these could have : in them!
			str = lexer.ByteStringNoSemicolon()
		} else {
			if la, _ = lexer.LookAheadK(0); la == '"' {
				str, _ = lexer.QuotedString()
				quoted = true
			} else {
				lexer.Match(core.LexerCore_ID)
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

	/*} catch (ParseException ex) {
	          return new NameValue(name.getTokenValue(),null);
	  }

	  } finally {
	          if (debug) dbg_leave("nameValue");
	  }*/
}

func (this *ViaParser) Parse() (sh header.SIPHeaderHeader, ParseException error) {
	//  if (debug) dbg_enter("parse");
	//  try {
	lexer := this.GetLexer()

	viaList := header.NewViaList()
	// The first via header.
	lexer.Match(TokenTypes_VIA)
	lexer.SPorHT()   // ignore blanks
	lexer.Match(':') // expect a colon.
	lexer.SPorHT()   // ingore blanks.

	for {
		v := header.NewVia()
		this.ParseVia(v)
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
	// } finally {
	//if (debug) dbg_leave("parse");
	// }

}
