/**
 * ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
 * Module Name   : GoSIP Specification
 * File Name     : ExtensionHeader.go
 * Author        : Rain Liu
 *~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
 */

package header

/**
 * This interface represents an Extension SIP header that is not currently
 * defined by JAIN SIP. Extension Headers can be added as required by extending
 * this interface assuming other endpoints understand the Header. Any Header that
 * extends this class must define a "NAME" String constant identifying the name
 * of the extension Header. A server must ignore Headers that it does not
 * understand. A proxy must not remove or modify Headers that it does not
 * understand.
 */

type ExtensionHeader interface {
	Header

	/**
	 * Sets the value parameter of the ExtensionHeader.
	 *
	 * @param value - the new value of the ExtensionHeader
	 * @throws ParseException which signals that an error has been reached
	 * unexpectedly while parsing the value parameter.
	 */
	SetValue(value string) (ParseException error)

	/**
	 * Gets the value of the ExtensionHeader.
	 *
	 * @return the string of the value parameter.
	 */
	//GetValue() string;

}
