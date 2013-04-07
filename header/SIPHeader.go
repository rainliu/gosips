package header

import (

)

type SIPHeader interface{
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