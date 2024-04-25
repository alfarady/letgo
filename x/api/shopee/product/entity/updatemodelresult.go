package entity

import (
	"github.com/alfarady/letgo/lib"
	"github.com/alfarady/letgo/x/api/shopee/commonentity"
)


//UpdateModelResult
type UpdateModelResult struct{
	commonentity.Result
	Warning string `json:"warning"`
}

//String
func(g UpdateModelResult)String()string{
	return lib.ObjectToString(g)
}