package header

import (
	//"errors"
	"container/list"
	"gosip/core"
)


const ParameterNames_NEXT_NONCE = "nextnonce";
const ParameterNames_TAG="tag";
const ParameterNames_USERNAME="username";
const ParameterNames_URI="uri";
const ParameterNames_DOMAIN="domain";
const ParameterNames_CNONCE="cnonce";
const ParameterNames_PASSWORD="password";
const ParameterNames_RESPONSE ="response";
const ParameterNames_RESPONSE_AUTH="rspauth";
const ParameterNames_OPAQUE="opaque";
const ParameterNames_ALGORITHM="algorithm";
const ParameterNames_DIGEST="Digest";
const ParameterNames_SIGNED_BY="signed-by";
const ParameterNames_SIGNATURE="signature";
const ParameterNames_NONCE="nonce";
const ParameterNames_NONCE_COUNT="nc";
const ParameterNames_PUBKEY="pubkey";
const ParameterNames_COOKIE="cookie";
const ParameterNames_REALM="realm";
const ParameterNames_VERSION="version";
const ParameterNames_STALE="stale";
const ParameterNames_QOP="qop";
const ParameterNames_NC="nc";
const ParameterNames_PURPOSE="purpose";                           
const ParameterNames_CARD="card";                           
const ParameterNames_INFO="info";                           
const ParameterNames_ACTION="action";
const ParameterNames_PROXY="proxy";
const ParameterNames_REDIRECT="redirect"; 
const ParameterNames_EXPIRES = "expires";
const ParameterNames_Q="q";
const ParameterNames_RENDER = "render";
const ParameterNames_SESSION = "session";
const ParameterNames_ICON	= "icon";
const ParameterNames_ALERT   = "alert";
const ParameterNames_HANDLING = "handling"; 
const ParameterNames_REQUIRED = "required";
const ParameterNames_OPTIONAL = "optional";
const ParameterNames_EMERGENCY="emergency";
const ParameterNames_URGENT="urgent";
const ParameterNames_NORMAL="normal";
const ParameterNames_NON_URGENT="non-urgent";
const ParameterNames_DURATION="duration";
const ParameterNames_BRANCH="branch";
const ParameterNames_HIDDEN="hidden";
const ParameterNames_RECEIVED="received";
const ParameterNames_MADDR="maddr";
const ParameterNames_TTL="ttl";
const ParameterNames_TRANSPORT="transport";
const ParameterNames_TEXT="text";
const ParameterNames_CAUSE="cause";
const ParameterNames_ID="id";
        
/**
* Parameters header. Suitable for extension by headers that have parameters.
*/


type ParametersHeaderImpl struct{//implements javax.sip.header.Parameters {
	SIPHeaderImpl 
        
	parameters *core.NameValueList;
}
	
	//protected ParametersHeaderImpl() {
	//	this.parameters = new NameValueList();
	//}

	func NewParametersHeaderImpl(hdrName string) *ParametersHeaderImpl {
		this := &ParametersHeaderImpl{}
		this.SIPHeaderImpl.super(hdrName);
		this.parameters = core.NewNameValueList(hdrName);
		
		return this;
	}
	
	func (this *ParametersHeaderImpl) super(hdrName string){	
		this.SIPHeaderImpl.super(hdrName);
		this.parameters = core.NewNameValueList(hdrName);
	}
        
     /**
      * Returns the value of the named parameter, or null if it is not set. A
      * zero-length String indicates flag parameter.
      *
      * @param <var>name</var> name of parameter to retrieve
      * @return the value of specified parameter
      */

     func (this *ParametersHeaderImpl) GetParameter( name string) string {
            return this.parameters.GetParameter(name);
            
     }


      /** Return the parameter as an object (dont convert to string).
       *
       *@param name is the name of the parameter to get.
       *@return the object associated with the name.
       *
       */
      func (this *ParametersHeaderImpl) GetParameterValue( name string) interface{} {
            return this.parameters.GetValue(name);
      }

     /**
      * Returns an Iterator over the names (Strings) of all parameters present
      * in this ParametersHeaderImpl.
      *
      * @return an Iterator over all the parameter names
      */

    func (this *ParametersHeaderImpl) GetParameterNames() *list.List {
        return this.parameters.GetNames();
    }

     /** Return true if you have a parameter and false otherwise.
      *
      *@return true if the parameters list is non-empty.
      */

      func (this *ParametersHeaderImpl) HasParameters() bool{
		return this.parameters != nil && this.parameters.Len()!=0;
      }
    
     /**
     * Removes the specified parameter from Parameters of this ParametersHeaderImpl.
     * This method returns silently if the parameter is not part of the
     * ParametersHeaderImpl.
     *
     * @param name - a String specifying the parameter name
     */

    func (this *ParametersHeaderImpl) RemoveParameter( name string) {
    	this.parameters.Delete(name);
    }
    

    /**
     * Sets the value of the specified parameter. If the parameter already had
     *
     * a value it will be overwritten. A zero-length String indicates flag
     *
     * parameter.
     *
     *
     *
     * @param name - a String specifying the parameter name
     *
     * @param value - a String specifying the parameter value
     *
     * @throws ParseException which signals that an error has been reached
     *
     * unexpectedly while parsing the parameter name or value.
     *
     */
    func (this *ParametersHeaderImpl) SetParameter( name,  value string) (ParseException error) {
		 nv := this.parameters.GetNameValue(name);
		if nv != nil {
		   nv.SetValue(value);
		} else {
           nv  = core.NewNameValue(name,value);
           this.parameters.AddNameValue(nv);
		}
		return nil;
    }

    /**
     * Sets the value of the specified parameter. If the parameter already had
     *
     * a value it will be overwritten. A zero-length String indicates flag
     *
     * parameter.
     *
     *
     *
     * @param name - a String specifying the parameter name
     *
     * @param value - a String specifying the parameter value
     *
     * @throws ParseException which signals that an error has been reached
     *
     * unexpectedly while parsing the parameter name or value.
     *
     */
    func (this *ParametersHeaderImpl) SetQuotedParameter( name,  value string) {
		nv := this.parameters.GetNameValue(name);
		if nv != nil {
		 	nv.SetValue(value);
		 	nv.SetQuotedValue();
		} else {
           	nv  = core.NewNameValue(name,value);
	   	   	nv.SetQuotedValue();
           	this.parameters.AddNameValue(nv);
		}
    }

    /**
     * Sets the value of the specified parameter. If the parameter already had
     *
     * a value it will be overwritten.
     *
     *
     * @param name - a String specifying the parameter name
     *
     * @param value - an int specifying the parameter value
     *
     * @throws ParseException which signals that an error has been reached
     *
     * unexpectedly while parsing the parameter name or value.
     *
     */
    /*func (this *ParametersHeaderImpl) SetParameter( name string,  value int) {
	Integer val = new Integer(value);
	NameValue nv = parameters.getNameValue(name);
	if (nv != null) {
		 nv.setValue(val);
	} else {
           nv  = new NameValue(name,val);
           this.parameters.set(nv);
	}
    }*/

    /**
     * Sets the value of the specified parameter. If the parameter already had
     *
     * a value it will be overwritten. 
     *
     *
     * @param name - a String specifying the parameter name
     *
     * @param value - a boolean specifying the parameter value
     *
     * @throws ParseException which signals that an error has been reached
     *
     * unexpectedly while parsing the parameter name or value.
     *
     */
    /*protected void setParameter(String name, boolean value) {
	Boolean val = new Boolean(value);
	NameValue nv = parameters.getNameValue(name);
	if (nv != null) {
		 nv.setValue(val);
	} else {
           nv  = new NameValue(name,val);
           this.parameters.set(nv);
	}
    }*/

    /**
     * Sets the value of the specified parameter. If the parameter already had
     *
     * a value it will be overwritten. 
     *
     * @param name - a String specifying the parameter name
     *
     * @param value - a boolean specifying the parameter value
     *
     * @throws ParseException which signals that an error has been reached
     *
     * unexpectedly while parsing the parameter name or value.
     *
     */
    /*protected void setParameter(String name, float value) {
	Float val = new Float(value);
	NameValue nv = parameters.getNameValue(name);
	if (nv != null) {
		 nv.setValue(val);
	} else {
           nv  = new NameValue(name,val);
           this.parameters.set(nv);
	}
    }*/


    /**
     * Sets the value of the specified parameter. If the parameter already had
     *
     * a value it will be overwritten. A zero-length String indicates flag
     *
     * parameter.
     *
     *
     *
     * @param name - a String specifying the parameter name
     *
     * @param value - a String specifying the parameter value
     *
     * @throws ParseException which signals that an error has been reached
     *
     * unexpectedly while parsing the parameter name or value.
     *
     */
    /*protected void setParameter(String name, Object value) {
	NameValue nv = parameters.getNameValue(name);
	if (nv != null) {
		 nv.setValue(value);
	} else {
           nv  = new NameValue(name,value);
           this.parameters.set(nv);
	}
    }*/


     /** Return true if has a parameter.
      *
      *@param paramName is the name of the parameter.
      *
      *@return true if the parameter exists and false if not.
      */
     func (this *ParametersHeaderImpl) HasParameter( parameterName string) bool {
		return this.parameters.HasNameValue(parameterName);
     }
	
   /**
    *Remove all parameters.
    */
   func (this *ParametersHeaderImpl) RemoveParameters() {
       this.parameters.Init();// = new NameValueList();
   }
    
   /**
    * get the parameter list.
    * @return parameter list
    */
    func (this *ParametersHeaderImpl) GetParameters () *core.NameValueList{
        return this.parameters;
    }

    /** Set the parameter given a name and value.
     *
     * @param nameValue - the name value of the parameter to set.
     */
    //func (this *ParametersHeaderImpl) SetParameter( nameValue *NameValue) {
	//System.out.println("setParameter " + this + " nbv = " + nameValue);
	//	this.parameters.AddNameValue(nameValue);
    //}
    
    /** Set the parameter list.
     *
     *@param nameValueList - the name value list to set as the parameter list.
     */
    func (this *ParametersHeaderImpl) SetParameters(  parameters *core.NameValueList) {
        this.parameters = parameters;
    }
    
    
    /** Get the parameter as an integer value.
     *
     *@param parameterName -- the parameter name to fetch.
     *
     *@return -1 if the parameter is not defined in the header.
     */
    /*func (this *ParametersHeaderImpl) GetParameterAsInt(String parameterName) int{
        if (this.getParameterValue(parameterName) != null) {
                try {
                  if (this.getParameterValue(parameterName) 
                        instanceof String) {
                     return Integer.parseInt
                        (this.getParameter(parameterName));
                  } else {
                     return 
                       ((Integer)getParameterValue(parameterName)).intValue();
                  }
                } catch (NumberFormatException ex) {
                    return -1;
                }
            } else return -1;
    }*/

    /** Get the parameter as an integer when it is entered as a hex.
     *
     *@param parameterName -- The parameter name to fetch.
     *
     *@return -1 if the parameter is not defined in the header.
     */
    /*protected int getParameterAsHexInt(String parameterName) {
        if (this.getParameterValue(parameterName) != null) {
                try {
                  if (this.getParameterValue(parameterName) 
                        instanceof String) {
                     return Integer.parseInt
                        (this.getParameter(parameterName),16);
                  } else {
                     return 
                       ((Integer)getParameterValue(parameterName)).intValue();
                  }
                } catch (NumberFormatException ex) {
                    return -1;
                }
            } else return -1;
    }*/


    
    /** Get the parameter as a float value.
     *
     *@param parameterName -- the parameter name to fetch
     *
     *@return -1 if the parameter is not defined or the parameter as a float.
     */
    /*protected float getParameterAsFloat(String parameterName) {
        
      if (this.getParameterValue(parameterName) != null) {
                try {
                  if (this.getParameterValue(parameterName) 
                        instanceof String) {
                     return Float.parseFloat
                        (this.getParameter(parameterName));
                  } else {
                     return 
                       ((Float)getParameterValue(parameterName)).floatValue();
                  }
                } catch (NumberFormatException ex) {
                    return -1;
                }
            } else return -1;
    }*/

    /** Get the parameter as a long value.
     *
     *@param parameterName -- the parameter name to fetch.
     *
     *@return -1 if the parameter is not defined or the parameter as a long.
     */
    
     /*protected long getParameterAsLong(String parameterName) {
        if (this.getParameterValue(parameterName) != null) {
                try {
                  if (this.getParameterValue(parameterName) 
                        instanceof String) {
                     return Long.parseLong
                        (this.getParameter(parameterName));
                  } else {
                     return 
                       ((Long)getParameterValue(parameterName)).longValue();
                  }
                } catch (NumberFormatException ex) {
                    return -1;
                }
            } else return -1;
     }*/
    
    /** get the parameter value as a URI.
     *
     *@param parameterName -- the parameter name 
     *
     *@return value of the parameter as a URI or null if the parameter
     *  not present.
     */
    /*protected GenericURI getParameterAsURI(String parameterName) {
      Object val =   getParameterValue (parameterName);
	    if (val instanceof GenericURI)  return (GenericURI) val;
	    else {
                try{
                    return new GenericURI((String)val);
                }
                catch (ParseException ex) {
                //catch ( URISyntaxException ex) { 
                      return null; 
                }
            }
    }*/

	/** Get the parameter value as a boolean.
	*
	*@param parameterName -- the parameter name
	*@return boolean value of the parameter.
	*/
	/*protected boolean getParameterAsBoolean(String parameterName) {
		Object val = getParameterValue(parameterName);
		if (val == null) {
                        return false;
                } else if (val instanceof Boolean) {
			return ((Boolean) val).booleanValue();
		} else if (val instanceof  String) {
			    return Boolean.getBoolean((String)val);
		} else return false;
	}*/
		

	/** This is for the benifit of the TCK.
	*
	*@return the name value pair for the given parameter name.
	*/
	func (this *ParametersHeaderImpl) GetNameValue( parameterName string) *core.NameValue {
		return this.parameters.GetNameValue(parameterName);
	}
	
    
    
    func (this *ParametersHeaderImpl) EncodeBody() string{
    	return ""
    }