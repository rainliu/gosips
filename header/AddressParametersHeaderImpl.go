package header

import (
	"gosip/address"
)


/** An abstract class for headers that take an address and parameters. 
*/

type AddressParametersHeaderImpl struct{
	ParametersHeaderImpl

  	addr *address.AddressImpl;
}
	
	/** Constructor given the name of the header.
	*/
	func NewAddressParametersHeaderImpl( name string) *AddressParametersHeaderImpl{
		this := &AddressParametersHeaderImpl{};
		this.ParametersHeaderImpl.super(name);
		return this;
	}
	
	func (this *AddressParametersHeaderImpl) super(name string) {
		this.ParametersHeaderImpl.super(name);
	}
	
        /** get the Address field
         * @return the imbedded  Address
         */        
    func (this *AddressParametersHeaderImpl) GetAddress() *address.AddressImpl {
		return this.addr;
	}
        
        /** set the Address field
         * @param address Address to set
         */        
	func (this *AddressParametersHeaderImpl) SetAddress(addr *address.AddressImpl) {
		this.addr = addr
	}

