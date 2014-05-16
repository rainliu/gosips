package header

/**

 * The Accept-Language header field is used in requests to indicate the

 * preferred languages for reason phrases, session descriptions, or

 * status responses carried as message bodies in the response.  If no

 * Accept-Language header field is present, the server SHOULD assume all

 * languages are acceptable to the client. The q-value is used in a similar

 * manner to AcceptHeader to indicate degrees of preference.

 * <p>

 * For Example:<br>

 * <code>Accept-Language: da, en-gb;q=0.8, en;q=0.7</code>

 *

 * @see AcceptHeader

 */

type AcceptLanguageHeader interface {
	ParametersHeader
	QValue
	/**

	 * Gets the language value of the AcceptLanguageHeader.

	 *

	 * @return the language Locale value of this AcceptLanguageHeader

	 */

	GetAcceptLanguage() string

	/**

	 * Sets the language parameter of this AcceptLanguageHeader.

	 *

	 * @param language - the new Locale value of the language of

	 * AcceptLanguageHeader
	 */

	SetAcceptLanguage(language string)
}
