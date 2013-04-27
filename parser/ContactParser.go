package parser

import (
	"gosip/core"
	"gosip/header"
)


/** 
* A parser for The SIP contact header.
*/

type ContactParser struct{
 	AddressParametersParser
}
    
	func NewContactParser(contact string) *ContactParser {
		this := &ContactParser{}
		this.AddressParametersParser.super(contact);
		return this;
	}

    func NewContactParserFromLexer(lexer core.Lexer) *ContactParser {
    	this := &ContactParser{}
		this.AddressParametersParser.superFromLexer(lexer);
		return this;
	}

	func (this *ContactParser) Parse() (sh header.SIPHeader, ParseException error) {
		retval := header.NewContactListImpl();
		// past the header name and the colon.
		lexer := this.GetLexer();
		var la byte;
		this.HeaderName(TokenTypes_CONTACT);
		for {
		   contact := header.NewContactImpl();
		   if la,_=lexer.LookAheadK(0); la == '*'  {
			 lexer.Match('*');
			 contact.SetWildCardFlag(true);
		   } else {
		   	 this.AddressParametersParser.Parse(contact);
		   }
		   retval.AddContact(contact);
		   lexer.SPorHT();
		   if la,_ = lexer.LookAheadK(0); la == ',' {
			lexer.Match(',');
			lexer.SPorHT();
		   } else if la,_=lexer.LookAheadK(0); la == '\n' {
		    break;
		   }else{
		    //println(lexer.GetRest());
		    return nil, this.CreateParseException("unexpected char");
		   }
		}
		return retval, nil;
	}





