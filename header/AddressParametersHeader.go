package header

import (
	"gosip/address"
)

type AddressParametersHeader interface{
	ParametersHeader
	
	GetAddress() *address.AddressImpl
	SetAddress(addr *address.AddressImpl)
}