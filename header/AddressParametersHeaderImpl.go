package header

import (
	"gosip/address"
)


/** An abstract class for headers that take an address and parameters. 
*/

type AddressParametersHeader struct{
	ParametersHeader

  	addr *address.AddressImpl;
}
	
	/** Constructor given the name of the header.
	*/
	func NewAddressParametersHeader( name string) *AddressParametersHeader{
		this := &AddressParametersHeader{};
		this.ParametersHeader.super(name);
		return this;
	}
	
	func (this *AddressParametersHeader) super(name string) {
		this.ParametersHeader.super(name);
	}
	
        /** get the Address field
         * @return the imbedded  Address
         */        
    func (this *AddressParametersHeader) GetAddress() *address.AddressImpl {
		return this.addr;
	}
        
        /** set the Address field
         * @param address Address to set
         */        
	func (this *AddressParametersHeader) SetAddress(addr *address.AddressImpl) {
		this.addr = addr
	}

