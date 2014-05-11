package core

import (
	"strconv"
)

type Token struct {
	tokenValue string
	tokenType  int
}

func NewToken(tvalue string, ttype int) *Token {
	return &Token{tokenValue: tvalue, tokenType: ttype}
}

func (this *Token) GetTokenValue() string {
	return this.tokenValue
}
func (this *Token) GetTokenType() int {
	return this.tokenType
}
func (this *Token) SetTokenValue(tvalue string) {
	this.tokenValue = tvalue
}
func (this *Token) SetTokenType(ttype int) {
	this.tokenType = ttype
}

func (this *Token) String() string {
	return "tokenValue = " + this.tokenValue + " / tokenType = " + strconv.Itoa(this.tokenType)
}

func (this *Token) Clone() interface{} {
	retval := &Token{}

	retval.tokenType = this.tokenType
	retval.tokenValue = this.tokenValue

	return retval
}
