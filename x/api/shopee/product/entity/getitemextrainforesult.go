package entity

import (
	"github.com/alfarady/letgo/lib"
	"github.com/alfarady/letgo/x/api/shopee/commonentity"
)

//GetItemExtraInfoResult
type GetItemExtraInfoResult struct{
	commonentity.Result
	Response GetItemExtraInfoResultResponse `json:"response"`
	Warning string `json:"warning"`
}

//String
func(g GetItemExtraInfoResult)String()string{
	return lib.ObjectToString(g)
}

//GetItemExtraInfoResultResponse
type GetItemExtraInfoResultResponse struct{
	ItemList []ItemExtraEntity `json:"item_list"`
}

//String
func(g GetItemExtraInfoResultResponse)String()string{
	return lib.ObjectToString(g)
}