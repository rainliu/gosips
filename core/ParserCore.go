package core

import "bytes"

/** Generic parser class.
* All parsers inherit this class.
 */
type ParserCore struct {
	//public static final boolean debug = Debug.parserDebug;
	nesting_level int //protected static int

	lexer Lexer //*LexerCore
}

func NewParserCore(buffer string) *ParserCore {
	this := &ParserCore{}

	this.lexer = NewLexerCore("CharLexer", buffer)

	return this
}

func (this *ParserCore) Super(buffer string) {
	this.lexer = NewLexerCore("CharLexer", buffer)
}

func (this *ParserCore) GetLexer() Lexer {
	return this.lexer
}
func (this *ParserCore) SetLexer(lexer Lexer) {
	this.lexer = lexer
}

func (this *ParserCore) NameValue(separator byte) *NameValue {
	if Debug.ParserDebug {
		this.Dbg_enter("nameValue")
		defer this.Dbg_leave("nameValue")
	}

	this.lexer.Match(LexerCore_ID)
	name := this.lexer.GetNextToken()
	// eat white space.
	//println(name.String())
	this.lexer.SPorHT()
	//try {
	quoted := false
	la, err := this.lexer.LookAheadK(0)
	if la == separator && err == nil {
		//println(this.lexer.GetRest())
		this.lexer.ConsumeK(1)
		//println(this.lexer.GetRest())
		this.lexer.SPorHT()
		//println(this.lexer.GetRest())
		var str string

		if la, err = this.lexer.LookAheadK(0); la == '"' && err == nil {
			str, _ = this.lexer.QuotedString()
			quoted = true
		} else {
			//fmt.Printf("%c\n", la)
			//println(this.lexer.GetRest())
			this.lexer.Match(LexerCore_ID)
			//println(this.lexer.GetRest())
			value := this.lexer.GetNextToken()
			str = value.tokenValue
			//println(value.String())
		}
		nv := NewNameValue(name.tokenValue, str)
		if quoted {
			nv.SetQuotedValue()
		}
		//println(nv.GetValue().(string))
		return nv
	} //else {
	return NewNameValue(name.tokenValue, "")
	//}
	//} catch (ParseException ex) {
	//	return new NameValue(name.tokenValue,null);
	//}
}

func (this *ParserCore) Dbg_enter(rule string) {
	var stringBuffer bytes.Buffer //= new StringBuffer();
	for i := 0; i < this.nesting_level; i++ {
		stringBuffer.WriteString(">")
	}
	if Debug.ParserDebug {
		println(stringBuffer.String() + rule + "\nlexer buffer = \n" + this.lexer.GetRest())
	}
	this.nesting_level++
}

func (this *ParserCore) Dbg_leave(rule string) {
	var stringBuffer bytes.Buffer //= new StringBuffer();
	for i := 0; i < this.nesting_level; i++ {
		stringBuffer.WriteString("<")
	}
	if Debug.ParserDebug {
		println(stringBuffer.String() + rule + "\nlexer buffer = \n" + this.lexer.GetRest())
	}
	this.nesting_level--
}

/*func (this *ParserCore)  NameValue nameValue() * {
	return nameValue('=');
}*/

func (this *ParserCore) PeekLine(rule string) {
	if Debug.ParserDebug {
		Debug.println(rule + " " + this.lexer.PeekLine())
	}
}
