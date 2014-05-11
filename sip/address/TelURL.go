package address

import (
	"container/list"
)

/**
 * This class represents Tel URLs, which are used for addressing. The Tel URL
 * starts with the scheme <code>tel:</code>. This tells the
 * local entity that what follows is a URL that should be parsed as described
 * in <a href = "http://www.ietf.org/rfc/rfc2806.txt">RFC2806</a>. After that,
 * the URL contains the phone number of the remote entity.
 * <p>
 * Within a SIP Message, TelURLs can be used to indicate the source and intended
 * destination of a Request, redirection addresses and the current destination
 * of a Request. All these Headers may contain TelURLs.
 * <p>
 * The TelURL interface extends the generic URI interface and provides
 * additional convenience methods for the following components of a TelURL
 * address, above the generic URI class:
 * <ul>
 * <li>ISDN Subaddress - Phone numbers can also contain subaddresses, which
 * are used to identify different remote entities under the same phone number.
 * <li>Post Dial - Phone numbers can also contain a post-dial sequence.
 * This is what is often used with voice mailboxes and other services that
 * are controlled by dialing numbers from your phone keypad while the call is
 * in progress.
 * <li>Global - Phone numbers can be either "global" or "local". Global numbers
 * are unambiguous everywhere. Local numbers are usable only within a certain
 * area.
 * <li>URL parameters - Parameters affecting a request constructed from this
 * URL. URL parameters are added to the end of the URL component and are
 * separated by semi-colons. URL parameters take the form:<br>
 * parameter-name "=" parameter-value
 * </ul>
 * See <a href = "http://www.ietf.org/rfc/rfc2806.txt">RFC2806</a> for more
 * information on the use of TelURL's.
 */

type TelURL interface {
	URI

	/**
	 * Returns the value of the named parameter, or null if it is not set. A
	 * zero-length String indicates flag parameter.
	 *
	 * @param <var>name</var> name of parameter to retrieve
	 * @return the value of specified parameter
	 */
	GetParameter(name string) string

	/**
	 * Sets the value of the specified parameter. If the parameter already had
	 * a value it will be overwritten. A zero-length String indicates flag
	 * parameter.
	 *
	 * @param name - a String specifying the parameter name
	 * @param value - a String specifying the parameter value
	 * @throws ParseException which signals that an error has been reached
	 * unexpectedly while parsing the parameter name or value.
	 */
	SetParameter(name, value string) (ParseException error)

	/**
	 * Returns an Iterator over the names (Strings) of all parameters present
	 * in this ParametersHeader.
	 *
	 * @return an Iterator over all the parameter names
	 */
	GetParameterNames() *list.List //Iterator

	/**
	 * Removes the specified parameter from Parameters of this ParametersHeader.
	 * This method returns silently if the parameter is not part of the
	 * ParametersHeader.
	 *
	 * @param name - a String specifying the parameter name
	 */
	RemoveParameter(name string)

	/**
	 * Returns <code>true</code> if this TelURL is global i.e. if the TelURI
	 * has a global phone user.
	 *
	 * @return <code>true</code> if this TelURL represents a global phone user,
	 * and <code>false</code> otherwise.
	 */
	IsGlobal() bool

	/**
	 * Sets phone user of this TelURL to be either global or local. The default
	 * value is false, hence the TelURL is defaulted to local.
	 *
	 * @param global - the boolean value indicating if the TelURL has a global
	 * phone user.
	 */
	SetGlobal(global bool)

	/**
	 * Sets post dial of this TelURL. The post-dial sequence describes what and
	 * when the local entity should send to the phone line.
	 *
	 * @param postDial - new value of the <code>postDial</code> parameter
	 * @throws ParseException  which signals that an error has been reached
	 * unexpectedly while parsing the postDial value.
	 */
	SetPostDial(postDial string) (ParseException error)

	/**
	 * Returns the value of the <code>postDial</code> parameter, or null if it
	 * is not set.
	 *
	 * @return  the value of the <code>postDial</code> parameter
	 */
	GetPostDial() string

	/**
	 * Sets phone number of this TelURL. The phone number may either be local or
	 * global determined by the isGlobal method in this interface. The phoneNumber
	 * argument should not contain the "+" associated with telephone numbers.
	 *
	 * @param phoneNumber - new value of the <code>phoneNumber</code> parameter
	 * @throws ParseException  which signals that an error has been reached
	 * unexpectedly while parsing the phoneNumber value.
	 */
	SetPhoneNumber(phoneNumber string) (ParseException error)

	/**
	 * Returns the value of the <code>phoneNumber</code> parameter. This method
	 * will not return the "+" associated with telephone numbers.
	 *
	 * @return  the value of the <code>phoneNumber</code> parameter
	 */
	GetPhoneNumber() string

	/**
	 * Sets ISDN subaddress of this TelURL. If a subaddress is present, it is
	 * appended to the phone number after ";isub=".
	 *
	 * @param isdnSubAddress - new value of the <code>isdnSubAddress</code>
	 * parameter
	 * @throws ParseException which signals that an error has been reached
	 * unexpectedly while parsing the isdnSubAddress value.
	 */
	SetIsdnSubAddress(isdnSubAddress string) (ParseException error)

	/**
	 * Returns the value of the <code>isdnSubAddress</code> parameter, or null
	 * if it is not set.
	 *
	 * @return  the value of the <code>isdnSubAddress</code> parameter
	 */
	GetIsdnSubAddress() string
}
