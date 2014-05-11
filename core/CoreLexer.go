package core

import (
	"bytes"
	"errors"
	"strconv"
	"strings"
)

/** A lexical analyzer that is used by all parsers in our implementation.
 */
// IMPORTANT - All keyword matches should be between START and END
const CORELEXER_START = 2048
const CORELEXER_END = CORELEXER_START + 2048
const CORELEXER_ID = CORELEXER_END - 1 // IMPORTANT -- This should be < END

// Individial token classes.
const CORELEXER_WHITESPACE = CORELEXER_END + 1
const CORELEXER_DIGIT = CORELEXER_END + 2
const CORELEXER_ALPHA = CORELEXER_END + 3
const CORELEXER_BACKSLASH = (int)('\\')
const CORELEXER_QUOTE = (int)('\'')
const CORELEXER_AT = (int)('@')
const CORELEXER_SP = (int)(' ')
const CORELEXER_HT = (int)('\t')
const CORELEXER_COLON = (int)(':')
const CORELEXER_STAR = (int)('*')
const CORELEXER_DOLLAR = (int)('$')
const CORELEXER_PLUS = (int)('+')
const CORELEXER_POUND = (int)('#')
const CORELEXER_MINUS = (int)('-')
const CORELEXER_DOUBLEQUOTE = (int)('"')
const CORELEXER_TILDE = (int)('~')
const CORELEXER_BACK_QUOTE = (int)('`')
const CORELEXER_NULL = (int)(0) //('\0')	;
const CORELEXER_EQUALS = (int)('=')
const CORELEXER_SEMICOLON = (int)(';')
const CORELEXER_SLASH = (int)('/')
const CORELEXER_L_SQUARE_BRACKET = (int)('[')
const CORELEXER_R_SQUARE_BRACKET = (int)(']')
const CORELEXER_R_CURLY = (int)('}')
const CORELEXER_L_CURLY = (int)('{')
const CORELEXER_HAT = (int)('^')
const CORELEXER_BAR = (int)('|')
const CORELEXER_DOT = (int)('.')
const CORELEXER_EXCLAMATION = (int)('!')
const CORELEXER_LPAREN = (int)('(')
const CORELEXER_RPAREN = (int)(')')
const CORELEXER_GREATER_THAN = (int)('>')
const CORELEXER_LESS_THAN = (int)('<')
const CORELEXER_PERCENT = (int)('%')
const CORELEXER_QUESTION = (int)('?')
const CORELEXER_AND = (int)('&')
const CORELEXER_UNDERSCORE = (int)('_')

type CoreLexer struct {
	StringTokenizer

	globalSymbolTable map[int]string
	lexerTables       map[string]LexerMap
	currentLexer      LexerMap
	currentLexerName  string
	currentMatch      *Token
}

func NewCoreLexer(lexerName string, buffer string) *CoreLexer {
	this := &CoreLexer{}

	this.StringTokenizer.super(buffer)

	this.globalSymbolTable = make(map[int]string)
	this.lexerTables = make(map[string]LexerMap)
	this.currentLexer = make(LexerMap)
	this.currentLexerName = lexerName

	return this
}

func (this *CoreLexer) Super(lexerName, buffer string) {
	this.StringTokenizer.super(buffer)

	this.globalSymbolTable = make(map[int]string)
	this.lexerTables = make(map[string]LexerMap)
	this.currentLexer = make(LexerMap)
	this.currentLexerName = lexerName
}

func (this *CoreLexer) SetLexerName(lexerName string) {
	this.currentLexerName = lexerName
}

func (this *CoreLexer) GetLexerName() string {
	return this.currentLexerName
}

func (this *CoreLexer) AddKeyword(name string, value int) {
	this.currentLexer[name] = value
	if _, ok := this.globalSymbolTable[value]; !ok {
		this.globalSymbolTable[value] = name
	}
}

func (this *CoreLexer) LookupToken(value int) string {
	if value > CORELEXER_START {
		return this.globalSymbolTable[value]
	} else {
		return strconv.Itoa(value)
	}
}

func (this *CoreLexer) AddLexer(lexerName string) LexerMap {
	var ok bool
	this.currentLexer, ok = this.lexerTables[lexerName]
	if !ok {
		this.currentLexer = make(LexerMap)
		this.lexerTables[lexerName] = this.currentLexer
	}
	return this.currentLexer
}

func (this *CoreLexer) SelectLexer(lexerName string) {
	this.currentLexer = this.lexerTables[lexerName]
	this.currentLexerName = lexerName
}

func (this *CoreLexer) CurrentLexer() LexerMap {
	return this.currentLexer
}

/** Peek the next id but dont move the buffer pointer forward.
 */
func (this *CoreLexer) PeekNextId() string {
	oldPtr := this.ptr
	retval := this.Ttoken()
	this.savedPtr = this.ptr
	this.ptr = oldPtr
	return retval
}

/** Get the next id.
 */
func (this *CoreLexer) GetNextId() string {
	return this.Ttoken()
}

// call this after you call match
func (this *CoreLexer) GetNextToken() *Token {
	return this.currentMatch

}

/** Look ahead for one token.
 */
func (this *CoreLexer) PeekNextToken() (*Token, error) {
	tok, err := this.PeekNextTokenK(1)
	if err != nil {
		return nil, err
	} else {
		return tok[0], nil
	}
}

func (this *CoreLexer) PeekNextTokenK(ntokens int) ([]*Token, error) {
	old := this.ptr
	retval := make([]*Token, ntokens)
	var err error
	for i := 0; i < ntokens; i++ {
		tok := &Token{}
		if this.StartsId() {
			id := this.Ttoken()
			tok.tokenValue = id
			if _, ok := this.currentLexer[strings.ToUpper(id)]; ok {
				tok.tokenType = this.currentLexer[strings.ToUpper(id)]
			} else {
				tok.tokenType = CORELEXER_ID
			}
		} else {
			nextChar, err := this.GetNextChar()
			if err != nil {
				break
			}
			tok.tokenValue += string(nextChar)
			if this.IsAlpha(nextChar) {
				tok.tokenType = CORELEXER_ALPHA
			} else if this.IsDigit(nextChar) {
				tok.tokenType = CORELEXER_DIGIT
			} else {
				tok.tokenType = (int)(nextChar)
			}
		}
		retval[i] = tok
	}
	this.savedPtr = this.ptr
	this.ptr = old
	return retval, err
}

/** Match the given token or throw an exception if no such token
 * can be matched.
 */
func (this *CoreLexer) Match(tok int) (t *Token, ParseException error) {
	if Debug.ParserDebug {
		Debug.println("match " + strconv.Itoa(tok))
	}
	if tok > CORELEXER_START && tok < CORELEXER_END {
		if tok == CORELEXER_ID {
			// Generic ID sought.
			if !this.StartsId() {
				return nil, errors.New("ParseException: ID expected")
			}
			id := this.GetNextId()
			this.currentMatch = &Token{}
			this.currentMatch.tokenValue = id
			this.currentMatch.tokenType = CORELEXER_ID
		} else {
			nexttok := this.GetNextId()
			cur, ok := this.currentLexer[strings.ToUpper(nexttok)]
			if !ok || cur != tok {
				return nil, errors.New("ParseException: Unexpected Token")
			}
			this.currentMatch = &Token{}
			this.currentMatch.tokenValue = nexttok
			this.currentMatch.tokenType = tok
		}
	} else if tok > CORELEXER_END {
		// Character classes.
		next, err := this.LookAheadK(0)
		if err != nil {
			return nil, errors.New("ParseException: Expecting DIGIT")
		}
		if tok == CORELEXER_DIGIT {
			if !this.IsDigit(next) {
				return nil, errors.New("ParseException: Expecting DIGIT")
			}
			this.currentMatch = &Token{}
			this.currentMatch.tokenValue = string(next)
			this.currentMatch.tokenType = tok
			this.ConsumeK(1)
		} else if tok == CORELEXER_ALPHA {
			if !this.IsAlpha(next) {
				return nil, errors.New("ParseException: Expecting ALPHA")
			}
			this.currentMatch = &Token{}
			this.currentMatch.tokenValue = string(next)
			this.currentMatch.tokenType = tok
			this.ConsumeK(1)
		}
	} else {
		// This is a direct character spec.
		ch := byte(tok)
		next, err := this.LookAheadK(0)
		if err != nil {
			return nil, errors.New("ParseException: Expecting DIGIT")
		}
		if next == ch {
			this.currentMatch = &Token{}
			this.currentMatch.tokenValue = string(ch)
			this.currentMatch.tokenType = tok
			this.ConsumeK(1)
		} else {
			return nil, errors.New("ParseException: Expecting")
		}
	}
	return this.currentMatch, nil
}

func (this *CoreLexer) SPorHT() {
	var ch byte

	for ch, _ = this.LookAheadK(0); ch == ' ' || ch == '\t'; ch, _ = this.LookAheadK(0) {
		this.ConsumeK(1)
	}
}

func (this *CoreLexer) StartsId() bool {
	nextChar, err := this.LookAheadK(0)
	if err != nil {
		return false
	}
	return (this.IsAlpha(nextChar) ||
		this.IsDigit(nextChar) ||
		nextChar == '-' ||
		nextChar == '.' ||
		nextChar == '!' ||
		nextChar == '%' ||
		nextChar == '*' ||
		nextChar == '_' ||
		nextChar == '+' ||
		nextChar == '`' ||
		nextChar == '\'' ||
		nextChar == '~')
}

func (this *CoreLexer) Ttoken() string {
	var nextId bytes.Buffer

	for this.HasMoreChars() {
		nextChar, err := this.LookAheadK(0)
		if err != nil {
			break
		}

		if this.IsAlpha(nextChar) ||
			this.IsDigit(nextChar) ||
			nextChar == '-' ||
			nextChar == '.' ||
			nextChar == '!' ||
			nextChar == '%' ||
			nextChar == '*' ||
			nextChar == '_' ||
			nextChar == '+' ||
			nextChar == '`' ||
			nextChar == '\'' ||
			nextChar == '~' {
			this.ConsumeK(1)
			nextId.WriteByte(nextChar)
		} else {
			break
		}
	}
	return nextId.String()
}

func (this *CoreLexer) TtokenAllowSpace() string {
	var nextId bytes.Buffer

	for this.HasMoreChars() {
		nextChar, err := this.LookAheadK(0)
		if err != nil {
			break
		}

		if this.IsAlpha(nextChar) ||
			this.IsDigit(nextChar) ||
			nextChar == '-' ||
			nextChar == '.' ||
			nextChar == '!' ||
			nextChar == '%' ||
			nextChar == '*' ||
			nextChar == '_' ||
			nextChar == '+' ||
			nextChar == '`' ||
			nextChar == '\'' ||
			nextChar == '~' ||
			nextChar == ' ' ||
			nextChar == '\t' {

			nextId.WriteByte(nextChar)
			this.ConsumeK(1)
		} else {
			break
		}
	}
	return nextId.String()
}

// Assume the cursor is at a quote.
func (this *CoreLexer) QuotedString() (s string, err error) {
	var retval bytes.Buffer
	var next byte

	if next, err = this.LookAheadK(0); next != '"' || err != nil {
		return "", err
	}
	this.ConsumeK(1)
	for {
		if next, err = this.GetNextChar(); err != nil {
			break
		}
		if next == '"' {
			// Got to the terminating quote.
			break
			// } else if next == 0 { //'\0' {
			// 	return "", errors.New("ParseException: unexpected EOL")
		} else if next == '\\' {
			retval.WriteByte(next)
			next, _ = this.GetNextChar()
			retval.WriteByte(next)
		} else {
			retval.WriteByte(next)
		}
	}
	return retval.String(), err
}

// Assume the cursor is at a "("
func (this *CoreLexer) Comment() (s string, err error) {
	var retval bytes.Buffer
	var next byte

	if next, err = this.LookAheadK(0); next != '(' || err != nil {
		return "", err
	}
	this.ConsumeK(1)
	for {
		if next, err = this.GetNextChar(); err != nil {
			break
		}
		if next == ')' {
			break
			// } else if next == 0 { //'\0' {
			// 	return "", errors.New("ParseException: unexpected EOL")
		} else if next == '\\' {
			retval.WriteByte(next)
			if next, err = this.GetNextChar(); err != nil {
				break
			}
			// if next == 0 { //'\0'{
			// 	return "", errors.New("ParseException: unexpected EOL")
			// }
			retval.WriteByte(next)
		} else {
			retval.WriteByte(next)
		}
	}
	return retval.String(), err
}

func (this *CoreLexer) ByteStringNoSemicolon() string {
	var retval bytes.Buffer

	for {
		next, err := this.LookAheadK(0)
		if err != nil {
			break
		}
		if /*next == 0*/ /*'\0'*/ /*||*/ next == '\n' || next == ';' {
			break
		} else {
			this.ConsumeK(1)
			retval.WriteByte(next)
		}
	}

	return retval.String()
}

func (this *CoreLexer) ByteStringNoComma() string {
	var retval bytes.Buffer

	for {
		next, err := this.LookAheadK(0)
		if err != nil {
			break
		}
		if next == '\n' || next == ',' {
			break
		} else {
			this.ConsumeK(1)
			retval.WriteByte(next)
		}
	}

	return retval.String()
}

func (this *CoreLexer) CharAsString(ch byte) string {
	var retval bytes.Buffer
	retval.WriteByte(ch)
	return retval.String()
}

/** Lookahead in the inputBuffer for n chars and return as a string.
 * Do not consume the input.
 */
func (this *CoreLexer) NCharAsString(nchars int) string {
	var retval bytes.Buffer

	for i := 0; i < nchars; i++ {
		next, err := this.LookAheadK(i)
		if err != nil {
			break
		}
		retval.WriteByte(next)
	}

	return retval.String()
}

/** Get and consume the next number.
 */
func (this *CoreLexer) Number() (n int, ParseException error) {
	var retval bytes.Buffer

	next, err := this.LookAheadK(0)
	if err != nil {
		return -1, err
	}
	if !this.IsDigit(next) {
		return -1, errors.New("ParseException: unexpected token \"" + string(next) + "\"")
	}

	retval.WriteByte(next)
	this.ConsumeK(1)
	for {
		next, err := this.LookAheadK(0)
		if err == nil && this.IsDigit(next) {
			retval.WriteByte(next)
			this.ConsumeK(1)
		} else {
			break
		}
	}

	if n, err = strconv.Atoi(retval.String()); err != nil {
		return -1, err
	} else {
		return n, nil
	}
}

/** Mark the position for backtracking.
 */
func (this *CoreLexer) MarkInputPosition() int {
	return this.ptr
}

/** Rewind the input ptr to the marked position.
 */
func (this *CoreLexer) RewindInputPosition(position int) {
	this.ptr = position
}

/** Get the rest of the String
 * @return String
 */
func (this *CoreLexer) GetRest() string {
	if this.ptr >= len(this.buffer) {
		return ""
	} else {
		return this.buffer[this.ptr:]
	}
}

/** Get the sub-String until the character is encountered.
 * Acknowledgement - Sylvian Corre submitted a bug fix for this
 * method.
 * @param char c the character to match
 * @return matching string.
 */
func (this *CoreLexer) GetString(c byte) (s string, err error) {
	var savedPtr int = this.ptr
	var retval bytes.Buffer
	var next byte

	for {
		next, err = this.LookAheadK(0)

		if err != nil /*next == 0*/ { //'\0'
			this.ptr = savedPtr
			break //return "", errors.New("ParseException: unexpected EOL")
		} else if next == c {
			this.ConsumeK(1)
			break
		} else if next == '\\' {
			this.ConsumeK(1)
			next, err = this.LookAheadK(0)
			if err != nil /*nextchar == 0*/ { //'\0'  {
				this.ptr = savedPtr
				break //return "", errors.New("ParseException: unexpected EOL")
			} else {
				this.ConsumeK(1)
				retval.WriteByte(next)
			}
		} else {
			this.ConsumeK(1)
			retval.WriteByte(next)
		}
	}
	return retval.String(), err
}

/** Get the read pointer.
 */
func (this *CoreLexer) GetPtr() int {
	return this.ptr
}

/** Get the buffer.
 */
func (this *CoreLexer) GetBuffer() string {
	return this.buffer
}
