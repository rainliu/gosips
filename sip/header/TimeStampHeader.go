package header

/**
 * The Timestamp header field describes when the UAC sent the request to the
 * UAS. When a 100 (Trying) response is generated, any Timestamp header field
 * present in the request MUST be copied into this 100 (Trying) response. If
 * there is a delay in generating the response, the UAS SHOULD add a delay
 * value into the Timestamp value in the response. This value MUST contain the
 * difference between the time of sending of the response and receipt of the
 * request, measured in seconds. Although there is no normative behavior
 * defined here that makes use of the header, it allows for extensions or
 * SIP applications to obtain RTT estimates, that may be used to adjust the
 * timeout value for retransmissions.
 * <p>
 * For Example:<br>
 * <code>Timestamp: 54</code>
 */

type TimeStampHeader interface {
	Header

	/**
	 * Sets the timestamp value of this TimeStampHeader to the new timestamp
	 * value passed to this method.
	 *
	 * @param timestamp - the new float timestamp value
	 * @throws InvalidArgumentException if the timestamp value argument is a
	 * negative value.
	 */
	SetTimeStamp(timeStamp float32) (InvalidArgumentException error)

	/**
	 * Gets the timestamp value of this TimeStampHeader.
	 *
	 * @return the timestamp value of this TimeStampHeader
	 */
	GetTimeStamp() float32

	/**
	 * Gets delay of TimeStampHeader. This method returns <code>-1</code> if the
	 * delay parameter is not set.
	 *
	 * @return the delay value of this TimeStampHeader
	 */

	GetDelay() float32

	/**
	 * Sets the new delay value of the TimestampHeader to the delay parameter
	 * passed to this method
	 *
	 * @param delay - the new float delay value
	 * @throws InvalidArgumentException if the delay value argumenmt is a
	 * negative value other than the default value <code>-1</code>.
	 */

	SetDelay(delay float32) (InvalidArgumentException error)
}
