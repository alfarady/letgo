package entity

import (
	"github.com/alfarady/letgo/lib"
	"github.com/alfarady/letgo/x/api/shopee/commonentity"
)


//UpdateGlobalModelResult
type UpdateGlobalModelResult struct{
	commonentity.Result
	Warning string `json:"warning"`
}

//String
func(g UpdateGlobalModelResult)String()string{
	return lib.ObjectToString(g)
}