package entity

import (
	"github.com/alfarady/letgo/lib"
	"github.com/alfarady/letgo/x/api/shopee/commonentity"
)

//SupportSizeChartResult
type SupportGlobalSizeChartResult struct{
	commonentity.Result
	Warning string `json:"warning"`
	Response SupportSizeChartResultResponse `json:"response"`
}

//String
func(r SupportGlobalSizeChartResult)String()string{
	return lib.ObjectToString(r)
}
//SupportSizeChartResultResponse
type SupportSizeChartResultResponse struct{
	SupportSizeChart bool `json:"support_size_chart"`
}

//String
func(r SupportSizeChartResultResponse)String()string{
	return lib.ObjectToString(r)
}