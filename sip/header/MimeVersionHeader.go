package header

/**
 * SIP messages MAY include a single MIME-Version general-header field to
 * indicate what version of the MIME protocol was used to construct the
 * message. Use of the MIME-Version header field indicates that the message is
 * in full compliance with the MIME protocol as defined in
 * <a href = "http://www.ietf.org/rfc/rfc2405.txt">RFC2045</a>. Proxies/gateways
 * are responsible for ensuring full compliance (where possible) when exporting
 * SIP messages to strict MIME environments.
 * <p>
 * For Example:<br>
 * <code>MIME-Version: 1.0</code>
 */
type MimeVersionHeader interface {
	Header

	/**
	 * Gets the Minor version value of this MimeVersionHeader.
	 *
	 * @return the Minor version of this MimeVersionHeader
	 */
	GetMinorVersion() int

	/**
	 * Sets the Minor-Version argument of this MimeVersionHeader to the supplied
	 * minorVersion value.
	 *
	 * @param minorVersion - the new minor MIME version
	 * @throws InvalidArgumentException if the supplied value is less than zero.
	 */
	SetMinorVersion(minorVersion int) (InvalidArgumentException error)

	/**
	 * Gets the Major version value of this MimeVersionHeader.
	 *
	 * @return the Major version of this MimeVersionHeader
	 */
	GetMajorVersion() int

	/**
	 * Sets the Major-Version argument of this MimeVersionHeader to the supplied
	 * majorVersion value.
	 *
	 * @param majorVersion - the new major MIME version
	 * @throws InvalidArgumentException if the supplied version is less than zero.
	 */
	SetMajorVersion(majorVersion int) (InvalidArgumentException error)
}
