package entity

import (
	"github.com/alfarady/letgo/lib"
)

//ParentBrandEntity
type ParentBrandEntity struct{
	ParentBrandID int64 `json:"parent_brand_id"`
}

//String
func(a ParentBrandEntity)String()string{
	return lib.ObjectToString(a)
}