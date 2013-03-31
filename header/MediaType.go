/**
 * ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
 * Module Name   : GoSIP Specification
 * File Name     : MediaType.go
 * Author        : Rain Liu  
 *~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
 */
 
package header

import (

)

/**
 * This interface represents media type methods for any header that contain
 * content type and content sub-type values. 
 *
 * @see AcceptHeader
 * @see ContentTypeHeader
 *
 * @since 1.1
 * @author Sun Microsystems
 */

type MediaType interface {

    /**
     * Sets value of media type of Header with Content Type.
     *
     * @param contentType - the new string value of the content type
     * @throws ParseException which signals that an error has been reached
     * unexpectedly while parsing the contentType value.
     */
    SetContentType(contentType string) (ParseException error);

    /**
     * Gets media type of Header with Content type.
     *
     * @return media type of Header with Content type.
     */
    GetContentType() string;

    /**
     * Sets value of media subtype of Header with Content sub-type.
     *
     * @param contentSubType - the new string value of the content sub-type.
     * @throws ParseException which signals that an error has been reached
     * unexpectedly while parsing the contentSubType value.
     */
    SetContentSubType(contentSubType string) (ParseException error);

    /**
     * Gets media sub-type of Header with Content sub-type.
     *
     * @return media sub-type of Header with Content sub-type.
     */
    GetContentSubType() string;

}

