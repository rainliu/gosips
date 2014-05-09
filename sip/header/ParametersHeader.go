package header

import (
	"container/list"
	"gosips/core"
)

type ParametersHeader interface {
	Header
	GetParameter(name string) string
	GetParameterValue(name string) string //interface{}
	GetParameterNames() *list.List
	HasParameters() bool
	RemoveParameter(name string)
	SetParameter(name, value string) (ParseException error)
	SetQuotedParameter(name, value string)
	HasParameter(parameterName string) bool
	RemoveParameters()
	GetParameters() *core.NameValueList
	SetParameters(parameters *core.NameValueList)
	GetNameValue(parameterName string) *core.NameValue
	//EncodeBody() string
}
