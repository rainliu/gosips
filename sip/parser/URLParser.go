package parser

import (
	"bytes"
	//"fmt"
	"gosips/core"
	"gosips/sip/address"
	"strings"
)

/** SIPParser For SIP and Tel URLs. Other kinds of URL's are handled by the
 */
type URLParser struct {
	SIPParser
}

func NewURLParser(url string) *URLParser {
	this := &URLParser{}
	this.SIPParser.SetLexer(NewSIPLexer("sip_urlLexer", url))
	return this
}

func NewURLParserFromLexer(lexer core.Lexer) *URLParser {
	this := &URLParser{}
	this.SIPParser.SetLexer(lexer)
	this.SIPParser.GetLexer().SelectLexer("sip_urlLexer")
	return this
}

func (this *URLParser) IsMark(next byte) bool {
	return next == '-' ||
		next == '_' ||
		next == '.' ||
		next == '!' ||
		next == '~' ||
		next == '*' ||
		next == '\'' ||
		next == '(' ||
		next == ')'
}

func (this *URLParser) IsUnreserved(next byte) bool {
	return this.GetLexer().IsAlpha(next) || this.GetLexer().IsDigit(next) || this.IsMark(next)

}

func (this *URLParser) IsReservedNoSlash(next byte) bool {
	return next == ';' ||
		next == '?' ||
		next == ':' ||
		next == '@' ||
		next == '&' ||
		next == '+' ||
		next == '$' ||
		next == ','
}

// Missing '=' bug in character set - discovered by interop testing
// at SIPIT 13 by Bob Johnson and Scott Holben.
func (this *URLParser) IsUserUnreserved(la byte) bool {
	return la == '&' ||
		la == '=' ||
		la == '+' ||
		la == '$' ||
		la == ',' ||
		la == ';' ||
		la == '?' ||
		la == '/'
}

func (this *URLParser) Unreserved() (s string, ParseException error) {
	next, _ := this.GetLexer().LookAheadK(0)
	if this.IsUnreserved(next) {
		this.GetLexer().ConsumeK(1)

		var retval bytes.Buffer
		retval.WriteByte(next)
		return retval.String(), nil
	} else {
		return "", this.CreateParseException("unreserved")
	}
}

/** Name or value of a parameter.
 */
func (this *URLParser) ParamNameOrValue() (s string, ParseException error) {
	var retval bytes.Buffer
	for this.GetLexer().HasMoreChars() {
		next, _ := this.GetLexer().LookAheadK(0)
		if next == '[' || next == '[' || next == '/' ||
			next == ':' || next == '&' || next == '+' ||
			next == '$' || this.IsUnreserved(next) {
			retval.WriteByte(next)
			this.GetLexer().ConsumeK(1)
		} else if this.IsEscaped() {
			esc := this.GetLexer().NCharAsString(3)
			this.GetLexer().ConsumeK(3)
			retval.WriteString(esc)
		} else {
			break
		}
	}
	return retval.String(), nil
}

func (this *URLParser) UriParam() (nv *core.NameValue, ParseException error) {
	//if (debug) dbg_enter("uriParam");
	//try {
	pname, _ := this.ParamNameOrValue()
	next, _ := this.GetLexer().LookAheadK(0)
	if next == '=' {
		this.GetLexer().ConsumeK(1)
		pvalue, _ := this.ParamNameOrValue()
		return core.NewNameValue(pname, pvalue), nil
	} else {
		//println(pname + "=" + pvalue);
		return core.NewNameValue(pname, nil), nil
	}
	//} finally {
	//   if (debug) dbg_leave("uriParam");
	//}
}

func (this *URLParser) IsReserved(next byte) bool {
	return next == ';' ||
		next == '/' ||
		next == '?' ||
		next == ':' ||
		next == '@' ||
		next == '&' ||
		next == '+' ||
		next == '$' ||
		next == ','
}

func (this *URLParser) Reserved() (s string, ParseException error) {
	next, _ := this.GetLexer().LookAheadK(0)
	if this.IsReserved(next) {
		this.GetLexer().ConsumeK(1)
		var retval bytes.Buffer
		retval.WriteByte(next)
		return retval.String(), nil
	} else {
		return "", this.CreateParseException("reserved")
	}
}

func (this *URLParser) IsEscaped() bool {
	//try {
	next, _ := this.GetLexer().LookAheadK(0)
	next1, _ := this.GetLexer().LookAheadK(1)
	next2, _ := this.GetLexer().LookAheadK(2)
	return next == '%' && this.GetLexer().IsHexDigit(next1) && this.GetLexer().IsHexDigit(next2)
	//} catch (Exception ex) {
	//    return false;
	//}
}

func (this *URLParser) Escaped() (s string, ParseException error) {
	//if (debug) dbg_enter("escaped");
	//try {
	var retval bytes.Buffer
	next, _ := this.GetLexer().LookAheadK(0)
	next1, _ := this.GetLexer().LookAheadK(1)
	next2, _ := this.GetLexer().LookAheadK(2)
	if next == '%' && this.GetLexer().IsHexDigit(next1) && this.GetLexer().IsHexDigit(next2) {
		this.GetLexer().ConsumeK(3)
		retval.WriteByte(next)
		retval.WriteByte(next1)
		retval.WriteByte(next2)
		return retval.String(), nil
	} else {
		return "", this.CreateParseException("escaped")
	}

	//} finally {
	//    if (debug) dbg_leave("escaped");
	//}
}

func (this *URLParser) Mark() (s string, ParseException error) {
	//if (debug) dbg_enter("mark");
	//try {
	next, _ := this.GetLexer().LookAheadK(0)
	if this.IsMark(next) {
		this.GetLexer().ConsumeK(1)
		var retval bytes.Buffer
		retval.WriteByte(next)
		return retval.String(), nil
	} else {
		return "", this.CreateParseException("mark")
	}
	//} finally {
	//    if (debug) dbg_leave("mark");
	//}
}

func (this *URLParser) Uric() string {
	//if (debug) dbg_enter("uric");
	//try {
	//try {
	la, _ := this.GetLexer().LookAheadK(0)
	if this.IsUnreserved(la) {
		this.GetLexer().ConsumeK(1)
		return this.GetLexer().CharAsString(la)
	} else if this.IsReserved(la) {
		this.GetLexer().ConsumeK(1)
		return this.GetLexer().CharAsString(la)
	} else if this.IsEscaped() {
		retval := this.GetLexer().NCharAsString(3)
		this.GetLexer().ConsumeK(3)
		return retval
	} else {
		return ""
	}
	//} catch (Exception ex) {
	//    return null;
	//}
	//} finally {
	//    if (debug) dbg_leave("uric");
	//}

}

func (this *URLParser) UricNoSlash() string {
	//if (debug) dbg_enter("uricNoSlash");
	//try {
	//     try {
	la, _ := this.GetLexer().LookAheadK(0)
	if this.IsEscaped() {
		retval := this.GetLexer().CharAsString(3)
		this.GetLexer().ConsumeK(3)
		return retval
	} else if this.IsUnreserved(la) {
		this.GetLexer().ConsumeK(1)
		return this.GetLexer().CharAsString(la)
	} else if this.IsReservedNoSlash(la) {
		this.GetLexer().ConsumeK(1)
		return this.GetLexer().CharAsString(la)
	} else {
		return ""
	}
	// } catch (ParseException ex) {
	//     return null;
	// }
	//} finally {
	//     if (debug) dbg_leave("uricNoSlash");
	//}
}

func (this *URLParser) UricString() string {
	var retval bytes.Buffer
	for {
		next := this.Uric()
		if next == "" {
			break
		}
		retval.WriteString(next)
	}
	return retval.String()
}

/** Parse and return a structure for a generic URL.
 * Note that non SIP URLs are just stored as a string (not parsed).
 *@return URI is a URL structure for a SIP url.
 *@throws ParsException if there was a problem parsing.
 */
func (this *URLParser) UriReference() (url address.URI, ParseException error) {
	///if (debug) dbg_enter("uriReference");
	var retval address.URI
	vect, _ := this.GetLexer().PeekNextTokenK(2)
	t1 := vect[0]
	t2 := vect[1]
	//try {
	//println("URLParser::UriReference():"+this.GetLexer().GetRest());
	//println("t2 :" + t2.GetTokenValue());
	//println("t1 :" + t1.GetTokenValue());
	//fmt.Printf("tokenval = %d\n", t1.GetTokenType());

	if t1.GetTokenType() == TokenTypes_SIP {
		if t2.GetTokenType() == ':' {
			retval, _ = this.SipURL()
		} else {
			return nil, this.CreateParseException("Expecting ':'")
		}
	} else if t1.GetTokenType() == TokenTypes_TEL {
		if t2.GetTokenType() == ':' {
			retval, _ = this.TelURL()
		} else {
			return nil, this.CreateParseException("Expecting ':'")
		}
	} else {
		//println("i'm UriString()");
		urlString := this.UricString()
		//try {
		retval = address.NewURIImpl(urlString)
		//}
		//catch (ParseException ex) {
		//   throw createParseException(ex.getMessage());
		//}
	}
	//} finally {
	//    if (debug) dbg_leave("uriReference");
	//}
	return retval, nil
}

/** SIPParser for the base phone number.
 */
func (this *URLParser) Base_phone_number() (s string, ParseException error) {
	var retval bytes.Buffer //new StringBuffer() ;

	//if (debug) dbg_enter("base_phone_number");
	//try {
	lc := 0
	for this.GetLexer().HasMoreChars() {
		w, _ := this.GetLexer().LookAheadK(0)
		if this.GetLexer().IsDigit(w) || w == '-' || w == '.' || w == '(' || w == ')' {
			this.GetLexer().ConsumeK(1)
			retval.WriteByte(w)
			lc++
		} else if lc > 0 {
			break
		} else {
			return "", this.CreateParseException("unexpected " + string(w))
		}
	}
	return retval.String(), nil
	//} finally {
	//    if (debug) dbg_leave("base_phone_number");
	//}
}

/** SIPParser for the local phone #.
 */
func (this *URLParser) Local_number() (s string, ParseException error) {
	var retval bytes.Buffer //StringBuffer s = new StringBuffer() ;

	//if (debug) dbg_enter("local_number");
	//try {
	lc := 0
	for this.GetLexer().HasMoreChars() {
		la, _ := this.GetLexer().LookAheadK(0)
		if la == '*' || la == '#' || la == '-' ||
			la == '.' || la == '(' || la == ')' ||
			this.GetLexer().IsDigit(la) {
			this.GetLexer().ConsumeK(1)
			retval.WriteByte(la)
			lc++
		} else if lc > 0 {
			break
		} else {
			return "", this.CreateParseException("unexepcted " + string(la))
		}
	}
	return retval.String(), nil
	//} finally {
	//    if (debug) dbg_leave("local_number");
	//}

}

/** SIPParser for telephone subscriber.
 *
 *@return the parsed telephone number.
 */
func (this *URLParser) ParseTelephoneNumber() (tn *address.TelephoneNumber, ParseException error) {
	//TelephoneNumber tn  ;
	tn = address.NewTelephoneNumber()
	//if (debug) dbg_enter("telephone_subscriber");
	this.GetLexer().SelectLexer("charLexer")
	//try {
	c, _ := this.GetLexer().LookAheadK(0)
	if c == '+' {
		tn, _ = this.Global_phone_number()
	} else if this.GetLexer().IsAlpha(c) || this.GetLexer().IsDigit(c) ||
		c == '-' || c == '*' || c == '.' ||
		c == '(' || c == ')' || c == '#' {
		tn, _ = this.Local_phone_number()
	} else {
		return nil, this.CreateParseException("unexpected char " + string(c))
	}

	return tn, nil
	//} finally {
	//    if (debug) dbg_leave("telephone_subscriber");
	//}

}

func (this *URLParser) Global_phone_number() (tn *address.TelephoneNumber, ParseException error) {
	//if (debug) dbg_enter("global_phone_number");
	//try {
	tn = address.NewTelephoneNumber()
	tn.SetGlobal(true)
	//var nv NameValueList;
	this.GetLexer().Match(core.CORELEXER_PLUS)
	b, _ := this.Base_phone_number()
	tn.SetPhoneNumber(b)
	if this.GetLexer().HasMoreChars() {
		tok, _ := this.GetLexer().LookAheadK(0)
		if tok == ';' {
			this.GetLexer().ConsumeK(1)
			nv, _ := this.Tel_parameters()
			tn.SetParameters(nv)
		}
	}
	return tn, nil
	//} finally {
	//    if (debug) dbg_leave("global_phone_number");
	//}
}

func (this *URLParser) Local_phone_number() (tn *address.TelephoneNumber, ParseException error) {
	//if (debug) dbg_enter("local_phone_number");
	tn = address.NewTelephoneNumber()
	tn.SetGlobal(false)
	//NameValueList nv = null;
	//String b = null;
	//try {
	b, _ := this.Local_number()
	tn.SetPhoneNumber(b)
	if this.GetLexer().HasMoreChars() {
		tok, _ := this.GetLexer().PeekNextToken()
		switch tok.GetTokenType() {
		case TokenTypes_SEMICOLON:
			{
				this.GetLexer().ConsumeK(1)
				nv, _ := this.Tel_parameters()
				tn.SetParameters(nv)
				//break;
			}
		default:
			{
				//break;
			}
		}
	}
	//} finally {
	//    if (debug) dbg_leave("local_phone_number");
	//}
	return tn, nil
}

func (this *URLParser) Tel_parameters() (nvl *core.NameValueList, ParseException error) {
	nvList := core.NewNameValueList("tel_parameters")
	for {
		nv := this.NameValue('=')
		nvList.AddNameValue(nv)
		tok, _ := this.GetLexer().LookAheadK(0)
		if tok == ';' {
			continue
		} else {
			break
		}
	}
	return nvList, nil
}

/** Parse and return a structure for a Tel URL.
 *@return a parsed tel url structure.
 */
func (this *URLParser) TelURL() (telUrl *address.TelURLImpl, ParseException error) {
	this.GetLexer().Match(TokenTypes_TEL)
	this.GetLexer().Match(':')
	tn, _ := this.ParseTelephoneNumber()
	telUrl = address.NewTelURLImpl()
	telUrl.SetTelephoneNumber(tn)
	return telUrl, nil

}

/** Parse and return a structure for a SIP URL.
 *@return a URL structure for a SIP url.
 *@throws ParsException if there was a problem parsing.
 */

func (this *URLParser) SipURL() (sipurl *address.SipURIImpl, ParseException error) {
	//if (debug) dbg_enter("sipURL");
	retval := address.NewSipURIImpl()
	//try{
	this.GetLexer().Match(TokenTypes_SIP)
	this.GetLexer().Match(':')
	retval.SetScheme(core.SIPTransportNames_SIP) //TokenNames_SIP);
	//m := this.GetLexer().MarkInputPosition();
	//println("sipulr" + this.GetLexer().GetRest())

	buffer := this.GetLexer().GetRest()
	if n := strings.Index(buffer, "@"); n == -1 {
		// hostPort
		//fmt.Printf("0:%s\n", this.GetLexer().GetRest())
		hnp := core.NewHostNameParserFromLexer(this.GetLexer())
		hp, _ := hnp.GetHostPort()
		retval.SetHostPort(hp)
	} else {
		if n = strings.Index(buffer, ":"); n == -1 {
			// name@hostPort
			//fmt.Printf("1:%s\n", this.GetLexer().GetRest())
			user, _ := this.User()
			//fmt.Printf("1.1:%s\n", this.GetLexer().GetRest())
			this.GetLexer().Match('@')
			//fmt.Printf("1.2:%s\n", this.GetLexer().GetRest())
			hnp := core.NewHostNameParserFromLexer(this.GetLexer())
			hp, _ := hnp.GetHostPort()
			//fmt.Printf("1.3:%s, %s\n",this.GetLexer().GetRest(), hp.String());
			retval.SetUser(user)
			retval.SetHostPort(hp)
		} else {
			//fmt.Printf("2:%s\n", this.GetLexer().GetRest())
			user, _ := this.User()
			//la,_ := this.GetLexer().LookAheadK(0);
			// name:password@hostPort
			//fmt.Printf("3:%s\n", this.GetLexer().GetRest())
			this.GetLexer().Match(':')
			password, _ := this.Password()
			//fmt.Printf("4:%s\n", this.GetLexer().GetRest())
			this.GetLexer().Match('@')
			//fmt.Printf("5:%s\n", this.GetLexer().GetRest())
			hnp := core.NewHostNameParserFromLexer(this.GetLexer())
			hp, _ := hnp.GetHostPort()
			retval.SetUser(user)
			retval.SetUserPassword(password)
			retval.SetHostPort(hp)
		}
	}
	//println(this.GetLexer().GetRest());
	this.GetLexer().SelectLexer("charLexer")
	for this.GetLexer().HasMoreChars() {
		if la, _ := this.GetLexer().LookAheadK(0); la != ';' {
			break
		}
		this.GetLexer().ConsumeK(1)
		parms, _ := this.UriParam()
		retval.SetUriParameter(parms)
	}

	if la, _ := this.GetLexer().LookAheadK(0); this.GetLexer().HasMoreChars() && la == '?' {
		this.GetLexer().ConsumeK(1)
		for this.GetLexer().HasMoreChars() {
			parms, _ := this.Qheader()
			retval.SetQHeader(parms)
			if la, _ = this.GetLexer().LookAheadK(0); this.GetLexer().HasMoreChars() && la != '&' {
				break
			} else {
				this.GetLexer().ConsumeK(1)
			}
		}
	}
	return retval, nil
	//} finally {
	//    if (debug) dbg_leave("sipURL");
	//}
}

func (this *URLParser) PeekScheme() (s string, ParseException error) {
	tokens, _ := this.GetLexer().PeekNextTokenK(1)
	if len(tokens) == 0 {
		return "", nil
	}
	scheme := tokens[0].GetTokenValue()
	return scheme, nil
}

/** Get a name value for a given query header (ie one that comes
 * after the ?).
 */
func (this *URLParser) Qheader() (nv *core.NameValue, ParseException error) {
	name, _ := this.GetLexer().GetNextTokenByDelim('=')
	this.GetLexer().ConsumeK(1)
	value, _ := this.Hvalue()
	return core.NewNameValue(name, value), nil

}

func (this *URLParser) Hvalue() (s string, ParseException error) {
	var retval bytes.Buffer //= new StringBuffer();
	for this.GetLexer().HasMoreChars() {
		la, _ := this.GetLexer().LookAheadK(0)
		// Look for a character that can terminate a URL.
		if la == '+' || la == '?' || la == ':' ||
			la == '[' || la == ']' || la == '/' || la == '$' ||
			la == '_' || la == '-' || la == '"' || la == '!' ||
			la == '~' || la == '*' || la == '.' || la == '(' ||
			la == ')' || this.GetLexer().IsAlpha(la) || this.GetLexer().IsDigit(la) {
			this.GetLexer().ConsumeK(1)
			retval.WriteByte(la)
		} else if la == '%' {
			str, _ := this.Escaped()
			retval.WriteString(str)
		} else {
			break
		}
	}
	return retval.String(), nil
}

/** Scan forward until you hit a terminating character for a URL.
 * We do not handle non sip urls in this implementation.
 *@return the string that takes us to the end of this URL (i.e. to
 * the next delimiter).
 */
func (this *URLParser) UrlString() (s string, ParseException error) {
	var retval bytes.Buffer //StringBuffer retval = new StringBuffer();
	this.GetLexer().SelectLexer("charLexer")

	for this.GetLexer().HasMoreChars() {
		la, _ := this.GetLexer().LookAheadK(0)
		// Look for a character that can terminate a URL.
		if la == ' ' || la == '\t' || la == '\n' ||
			la == '>' || la == '<' {
			break
		}
		this.GetLexer().ConsumeK(0)
		retval.WriteByte(la)
	}
	return retval.String(), nil
}

func (this *URLParser) User() (s string, ParseException error) {

	//if (debug) dbg_enter("user");
	//try {
	var retval bytes.Buffer //StringBuffer retval = new StringBuffer();
	for this.GetLexer().HasMoreChars() {
		la, _ := this.GetLexer().LookAheadK(0)
		//if (la == '=') break;
		if this.IsUnreserved(la) ||
			this.IsUserUnreserved(la) {
			retval.WriteByte(la)
			this.GetLexer().ConsumeK(1)
		} else if this.IsEscaped() {
			//print(this.GetLexer().GetRest())
			esc := this.GetLexer().NCharAsString(3)
			//print(esc)
			this.GetLexer().ConsumeK(3)
			retval.WriteString(esc)
		} else {
			break
		}
	}
	return retval.String(), nil
	//}  finally {
	//	if (debug) dbg_leave("user");
	//}
}

func (this *URLParser) Password() (s string, ParseException error) {
	var retval bytes.Buffer //StringBuffer retval = new StringBuffer();
	for {
		la, _ := this.GetLexer().LookAheadK(0)
		if this.IsUnreserved(la) || la == '&' || la == '=' ||
			la == '+' || la == '$' || la == ',' {
			retval.WriteByte(la)
			this.GetLexer().ConsumeK(1)
		} else if this.IsEscaped() {
			esc := this.GetLexer().NCharAsString(3)
			retval.WriteString(esc)
		} else {
			break
		}
	}
	return retval.String(), nil

}

/** Default parse method. This method just calls uriReference.
 */

func (this *URLParser) Parse() (url address.URI, ParseException error) {
	return this.UriReference()
}
