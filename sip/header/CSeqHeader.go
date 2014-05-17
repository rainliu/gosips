package header

/**
 * A CSeq header field in a request contains a single decimal sequence number
 * and the request method. The CSeq header field serves to identify and order
 * transactions within a dialog, to provide a means to uniquely identify
 * transactions, and to differentiate between new requests and request
 * retransmissions.  Two CSeq header fields are considered equal if the
 * sequence number and the request method are identical.
 * <ul>
 * <li>Method - The method part of CSeq is case-sensitive and MUST match that
 * of the request.
 * <li>Sequence number - The sequence number is chosen by the requesting
 * client and is unique within a single value of Call-ID. The sequence number
 * MUST be expressible as a 32-bit unsigned integer and MUST be less than
 * 2**31. For non-REGISTER requests outside of a dialog, the sequence number
 * value is arbitrary. Consecutive Requests that differ in method, headers or
 * body, but have the same CallIdHeader must contain strictly monotonically
 * increasing and contiguous sequence numbers; sequence numbers do not wrap
 * around. Retransmissions of the same Request carry the same sequence number,
 * but an INVITE Request with a different message body or different headers
 * (a "re-invitation") acquires a new, higher sequence number. A server must
 * echo the CSeqHeader from the Request in its Response. If the method value is
 * missing in the received CSeqHeader, the server fills it in appropriately.
 * ACK and CANCEL Requests must contain the same CSeqHeader sequence number
 * (but not method) as the INVITE Request they refer to, while a BYE Request
 * cancelling an invitation must have a higher sequence number. An user agent
 * server must remember the highest sequence number for any INVITE Request
 * with the same CallIdHeader. The server must respond to, and then discard,
 * any INVITE Request with a lower sequence number.
 * </ul>
 * As long as a client follows the above guidelines, it may use any mechanism
 * it would like to select CSeq header field values.
 * <p>
 * <b>Forked Requests:</b><br>
 * Forked Requests must have the same CSeqHeader as there would be ambiguity
 * otherwise between these forked Requests and later BYE Requests issued by the
 * client user agent.
 * <p>
 * For Example:<br>
 * <code>CSeq: 4711 INVITE</code>
 *
 */

type CSeqHeader interface {
	Header

	/**
	 * Sets the method of CSeqHeader
	 *
	 * @param method - the method of the Request of this CSeqHeader
	 * @throws ParseException which signals that an error has been reached
	 * unexpectedly while parsing the method value.
	 */
	SetMethod(method string) (ParseException error)

	/**
	 * Gets the method of CSeqHeader
	 *
	 * @return method of CSeqHeader
	 */
	GetMethod() string

	/**
	 * Sets the sequence number value of the CSeqHeader. The sequence number
	 * MUST be expressible as a 32-bit unsigned integer and MUST be less than
	 * 2**31.
	 *
	 * @param sequenceNumber - the new sequence number of this CSeqHeader
	 * @throws InvalidArgumentException if supplied value is less than zero.
	 *
	 */
	SetSequenceNumber(sequenceNumber int) (InvalidArgumentException error)

	/**
	 * Gets the sequence number of this CSeqHeader.
	 *
	 * @return sequence number of the CSeqHeader
	 *
	 */
	GetSequenceNumber() int
}
