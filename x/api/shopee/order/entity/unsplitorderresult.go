package entity

import(
	"github.com/alfarady/letgo/x/api/shopee/commonentity"
	"github.com/alfarady/letgo/lib"
)

//UnSplitOrderResult
type UnSplitOrderResult struct{
	commonentity.Result
	Response OrderUpdateTimeResponse `json:"response"`
}

//String
func(u UnSplitOrderResult)String()string{
	return lib.ObjectToString(u)
}