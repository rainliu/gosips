package header

import (
	"bytes"
	"errors"
	"gosips/core"
	"strconv"
)

/**
* Via Header (these are strung together in a ViaList).
 */

type Via struct {
	Parameters

	/** sentProtocol field.
	 */
	sentProtocol *Protocol

	/** sentBy field.
	 */
	sentBy *core.HostPort

	/** comment field
	 */
	comment string
}

/** Default constructor
 */
func NewVia() *Via {
	this := &Via{}

	this.Parameters.super(core.SIPHeaderNames_VIA)
	this.sentProtocol = NewProtocol()

	return this
}

func (this *Via) super(name string) {
	this.Parameters.super(name)
	this.sentProtocol = NewProtocol()
}

/**
 *Compare two via headers for equaltiy.
 * @param other Object to set.
 * @return true if the two via headers are the same.
 */
/*func (this *Via) equals(Object other) {
    if (! this.getClass().equals(other.getClass())) {
        return false;
    }
    Via that = (Via) other;

    if (! this.sentProtocol.equals(that.sentProtocol)) {
        return false;
    }
    if ( ! this.sentBy.equals(that.sentBy)) {
        return false;
    }
    return true;
}*/

/** get the Protocol Version
 * @return String
 */
func (this *Via) GetProtocolVersion() string {
	if this.sentProtocol == nil {
		return ""
	} else {
		return this.sentProtocol.GetProtocolVersion()
	}
}

/**
 * Accessor for the sentProtocol field.
 * @return Protocol field
 */
func (this *Via) GetSentProtocol() *Protocol {
	return this.sentProtocol
}

/**
 * Accessor for the sentBy field
 *@return SentBy field
 */
func (this *Via) GetSentBy() *core.HostPort {
	return this.sentBy
}

/**
 * Accessor for the parameters field
 * @return parameters field
 */
func (this *Via) GetViaParms() *core.NameValueList {
	return this.parameters
}

/**
 * Accessor for the comment field.
 * @return comment field.
 */
func (this *Via) GetComment() string {
	return this.comment
}

/**
 *  Get the maddr parameter if it exists.
 * @return maddr parameter.
 */
func (this *Via) GetMaddr() *core.Host {
	return this.parameters.GetValue(ParameterNames_MADDR).(*core.Host)
}

/** port of the Via Header.
 * @return true if Port exists.
 */
func (this *Via) HasPort() bool {
	if this.sentBy == nil {
		return false
	}
	return this.sentBy.HasPort()
}

/** comment of the Via Header.
 *
 * @return false if comment does not exist and true otherwise.
 */
func (this *Via) HasComment() bool {
	return this.comment != ""
}

/** remove the port.
 */
func (this *Via) RemovePort() {
	this.sentBy = nil
}

/** remove the comment field.
 */
func (this *Via) RemoveComment() {
	this.comment = ""
}

/** set the Protocol Version
 * @param protocolVersion String to set
 */
func (this *Via) SetProtocolVersion(protocolVersion string) {
	if this.sentProtocol == nil {
		this.sentProtocol = NewProtocol()
	}
	this.sentProtocol.SetProtocolVersion(protocolVersion)
}

/** set the Host of the Via Header
 * @param host String to set
 */
func (this *Via) SetHost(host *core.Host) {
	if this.sentBy == nil {
		this.sentBy = core.NewHostPort()
	}
	this.sentBy.SetHost(host)
}

/**
 * Set the sentProtocol member
 * @param s Protocol to set.
 */
func (this *Via) SetSentProtocol(s *Protocol) {
	this.sentProtocol = s
}

/**
 * Set the sentBy member
 * @param s HostPort to set.
 */
func (this *Via) SetSentBy(s *core.HostPort) {
	this.sentBy = s
}

/**
 * Set the comment member
 * @param c String to set.
 */
func (this *Via) SetComment(c string) {
	this.comment = c
}

/** Encode the body of this header (the stuff that follows headerName).
 * A.K.A headerValue.
 */
func (this *Via) EncodeBody() string {
	var encoding bytes.Buffer
	encoding.WriteString(this.sentProtocol.String())
	encoding.WriteString(core.SIPSeparatorNames_SP)
	encoding.WriteString(this.sentBy.String())
	// Add the default port if there is no port specified.
	// if !this.sentBy.HasPort() {
	// 	encoding.WriteString(core.SIPSeparatorNames_COLON)
	// 	encoding.WriteString("5060")
	// }
	if this.comment != "" {
		encoding.WriteString(core.SIPSeparatorNames_SP)
		encoding.WriteString(core.SIPSeparatorNames_LPAREN)
		encoding.WriteString(this.comment)
		encoding.WriteString(core.SIPSeparatorNames_RPAREN)
	}
	if this.parameters.Len() > 0 {
		encoding.WriteString(core.SIPSeparatorNames_SEMICOLON)
		encoding.WriteString(this.parameters.String())
	}
	return encoding.String()
}

/**
 * Set the host part of this ViaHeader to the newly supplied <code>host</code>
 * parameter.
 *
 * @return host - the new interger value of the host of this ViaHeader
 * @throws ParseException which signals that an error has been reached
 * unexpectedly while parsing the host value.
 */
func (this *Via) SetHostFromString(host string) (ParseException error) {
	if this.sentBy == nil {
		this.sentBy = core.NewHostPort()
	}
	//try {
	h := core.NewHost(host)
	this.sentBy.SetHost(h)
	//}
	//catch(Exception e) {
	//    throw new  NullPointerException(" host parameter is null");
	//}
	return nil
}

/**
 * Returns the host part of this ViaHeader.
 *
 * @return  the string value of the host
 */
func (this *Via) GetHost() string {
	if this.sentBy == nil {
		return ""
	} else {
		host := this.sentBy.GetHost()
		if host == nil {
			return ""
		} else {
			return host.GetHostName()
		}
	}
}

/**
 * Set the port part of this ViaHeader to the newly supplied <code>port</code>
 * parameter.
 *
 * @param port - the new interger value of the port of this ViaHeader
 */
func (this *Via) SetPort(port int) {
	if this.sentBy == nil {
		this.sentBy = core.NewHostPort()
	}
	this.sentBy.SetPort(port)
}

/**
 * Returns the port part of this ViaHeader.
 *
 * @return the integer value of the port
 */
func (this *Via) GetPort() int {
	if this.sentBy == nil {
		return -1
	}
	return this.sentBy.GetPort()
}

/**
 * Returns the value of the transport parameter.
 *
 * @return the string value of the transport paramter of the ViaHeader
 */
func (this *Via) GetTransport() string {
	if this.sentProtocol == nil {
		return ""
	}
	return this.sentProtocol.GetTransport()
}

/**
 * Sets the value of the transport. This parameter specifies
 * which transport protocol to use for sending requests and responses to
 * this entity. The following values are defined: "udp", "tcp", "sctp",
 * "tls", but other values may be used also.
 *
 * @param transport - new value for the transport parameter
 * @throws ParseException which signals that an error has been reached
 * unexpectedly while parsing the transport value.
 */
func (this *Via) setTransport(transport string) (ParseException error) {
	if transport == "" {
		return errors.New("NullPointerException: GoSIP Exception, Via, setTransport(), the transport parameter is null.")
	}
	if this.sentProtocol == nil {
		this.sentProtocol = NewProtocol()
	}
	this.sentProtocol.SetTransport(transport)
	return nil
}

/**
 * Returns the value of the protocol used.
 *
 * @return the string value of the protocol paramter of the ViaHeader
 */
func (this *Via) GetProtocol() string {
	if this.sentProtocol == nil {
		return ""
	}
	return this.sentProtocol.GetProtocolName()
}

/**
 * Sets the value of the protocol parameter. This parameter specifies
 * which protocol is used, for example "SIP/2.0".
 *
 * @param protocol - new value for the protocol parameter
 * @throws ParseException which signals that an error has been reached
 * unexpectedly while parsing the protocol value.
 */
func (this *Via) SetProtocol(protocol string) (ParseException error) {
	if protocol == "" {
		return errors.New("NullPointerException: GoSIP Exception, Via, setProtocol(), the protocol parameter is null.")
	}
	if this.sentProtocol == nil {
		this.sentProtocol = NewProtocol()
	}
	this.sentProtocol.SetProtocolName(protocol)
	return nil
}

/**
 * Returns the value of the ttl parameter, or -1 if this is not set.
 *
 * @return the integer value of the <code>ttl</code> parameter
 */
func (this *Via) GetTTL() int {
	ttl, _ := strconv.Atoi(this.GetParameter(ParameterNames_TTL))
	return ttl
}

/**
 * Sets the value of the ttl parameter. The ttl parameter specifies the
 * time-to-live value when packets are sent using UDP multicast.
 *
 * @param ttl - new value of the ttl parameter
 * @throws InvalidArgumentException if supplied value is less than zero or
 * greater than 255, excluding -1 the default not set value.
 */
func (this *Via) SetTTL(ttl int) (InvalidArgumentException error) {
	if ttl < 0 && ttl != -1 {
		return errors.New("InvalidArgumentException: GoSIP Exception, Via, setTTL(), the ttl parameter is < 0")
	}
	this.SetParameter(ParameterNames_TTL, strconv.Itoa(ttl))
	return nil
}

/**
 * Returns the value of the <code>maddr</code> parameter, or null if this
 * is not set.
 *
 * @return the string value of the maddr parameter
 */
func (this *Via) GetMAddr() string {
	return this.GetParameter(ParameterNames_MADDR)
}

/**
 * Sets the value of the <code>maddr</code> parameter of this ViaHeader. The
 * maddr parameter indicates the server address to be contacted for this
 * user, overriding any address derived from the host field.
 *
 * @param  method - new value of the <code>maddr</code> parameter
 * @throws ParseException which signals that an error has been reached
 * unexpectedly while parsing the mAddr value.
 */
func (this *Via) SetMAddr(mAddr string) (ParseException error) {
	if mAddr == "" {
		return errors.New("NullPointerException: GoSIP Exception, Via, setMAddr(), the mAddr parameter is null.")
	}
	host := core.NewHost(mAddr)
	//host.SetAddress(mAddr)
	//println(host.String())
	this.SetParameter(ParameterNames_MADDR, host.String())
	return nil
}

/**
 * Gets the received paramater of the ViaHeader. Returns null if received
 * does not exist.
 *
 * @return the string received value of ViaHeader
 */
func (this *Via) GetReceived() string {
	return this.GetParameter(ParameterNames_RECEIVED)
}

/**
 * Sets the received parameter of ViaHeader.
 *
 * @param received - the newly supplied received parameter.
 * @throws ParseException which signals that an error has been reached
 * unexpectedly while parsing the received value.
 */
func (this *Via) SetReceived(received string) (ParseException error) {
	if received == "" {
		return errors.New("NullPointerException: GoSIP Exception, Via, setReceived(), the received parameter is null.")
	}
	this.SetParameter(ParameterNames_RECEIVED, received)
	return nil
}

/**
 * Gets the branch paramater of the ViaHeader. Returns null if branch
 * does not exist.
 *
 * @return the string branch value of ViaHeader
 */
func (this *Via) GetBranch() string {
	return this.GetParameter(ParameterNames_BRANCH)
}

/**
 * Sets the branch parameter of the ViaHeader to the newly supplied
 * branch value.
 *
 * @param branch - the new string branch parmameter of the ViaHeader.
 * @throws ParseException which signals that an error has been reached
 * unexpectedly while parsing the branch value.
 */
func (this *Via) SetBranch(branch string) (ParseException error) {
	if branch == "" {
		return errors.New("NullPointerException: GoSIP Exception, Via, setBranch(), the branch parameter is null.")
	}
	this.SetParameter(ParameterNames_BRANCH, branch)
	return nil
}
