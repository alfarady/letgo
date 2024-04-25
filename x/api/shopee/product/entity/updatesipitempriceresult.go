package entity

import (
	"github.com/alfarady/letgo/lib"
	"github.com/alfarady/letgo/x/api/shopee/commonentity"
)

//UpdateSipItemPriceResult
type UpdateSipItemPriceResult struct{
	commonentity.Result
	Warning string `json:"warning"`
}

//String
func(r UpdateSipItemPriceResult)String()string{
	return lib.ObjectToString(r)
}