package header

import ()

/**
 * This interface represents the AllowEvents SIP header, as defined by
 * <a href = "http://www.ietf.org/rfc/rfc3265.txt">RFC3265</a>, this header is
 * not part of RFC3261.
 * <p>
 * The AllowEventsHeader includes a list of tokens which indicates the event
 * packages supported by the client (if sent in a request) or server (if sent
 * in a response). In other words, a node sending an AllowEventsHeader is
 * advertising that it can process SUBSCRIBE requests and generate NOTIFY
 * requests for all of the event packages listed in that header.
 * <p>
 * Any node implementing one or more event packages SHOULD include an appropriate
 * AllowEventsHeader indicating all supported events in all methods which
 * initiate dialogs and their responses (such as INVITE) and OPTIONS responses.
 * This information is very useful, for example, in allowing user agents to
 * render particular interface elements appropriately according to whether the
 * events required to implement the features they represent are supported by
 * the appropriate nodes.
 * <p>
 * Note that "Allow-Events" headers MUST NOT be inserted by proxies.
 *
 * @since v1.1
 * @author Sun Microsystems
 */

type AllowEventsHeader interface {
	Header

	/**
	 * Sets the eventType defined in this AllowEventsHeader.
	 *
	 * @param eventType - the String defining the method supported
	 * in this AllowEventsHeader
	 * @throws ParseException which signals that an error has been reached
	 * unexpectedly while parsing the Strings defining the eventType supported
	 */
	SetEventType(eventType string) //throws ParseException;

	/**
	 * Gets the eventType of the AllowEventsHeader.
	 *
	 * @return the String object identifing the eventTypes of AllowEventsHeader.
	 */
	GetEventType() string

	/**
	 * Name of AllowEventsHeader
	 */
	//public final static String NAME = "Allow-Events";
}
