package parser

import (
	"bytes"
	"errors"
	"gosips/core"
	"gosips/sip/address"
	"gosips/sip/header"
	"gosips/sip/message"
	"strings"
)

/**
 * Parse SIP message and parts of SIP messages such as URI's etc
 * from memory and return a structure.
 * Intended use:  UDP message processing.
 * This class is used when you have an entire SIP message or SIPHeader
 * or SIP URL in memory and you want to generate a parsed structure from
 * it. For SIP messages, the payload can be binary or String.
 * If you have a binary payload,
 * use parseSIPMessage(byte[]) else use parseSIPMessage(String)
 * The payload is accessible from the parsed message using the GetContent and
 * GetContentBytes methods provided by the SIPMessage class. If SDP parsing
 * is enabled using the parseContent method, then the SDP body is also parsed
 * and can be accessed from the message using the GetSDPAnnounce method.
 * Currently only eager parsing of the message is supported (i.e. the
 * entire message is parsed in one feld swoop).
 */

type StringMsgParser struct {
	readBody bool

	rawMessage     string // Unprocessed message  (for error reporting)
	rawMessage1    string // Unprocessed message  (for error reporting)
	currentMessage string // the message being parsed. (for error reporting)

	//private ParseExceptionListener parseExceptionListener;

	messageHeaders map[int]string // Message headers

	bufferPointer int

	bodyIsString bool

	currentMessageBytes []byte

	contentLength int

	debugFlag bool

	currentLine int

	currentHeader string
}

/**
 *Constructor
 */
func NewStringMsgParser() *StringMsgParser {
	this := &StringMsgParser{}
	this.bufferPointer = 0
	this.currentLine = 0
	this.readBody = true
	return this
}

/** Get the message body.
 */
func (this *StringMsgParser) GetMessageBody() string {
	if this.contentLength == 0 {
		return ""
	} else {
		endIndex := this.bufferPointer + this.contentLength
		var body string
		// guard against bad specifications.
		if endIndex > len(this.currentMessage) {
			endIndex = len(this.currentMessage)
			body = this.currentMessage[this.bufferPointer:endIndex]
			this.bufferPointer = endIndex
		} else {
			body = this.currentMessage[this.bufferPointer:endIndex]
			this.bufferPointer = endIndex + 1
		}
		this.contentLength = 0
		return body
	}

}

/** Get the message body as a byte array.
 */
func (this *StringMsgParser) GetBodyAsBytes() []byte {
	if this.contentLength == 0 {
		return nil
	} else {
		endIndex := this.bufferPointer + this.contentLength
		// guard against bad specifications.
		if endIndex > len(this.currentMessageBytes) {
			endIndex = len(this.currentMessageBytes)
		}
		body := make([]byte, endIndex-this.bufferPointer)

		copy(body, this.currentMessageBytes[this.bufferPointer:endIndex])

		this.bufferPointer = endIndex
		this.contentLength = 0
		return body
	}

}

/** Return the contents till the end of the buffer (this is useful when
 * you encounter an error.
 */
func (this *StringMsgParser) ReadToEnd() string {
	body := this.currentMessage[this.bufferPointer:]
	this.bufferPointer += len(body)
	return body
}

/** Return tbe bytes to the end of the message.
 * This is invoked when the parser is invoked with an array of bytes
 * rather than with a string.
 */
func (this *StringMsgParser) ReadBytesToEnd() []byte {
	body := make([]byte, len(this.currentMessageBytes)-this.bufferPointer)
	endIndex := len(this.currentMessageBytes)
	copy(body, this.currentMessageBytes[this.bufferPointer:endIndex])
	this.bufferPointer = endIndex
	this.contentLength = 0
	return body
}

/**
 * add a handler for header parsing errors.
 * @param  pexhadler is a class
 *  	that implements the ParseExceptionListener interface.
 */
// public void SetParseExceptionListener
// ( ParseExceptionListener pexhandler ) {
//     parseExceptionListener = pexhandler;
// }

/** Return true if the body is encoded as a string.
 * If the parseSIPMessage(String) method is invoked then the body
 * is assumed to be a string.
 */
func (this *StringMsgParser) IsBodyString() bool {
	return this.bodyIsString
}

/** Parse a buffer containing a single SIP Message where the body
 * is an array of un-interpreted bytes. This is intended for parsing
 * the message from a memory buffer when the buffer.
 * Incorporates a bug fix for a bug that was noted by Will Sullin of
 * Callcast
 * @param msgBuffer a byte buffer containing the messages to be parsed.
 *   This can consist of multiple SIP Messages concatenated toGether.
 * @return a SIPMessage[] structure (request or response)
 * 			containing the parsed SIP message.
 * @exception SIPIllegalMessageException is thrown when an
 * 			illegal message has been encountered (and
 *			the rest of the buffer is discarded).
 * @see ParseExceptionListener
 */
func (this *StringMsgParser) ParseSIPMessageFromByte(msgBuffer []byte) (message.Message, error) {
	this.bufferPointer = 0
	this.bodyIsString = false

	this.currentMessageBytes = msgBuffer
	var s int
	// Squeeze out leading CRLF
	// Squeeze out the leading nulls (otherwise the parser will crash)
	// Bug noted by Will Sullin of Callcast
	for s = this.bufferPointer; s < len(msgBuffer); s++ {
		if msgBuffer[s] != '\r' && msgBuffer[s] != '\n' {
			break
		}
	}

	if s == len(msgBuffer) {
		return nil, nil
	}

	// Find the end of the SIP message.
	var f int
	for f = s; f < len(msgBuffer)-4; f++ {
		if msgBuffer[f] == '\r' &&
			msgBuffer[f+1] == '\n' &&
			msgBuffer[f+2] == '\r' &&
			msgBuffer[f+3] == '\n' {
			break
		}
	}
	if f < len(msgBuffer) {
		f += 4
	} else {
		// Could not find CRLFCRLF end of message so look for LFLF
		for f = s; f < len(msgBuffer)-2; f++ {
			if msgBuffer[f] == '\n' &&
				msgBuffer[f] == '\n' {
				break
			}
		}
		if f < len(msgBuffer) {
			f += 2
		} else {
			return nil, errors.New("ParseException: Message not terminated")
		}
	}

	// Encode the body as a UTF-8 string.
	var messageString string
	messageString = string(msgBuffer[s:f]) //, "UTF-8");

	this.bufferPointer = f
	message := []byte(messageString)
	length := len(message)
	// Get rid of CR to make it uniform for the parser.
	for k := 0; k < length; k++ {
		if message[k] == '\r' {
			copy(message[k:length-1], message[k+1:length])
			length--
		}
	}

	// The following can be written more efficiently in a single pass
	// but it is somewhat tricky.
	tokenizer := core.NewStringTokenizer(string(message[:length]))
	var cooked_message bytes.Buffer

	for tokenizer.HasMoreChars() {
		nexttok := tokenizer.NextToken()
		// Ignore blank lines with leading spaces or tabs.
		if strings.TrimSpace(nexttok) == "" {
			cooked_message.WriteString("\n")
		} else {
			cooked_message.WriteString(nexttok)
		}
	}

	message1 := cooked_message.String()
	length = strings.Index(message1, "\n\n") + 2

	// Handle continuations - look for a space or a tab at the start
	// of the line and append it to the previous line.
	var cooked_message1 bytes.Buffer
	for k := 0; k < length-1; {
		if message1[k] == '\n' {
			if message1[k+1] == '\t' || message1[k+1] == ' ' {
				k++
				k++
				if k == length {
					break
				} else {
					continue
				}
			} else {
				cooked_message1.WriteByte(message1[k])
			}

			if message1[k+1] == '\n' {
				cooked_message1.WriteByte('\n')
			}
		} else {
			cooked_message1.WriteByte(message1[k])
		}
		k++
	}
	cooked_message1.WriteString("\n\n")
	cooked_message1.WriteString(message1[length:])

	// Separate the string out into substrings for
	// error reporting.
	this.currentMessage = cooked_message1.String()

	sipmsg, err := this.ParseMessage(this.currentMessage)
	if err != nil {
		return nil, err
	}

	if this.readBody && sipmsg.GetContentLength() != nil && sipmsg.GetContentLength().GetContentLength() != 0 {
		this.contentLength = sipmsg.GetContentLength().GetContentLength()
		body := this.GetBodyAsBytes()
		sipmsg.SetMessageContentFromByte(body)
	}

	return sipmsg, nil
}

/**
 * Parse a buffer containing one or more SIP Messages  and return an array of
 * SIPMessage parsed structures. Note that the current limitation is that
 * this does not handle content encoding properly. The message content is
 * just assumed to be encoded using the same encoding as the sip message
 * itself (i.e. binary encodings such as gzip are not supported).
 * @param sipMessages a String containing the messages to be parsed.
 *   This can consist of multiple SIP Messages concatenated toGether.
 * @return a SIPMessage structure (request or response)
 * 			containing the parsed SIP message.
 * @exception SIPIllegalMessageException is thrown when an
 * 			illegal message has been encountered (and
 *			the rest of the buffer is discarded).
 * @see ParseExceptionListener
 */

func (this *StringMsgParser) ParseSIPMessage(sipMessage string) (message.Message, error) {
	return this.ParseSIPMessageFromByte([]byte(sipMessage))
}

/** This is called repeatedly by parseSIPMessage to parse
 * the contents of a message buffer. This assumes the message
 * already has continuations etc. taken care of.
 * prior to its being called.
 */
func (this *StringMsgParser) ParseMessage(currentMessage string) (message.Message, error) {
	var err error

	sip_message_size := 0 // # of lines in the sip message
	var sipmsg message.Message

	tokenizer := core.NewStringTokenizer(currentMessage)
	this.messageHeaders = make(map[int]string) // A list of headers for error reporting

	for tokenizer.HasMoreChars() {
		nexttok := tokenizer.NextToken()
		if nexttok == "\n" {
			nextnexttok := tokenizer.NextToken()
			if nextnexttok == "\n" {
				break
			} else {
				this.messageHeaders[sip_message_size] = nextnexttok
			}
		} else {
			this.messageHeaders[sip_message_size] = nexttok
		}

		sip_message_size++
	}

	this.currentLine = 0
	this.currentHeader = this.messageHeaders[this.currentLine]
	firstLine := this.currentHeader
	if !strings.HasPrefix(firstLine, header.SIPConstants_SIP_VERSION_STRING) {
		sipmsg = message.NewSIPRequest()
		var rl *header.RequestLine
		if rl, err = NewRequestLineParser(firstLine + "\n").Parse(); err != nil {
			return nil, err
		}
		sipmsg.(*message.SIPRequest).SetRequestLine(rl)
	} else {
		sipmsg = message.NewSIPResponse()
		var sl *header.StatusLine
		if sl, err = NewStatusLineParser(firstLine + "\n").Parse(); err != nil {
			return nil, err
		}
		sipmsg.(*message.SIPResponse).SetStatusLine(sl)
	}

	for i := 1; i < len(this.messageHeaders); i++ {
		hdrstring := this.messageHeaders[i]
		if hdrstring == "" || strings.TrimSpace(hdrstring) == "" {
			continue
		}

		var hdrParser Parser
		if hdrParser, err = CreateParser(hdrstring + "\n"); err != nil {
			return nil, err
		}

		var sipHeader header.Header
		if sipHeader, err = hdrParser.Parse(); err != nil {
			return nil, err
		}

		if _, ok := sipmsg.(*message.SIPRequest); ok {
			sipmsg.(*message.SIPRequest).AttachHeader2(sipHeader, false)
		} else {
			sipmsg.(*message.SIPResponse).AttachHeader2(sipHeader, false)
		}
	}
	return sipmsg, nil
}

/**
 * Parse an address (nameaddr or address spec)  and return and address
 * structure.
 * @param address is a String containing the address to be parsed.
 * @return a parsed address structure.
 * @exception  ParseException when the address is badly formatted.
 */

func (this *StringMsgParser) ParseAddress(address string) (*address.AddressImpl, error) {
	addressParser := NewAddressParser(address)
	return addressParser.Address()
}

/**
 * Parse a host:port and return a parsed structure.
 * @param hostport is a String containing the host:port to be parsed
 * @return a parsed address structure.
 * @exception throws a ParseException when the address is badly formatted.
 */
func (this *StringMsgParser) ParseHostPort(hostport string) (*core.HostPort, error) {
	lexer := NewSIPLexer("charLexer", hostport)
	return core.NewHostNameParserFromLexer(lexer).GetHostPort()
}

/**
 * Parse a host name and return a parsed structure.
 * @param host is a String containing the host name to be parsed
 * @return a parsed address structure.
 * @exception throws a ParseException when the hostname is badly formatted.
 */
func (this *StringMsgParser) ParseHost(host string) (*core.Host, error) {
	lexer := NewSIPLexer("charLexer", host)
	return core.NewHostNameParserFromLexer(lexer).GetHost()
}

/**
 * Parse a telephone number return a parsed structure.
 * @param telphone_number is a String containing
 * the telephone # to be parsed
 * @return a parsed address structure.
 * @exception throws a ParseException when the address is badly formatted.
 */
func (this *StringMsgParser) ParSetelephoneNumber(telephone_number string) (*address.TelephoneNumber, error) {
	return NewURLParser(telephone_number).ParseTelephoneNumber()
}

/**
 * Parse a  SIP url from a string and return a URI structure for it.
 * @param url a String containing the URI structure to be parsed.
 * @return A parsed URI structure
 * @exception ParseException  if there was an error parsing the message.
 */
func (this *StringMsgParser) ParseSIPUrl(url string) (*address.SipURIImpl, error) {
	sipuri, err := NewURLParser(url).Parse()
	return sipuri.(*address.SipURIImpl), err
}

/**
 * Parse a  uri from a string and return a URI structure for it.
 * @param url a String containing the URI structure to be parsed.
 * @return A parsed URI structure
 * @exception ParseException  if there was an error parsing the message.
 */
func (this *StringMsgParser) ParseUrl(url string) (*address.URIImpl, error) {
	guri, err := NewURLParser(url).Parse()
	return guri.(*address.URIImpl), err
}

/**
 * Parse an individual SIP message header from a string.
 * @param header String containing the SIP header.
 * @return a SIPHeader structure.
 * @exception ParseException  if there was an error parsing the message.
 */
func (this *StringMsgParser) ParseSIPHeader(h string) (*header.SIPHeader, error) {
	var err error

	h += "\n\n"
	// Handle line folding.
	var nmessage string
	//counter := 0
	// eat leading spaces and carriage returns (necessary??)
	i := 0
	for h[i] == '\n' || h[i] == '\t' || h[i] == ' ' {
		i++
	}
	for ; i < len(h); i++ {
		if i < len(h)-1 && (h[i] == '\n' && (h[i+1] == '\t' || h[i+1] == ' ')) {
			nmessage += " "
			i++
		} else {
			nmessage += string(h[i])
		}
	}

	nmessage += "\n"

	var hp Parser
	if hp, err = CreateParser(nmessage); err != nil {
		return nil, err
	}

	shp, err := hp.Parse()
	return shp.(*header.SIPHeader), err
}

/**
 * Parse the SIP Request Line
 * @param  requestLine a String  containing the request line to be parsed.
 * @return  a RequestLine structure that has the parsed RequestLine
 * @exception ParseException  if there was an error parsing the requestLine.
 */

func (this *StringMsgParser) ParseSIPRequestLine(requestLine string) (*header.RequestLine, error) {
	//throws ParseException {
	requestLine += "\n"
	return NewRequestLineParser(requestLine).Parse()
}

/**
 * Parse the SIP Response message status line
 * @param statusLine a String containing the Status line to be parsed.
 * @return StatusLine class corresponding to message
 * @exception ParseException  if there was an error parsing
 * @see StatusLine
 */

func (this *StringMsgParser) ParseSIPStatusLine(statusLine string) (*header.StatusLine, error) {
	//throws ParseException {
	statusLine += "\n"
	return NewStatusLineParser(statusLine).Parse()
}

/**
 * Get the current header.
 */
func (this *StringMsgParser) GetCurrentHeader() string {
	return this.currentHeader
}

/**
 * Get the current line number.
 */
func (this *StringMsgParser) GetCurrentLineNumber() int {
	return this.currentLine
}
