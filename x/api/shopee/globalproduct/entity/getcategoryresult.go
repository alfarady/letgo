package entity

import (
	"github.com/alfarady/letgo/lib"
	"github.com/alfarady/letgo/x/api/shopee/commonentity"
)

//GetCategoryResult
type GetCategoryResult struct{
	commonentity.Result
	Warning string `json:"warning"`
	Response GetCategoryResultResponse `json:"response"`
}

//String
func(g GetCategoryResult)String()string{
	return lib.ObjectToString(g)
}
//GetCategoryResultResponse
type GetCategoryResultResponse struct{
	CategoryList []CategoryEntity `json:"category_list"`
}

//String
func(g GetCategoryResultResponse)String()string{
	return lib.ObjectToString(g)
}