package header

import (
	"gosips/sip/address"
)

/** An abstract class for headers that take an address and parameters.
 */

type AddressParameters struct {
	Parameters

	addr address.Address
}

/** Constructor given the name of the header.
 */
func NewAddressParameters(name string) *AddressParameters {
	this := &AddressParameters{}
	this.Parameters.super(name)
	return this
}

func (this *AddressParameters) super(name string) {
	this.Parameters.super(name)
}

/** get the Address field
 * @return the imbedded  Address
 */
func (this *AddressParameters) GetAddress() address.Address {
	return this.addr
}

/** set the Address field
 * @param address Address to set
 */
func (this *AddressParameters) SetAddress(addr address.Address) {
	this.addr = addr
}
