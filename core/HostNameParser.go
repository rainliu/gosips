package core

import (
	"bytes"
	"errors"
)

/** SIPParser for host names.
 */
type HostNameParser struct {
	CoreParser
}

func NewHostNameParser(hname string) *HostNameParser {
	this := &HostNameParser{}

	this.lexer = NewCoreLexer("charLexer", hname)

	return this
}

/** The lexer is initialized with the buffer.
 */
func NewHostNameParserFromLexer(lexer Lexer) *HostNameParser {
	this := &HostNameParser{}

	this.CoreParser.SetLexer(lexer)
	this.CoreParser.GetLexer().SelectLexer("charLexer")

	return this
}

func (this *HostNameParser) DomainLabel() (s string, ParseException error) {
	var retval bytes.Buffer
	if Debug.ParserDebug {
		this.Dbg_enter("domainLabel")
		defer this.Dbg_leave("domainLabel")
	}

	for this.lexer.HasMoreChars() {
		la, err := this.lexer.LookAheadK(0)
		if err != nil {
			return retval.String(), err
		}
		if this.lexer.IsAlpha(la) {
			this.lexer.ConsumeK(1)
			retval.WriteByte(la)
		} else if this.lexer.IsDigit(la) {
			this.lexer.ConsumeK(1)
			retval.WriteByte(la)
		} else if la == '-' {
			this.lexer.ConsumeK(1)
			retval.WriteByte(la)
		} else {
			break
		}
	}

	return retval.String(), nil
}

func (this *HostNameParser) Ipv6Reference() (s string, ParseException error) {
	var retval bytes.Buffer
	if Debug.ParserDebug {
		this.Dbg_enter("ipv6Reference")
		defer this.Dbg_leave("ipv6Reference")
	}

	for this.lexer.HasMoreChars() {
		la, err := this.lexer.LookAheadK(0)
		if err != nil {
			return retval.String(), err
		}
		if this.lexer.IsHexDigit(la) {
			this.lexer.ConsumeK(1)
			retval.WriteByte(la)
		} else if la == '.' ||
			la == ':' ||
			la == '[' {
			this.lexer.ConsumeK(1)
			retval.WriteByte(la)
		} else if la == ']' {
			this.lexer.ConsumeK(1)
			retval.WriteByte(la)
			return retval.String(), nil
		} else {
			break
		}
	}

	return retval.String(), errors.New("ParseException: Illegal Host name")
}

func (this *HostNameParser) GetHost() (h *Host, err error) {
	if Debug.ParserDebug {
		this.Dbg_enter("host")
		defer this.Dbg_leave("host")
	}

	var hname bytes.Buffer
	var next byte
	var nextToks string

	//IPv6 referene
	if next, err = this.lexer.LookAheadK(0); err == nil && next == '[' {
		if nextToks, err = this.Ipv6Reference(); err != nil {
			return nil, err
		}
		hname.WriteString(nextToks)
	} else { //IPv4 address or hostname
		if nextToks, err = this.DomainLabel(); err != nil {
			return nil, err
		}
		hname.WriteString(nextToks)
		// Bug reported by Stuart Woodsford (used to barf on
		// more than 4 components to the name).
		for this.lexer.HasMoreChars() {
			// Reached the end of the buffer.
			if next, err = this.lexer.LookAheadK(0); err == nil && next == '.' {
				this.lexer.ConsumeK(1)
				if nextToks, err = this.DomainLabel(); err != nil {
					return nil, err
				}
				hname.WriteString(".")
				hname.WriteString(nextToks)
			} else {
				break
			}
		}
	}

	hostname := hname.String()

	if hostname == "" {
		return nil, errors.New("ParseException: Illegal Host name")
	} else {
		return NewHost(hostname), nil
	}
}

func (this *HostNameParser) GetHostPort() (hp *HostPort, ParseException error) {
	if Debug.ParserDebug {
		this.Dbg_enter("hostPort")
		defer this.Dbg_leave("hostPort")
	}

	host, err := this.GetHost()
	if err != nil {
		return nil, err
	}
	hp = &HostPort{host: host, port: -1}
	// Has a port?
	if this.lexer.HasMoreChars() {
		if next, err := this.lexer.LookAheadK(0); err == nil && next == ':' {
			this.lexer.ConsumeK(1)

			port, err := this.lexer.Number()
			if err != nil {
				return nil, err
			}
			hp.SetPort(port)

		}
	}
	return hp, nil
}
