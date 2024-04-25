package entity

import (
	"github.com/alfarady/letgo/lib"
	"github.com/alfarady/letgo/x/api/shopee/commonentity"
)


//UpdateGlobalPriceResult
type UpdateGlobalPriceResult struct{
	commonentity.Result
	Warning string `json:"warning"`
}

//String
func(g UpdateGlobalPriceResult)String()string{
	return lib.ObjectToString(g)
}