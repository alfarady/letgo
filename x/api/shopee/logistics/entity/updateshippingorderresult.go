package entity

import(
	"github.com/alfarady/letgo/x/api/shopee/commonentity"
	"github.com/alfarady/letgo/lib"
)

//UpdateShippingOrderResult
type UpdateShippingOrderResult struct{
	commonentity.Result
}

//String
func(u UpdateShippingOrderResult)String()string{
	return lib.ObjectToString(u)
}