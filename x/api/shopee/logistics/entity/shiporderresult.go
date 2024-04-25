package entity

import(
	"github.com/alfarady/letgo/x/api/shopee/commonentity"
	"github.com/alfarady/letgo/lib"
)

//ShipOrderResult
type ShipOrderResult struct{
	commonentity.Result
}

//String
func(s ShipOrderResult)String()string{
	return lib.ObjectToString(s)
}