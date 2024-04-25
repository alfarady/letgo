package entity

import (
	"github.com/alfarady/letgo/lib"
	"github.com/alfarady/letgo/x/api/shopee/commonentity"
)

//UpdateGlobalStockResult
type UpdateGlobalStockResult struct{
	commonentity.Result
	Warning string `json:"warning"`
}

//String
func(r UpdateGlobalStockResult)String()string{
	return lib.ObjectToString(r)
}