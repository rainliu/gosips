/**
 * ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
 * Module Name   : GoSIP Specification
 * File Name     : URI.go
 * Author        : Rain Liu
 *~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
 */
 
package address

import (

)

/**
 * This class represents a generic URI. This is the base interface for any 
 * type of URI. These are used in SIP requests to identify the callee and also 
 * in Contact, From, and To headers. 
 * <p>
 * The generic syntax of URIs is defined in 
 * <a href = http://www.ietf.org/rfc/rfc2396.txt>RFC 2396</a>. 
 *
 * @see TelURL
 * @see SipURI
 *
 * @author Sun Microsystems
 * @since 1.1
 */

type URI interface{// extends Cloneable, Serializable{

    
    /**
     * Returns the value of the "scheme" of this URI, for example "sip", "sips" 
     * or "tel".
     *
     * @return the scheme paramter of the URI
     */
    GetScheme() string;

    /**
     * Creates and returns a deep copy of the URI. This methods must ensure a
     * deep copy of the URI, so that when a URI is cloned the URI can be 
     * modified without effecting the original URI. This provides useful 
     * functionality for proxying Requests and Responses. This method overrides 
     * the clone method in java.lang.Object.
     *
     * @return a deep copy of URI
     */
    Clone() interface{};
    
    /**
     * This method determines if this is a URI with a scheme of "sip" or "sips". 
     *
     * @return true if the scheme is "sip" or "sips", false otherwise.
     */        
    IsSipURI() bool;
    
    /**
     * This method returns the URI as a string. 
     *
     * @return String The stringified version of the URI
     */    
    Encode() string;
            
}

