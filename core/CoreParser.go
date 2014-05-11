package core

import "bytes"

/** Generic parser class.
* All parsers inherit this class.
 */
type CoreParser struct {
	nesting_level int

	lexer Lexer //*CoreLexer
}

func NewCoreParser(buffer string) *CoreParser {
	this := &CoreParser{}

	this.lexer = NewCoreLexer("CharLexer", buffer)

	return this
}

func (this *CoreParser) Super(buffer string) {
	this.lexer = NewCoreLexer("CharLexer", buffer)
}

func (this *CoreParser) GetLexer() Lexer {
	return this.lexer
}
func (this *CoreParser) SetLexer(lexer Lexer) {
	this.lexer = lexer
}

func (this *CoreParser) NameValue(separator byte) *NameValue {
	if Debug.ParserDebug {
		this.Dbg_enter("nameValue")
		defer this.Dbg_leave("nameValue")
	}

	this.lexer.Match(CORELEXER_ID)
	name := this.lexer.GetNextToken()
	// eat white space.
	this.lexer.SPorHT()

	quoted := false
	la, err := this.lexer.LookAheadK(0)
	if la == separator && err == nil {
		this.lexer.ConsumeK(1)
		this.lexer.SPorHT()

		var str string

		if la, err = this.lexer.LookAheadK(0); la == '"' && err == nil {
			str, _ = this.lexer.QuotedString()
			quoted = true
		} else {
			this.lexer.Match(CORELEXER_ID)
			value := this.lexer.GetNextToken()
			str = value.tokenValue
		}
		nv := NewNameValue(name.tokenValue, str)
		if quoted {
			nv.SetQuotedValue()
		}

		return nv
	} else {
		return NewNameValue(name.tokenValue, "")
	}
}

func (this *CoreParser) Dbg_enter(rule string) {
	var stringBuffer bytes.Buffer
	for i := 0; i < this.nesting_level; i++ {
		stringBuffer.WriteString(">")
	}
	if Debug.ParserDebug {
		println(stringBuffer.String() + rule + "\nlexer buffer = \n" + this.lexer.GetRest())
	}
	this.nesting_level++
}

func (this *CoreParser) Dbg_leave(rule string) {
	var stringBuffer bytes.Buffer
	for i := 0; i < this.nesting_level; i++ {
		stringBuffer.WriteString("<")
	}
	if Debug.ParserDebug {
		println(stringBuffer.String() + rule + "\nlexer buffer = \n" + this.lexer.GetRest())
	}
	this.nesting_level--
}

func (this *CoreParser) PeekLine(rule string) {
	if Debug.ParserDebug {
		Debug.println(rule + " " + this.lexer.PeekLine())
	}
}
