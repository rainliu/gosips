package parser

import (
	"gosip/header"
)


/** parameters parser header.
*/
type ParametersParser struct{ 
	HeaderParserImpl
}

	/*protected ParametersParser(Lexer lexer) {
		super((Lexer)lexer);
	}*/

	func NewParametersParser(buffer string) *ParametersParser {
		this := &ParametersParser{};
       	this.HeaderParserImpl.super(buffer);
		return this;
	}
	
	func (this *ParametersParser) super(buffer string){
		this.HeaderParserImpl.super(buffer);
	}

	func (this *ParametersParser) Parse(parametersHeader *header.ParametersHeader) (ParseException error) {
		var ch byte;
		var err error;
		
		lexer := this.GetLexer();
		
		lexer.SPorHT();
		if ch, err = lexer.LookAheadK(0); err!=nil{
			return err;
		}
		
		for ch==';' {
		   lexer.ConsumeK(1);
		   // eat white space
           lexer.SPorHT();
		   nv := this.NameValue('=');
		   parametersHeader.SetParameter(nv.GetName(), nv.GetValue().(string));
		   // eat white space
           lexer.SPorHT();
           
           if ch, err = lexer.LookAheadK(0); err!=nil{
           	  return err;
           }
		}
		return nil;
	}
