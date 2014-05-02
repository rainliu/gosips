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

func (this *AcceptLanguageList) GetFirst() Header {
	retval := this.SIPHeaderList.Front().Value.(Header) //super.getFirst();
	if retval != nil {
		return retval
	} else {
		return NewAcceptLanguage()
	}
}

func (this *AcceptLanguageList) GetLast() Header {
	retval := this.SIPHeaderList.Back().Value.(Header) //Header retval = super.getLast();
	if retval != nil {
		return retval
	} else {
		return NewAcceptLanguage()
	}
}
