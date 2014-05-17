package header

/**

 * The OrganizationHeader conveys the name of the organization to which the

 * entity issuing the Request or Response belongs. It may be used by client

 * software to filter calls.

 */

type OrganizationHeader interface {
	Header

	/**

	 * Sets the organization value of the OrganizationHeader to the

	 * organization parameter supplied.

	 *

	 * @param organization - the new string organization value

	 * @throws ParseException which signals that an error has been reached

	 * unexpectedly while parsing the organization value.

	 */

	SetOrganization(organization string) (ParseException error)

	/**

	 * Gets the organization value of OrganizationHeader.

	 *

	 * @return organization of OrganizationHeader

	 */

	GetOrganization() string
}
