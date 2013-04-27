package core

import (
	//"fmt"
    "bytes"
    "errors"
    "strconv"
    "strings"
)

/** A lexical analyzer that is used by all parsers in our implementation.
 *
 */
// IMPORTANT - All keyword matches should be between START and END
const LexerCore_START = 2048
const LexerCore_END = LexerCore_START + 2048
const LexerCore_ID = LexerCore_END - 1 // IMPORTANT -- This should be < END

// Individial token classes.
const LexerCore_WHITESPACE = LexerCore_END + 1
const LexerCore_DIGIT = LexerCore_END + 2
const LexerCore_ALPHA = LexerCore_END + 3
const LexerCore_BACKSLASH = (int)('\\')
const LexerCore_QUOTE = (int)('\'')
const LexerCore_AT = (int)('@')
const LexerCore_SP = (int)(' ')
const LexerCore_HT = (int)('\t')
const LexerCore_COLON = (int)(':')
const LexerCore_STAR = (int)('*')
const LexerCore_DOLLAR = (int)('$')
const LexerCore_PLUS = (int)('+')
const LexerCore_POUND = (int)('#')
const LexerCore_MINUS = (int)('-')
const LexerCore_DOUBLEQUOTE = (int)('"')
const LexerCore_TILDE = (int)('~')
const LexerCore_BACK_QUOTE = (int)('`')
const LexerCore_NULL = (int)(0) //('\0')	;
const LexerCore_EQUALS = (int)('=')
const LexerCore_SEMICOLON = (int)(';')
const LexerCore_SLASH = (int)('/')
const LexerCore_L_SQUARE_BRACKET = (int)('[')
const LexerCore_R_SQUARE_BRACKET = (int)(']')
const LexerCore_R_CURLY = (int)('}')
const LexerCore_L_CURLY = (int)('{')
const LexerCore_HAT = (int)('^')
const LexerCore_BAR = (int)('|')
const LexerCore_DOT = (int)('.')
const LexerCore_EXCLAMATION = (int)('!')
const LexerCore_LPAREN = (int)('(')
const LexerCore_RPAREN = (int)(')')
const LexerCore_GREATER_THAN = (int)('>')
const LexerCore_LESS_THAN = (int)('<')
const LexerCore_PERCENT = (int)('%')
const LexerCore_QUESTION = (int)('?')
const LexerCore_AND = (int)('&')
const LexerCore_UNDERSCORE = (int)('_')

type LexerCore struct {
    StringTokenizer

    globalSymbolTable map[int]string
    lexerTables       map[string]LexerMap
    currentLexer      LexerMap
    currentLexerName  string
    currentMatch      *Token
}

/*static{
    globalSymbolTable  = new Hashtable();
    lexerTables = new Hashtable();
}*/
func NewLexerCore(lexerName string, buffer string) *LexerCore {
    this := &LexerCore{}

	this.StringTokenizer.super(buffer);

    this.globalSymbolTable = make(map[int]string)
    this.lexerTables = make(map[string]LexerMap)
    this.currentLexer = make(LexerMap)
    this.currentLexerName = lexerName

    return this
}

func (this *LexerCore) Super(lexerName, buffer string){
	this.StringTokenizer.super(buffer);

    this.globalSymbolTable = make(map[int]string)
    this.lexerTables = make(map[string]LexerMap)
    this.currentLexer = make(LexerMap)
    this.currentLexerName = lexerName
}

func (this *LexerCore) SetLexerName(lexerName string){
	this.currentLexerName = lexerName;
}

func (this *LexerCore) GetLexerName() string{
	return this.currentLexerName;
}
	
func (this *LexerCore) AddKeyword(name string, value int) {
    // System.out.println("addKeyword " + name + " value = " + value);
    // new Exception().printStackTrace();
    //Integer val = new Integer(value);
    this.currentLexer[name] = value
    if _, ok := this.globalSymbolTable[value]; !ok {
        this.globalSymbolTable[value] = name
    }
}

func (this *LexerCore) LookupToken(value int) string {
    if value > LexerCore_START {
        return this.globalSymbolTable[value]
    }   //else {
    return strconv.Itoa(value)
    //}
}

func (this *LexerCore) AddLexer(lexerName string) LexerMap {
    var ok bool
    this.currentLexer, ok = this.lexerTables[lexerName]
    if !ok {
        this.currentLexer = make(LexerMap)
        this.lexerTables[lexerName] = this.currentLexer
    }
    return this.currentLexer
}

//public abstract void selectLexer(String lexerName);

func (this *LexerCore) SelectLexer(lexerName string) {
	this.currentLexer = this.lexerTables[lexerName];
    this.currentLexerName = lexerName
}

func (this *LexerCore) CurrentLexer() LexerMap{
	return this.currentLexer;	
}

/*protected LexerCore() {
    this.currentLexer = new Hashtable();
    this.currentLexerName = "charLexer";
}*/

/*public LexerCore(String lexerName) {
	this();
        this.currentLexerName = lexerName;
    }*/

/** Initialize the lexer with a buffer.
 */
/*public LexerCore(String lexerName, String buffer) {
    this(lexerName);
    this.buffer =  buffer;
}*/

/** Peek the next id but dont move the buffer pointer forward.
 */

func (this *LexerCore) PeekNextId() string {
    oldPtr := this.ptr
    retval := this.Ttoken()
    this.savedPtr = this.ptr
    this.ptr = oldPtr
    return retval
}

/** Get the next id.
 */
func (this *LexerCore) GetNextId() string {
    return this.Ttoken()
}

// call this after you call match
func (this *LexerCore) GetNextToken() *Token {
    return this.currentMatch

}

/** Look ahead for one token.
 */
func (this *LexerCore) PeekNextToken() *Token {
    return this.PeekNextTokenK(1)[0]
}

func (this *LexerCore) PeekNextTokenK(ntokens int) []*Token {
    old := this.ptr
    //fmt.Printf("len=%d\n",len(this.currentLexer));
    retval := make([]*Token, ntokens)
    for i := 0; i < ntokens; i++ {
        tok := &Token{}
        if this.StartsId() {
            id := this.Ttoken()
            tok.tokenValue = id
            if _, ok := this.currentLexer[strings.ToUpper(id)]; ok {
                tok.tokenType = this.currentLexer[strings.ToUpper(id)]
            } else {
                tok.tokenType = LexerCore_ID
            }
        } else {
            nextChar, _ := this.GetNextChar()
            tok.tokenValue += string(nextChar)
            if this.IsAlpha(nextChar) {
                tok.tokenType = LexerCore_ALPHA
            } else if this.IsDigit(nextChar) {
                tok.tokenType = LexerCore_DIGIT
            } else {
                tok.tokenType = (int)(nextChar)
            }
        }
        retval[i] = tok
    }
    this.savedPtr = this.ptr
    this.ptr = old
    return retval
}

/** Match the given token or throw an exception if no such token
 * can be matched.
 */
func (this *LexerCore) Match(tok int) (t *Token, ParseException error) {
    if Debug.ParserDebug {
        Debug.println("match " + strconv.Itoa(tok))
    }
    if tok > LexerCore_START && tok < LexerCore_END {
        if tok == LexerCore_ID {
            // Generic ID sought.
            if !this.StartsId() {
                return nil, errors.New("ParseException: ID expected")
            }
            id := this.GetNextId()
            this.currentMatch = &Token{}
            this.currentMatch.tokenValue = id
            this.currentMatch.tokenType = LexerCore_ID
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
    } else if tok > LexerCore_END {
        // Character classes.
        next, err := this.LookAheadK(0)
        if err != nil {
            return nil, errors.New("ParseException: Expecting DIGIT")
        }
        if tok == LexerCore_DIGIT {
            if !this.IsDigit(next) {
                return nil, errors.New("ParseException: Expecting DIGIT")
            }
            this.currentMatch = &Token{}
            this.currentMatch.tokenValue = string(next)
            this.currentMatch.tokenType = tok
            this.ConsumeK(1)
        } else if tok == LexerCore_ALPHA {
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

func (this *LexerCore) SPorHT() {
    //try {
    var ch byte;
    
    ch, _ = this.LookAheadK(0);
    for ch == ' ' || ch == '\t' {
    	this.ConsumeK(1)
    	ch, _ = this.LookAheadK(0);
    }
	//} catch (ParseException ex) {
    // Ignore
    //}
}
func (this *LexerCore) StartsId() bool {
    //try {
    nextChar, err := this.LookAheadK(0)
    if err != nil {
        return false
    }
    return (this.IsAlpha(nextChar) ||
        this.IsDigit(nextChar) ||
        nextChar == '_' ||
        nextChar == '+' ||
        nextChar == '-' ||
        nextChar == '!' ||
        nextChar == '`' ||
        nextChar == '\'' ||
        nextChar == '~' ||
        nextChar == '.' ||
        nextChar == '*')
    //} catch (ParseException ex) {
    //return false
    //}
}

func (this *LexerCore) Ttoken() string {
    var nextId bytes.Buffer //=  new StringBuffer();
    //try {
    for this.HasMoreChars() {
        nextChar, err := this.LookAheadK(0)
        if err != nil {
            break
        }
        //Debug.println("nextChar = " + nextChar);
        if this.IsAlpha(nextChar) ||
            this.IsDigit(nextChar) ||
            nextChar == '_' ||
            nextChar == '+' ||
            nextChar == '-' ||
            nextChar == '!' ||
            nextChar == '`' ||
            nextChar == '\'' ||
            nextChar == '~' ||
            nextChar == '.' ||
            nextChar == '*' {
            this.ConsumeK(1)
            nextId.WriteByte(nextChar)
        } else {
            break
        }
    }
    return nextId.String()
    //} catch (ParseException ex) {
    //    return nextId.toString();
    //}
}

func (this *LexerCore) TtokenAllowSpace() string {
    var nextId bytes.Buffer //=  new StringBuffer();
    //try {
    for this.HasMoreChars() {
        nextChar, err := this.LookAheadK(0)
        if err != nil {
            break
        }
        //Debug.println("nextChar = " + nextChar);
        if this.IsAlpha(nextChar) ||
            this.IsDigit(nextChar) ||
            nextChar == '_' ||
            nextChar == '+' ||
            nextChar == '-' ||
            nextChar == '!' ||
            nextChar == '`' ||
            nextChar == '\'' ||
            nextChar == '~' ||
            nextChar == '.' ||
            nextChar == ' ' ||
            nextChar == '\t' ||
            nextChar == '*' {
            nextId.WriteByte(nextChar)
            this.ConsumeK(1)
        } else {
            break
        }
    }
    return nextId.String()
    //}  catch (ParseException ex) {
    //    return nextId.toString();
    //}
}

// Assume the cursor is at a quote.
func (this *LexerCore) QuotedString() (s string, ParseException error) {
    var retval bytes.Buffer //= new StringBuffer();
    if next, err := this.LookAheadK(0); next != '"' || err != nil {
        return "", nil
    }
    this.ConsumeK(1)
    for {
        next, err := this.GetNextChar()
        if err != nil {
            break
        }
        if next == '"' {
            // Got to the terminating quote.
            break
        } else if next == 0 { //'\0' {
            return "", errors.New("ParseException: unexpected EOL")
        } else if next == '\\' {
            retval.WriteByte(next)
            next, _ = this.GetNextChar()
            retval.WriteByte(next)
        } else {
            retval.WriteByte(next)
        }
    }
    return retval.String(), nil
}

// Assume the cursor is at a "("
func (this *LexerCore) Comment() (s string, ParseException error) {
    var retval bytes.Buffer //= new StringBuffer();
    if next, err := this.LookAheadK(0); next != '(' || err != nil {
        return "", nil
    }
    this.ConsumeK(1)
    for {
        next, err := this.GetNextChar()
        if err != nil {
            break
        }
        if next == ')' {
            break
        } else if next == 0 { //'\0' {
            return "", errors.New("ParseException: unexpected EOL")
        } else if next == '\\' {
            retval.WriteByte(next)
            next, _ = this.GetNextChar()
            if next == 0 { //'\0'{
                return "", errors.New("ParseException: unexpected EOL")
            }
            retval.WriteByte(next)
        } else {
            retval.WriteByte(next)
        }
    }
    return retval.String(), nil
}

func (this *LexerCore) ByteStringNoSemicolon() string {
    var retval bytes.Buffer //= new StringBuffer();
    //try  {
    for {
        next, err := this.LookAheadK(0)
        if err != nil {
            break
        }
        if next == 0 /*'\0'*/ || next == '\n' || next == ';' {
            break
        } else {
            this.ConsumeK(1)
            retval.WriteByte(next)
        }
    }
    //} catch (ParseException ex) {
    //    return retval.toString();
    //}
    return retval.String()
}

func (this *LexerCore) ByteStringNoComma() string {
    var retval bytes.Buffer //= new StringBuffer();
    //try {
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
    //} catch (ParseException ex) {
    //}
    return retval.String()
}

func (this *LexerCore) CharAsString(ch byte) string {
	var retval bytes.Buffer
	retval.WriteByte(ch);
    return retval.String();
}

/** Lookahead in the inputBuffer for n chars and return as a string.
 * Do not consume the input.
 */
func (this *LexerCore) NCharAsString(nchars int) string {
    var retval bytes.Buffer // new StringBuffer();
    //try {
    for i := 0; i < nchars; i++ {
        next, err := this.LookAheadK(i)
        if err != nil {
            break
        }
        retval.WriteByte(next)
    }
    return retval.String()
    //} catch (ParseException ex) {
    //    return retval.toString();
    //}
}

/** Get and consume the next number.
 */
func (this *LexerCore) Number() (n int, ParseException error) {
    var retval bytes.Buffer //= new StringBuffer();
    //try {
    next, err := this.LookAheadK(0)
    if err != nil {
        return -1, err;
    }
    if !this.IsDigit(next) {
    	return -1, errors.New("ParseException: unexpected token \""+string(next)+"\"");
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
    
    if n, err = strconv.Atoi(retval.String()); err!=nil{
    	return -1, err;
    }else{
    	return n, nil
    }
    //} catch (ParseException ex) {
    //    return retval.toString();
    //}
}

/** Mark the position for backtracking.
 */
func (this *LexerCore) MarkInputPosition() int {
    return this.ptr
}

/** Rewind the input ptr to the marked position.
 */
func (this *LexerCore) RewindInputPosition(position int) {
    this.ptr = position
}

/** Get the rest of the String
 * @return String
 */
func (this *LexerCore) GetRest() string {
    if this.ptr >= len(this.buffer) {
        return ""
    }   //else {
    return this.buffer[this.ptr:]
    //}
}

/** Get the sub-String until the character is encountered.
 * Acknowledgement - Sylvian Corre submitted a bug fix for this
 * method.
 * @param char c the character to match
 * @return matching string.
 */
func (this *LexerCore) GetString(c byte) (s string, ParseException error) {
    var savedPtr int = this.ptr
    var retval bytes.Buffer // = new StringBuffer();
    for {
        next, _ := this.LookAheadK(0)
        //System.out.println(" next = [" + next + ']' + "ptr = " + ptr);
        //System.out.println(next == '\0');

        if next == 0 { //'\0'   {
            this.ptr = savedPtr
            return "", errors.New("ParseException: unexpected EOL")
        } else if next == c {
            this.ConsumeK(1)
            break
        } else if next == '\\' {
            this.ConsumeK(1)
            nextchar, _ := this.LookAheadK(0)
            if nextchar == 0 { //'\0'  {
                this.ptr = savedPtr
                return "", errors.New("ParseException: unexpected EOL")
            } else {
                this.ConsumeK(1)
                retval.WriteByte(nextchar)
            }
        } else {
            this.ConsumeK(1)
            retval.WriteByte(next)
        }
    }
    return retval.String(), nil
}

/** Get the read pointer.
 */
func (this *LexerCore) GetPtr() int {
    return this.ptr
}

/** Get the buffer.
 */
func (this *LexerCore) GetBuffer() string {
    return this.buffer
}

/** Create a parse exception.
 */
/*public ParseException createParseException() {
	  return new ParseException(this.buffer,this.ptr);
     }*/
