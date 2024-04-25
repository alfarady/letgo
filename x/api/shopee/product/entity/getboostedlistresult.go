package entity

import (
	"github.com/alfarady/letgo/lib"
	"github.com/alfarady/letgo/x/api/shopee/commonentity"
)

//GetBoostedListResult
type GetBoostedListResult struct{
	commonentity.Result
	Response GetBoostedListResultResponse `json:"response"`
	Warning string `json:"warning"`
}

//String
func(g GetBoostedListResult)String()string{
	return lib.ObjectToString(g)
}
//GetBoostedListResultResponse
type GetBoostedListResultResponse struct{
	ItemList []GetBoostedListItemListEntity `json:"item_list"`
}

//String
func(g GetBoostedListResultResponse)String()string{
	return lib.ObjectToString(g)
}