package header

import (
	"gosips/sip/address"
)

/**
 * When present in an INVITE request, the Alert-Info header field
 * specifies an alternative ring tone to the UAS.  When present in a 180
 * (Ringing) response, the Alert-Info header field specifies an
 * alternative ringback tone to the UAC.  A typical usage is for a proxy
 * to insert this header field to provide a distinctive ring feature.
 * <p>
 * The Alert-Info header field can introduce security risks, which are
 * identical to the Call-Info header field risk, see section 20.9 of
 * <a href = "http://www.ietf.org/rfc/rfc3261.txt">RFC3261</a>.
 * In addition, a user SHOULD be able to disable this feature selectively.
 * This helps prevent disruptions that could result from the use of this
 * header field by untrusted elements.
 * <p>
 * For Example:<br>
 * <code>Alert-Info: <http://jcp.org/yeeha.wav></code>
 */

type AlertInfoHeader interface {
	ParametersHeader

	/**
	 * Sets the AlertInfo of the AlertInfoHeader to the <var>alertInfo</var>
	 * parameter value.
	 *
	 * @param alertInfo the new Alert Info URI of this AlertInfoHeader.
	 */
	SetAlertInfo(alertInfo address.URI)

	/**
	 * Returns the AlertInfo value of this AlertInfoHeader.
	 *
	 * @return the URI representing the AlertInfo.
	 */
	GetAlertInfo() address.URI
}
