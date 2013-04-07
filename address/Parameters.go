/**
 * ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
 * Module Name   : GoSIP Specification
 * File Name     : Parameters.go
 * Author        : Rain Liu
 *~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
 */
 
package address

import (
	"container/list"
)

/**
 * This interface defines methods for accessing generic parameters for 
 * Headers that contain generic parameter values.
 *
 * @version 1.1
 * @author Sun Microsystems
 */
type Parameters interface {

    /**
     * Returns the value of the named parameter, or null if it is not set. A
     * zero-length String indicates flag parameter.
     *
     * @param <var>name</var> name of parameter to retrieve
     * @return the value of specified parameter
     */
    GetParameter(name string) string;

    /**
     * Sets the value of the specified parameter. If the parameter already had
     * a value it will be overwritten. A zero-length String indicates flag
     * parameter.
     *
     * @param name - a String specifying the parameter name
     * @param value - a String specifying the parameter value
     * @throws ParseException which signals that an error has been reached
     * unexpectedly while parsing the parameter name or value.
     */
    SetParameter(name, value string) (ParseException error);

    /**
     * Returns an Iterator over the names (Strings) of all parameters present
     * in this ParametersHeader.
     *
     * @return an Iterator over all the parameter names
     */
    GetParameterNames() *list.List;//Iterator

    /**
     * Removes the specified parameter from Parameters of this ParametersHeader.
     * This method returns silently if the parameter is not part of the
     * ParametersHeader.
     *
     * @param name - a String specifying the parameter name
     */
    RemoveParameter(name string);

}

