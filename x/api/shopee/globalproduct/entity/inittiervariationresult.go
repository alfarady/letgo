package entity

import (
	"github.com/alfarady/letgo/lib"
	"github.com/alfarady/letgo/x/api/shopee/commonentity"
)

//InitTierVariationResult
type InitTierVariationResult struct{
	commonentity.Result
	Warning string `json:"warning"`
}

//String
func(g InitTierVariationResult)String()string{
	return lib.ObjectToString(g)
}