package address

import (
	"container/list"
	"gosips/core"
)

/** Implementation of the TelURL interface.
*<a href="{@docRoot}/uncopyright.html">This code is in the public domain.</a>
*
 */
type TelURLImpl struct {
	GenericURI

	telephoneNumber *core.TelephoneNumber
}

/** Creates a new instance of TelURLImpl */
func NewTelURLImpl() *TelURLImpl {
	this := &TelURLImpl{}

	this.scheme = "tel"

	return this
}

/** Set the telephone number.
 *@param telephoneNumber -- telephone number to Set.
 */

func (this *TelURLImpl) SetTelephoneNumber(telephoneNumber *core.TelephoneNumber) {
	this.telephoneNumber = telephoneNumber
}

/** Returns the value of the <code>isdnSubAddress</code> parameter, or null
 * if it is not Set.
 *
 * @return  the value of the <code>isdnSubAddress</code> parameter
 */
func (this *TelURLImpl) GetIsdnSubAddress() string {
	return this.telephoneNumber.GetIsdnSubaddress()
}

/** Returns the value of the <code>postDial</code> parameter, or null if it
 * is not Set.
 *
 * @return  the value of the <code>postDial</code> parameter
 */
func (this *TelURLImpl) GetPostDial() string {
	return this.telephoneNumber.GetPostDial()
}

/** Returns the value of the "scheme" of this URI, for example "sip", "sips"
 * or "tel".
 *
 * @return the scheme paramter of the URI
 */
func (this *TelURLImpl) GetScheme() string {
	return this.scheme
}

/** Returns <code>true</code> if this TelURL is global i.e. if the TelURI
 * has a global phone user.
 *
 * @return <code>true</code> if this TelURL represents a global phone user,
 * and <code>false</code> otherwise.
 */
func (this *TelURLImpl) IsGlobal() bool {
	return this.telephoneNumber.IsGlobal()
}

/** This method determines if this is a URI with a scheme of "sip" or "sips".
 *
 * @return true if the scheme is "sip" or "sips", false otherwise.
 */
func (this *TelURLImpl) IsSipURI() bool {
	return false
}

/** Sets phone user of this TelURL to be either global or local. The default
 * value is false, hence the TelURL is defaulted to local.
 *
 * @param global - the boolean value indicating if the TelURL has a global
 * phone user.
 */
func (this *TelURLImpl) SetGlobal(global bool) {
	this.telephoneNumber.SetGlobal(global)
}

/** Sets ISDN subaddress of this TelURL. If a subaddress is present, it is
 * appended to the phone number after ";isub=".
 *
 * @param isdnSubAddress - new value of the <code>isdnSubAddress</code>
 * parameter
 */
func (this *TelURLImpl) SetIsdnSubAddress(isdnSubAddress string) {
	this.telephoneNumber.SetIsdnSubaddress(isdnSubAddress)
}

/** Sets post dial of this TelURL. The post-dial sequence describes what and
 * when the local entity should send to the phone line.
 *
 * @param postDial - new value of the <code>postDial</code> parameter
 */
func (this *TelURLImpl) SetPostDial(postDial string) {
	this.telephoneNumber.SetPostDial(postDial)
}

/** Set the telephone number.
 * @param telphoneNumber -- long phone number to Set.
 */
func (this *TelURLImpl) SetPhoneNumber(telephoneNumber string) {
	this.telephoneNumber.SetPhoneNumber(telephoneNumber)
}

/** Get the telephone number.
 *
 *@return -- the telephone number.
 */
func (this *TelURLImpl) GetPhoneNumber() string {
	return this.telephoneNumber.GetPhoneNumber()
}

/** Return the string encoding.
 *
 *@return -- the string encoding.
 */
/*func (this *TelURLImpl) toString() {
    return this.scheme + ":" + telephoneNumber.encode();
}*/

func (this *TelURLImpl) String() string {
	return this.scheme + ":" + this.telephoneNumber.String()
}

/** Deep copy clone operation.
 *
 *@return -- a cloned version of this telephone number.
 */
func (this *TelURLImpl) Clone() interface{} {
	retval := NewTelURLImpl()
	retval.scheme = this.scheme
	if this.telephoneNumber != nil {
		retval.telephoneNumber = this.telephoneNumber.Clone().(*core.TelephoneNumber)
	}
	return retval
}

func (this *TelURLImpl) GetParameter(parameterName string) string {
	return this.telephoneNumber.GetParameter(parameterName)
}

func (this *TelURLImpl) SetParameter(name, value string) {
	this.telephoneNumber.SetParameter(name, value)
}

func (this *TelURLImpl) GetParameterNames() *list.List {
	return this.telephoneNumber.GetParameterNames()
}

func (this *TelURLImpl) RemoveParameter(name string) {
	this.telephoneNumber.RemoveParameter(name)
}
