package header

/**
 * The Unsupported header field lists the features not supported by the UAS.
 * If a server does not understand the option, it must respond by returning a
 * 420 (Bad Extension) Response and list those options it does not understand in
 * the UnsupportedHeader.
 *
 * @see SupportedHeader
 * @see RequireHeader
 * @version 1.1
 * @author Sun Microsystems
 */

type UnsupportedHeader interface {
	OptionTag
	Header

	/**
	 * Name of UnsupportedHeader
	 */
	//public final static String NAME = "Unsupported";

}
