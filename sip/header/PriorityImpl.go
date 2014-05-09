package header

import "gosips/core"

/**
* the Priority header.
*
*@author Olivier Deruelle <deruelle@nist.gov><br/>
*
*<a href="{@docRoot}/uncopyright.html">This code is in the public domain.</a>
*
 */
type Priority struct {
	SIPHeader
	//implements PriorityHeader {

	/** constant EMERGENCY field
	 */
	//public static final String EMERGENCY=ParameterNames.EMERGENCY;

	/** constant URGENT field
	 */
	//public static final String URGENT= ParameterNames.URGENT;

	/** constant NORMAL field
	 */
	//public static final String NORMAL= ParameterNames.NORMAL;

	/** constant NON_URGENT field
	 */
	//public static final String NON_URGENT= ParameterNames.NON_URGENT;
	/** priority field
	 */
	priority string
}

/** Default constructor
 */
func NewPriority() *Priority {
	this := &Priority{}
	this.SIPHeader.super(core.SIPHeaderNames_PRIORITY)
	return this
}

func (this *Priority) String() string {
	return this.headerName + core.SIPSeparatorNames_COLON +
		core.SIPSeparatorNames_SP + this.EncodeBody() + core.SIPSeparatorNames_NEWLINE
}

/**
 * Encode into canonical form.
 * @return String
 */
func (this *Priority) EncodeBody() string {
	return this.priority
}

/**
 * get the priority value.
 * @return String
 */
func (this *Priority) GetPriority() string {
	return this.priority
}

/**
 * Set the priority member
 * @param p String to set
 */
func (this *Priority) SetPriority(p string) { //throws ParseException{
	//if (p==null) throw new  NullPointerException("JAIN-SIP Exception,"+
	//"Priority, setPriority(), the priority parameter is null");
	this.priority = p
}
