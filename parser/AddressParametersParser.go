package parser

import (
	"gosip/header"
	"gosip/core"
)


/** Address parameters parser.
*/
type AddressParametersParser struct{
 	ParametersParser
}

	func NewAddressParametersParserFromLexer(lexer core.Lexer) *AddressParametersParser {
		this := &AddressParametersParser{};
		this.ParametersParser.superFromLexer(lexer);
		return this;
	}

	func NewAddressParametersParser(buffer string) *AddressParametersParser {
		this := &AddressParametersParser{}
		this.ParametersParser.super(buffer);
		return this;
	}
	
	func (this *AddressParametersParser) super(buffer string){
		this.ParametersParser.super(buffer);
	}

	func (this *AddressParametersParser) superFromLexer(lexer core.Lexer){
		this.ParametersParser.superFromLexer(lexer);
	}
	
	func (this *AddressParametersParser) Parse(addressParametersHeader header.AddressParametersHeader) (ParseException error) {
		//dbg_enter("AddressParametersParser.parse");
		//try {
		addressParser := NewAddressParserFromLexer(this.GetLexer());
		addr,_ := addressParser.Address();
		addressParametersHeader.SetAddress(addr);
		this.ParametersParser.Parse(addressParametersHeader);
		//} finally {
		//   dbg_leave("AddressParametersParser.parse");
		//}
		
		return nil;
	}

