package entity

import (
	"github.com/alfarady/letgo/lib"
	"github.com/alfarady/letgo/x/api/shopee/commonentity"
)

//AddGlobalModelResult
type AddGlobalModelResult struct{
	commonentity.Result
	Warning string `json:"warning"`
}

//String
func(g AddGlobalModelResult)String()string{
	return lib.ObjectToString(g)
}