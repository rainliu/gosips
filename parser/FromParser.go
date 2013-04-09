package parser

import (
	"gosip/address"
	"gosip/header"
)

/** From header parser.
*/

type FromParser struct{
 	AddressParametersParser
}
    
	func NewFromParser(from string) *FromParser {
		this := &FromParser{};
		this.AddressParametersParser.super(from);
		
		return this;
	}
	
	func (this *FromParser) super(from string){
		this.AddressParametersParser.super(from);
	}
	
	

    /*protected FromParser(Lexer lexer) {
		super(lexer);
	}*/
	
	func (this *FromParser) Parse() (sh header.SIPHeader, ParseException error) {
	 	from := header.NewFrom();
		lexer := this.GetLexer()
		lexer.Match (TokenTypes_FROM);
		lexer.SPorHT();
		lexer.Match(':');
		lexer.SPorHT();
		this.AddressParametersParser.Parse(from);
		lexer.Match('\n');
		if from.GetAddress().GetAddressType() == address.ADDRESS_SPEC {
			// the parameters are header parameters.
			if from.GetAddress().GetURI().IsSipURI() {
			  sipUri,_ := from.GetAddress().GetURI().(*address.SipUri);
			  parms := sipUri.GetUriParms();
			  if parms != nil && parms.Len()>0 {
			     from.SetParameters(parms);
			     sipUri.RemoveUriParms();
			  }
			}
		}
			
        return from, nil;  
	}
