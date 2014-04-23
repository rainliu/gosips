package header

import ()

type SIPHeaderHeader interface {
	EncodeBody() string
	GetHeaderName() string
	GetHeaderValue() string
	GetName() string
	GetValue() string
	IsHeaderList() bool
	SetHeaderName(hdrname string)
	String() string
	Clone() interface{}
}
