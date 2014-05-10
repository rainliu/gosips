package core

import (
	"bytes"
	"errors"
)

/** Parser for host names.
 */

type HostNameParser struct {
	ParserCore
}

func NewHostNameParser(hname string) *HostNameParser {
	this := &HostNameParser{}

	this.lexer = NewLexerCore("charLexer", hname)

	return this
}

/** The lexer is initialized with the buffer.
 */

func NewHostNameParserFromLexer(lexer Lexer) *HostNameParser {
	this := &HostNameParser{}

	this.ParserCore.SetLexer(lexer)
	this.ParserCore.GetLexer().SelectLexer("charLexer")

	return this
}

func (this *HostNameParser) DomainLabel() (s string, ParseException error) {
	var retval bytes.Buffer //= new StringBuffer();
	if Debug.ParserDebug {
		this.Dbg_enter("domainLabel")
		defer this.Dbg_leave("domainLabel")
	}
	//try {
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
	//Debug.println("returning " + retval.toString());
	return retval.String(), nil
	//} finally {
	//    if (debug) dbg_leave("domainLabel");
	//}
}

func (this *HostNameParser) Ipv6Reference() (s string, ParseException error) {
	var retval bytes.Buffer //= new StringBuffer();
	if Debug.ParserDebug {
		this.Dbg_enter("ipv6Reference")
		defer this.Dbg_leave("ipv6Reference")
	}

	//try {
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
	//} finally {
	//    if (debug) dbg_leave("ipv6Reference");
	//}
}

func (this *HostNameParser) GetHost() (h *Host, ParseException error) {
	if Debug.ParserDebug {
		this.Dbg_enter("host")
		defer this.Dbg_leave("host")
	}
	//try {
	var hname bytes.Buffer //= new StringBuffer();

	//IPv6 referene
	if next, err := this.lexer.LookAheadK(0); err == nil && next == '[' {
		nextToks, _ := this.Ipv6Reference()
		hname.WriteString(nextToks)
	} else { //IPv4 address or hostname
		nextToks, _ := this.DomainLabel()
		hname.WriteString(nextToks)
		// Bug reported by Stuart Woodsford (used to barf on
		// more than 4 components to the name).
		for this.lexer.HasMoreChars() {
			// Reached the end of the buffer.
			if nextTok, err := this.lexer.LookAheadK(0); err == nil && nextTok == '.' {
				this.lexer.ConsumeK(1)
				nextToks, err = this.DomainLabel()
				hname.WriteString(".")
				hname.WriteString(nextToks)
			} else {
				break
			}
		}
	}

	hostname := hname.String()
	//println(hostname);
	if hostname == "" {
		return nil, errors.New("ParseException: Illegal Host name")
	} //else{
	return NewHost(hostname), nil
	//}
	//} finally {
	//    if (debug) dbg_leave("host");
	//}
}

func (this *HostNameParser) GetHostPort() (hp *HostPort, ParseException error) {
	if Debug.ParserDebug {
		this.Dbg_enter("hostPort")
		defer this.Dbg_leave("hostPort")
	}

	//try {
	host, _ := this.GetHost()
	hp = &HostPort{host: host, port: -1}
	// Has a port?
	if this.lexer.HasMoreChars() {
		if next, err := this.lexer.LookAheadK(0); err == nil && next == ':' {
			this.lexer.ConsumeK(1)
			//try {
			port, _ := this.lexer.Number()
			//p, _ := strconv.Atoi(port)
			hp.SetPort(port)
			/*} catch (NumberFormatException nfe) {
			  throw new ParseException
			          (lexer.getBuffer() + " :Error parsing port ",
			                  lexer.getPtr());*/
		} // else {
		//   return nil, errors.New("ParseException: Error parsing port")
		//}
	}
	return hp, nil
	//} finally {
	//    if (debug) dbg_leave("hostPort");
	//}
}
