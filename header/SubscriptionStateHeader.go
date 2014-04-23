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
*
* @since v1.1
* @author Sun Microsystems
*/

type SubscriptionStateHeader interface {
	ParametersHeader
	Header

	/**
	 * Sets the relative expires value of the SubscriptionStateHeader. The
	 * expires value MUST be greater than zero and MUST be less than 2**31.
	 *
	 * @param expires - the new expires value of this SubscriptionStateHeader.
	 * @throws InvalidArgumentException if supplied value is less than zero.
	 */
	SetExpires(expires int) // throws InvalidArgumentException;

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
	SetRetryAfter(retryAfter int) //throws InvalidArgumentException;

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
	SetReasonCode(reasonCode string) //throws ParseException;

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
	SetState(state string) //throws ParseException;

	/**
	 * Name of SubscriptionStateHeader
	 */
	//public final static String NAME = "Subscription-State";

	//Reason Code Constants

	/**
	 * Reason Code: The reason why the subscription was terminated is Unknown.
	 */
	//public final static String UNKNOWN = "Unknown";

	/**
	 * Reason Code: The subscription has been terminated, but the subscriber SHOULD retry
	 * immediately with a new subscription. One primary use of such a status
	 * code is to allow migration of subscriptions between nodes. The
	 * "retry-after" parameter has no semantics for "deactivated".
	 */
	//public final static String DEACTIVATED = "Deactivated";

	/**
	 * Reason Code: The subscription has been terminated, but the client SHOULD retry at
	 * some later time. If a "retry-after" parameter is also present, the client
	 * SHOULD wait at least the number of seconds specified by that parameter
	 * before attempting to re-subscribe.
	 */
	//public final static String PROBATION = "Probation";

	/**
	 * Reason Code: The subscription has been terminated due to change in authorization
	 * policy. Clients SHOULD NOT attempt to re-subscribe. The "retry-after"
	 * parameter has no semantics for "rejected".
	 */
	//public final static String REJECTED = "Rejected";

	/**
	 * Reason Code: The subscription has been terminated because it was not refreshed before
	 * it expired. Clients MAY re-subscribe immediately. The "retry-after"
	 * parameter has no semantics for "timeout".
	 */
	//public final static String TIMEOUT = "Timeout";

	/**
	 * Reason Code: The subscription has been terminated because the notifier could not
	 * obtain authorization in a timely fashion. If a "retry-after" parameter
	 * is also present, the client SHOULD wait at least the number of seconds
	 * specified by that parameter before attempting to re-subscribe; otherwise,
	 * the client MAY retry immediately, but will likely Get put back into
	 * pending state.
	 */
	//public final static String GIVE_UP = "Give-Up";

	/**
	 * Reason Code: The subscription has been terminated because the resource state which was
	 * being monitored no longer exists. Clients SHOULD NOT attempt to
	 * re-subscribe. The "retry-after" parameter has no semantics for "noresource".
	 */
	// public final static String NO_RESOURCE = "No-Resource";

	// State constants

	/**
	 * State: The subscription has been accepted and (in general) has been
	 * authorized.
	 */
	//public final static String ACTIVE = "Active";

	/**
	 * State: The subscription has been terminated, if a reason code is present,
	 * the client should behave as described in the reason code.
	 */
	// public final static String TERMINATED = "Terminated";

	/**
	 * State: The subscription has been received by the notifier, but there is
	 * insufficient policy information to grant or deny the subscription yet.
	 */
	// public final static String PENDING = "Pending";

}
