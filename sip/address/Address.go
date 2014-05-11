package address

/**
 * This interface represents a user's display name and URI address. The display
 * name of an address is optional but if included can be displayed to an end-user.
 * The address URI (most likely a SipURI) is the user's address. For example a
 * 'To' address of <code>To: Bob sip:duke@jcp.org</code> would have a display
 * name attribute of <code>Bob</code> and an address of
 * <code>sip:duke@jcp.org</code>.
 *
 * @see SipURI
 * @see TelURL
 */

type Address interface {
	/**
	 * Sets the display name of the Address. The display name is an
	 * additional user friendly personalized text that accompanies the address.
	 *
	 * @param displayName - the new string value of the display name.
	 * @throws ParseException which signals that an error has been reached
	 * unexpectedly while parsing the displayName value.
	 */
	SetDisplayName(displayName string) (ParseException error)

	/**
	 * Gets the display name of this Address, or null if the attribute is
	 * not set.
	 *
	 * @return the display name of this Address
	 */
	GetDisplayName() string

	/**
	 * Sets the URI of this Address. The URI can be either a TelURL or a SipURI.
	 *
	 * @param uri - the new URI value of this Address.
	 */
	SetURI(uri URI)

	/**
	 * Returns the URI  of this Address. The type of URI can be
	 * determined by the scheme.
	 *
	 * @return URI parmater of the Address object
	 */
	GetURI() URI

	/**
	 * Returns a string representation of this Address.
	 *
	 * @return the stringified representation of the Address
	 */
	String() string //ToString() string;

	/**
	 * Indicates whether some other Object is "equal to" this Address.
	 * The actual implementation class of a Address object must override
	 * the Object.equals method. The new equals method must ensure that the
	 * implementation of the method is reflexive, symmetric, transitive and
	 * for any non null value X, X.equals(null) returns false.
	 *
	 * @param obj - the Object with which to compare this Address
	 * @return true if this Address is "equal to" the object argument and
	 * false otherwise.
	 * @see Object
	 */
	//Equals(obj interface{}) bool;

	/**
	 * This determines if this address is a wildcard address. That is
	 * <code>((SipURI)Address.getURI()).getUser() == *;</code>. This method
	 * is specific to SIP and SIPS schemes.
	 *
	 * @return true if this address is a wildcard, false otherwise.
	 */
	IsWildcard() bool
}
