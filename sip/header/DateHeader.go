package header

import "time"

/**
 * The Date header field reflects the time when the request or response is
 * first sent. Retransmissions have the same Date header field value as the
 * original. The Date header field contains the date and time. Unlike
 * HTTP/1.1, SIP only supports the most recent
 * <a href = "http://www.ietf.org/rfc/rfc1123.txt">RFC 1123</a> format for dates.
 * SIP restricts the time zone in SIP-date to "GMT", while RFC 1123 allows any
 * time zone.
 * <p>
 * The Date header field can be used by simple end systems without a
 * battery-backed clock to acquire a notion of current time. However, in its
 * GMT form, it requires clients to know their offset from GMT.
 * <p>
 * Example:<br>
 * Date: Sat, 13 Nov 2010 23:29:00 GMT
 */

type DateHeader interface {
	Header

	/**

	 * Sets date of DateHeader. The date is repesented by the Calendar object.

	 *

	 * @param date the Calendar object date of this header.

	 */

	SetDate(date *time.Time)

	/**

	 * Gets the date of DateHeader. The date is repesented by the Calender

	 * object.

	 *

	 * @return the Calendar object representing the date of DateHeader

	 */

	GetDate() *time.Time
}
