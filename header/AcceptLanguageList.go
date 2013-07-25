package header

import (
	"gosip/core"
)

/**
* AcceptLanguageList: Strings together a list of AcceptLanguage SIPHeaders.
 */
type AcceptLanguageList struct {
	SIPHeaderList
}

func NewAcceptLanguageList() *AcceptLanguageList {
	this := &AcceptLanguageList{}
	this.SIPHeaderList.super(core.SIPHeaderNames_ACCEPT_LANGUAGE)
	return this
}

func (this *AcceptLanguageList) GetFirst() SIPHeader {
	retval := this.SIPHeaderList.Front().Value.(SIPHeader) //super.getFirst();
	if retval != nil {
		return retval
	} else {
		return NewAcceptLanguage()
	}
}

func (this *AcceptLanguageList) GetLast() SIPHeader {
	retval := this.SIPHeaderList.Back().Value.(SIPHeader) //SIPHeader retval = super.getLast();
	if retval != nil {
		return retval
	} else {
		return NewAcceptLanguage()
	}
}
