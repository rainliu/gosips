package header

import "gosips/core"

/**
* Challenge part of the Auth header. This is only used by the parser interface
 */
type Challenge struct {
	/** scheme field
	 */
	scheme string

	/** authParms list
	 */
	authParams *core.NameValueList
}

/** Default constructor
 */
func NewChallenge() *Challenge {
	this := &Challenge{}
	this.authParams = core.NewNameValueList("authParams")
	this.authParams.SetSeparator(core.SIPSeparatorNames_COMMA)
	return this
}

/**
 * Encode the challenge in canonical form.
 * @return String
 */
func (this *Challenge) String() string {
	return this.scheme + core.SIPSeparatorNames_SP + this.authParams.String()
}

/** Get the scheme field
 * @return String
 */
func (this *Challenge) GetScheme() string {
	return this.scheme
}

/** Get AuthParms list.
 * @return NameValueList
 */
func (this *Challenge) GetAuthParams() *core.NameValueList {
	return this.authParams
}

/** Get the domain
 * @return String
 */
func (this *Challenge) GetDomain() string {
	return this.authParams.GetValue(ParameterNames_DOMAIN).(string)
}

/** Get the URI field
 * @return String
 */
func (this *Challenge) GetURI() string {
	return this.authParams.GetValue(ParameterNames_URI).(string)
}

/** Get the Opaque field
 * @return String
 */
func (this *Challenge) GetOpaque() string {
	return this.authParams.GetValue(ParameterNames_OPAQUE).(string)
}

/** Get QOP value
 * @return String
 */
func (this *Challenge) GetQOP() string {
	return this.authParams.GetValue(ParameterNames_QOP).(string)
}

/** Get the Algorithm value.
 * @return String
 */
func (this *Challenge) GetAlgorithm() string {
	return this.authParams.GetValue(ParameterNames_ALGORITHM).(string)
}

/** Get the State value.
 * @return String
 */
func (this *Challenge) GetStale() string {
	return this.authParams.GetValue(ParameterNames_STALE).(string)
}

/** Get the Signature value.
 * @return String
 */
func (this *Challenge) GetSignature() string {
	return this.authParams.GetValue(ParameterNames_SIGNATURE).(string)
}

/** Get the signedBy value.
 * @return String
 */
func (this *Challenge) GetSignedBy() string {
	return this.authParams.GetValue(ParameterNames_SIGNED_BY).(string)
}

/** Get the Response value.
 * @return String
 */
func (this *Challenge) GetResponse() string {
	return this.authParams.GetValue(ParameterNames_RESPONSE).(string)
}

/** Get the realm value.
 * @return String.
 */
func (this *Challenge) GetRealm() string {
	return this.authParams.GetValue(ParameterNames_REALM).(string)
}

/** Get the specified parameter
 * @param name String to Set
 * @return String to Set
 */
func (this *Challenge) GetParameter(name string) string {
	return this.authParams.GetValue(name).(string)
}

/** boolean function
 * @param name String to Set
 * @return true if this header has the specified parameter, false otherwise.
 */
func (this *Challenge) HasParameter(name string) bool {
	return this.authParams.GetNameValue(name) != nil
}

/** Boolean function
 * @return true if this header has some parameters.
 */
func (this *Challenge) HasParameters() bool {
	return this.authParams.Len() != 0
}

/** delete the specified parameter
 * @param name String
 * @return true if the specified parameter has been removed, false
 * otherwise.
 */
func (this *Challenge) RemoveParameter(name string) bool {
	return this.authParams.Delete(name)
}

/** remove all parameters
 */
func (this *Challenge) RemoveParameters() {
	this.authParams = core.NewNameValueList("authParams")
}

/** Set the specified parameter
 * @param nv NameValue to Set
 */
func (this *Challenge) SetParameter(nv *core.NameValue) {
	this.authParams.AddNameValue(nv)
}

/**
 * Set the scheme member
 * @param s String to Set
 */
func (this *Challenge) SetScheme(s string) {
	this.scheme = s
}

/**
 * Set the authParams member
 * @param a NameValueList to Set
 */
func (this *Challenge) SetAuthParams(a *core.NameValueList) {
	this.authParams = a
}
