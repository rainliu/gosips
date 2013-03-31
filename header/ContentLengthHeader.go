/**
 * ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
 * Module Name   : GoSIP Specification
 * File Name     : ContentLengthHeader.go
 * Author        : Rain Liu
 *~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
 */
 
package header

import (

)


/**

 * The Content-Length header field indicates the size of the message-body, in

 * decimal number of octets, sent to the recipient. Applications SHOULD use

 * this field to indicate the size of the message-body to be transferred,

 * regardless of the media type of the entity.  If a stream-based protocol

 * (such as TCP) is used as transport, the header field MUST be used.

 * <p>

 * The size of the message-body does not include the CRLF separating header

 * fields and body.  Any Content-Length greater than or equal to zero is a

 * valid value.  If no body is present in a message, then the Content-Length

 * header field value MUST be set to zero.

 *

 * @see ContentDispositionHeader

 * @see ContentTypeHeader

 * @see ContentEncodingHeader

 * @see ContentLanguageHeader

 *

 * @version 1.1

 * @author Sun Microsystems

 */



type ContentLengthHeader interface{
	Header

    /**

     * Set content-length of ContentLengthHeader. The content-length must be

     * greater than or equal to zero.

     *

     * @param <var>contentLength</var> the content-length of the message body

     * as a decimal number of octets.

     * @throws InvalidArgumentException if contentLength is less than zero.

     */

    SetContentLength(contentLength int) (InvalidArgumentException error);



    /**

     * Gets content-length of the message body.

     *

     * @return content-length of the message body as a decimal number of octets.

     */

    GetContentLength() int;



    /**

     * Name of ContentLengthHeader

     */

    //public final static String NAME = "Content-Length";

}

