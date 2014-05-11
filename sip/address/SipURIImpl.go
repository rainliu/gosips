package address

import (
	"bytes"
	"container/list"
	"gosips/core"
	"strconv"
	"strings"
)

/**
* The SipUri structure.
 */
type SipUri struct {
	Uri

	/** Authority for the uri.
	 */

	authority *Authority

	/** uriParms list
	 */
	uriParms *core.NameValueList

	/** qheaders list
	 */
	qheaders *core.NameValueList

	/** telephoneSubscriber field
	 */
	telephoneSubscriber *TelephoneNumber
}

func NewSipUri() *SipUri {
	this := &SipUri{}

	this.scheme = core.SIPTransportNames_SIP
	this.uriParms = core.NewNameValueList("uriparms")
	this.qheaders = core.NewNameValueList("qheaders")
	this.qheaders.SetSeparator("&")

	return this
}

/** Constructor given the scheme.
 * The scheme must be either Sip or Sips
 */
func (this *SipUri) SetScheme(scheme string) {
	scheme = strings.ToLower(scheme)
	if scheme != core.SIPTransportNames_SIP && scheme != core.SIPTransportNames_SIPS {
		return //throw new IllegalArgumentException("bad scheme "+scheme);
	}
	this.scheme = scheme
}

/** Get the scheme.
 */
func (this *SipUri) GetScheme() string {
	return this.scheme
}

/**
 * clear all URI Parameters.
 * @since v1.0
 */
func (this *SipUri) ClearUriParms() {
	this.uriParms = core.NewNameValueList("uriparms")
}

/**
*Clear the password from the user part if it exists.
 */
func (this *SipUri) ClearPassword() {
	if this.authority != nil {
		userInfo := this.authority.GetUserInfo()
		if userInfo != nil {
			userInfo.ClearPassword()
		}
	}
}

/** Get the authority.
 */
func (this *SipUri) GetAuthority() *Authority {
	return this.authority
}

/**
 * Clear all Qheaders.
 */
func (this *SipUri) ClearQheaders() {
	this.qheaders = core.NewNameValueList("qheaders")
}

/**
 * Compare two URIs and return true if they are equal.
 * @param that the object to compare to.
 * @return true if the object is equal to this object.
 */
/*public boolean   equals ( Object that ) {

	if (that == null) return false;

        if (!this.GetClass().equals(that.GetClass())){
            return false;
        }

        SipUri other = (SipUri) that;

	// Compare the authority portion.
	if (!this.authority.equals(other.authority)) return false;

        // compare the parameter lists.
	ListIterator li = this.uriParms.listIterator();
	NameValueList hisParms = other.uriParms;
	while(li.hasNext()) {
		NameValue nv = (NameValue) li.next();
		// transport string defaults to udp.
		if (nv.GetName().equals(TRANSPORT) ) {
			String value = (String) nv.GetValue();
			String hisTransport =
			   (String) hisParms.GetValue(TRANSPORT);
			if (hisTransport == null &&
				value.compareToIgnoreCase(UDP) == 0)  {
			        continue;
			} else if ( hisTransport == null) {
				 return false;
			} else if
			    (hisTransport.compareToIgnoreCase(value) == 0)  {
				continue;
			}
		} else {
			NameValue hisnv = hisParms.GetNameValue(nv.GetName());
			if (hisnv == null) {
				return false;
			} else if (! hisnv.equals(nv)) {
				 return false;
			}
		}
	}

        // leave headers alone - they are just a screwy way of constructing
        // an entire sip message header as part of a URL.
        return true;
    }*/

/**
 * Construct a URL from the parsed structure.
 * @return String
 */
func (this *SipUri) String() string {
	var retval bytes.Buffer
	retval.WriteString(this.scheme)
	retval.WriteString(core.SIPSeparatorNames_COLON)
	if this.authority != nil {
		retval.WriteString(this.authority.String())
	}
	if this.uriParms.Len() != 0 {
		retval.WriteString(core.SIPSeparatorNames_SEMICOLON)
		retval.WriteString(this.uriParms.String())
	}
	if this.qheaders.Len() != 0 {
		retval.WriteString(core.SIPSeparatorNames_QUESTION)
		retval.WriteString(this.qheaders.String())
	}
	return retval.String()
}

/** Return a string representation.
*
*@return the String representation of this URI.
*
 */
//public String toString() { return this.String(); }

/**
 * GetUser@host
 * @return user@host portion of the uri (null if none exists).
 */
func (this *SipUri) GetUserAtHost() string {
	user := this.authority.GetUserInfo().GetUser()
	host := this.authority.GetHost().String()
	return user + core.SIPSeparatorNames_AT + host
}

/**
 * GetUser@host
 * @return user@host portion of the uri (null if none exists).
 */
func (this *SipUri) GetUserAtHostPort() string {
	var user string
	if this.authority.GetUserInfo() != nil {
		user = this.authority.GetUserInfo().GetUser()
	}
	host := this.authority.GetHost().String()
	port := this.authority.GetPort()
	// If port not set assign the default.
	var s bytes.Buffer
	if user != "" {
		s.WriteString(user)
		s.WriteString(core.SIPSeparatorNames_AT)
	}
	s.WriteString(host)
	if port != -1 {
		s.WriteString(core.SIPSeparatorNames_COLON)
		s.WriteString(strconv.Itoa(port))
	}
	//else
	return s.String()
}

/**
 * Get the parameter (do a name lookup) and return null if none exists.
 * @param parmname Name of the parameter to Get.
 * @return Parameter of the given name (null if none exists).
 */
func (this *SipUri) GetParm(parmname string) interface{} {
	return this.uriParms.GetValue(parmname)
}

/**
 * Get the method parameter.
 * @return Method parameter.
 */
func (this *SipUri) GetMethod() string {
	return this.GetParm(core.SIPTransportNames_METHOD).(string)
}

/**
 * Accessor for URI parameters
 * @return A name-value list containing the parameters.
 */
func (this *SipUri) GetUriParms() *core.NameValueList {
	return this.uriParms
}

/** Remove the URI parameters.
*
 */
func (this *SipUri) RemoveUriParms() {
	this.uriParms = core.NewNameValueList("uriparms")
}

/**
 * Accessor forSIPObjects
 * @return Get the query headers (that appear after the ? in
 * the URL)
 */
func (this *SipUri) GetQheaders() *core.NameValueList {
	return this.qheaders
}

/**
 * Get the urse parameter.
 * @return User parameter (user= phone or user=ip).
 */
func (this *SipUri) GetUserType() string {
	return this.uriParms.GetValue(core.SIPTransportNames_USER).(string)
}

/**
 * Get the password of the user.
 * @return User password when it embedded as part of the uri
 * ( a very bad idea).
 */
func (this *SipUri) GetUserPassword() string {
	if this.authority == nil {
		return ""
	}
	return this.authority.GetPassword()
}

/** Set the user password.
 *@param password - password to set.
 */
func (this *SipUri) SetUserPassword(password string) {
	if this.authority == nil {
		this.authority = NewAuthority()
	}
	this.authority.SetPassword(password)
}

/**
 * Returns the stucture corresponding to the telephone number
 * provided that the user is a telephone subscriber.
 * @return TelephoneNumber part of the url (only makes sense
 * when user = phone is specified)
 */
func (this *SipUri) GetTelephoneSubscriber() *TelephoneNumber {
	if this.telephoneSubscriber == nil {
		this.telephoneSubscriber = NewTelephoneNumber()
	}
	return this.telephoneSubscriber
}

/**
 * Get the host and port of the server.
 * @return Get the host:port part of the url parsed into a
 * structure.
 */
func (this *SipUri) GetHostPort() *core.HostPort {
	if this.authority == nil {
		return nil
	} //else {
	return this.authority.GetHostPort()
	//}
}

/** Get the port from the authority field.
*
*@return the port from the authority field.
 */
func (this *SipUri) GetPort() int {
	hp := this.GetHostPort()
	if hp == nil {
		return -1
	}
	return hp.GetPort()
}

/** Get the host protion of the URI.
* @return the host portion of the url.
 */
func (this *SipUri) GetHost() string {
	return this.authority.GetHost().String()
}

/**
 * returns true if the user is a telephone subscriber.
 *  If the host is an Internet telephony
 * gateway, a telephone-subscriber field MAY be used instead
 * of a user field. The telephone-subscriber field uses the
 * notation of RFC 2806 [19]. Any characters of the un-escaped
 * "telephone-subscriber" that are not either in the set
 * "unreserved" or "user-unreserved" MUST be escaped. The set
 * of characters not reserved in the RFC 2806 description of
 * telephone-subscriber contains a number of characters in
 * various syntax elements that need to be escaped when used
 * in SIP URLs, for example quotation marks (%22), hash (%23),
 * colon (%3a), at-sign (%40) and the "unwise" characters,
 * i.e., punctuation of %5b and above.
 *
 * The telephone number is a special case of a user name and
 * cannot be distinguished by a BNF. Thus, a URL parameter,
 * user, is added to distinguish telephone numbers from user
 * names.
 *
 * The user parameter value "phone" indicates that the user
 * part contains a telephone number. Even without this
 * parameter, recipients of SIP URLs MAY interpret the pre-@
 * part as a telephone number if local restrictions on the
 * @return true if the user is a telephone subscriber.
 */
func (this *SipUri) IsUserTelephoneSubscriber() bool {
	usrtype := this.uriParms.GetValue(core.SIPTransportNames_USER).(string)
	if usrtype == "" {
		return false
	}
	return usrtype == (core.SIPTransportNames_PHONE)
}

/**
 *Remove the ttl value from the parameter list if it exists.
 */
func (this *SipUri) RemoveTTL() {
	if this.uriParms != nil {
		this.uriParms.Delete(core.SIPTransportNames_TTL)
	}
}

/**
 *Remove the maddr param if it exists.
 */
func (this *SipUri) RemoveMAddr() {
	if this.uriParms != nil {
		this.uriParms.Delete(core.SIPTransportNames_MADDR)
	}
}

/**
 *Delete the transport string.
 */
func (this *SipUri) RemoveTransport() {
	if this.uriParms != nil {
		this.uriParms.Delete(core.SIPTransportNames_TRANSPORT)
	}
}

/** Remove a header given its name (provided it exists).
 * @param name name of the header to Remove.
 */
func (this *SipUri) RemoveHeader(name string) {
	if this.qheaders != nil {
		this.qheaders.Delete(name)
	}
}

/** Remove all headers.
 */
func (this *SipUri) RemoveHeaders() {
	this.qheaders = core.NewNameValueList("qheaders")
}

/**
 * Set the user type.
 */
func (this *SipUri) RemoveUserType() {
	if this.uriParms != nil {
		this.uriParms.Delete(core.SIPTransportNames_USER)
	}
}

/**
 *Remove the port setting.
 */
func (this *SipUri) RemovePort() {
	this.authority.RemovePort()
}

/**
 * Remove the Method.
 */
func (this *SipUri) RemoveMethod() {
	if this.uriParms != nil {
		this.uriParms.Delete(core.SIPTransportNames_METHOD)
	}
}

/** Sets the user of SipURI. The identifier of a particular resource at
 * the host being addressed. The user and the user password including the
 * "at" sign make up the user-info.
 *
 * @param user - the new String value of the user.
 * @throws ParseException which signals that an error has been reached
 * unexpectedly while parsing the user value.
 */
func (this *SipUri) SetUser(uname string) {
	if this.authority == nil {
		this.authority = NewAuthority()
	}

	this.authority.SetUser(uname)
}

/** Remove the user.
 */
func (this *SipUri) RemoveUser() {
	this.authority.RemoveUserInfo()
}

/** Set the default parameters for this URI.
 * Do nothing if the parameter is already set to some value.
 * Otherwise set it to the given value.
 * @param name Name of the parameter to set.
 * @param value value of the parameter to set.
 */
func (this *SipUri) SetDefaultParm(name string, value interface{}) {
	if this.uriParms.GetValue(name) == nil {
		nv := core.NewNameValue(name, value)
		this.uriParms.AddNameValue(nv)
	}
}

/** Set the authority member
 * @param authority Authority to set.
 */
func (this *SipUri) SetAuthority(authority *Authority) {
	this.authority = authority
}

/** Set the host for this URI.
 * @param h host to set.
 */
func (this *SipUri) SetHost(h *core.Host) {
	if this.authority == nil {
		this.authority = NewAuthority()
	}
	this.authority.SetHost(h)
}

/** Set the uriParms member
 * @param parms URI parameters to set.
 */
func (this *SipUri) SetUriParms(parms *core.NameValueList) {
	this.uriParms = parms
}

/**
         * Set a given URI parameter. Note - parameter must be properly
	 *  encoded before the function is called.
         * @param name Name of the parameter to set.
         * @param value value of the parameter to set.
*/
func (this *SipUri) SetUriParm(name string, value interface{}) {
	nv := core.NewNameValue(name, value)
	this.uriParms.AddNameValue(nv)
}

/** Set the qheaders member
 * @param parms query headers to set.
 */
func (this *SipUri) SetQheaders(parms *core.NameValueList) {
	this.qheaders = parms
}

/**
 * Set the MADDR parameter .
 * @param mAddr Host Name to set
 */
func (this *SipUri) SetMAddr(mAddr string) {
	nameValue := this.uriParms.GetNameValue(core.SIPTransportNames_MADDR)
	host := &core.Host{}
	host.SetAddress(mAddr)
	if nameValue != nil {
		nameValue.SetValue(host)
	} else {
		nameValue = core.NewNameValue(core.SIPTransportNames_MADDR, host)
		this.uriParms.AddNameValue(nameValue)
	}
}

/** Sets the value of the user parameter. The user URI parameter exists to
 * distinguish telephone numbers from user names that happen to look like
 * telephone numbers.  This is equivalent to setParameter("user", user).
 *
 * @param  userParam - new value String value of the method parameter
 */
func (this *SipUri) SetUserParam(usertype string) {
	this.uriParms.Delete(core.SIPTransportNames_USER)
	this.uriParms.AddNameAndValue(core.SIPTransportNames_USER, usertype)
}

/** set the Method
 * @param method method parameter
 */
func (this *SipUri) SetMethod(method string) {
	this.uriParms.AddNameAndValue(core.SIPTransportNames_METHOD, method)
}

/**
 * Sets ISDN subaddress of SipURL
 * @param <var>isdnSubAddress</var> ISDN subaddress
 */
func (this *SipUri) SetIsdnSubAddress(isdnSubAddress string) {
	if this.telephoneSubscriber == nil {
		this.telephoneSubscriber = NewTelephoneNumber()
	}
	this.telephoneSubscriber.SetIsdnSubaddress(isdnSubAddress)
}

/** Set the telephone subscriber field.
 * @param tel Telephone subscriber field to set.
 */
func (this *SipUri) SetTelephoneSubscriber(tel *TelephoneNumber) {
	this.telephoneSubscriber = tel
}

/** set the port to a given value.
 * @param p Port to set.
 */
func (this *SipUri) SetPort(p int) {
	if this.authority == nil {
		this.authority = NewAuthority()
	}
	this.authority.SetPort(p)
}

/** Boolean to check if a parameter of a given name exists.
 * @param name Name of the parameter to check on.
 * @return a boolean indicating whether the parameter exists.
 */
func (this *SipUri) HasParameter(name string) bool {
	return this.uriParms.GetValue(name) != nil
}

/** Set the query header when provided as a name-value pair.
 *@param qHeader - qeuery header provided as a name,value pair.
 */
func (this *SipUri) SetQHeader(nameValue *core.NameValue) {
	this.qheaders.PushBack(nameValue)
}

/** Set the parameter as given.
 *@param nameValue - parameter to set.
 */
func (this *SipUri) SetUriParameter(nameValue *core.NameValue) {
	this.uriParms.PushBack(nameValue)
}

/** Return true if the transport parameter is defined.
 * @return true if transport appears as a parameter and false otherwise.
 */
func (this *SipUri) HasTransport() bool {
	return this.HasParameter(core.SIPTransportNames_TRANSPORT)
}

/**
 * Remove a parameter given its name
 * @param name -- name of the parameter to Remove.
 */
func (this *SipUri) RemoveParameter(name string) {
	this.uriParms.Delete(name)
}

/** Set the hostPort field of the imbedded authority field.
 *@param hostPort is the hostPort to set.
 */
func (this *SipUri) SetHostPort(hostPort *core.HostPort) {
	if this.authority == nil {
		this.authority = NewAuthority()
	}
	this.authority.SetHostPort(hostPort)
}

/**/
/** clone this.
 */
func (this *SipUri) Clone() interface{} {
	retval := NewSipUri()

	retval.uriString = this.uriString
	retval.scheme = this.scheme

	if this.authority != nil {
		retval.authority = this.authority.Clone().(*Authority)
	}
	if this.uriParms != nil {
		retval.uriParms = this.uriParms.Clone().(*core.NameValueList)
	}
	if this.qheaders != nil {
		retval.qheaders = this.qheaders.Clone().(*core.NameValueList)
	}
	if this.telephoneSubscriber != nil {
		retval.telephoneSubscriber = this.telephoneSubscriber.Clone().(*TelephoneNumber)
	}

	return retval
}

/** Returns the value of the named header, or null if it is not set.
 * SIP/SIPS URIs may specify headers. As an example, the URI
 * sip:joe@jcp.org?priority=urgent has a header "priority" whose
 * value is "urgent".
 *
 * @param <var>name</var> name of header to retrieve
 * @return the value of specified header
 */
func (this *SipUri) GetHeader(name string) string {
	if this.qheaders.GetValue(name) == nil {
		return ""
	}

	return this.qheaders.GetValue(name).(string)

}

/** Returns an Iterator over the names (Strings) of all headers present
 * in this SipURI.
 *
 * @return an Iterator over all the header names
 */
func (this *SipUri) GetHeaderNames() *core.NameValueList {
	return this.qheaders
}

/** Returns the value of the <code>lr</code> parameter, or null if this
 * is not Set. This is equivalent to GetParameter("lr").
 *
 * @return the value of the <code>lr</code> parameter
 */
func (this *SipUri) GetLrParam() string {
	if this.HasParameter(core.SIPTransportNames_LR) {
		return "true"
	}
	return ""
}

/** Returns the value of the <code>maddr</code> parameter, or null if this
 * is not Set. This is equivalent to GetParameter("maddr").
 *
 * @return the value of the <code>maddr</code> parameter
 */
func (this *SipUri) GetMAddrParam() string {
	maddr := this.uriParms.GetNameValue(core.SIPTransportNames_MADDR)
	if maddr == nil {
		return ""
	}
	return maddr.GetValue().(string)
}

/** Returns the value of the <code>method</code> parameter, or null if this
 * is not Set. This is equivalent to GetParameter("method").
 *
 * @return  the value of the <code>method</code> parameter
 */
func (this *SipUri) GetMethodParam() string {
	return this.GetParameter(core.SIPTransportNames_METHOD)
}

/**
 * Returns the value of the named parameter, or null if it is not Set. A
 *
 * zero-length String indicates flag parameter.
 *
 *
 *
 * @param <var>name</var> name of parameter to retrieve
 *
 * @return the value of specified parameter
 *
 */
func (this *SipUri) GetParameter(name string) string {
	val := this.uriParms.GetValue(name)
	if val == nil {
		return ""
	}

	return val.(string)
}

/**
 * Returns an Iterator over the names (Strings) of all parameters present
 *
 * in this ParametersHeader.
 *
 *
 *
 * @return an Iterator over all the parameter names
 *
 */
func (this *SipUri) GetParameterNames() *list.List {
	return this.uriParms.GetNames()
}

/** Returns the value of the "ttl" parameter, or -1 if this is not Set.
 * This method is equivalent to GetParameter("ttl").
 *
 * @return the value of the <code>ttl</code> parameter
 */
func (this *SipUri) GetTTLParam() int {
	ttl := this.uriParms.GetValue("ttl")
	if ttl != nil {
		return ttl.(int)
	}
	return -1
}

/** Returns the value of the "transport" parameter, or null if this is not
 * Set. This is equivalent to GetParameter("transport").
 *
 * @return the transport paramter of the SipURI
 */
func (this *SipUri) GetTransportParam() string {
	if this.uriParms != nil {
		return this.uriParms.GetValue(core.SIPTransportNames_TRANSPORT).(string)
	}
	return ""
}

/** Returns the value of the <code>userParam</code>,
 *or null if this is not Set.
 * <p>
 * This is equivalent to GetParameter("user").
 *
 * @return the value of the <code>userParam</code> of the SipURI
 */
func (this *SipUri) GetUser() string {
	return this.authority.GetUser()
}

/** Returns true if this SipURI is secure i.e. if this SipURI represents a
 * sips URI. A sip URI returns false.
 *
 * @return  <code>true</code> if this SipURI represents a sips URI, and
 * <code>false</code> if it represents a sip URI.
 */
func (this *SipUri) IsSecure() bool {
	return strings.ToLower(this.GetScheme()) == (core.SIPTransportNames_SIPS)
}

/** This method determines if this is a URI with a scheme of "sip" or "sips".
 *
 * @return true if the scheme is "sip" or "sips", false otherwise.
 */
func (this *SipUri) IsSipURI() bool {
	return true
}

/** Sets the value of the specified header fields to be included in a
 * request constructed from the URI. If the header already had a value it
 * will be overwritten.
 *
 * @param name - a String specifying the header name
 * @param value - a String specifying the header value
 */
func (this *SipUri) SetHeader(name, value string) {
	if this.qheaders.GetValue(name) == nil {
		nv := core.NewNameValue(name, value)
		this.qheaders.AddNameValue(nv)
	} else {
		nv := this.qheaders.GetNameValue(name)
		nv.SetValue(value)
	}
}

/** Returns the host part of this SipURI.
 *
 * @return  the host part of this SipURI
 */
func (this *SipUri) SetHostString(host string) {
	h := core.NewHost(host)
	this.SetHost(h)
}

/** Sets the value of the <code>lr</code> parameter of this SipURI. The lr
 * parameter, when present, indicates that the element responsible for
 * this resource implements the routing mechanisms specified in RFC 3261.
 * This parameter will be used in the URIs proxies place in the
 * Record-Route header field values, and may appear in the URIs in a
 * pre-existing route Set.
 */
func (this *SipUri) SetLrParam() {
	if this.uriParms.GetValue("lr") != nil {
		return
	}
	nv := core.NewNameValue("lr", nil)
	this.uriParms.AddNameValue(nv)
}

/** Sets the value of the <code>maddr</code> parameter of this SipURI. The
 * maddr parameter indicates the server address to be contacted for this
 * user, overriding any address derived from the host field. This is
 * equivalent to SetParameter("maddr", maddr).
 *
 * @param  method - new value of the <code>maddr</code> parameter
 */
func (this *SipUri) SetMAddrParam(maddr string) {
	if maddr == "" {
		//throw new NullPointerException("bad maddr");
		return
	}
	this.SetParameter("maddr", maddr)
}

/** Sets the value of the <code>method</code> parameter. This specifies
 * which SIP method to use in requests directed at this URI. This is
 * equivalent to SetParameter("method", method).
 *
 * @param  method - new value String value of the method parameter
 */
func (this *SipUri) SetMethodParam(method string) {
	this.SetParameter("method", method)
}

/**
 * Sets the value of the specified parameter. If the parameter already had
 *
 * a value it will be overwritten. A zero-length String indicates flag
 *
 * parameter.
 *
 *
 *
 * @param name - a String specifying the parameter name
 *
 * @param value - a String specifying the parameter value
 *
 * @throws ParseException which signals that an error has been reached
 *
 * unexpectedly while parsing the parameter name or value.
 *
 */
func (this *SipUri) SetParameter(name, value string) {
	if name == "ttl" {
		//try {
		if _, err := strconv.Atoi(value); err != nil {
			return
		}
	}
	nv := core.NewNameValue(name, value)
	this.uriParms.Delete(name)
	this.uriParms.AddNameValue(nv)
}

/** Sets the scheme of this URI to sip or sips depending on whether the
 * argument is true or false. The default value is false.
 *
 * @param secure - the boolean value indicating if the SipURI is secure.
 */
func (this *SipUri) SetSecure(secure bool) {
	if secure {
		this.scheme = core.SIPTransportNames_SIPS
	} else {
		this.scheme = core.SIPTransportNames_SIP
	}
}

/** Sets the value of the <code>ttl</code> parameter. The ttl parameter
 * specifies the time-to-live value when packets are sent using UDP
 * multicast. This is equivalent to SetParameter("ttl", ttl).
 *
 * @param ttl - new value of the <code>ttl</code> parameter
 */
func (this *SipUri) SetTTLParam(ttl int) {
	if ttl <= 0 {
		//throw new IllegalArgumentException ("Bad ttl value");
		return
	}
	if this.uriParms != nil {
		this.uriParms.Delete("ttl")
		nv := core.NewNameValue("ttl", ttl)
		this.uriParms.AddNameValue(nv)
	}
}

/** Sets the value of the "transport" parameter. This parameter specifies
 * which transport protocol to use for sending requests and responses to
 * this entity. The following values are defined: "udp", "tcp", "sctp",
 * "tls", but other values may be used also. This method is equivalent to
 * SetParameter("transport", transport). Transport parameter constants
 * are defined in the {@link javax.sip.ListeningPoint}.
 *
 * @param transport - new value for the "transport" parameter
 * @see javax.sip.ListeningPoint
 */
func (this *SipUri) SetTransportParam(transport string) {
	if transport == "" {
		//throw new NullPointerException("null arg");
		return
	}
	if strings.ToUpper(transport) == "UDP" ||
		strings.ToUpper(transport) == "TCP" {
		nv := core.NewNameValue(core.SIPTransportNames_TRANSPORT, strings.ToLower(transport))
		this.uriParms.Delete(core.SIPTransportNames_TRANSPORT)
		this.uriParms.AddNameValue(nv)
	} // else throw new ParseException ("bad transport " + transport, 0);
}

/** Returns the user part of this SipURI, or null if it is not Set.
 *
 * @return  the user part of this SipURI
 */
func (this *SipUri) GetUserParam() string {
	return this.GetParameter("user")

}

/** Returns whether the the <code>lr</code> parameter is Set. This is
 * equivalent to hasParameter("lr"). This interface has no GetLrParam as
 * RFC3261 does not specify any values for the "lr" paramater.
 *
 * @return true if the "lr" parameter is Set, false otherwise.
 */
func (this *SipUri) HasLrParam() bool {
	return this.uriParms.GetNameValue("lr") != nil
}
