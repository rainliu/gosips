package parser

import (
	"gosips/core"
	"gosips/sip/address"
	"strings"
)

/** Parser for addresses.
 */

type AddressParser struct {
	Parser
}

func NewAddressParser(addr string) *AddressParser {
	this := &AddressParser{}

	this.Parser.super(addr)
	this.Parser.GetLexer().SetLexerName("charLexer")

	return this
}

func NewAddressParserFromLexer(lexer core.Lexer) *AddressParser {
	this := &AddressParser{}

	this.SetLexer(lexer)
	this.GetLexer().SelectLexer("charLexer")

	return this
}

func (this *AddressParser) super(addr string) {
	this.Parser.super(addr)
	this.Parser.GetLexer().SetLexerName("charLexer")
}

func (this *AddressParser) superFromLexer(lexer core.Lexer) {
	this.SetLexer(lexer)
	this.GetLexer().SelectLexer("charLexer")
}

func (this *AddressParser) NameAddr() (addr *address.AddressImpl, ParseException error) {
	//if (debug) dbg_enter("nameAddr");
	//try {
	var ch byte
	//var err error;
	lexer := this.GetLexer()

	//println("AddressParser::NameAddr():" + lexer.GetRest())

	if ch, _ = lexer.LookAheadK(0); ch == '<' {
		lexer.Match('<')
		lexer.SelectLexer("sip_urlLexer")
		lexer.SPorHT()
		uriParser := NewURLParserFromLexer(lexer)
		uri, _ := uriParser.UriReference()
		addr = address.NewAddressImpl()
		addr.SetAddressType(address.NAME_ADDR)
		addr.SetURI(uri)
		lexer.SPorHT()
		lexer.Match('>')
		return addr, nil
	} else {
		addr = address.NewAddressImpl()
		addr.SetAddressType(address.NAME_ADDR)
		var name string
		if ch, _ = lexer.LookAheadK(0); ch == '"' {
			name, _ = lexer.QuotedString()
			lexer.SPorHT()
		} else {
			name, _ = lexer.GetNextTokenByDelim('<')
		}
		addr.SetDisplayName(strings.TrimSpace(name))
		lexer.Match('<')
		lexer.SPorHT()
		uriParser := NewURLParserFromLexer(lexer)
		uri, _ := uriParser.UriReference()
		addr.SetAddressType(address.NAME_ADDR)
		addr.SetURI(uri)
		lexer.SPorHT()
		lexer.Match('>')
		return addr, nil
	}
	// } finally {
	//if (debug) dbg_leave("nameAddr");
	//}

}

func (this *AddressParser) Address() (retval *address.AddressImpl, ParseException error) {
	//if (debug) dbg_enter("address");
	//AddressImpl retval = null;
	//try {
	var ch byte
	//var err error
	lexer := this.GetLexer()
	k := 0
	//println(lexer.GetRest())
	for lexer.HasMoreChars() {
		if ch, ParseException = lexer.LookAheadK(k); ch == '<' ||
			ch == '"' ||
			ch == ':' ||
			ch == '/' {
			break
		} else if ParseException != nil /*ch == 0*/ { //'\0'
			return nil, this.CreateParseException("unexpected EOL")
		} else {
			k++
		}
	}
	//println(lexer.GetRest())
	if ch, _ = lexer.LookAheadK(k); ch == '<' ||
		ch == '"' {
		retval, _ = this.NameAddr()
	} else if ch == ':' ||
		ch == '/' {
		retval = address.NewAddressImpl()
		uriParser := NewURLParserFromLexer(lexer)
		uri, _ := uriParser.UriReference()
		retval.SetAddressType(address.ADDRESS_SPEC)
		retval.SetURI(uri)
	} else {
		return nil, this.CreateParseException("Bad address spec")
	}
	//println(lexer.GetRest());
	return retval, nil
	// } finally {
	//if (debug) dbg_leave("address");
	// }

}
