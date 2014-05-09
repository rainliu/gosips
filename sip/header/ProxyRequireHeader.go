package header

/**

 * The Proxy-Require header field is used to indicate proxy-sensitive features

 * that must be supported by the proxy. The Proxy-Require header field contains

 * a list of option tags. Each option tag defines a SIP extension that MUST be

 * understood by the proxy to process the request. Any ProxyRequireHeader

 * features that are not supported by the proxy must be negatively acknowledged

 * by the proxy to the client if not supported. Proxy servers treat this field

 * identically to the RequireHeader.

 * <p>

 * For Example:<br>

 * <code>Proxy-Require: foo</code>

 *

 * @see RequireHeader

 * @version 1.1

 * @author Sun Microsystems

 */

type ProxyRequireHeader interface {
	RequireHeader

	/**

	 * Name of ProxyRequireHeader

	 */

	//public final static String NAME = "Proxy-Require";

}
