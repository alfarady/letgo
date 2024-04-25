package entity

import (
	"github.com/alfarady/letgo/lib"
	"github.com/alfarady/letgo/x/api/shopee/commonentity"
)

//UpdateGlobalSizeChartResult
type UpdateGlobalSizeChartResult struct{
	commonentity.Result
	Warning string `json:"warning"`
}

//String
func(r UpdateGlobalSizeChartResult)String()string{
	return lib.ObjectToString(r)
}