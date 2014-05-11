package core

import "bytes"

/** Generic parser class.
* All parsers inherit this class.
 */
type ParserCore struct {
	nesting_level int

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
			this.lexer.Match(LexerCore_ID)
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

func (this *ParserCore) Dbg_enter(rule string) {
	var stringBuffer bytes.Buffer
	for i := 0; i < this.nesting_level; i++ {
		stringBuffer.WriteString(">")
	}
	if Debug.ParserDebug {
		println(stringBuffer.String() + rule + "\nlexer buffer = \n" + this.lexer.GetRest())
	}
	this.nesting_level++
}

func (this *ParserCore) Dbg_leave(rule string) {
	var stringBuffer bytes.Buffer
	for i := 0; i < this.nesting_level; i++ {
		stringBuffer.WriteString("<")
	}
	if Debug.ParserDebug {
		println(stringBuffer.String() + rule + "\nlexer buffer = \n" + this.lexer.GetRest())
	}
	this.nesting_level--
}

func (this *ParserCore) PeekLine(rule string) {
	if Debug.ParserDebug {
		Debug.println(rule + " " + this.lexer.PeekLine())
	}
}
