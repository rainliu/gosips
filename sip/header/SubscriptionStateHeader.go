package header

/**
 * This interface represents the Subscription State header, as
* defined by <a href = "http://www.ietf.org/rfc/rfc3265.txt">RFC3265</a>, this
* header is not part of RFC3261.
* <p>
* NOTIFY requests MUST contain SubscriptionState headers which indicate the
* status of the subscription. The subscription states are:
* <ul>
* <li> active - If the SubscriptionState header value is "active", it means
* that the subscription has been accepted and (in general) has been authorized.
* If the header also contains an "expires" parameter, the subscriber SHOULD
* take it as the authoritative subscription duration and adjust accordingly.
* The "retry-after" and "reason" parameters have no semantics for "active".
* <li> pending - If the SubscriptionState value is "pending", the
* subscription has been received by the notifier, but there is insufficient
* policy information to grant or deny the subscription yet. If the header also
* contains an "expires" parameter, the subscriber SHOULD take it as the
* authoritative subscription duration and adjust accordingly. No further
* action is necessary on the part of the subscriber. The "retry-after" and
* "reason" parameters have no semantics for "pending".
* <li> terminated - If the SubscriptionState value is "terminated", the
* subscriber should consider the subscription terminated. The "expires"
* parameter has no semantics for "terminated". If a reason code is present, the
* client should behave as described in the reason code defined in this Header.
* If no reason code or an unknown reason code is present, the client MAY
* attempt to re-subscribe at any time (unless a "retry-after" parameter is
* present, in which case the client SHOULD NOT attempt re-subscription until
* after the number of seconds specified by the "retry-after" parameter).
* </ul>
*/

type SubscriptionStateHeader interface {
	ParametersHeader

	/**
	 * Sets the relative expires value of the SubscriptionStateHeader. The
	 * expires value MUST be greater than zero and MUST be less than 2**31.
	 *
	 * @param expires - the new expires value of this SubscriptionStateHeader.
	 * @throws InvalidArgumentException if supplied value is less than zero.
	 */
	SetExpires(expires int) (InvalidArgumentException error)

	/**
	 * Gets the expires value of the SubscriptionStateHeader. This expires value is
	 * relative time.
	 *
	 * @return the expires value of the SubscriptionStateHeader.
	 */
	GetExpires() int

	/**
	 * Sets the retry after value of the SubscriptionStateHeader. The retry after value
	 * MUST be greater than zero and MUST be less than 2**31.
	 *
	 * @param retryAfter - the new retry after value of this SubscriptionStateHeader
	 * @throws InvalidArgumentException if supplied value is less than zero.
	 */
	SetRetryAfter(retryAfter int) (InvalidArgumentException error)

	/**
	 * Gets the retry after value of the SubscriptionStateHeader. This retry after
	 * value is relative time.
	 *
	 * @return the retry after value of the SubscriptionStateHeader.
	 */
	GetRetryAfter() int

	/**
	 * Gets the reason code of SubscriptionStateHeader.
	 *
	 * @return the comment of this SubscriptionStateHeader, return null if no reason code
	 * is available.
	 */
	GetReasonCode() string

	/**
	 * Sets the reason code value of the SubscriptionStateHeader.
	 *
	 * @param reasonCode - the new reason code string value of the SubscriptionStateHeader.
	 * @throws ParseException which signals that an error has been reached
	 * unexpectedly while parsing the reason code.
	 */
	SetReasonCode(reasonCode string) (ParseException error)

	/**
	 * Gets the state of SubscriptionStateHeader.
	 *
	 * @return the state of this SubscriptionStateHeader.
	 */
	GetState() string

	/**
	 * Sets the state value of the SubscriptionStateHeader.
	 *
	 * @param state - the new state string value of the SubscriptionStateHeader.
	 * @throws ParseException which signals that an error has been reached
	 * unexpectedly while parsing the state.
	 */
	SetState(state string) (ParseException error)
}
