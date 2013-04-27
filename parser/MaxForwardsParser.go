package parser

import (
	"gosip/core"
	"gosip/header"
)

/** Parser for Max Forwards Header.*/
type MaxForwardsParser struct{
 	HeaderParserImpl
}

       func NewMaxForwardsParser(contentLength string) *MaxForwardsParser {
       	this := &MaxForwardsParser{}
       	this.HeaderParserImpl.super(contentLength);
       	return this;
       }

        func NewMaxForwardsParserFromLexer(lexer core.Lexer) *MaxForwardsParser {
			this := &MaxForwardsParser{}
       		this.HeaderParserImpl.superFromLexer(lexer);
       		return this;
		}
	
	func (this *MaxForwardsParser) super(hname string){
    	this.HeaderParserImpl.super(hname);
    }
    
	func (this *MaxForwardsParser) Parse() (sh header.SIPHeader, ParseException error) {
	     //if (debug) dbg_enter("MaxForwardsParser.enter");
         //    try {
		contentLength := header.NewMaxForwards();
		this.HeaderName (TokenTypes_MAX_FORWARDS);
		lexer := this.GetLexer();
        number,_:= lexer.Number();
        contentLength.SetMaxForwards(number);
        lexer.SPorHT();
		lexer.Match('\n');
        return contentLength,nil;
             /* } catch (InvalidArgumentException ex) {
		   throw createParseException(ex.getMessage());
              } catch (NumberFormatException ex) {
		   throw createParseException(ex.getMessage());
              }  finally {
			if (debug) dbg_leave("MaxForwardsParser.leave");
	      }*/
	}

	
       
