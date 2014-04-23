package header

/**

 * The Retry-After header field identifies the time to retry the request after
 * recipt of the response. It can be used with a 500 (Server Internal Error) or 503
 * (Service Unavailable) response to indicate how long the service is

 * expected to be unavailable to the requesting client and with a 404

 * (Not Found), 413 (Request Entity Too Large), 480 (Temporarily Unavailable),

 * 486 (Busy Here), 600 (Busy), or 603 (Decline) response to indicate when the

 * called party anticipates being available again. The value of this field is

 * a positive integer number of seconds (in decimal) after the time of the

 * response.

 * <p>

 * An optional comment can be used to indicate additional information about the

 * time of callback.  An optional "duration" parameter indicates how long the

 * called party will be reachable starting at the initial time of availability.

 * If no duration parameter is given, the service is assumed to be available

 * indefinitely.

 * <p>

 * For Examples:<br>

 * <code>Retry-After: 18000;duration=3600<br>

 * Retry-After: 120 (I'm in a meeting)</code>

 *

 * @see Parameters

 * @see Header

 * @version 1.1

 * @author Sun Microsystems

 */

type RetryAfterHeader interface {
	Header
	ParametersHeader

	/**

	 * Sets the retry after value of the RetryAfterHeader. The retry after value

	 * MUST be greater than zero and MUST be less than 2**31.

	 *

	 * @param retryAfter - the new retry after value of this RetryAfterHeader

	 * @throws InvalidArgumentException if supplied value is less than zero.

	 * @since v1.1

	 */

	SetRetryAfter(retryAfter int) // throws InvalidArgumentException;

	/**

	 * Gets the retry after value of the RetryAfterHeader. This retry after

	 * value is relative time.

	 *

	 * @return the retry after value of the RetryAfterHeader.

	 * @since v1.1

	 */

	GetRetryAfter() int

	/**

	 * Gets the comment of RetryAfterHeader.

	 *

	 * @return the comment of this RetryAfterHeader, return null if no comment

	 * is available.

	 */

	GetComment() string

	/**

	 * Sets the comment value of the RetryAfterHeader.

	 *

	 * @param comment - the new comment string value of the RetryAfterHeader.

	 * @throws ParseException which signals that an error has been reached

	 * unexpectedly while parsing the comment.

	 */

	SetComment(comment string) // throws ParseException;

	/**

	 * Sets the duration value of the RetryAfterHeader. The retry after value

	 * MUST be greater than zero and MUST be less than 2**31.

	 *

	 * @param duration - the new duration value of this RetryAfterHeader

	 * @throws InvalidArgumentException if supplied value is less than zero.

	 * @since v1.1

	 */

	SetDuration(duration int) //throws InvalidArgumentException;

	/**

	 * Gets the duration value of the RetryAfterHeader. This duration value

	 * is relative time.

	 *

	 * @return the duration value of the RetryAfterHeader, return zero if not

	 * Set.

	 * @since v1.1

	 */

	GetDuration() int

	/**

	 * Name of RetryAfterHeader

	 */

	// public final static String NAME = "Retry-After";

}
