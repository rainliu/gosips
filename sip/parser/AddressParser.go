package parser

import (
	"gosips/core"
	"gosips/sip/address"
	"strings"
)

/** SIPParser for addresses.
 */
type AddressParser struct {
	SIPParser
}

func NewAddressParser(addr string) *AddressParser {
	this := &AddressParser{}

	this.SIPParser.super(addr)
	this.SIPParser.GetLexer().SetLexerName("charLexer")

	return this
}

func NewAddressParserFromLexer(lexer core.Lexer) *AddressParser {
	this := &AddressParser{}

	this.SetLexer(lexer)
	this.GetLexer().SelectLexer("charLexer")

	return this
}

func (this *AddressParser) super(addr string) {
	this.SIPParser.super(addr)
	this.SIPParser.GetLexer().SetLexerName("charLexer")
}

func (this *AddressParser) superFromLexer(lexer core.Lexer) {
	this.SetLexer(lexer)
	this.GetLexer().SelectLexer("charLexer")
}

func (this *AddressParser) NameAddr() (addr *address.AddressImpl, ParseException error) {
	var ch byte
	var uri address.URI

	lexer := this.GetLexer()

	if ch, _ = lexer.LookAheadK(0); ch == '<' {
		lexer.Match('<')
		lexer.SelectLexer("sip_urlLexer")
		lexer.SPorHT()
		uriParser := NewURLParserFromLexer(lexer)
		if uri, ParseException = uriParser.UriReference(); ParseException != nil {
			return nil, ParseException
		}
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
			if name, ParseException = lexer.QuotedString(); ParseException != nil {
				return nil, ParseException
			}
			lexer.SPorHT()
		} else {
			if name, ParseException = lexer.GetNextTokenByDelim('<'); ParseException != nil {
				return nil, ParseException
			}
		}
		addr.SetDisplayName(strings.TrimSpace(name))
		lexer.Match('<')
		lexer.SPorHT()
		uriParser := NewURLParserFromLexer(lexer)
		if uri, ParseException = uriParser.UriReference(); ParseException != nil {
			return nil, ParseException
		}
		addr.SetAddressType(address.NAME_ADDR)
		addr.SetURI(uri)
		lexer.SPorHT()
		lexer.Match('>')
		return addr, nil
	}
}

func (this *AddressParser) Address() (retval *address.AddressImpl, ParseException error) {
	var ch byte
	lexer := this.GetLexer()
	k := 0
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

	if ch, _ = lexer.LookAheadK(k); ch == '<' ||
		ch == '"' {
		if retval, ParseException = this.NameAddr(); ParseException != nil {
			return nil, ParseException
		}
	} else if ch == ':' ||
		ch == '/' {
		retval = address.NewAddressImpl()
		uriParser := NewURLParserFromLexer(lexer)

		var uri address.URI
		if uri, ParseException = uriParser.UriReference(); ParseException != nil {
			return nil, ParseException
		}
		retval.SetAddressType(address.ADDRESS_SPEC)
		retval.SetURI(uri)
	} else {
		return nil, this.CreateParseException("Bad address spec")
	}

	return retval, nil
}
