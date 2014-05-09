package header

import "container/list"

/**
 * The User-Agent header field contains information about the UAC originating
 * the request. This is for statistical purposes, the tracing of protocol
 * violations, and automated recognition of user agents for the sake of
 * tailoring Responses to avoid particular user agent limitations. However
 * revealing the specific software version of the user agent might allow the
 * user agent to become more vulnerable to attacks against software that is
 * known to contain security holes. Implementers SHOULD make the User-Agent
 * header field a configurable option.
 * <p>
 * For Example:<br>
 * <code>User-Agent: Softphone Beta1.5</code>
 *
 * @see ServerHeader
 * @see ViaHeader
 *
 * @version 1.1
 * @author Sun Microsystems
 */
type UserAgentHeader interface {
	Header

	/**
	 * Returns the List of product values.
	 *
	 * @return the List of strings identifying the software of this ServerHeader
	 */
	GetProduct() *list.List

	/**
	 * Sets the List of product values of the ServerHeader.
	 *
	 * @param product - a List of Strings specifying the product values
	 * @throws ParseException which signals that an error has been reached
	 * unexpectedly while parsing the List of product value strings.
	 */
	SetProduct(product *list.List) //throws ParseException;

	/**
	 * Name of UserAgentHeader
	 */
	// public final static String NAME = "User-Agent";

}
