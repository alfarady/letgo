package entity

import (
	"github.com/alfarady/letgo/lib"
	"github.com/alfarady/letgo/x/api/shopee/commonentity"
)

//AddItemResult
type AddItemResult struct{
	commonentity.Result
	Warning string `json:"warning"`
	ItemDangerous int `json:"item_dangerous"`
	Response ItemEntity `json:"response"`
}

//String
func(g AddItemResult)String()string{
	return lib.ObjectToString(g)
}
