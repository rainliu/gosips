package address

import (
	"strings"
)

/** Implementation of the URI class. This relies on the 1.4 URI class.
 */

type URIImpl struct {
	/** Imbedded URI
	 */
	uriString string

	scheme string
}

/** Constructor given the URI string
 * @param uriString The imbedded URI string.
 * @throws URISyntaxException When there is a syntaz error in the imbedded URI.
 */
func NewURIImpl(uriString string) *URIImpl {
	i := strings.Index(uriString, ":")
	if i > 0 {
		return &URIImpl{uriString: uriString, scheme: uriString[0:i]}
	}

	return nil
}

/** Encode the URI.
 * @return The encoded URI
 */
func (this *URIImpl) String() string {
	return this.uriString

}

/** Encode this URI.
 * @return The encoded URI
 */
/*func (this *URIImpl) ToString() string {
    return this.Encode();

}*/

/** Overrides the base clone method
 * @return The Cloned strucutre,
 */
func (this *URIImpl) Clone() interface{} {
	//try {
	return NewURIImpl(this.uriString)

	//}
	//catch ( Exception ex){
	//    throw new RuntimeException(ex.getMessage() + this.uriString);
	//}
}

/** Returns the value of the "scheme" of
 * this URI, for example "sip", "sips" or "tel".
 *
 * @return the scheme paramter of the URI
 */
func (this *URIImpl) GetScheme() string {
	return this.scheme
}

/** This method determines if this is a URI with a scheme of
 * "sip" or "sips".
 *
 * @return true if the scheme is "sip" or "sips", false otherwise.
 */
func (this *URIImpl) IsSipURI() bool {
	var uri URI = this
	_, ok := uri.(SipURI)
	return ok //this instanceof SipURIImpl;
}
