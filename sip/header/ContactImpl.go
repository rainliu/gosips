package header

import (
	"bytes"
	"errors"
	"gosips/core"
	"gosips/sip/address"
	"strconv"
)

/**
* Contact Item.
 */
type Contact struct { //implements javax.sip.header.ContactHeader
	AddressParameters

	// This must be private or the toString will go for a loop!
	contactList *ContactList

	/** wildCardFlag field.
	 */
	wildCardFlag bool
}

/** Default constructor.
 */
func NewContact() *Contact {
	this := &Contact{}

	this.AddressParameters.super(core.SIPHeaderNames_CONTACT)
	this.contactList = NewContactList()

	return this
}

func (this *Contact) super(hname string) {
	this.AddressParameters.super(hname)
	this.contactList = NewContactList()
}

/** Set a parameter.
 */
/*
func (this *Contact) SetParameter(name, value string) (ParseException error) {
	nv := this.parameters.GetNameValue(name)
	if nv != nil {
		nv.SetValue(value)
	} else {
		nv = core.NewNameValue(name, value)
		if strings.ToLower(name) == "methods" {
			nv.SetQuotedValue()
		}
		this.parameters.AddNameValue(nv)
	}
	return nil
}*/

/**
 * Encode body of the header into a cannonical String.
 * @return string encoding of the header value.
 */
func (this *Contact) EncodeBody() string {
	var encoding bytes.Buffer //= new StringBuffer();
	if this.wildCardFlag {
		encoding.WriteString("*")
		return encoding.String()
	}
	addr, _ := this.addr.(*address.AddressImpl)
	// Bug report by Joao Paulo
	if addr.GetAddressType() == address.NAME_ADDR {
		encoding.WriteString(this.addr.String())
	} else {
		// Encoding in canonical form must have <> around address.
		encoding.WriteString("<")
		encoding.WriteString(this.addr.String())
		encoding.WriteString(">")
	}
	if this.parameters.Len() > 0 {
		encoding.WriteString(core.SIPSeparatorNames_SEMICOLON)
		encoding.WriteString(this.parameters.String())
	}

	return encoding.String()
}

/** get the Contact list.
 * @return ContactList
 */
func (this *Contact) GetContactList() *ContactList {
	return this.contactList
}

/** get the WildCardFlag field
 * @return boolean
 */
func (this *Contact) GetWildCardFlag() bool {
	return this.wildCardFlag
}

/** get the address field.
 * @return Address
 */
func (this *Contact) GetAddress() address.Address {
	// JAIN-SIP stores the wild card as an address!
	return this.addr
}

/** get the parameters List
 * @return NameValueList
 */
func (this *Contact) GetContactParms() *core.NameValueList {
	return this.parameters
}

/** get Expires parameter.
 * @return the Expires parameter.
 */
func (this *Contact) GetExpires() int {
	retval, _ := strconv.Atoi(this.GetParameter(core.SIPHeaderNames_EXPIRES))
	return retval
}

/** Set the expiry time in seconds.
*@param expiryDeltaSeconds exipry time.
 */

func (this *Contact) SetExpires(expiryDeltaSeconds int) (InvalidArgumentException error) {
	//Integer deltaSeconds = new Integer(expiryDeltaSeconds);
	this.parameters.AddNameValue(core.NewNameValue(core.SIPHeaderNames_EXPIRES, strconv.Itoa(expiryDeltaSeconds)))
	return nil
}


/** set the Contact List
 * @param cl ContactList to set
 */
func (this *Contact) SetContactList(cl *ContactList) {
	this.contactList = cl
}

/**
 * Set the wildCardFlag member
 * @param w boolean to set
 */
func (this *Contact) SetWildCardFlag(w bool) {
	this.wildCardFlag = true
	addr := address.NewAddressImpl()
	addr.SetWildCardFlag()
	this.SetAddress(addr)
}

/**
 * Set the address member
 *
 * @param address Address to set
 */
func (this *Contact) SetAddress(addr address.Address) {
	// Canonical form must have <> around the address.
	this.AddressParameters.SetAddress(addr)
	this.wildCardFlag = false
}



/** get the QValue field. Return -1 if the parameter has not been
 * set.
 * @return float
 */
func (this *Contact) GetQValue() float32 {
	if !this.HasParameter(ParameterNames_Q) {
		return -1
	}
	qstr := this.GetParameterValue(ParameterNames_Q)
	q, _ := strconv.ParseFloat(qstr, 32)
	return float32(q)
}

/**
 * Return true if the q value has been set.
 * @return boolean
 */
func (this *Contact) HasQValue() bool {
	return this.HasParameter(ParameterNames_Q)
}

/**
 * Remove the q value.
 */
func (this *Contact) RemoveQValue() {
	this.RemoveParameter(ParameterNames_Q)
}

/**
 * Sets q-value for media-range. Q-values allow the
 *
 * user to indicate the relative degree of preference for that media-range,
 *
 * using the qvalue scale from 0 to 1. If no q-value is present, the
 *
 * media-range should be treated as having a q-value of 1.
 *
 *
 *
 * @param qValue - the new float value of the q-value
 *
 * @throws InvalidArgumentException if the q parameter value is not between <code>0 and 1</code>.
 *
 */
func (this *Contact) SetQValue(q float32) (InvalidArgumentException error) {
	if q < 0.0 || q > 1.0 {
		return errors.New("qvalue out of range!")
	}
	this.SetParameter(ParameterNames_Q, strconv.FormatFloat(float64(q), 'f', -1, 32))
	return nil
}
