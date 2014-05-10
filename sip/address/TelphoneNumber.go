package address

import (
	"container/list"
	"gosips/core"
)

/** Telephone number class.
 */
type TelephoneNumber struct {
	//NetObject

	/** isglobal field
	 */
	isglobal bool

	/** phoneNumber field
	 */
	phoneNumber string

	/** parmeters list
	 */
	parms *core.NameValueList
}

/** Creates new TelephoneNumber */
func NewTelephoneNumber() *TelephoneNumber {
	this := &TelephoneNumber{}

	this.parms = core.NewNameValueList("telparms")

	return this
}

/** delete the specified parameter.
 * @param name String to set
 */
func (this *TelephoneNumber) DeleteParm(name string) {
	this.parms.Delete(name)
}

/** get the PhoneNumber field
 * @return String
 */
func (this *TelephoneNumber) GetPhoneNumber() string {
	return this.phoneNumber
}

/** get the PostDial field
 * @return String
 */
func (this *TelephoneNumber) GetPostDial() string {
	return this.parms.GetValue(core.SIPTransportNames_POSTDIAL).(string)
}

/**
 * Get the isdn subaddress for this number.
 * @return String
 */
func (this *TelephoneNumber) GetIsdnSubaddress() string {
	return this.parms.GetValue(core.SIPTransportNames_ISUB).(string)
}

/** returns true if th PostDial field exists
 * @return boolean
 */
func (this *TelephoneNumber) HasPostDial() bool {
	return this.parms.GetValue(core.SIPTransportNames_POSTDIAL) != nil
}

/** return true if this header has parameters.
 * @param pname String to set
 * @return boolean
 */
func (this *TelephoneNumber) HasParm(pname string) bool {
	return this.parms.HasNameValue(pname)
}

/**
 * return true if the isdn subaddress exists.
 * @return boolean
 */
func (this *TelephoneNumber) HasIsdnSubaddress() bool {
	return this.HasParm(core.SIPTransportNames_ISUB)
}

/**
 * is a global telephone number.
 * @return boolean
 */
func (this *TelephoneNumber) IsGlobal() bool {
	return this.isglobal
}

/** remove the PostDial field
 */
func (this *TelephoneNumber) RemovePostDial() {
	this.parms.Delete(core.SIPTransportNames_POSTDIAL)
}

/**
 * Remove the isdn subaddress (if it exists).
 */
func (this *TelephoneNumber) RemoveIsdnSubaddress() {
	this.DeleteParm(core.SIPTransportNames_ISUB)
}

/**
 * Set the list of parameters.
 * @param p NameValueList to set
 */
func (this *TelephoneNumber) SetParameters(p *core.NameValueList) {
	this.parms = p
}

/** set the Global field
 * @param g boolean to set
 */
func (this *TelephoneNumber) SetGlobal(g bool) {
	this.isglobal = g
}

/** set the PostDial field
 * @param p String to set
 */
func (this *TelephoneNumber) SetPostDial(p string) {
	nv := core.NewNameValue(core.SIPTransportNames_POSTDIAL, p)
	this.parms.AddNameValue(nv)
}

/** set the specified parameter
 * @param name String to set
 * @param value Object to set
 */
func (this *TelephoneNumber) SetParm(name string, value interface{}) {
	nv := core.NewNameValue(name, value)
	this.parms.AddNameValue(nv)
}

/**
 * set the isdn subaddress for this structure.
 * @param isub String to set
 */
func (this *TelephoneNumber) SetIsdnSubaddress(isub string) {
	this.SetParm(core.SIPTransportNames_ISUB, isub)
}

/** set the PhoneNumber field
 * @param num String to set
 */
func (this *TelephoneNumber) SetPhoneNumber(num string) {
	this.phoneNumber = num
}

func (this *TelephoneNumber) Clone() interface{} {
	retval := &TelephoneNumber{}

	retval.isglobal = this.isglobal
	retval.phoneNumber = this.phoneNumber
	retval.parms = core.NewNameValueList("telparms")

	for e := this.parms.Front(); e != nil; e = e.Next() {
		nv := e.Value.(*core.NameValue)
		retval.parms.AddNameValue(nv.Clone().(*core.NameValue))
	}

	return retval
}

func (this *TelephoneNumber) String() string {
	var retval string //= "";
	if this.isglobal {
		retval += "+"
	}
	retval += this.phoneNumber
	if this.parms.Len() != 0 {
		retval += core.SIPSeparatorNames_SEMICOLON
		retval += this.parms.String()
	}
	return retval
}

/**
 * Returns the value of the named parameter, or null if it is not set. A
 * zero-length String indicates flag parameter.
 *
 * @param <var>name</var> name of parameter to retrieve
 *
 * @return the value of specified parameter
 *
 */
func (this *TelephoneNumber) GetParameter(name string) string {
	val := this.parms.GetValue(name)
	if val == nil {
		return ""
	}
	
	return val.(string)
}

/**
 *
 * Returns an Iterator over the names (Strings) of all parameters.
 *
 * @return an Iterator over all the parameter names
 *
 */
func (this *TelephoneNumber) GetParameterNames() *list.List {
	return this.parms.GetNames()
}

func (this *TelephoneNumber) RemoveParameter(parameter string) {
	this.parms.Delete(parameter)
}

func (this *TelephoneNumber) SetParameter(name, value string) {
	nv := core.NewNameValue(name, value)
	this.parms.AddNameValue(nv)
}
