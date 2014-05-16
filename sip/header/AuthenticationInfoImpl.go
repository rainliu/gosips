package header

import (
	"errors"
	"gosips/core"
	"strconv"
	"strings"
)

/** Authentication info SIP Header.
 *
 */
type AuthenticationInfo struct {
	Parameters
}

/** Default contstructor.
 */
func NewAuthenticationInfo() *AuthenticationInfo {
	this := &AuthenticationInfo{}
	this.Parameters.super(core.SIPHeaderNames_AUTHENTICATION_INFO)
	this.parameters.SetSeparator(core.SIPSeparatorNames_COMMA)
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
 */
func (this *AuthenticationInfo) GetAuthenticationInfo() string {
	return this.EncodeBody()
}

/** Returns the CNonce value of this AuthenticationInfoHeader.
 *
 * @return the String representing the cNonce information, null if value is
 * not set.
 *
 */
func (this *AuthenticationInfo) GetCNonce() string {
	return this.GetParameter(ParameterNames_CNONCE)
}

/** Returns the nextNonce value of this AuthenticationInfoHeader.
 *
 * @return the String representing the nextNonce
 * information, null if value is not set.
 *
 */
func (this *AuthenticationInfo) GetNextNonce() string {
	return this.GetParameter(ParameterNames_NEXT_NONCE)
}

/** Returns the Nonce Count value of this AuthenticationInfoHeader.
 *
 * @return the integer representing the nonceCount information, -1 if value is
 * not set.
 *
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
 *
 */
func (this *AuthenticationInfo) GetQop() string {
	return this.GetParameter(ParameterNames_QOP)
}

/** Returns the Response value of this AuthenticationInfoHeader.
 *
 * @return the String representing the Response information.
 *
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
 *
 */
func (this *AuthenticationInfo) SetCNonce(cNonce string) (ParseException error) {
	this.SetParameter(ParameterNames_CNONCE, cNonce)
	return nil
}

/** Sets the NextNonce of the AuthenticationInfoHeader to the <var>nextNonce</var>
 * parameter value.
 *
 * @param nextNonce - the new nextNonce String of this AuthenticationInfoHeader.
 * @throws ParseException which signals that an error has been reached
 * unexpectedly while parsing the nextNonce value.
 *
 */
func (this *AuthenticationInfo) SetNextNonce(nextNonce string) (ParseException error) {
	this.SetParameter(ParameterNames_NEXT_NONCE, nextNonce)
	return nil
}

/** Sets the Nonce Count of the AuthenticationInfoHeader to the <var>nonceCount</var>
 * parameter value.
 *
 * @param nonceCount - the new nonceCount integer of this AuthenticationInfoHeader.
 * @throws ParseException which signals that an error has been reached
 * unexpectedly while parsing the nonceCount value.
 *
 */
func (this *AuthenticationInfo) SetNonceCount(nonceCount int) (ParseException error) {
	if nonceCount < 0 {
		return errors.New("ParseException: bad value")
	}
	nc := strconv.FormatUint(uint64(nonceCount), 16)

	base := "00000000"
	nc = base[0:8-len(nc)] + nc
	this.SetParameter(ParameterNames_NC, nc)
	return nil
}

/** Sets the Qop value of the AuthenticationInfoHeader to the new
 * <var>qop</var> parameter value.
 *
 * @param qop - the new Qop string of this AuthenticationInfoHeader.
 * @throws ParseException which signals that an error has been reached
 * unexpectedly while parsing the Qop value.
 *
 */
func (this *AuthenticationInfo) SetQop(qop string) (ParseException error) {
	this.SetParameter(ParameterNames_QOP, qop)
	return nil
}

/** Sets the Response of the
 * AuthenticationInfoHeader to the new <var>response</var>
 * parameter value.
 *
 * @param response - the new response String of this
 * AuthenticationInfoHeader.
 * @throws ParseException which signals that an error has been reached
 * unexpectedly while parsing the Response.
 *
 */
func (this *AuthenticationInfo) SetResponse(response string) (ParseException error) {
	this.SetParameter(ParameterNames_RESPONSE, response)
	return nil
}

func (this *AuthenticationInfo) SetParameter(name, value string) error {
	if name == "" {
		return errors.New("NullPointerException: null name")
	}

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
			strings.ToLower(name) == (ParameterNames_ALGORITHM) ||
			strings.ToLower(name) == (ParameterNames_RESPONSE_AUTH) {
			//if value == "" {//TODO by LY
			//	return errors.New("NullPointerException: null value")
			//}
			if strings.HasPrefix(value, core.SIPSeparatorNames_DOUBLE_QUOTE) {
				return errors.New("ParseException: " +
					value + " : Unexpected DOUBLE_QUOTE")
			}
			nv.SetQuotedValue()
		}
		this.parameters.SetNameValue(nv)
	} else {
		nv.SetValue(value)
	}

	return nil
}
