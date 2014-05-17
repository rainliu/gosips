package header

import "container/list"

/**
 * The Server header field contains information about the software used by
 * the UAS to handle the request. Revealing the specific software version of
 * the server might allow the server to become more vulnerable to attacks
 * against software that is known to contain security holes. Implementers
 * SHOULD make the Server header field a configurable option. If the Response
 * is being forwarded through a proxy, the proxy application must not modify
 * the ServerHeaders. Instead, it should include a ViaHeader.
 * <p>
 * For Example:<br>
 * <code>Server: HomeServer v2</code>
 *
 * @see ViaHeader
 * @see UserAgentHeader
 */
type ServerHeader interface {
	Header

	/**
	 * Returns a ListIterator over the List of product values.
	 *
	 * @return a ListIterator over the List of strings identifying the
	 * software of this ServerHeader
	 */
	GetProduct() *list.List

	/**
	 * Sets the List of product values of the ServerHeader.
	 *
	 * @param product - a List of Strings specifying the product values
	 * @throws ParseException which signals that an error has been reached
	 * unexpectedly while parsing the List of product value strings.
	 */
	SetProduct(product *list.List) (ParseException error)
}
