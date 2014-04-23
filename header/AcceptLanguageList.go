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

func (this *AcceptLanguageList) GetFirst() SIPHeaderHeader {
	retval := this.SIPHeaderList.Front().Value.(SIPHeaderHeader) //super.getFirst();
	if retval != nil {
		return retval
	} else {
		return NewAcceptLanguage()
	}
}

func (this *AcceptLanguageList) GetLast() SIPHeaderHeader {
	retval := this.SIPHeaderList.Back().Value.(SIPHeaderHeader) //SIPHeaderHeader retval = super.getLast();
	if retval != nil {
		return retval
	} else {
		return NewAcceptLanguage()
	}
}
