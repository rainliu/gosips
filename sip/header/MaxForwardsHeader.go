package header

/**

 * The Max-Forwards header field must be used with any SIP method to limit

 * the number of proxies or gateways that can forward the request to the next

 * downstream server.  This can also be useful when the client is attempting

 * to trace a request chain that appears to be failing or looping in mid-chain.

 * <p>

 * The Max-Forwards value is an integer in the range 0-255 indicating the

 * remaining number of times this request message is allowed to be forwarded.

 * This count is decremented by each server that forwards the request. The

 * recommended initial value is 70.

 * <p>

 * This header field should be inserted by elements that can not otherwise

 * guarantee loop detection.  For example, a B2BUA should insert a Max-Forwards

 * header field.

 * <p>

 * For Example:<br>

 * <code>Max-Forwards: 6</code>

 */

type MaxForwardsHeader interface {
	Header

	/**

	 * This convenience function decrements the number of max-forwards by one.

	 * This utility is useful for proxy functionality.

	 *

	 * @throws TooManyHopsException if implementation cannot decrement

	 * max-fowards i.e. max-forwards has reached zero

	 */

	DecrementMaxForwards() (TooManyHopsException error)

	/**

	 * Gets the maximum number of forwards value of this MaxForwardsHeader.

	 *

	 * @return the maximum number of forwards of this MaxForwardsHeader

	 */

	GetMaxForwards() int

	/**

	 * Sets the max-forwards argument of this MaxForwardsHeader to the supplied

	 * <var>maxForwards</var> value.

	 *

	 * @param maxForwards - the number of max-forwards

	 * @throws InvalidArgumentException if the maxForwards argument is less

	 * than 0 or greater than 255.

	 */

	SetMaxForwards(maxForwards int) (InvalidArgumentException error)
}
