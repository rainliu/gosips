/**
 * ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
 * Module Name   : GoSIP Specification
 * File Name     : ContentLanguageHeader.go
 * Author        : Rain Liu
 *~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
 */

package header

import (

)

/**
 * The Content-Language header field is used to indicate the language of the
 * message body.
 * <p>
 * For Example:<br>
 * <code>Content-Language: fr</code>
 *
 * @see ContentDispositionHeader
 * @see ContentLengthHeader
 * @see ContentEncodingHeader
 * @see ContentTypeHeader
 *
 * @since 1.1
 * @author Sun Microsystems
 *
 */
type ContentLanguageHeader interface {
	Header

    /**
     * Gets the language value of the ContentLanguageHeader.
     *
     * @return the Locale value of this ContentLanguageHeader
     */
    GetContentLanguage() string;//Locale;

    /**
     * Sets the language parameter of this ContentLanguageHeader.
     *
     * @param language - the new Locale value of the language of
     * ContentLanguageHeader
     */
    SetContentLanguage(language string);//Locale);

    /**
     * Name of ContentLanguageHeader
     */
    //public final static String NAME = "Content-Language";

}

