package header

import (
	"strings"
	"gosip/core"
)

/**
*  Protocol name and version.
*/
type Protocol struct{// SIPObject {
    
        /** protocolName field
         */    
	protocolName string;
        
        /** protocolVersion field
         */        
	protocolVersion string;
        
        /** transport field
         */        
	transport string;
}
	/** 
	* Default constructor.
	*/
	func NewProtocol() *Protocol{
		this := &Protocol{}
		this.protocolName = "SIP";
		this.protocolVersion = "2.0";
		this.transport = "UDP";	
		
		return this;
	}
	
	func (this *Protocol) super() {
		this.protocolName = "SIP";
		this.protocolVersion = "2.0";
		this.transport = "UDP";	
	}
	
        /**
         * Compare two To headers for equality.
         * @return true if the two headers are the same.
         * @param other Object to set
         */   
        /*public boolean equals(Object other) {
            if (! other.getClass().equals(this.getClass())) {
                return false;
            }
            Protocol that = (Protocol) other;
            if (this.protocolName.compareToIgnoreCase(that.protocolName) != 0) {
                return false;
            }
            if (this.transport.compareToIgnoreCase(that.transport) != 0) {
                return false;
            }
            if (this.protocolVersion.compareTo(that.protocolVersion) != 0) {
                return false;
            }
            return true;
        }*/
        
        /**
         * Return canonical form.
         * @return String
         */  
        func (this *Protocol) String() string {
            return  strings.ToUpper(this.protocolName) + 
            	    core.SIPSeparatorNames_SLASH + 
					this.protocolVersion +
                	core.SIPSeparatorNames_SLASH + 
                	strings.ToUpper(this.transport);
        }

        /** get the protocol name
         * @return String
         */        
	func (this *Protocol) GetProtocolName() string{
            return this.protocolName ;
        }
            
        /** get the protocol version
         * @return String
         */        
	func (this *Protocol) GetProtocolVersion() string {
            return this.protocolVersion ;
        }
        
        /** get the transport
         * @return String
         */        
	func (this *Protocol) GetTransport() string{
            return this.transport ;
        }
        
	/**
         * Set the protocolName member
         * @param p String to set
         */
	func (this *Protocol) SetProtocolName( p string) {
            this.protocolName = p ;
        }
        
	/**
         * Set the protocolVersion member
         * @param p String to set
         */
	func (this *Protocol) SetProtocolVersion( p string) {
            this.protocolVersion = p ;
        }
        
	/**
         * Set the transport member
         * @param t String to set
         */
	func (this *Protocol) SetTransport( t string) {
            this.transport = t ;
        }


	
       