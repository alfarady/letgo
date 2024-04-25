package entity

import(
	"github.com/alfarady/letgo/x/api/shopee/commonentity"
	"github.com/alfarady/letgo/lib"
)

//HandleBuyerCancellationResult
type HandleBuyerCancellationResult struct{
	commonentity.Result
}

//String
func(h HandleBuyerCancellationResult)String()string{
	return lib.ObjectToString(h)
}