package entity

import(
	"github.com/alfarady/letgo/x/api/shopee/commonentity"
	"github.com/alfarady/letgo/lib"
)

//UnSplitOrderResult
type AddInvoiceDataResult struct{
	commonentity.Result
}

//String
func(a AddInvoiceDataResult)String()string{
	return lib.ObjectToString(a)
}