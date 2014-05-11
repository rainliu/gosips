package core

import (
	"bytes"
	"errors"
	"strings"
)

/** Base string token splitter.
 */
type StringTokenizer struct {
	buffer   string
	ptr      int
	savedPtr int
}

func NewStringTokenizer(buffer string) *StringTokenizer {
	this := &StringTokenizer{}
	this.buffer = buffer
	this.ptr = 0

	return this
}

func (this *StringTokenizer) super(buffer string) {
	this.buffer = buffer
	this.ptr = 0
}

func (this *StringTokenizer) NextToken() string {
	var retval bytes.Buffer

	for this.ptr < len(this.buffer) {
		if this.buffer[this.ptr] == '\n' {
			retval.WriteByte(this.buffer[this.ptr])
			this.ptr++
			break
		} else {
			retval.WriteByte(this.buffer[this.ptr])
			this.ptr++
		}
	}

	return retval.String()
}

func (this *StringTokenizer) HasMoreChars() bool {
	return this.ptr < len(this.buffer)
}

func (this *StringTokenizer) IsHexDigit(ch byte) bool {
	if this.IsDigit(ch) {
		return true
	}
	ch1 := strings.ToUpper(string(ch))[0]
	return ch1 == 'A' || ch1 == 'B' || ch1 == 'C' ||
		ch1 == 'D' || ch1 == 'E' || ch1 == 'F'

}

func (this *StringTokenizer) IsAlpha(ch byte) bool {
	return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z')
}

func (this *StringTokenizer) IsDigit(ch byte) bool {
	return ch == '0' || ch == '1' || ch == '2' || ch == '3' || ch == '4' ||
		ch == '5' || ch == '6' || ch == '7' || ch == '8' || ch == '9'
}

func (this *StringTokenizer) GetLine() string {
	var retval bytes.Buffer
	for this.ptr < len(this.buffer) && this.buffer[this.ptr] != '\n' {
		retval.WriteByte(this.buffer[this.ptr])
		this.ptr++
	}
	if this.ptr < len(this.buffer) && this.buffer[this.ptr] == '\n' {
		retval.WriteString("\n")
		this.ptr++
	}
	return retval.String()
}

func (this *StringTokenizer) PeekLine() string {
	curPos := this.ptr
	retval := this.GetLine()
	this.ptr = curPos
	return retval
}

func (this *StringTokenizer) LookAhead() (byte, error) {
	return this.LookAheadK(0)
}

func (this *StringTokenizer) LookAheadK(k int) (byte, error) {
	if this.ptr+k < len(this.buffer) {
		return this.buffer[this.ptr+k], nil
	}
	return 0, errors.New("StringTokenizer::LookAheadK: End of buffer")
}

func (this *StringTokenizer) GetNextChar() (byte, error) {
	if this.ptr >= len(this.buffer) {
		return 0, errors.New("StringTokenizer::GetNextChar: End of buffer")
	}
	ch := this.buffer[this.ptr]
	this.ptr++
	return ch, nil
}

func (this *StringTokenizer) Consume() {
	this.ptr = this.savedPtr
}

func (this *StringTokenizer) ConsumeK(k int) {
	this.ptr += k
}

/** Get a Vector of the buffer tokenized by lines
 */
func (this *StringTokenizer) GetLines() map[int]string {
	result := make(map[int]string)
	for this.HasMoreChars() {
		line := this.GetLine()
		result[len(result)] = line
	}
	return result
}

/** Get the next token from the buffer.
 */
func (this *StringTokenizer) GetNextTokenByDelim(delim byte) (string, error) {
	var retval bytes.Buffer
	for {
		la, err := this.LookAheadK(0)
		if err != nil {
			return "", err
		}
		if la == delim {
			break
		}
		retval.WriteByte(this.buffer[this.ptr])
		this.ConsumeK(1)
	}
	return retval.String(), nil
}

/** get the SDP field name of the line
 *  @return String
 */
func (this *StringTokenizer) GetSDPFieldName(line string) string {
	if line == "" {
		return ""
	}

	begin := strings.Index(line, "=")
	if begin != -1 {
		return line[0:begin]
	} else {
		return ""
	}
}
