package header

import (
	"gosip/address"
)

type AddressParametersHeader interface{
	ParametersHeader
	
	GetAddress() address.Address
	SetAddress(addr address.Address)
}