package address

import (
	//"errors"
	"strings"
)
    
    
/** Implementation of the URI class. This relies on the 1.4 URI class.
 *
 *@author Rain Liu <br/>
 *
 *<a href="{@docRoot}/uncopyright.html">This code is in the public domain.</a>
 */
    
type GenericURI struct {// implements gosip.address.URI
       
    /** Imbedded URI
     */    
    uriString string;
    
    scheme string;
}

    /** Consturctor
     */    
    //protected GenericURI() {}
    
    /** Constructor given the URI string
     * @param uriString The imbedded URI string.
     * @throws URISyntaxException When there is a syntaz error in the imbedded URI.
     */    
    func NewGenericURI(uriString string) *GenericURI {
        	i := strings.Index(uriString, ":");
            if i>0 {
            	return &GenericURI{uriString:uriString, scheme:uriString[0:i]};
        	}
        	
			return nil;
    }
    
    /** Encode the URI.
     * @return The encoded URI
     */    
    func (this *GenericURI) Encode() string {
       return this.uriString;
       
    }
    
    /** Encode this URI.
     * @return The encoded URI
     */
    /*func (this *GenericURI) ToString() string {
        return this.Encode(); 
     
    }*/
    
    /** Overrides the base clone method
     * @return The Cloned strucutre,
     */    
    func (this *GenericURI) Clone() interface{}  {
        //try {
            return NewGenericURI(this.uriString);
            
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
    func (this *GenericURI) GetScheme() string {
       return this.scheme;
    }
    
    /** This method determines if this is a URI with a scheme of
     * "sip" or "sips".
     *
     * @return true if the scheme is "sip" or "sips", false otherwise.
     */
    func (this *GenericURI) IsSipURI() bool {
    	var uri URI = this;
    	_, ok := uri.(SipURI);
        return ok; //this instanceof SipUri;
    }

