package core

type NameValue struct {
	isQuotedString bool
	separator      string
	quotes         string
	name           string
	value          interface{} //only accept nil and string
}

func NewNameValue(n string, v interface{}) *NameValue {
	if v != nil {
		if _, ok := v.(string); !ok {
			panic("value must be nil or string type")
		}
	}

	this := &NameValue{}

	this.name = n
	this.value = v
	this.separator = SIPSeparatorNames_EQUALS
	this.quotes = ""

	return this
}

/**
* Set the separator for the encoding method below.
 */
func (this *NameValue) SetSeparator(sep string) {
	this.separator = sep
}

/** A flag that indicates that doublequotes should be put around the
* value when encoded
*(for example name=value when value is doublequoted).
 */
func (this *NameValue) SetQuotedValue() {
	this.isQuotedString = true
	this.quotes = SIPSeparatorNames_DOUBLE_QUOTE
}

/** Return true if the value is quoted in doublequotes.
 */
func (this *NameValue) IsValueQuoted() bool {
	return this.isQuotedString
}

func (this *NameValue) GetName() string {
	return this.name
}

func (this *NameValue) GetValue() interface{} {
	return this.value
}

/**
* Set the name member
 */
func (this *NameValue) SetName(n string) {
	this.name = n
}

/**
* Set the value member
 */
func (this *NameValue) SetValue(v interface{}) {
	if v != nil {
		if _, ok := v.(string); !ok {
			panic("value must be nil or string type")
		}
	}
	this.value = v
}

/**
	* Get the encoded representation of this namevalue object.
        * Added doublequote for encoding doublequoted values
	* (bug reported by Kirby Kiem).
	*@since 1.0
	*@return an encoded name value (eg. name=value) string.
*/
func (this *NameValue) String() string {
	if this.name != "" && this.value != nil {
		return this.name + this.separator + this.quotes + this.value.(string) + this.quotes
	} else if this.name == "" && this.value != nil {
		return this.quotes + this.value.(string) + this.quotes
	} else if this.name != "" && this.value == nil {
		return this.name
	} else {
		return ""
	}
}

func (this *NameValue) Clone() interface{} {
	retval := &NameValue{}
	retval.separator = this.separator
	retval.isQuotedString = this.isQuotedString
	retval.quotes = this.quotes
	retval.name = this.name
	if this.value != nil {
		retval.value = this.value.(string)
	}
	return retval
}

/**
* Equality comparison predicate.
 */
/*public boolean equals( Object other) {
	if (! other.getClass().equals(this.getClass()))  return false;
        NameValue that = (NameValue) other;
	if (this == that) return true;
	if (this.name  == null && that.name != null ||
	   this.name != null && that.name == null) return false;
	if (this.name != null && that.name != null &&
		this.name.compareToIgnoreCase(that.name) != 0)
			return false;
	if ( this.value != null && that.value == null ||
	     this.value == null && that.value != null) return false;
	if (this.value == that.value) return true;
	if (value instanceof String) {
		// Quoted string comparisions are case sensitive.
	        if (isQuotedString)
			return this.value.equals(that.value);
		String val = (String) this.value;
		String val1 = (String) that.value;
		return val.compareToIgnoreCase(val1) == 0 ;
	} else return this.value.equals(that.value);
}*/
