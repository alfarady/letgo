package entity

import (
	"github.com/alfarady/letgo/lib"
	"github.com/alfarady/letgo/x/api/shopee/commonentity"
)

//CategoryRecommendResult
type CategoryRecommendResult struct{
	commonentity.Result
	Warning string `json:"warning"`
	Response CategoryRecommendResultResponse `json:"response"` 
}

//String
func(c CategoryRecommendResult)String()string{
	return lib.ObjectToString(c)
}
//CategoryRecommendResultResponse
type CategoryRecommendResultResponse struct{
	CategoryID []int64 `json:"category_id"`
}

//String
func(c CategoryRecommendResultResponse)String()string{
	return lib.ObjectToString(c)
}