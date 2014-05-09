package address

import (
	"bytes"
	"errors"
	"gosips/core"
)

/**
 * Address structure. Imbeds a URI and adds a display name.
 *
 *@author M. Ranganathan <mranga@nist.gov>  <br/>
 *
 *<a href="{@docRoot}/uncopyright.html">This code is in the public domain.</a>
 *
 *@version JAIN-SIP-1.1
 *
 */
/** Constant field.
 */
const (
	NAME_ADDR = iota //0;

	/** constant field.
	 */
	ADDRESS_SPEC //1;

	/** Constant field.
	 */
	WILD_CARD //2;
)

type AddressImpl struct { //implements gosip/address/Address

	addressType int

	/** displayName field
	 */
	displayName string

	/** address field
	 */
	address URI //*GenericURI;

}

func NewAddressImpl() *AddressImpl {
	return &AddressImpl{}
}

/** Match on the address only.
 * Dont care about the display name.
 */

//func (this *AddressImpl) Match(other Address) bool  {
// TODO -- add the matcher;
//if other == nil {
//	return true;
//}
//if (!(other instanceof Address)) return false;
//else {
/*
   AddressImpl that = (AddressImpl)other;
   if (that.getMatcher() != null)
       return that.getMatcher().match(this.encode());
   else if (that.displayName != null &&
   this.displayName == null)  return false;
   else if (that.displayName == null)
       return  address.match(that.address);
   else return displayName.equalsIgnoreCase
   (that.displayName) &&
   address.match(that.address);
*/
//}
//    return false;
//}

/** Get the host port portion of the address spec.
 *@return host:port in a HostPort structure.
 */
func (this *AddressImpl) GetHostPort() (hp *core.HostPort, RuntimeException error) {
	if sipuri, ok := this.address.(*SipUri); ok {
		return sipuri.GetHostPort(), nil
	}

	return nil, errors.New("RuntimeException: address is not a SipUri")
}

/** Get the port from the imbedded URI. This assumes that a SIP URL
 * is encapsulated in this address object.
 *
 *@return the port from the address.
 *
 */
func (this *AddressImpl) GetPort() (int, error) {
	if sipuri, ok := this.address.(*SipUri); ok {
		return sipuri.GetHostPort().GetPort(), nil
	}
	return -1, errors.New("RuntimeException: address is not a SipUri")
}

/** Get the user@host:port for the address field. This assumes
 * that the encapsulated object is a SipUri.
 *
 * BUG Fix from Antonis Kadris.
 *
 *@return string containing user@host:port.
 */
func (this *AddressImpl) GetUserAtHostPort() string {
	if sipuri, ok := this.address.(*SipUri); ok {
		sipuri.GetUserAtHostPort()
	}
	return this.address.String()
}

/** Get the host name from the address.
 *
 *@return the host name.
 */
func (this *AddressImpl) GetHost() (string, error) {
	if sipuri, ok := this.address.(*SipUri); ok {
		return sipuri.GetHostPort().GetHost().GetHostName(), nil
	}

	return "", errors.New("RuntimeException: address is not a SipUri")
}

/** Remove a parameter from the address.
 *
 *@param parameterName is the name of the parameter to remove.
 */
func (this *AddressImpl) RemoveParameter(parameterName string) {
	if sipuri, ok := this.address.(*SipUri); ok {
		sipuri.RemoveParameter(parameterName)
	}
	/*
	   if (! (address instanceof SipUri) )
	       throw new RuntimeException
	       ("address is not a SipUri");
	   SipUri uri = (SipUri) address;
	*/
}

/**
 * Encode the address as a string and return it.
 * @return String canonical encoded version of this address.
 */
func (this *AddressImpl) String() string {
	if this.addressType == WILD_CARD {
		return "*"
	}

	var encoding bytes.Buffer //= new StringBuffer();
	if this.displayName != "" {
		encoding.WriteString(core.SIPSeparatorNames_DOUBLE_QUOTE)
		encoding.WriteString(this.displayName)
		encoding.WriteString(core.SIPSeparatorNames_DOUBLE_QUOTE)
		encoding.WriteString(core.SIPSeparatorNames_SP)
	}
	if this.address != nil {
		if this.addressType == NAME_ADDR || this.displayName != "" {
			encoding.WriteString(core.SIPSeparatorNames_LESS_THAN)
		}
		encoding.WriteString(this.address.String())
		if this.addressType == NAME_ADDR || this.displayName != "" {
			encoding.WriteString(core.SIPSeparatorNames_GREATER_THAN)
		}
	}
	return encoding.String()
}

//public AddressImpl() { this.addressType = NAME_ADDR; } init by golang

/**
 * Get the address type;
 * @return int
 */
func (this *AddressImpl) GetAddressType() int {
	return this.addressType
}

/**
 * Set the address type. The address can be NAME_ADDR, ADDR_SPEC or
 * WILD_CARD
 *
 * @param atype int to set
 *
 */
func (this *AddressImpl) SetAddressType(atype int) {
	this.addressType = atype
}

/**
 * get the display name
 *
 * @return String
 *
 */
func (this *AddressImpl) GetDisplayName() string {
	return this.displayName
}

/**
 * Set the displayName member
 *
 * @param displayName String to set
 *
 */
func (this *AddressImpl) SetDisplayName(displayName string) (ParseException error) {
	this.displayName = displayName
	this.addressType = NAME_ADDR
	return nil
}

/**
 * Set the address field
 *
 * @param address SipUri to set
 *
 */
func (this *AddressImpl) SetAddess(address URI) {
	this.address = address
}

/**
 * Compare two address specs for equality.
 *
 * @param other Object to compare this this address
 *
 * @return boolean
 *
 */
/*public boolean equals(Object other) {

    if (! this.getClass().equals(other.getClass())) {
        return false;
    }
    AddressImpl that = (AddressImpl) other;
    if (this.addressType == WILD_CARD &&
    that.addressType != WILD_CARD) return false;

    // Ignore the display name; only compare the address spec.
    boolean retval =  this.address.equals(that.address);
    return retval;
}*/

/** return true if DisplayName exist.
 *
 * @return boolean
 */
func (this *AddressImpl) HasDisplayName() bool {
	return (this.displayName != "")
}

/** remove the displayName field
 */
func (this *AddressImpl) RemoveDisplayName() {
	this.displayName = ""
}

/** Return true if the imbedded URI is a sip URI.
 *
 * @return true if the imbedded URI is a SIP URI.
 *
 */
func (this *AddressImpl) IsSIPAddress() bool {
	_, ok := this.address.(SipURI)
	return ok
}

/** Returns the URI address of this Address. The type of URI can be
 * determined by the scheme.
 *
 * @return address parmater of the Address object
 */
func (this *AddressImpl) GetURI() URI {
	return this.address
}

/** This determines if this address is a wildcard address. That is
 * <code>Address.getAddress.getUserInfo() == *;</code>
 *
 * @return true if this name address is a wildcard, false otherwise.
 */
func (this *AddressImpl) IsWildcard() bool {
	return this.addressType == WILD_CARD
}

/** Sets the URI address of this Address. The URI can be either a
 * TelURL or a SipURI.
 *
 * @param address - the new URI address value of this NameAddress.
 */
func (this *AddressImpl) SetURI(address URI) {
	this.address = address
}

/** Set the user name for the imbedded URI.
 *
 *@param user -- user name to set for the imbedded URI.
 */
func (this *AddressImpl) SetUser(user string) {
	if sipuri, ok := this.address.(*SipUri); ok {
		sipuri.SetUser(user)
	}
}

/** Mark this a wild card address type.
 */
func (this *AddressImpl) SetWildCardFlag() {
	this.addressType = WILD_CARD
}
