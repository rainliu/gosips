/**
 * ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
 * Module Name   : GoSIP Specification
 * File Name     : AddressFactory.go
 * Author        : Rain Liu
 *~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
 */
 
package address

import (

)

/**
 * This interface provides factory methods that allow an application to create 
 * Address objects, URI's, SipURI's and TelURL's from a particular 
 * implementation of this specification. This class is a singleton and can be 
 * retrieved from the {@link javax.sip.SipFactory#createAddressFactory()}.
 */
type AddressFactory interface{

    /**
     * Creates a URI based on given URI string. The URI string is parsed in 
     * order to create the new URI instance. Depending on the scheme the 
     * returned may or may not be a SipURI or TelURL cast as a URI.
     *
     * @param uri - the new string value of the URI.
     * @throws ParseException if the URI string is malformed.
     */
    CreateURI(uriStr string) (uri URI, ParseException error);

    /**
     * Creates a SipURI based on the given user and host components. The user
     * component may be null. 
     * <p>
     * This create method first builds a URI in string form using the given 
     * components as follows:
     * <ul>
     * <li>Initially, the result string is empty.
     * <li>The scheme followed by a colon character ('sip:') is appended to 
     * the result.
     * <li>The user and host are then appended. Any character that is not a 
     * legal URI character is quoted.
     * </ul>
     * <br>
     * The resulting URI string is then parsed in order to create the new 
     * SipURI instance as if by invoking the createURI(String) constructor; 
     * this may cause a URISyntaxException to be thrown.
     * <p>
     * An application that wishes to create a 'sips' URI should call the 
     * {@link SipURI#setSecure(boolean)} with an argument of 'true' on the 
     * returned SipURI. 
     *
     * @param user - the new string value of the user, this value may be null.
     * @param host - the new string value of the host.
     * @throws ParseException if the URI string is malformed. 
     */
    CreateSipURI(user, host string) (sipuri SipURI, ParseException error);    

    /**
     * Creates a TelURL based on given URI string. The scheme or '+' should 
     * not be included in the phoneNumber string argument.
     *
     * @param uri - the new string value of the phoneNumber.
     * @throws ParseException if the URI string is malformed. 
     */
     CreateTelURL(phoneNumber string) (telurl TelURL, ParseException error);
    
    /**
     * Creates an Address with the new address string value. The address 
     * string is parsed in order to create the new Address instance. Create
     * with a String value of "*" creates a wildcard address. The wildcard 
     * can be determined if <code>((SipURI)Address.getURI).getUser() == *;</code>.
     *
     * @param address - the new string value of the address.
     * @throws ParseException which signals that an error has been reached
     * unexpectedly while parsing the address value. 
     */
    CreateAddressFromString(addrStr string) (addr Address, ParseException error);

    /**
     * Creates an Address with the new URI attribute value. 
     *
     * @param uri - the URI value of the address.
     */
     CreateAddressFromURI(uri URI) (addr Address);
    
    /**
     * Creates an Address with the new display name and URI attribute
     * values.
     *
     * @param displayName - the new string value of the display name of the
     * address. A <code>null</code> value does not set the display name.
     * @param uri - the new URI value of the address.
     * @throws ParseException which signals that an error has been reached
     * unexpectedly while parsing the displayName value. 
     */
    CreateAddressFromURIWithDisplayName(displayName string, uri URI) (addr Address, ParseException error);
}