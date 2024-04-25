package entity

import (
	"github.com/alfarady/letgo/lib"
	"github.com/alfarady/letgo/x/api/shopee/commonentity"
)

//UpdateTierVariationResult
type UpdateTierVariationResult struct{
	commonentity.Result
	Warning string `json:"warning"`
}

//String
func(r UpdateTierVariationResult)String()string{
	return lib.ObjectToString(r)
}