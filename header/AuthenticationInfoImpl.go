package header

import (
	"gosip/core"
	"strconv"
	"strings"
)

/** Authentication info SIP Header.
 *
 *@author M. Ranganathan <mranga@nist.gov>  NIST/ITL/ANTD
 *
 *<a href="{@docRoot}/uncopyright.html">This code is in the public domain.</a>
 *
 *@version JAIN-1.1
 *
 */
type AuthenticationInfo struct {
	Parameters
	//implements javax.sip.header.AuthenticationInfoHeader {
}

/** Default contstructor.
 */

func NewAuthenticationInfo() *AuthenticationInfo {
	this := &AuthenticationInfo{}
	this.Parameters.super(core.SIPHeaderNames_AUTHENTICATION_INFO)
	this.parameters.SetSeparator(core.SIPSeparatorNames_COMMA) // Odd ball.
	return this
}

func (this *AuthenticationInfo) Add(nv *core.NameValue) {
	this.parameters.AddNameValue(nv)
}

/** Value of header encoded in canonical form.
 */
func (this *AuthenticationInfo) String() string {
	return this.headerName + core.SIPSeparatorNames_COLON +
		core.SIPSeparatorNames_SP + this.EncodeBody() + core.SIPSeparatorNames_NEWLINE
}

func (this *AuthenticationInfo) EncodeBody() string {
	return this.parameters.String()
}

/** Get the name value pair for a given authentication info parameter.
 *
 *@param name is the name for which we want to retrieve the name value
 *  list.
 */

func (this *AuthenticationInfo) GetAuthInfo(name string) *core.NameValue {
	return this.parameters.GetNameValue(name)
}

/**
 * Returns the AuthenticationInfo value of this AuthenticationInfoHeader.
 *
 *
 *
 * @return the String representing the AuthenticationInfo
 *
 * @since JAIN SIP v1.1
 *
 */
func (this *AuthenticationInfo) GetAuthenticationInfo() string {
	return this.EncodeBody()
}

/** Returns the CNonce value of this AuthenticationInfoHeader.
 *
 * @return the String representing the cNonce information, null if value is
 * not set.
 * @since v1.1
 */
func (this *AuthenticationInfo) GetCNonce() string {
	return this.GetParameter(ParameterNames_CNONCE)
}

/** Returns the nextNonce value of this AuthenticationInfoHeader.
 *
 * @return the String representing the nextNonce
 * information, null if value is not set.
 * @since v1.1
 */
func (this *AuthenticationInfo) GetNextNonce() string {
	return this.GetParameter(ParameterNames_NEXT_NONCE)
}

/** Returns the Nonce Count value of this AuthenticationInfoHeader.
 *
 * @return the integer representing the nonceCount information, -1 if value is
 * not set.
 * @since v1.1
 */
func (this *AuthenticationInfo) GetNonceCount() int {
	s := this.GetParameter(ParameterNames_NONCE_COUNT)
	nCount, _ := strconv.ParseInt(s, 10, 32)
	return int(nCount)
}

/** Returns the messageQop value of this AuthenticationInfoHeader.
 *
 * @return the string representing the messageQop information, null if the
 * value is not set.
 * @since v1.1
 */
func (this *AuthenticationInfo) GetQop() string {
	return this.GetParameter(ParameterNames_QOP)
}

/** Returns the Response value of this AuthenticationInfoHeader.
 *
 * @return the String representing the Response information.
 * @since v1.1
 */
func (this *AuthenticationInfo) GetResponse() string {
	return this.GetParameter(ParameterNames_RESPONSE_AUTH)
}

/** Sets the CNonce of the AuthenticationInfoHeader to the <var>cNonce</var>
 * parameter value.
 *
 * @param cNonce - the new cNonce String of this AuthenticationInfoHeader.
 * @throws ParseException which signals that an error has been reached
 * unexpectedly while parsing the cNonce value.
 * @since v1.1
 */
func (this *AuthenticationInfo) SetCNonce(cNonce string) { //throws ParseException {
	this.SetParameter(ParameterNames_CNONCE, cNonce)
}

/** Sets the NextNonce of the AuthenticationInfoHeader to the <var>nextNonce</var>
 * parameter value.
 *
 * @param nextNonce - the new nextNonce String of this AuthenticationInfoHeader.
 * @throws ParseException which signals that an error has been reached
 * unexpectedly while parsing the nextNonce value.
 * @since v1.1
 */
func (this *AuthenticationInfo) SetNextNonce(nextNonce string) { //throws ParseException {
	this.SetParameter(ParameterNames_NEXT_NONCE, nextNonce)
}

/** Sets the Nonce Count of the AuthenticationInfoHeader to the <var>nonceCount</var>
 * parameter value.
 *
 * @param nonceCount - the new nonceCount integer of this AuthenticationInfoHeader.
 * @throws ParseException which signals that an error has been reached
 * unexpectedly while parsing the nonceCount value.
 * @since v1.1
 */
func (this *AuthenticationInfo) SetNonceCount(nonceCount int) { //throws ParseException {
	//if (nonceCount  < 0 ) throw new ParseException("bad value",0);
	nc := strconv.FormatUint(uint64(nonceCount), 16) //.toHexString(nonceCount)

	base := "00000000"
	nc = base[0:8-len(nc)] + nc
	this.SetParameter(ParameterNames_NC, nc)
}

/** Sets the Qop value of the AuthenticationInfoHeader to the new
 * <var>qop</var> parameter value.
 *
 * @param qop - the new Qop string of this AuthenticationInfoHeader.
 * @throws ParseException which signals that an error has been reached
 * unexpectedly while parsing the Qop value.
 * @since v1.1
 */
func (this *AuthenticationInfo) SetQop(qop string) { // throws ParseException {
	this.SetParameter(ParameterNames_QOP, qop)
}

/** Sets the Response of the
 * AuthenticationInfoHeader to the new <var>response</var>
 * parameter value.
 *
 * @param response - the new response String of this
 * AuthenticationInfoHeader.
 * @throws ParseException which signals that an error has been reached
 * unexpectedly while parsing the Response.
 * @since v1.1
 */
func (this *AuthenticationInfo) SetResponse(response string) { //throws ParseException {
	this.SetParameter(ParameterNames_RESPONSE, response)
}

func (this *AuthenticationInfo) SetParameter(name, value string) error {
	//throws ParseException {
	//if (name == null) throw new NullPointerException("null name");
	nv := this.parameters.GetNameValue(strings.ToLower(name))
	if nv == nil {
		nv = core.NewNameValue(name, value)
		if strings.ToLower(name) == (ParameterNames_QOP) ||
			strings.ToLower(name) == (ParameterNames_NEXT_NONCE) ||
			strings.ToLower(name) == (ParameterNames_REALM) ||
			strings.ToLower(name) == (ParameterNames_CNONCE) ||
			strings.ToLower(name) == (ParameterNames_NONCE) ||
			strings.ToLower(name) == (ParameterNames_OPAQUE) ||
			strings.ToLower(name) == (ParameterNames_USERNAME) ||
			strings.ToLower(name) == (ParameterNames_DOMAIN) ||
			strings.ToLower(name) == (ParameterNames_NEXT_NONCE) ||
			strings.ToLower(name) == (ParameterNames_RESPONSE_AUTH) {
			// if (value ==
			//     throw new NullPointerException("null value");
			// if (value.startsWith(Separators.DOUBLE_QUOTE))
			//     throw new ParseException
			//     (value + " : Unexpected DOUBLE_QUOTE",0);
			nv.SetQuotedValue()
		}
		this.parameters.SetNameValue(nv)
	} else {
		nv.SetValue(value)
	}

	return nil
}
