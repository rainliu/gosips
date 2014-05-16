package header

import (
	"gosips/sip/address"
)

/**
 * This interface represents methods for manipulating Address object
 * values for any header that contains a Address value.
 * <p>
 * When the header field value contains a display name encapsulated in the
 * Address, the URI including all URI parameters is enclosed in "<" and ">".
 * If no "<" and ">" are present, all parameters after the URI are header
 * parameters, not URI parameters. The display name can be tokens, or a
 * quoted string, if a larger character set is desired.
 *
 * @see Address
 * @see ContactHeader
 * @see FromHeader
 * @see RecordRouteHeader
 * @see ReplyToHeader
 * @see RouteHeader
 * @see ToHeader
 */
type AddressHeader interface {

	/**
	 * Sets the Address parameter of this Address.
	 *
	 * @param address - the Address object that represents the new
	 *  address of this Address.
	 */
	SetAddress(addr address.Address)

	/**
	 * Gets the address parameter of this Address.
	 *
	 * @return the Address of this Address
	 */
	GetAddress() address.Address
}
