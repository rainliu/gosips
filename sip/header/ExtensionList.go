package header

import "bytes"

type ExtensionHeaderLister interface {
	SIPHeaderLister
}

/**
* A generic extension header list.
 */

type ExtensionHeaderList struct {
	SIPHeaderList
}

func NewExtensionHeaderList(hName string) *ExtensionHeaderList {
	this := &ExtensionHeaderList{}
	this.SIPHeaderList.super(hName)
	return this
}

func (this *ExtensionHeaderList) String() string {
	var encoding bytes.Buffer
	for e := this.Front(); e != nil; e = e.Next() {
		eh := e.Value.(*Extension)
		encoding.WriteString(eh.String())
	}
	return encoding.String()
}
