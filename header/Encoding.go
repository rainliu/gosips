/**
 * ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
 * Module Name   : GoSIP Specification
 * File Name     : Encoding.go
 * Author        : Rain Liu
 *~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
 */

package header

import (

)


/**
 * This interface represents encoding methods for any header that contains an
 * encoding value. 
 *
 * @see AcceptEncodingHeader
 * @see ContentEncodingHeader
 *
 * @since 1.1
 * @author Sun Microsystems
 */
type Encoding interface {

    /**
     * Sets the encoding of an EncodingHeader.
     *
     * @param encoding - the new string value defining the encoding.
     * @throws ParseException which signals that an error has been reached
     * unexpectedly while parsing the encoding value.
     */

    SetEncoding(encoding string) (ParseException error);

    /**
     * Gets the encoding of an EncodingHeader. Returns null if no 
     * encoding is defined in an EncodingHeader.
     *
     * @return the string value identifing the encoding
     */
    GetEncoding() string;    

}


