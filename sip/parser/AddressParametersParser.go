package parser

import (
	"gosips/core"
	"gosips/sip/address"
	"gosips/sip/header"
)

/** Address parameters parser.
 */
type AddressParametersParser struct {
	ParametersParser
}

func NewAddressParametersParserFromLexer(lexer core.Lexer) *AddressParametersParser {
	this := &AddressParametersParser{}
	this.ParametersParser.superFromLexer(lexer)
	return this
}

func NewAddressParametersParser(buffer string) *AddressParametersParser {
	this := &AddressParametersParser{}
	this.ParametersParser.super(buffer)
	return this
}

func (this *AddressParametersParser) super(buffer string) {
	this.ParametersParser.super(buffer)
}

func (this *AddressParametersParser) superFromLexer(lexer core.Lexer) {
	this.ParametersParser.superFromLexer(lexer)
}

func (this *AddressParametersParser) Parse(addressParametersHeader header.AddressParametersHeader) (ParseException error) {
	addressParser := NewAddressParserFromLexer(this.GetLexer())
	var addr *address.AddressImpl
	if addr, ParseException = addressParser.Address(); ParseException != nil {
		return ParseException
	}
	addressParametersHeader.SetAddress(addr)
	ParseException = this.ParametersParser.Parse(addressParametersHeader)
	return ParseException
}
