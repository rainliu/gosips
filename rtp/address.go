package rtp

/** Identifies the actual implementation being used. */
type AddressType uint8

const (
	IPv4AddressType        AddressType = iota /**< Used by the UDP over IPv4 transmitter. */
	IPv6AddressType                           /**< Used by the UDP over IPv6 transmitter. */
	HostAddressType                           /**< A very general type of address, consisting of a port number and a number of bytes representing the host address. */
	UserDefinedAddressType                    /**< Can be useful for a user-defined transmitter. */
)

/** This class is an abstract class which is used to specify destinations, multicast groups etc. */
type RTPAddress interface {
	/** Returns the type of address the actual implementation represents. */
	GetAddressType() AddressType

	/** Creates a copy of the RTPAddress instance.
	 *  Creates a copy of the RTPAddress instance. If \c mgr is not NULL, the
	 *  corresponding memory manager will be used to allocate the memory for the address
	 *  copy.
	 */
	Clone() RTPAddress

	/** Checks if the address \c addr is the same address as the one this instance represents.
	 *  Checks if the address \c addr is the same address as the one this instance represents.
	 *  Implementations must be able to handle a NULL argument.
	 */
	IsSameAddress(addr RTPAddress) bool

	/** Checks if the address \c addr represents the same host as this instance.
	 *  Checks if the address \c addr represents the same host as this instance. Implementations
	 *  must be able to handle a NULL argument.
	 */
	IsFromSameHost(addr RTPAddress) bool

	String() string
}
