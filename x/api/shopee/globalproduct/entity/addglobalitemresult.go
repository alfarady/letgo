package entity

import (
	"github.com/alfarady/letgo/lib"
	"github.com/alfarady/letgo/x/api/shopee/commonentity"
)

//AddGlobalItemResult
type AddGlobalItemResult struct{
	commonentity.Result
	Warning string `json:"warning"`
	Response AddGlobalItemResultResponse `json:"response"`
}

//String
func(g AddGlobalItemResult)String()string{
	return lib.ObjectToString(g)
}
//AddGlobalItemResultResponse
type AddGlobalItemResultResponse struct{
	GlobalItemID int64 `json:"global_item_id"`
}
//String
func(g AddGlobalItemResultResponse)String()string{
	return lib.ObjectToString(g)
}