package header

/**

 * The Min-Expires header field conveys the minimum refresh interval supported

 * for soft-state elements managed by that server. This includes Contact

 * header fields that are stored by a registrar. The header field contains a

 * decimal integer number of seconds from 0 to (2**32)-1.

 * <p>

 * Allowing the registrar to set the registration interval protects it against

 * excessively frequent registration refreshes while limiting the state that it

 * needs to maintain and decreasing the likelihood of registrations going

 * stale. The expiration interval of a registration is frequently used in the

 * creation of services.  An example is a follow-me service, where the user may

 * only be available at a terminal for a brief period. Therefore, registrars

 * should accept brief registrations; a request should only be rejected if the

 * interval is so short that the refreshes would degrade registrar performance.

 * <p>

 * If a User Agent receives a 423 (Interval Too Brief) response to a REGISTER request,

 * it MAY retry the registration after making the expiration interval of all

 * contact addresses in the REGISTER request equal to or greater than the

 * expiration interval within the Min-Expires header field of the 423

 * (Interval Too Brief) response. The Min-Expires header field that states the

 * minimum expiration interval the registrar is willing to honor.

 * <p>

 * For Example:<br>

 * <code>Min-Expires: 60</code>

 */

type MinExpiresHeader interface {
	ExpiresHeader
}
