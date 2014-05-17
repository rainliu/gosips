package header

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
 */
type ContentLanguageHeader interface {
	Header

	/**
	 * Gets the language value of the ContentLanguageHeader.
	 *
	 * @return the Locale value of this ContentLanguageHeader
	 */
	GetContentLanguage() string

	/**
	 * Sets the language parameter of this ContentLanguageHeader.
	 *
	 * @param language - the new Locale value of the language of
	 * ContentLanguageHeader
	 */
	SetContentLanguage(language string)
}
